package model

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type HTTPResponse struct {
	Result      []Movie `json:"Search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"Response"`
}

type SearchRequest struct {
	MovieName string
	Page      int
}

type SearchResponse struct {
	Response     []Movie `json:"Search"`
	ErrorMessage string  `json:"Error,omitempty"`
}
