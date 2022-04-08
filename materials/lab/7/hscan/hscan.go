package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) bool {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := false
	// Create check function to avoid branching in loop for possibly better speed.
	var checkFn = func (password string) {}
	if len(sourceHash) == 32 {
		checkFn = func (password string) {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				res = true
			}
		}
	} else {
		checkFn = func (password string) {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				res = true
			}
		}
	}
	for scanner.Scan() {
		password := scanner.Text()
		checkFn(password)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return res
}

// GenHashMaps generates hashmaps for use later. Returns the number of passwords hashed.
func GenHashMaps(filename string) int {

	md5lookup = make(map[string]string)
	shalookup = make(map[string]string)

	var wg sync.WaitGroup
	var mx = &sync.RWMutex{}

	//OPTIONAL -- Can you use workers to make this even faster

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0

	worker := func (jobs <- chan string) {
		for p := range jobs {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(p)))
			hash2 := fmt.Sprintf("%x", sha256.Sum256([]byte(p)))
			mx.Lock()
			md5lookup[hash] = p
			shalookup[hash2] = p
			mx.Unlock()
			wg.Done()
		}
	}

	jobs := make(chan string)

	for w := 1; w <= 100; w++ {
		go worker(jobs)
	}

	for scanner.Scan() {
		password := scanner.Text()
		wg.Add(1)
		jobs <- password
		res += 1
	}

	wg.Wait()
	return res
}

// GenHashMapsSlow generates hashmaps for use later. Returns the number of passwords hashed.
func GenHashMapsSlow(filename string) int {

	md5lookup = make(map[string]string)
	shalookup = make(map[string]string)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0

	for scanner.Scan() {
		password := scanner.Text()
		
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		md5lookup[hash] = password

		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		shalookup[hash] = password

		res += 1
	}

	return res
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil
	} else {
		return "", errors.New("password does not exist")
	}
}
