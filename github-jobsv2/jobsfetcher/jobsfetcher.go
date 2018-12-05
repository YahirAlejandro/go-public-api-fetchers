package jobsfetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type githubJob struct {
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

func queryAPIEndpoint(jobDescriptionFlag, jobLocationFlag *string) []byte {
	baseurl := "https://jobs.github.com/positions.json?"
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

func FetchSource() {

}
