package main

import (
	"flag"
	"os"
	"strings"
)

// Config stores application configuration
type Config struct {
	// App debug mode
	Debug     bool
	// Pushgateway URL
	PushgatewayURL string
}

// defaultConfig stores default configuration (falses and empty strings)
var defaultConfig = Config{}

// loadEnvConfig loads configuration from environment variables
func loadEnvConfig() *Config {
	var appDebug bool
	debug := strings.ToLower(os.Getenv("CFIPRMON_DEBUG"))
	pushgatewayURL := os.Getenv("CFIPRMON_PUSHGATEWAY_URL")

	// Check values
	if debug == "true" || debug == "1" {
		appDebug = true
	}

	return &Config{
		Debug:            appDebug,
		PushgatewayURL: pushgatewayURL,
	}
}

// loadConfig loads configuration from all sources in order of priority:
// 1. Command line arguments (flags)
// 2. Environment variables
// 3. Default configuration
func loadConfig() (*Config, error) {
	// Start with default configuration
	config := defaultConfig

	// Load environment variables
	envConfig := loadEnvConfig()
	config = *envConfig
	// Load flags
	flagConfig := parseFlags()
	// Visit only set flags
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "debug":
			config.Debug = flagConfig.Debug
		case "pushgateway-url":
			config.PushgatewayURL = flagConfig.PushgatewayURL
		}
	})

	// Validate required fields
	if config.PushgatewayURL == "" {
		config.PushgatewayURL = "prometheus-pushgateway.monitoring.svc.cluster.local:9091"
	}

	return &config, nil
}
