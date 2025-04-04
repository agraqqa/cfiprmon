package main

import (
	"flag"
	"fmt"
	"os"
)

// Function to parse command line arguments
func parseFlags() *Config {
	var help bool
	var flags = Config{}
	var usage = `
Usage: op2vault [flags]

Flags:
  --debug=<true/false>
    Show debug information. Default: false

  --pushgateway-url <url>
    Pushgateway URL. Default: prometheus-pushgateway.monitoring.svc.cluster.local:9091
`
	// Define flags
	flag.BoolVar(&help, "help", false, "Display help")
	flag.BoolVar(&flags.Debug, "debug", false, "Show debug information")
	flag.StringVar(&flags.PushgatewayURL, "pushgateway-url", "prometheus-pushgateway.monitoring.svc.cluster.local:9091", "Pushgateway URL")

	// Parse the flags
	flag.Parse()

	// Display help if --help or -h is provided
	if help {
		fmt.Println(usage)
		os.Exit(0)
	}

	return &flags
}
