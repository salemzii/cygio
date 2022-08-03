package app

type Alert struct {
	Headline  string     `json:"headline"`
	Media     []Media    `json:"media"`
	Body      string     `json:"body"`
	Platforms []Platform `json:"platforms"`
	Urls      []Url      `json:"urls"`
}

type Url struct {
	Uri string `json:"uri"`
}

type Media struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Platform struct {
	Name string `json:"name"`
}
