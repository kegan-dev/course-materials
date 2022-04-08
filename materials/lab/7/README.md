# Lab 7
Due April 9th at 11:59PM

## Submission

Run using the command `./main --file <filename>`

Using the Top1pt6Million-probable-v2.txt you can find the passwords for the two drmike hashes. wordlist.txt let's you find the rest of the hashes.

A test has been added to hscan_test.go for timing the hashmap creation. It will output the timing information for a fast and slow implementation.

Another test has been added to verify that the hashmaps are generated correctly. It generates the hashmaps and checks for known values to exist for passwords.

## Development Work [18 points]
- [3pt] Complete the TODOs in [main.go](course-materials/materials/lab/7/main/main.go)
- [12pt] Complete the TODOs in [hscan.go](course-materials/materials/lab/7/hscan/hscan.go)
- [3pt] Create at least one new test in [hscan_test.go](course-materials/materials/lab/7/hscan/hscan_test.go)

## Capture  details [2pts]
- Capture Timing Details (per hscan.go) for various implementation of creating the hash maps

## Submit 
1. Link to your Github Repository [16pts]
2. Report the numbers [2pts]
2. List of Collaborator
