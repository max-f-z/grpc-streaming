package srvstream

import (
	"sync"

	grpc "google.golang.org/grpc"
)

type srvstreamServer struct{}

// Register register the Service to gRPC Server
func Register(server *grpc.Server) {
	RegisterSrvStreamServer(server, &srvstreamServer{})
}

// SrvStreaming -
func (s *srvstreamServer) SrvStreaming(in *SrvStreamReq, gs SrvStream_SrvStreamingServer) error {
	content := []string{
		"page 1 content",
		"page 2 content",
		"page 3 content",
		"page 4 content",
		"page 5 content",
		"page 6 content",
		"page 7 content",
		"page 8 content",
		"page 9 content",
		"page 10 content",
	}

	var wg sync.WaitGroup
	wg.Add(10)

	if in != nil && in.Name != "" {
		for i, v := range content {
			go func(i int, v string) {
				gs.Send(&SrvStreamResp{
					MsgId:   int64(i),
					Message: v,
				})
				wg.Done()
			}(i, v)
		}
	}
	wg.Wait()

	return nil
}
