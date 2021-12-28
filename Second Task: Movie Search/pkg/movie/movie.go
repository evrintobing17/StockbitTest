package movie

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/evrintobing17/StockbitTest/api/proto"
	"github.com/evrintobing17/StockbitTest/internal/util/model"
	"github.com/go-kit/kit/log"
)

type movieService struct{}

func NewMovieSearchService() Service {
	return &movieService{}
}

func (m *movieService) Search(ctx context.Context, url string, req proto.SearchRequest) ([]model.Movie, error) {

	var httpResp model.HTTPResponse

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("http request failed. error : %s\n", err)
		return httpResp.Result, err
	}

	if response.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&httpResp)
	} else {
		return httpResp.Result, errors.New("http response is not 200")
	}

	return httpResp.Result, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
