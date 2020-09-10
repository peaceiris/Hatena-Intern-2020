package grpc

import (
	"context"
	"testing"

	pb "github.com/peaceiris/Hatena-Intern-2020/services/fetcher/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetch_Title(t *testing.T) {
	s := NewServer()
	url := "https://example.com/"
	extected := "Example Domain"
	reply, err := s.Fetch(context.Background(), &pb.FetchRequest{Src: url})
	assert.NoError(t, err)
	assert.Equal(t, extected, reply.Title)
}
