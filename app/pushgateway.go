package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

const (
	namespace = "cloudflare"
	subsystem = "ip_ranges"
)

// Metrics represents the Prometheus metrics we collect
type Metrics struct {
    ipv4Checksum *prometheus.GaugeVec
    ipv6Checksum *prometheus.GaugeVec
}

// NewMetrics creates and returns new Metrics
func NewMetrics() *Metrics {
    return &Metrics{
        ipv4Checksum: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: namespace,
                Subsystem: subsystem,
                Name:      "ipv4_checksum",
                Help:      "Checksum of Cloudflare IPv4 ranges file",
            },
            []string{},
        ),
        ipv6Checksum: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: namespace,
                Subsystem: subsystem,
                Name:      "ipv6_checksum",
                Help:      "Checksum of Cloudflare IPv6 ranges file",
            },
            []string{},
        ),
    }
}

// PushMetrics sends metrics to Pushgateway
func PushMetrics(pushgatewayURL string, ipv4Checksum, ipv6Checksum uint32) error {
    metrics := NewMetrics()

    // Set the checksum values directly as float64
    metrics.ipv4Checksum.WithLabelValues().Set(float64(ipv4Checksum))
    metrics.ipv6Checksum.WithLabelValues().Set(float64(ipv6Checksum))

    pusher := push.New(pushgatewayURL, "cloudflare_ip_ranges_monitor")

    pusher.Collector(metrics.ipv4Checksum)
    pusher.Collector(metrics.ipv6Checksum)

    if err := pusher.Push(); err != nil {
        return fmt.Errorf("could not push metrics to Pushgateway: %v", err)
    }

    return nil
}
