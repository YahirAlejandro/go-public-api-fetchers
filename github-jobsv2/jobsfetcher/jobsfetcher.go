package jobsfetcher

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
	const
}
