// Package parser is providing with cmdline tool to process WKT or epsg.properties
package main

import (
	"log"
	"os"

	"github.com/fikin/wkt-crs-go/internal/cmdline"
)

func main() {
	l := log.Default()

	fn := cmdline.ParseMainOpts(l, os.Args[1:])

	// nolint:staticcheck
	if usage, err := fn(); err != nil {
		l.Printf("[ERR] %v\n", err)
		usage()
		os.Exit(1)
	}
}
