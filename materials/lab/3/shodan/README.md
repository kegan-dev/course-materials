# Option 3

help.go is a new file that pulls currently available facets and filters from the Shodan API. The function Helper does the querying of the Shodan API and returns a struct with the lists of facets and filters.

main.go has been modified to use the new features.
