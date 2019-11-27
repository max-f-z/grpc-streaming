package srvstream

import (
	"context"
	"io"
	"log"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestSrvStream(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewSrvStreamClient(conn)

	gs, err := client.SrvStreaming(context.Background(), &SrvStreamReq{
		Name: "zack",
	})

	if err != nil {
		t.Error(err.Error())
	}

	for {
		resp, err := gs.Recv()
		if err == io.EOF {
			break
		}
		t.Log(resp.GetMessage())
	}
}
