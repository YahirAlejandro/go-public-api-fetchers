package jobsfetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/YahirAlejandro/go-public-api-fetchers/github-jobsv2/telegrambot"
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

func checkError(src string, e error) {
	if e != nil {
		log.Fatal(src, e)
	}
}

func queryAPIEndpoint() []byte {
	baseurl := "https://jobs.github.com/positions.json?"
	jobDescription := "description=devops"
	jobLocation := "location=ny"
	url := baseurl + jobDescription + "&" + jobLocation

	req, err := http.Get(url)
	checkError("Getting JSON jobs via GET: ", err)

	resp, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	return resp
}

func unmarshalResponse(respByte []byte) {
	rJobs := []githubJob{}

	err := json.Unmarshal([]byte(respByte), &rJobs)
	checkError("Unmarshaling job response: ", err)

	var jobList string

	for i, v := range rJobs {
		i++
		// These aren't really useful until sanitized...
		//jobDescription := fmt.Sprintf("%v...", v.Description[3:160])
		//jobPostingMessage := fmt.Sprintf("%d.- %v:\n %v.\n⚐ %v ⚐\n\n\t\t • \"%v\"\n ➤ %v\n\n\n", i, v.Company, v.Title, v.Location, jobDescription, v.URL)

		jobPostingMessage := fmt.Sprintf("%d.- %v:\n • %v.\n⚐ %v ⚐\n➤ %v\n\n", i, v.Company, v.Title, v.Location, v.URL)

		jobList = jobList + jobPostingMessage
	}

	telegrambot.SendMessage(jobList)
}

func Fetch() {
	jobs := queryAPIEndpoint()
	unmarshalResponse(jobs)
}
