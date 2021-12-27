package movie

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

type movieService struct{}

func NewService() Service {
	return &movieService{}
}

func (m *MovieSearchService) Search(ctx context.Context, url string, req SearchRequest) ([]Movie, error) {

	var httpResp HTTPResponse

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
