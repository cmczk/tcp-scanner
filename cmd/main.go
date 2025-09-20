package main

import (
	"github.com/cmczk/tcp-scanner/cli"
	"github.com/cmczk/tcp-scanner/scanner"
)

func main() {
	opts := cli.ParseFlags()

	s := scanner.New(opts.Host, opts.MinPort, opts.MaxPort)
	s.Scan()
}
