package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	binrpc "github.com/florentchauveau/go-kamailio-binrpc/v3"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	URI     string
	Timeout time.Duration

	url   *url.URL
	mutex sync.Mutex
	conn  net.Conn

	up            prometheus.Gauge
	failedScrapes prometheus.Counter
	totalScrapes  prometheus.Counter
}

const (
	namespace = "kamailio"
)

var (
	metricNamePattern = regexp.MustCompile("[^a-zA-Z0-9]")
)

func NewCollector(uri string, timeout time.Duration) (*Collector, error) {
	collector := Collector{}

	collector.URI = uri
	collector.Timeout = timeout

	var url *url.URL
	var err error

	if url, err = url.Parse(collector.URI); err != nil {
		return nil, fmt.Errorf("cannot parse URI: %w", err)
	}

	collector.url = url

	collector.up = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "up",
		Help:      "Was the last scrape successful.",
	})

	collector.totalScrapes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "exporter_total_scrapes",
		Help:      "Number of total kamailio scrapes",
	})

	collector.failedScrapes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "exporter_failed_scrapes",
		Help:      "Number of failed kamailio scrapes",
	})

	return &collector, nil
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.scrape(ch)

	if err != nil {
		c.failedScrapes.Inc()
		c.up.Set(0)
		log.Println("[error] Failed to scrape metrics:", err)
	} else {
		c.up.Set(1)
	}

	ch <- c.up
	ch <- c.totalScrapes
	ch <- c.failedScrapes
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *Collector) fetchBINRPC() ([]binrpc.Record, error) {
	cookie, err := binrpc.WritePacket(c.conn, "stats.fetch", "all")
	if err != nil {
		return nil, err
	}

	records, err := binrpc.ReadPacket(c.conn, cookie)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (c *Collector) scrape(ch chan<- prometheus.Metric) error {
	c.totalScrapes.Inc()

	var err error

	address := c.url.Host
	if c.url.Scheme == "unix" {
		address = c.url.Path
	}

	c.conn, err = net.DialTimeout(c.url.Scheme, address, c.Timeout)
	if err != nil {
		return fmt.Errorf(`Connection timed out:  %s`, err)
	}

	c.conn.SetDeadline(time.Now().Add(c.Timeout))

	defer c.conn.Close()

	records, err := c.fetchBINRPC()
	if err != nil {
		return fmt.Errorf(`BINRPC fetch error:  %s`, err)
	}

	// We expect just 1 record of type map.
	if len(records) == 2 && records[0].Type == binrpc.TypeInt && records[0].Value.(int) == 500 {
		return fmt.Errorf(`invalid response: [500] %s`, records[1].Value.(string))
	} else if len(records) != 1 {
		return fmt.Errorf(`invalid response, expected %d record, got %d`, 1, len(records))
	}

	items, err := records[0].StructItems()
	if err != nil {
		return fmt.Errorf(`BINRPC StructItems not as expected:  %s`, err)
	}

	for _, item := range items {
		rawMetricName := item.Key
		normalizedMetricName := namespace + "_" + normalizeMetricName(rawMetricName)
		metricString, _ := item.Value.String()
		metricValue, _ := strconv.Atoi(metricString)
		metric, err := prometheus.NewConstMetric(
			prometheus.NewDesc(normalizedMetricName, rawMetricName, []string{}, nil),
			guessMetricValueTypeForName(rawMetricName),
			float64(metricValue),
			[]string{}...,
		)
		if err != nil {
			return fmt.Errorf(`Unable to construct prometheus metric:  %s`, err)
		}

		ch <- metric
	}

	return nil
}

func guessMetricValueTypeForName(name string) prometheus.ValueType {
	valueTypeByMetricName := map[string]prometheus.ValueType{
		"core.bad_msg_hdr":                             prometheus.CounterValue,
		"core.bad_URIs_rcvd":                           prometheus.CounterValue,
		"core.drop_replies":                            prometheus.CounterValue,
		"core.drop_requests":                           prometheus.CounterValue,
		"core.err_replies":                             prometheus.CounterValue,
		"core.err_requests":                            prometheus.CounterValue,
		"core.fwd_replies":                             prometheus.CounterValue,
		"core.fwd_requests":                            prometheus.CounterValue,
		"core.rcv_replies_18x":                         prometheus.CounterValue,
		"core.rcv_replies_1xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_1xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_1xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_1xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_1xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_1xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_1xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_1xx":                         prometheus.CounterValue,
		"core.rcv_replies_1xx_update":                  prometheus.CounterValue,
		"core.rcv_replies_2xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_2xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_2xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_2xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_2xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_2xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_2xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_2xx":                         prometheus.CounterValue,
		"core.rcv_replies_2xx_update":                  prometheus.CounterValue,
		"core.rcv_replies_3xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_3xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_3xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_3xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_3xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_3xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_3xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_3xx":                         prometheus.CounterValue,
		"core.rcv_replies_3xx_update":                  prometheus.CounterValue,
		"core.rcv_replies_401":                         prometheus.CounterValue,
		"core.rcv_replies_404":                         prometheus.CounterValue,
		"core.rcv_replies_407":                         prometheus.CounterValue,
		"core.rcv_replies_480":                         prometheus.CounterValue,
		"core.rcv_replies_486":                         prometheus.CounterValue,
		"core.rcv_replies_4xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_4xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_4xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_4xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_4xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_4xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_4xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_4xx":                         prometheus.CounterValue,
		"core.rcv_replies_4xx_update":                  prometheus.CounterValue,
		"core.rcv_replies_5xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_5xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_5xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_5xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_5xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_5xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_5xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_5xx":                         prometheus.CounterValue,
		"core.rcv_replies_5xx_update":                  prometheus.CounterValue,
		"core.rcv_replies_6xx_bye":                     prometheus.CounterValue,
		"core.rcv_replies_6xx_cancel":                  prometheus.CounterValue,
		"core.rcv_replies_6xx_invite":                  prometheus.CounterValue,
		"core.rcv_replies_6xx_message":                 prometheus.CounterValue,
		"core.rcv_replies_6xx_prack":                   prometheus.CounterValue,
		"core.rcv_replies_6xx_refer":                   prometheus.CounterValue,
		"core.rcv_replies_6xx_reg":                     prometheus.CounterValue,
		"core.rcv_replies_6xx":                         prometheus.CounterValue,
		"core.rcv_replies_6xx_update":                  prometheus.CounterValue,
		"core.rcv_replies":                             prometheus.CounterValue,
		"core.rcv_requests_ack":                        prometheus.CounterValue,
		"core.rcv_requests_bye":                        prometheus.CounterValue,
		"core.rcv_requests_cancel":                     prometheus.CounterValue,
		"core.rcv_requests_info":                       prometheus.CounterValue,
		"core.rcv_requests_invite":                     prometheus.CounterValue,
		"core.rcv_requests_message":                    prometheus.CounterValue,
		"core.rcv_requests_notify":                     prometheus.CounterValue,
		"core.rcv_requests_options":                    prometheus.CounterValue,
		"core.rcv_requests_prack":                      prometheus.CounterValue,
		"core.rcv_requests_publish":                    prometheus.CounterValue,
		"core.rcv_requests_refer":                      prometheus.CounterValue,
		"core.rcv_requests_register":                   prometheus.CounterValue,
		"core.rcv_requests":                            prometheus.CounterValue,
		"core.rcv_requests_subscribe":                  prometheus.CounterValue,
		"core.rcv_requests_update":                     prometheus.CounterValue,
		"core.unsupported_methods":                     prometheus.CounterValue,
		"dialog.active_dialogs":                        prometheus.GaugeValue,
		"dialog.early_dialogs":                         prometheus.GaugeValue,
		"dialog.expired_dialogs":                       prometheus.CounterValue,
		"dialog.failed_dialogs":                        prometheus.CounterValue,
		"dialog.processed_dialogs":                     prometheus.CounterValue,
		"dns.failed_dns_request":                       prometheus.CounterValue,
		"dns.slow_dns_request":                         prometheus.CounterValue,
		"mysql.driver_errors":                          prometheus.CounterValue,
		"pike.blocked_ips":                             prometheus.CounterValue,
		"registrar.accepted_regs":                      prometheus.CounterValue,
		"registrar.default_expire":                     prometheus.CounterValue,
		"registrar.default_expires_range":              prometheus.CounterValue,
		"registrar.expires_range":                      prometheus.CounterValue,
		"registrar.max_contacts":                       prometheus.CounterValue,
		"registrar.max_expires":                        prometheus.CounterValue,
		"registrar.rejected_regs":                      prometheus.CounterValue,
		"shmem.fragments":                              prometheus.GaugeValue,
		"shmem.free_size":                              prometheus.GaugeValue,
		"shmem.max_used_size":                          prometheus.GaugeValue,
		"shmem.real_used_size":                         prometheus.GaugeValue,
		"shmem.total_size":                             prometheus.GaugeValue,
		"shmem.used_size":                              prometheus.GaugeValue,
		"siptrace.traced_replies":                      prometheus.CounterValue,
		"siptrace.traced_requests":                     prometheus.CounterValue,
		"sl.1xx_replies":                               prometheus.CounterValue,
		"sl.200_replies":                               prometheus.CounterValue,
		"sl.202_replies":                               prometheus.CounterValue,
		"sl.2xx_replies":                               prometheus.CounterValue,
		"sl.300_replies":                               prometheus.CounterValue,
		"sl.301_replies":                               prometheus.CounterValue,
		"sl.302_replies":                               prometheus.CounterValue,
		"sl.3xx_replies":                               prometheus.CounterValue,
		"sl.400_replies":                               prometheus.CounterValue,
		"sl.401_replies":                               prometheus.CounterValue,
		"sl.403_replies":                               prometheus.CounterValue,
		"sl.404_replies":                               prometheus.CounterValue,
		"sl.407_replies":                               prometheus.CounterValue,
		"sl.408_replies":                               prometheus.CounterValue,
		"sl.483_replies":                               prometheus.CounterValue,
		"sl.4xx_replies":                               prometheus.CounterValue,
		"sl.500_replies":                               prometheus.CounterValue,
		"sl.5xx_replies":                               prometheus.CounterValue,
		"sl.6xx_replies":                               prometheus.CounterValue,
		"sl.failures":                                  prometheus.CounterValue,
		"sl.received_ACKs":                             prometheus.CounterValue,
		"sl.sent_err_replies":                          prometheus.CounterValue,
		"sl.sent_replies":                              prometheus.CounterValue,
		"sl.xxx_replies":                               prometheus.CounterValue,
		"tcp.connect_failed":                           prometheus.CounterValue,
		"tcp.connect_success":                          prometheus.CounterValue,
		"tcp.con_reset":                                prometheus.CounterValue,
		"tcp.con_timeout":                              prometheus.CounterValue,
		"tcp.current_opened_connections":               prometheus.GaugeValue,
		"tcp.current_write_queue_size":                 prometheus.GaugeValue,
		"tcp.established":                              prometheus.CounterValue,
		"tcp.local_reject":                             prometheus.CounterValue,
		"tcp.passive_open":                             prometheus.CounterValue,
		"tcp.sendq_full":                               prometheus.CounterValue,
		"tcp.send_timeout":                             prometheus.CounterValue,
		"tmx.2xx_transactions":                         prometheus.CounterValue,
		"tmx.3xx_transactions":                         prometheus.CounterValue,
		"tmx.4xx_transactions":                         prometheus.CounterValue,
		"tmx.5xx_transactions":                         prometheus.CounterValue,
		"tmx.6xx_transactions":                         prometheus.CounterValue,
		"tmx.active_transactions":                      prometheus.GaugeValue,
		"tmx.inuse_transactions":                       prometheus.GaugeValue,
		"tmx.rpl_absorbed":                             prometheus.CounterValue,
		"tmx.rpl_generated":                            prometheus.CounterValue,
		"tmx.rpl_received":                             prometheus.CounterValue,
		"tmx.rpl_relayed":                              prometheus.CounterValue,
		"tmx.rpl_sent":                                 prometheus.CounterValue,
		"tmx.UAC_transactions":                         prometheus.CounterValue,
		"tmx.UAS_transactions":                         prometheus.CounterValue,
		"tsilo.added_branches":                         prometheus.CounterValue,
		"tsilo.stored_ruris":                           prometheus.GaugeValue,
		"tsilo.stored_transactions":                    prometheus.GaugeValue,
		"tsilo.total_ruris":                            prometheus.CounterValue,
		"tsilo.total_transactions":                     prometheus.CounterValue,
		"usrloc.location_contacts":                     prometheus.GaugeValue,
		"usrloc.location_expires":                      prometheus.CounterValue,
		"usrloc.location_users":                        prometheus.GaugeValue,
		"usrloc.registered_users":                      prometheus.GaugeValue,
		"websocket.ws_current_connections":             prometheus.GaugeValue,
		"websocket.ws_failed_connections":              prometheus.CounterValue,
		"websocket.ws_failed_handshakes":               prometheus.CounterValue,
		"websocket.ws_local_closed_connections":        prometheus.CounterValue,
		"websocket.ws_max_concurrent_connections":      prometheus.CounterValue,
		"websocket.ws_msrp_current_connections":        prometheus.GaugeValue,
		"websocket.ws_msrp_failed_connections":         prometheus.CounterValue,
		"websocket.ws_msrp_local_closed_connections":   prometheus.CounterValue,
		"websocket.ws_msrp_max_concurrent_connections": prometheus.CounterValue,
		"websocket.ws_msrp_received_frames":            prometheus.CounterValue,
		"websocket.ws_msrp_remote_closed_connections":  prometheus.CounterValue,
		"websocket.ws_msrp_successful_handshakes":      prometheus.CounterValue,
		"websocket.ws_msrp_transmitted_frames":         prometheus.CounterValue,
		"websocket.ws_received_frames":                 prometheus.CounterValue,
		"websocket.ws_remote_closed_connections":       prometheus.CounterValue,
		"websocket.ws_sip_current_connections":         prometheus.GaugeValue,
		"websocket.ws_sip_failed_connections":          prometheus.CounterValue,
		"websocket.ws_sip_local_closed_connections":    prometheus.CounterValue,
		"websocket.ws_sip_max_concurrent_connections":  prometheus.CounterValue,
		"websocket.ws_sip_received_frames":             prometheus.CounterValue,
		"websocket.ws_sip_remote_closed_connections":   prometheus.CounterValue,
		"websocket.ws_sip_successful_handshakes":       prometheus.CounterValue,
		"websocket.ws_sip_transmitted_frames":          prometheus.CounterValue,
		"websocket.ws_successful_handshakes":           prometheus.CounterValue,
		"websocket.ws_transmitted_frames":              prometheus.CounterValue,
	}

	value, ok := valueTypeByMetricName[name]
	if ok {
		return value
	} else {
		// Make a guess!
		if strings.HasSuffix(name, "_bytes") || strings.HasSuffix(name, "_counter") ||
			strings.HasSuffix(name, "_seconds") || strings.HasSuffix(name, "_total") {
			return prometheus.CounterValue
		} else {
			// Defaulting to a "gauge" seems like it would be the lesser evil. It's easy to spot
			// a metric continuously growing as opposed to one the moves up and down but at an
			// unexpected scale.
			return prometheus.GaugeValue
		}
	}
}

func normalizeMetricName(metricName string) string {
	// Replace invalid characters with underscores.
	return metricNamePattern.ReplaceAllString(metricName, "_")
}
