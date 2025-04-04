package main

import (
	"io"
	"log"
)

// Main function
func main() {
	config, err := loadConfig()
	if err != nil {
		ErrorLogger.Fatal(err)
	}
	if !config.Debug {
		DebugLogger = log.New(io.Discard, "", 0)
	}
	DebugLogger.Printf("Configuration loaded")

	ipv4Checksum, err := GetCloudflareIPRangesChecksum(CloudflareIPv4URL)
	if err != nil {
	    ErrorLogger.Fatalf("Failed to get IPv4 hash: %v", err)
	}
	DebugLogger.Printf("IPv4 hash: %d", ipv4Checksum)

	ipv6Checksum, err := GetCloudflareIPRangesChecksum(CloudflareIPv6URL)
	if err != nil {
	    ErrorLogger.Fatalf("Failed to get IPv6 hash: %v", err)
	}
	DebugLogger.Printf("IPv6 hash: %d", ipv6Checksum)

	if err := PushMetrics(config.PushgatewayURL, ipv4Checksum, ipv6Checksum); err != nil {
		ErrorLogger.Fatalf("Failed to push metrics: %v", err)
	} else {
		InfoLogger.Println("Successfully pushed metrics to Pushgateway")
	}
}
