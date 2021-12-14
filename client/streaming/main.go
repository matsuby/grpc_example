package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "grpc_example/genproto"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	// Timeoutを5秒に設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	spew.Dump(r)

	stream, err2 := c.StreamExample(ctx, &pb.StreamExampleRequest{Msg: "hey hey ho"})
	if err2 != nil {
		log.Fatalf("could not streamExample: %v", err2)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		spew.Dump(reply)
	}
}
