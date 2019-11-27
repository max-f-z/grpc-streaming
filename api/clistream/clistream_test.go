package clistream

import (
	context "context"
	"log"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestCliStream(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewCliStreamClient(conn)

	gs, err := client.CliStreaming(context.Background())
	if err != nil {
		t.Fatalf("Cannot make connection to server %s", err.Error())
	}
	names := []string{"alex", "bob", "cain", "david", "erich", "frank", "gary"}

	for i, v := range names {
		gs.Send(&CliStreamReq{
			MsgId: int64(i),
			Name:  v,
		})
	}

	resp, err := gs.CloseAndRecv()

	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(resp)
}
