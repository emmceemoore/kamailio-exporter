# kamailio-exporter

An opinionated [Kamailio](https://www.kamailio.org/) exporter for [Prometheus](https://prometheus.io/).

## Prerequisites

**Kamailio**:
- You must be on a recent version of Kamailio (`5.1+`).
- You must have the [`ctl`](https://www.kamailio.net/docs/modules/stable/modules/ctl.html) module loaded.
  - The default settings should work; if you change anything (e.g. - `binrpc`, etc.) all bets are off!
- The exporter is expected to be running on the same host as `kamailio`.
**System Dependencies**:
- `git` and `go`
  - On a Debian-based machine: `apt-get install git golang-1.17`

## Build/Run

**Build:**
```
git clone ...
cd ./kamailio-exporter
go build -o ./kamailio-exporter
```

**Run:**
```
./kamailio-exporter &
curl localhost:9494/metrics
```

_In practice, you'll likely want to run the exporter using something like [systemd](https://systemd.io/)._

## Exported Metrics

The exported metrics reflect the values returned when you run `kamcmd stats.fetch all`. This will
depend on which modules you have loaded.

Example:
```
# HELP kamailio_core_bad_URIs_rcvd core.bad_URIs_rcvd
# TYPE kamailio_core_bad_URIs_rcvd counter
kamailio_core_bad_URIs_rcvd 0
# HELP kamailio_core_bad_msg_hdr core.bad_msg_hdr
# TYPE kamailio_core_bad_msg_hdr counter
kamailio_core_bad_msg_hdr 0
# HELP kamailio_core_drop_replies core.drop_replies
# TYPE kamailio_core_drop_replies counter
kamailio_core_drop_replies 0
# HELP kamailio_core_drop_requests core.drop_requests
# TYPE kamailio_core_drop_requests counter
kamailio_core_drop_requests 0
# HELP kamailio_core_err_replies core.err_replies
# TYPE kamailio_core_err_replies counter
kamailio_core_err_replies 0
# HELP kamailio_core_err_requests core.err_requests
# TYPE kamailio_core_err_requests counter
kamailio_core_err_requests 0
# HELP kamailio_core_fwd_replies core.fwd_replies
# TYPE kamailio_core_fwd_replies counter
kamailio_core_fwd_replies 0
# HELP kamailio_core_fwd_requests core.fwd_requests
# TYPE kamailio_core_fwd_requests counter
kamailio_core_fwd_requests 0
# HELP kamailio_core_rcv_replies core.rcv_replies
# TYPE kamailio_core_rcv_replies counter
kamailio_core_rcv_replies 693
# HELP kamailio_core_rcv_replies_18x core.rcv_replies_18x
# TYPE kamailio_core_rcv_replies_18x counter
kamailio_core_rcv_replies_18x 0
# HELP kamailio_core_rcv_replies_1xx core.rcv_replies_1xx
# TYPE kamailio_core_rcv_replies_1xx counter
kamailio_core_rcv_replies_1xx 0
# HELP kamailio_core_rcv_replies_1xx_bye core.rcv_replies_1xx_bye
# TYPE kamailio_core_rcv_replies_1xx_bye counter
kamailio_core_rcv_replies_1xx_bye 0
# HELP kamailio_core_rcv_replies_1xx_cancel core.rcv_replies_1xx_cancel
# TYPE kamailio_core_rcv_replies_1xx_cancel counter
kamailio_core_rcv_replies_1xx_cancel 0
# HELP kamailio_core_rcv_replies_1xx_invite core.rcv_replies_1xx_invite
# TYPE kamailio_core_rcv_replies_1xx_invite counter
kamailio_core_rcv_replies_1xx_invite 0
# HELP kamailio_core_rcv_replies_1xx_message core.rcv_replies_1xx_message
# TYPE kamailio_core_rcv_replies_1xx_message counter
kamailio_core_rcv_replies_1xx_message 0
# HELP kamailio_core_rcv_replies_1xx_prack core.rcv_replies_1xx_prack
# TYPE kamailio_core_rcv_replies_1xx_prack counter
kamailio_core_rcv_replies_1xx_prack 0
# HELP kamailio_core_rcv_replies_1xx_refer core.rcv_replies_1xx_refer
# TYPE kamailio_core_rcv_replies_1xx_refer counter
kamailio_core_rcv_replies_1xx_refer 0
# HELP kamailio_core_rcv_replies_1xx_reg core.rcv_replies_1xx_reg
# TYPE kamailio_core_rcv_replies_1xx_reg counter
kamailio_core_rcv_replies_1xx_reg 0
# HELP kamailio_core_rcv_replies_1xx_update core.rcv_replies_1xx_update
# TYPE kamailio_core_rcv_replies_1xx_update counter
kamailio_core_rcv_replies_1xx_update 0
# HELP kamailio_core_rcv_replies_2xx core.rcv_replies_2xx
# TYPE kamailio_core_rcv_replies_2xx counter
kamailio_core_rcv_replies_2xx 662
# HELP kamailio_core_rcv_replies_2xx_bye core.rcv_replies_2xx_bye
# TYPE kamailio_core_rcv_replies_2xx_bye counter
kamailio_core_rcv_replies_2xx_bye 0
# HELP kamailio_core_rcv_replies_2xx_cancel core.rcv_replies_2xx_cancel
# TYPE kamailio_core_rcv_replies_2xx_cancel counter
kamailio_core_rcv_replies_2xx_cancel 0
# HELP kamailio_core_rcv_replies_2xx_invite core.rcv_replies_2xx_invite
# TYPE kamailio_core_rcv_replies_2xx_invite counter
kamailio_core_rcv_replies_2xx_invite 0
# HELP kamailio_core_rcv_replies_2xx_message core.rcv_replies_2xx_message
# TYPE kamailio_core_rcv_replies_2xx_message counter
kamailio_core_rcv_replies_2xx_message 0
# HELP kamailio_core_rcv_replies_2xx_prack core.rcv_replies_2xx_prack
# TYPE kamailio_core_rcv_replies_2xx_prack counter
kamailio_core_rcv_replies_2xx_prack 0
# HELP kamailio_core_rcv_replies_2xx_refer core.rcv_replies_2xx_refer
# TYPE kamailio_core_rcv_replies_2xx_refer counter
kamailio_core_rcv_replies_2xx_refer 0
# HELP kamailio_core_rcv_replies_2xx_reg core.rcv_replies_2xx_reg
# TYPE kamailio_core_rcv_replies_2xx_reg counter
kamailio_core_rcv_replies_2xx_reg 0
# HELP kamailio_core_rcv_replies_2xx_update core.rcv_replies_2xx_update
# TYPE kamailio_core_rcv_replies_2xx_update counter
kamailio_core_rcv_replies_2xx_update 0
# HELP kamailio_core_rcv_replies_3xx core.rcv_replies_3xx
# TYPE kamailio_core_rcv_replies_3xx counter
kamailio_core_rcv_replies_3xx 0
# HELP kamailio_core_rcv_replies_3xx_bye core.rcv_replies_3xx_bye
# TYPE kamailio_core_rcv_replies_3xx_bye counter
kamailio_core_rcv_replies_3xx_bye 0
# HELP kamailio_core_rcv_replies_3xx_cancel core.rcv_replies_3xx_cancel
# TYPE kamailio_core_rcv_replies_3xx_cancel counter
kamailio_core_rcv_replies_3xx_cancel 0
# HELP kamailio_core_rcv_replies_3xx_invite core.rcv_replies_3xx_invite
# TYPE kamailio_core_rcv_replies_3xx_invite counter
kamailio_core_rcv_replies_3xx_invite 0
# HELP kamailio_core_rcv_replies_3xx_message core.rcv_replies_3xx_message
# TYPE kamailio_core_rcv_replies_3xx_message counter
kamailio_core_rcv_replies_3xx_message 0
# HELP kamailio_core_rcv_replies_3xx_prack core.rcv_replies_3xx_prack
# TYPE kamailio_core_rcv_replies_3xx_prack counter
kamailio_core_rcv_replies_3xx_prack 0
# HELP kamailio_core_rcv_replies_3xx_refer core.rcv_replies_3xx_refer
# TYPE kamailio_core_rcv_replies_3xx_refer counter
kamailio_core_rcv_replies_3xx_refer 0
# HELP kamailio_core_rcv_replies_3xx_reg core.rcv_replies_3xx_reg
# TYPE kamailio_core_rcv_replies_3xx_reg counter
kamailio_core_rcv_replies_3xx_reg 0
# HELP kamailio_core_rcv_replies_3xx_update core.rcv_replies_3xx_update
# TYPE kamailio_core_rcv_replies_3xx_update counter
kamailio_core_rcv_replies_3xx_update 0
# HELP kamailio_core_rcv_replies_401 core.rcv_replies_401
# TYPE kamailio_core_rcv_replies_401 counter
kamailio_core_rcv_replies_401 0
# HELP kamailio_core_rcv_replies_404 core.rcv_replies_404
# TYPE kamailio_core_rcv_replies_404 counter
kamailio_core_rcv_replies_404 0
# HELP kamailio_core_rcv_replies_407 core.rcv_replies_407
# TYPE kamailio_core_rcv_replies_407 counter
kamailio_core_rcv_replies_407 0
# HELP kamailio_core_rcv_replies_480 core.rcv_replies_480
# TYPE kamailio_core_rcv_replies_480 counter
kamailio_core_rcv_replies_480 0
# HELP kamailio_core_rcv_replies_486 core.rcv_replies_486
# TYPE kamailio_core_rcv_replies_486 counter
kamailio_core_rcv_replies_486 0
# HELP kamailio_core_rcv_replies_4xx core.rcv_replies_4xx
# TYPE kamailio_core_rcv_replies_4xx counter
kamailio_core_rcv_replies_4xx 8
# HELP kamailio_core_rcv_replies_4xx_bye core.rcv_replies_4xx_bye
# TYPE kamailio_core_rcv_replies_4xx_bye counter
kamailio_core_rcv_replies_4xx_bye 0
# HELP kamailio_core_rcv_replies_4xx_cancel core.rcv_replies_4xx_cancel
# TYPE kamailio_core_rcv_replies_4xx_cancel counter
kamailio_core_rcv_replies_4xx_cancel 0
# HELP kamailio_core_rcv_replies_4xx_invite core.rcv_replies_4xx_invite
# TYPE kamailio_core_rcv_replies_4xx_invite counter
kamailio_core_rcv_replies_4xx_invite 0
# HELP kamailio_core_rcv_replies_4xx_message core.rcv_replies_4xx_message
# TYPE kamailio_core_rcv_replies_4xx_message counter
kamailio_core_rcv_replies_4xx_message 0
# HELP kamailio_core_rcv_replies_4xx_prack core.rcv_replies_4xx_prack
# TYPE kamailio_core_rcv_replies_4xx_prack counter
kamailio_core_rcv_replies_4xx_prack 0
# HELP kamailio_core_rcv_replies_4xx_refer core.rcv_replies_4xx_refer
# TYPE kamailio_core_rcv_replies_4xx_refer counter
kamailio_core_rcv_replies_4xx_refer 0
# HELP kamailio_core_rcv_replies_4xx_reg core.rcv_replies_4xx_reg
# TYPE kamailio_core_rcv_replies_4xx_reg counter
kamailio_core_rcv_replies_4xx_reg 0
# HELP kamailio_core_rcv_replies_4xx_update core.rcv_replies_4xx_update
# TYPE kamailio_core_rcv_replies_4xx_update counter
kamailio_core_rcv_replies_4xx_update 0
# HELP kamailio_core_rcv_replies_5xx core.rcv_replies_5xx
# TYPE kamailio_core_rcv_replies_5xx counter
kamailio_core_rcv_replies_5xx 23
# HELP kamailio_core_rcv_replies_5xx_bye core.rcv_replies_5xx_bye
# TYPE kamailio_core_rcv_replies_5xx_bye counter
kamailio_core_rcv_replies_5xx_bye 0
# HELP kamailio_core_rcv_replies_5xx_cancel core.rcv_replies_5xx_cancel
# TYPE kamailio_core_rcv_replies_5xx_cancel counter
kamailio_core_rcv_replies_5xx_cancel 0
# HELP kamailio_core_rcv_replies_5xx_invite core.rcv_replies_5xx_invite
# TYPE kamailio_core_rcv_replies_5xx_invite counter
kamailio_core_rcv_replies_5xx_invite 0
# HELP kamailio_core_rcv_replies_5xx_message core.rcv_replies_5xx_message
# TYPE kamailio_core_rcv_replies_5xx_message counter
kamailio_core_rcv_replies_5xx_message 0
# HELP kamailio_core_rcv_replies_5xx_prack core.rcv_replies_5xx_prack
# TYPE kamailio_core_rcv_replies_5xx_prack counter
kamailio_core_rcv_replies_5xx_prack 0
# HELP kamailio_core_rcv_replies_5xx_refer core.rcv_replies_5xx_refer
# TYPE kamailio_core_rcv_replies_5xx_refer counter
kamailio_core_rcv_replies_5xx_refer 0
# HELP kamailio_core_rcv_replies_5xx_reg core.rcv_replies_5xx_reg
# TYPE kamailio_core_rcv_replies_5xx_reg counter
kamailio_core_rcv_replies_5xx_reg 0
# HELP kamailio_core_rcv_replies_5xx_update core.rcv_replies_5xx_update
# TYPE kamailio_core_rcv_replies_5xx_update counter
kamailio_core_rcv_replies_5xx_update 0
# HELP kamailio_core_rcv_replies_6xx core.rcv_replies_6xx
# TYPE kamailio_core_rcv_replies_6xx counter
kamailio_core_rcv_replies_6xx 0
# HELP kamailio_core_rcv_replies_6xx_bye core.rcv_replies_6xx_bye
# TYPE kamailio_core_rcv_replies_6xx_bye counter
kamailio_core_rcv_replies_6xx_bye 0
# HELP kamailio_core_rcv_replies_6xx_cancel core.rcv_replies_6xx_cancel
# TYPE kamailio_core_rcv_replies_6xx_cancel counter
kamailio_core_rcv_replies_6xx_cancel 0
# HELP kamailio_core_rcv_replies_6xx_invite core.rcv_replies_6xx_invite
# TYPE kamailio_core_rcv_replies_6xx_invite counter
kamailio_core_rcv_replies_6xx_invite 0
# HELP kamailio_core_rcv_replies_6xx_message core.rcv_replies_6xx_message
# TYPE kamailio_core_rcv_replies_6xx_message counter
kamailio_core_rcv_replies_6xx_message 0
# HELP kamailio_core_rcv_replies_6xx_prack core.rcv_replies_6xx_prack
# TYPE kamailio_core_rcv_replies_6xx_prack counter
kamailio_core_rcv_replies_6xx_prack 0
# HELP kamailio_core_rcv_replies_6xx_refer core.rcv_replies_6xx_refer
# TYPE kamailio_core_rcv_replies_6xx_refer counter
kamailio_core_rcv_replies_6xx_refer 0
# HELP kamailio_core_rcv_replies_6xx_reg core.rcv_replies_6xx_reg
# TYPE kamailio_core_rcv_replies_6xx_reg counter
kamailio_core_rcv_replies_6xx_reg 0
# HELP kamailio_core_rcv_replies_6xx_update core.rcv_replies_6xx_update
# TYPE kamailio_core_rcv_replies_6xx_update counter
kamailio_core_rcv_replies_6xx_update 0
# HELP kamailio_core_rcv_requests core.rcv_requests
# TYPE kamailio_core_rcv_requests counter
kamailio_core_rcv_requests 3532
# HELP kamailio_core_rcv_requests_ack core.rcv_requests_ack
# TYPE kamailio_core_rcv_requests_ack counter
kamailio_core_rcv_requests_ack 0
# HELP kamailio_core_rcv_requests_bye core.rcv_requests_bye
# TYPE kamailio_core_rcv_requests_bye counter
kamailio_core_rcv_requests_bye 0
# HELP kamailio_core_rcv_requests_cancel core.rcv_requests_cancel
# TYPE kamailio_core_rcv_requests_cancel counter
kamailio_core_rcv_requests_cancel 0
# HELP kamailio_core_rcv_requests_info core.rcv_requests_info
# TYPE kamailio_core_rcv_requests_info counter
kamailio_core_rcv_requests_info 0
# HELP kamailio_core_rcv_requests_invite core.rcv_requests_invite
# TYPE kamailio_core_rcv_requests_invite counter
kamailio_core_rcv_requests_invite 0
# HELP kamailio_core_rcv_requests_message core.rcv_requests_message
# TYPE kamailio_core_rcv_requests_message counter
kamailio_core_rcv_requests_message 0
# HELP kamailio_core_rcv_requests_notify core.rcv_requests_notify
# TYPE kamailio_core_rcv_requests_notify counter
kamailio_core_rcv_requests_notify 0
# HELP kamailio_core_rcv_requests_options core.rcv_requests_options
# TYPE kamailio_core_rcv_requests_options counter
kamailio_core_rcv_requests_options 2011
# HELP kamailio_core_rcv_requests_prack core.rcv_requests_prack
# TYPE kamailio_core_rcv_requests_prack counter
kamailio_core_rcv_requests_prack 0
# HELP kamailio_core_rcv_requests_publish core.rcv_requests_publish
# TYPE kamailio_core_rcv_requests_publish counter
kamailio_core_rcv_requests_publish 446
# HELP kamailio_core_rcv_requests_refer core.rcv_requests_refer
# TYPE kamailio_core_rcv_requests_refer counter
kamailio_core_rcv_requests_refer 0
# HELP kamailio_core_rcv_requests_register core.rcv_requests_register
# TYPE kamailio_core_rcv_requests_register counter
kamailio_core_rcv_requests_register 0
# HELP kamailio_core_rcv_requests_subscribe core.rcv_requests_subscribe
# TYPE kamailio_core_rcv_requests_subscribe counter
kamailio_core_rcv_requests_subscribe 746
# HELP kamailio_core_rcv_requests_update core.rcv_requests_update
# TYPE kamailio_core_rcv_requests_update counter
kamailio_core_rcv_requests_update 0
# HELP kamailio_core_unsupported_methods core.unsupported_methods
# TYPE kamailio_core_unsupported_methods counter
kamailio_core_unsupported_methods 0
# HELP kamailio_dns_failed_dns_request dns.failed_dns_request
# TYPE kamailio_dns_failed_dns_request counter
kamailio_dns_failed_dns_request 0
# HELP kamailio_dns_slow_dns_request dns.slow_dns_request
# TYPE kamailio_dns_slow_dns_request counter
kamailio_dns_slow_dns_request 0
# HELP kamailio_exporter_failed_scrapes Number of failed kamailio scrapes
# TYPE kamailio_exporter_failed_scrapes counter
kamailio_exporter_failed_scrapes 0
# HELP kamailio_exporter_total_scrapes Number of total kamailio scrapes
# TYPE kamailio_exporter_total_scrapes counter
kamailio_exporter_total_scrapes 2
# HELP kamailio_mysql_driver_errors mysql.driver_errors
# TYPE kamailio_mysql_driver_errors counter
kamailio_mysql_driver_errors 0
# HELP kamailio_registrar_accepted_regs registrar.accepted_regs
# TYPE kamailio_registrar_accepted_regs counter
kamailio_registrar_accepted_regs 0
# HELP kamailio_registrar_default_expire registrar.default_expire
# TYPE kamailio_registrar_default_expire counter
kamailio_registrar_default_expire 3600
# HELP kamailio_registrar_default_expires_range registrar.default_expires_range
# TYPE kamailio_registrar_default_expires_range counter
kamailio_registrar_default_expires_range 0
# HELP kamailio_registrar_expires_range registrar.expires_range
# TYPE kamailio_registrar_expires_range counter
kamailio_registrar_expires_range 0
# HELP kamailio_registrar_max_contacts registrar.max_contacts
# TYPE kamailio_registrar_max_contacts counter
kamailio_registrar_max_contacts 0
# HELP kamailio_registrar_max_expires registrar.max_expires
# TYPE kamailio_registrar_max_expires counter
kamailio_registrar_max_expires 0
# HELP kamailio_registrar_rejected_regs registrar.rejected_regs
# TYPE kamailio_registrar_rejected_regs counter
kamailio_registrar_rejected_regs 0
# HELP kamailio_shmem_fragments shmem.fragments
# TYPE kamailio_shmem_fragments gauge
kamailio_shmem_fragments 4
# HELP kamailio_shmem_free_size shmem.free_size
# TYPE kamailio_shmem_free_size gauge
kamailio_shmem_free_size 1.221669264e+09
# HELP kamailio_shmem_max_used_size shmem.max_used_size
# TYPE kamailio_shmem_max_used_size gauge
kamailio_shmem_max_used_size 3.46596e+06
# HELP kamailio_shmem_real_used_size shmem.real_used_size
# TYPE kamailio_shmem_real_used_size gauge
kamailio_shmem_real_used_size 3.067504e+06
# HELP kamailio_shmem_total_size shmem.total_size
# TYPE kamailio_shmem_total_size gauge
kamailio_shmem_total_size 1.224736768e+09
# HELP kamailio_shmem_used_size shmem.used_size
# TYPE kamailio_shmem_used_size gauge
kamailio_shmem_used_size 2.602616e+06
# HELP kamailio_sl_1xx_replies sl.1xx_replies
# TYPE kamailio_sl_1xx_replies counter
kamailio_sl_1xx_replies 0
# HELP kamailio_sl_200_replies sl.200_replies
# TYPE kamailio_sl_200_replies counter
kamailio_sl_200_replies 364
# HELP kamailio_sl_202_replies sl.202_replies
# TYPE kamailio_sl_202_replies counter
kamailio_sl_202_replies 0
# HELP kamailio_sl_2xx_replies sl.2xx_replies
# TYPE kamailio_sl_2xx_replies counter
kamailio_sl_2xx_replies 0
# HELP kamailio_sl_300_replies sl.300_replies
# TYPE kamailio_sl_300_replies counter
kamailio_sl_300_replies 0
# HELP kamailio_sl_301_replies sl.301_replies
# TYPE kamailio_sl_301_replies counter
kamailio_sl_301_replies 0
# HELP kamailio_sl_302_replies sl.302_replies
# TYPE kamailio_sl_302_replies counter
kamailio_sl_302_replies 0
# HELP kamailio_sl_3xx_replies sl.3xx_replies
# TYPE kamailio_sl_3xx_replies counter
kamailio_sl_3xx_replies 0
# HELP kamailio_sl_400_replies sl.400_replies
# TYPE kamailio_sl_400_replies counter
kamailio_sl_400_replies 0
# HELP kamailio_sl_401_replies sl.401_replies
# TYPE kamailio_sl_401_replies counter
kamailio_sl_401_replies 0
# HELP kamailio_sl_403_replies sl.403_replies
# TYPE kamailio_sl_403_replies counter
kamailio_sl_403_replies 0
# HELP kamailio_sl_404_replies sl.404_replies
# TYPE kamailio_sl_404_replies counter
kamailio_sl_404_replies 0
# HELP kamailio_sl_407_replies sl.407_replies
# TYPE kamailio_sl_407_replies counter
kamailio_sl_407_replies 0
# HELP kamailio_sl_408_replies sl.408_replies
# TYPE kamailio_sl_408_replies counter
kamailio_sl_408_replies 0
# HELP kamailio_sl_483_replies sl.483_replies
# TYPE kamailio_sl_483_replies counter
kamailio_sl_483_replies 0
# HELP kamailio_sl_4xx_replies sl.4xx_replies
# TYPE kamailio_sl_4xx_replies counter
kamailio_sl_4xx_replies 0
# HELP kamailio_sl_500_replies sl.500_replies
# TYPE kamailio_sl_500_replies counter
kamailio_sl_500_replies 0
# HELP kamailio_sl_5xx_replies sl.5xx_replies
# TYPE kamailio_sl_5xx_replies counter
kamailio_sl_5xx_replies 0
# HELP kamailio_sl_6xx_replies sl.6xx_replies
# TYPE kamailio_sl_6xx_replies counter
kamailio_sl_6xx_replies 0
# HELP kamailio_sl_failures sl.failures
# TYPE kamailio_sl_failures counter
kamailio_sl_failures 0
# HELP kamailio_sl_received_ACKs sl.received_ACKs
# TYPE kamailio_sl_received_ACKs counter
kamailio_sl_received_ACKs 0
# HELP kamailio_sl_sent_err_replies sl.sent_err_replies
# TYPE kamailio_sl_sent_err_replies counter
kamailio_sl_sent_err_replies 0
# HELP kamailio_sl_sent_replies sl.sent_replies
# TYPE kamailio_sl_sent_replies counter
kamailio_sl_sent_replies 364
# HELP kamailio_sl_xxx_replies sl.xxx_replies
# TYPE kamailio_sl_xxx_replies counter
kamailio_sl_xxx_replies 0
# HELP kamailio_tcp_con_reset tcp.con_reset
# TYPE kamailio_tcp_con_reset counter
kamailio_tcp_con_reset 0
# HELP kamailio_tcp_con_timeout tcp.con_timeout
# TYPE kamailio_tcp_con_timeout counter
kamailio_tcp_con_timeout 0
# HELP kamailio_tcp_connect_failed tcp.connect_failed
# TYPE kamailio_tcp_connect_failed counter
kamailio_tcp_connect_failed 0
# HELP kamailio_tcp_connect_success tcp.connect_success
# TYPE kamailio_tcp_connect_success counter
kamailio_tcp_connect_success 0
# HELP kamailio_tcp_current_opened_connections tcp.current_opened_connections
# TYPE kamailio_tcp_current_opened_connections gauge
kamailio_tcp_current_opened_connections 1
# HELP kamailio_tcp_current_write_queue_size tcp.current_write_queue_size
# TYPE kamailio_tcp_current_write_queue_size gauge
kamailio_tcp_current_write_queue_size 0
# HELP kamailio_tcp_established tcp.established
# TYPE kamailio_tcp_established counter
kamailio_tcp_established 1
# HELP kamailio_tcp_local_reject tcp.local_reject
# TYPE kamailio_tcp_local_reject counter
kamailio_tcp_local_reject 0
# HELP kamailio_tcp_passive_open tcp.passive_open
# TYPE kamailio_tcp_passive_open counter
kamailio_tcp_passive_open 1
# HELP kamailio_tcp_send_timeout tcp.send_timeout
# TYPE kamailio_tcp_send_timeout counter
kamailio_tcp_send_timeout 0
# HELP kamailio_tcp_sendq_full tcp.sendq_full
# TYPE kamailio_tcp_sendq_full counter
kamailio_tcp_sendq_full 0
# HELP kamailio_tmx_2xx_transactions tmx.2xx_transactions
# TYPE kamailio_tmx_2xx_transactions counter
kamailio_tmx_2xx_transactions 1469
# HELP kamailio_tmx_3xx_transactions tmx.3xx_transactions
# TYPE kamailio_tmx_3xx_transactions counter
kamailio_tmx_3xx_transactions 0
# HELP kamailio_tmx_4xx_transactions tmx.4xx_transactions
# TYPE kamailio_tmx_4xx_transactions counter
kamailio_tmx_4xx_transactions 232
# HELP kamailio_tmx_5xx_transactions tmx.5xx_transactions
# TYPE kamailio_tmx_5xx_transactions counter
kamailio_tmx_5xx_transactions 14
# HELP kamailio_tmx_6xx_transactions tmx.6xx_transactions
# TYPE kamailio_tmx_6xx_transactions counter
kamailio_tmx_6xx_transactions 0
# HELP kamailio_tmx_UAC_transactions tmx.UAC_transactions
# TYPE kamailio_tmx_UAC_transactions counter
kamailio_tmx_UAC_transactions 702
# HELP kamailio_tmx_UAS_transactions tmx.UAS_transactions
# TYPE kamailio_tmx_UAS_transactions counter
kamailio_tmx_UAS_transactions 1750
# HELP kamailio_tmx_active_transactions tmx.active_transactions
# TYPE kamailio_tmx_active_transactions gauge
kamailio_tmx_active_transactions 0
# HELP kamailio_tmx_inuse_transactions tmx.inuse_transactions
# TYPE kamailio_tmx_inuse_transactions gauge
kamailio_tmx_inuse_transactions 0
# HELP kamailio_tmx_rpl_absorbed tmx.rpl_absorbed
# TYPE kamailio_tmx_rpl_absorbed counter
kamailio_tmx_rpl_absorbed 14
# HELP kamailio_tmx_rpl_generated tmx.rpl_generated
# TYPE kamailio_tmx_rpl_generated counter
kamailio_tmx_rpl_generated 1036
# HELP kamailio_tmx_rpl_received tmx.rpl_received
# TYPE kamailio_tmx_rpl_received counter
kamailio_tmx_rpl_received 693
# HELP kamailio_tmx_rpl_relayed tmx.rpl_relayed
# TYPE kamailio_tmx_rpl_relayed counter
kamailio_tmx_rpl_relayed 679
# HELP kamailio_tmx_rpl_sent tmx.rpl_sent
# TYPE kamailio_tmx_rpl_sent counter
kamailio_tmx_rpl_sent 1715
# HELP kamailio_up Was the last scrape successful.
# TYPE kamailio_up gauge
kamailio_up 1
# HELP kamailio_usrloc_registered_users usrloc.registered_users
# TYPE kamailio_usrloc_registered_users gauge
kamailio_usrloc_registered_users 0
```

## Custom Metrics

Additional metrics can be added to the output of `kamcmd stats.fetch all` using the
[`statistics`](https://www.kamailio.net/docs/modules/stable/modules/statistics.html) module.

⚠️ If you want your custom metric to be exported as a [gauge](https://prometheus.io/docs/concepts/metric_types/#gauge),
then you should include the word "gauge" in your statistic variable name. Similarly, if you'd like
your custom metric to be a [counter](https://prometheus.io/docs/concepts/metric_types/#counter) then
include the word "counter" in your statistic name.

## Acknowledgements

I took a _lot_ of inspiration from the following projects:
- https://github.com/florentchauveau/kamailio_exporter
- https://github.com/pascomnet/kamailio_exporter

Why did I create yet another exporter?
- Metric naming
  - I didn't want metrics grouped by "label". (I'm sending these exported metrics to AWS CloudWatch.
    It's unclear from their [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-PrometheusEC2.html#CloudWatch-Agent-PrometheusEC2-configure-agent)
    how best to wire things up properly.)
  - I wanted names that closely matched the output of `kamcmd stats.fetch all`.
- Correctness
  - I haven't come across a project that had all of the exported metrics I was interested _and_ had
    them assigned with the correct value type. It's not to say that I'm 100% correct/comprehensive
    in this project either. There's even some code that _guesses_ what the type should be. I like
    to think that any mistakes will be quickly rectified.
- Simplicity
  - I'm only interested in the output of `kamcmd stats.fetch all` at this time. There's no extra
    options/configuration for things that _I_ don't need.
  - Convention over configuration for custom metrics. The names might be uglier than they would
    otherwise be and there's some magic that doesn't always work but in _my_ use case I can give
    the metrics better names when they're presented in a graph.
