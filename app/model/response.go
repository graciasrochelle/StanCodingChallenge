package model

type Response struct {
	Items []ResponseItem `json:"payload"`
}

type ResponseItem struct {
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Image        string  `json:"image"`
}
