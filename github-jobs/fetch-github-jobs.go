package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/YahirAlejandro/go-public-api-fetchers/github-jobs/telebotclient"
)

type fetchJobs struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	CompanyLogo string `json:"company_logo"`
	URL         string `json:"url"`
}

func flagSanitizer(stringToClean *string) {
	*stringToClean = strings.Replace(*stringToClean, " ", "+", -1)
}

func queryAPIEndpoint(jobDescriptionFlag, jobLocationFlag *string) []byte {
	// Building the URL
	const baseurl = "https://jobs.github.com/positions.json?"
	jobDescription := "description=" + *jobDescriptionFlag
	jobLocation := "location=" + *jobLocationFlag
	url := baseurl + jobDescription + "&" + jobLocation

	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	resp, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	return resp

}

func main() {
	// Flag definition, parsing and sanitizing
	jobDescriptionFlag := flag.String("desc", "devops", "It defaults to 'devops'")
	jobLocationFlag := flag.String("loc", "", "It defaults to nothing")
	flag.Parse()

	flagVarsSlice := []*string{jobDescriptionFlag, jobLocationFlag}

	for _, j := range flagVarsSlice {
		flagSanitizer(j)
	}

	resp := queryAPIEndpoint(jobDescriptionFlag, jobLocationFlag)

	rJobs := []fetchJobs{}

	err := json.Unmarshal([]byte(resp), &rJobs)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range rJobs {
		i++
		fmt.Printf("%d - %v: %v\n\t%v\n", i, v.Company, v.Title, v.Location)
	}
	telebotclient.SendTelegramMessage("Tis")
}
