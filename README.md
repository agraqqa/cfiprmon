# cfiprmon

`cfiprmon` is a simple tool to monitor changes of [Cloudflare IP ranges](https://www.cloudflare.com/ips/) files:
- [IPv4 text list](https://www.cloudflare.com/ips-v4/)
- [IPv6 text list](https://www.cloudflare.com/ips-v6/)

It delivers CRC of these files to Prometheus via Pushgateway as `cloudflare_ip_ranges_ipv4_checksum` and `cloudflare_ip_ranges_ipv6_checksum` metrics.

# Environment variables

- `CFIPRMON_PUSHGATEWAY_URL` - URL of the Prometheus Pushgateway. Default: `prometheus-pushgateway.monitoring.svc.cluster.local:9091`
- `CFIPRMON_DEBUG` - false by default
