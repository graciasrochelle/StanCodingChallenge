package model

type Request struct {
	Items []RequestItem `json:"payload"`
}

type RequestItem struct {
	Title        string `json:"title"`
	DRM          bool   `json:"drm"`
	Slug         string `json:"slug"`
	EpisodeCount int    `json:"episodeCount"`
	Image        Image  `json:"image"`
}

type Image struct {
	ShowImage string `json:"showImage"`
}
