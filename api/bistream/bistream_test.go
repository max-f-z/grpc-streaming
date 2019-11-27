package bistream

import (
	context "context"
	"fmt"
	"log"
	"sync"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestBiStream(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewBiStreamClient(conn)
	gs, err := client.BiStreaming(context.Background())
	if err != nil {
		t.Fatalf("Cannot make connection to server %s", err.Error())
	}

	msg, _ := gs.Recv()
	fmt.Println(msg)

	var wg sync.WaitGroup
	wg.Add(10)

	go func() {
		for {
			resp, err := gs.Recv()
			wg.Done()
			if resp != nil {
				fmt.Println(resp.Message)
			}
			if err != nil {
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go func(i int) {
			gs.Send(&BiStreamReq{
				MsgId: int64(i),
				Num:   int64(i),
			})
		}(i)
	}
	wg.Wait()
}
