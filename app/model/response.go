package model

type Response struct {
	Response []ResponseItem `json:"response"`
}

type ResponseItem struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}
