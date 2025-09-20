package main

import (
	"github.com/cmczk/tcp-scanner/scanner"
)

const host = "scanme.nmap.org"
const maxPort = 65535

func main() {
	s := scanner.New(host, nil, maxPort)
	s.Scan()
}
