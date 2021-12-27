package transport

import (
	"context"

	"github.com/evrintobing17/StockbitTest/api/proto"
	"github.com/evrintobing17/StockbitTest/pkg/movie/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	movie grpctransport.Handler
}

func NewGrpcHandler(ep endpoint.Set) proto.MovieSearchServer {
	return &grpcServer{
		movie: grpctransport.NewServer(
			ep.SearchEndpoint,
			decodeGRPCSearchMovieRequest,
			decodeGRPCSearchMovieResponse,
		),
	}

}

func decodeGRPCSearchMovieRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.SearchRequest)
	return endpoint.SearchRequest{
		MovieName: req.MovieName,
		Page:      int(req.Page),
	}, nil
}

func decodeGRPCSearchMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.SearchResponse)
	result := make([]*proto.Movie, 0)

	for _, movie := range resp.Response {
		result = append(result, &proto.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			Type:   movie.Type,
			ImdbID: movie.IMDBID,
			Poster: movie.Poster,
		})
	}

	return &proto.SearchResponse{
		MovieList: result,
		Err:       resp.ErrorMessage,
	}, nil
}
func (s *grpcServer) SearchMovie(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	_, resp, err := s.movie.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.SearchResponse), nil
}
