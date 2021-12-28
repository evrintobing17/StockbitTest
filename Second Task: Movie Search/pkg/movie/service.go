package movie

import (
	"context"

	"github.com/evrintobing17/StockbitTest/api/proto"
	"github.com/evrintobing17/StockbitTest/internal/util/model"
)

type Service interface {
	Search(ctx context.Context, url string, req proto.SearchRequest) ([]model.Movie, error)
}
