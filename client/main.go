package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpc_example/genproto"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var req = &pb.GetExampleRequest{
	SomeDouble:   1.23,
	SomeFloat:    2.34,
	SomeInt32:    -3,
	SomeInt64:    -4,
	SomeUint32:   5,
	SomeUint64:   6,
	SomeSint32:   -7,
	SomeSint64:   -8,
	SomeFixed32:  9,
	SomeFixed64:  10,
	SomeSfixed32: -11,
	SomeSfixed64: -12,
	SomeBool:     true,
	SomeString:   "some string",
	SomeBytes:    []byte("some bytes"),
	SomeMessage: &pb.GetExampleRequest_SomeMessage{
		Hoge: 13,
	},
	SomePartial: &pb.SomePartial{
		SomeString: "some partial string",
	},
	SomeEnum:           pb.GetExampleRequest_FUGA,
	SomeRepeatedString: []string{"one", "two", "three"},
	// SomeOneof:          &pb.GetExampleRequest_OneofInt32{OneofInt32: 32},
	SomeOneof: &pb.GetExampleRequest_OneofString{OneofString: "this is oneof string"},
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	spew.Dump(r)

	r2, err2 := c.GetExample(ctx, req)
	if err2 != nil {
		log.Fatalf("could not getExample: %v", err2)
	}
	spew.Dump(r2)
}
