# Usage

First build the main.go file with `go build`

## List facets and filters

`SHODAN_API_KEY=<apikey> ./main --help`

This will list the available facets and filters from the Shodan API.

## Search

`SHODAN_API_KEY=<apikey> ./main --search <searchterm>`

This will do a search of the term provided via the Shodan API.
