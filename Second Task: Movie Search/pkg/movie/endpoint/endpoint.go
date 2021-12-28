package endpoint

import (
	"context"
	"errors"
	"fmt"

	"github.com/evrintobing17/StockbitTest/api/proto"
	"github.com/evrintobing17/StockbitTest/internal/util/model"
	"github.com/evrintobing17/StockbitTest/pkg/movie"
	"github.com/go-kit/kit/endpoint"
	"github.com/mitchellh/mapstructure"
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
	key = "faf7e5bb&s"
	url = "http://www.omdbapi.com/?apikey=%s&s=%s&page=%d"
)

func MakeSearchEndpoint(srv movie.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req proto.SearchRequest
		err := mapstructure.Decode(request, &req)
		if err != nil {
			return err, nil
		}
		url := fmt.Sprintf(url, key, req.MovieName, req.Page)
		res, err := srv.Search(ctx, url, req)
		if err != nil {
			return SearchResponse{nil, err.Error()}, errors.New(err.Error())
		}
		return SearchResponse{res, ""}, nil
	}
}

func (e Set) Search(ctx context.Context) ([]model.Movie, error) {
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
