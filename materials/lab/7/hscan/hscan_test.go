// Testing for hscan. Run tests using go test.

package hscan

import (
	"testing"
	"time"
	"fmt"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "../main/wordlist.txt") // Currently function returns only number of open ports
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGenHashMapsTiming(t *testing.T) {
	start := time.Now()
	numPwd := GenHashMaps("../main/Top1pt6Million-probable-v2.txt")
	dur1 := time.Since(start)
	fmt.Println("Time to generate fast hashmaps: ", dur1)
	fmt.Println("Fast passwords per second: ", float64(numPwd) / dur1.Seconds())

	start = time.Now()
	numPwd = GenHashMapsSlow("../main/Top1pt6Million-probable-v2.txt")
	dur2 := time.Since(start)
	fmt.Println("Time to generate hashmaps: ", dur2)
	fmt.Println("Passwords per second: ", float64(numPwd) / dur2.Seconds())
}

func TestGenHashMapsWorks(t *testing.T) {
	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"
	GenHashMaps("../main/Top1pt6Million-probable-v2.txt")
	pwd, _ := GetMD5(drmike1)
	if pwd != "p@ssword" {
		t.Errorf("got %v, wanted %v", pwd, "p@ssword")
	}

	pwd, _ = GetSHA(drmike2)
	if pwd != "letmein" {
		t.Errorf("got %v, wanted %v", pwd, "letmein")
	}
}
