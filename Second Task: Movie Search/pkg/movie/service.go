package movie

import "context"

type Service interface {
	Search(ctx context.Context, url string, req SearchRequest) ([]Movie, error)
}
