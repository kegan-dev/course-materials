package main

import "bhg-scanner/scanner"

func main(){
	ps := scanner.PortScanner {
		StartPort: 1,
		EndPort: 1024,
		Timeout: 1,
	}
	ps.DoScan()
}