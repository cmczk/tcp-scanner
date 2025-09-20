package cli

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Host    string
	MinPort int
	MaxPort int
}

func ParseFlags() *Options {
	host := flag.String("host", "", "host to scan, required")
	minPort := flag.Int("min-port", 1, "lower bound for ports")
	maxPort := flag.Int("max-port", 65535, "upper bound for ports")

	flag.Parse()

	opts := Options{
		Host:    *host,
		MinPort: *minPort,
		MaxPort: *maxPort,
	}

	mustCheckFlags(opts)

	return &opts
}

func mustCheckFlags(opts Options) {
	if opts.Host == "" {
		fmt.Printf("[ERROR] host is required\n\n")
		flag.Usage()
		os.Exit(1)
	}
}
