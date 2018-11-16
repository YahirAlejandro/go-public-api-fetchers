package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func flagCleaner(flag *string) {
	// 'flag' is a pointer of type string
	// All values passed here are ponters to actual values, which this function will
	// modify, to manipulate them in the main function scope.
	// Remember to access the pointer by "*", otherwise the memory address will be modified.
	*flag = strings.Replace(*flag, " ", "+", -1)
}

func queryAPI(url string) []byte {
	// queryAPI: Fetch a url string using Get and returns a []byte slice.
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}

	resp, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	return resp
}

type availableJob struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Url         string `json:"company_url"`
	HowToApply  string
	Company     string
	CompanyUrl  string
	CompanyLogo string
}

func main() {

	// Catch flags
	// Valid flags to build $baseurlrequest
	jobDescriptionFlag := flag.String("desc", "devops", "The job description you're interested in")
	jobLocationFlag := flag.String("loc", "new+york", "The location of the job you're looking for")

	// Parse commandline output
	flag.Parse()

	// Append all flag vars to an array to easily for through it:
	// Here all pointers are passed into the sanitizer function.
	// Once pointers reach the sanitizer function, it will use the pointer to mutate
	// the value they're pointed to.
	flagVarsArray := []*string{jobDescriptionFlag, jobLocationFlag}

	// Send all flag vars to sanitize
	for _, n := range flagVarsArray {
		flagCleaner(n)
	}

	// Building the URL
	const baseurl = "https://jobs.github.com/positions.json?"
	jobDescription := "description=" + *jobDescriptionFlag
	jobLocation := "location=" + *jobLocationFlag
	url := baseurl + jobDescription + "&" + jobLocation

	resp := queryAPI(url)

	if len(resp) > 2 {
		fmt.Printf("%v\n", resp[0])
	} else {
		fmt.Printf("There were %d results, try another query.\n", cap(resp))
	}

	job := []availableJob{}
	err := json.Unmarshal(resp, &job)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	fmt.Printf("%d\n", len(job))
	fmt.Println(reflect.TypeOf(job))

}
