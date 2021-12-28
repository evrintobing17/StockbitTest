package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/evrintobing17/StockbitTest/pkg/movie/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

var (
	ErrUnknown = errors.New("unknown argument passed")

	ErrInvalidArgument = errors.New("invalid argument passed")
)

func NewHTTPHandler(ep endpoint.Set) http.Handler {
	h := http.NewServeMux()

	h.Handle("/search", httptransport.NewServer(
		ep.SearchEndpoint,
		decodeHTTPSearchMovieRequest,
		encodeResponse,
	))
	return h
}

func decodeHTTPSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	title := r.URL.Query().Get("title")
	if title == "" {
		return nil, errors.New("title must be provided via QueryParams")
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	req := endpoint.SearchRequest{
		MovieName: title,
		Page:      page,
	}
	return req, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
