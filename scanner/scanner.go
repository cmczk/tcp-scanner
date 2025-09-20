package scanner

import (
	"fmt"
	"net"
	"sort"
)

type Scanner struct {
	host    string
	minPort int
	maxPort int
}

func New(host string, minPort int, maxPort int) *Scanner {
	return &Scanner{
		host:    host,
		minPort: minPort,
		maxPort: maxPort,
	}
}

func (s *Scanner) Scan() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go s.worker(ports, results)
	}

	go func() {
		for i := s.minPort; i <= s.maxPort; i++ {
			ports <- i
		}
	}()

	for range s.maxPort {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("port %d open\n", port)
	}
}

func (s *Scanner) worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("[%s]:%d", s.host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}
