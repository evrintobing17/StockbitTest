package endpoint

import "github.com/evrintobing17/StockbitTest/internal/util/model"

type HTTPResponse struct {
	Result      []model.Movie `json:"Search"`
	TotalResult string        `json:"totalResults"`
	Response    string        `json:"Response"`
}

type SearchRequest struct {
	MovieName string
	Page      int
}

type SearchResponse struct {
	Response     []model.Movie `json:"Search"`
	ErrorMessage string        `json:"Error,omitempty"`
}
