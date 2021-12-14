package main

import (
	"context"
	"fmt"
	pb "grpc_example/genproto"
)

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) Greet(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	fmt.Println("greeeeet!")
	return &pb.Empty{}, nil
}

// オウム返しするだけのRPC
func (s *server) GetExample(ctx context.Context, in *pb.GetExampleRequest) (*pb.GetExampleResponse, error) {
	res := &pb.GetExampleResponse{
		SomeDouble:   in.SomeDouble,
		SomeFloat:    in.SomeFloat,
		SomeInt32:    in.SomeInt32,
		SomeInt64:    in.SomeInt64,
		SomeUint32:   in.SomeUint32,
		SomeUint64:   in.SomeUint64,
		SomeSint32:   in.SomeSint32,
		SomeSint64:   in.SomeSint64,
		SomeFixed32:  in.SomeFixed32,
		SomeFixed64:  in.SomeFixed64,
		SomeSfixed32: in.SomeSfixed32,
		SomeSfixed64: in.SomeSfixed64,
		SomeBool:     in.SomeBool,
		SomeString:   in.SomeString,
		SomeBytes:    in.SomeBytes,
		SomeMessage: &pb.GetExampleResponse_SomeMessage{
			Hoge: in.SomeMessage.Hoge,
		},
		SomePartial: &pb.SomePartial{
			SomeString: in.SomePartial.SomeString,
		},
		SomeEnum:           pb.GetExampleResponse_SomeEnum(in.SomeEnum),
		SomeRepeatedString: in.SomeRepeatedString,
	}

	// この辺のハンドリングはもうちょっとやり方あるかも..
	if in.GetOneofInt32() != 0 {
		res.SomeOneof = &pb.GetExampleResponse_OneofInt32{
			OneofInt32: in.GetOneofInt32(),
		}
	} else {
		res.SomeOneof = &pb.GetExampleResponse_OneofString{
			OneofString: in.GetOneofString(),
		}
	}

	return res, nil
}
