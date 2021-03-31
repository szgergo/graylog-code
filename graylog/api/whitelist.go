package api

type WhiteListedUrl struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type WhiteListApi struct {
	Entries *[]WhiteListedUrl `json:"entries"`
	Disabled bool `json:"disabled"`
}
