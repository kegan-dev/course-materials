package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "os"
    "path/filepath"
    "strconv"
	"strings"
	"regexp"
)


//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and 
// if responsewriter is passed, outputs to http 

var LOG_LEVEL int
var NUM_FILES int64 = 0
func walkFn(w http.ResponseWriter) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")
        for _, r := range regexes {
            if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)
                tfile.Filename = string(filename)
                tfile.Location = string(dir)

				exists := false
				var fileIx int64 = 0
				for i, f := range Files {
					if f.Filename == tfile.Filename && f.Location == tfile.Location {
						exists = true
						fileIx = int64(i+1)
						break
					}
				}

				if !exists {
					NUM_FILES += 1
					fileIx = NUM_FILES
					Files = append(Files, tfile)
				}

                if w != nil && NUM_FILES > 0 {
                    w.Write([]byte(`"`+(strconv.FormatInt(fileIx, 10))+`":  `))
                    json.NewEncoder(w).Encode(tfile)
                    w.Write([]byte(`,`))

                } 
                
				if LOG_LEVEL > 1 {
                	log.Printf("[+] HIT: %s\n", path)
				}
            }

        }
        return nil
    }

}

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")

		r := regexp.MustCompile(query)
		if r.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)

			exists := false
			var fileIx int64 = 0
			for i, f := range Files {
				if f.Filename == tfile.Filename && f.Location == tfile.Location {
					exists = true
					fileIx = int64(i+1)
					break
				}
			}

			if !exists {
				NUM_FILES += 1
				fileIx = NUM_FILES
				Files = append(Files, tfile)
			}

			if w != nil && NUM_FILES > 0 {
				w.Write([]byte(`"`+(strconv.FormatInt(fileIx, 10))+`":  `))
				json.NewEncoder(w).Encode(tfile)
				w.Write([]byte(`,`))
			}
			
			if LOG_LEVEL > 1 {
				log.Printf("[+] HIT: %s\n", path)
			}
		}

        return nil
    }
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{ "status" : "API is up and running ",`))
    var regexstrings []string
    
    for _, regex := range regexes{
        regexstrings = append(regexstrings, regex.String())
    }

    w.Write([]byte(` "regexs" :`))
    json.NewEncoder(w).Encode(regexstrings)
    w.Write([]byte(`}`))

	if LOG_LEVEL > 1 {
		log.Println(regexes)
	}

}


func MainPage(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

    w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html><body><H1>Welcome to my awesome file scraper API home page.</H1><p>This API allows you to scrape files that match patterns defined by the user.</p></body></html>")
}


func FindFile(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}

    q, ok := r.URL.Query()["q"]

    w.WriteHeader(http.StatusOK)
    if ok && len(q[0]) > 0 {
		if LOG_LEVEL > 0 {
        	log.Printf("Entering search with query=%s",q[0])
		}

        // ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
        // e.g., func finder(query string) []string { ... }
		found := false
        for _, File := range Files {
		    if File.Filename == q[0] {
                json.NewEncoder(w).Encode(File.Location)
                found = true
		    }
        }

		if !found {
			w.Write([]byte("File not found."))
		}
    } else {
        // didn't pass in a search term, show all that you've found
        w.Write([]byte(`"files":`))    
        json.NewEncoder(w).Encode(Files)
    }
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}

    w.Header().Set("Content-Type", "application/json")

    location, locOK := r.URL.Query()["location"]
	regex, regexOk := r.URL.Query()["regex"]
    
	rootDir := "/home/cabox"
    if locOK && len(location[0]) > 0 {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusFailedDependency)
        w.Write([]byte(`{ "parameters" : {"required": "location",`))    
        w.Write([]byte(`"optional": "regex"},`))    
        w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
        w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
        return 
    }

    //wrapper to make "nice json"
    w.Write([]byte(`{ `))
    
    // Define the logic required here to call the new function walkFn2(w,regex[0])
    // Hint, you need to grab the regex parameter (see how it's done for location above...) 
    
    // if regexOK
    //   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
    // else run code to locate files matching stored regular expression
	if regexOk {
		if err := filepath.Walk(rootDir + location[0], walkFn2(w, `(i?)`+regex[0])); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := filepath.Walk(rootDir + location[0], walkFn(w)); err != nil {
			log.Panicln(err)
		}
	}

    //wrapper to make "nice json"
    w.Write([]byte(` "status": "completed"} `))

}

func Reset(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
	
	resetRegEx()
	Files = nil
	NUM_FILES = 0
}

func Clear(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
	
	clearRegEx()
}

func AddRegex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
	
	addRegEx(`(?i)` + strings.TrimPrefix(r.URL.Path, "/addsearch/"))
}

// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint
