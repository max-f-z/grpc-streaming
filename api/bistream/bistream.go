package bistream

import (
	fmt "fmt"
	"io"

	grpc "google.golang.org/grpc"
)

type bistreamServer struct {
}

// Register register the Service to gRPC Server
func Register(server *grpc.Server) {
	RegisterBiStreamServer(server, &bistreamServer{})
}

// BiStreaming -
func (s *bistreamServer) BiStreaming(gs BiStream_BiStreamingServer) error {
	gs.Send(&BiStreamResp{
		Message: "input a number",
	})

	for {
		r, err := gs.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}

		gs.Send(&BiStreamResp{
			MsgId:   r.MsgId,
			Message: fmt.Sprintf("%d multiply by 2 is %d", r.Num, r.Num*2),
		})
	}
}
