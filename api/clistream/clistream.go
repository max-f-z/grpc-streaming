package clistream

import (
	"io"

	grpc "google.golang.org/grpc"
)

type clistreamServer struct{}

// Register register the Service to gRPC Server
func Register(server *grpc.Server) {
	RegisterCliStreamServer(server, &clistreamServer{})
}

func (s *clistreamServer) CliStreaming(gs CliStream_CliStreamingServer) error {
	resp := &CliStreamResp{}

	for {
		r, err := gs.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}
		resp.Message = resp.Message + " " + r.GetName()
	}
	return gs.SendAndClose(resp)
}
