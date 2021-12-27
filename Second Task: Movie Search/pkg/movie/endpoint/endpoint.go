package endpoint

import (
	"context"
	"errors"
	"fmt"

	"github.com/evrintobing17/StockbitTest/pkg/movie"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	SearchEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc movie.Service) Set {
	return Set{
		SearchEndpoint: MakeSearchEndpoint(svc),
	}
}

var (
	OMDBAPIKEY = "faf7e5bb&s"
	URL_FORMAT = "http://www.omdbapi.com/?apikey=%s&s=%s&page=%d"
)

func MakeSearchEndpoint(srv movie.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		url := fmt.Sprintf(URL_FORMAT, OMDBAPIKEY, req.MovieName, req.Page)
		res, err := srv.Search(ctx, url, req)
		if err != nil {
			return SearchResponse{nil, err.Error()}, errors.New(err.Error())
		}
		return SearchResponse{res, ""}, nil
	}
}

func (e Set) Search(ctx context.Context) ([]Movie, error) {
	req := SearchRequest{}
	resp, err := e.SearchEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(SearchResponse)
	if getResp.ErrorMessage != "" {
		return nil, errors.New(getResp.ErrorMessage)
	}
	return getResp.Response, nil
}
