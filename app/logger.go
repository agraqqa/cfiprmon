// This file contains the code for logging
package main

import (
	"log"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// init initializes loggers
func init() {
	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(log.Writer(), "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
