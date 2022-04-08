package main

import (
	"hscan/hscan"
	"flag"
	"log"
	"fmt"
)

func main() {
	filePtr := flag.String("file", "", "The search term to query.")
	flag.Parse()
	if filePtr == nil || *filePtr == "" {
		log.Fatalln("Usage: main --file <wordlist.txt>")
	}
	var file = *filePtr

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0" // p@ssword
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032" // letmein

	log.Print("GuessSingle(md5hash, file)")
	hscan.GuessSingle(md5hash, file)

	log.Print("GuessSingle(sha256hash, file)")
	hscan.GuessSingle(sha256hash, file)

	log.Print("GenHashMaps(file)")
	hscan.GenHashMaps(file)

	log.Print("GetSHA(sha256hash)")
	pwd, err := hscan.GetSHA(sha256hash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("[+] Password found (SHA): %s\n", pwd)

	log.Print("GetMD5(sha256hash)")
	pwd, err = hscan.GetMD5(sha256hash) // This should fail, but no task for it.
	if err == nil {
		fmt.Printf("[+] Password found (MD5): %s\n", pwd)
	}

	log.Print("GetMD5(md5hash)")
	pwd, err = hscan.GetMD5(md5hash) // Maybe will work.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[+] Password found (MD5): %s\n", pwd)
	}

	log.Print("GetMD5(drmike1)")
	pwd, err = hscan.GetMD5(drmike1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[+] Password found (MD5): %s\n", pwd)
	}

	log.Print("GetSHA(drmike2)")
	pwd, err = hscan.GetSHA(drmike2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[+] Password found (SHA): %s\n", pwd)
	}
}
