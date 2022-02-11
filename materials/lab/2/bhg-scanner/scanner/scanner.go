// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// In a main.go file use
// import "bhg-scanner/scanner"
// scanner.PortScanner()

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

type ResultPair struct {
	isOpen bool
	port int
}

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int  // notice the capitalization here. access limited!
var closedports []int

func worker(ports chan int, results chan ResultPair) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil { 
			results <- ResultPair { isOpen: false, port: p }
			continue
		}
		conn.Close()
		results <- ResultPair { isOpen: true, port: p }
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner() int {  

	ports := make(chan int, 100)   // Leaving this as advised in lab lecture.
	results := make(chan ResultPair)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port.isOpen {
			openports = append(openports, port.port)
		} else {
			closedports = append(closedports, port.port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	for _, port := range closedports {
		fmt.Printf("%d closed\n", port)
	}

	return len(openports) + len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
