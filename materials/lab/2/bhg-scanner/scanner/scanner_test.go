package scanner

import (
	"testing"
)

var ps1, smallPs, noPs PortScanner
func init() {
	ps1 = PortScanner {
		StartPort: 1,
		EndPort: 1024,
		Timeout: 1,
	}

	smallPs = PortScanner {
		StartPort: 70,
		EndPort: 80,
		Timeout: -1,
	}

	noPs = PortScanner {
		StartPort: 90,
		EndPort: 80,
		Timeout: -1,
	}
}

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

    got, _ := ps1.DoScan() // Currently function returns only number of open ports
    want := 1 // 80 is not open with this as expected from lab lecture.

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
    got1, got2 := ps1.DoScan()
    want := 1024

    if got1 + got2 != want {
        t.Errorf("got %d, wanted %d", got1 + got2, want)
    }
}

func TestSmallTotalPortsScanned(t *testing.T){
    got1, got2 := smallPs.DoScan()
    want := 11 // 80 - 70 + 1 for inclusive range.

    if got1 + got2 != want {
        t.Errorf("got %d, wanted %d", got1 + got2, want)
    }
}

func TestNoTotalPortsScanned(t *testing.T){
    got1, got2 := noPs.DoScan()
    want := 0 // 80 - 70 + 1 for inclusive range.

    if got1 + got2 != want {
        t.Errorf("got %d, wanted %d", got1 + got2, want)
    }
}
