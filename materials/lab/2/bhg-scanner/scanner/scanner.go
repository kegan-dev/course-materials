// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// In a main.go file use
// import "bhg-scanner/scanner"
// ps := scanner.PortScanner {StartPort: 1, EndPort: 1024, Timeout: 1}
// ps.DoScan()

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

// Port scanner with inclusive start, end, and timeout in seconds.
type PortScanner struct {
	StartPort int
	EndPort int
	Timeout int64
	openports []int
	closedports []int
}

type ResultPair struct {
	isOpen bool
	port int
}

func worker(timeout int64, ports chan int, results chan ResultPair) {
	if timeout < 1 {
		timeout = 1
	}
	
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, time.Duration(timeout) * time.Second)
		if err != nil { 
			results <- ResultPair { isOpen: false, port: p }
			continue
		}
		conn.Close()
		results <- ResultPair { isOpen: true, port: p }
	}
}

// Does a scan as a PortScanner.
// Returns the number of open and closed ports.
func (ps *PortScanner) DoScan() (int, int) {  
	ps.openports = make([]int, 0)
	ps.closedports = make([]int, 0)

	ports := make(chan int, 100)   // Leaving this as advised in lab lecture.
	results := make(chan ResultPair)

	for i := 0; i < cap(ports); i++ {
		go worker(ps.Timeout, ports, results)
	}

	go func() {
		for i := ps.StartPort; i <= ps.EndPort; i++ {
			ports <- i
		}
	}()

	for i := ps.StartPort; i <= ps.EndPort; i++ {
		port := <-results
		if port.isOpen {
			ps.openports = append(ps.openports, port.port)
		} else {
			ps.closedports = append(ps.closedports, port.port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(ps.openports)
	sort.Ints(ps.closedports)

	for _, port := range ps.openports {
		fmt.Printf("%d, open\n", port)
	}

	for _, port := range ps.closedports {
		fmt.Printf("%d, closed\n", port)
	}

	return len(ps.openports), len(ps.closedports)
}
