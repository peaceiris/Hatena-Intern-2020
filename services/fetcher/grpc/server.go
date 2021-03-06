package grpc

import (
	"context"

	fetcher "github.com/peaceiris/Hatena-Intern-2020/services/fetcher/fetcher"
	pb "github.com/peaceiris/Hatena-Intern-2020/services/fetcher/pb/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.FetcherServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetch は受け取った URL から title を取得する
func (s *Server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchReply, error) {
	title, err := fetcher.Fetch(ctx, in.Src)
	if err != nil {
		return nil, err
	}
	return &pb.FetchReply{Title: title}, nil
}
