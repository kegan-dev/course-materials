// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
	"flag"
)

func main() {
	searchPtr := flag.String("search", "", "The search term to query.")
	helpPtr := flag.Bool("help", false, "Get available facets and filters.")
	flag.Parse()

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	if (*helpPtr) {
		helpInfo , err := s.Helper()
		if err != nil {
			log.Panicln(err)
		}

		fmt.Println("\nAvailable Facets:")
		for _, facet := range helpInfo.AvailableFacets {
			fmt.Println(facet)
		}

		fmt.Println("\nAvailable filters:")
		for _, filter := range helpInfo.AvailableFilters {
			fmt.Println(filter)
		}
		return
	}

	if searchPtr == nil || *searchPtr == "" {
		log.Fatalln("Usage: main --search <searchterm>")
	}
	
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(*searchPtr)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h,_ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}


	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}


}