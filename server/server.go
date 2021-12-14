package main

import (
	"context"
	"fmt"
	pb "grpc_example/genproto"
	"log"
	"time"
)

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) Greet(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	fmt.Println("greeeeet!")
	return &pb.Empty{}, nil
}

// オウム返しするだけのRPC
func (s *server) GetExample(ctx context.Context, req *pb.GetExampleRequest) (*pb.GetExampleResponse, error) {
	res := &pb.GetExampleResponse{
		Res: &pb.GetExampleRequest{
			SomeDouble:         req.SomeDouble,
			SomeFloat:          req.SomeFloat,
			SomeInt32:          req.SomeInt32,
			SomeInt64:          req.SomeInt64,
			SomeUint32:         req.SomeUint32,
			SomeUint64:         req.SomeUint64,
			SomeSint32:         req.SomeSint32,
			SomeSint64:         req.SomeSint64,
			SomeFixed32:        req.SomeFixed32,
			SomeFixed64:        req.SomeFixed64,
			SomeSfixed32:       req.SomeSfixed32,
			SomeSfixed64:       req.SomeSfixed64,
			SomeBool:           req.SomeBool,
			SomeString:         req.SomeString,
			SomeBytes:          req.SomeBytes,
			SomeMessage:        &pb.GetExampleRequest_SomeMessage{Hoge: req.SomeMessage.Hoge},
			SomePartial:        &pb.SomePartial{SomeString: req.SomePartial.SomeString},
			SomeEnum:           pb.GetExampleRequest_SomeEnum(req.SomeEnum),
			SomeRepeatedString: req.SomeRepeatedString,
		},
	}

	// この辺のハンドリングはもうちょっとやり方あるかも..
	if req.GetOneofInt32() != 0 {
		res.Res.SomeOneof = &pb.GetExampleRequest_OneofInt32{
			OneofInt32: req.GetOneofInt32(),
		}
	} else {
		res.Res.SomeOneof = &pb.GetExampleRequest_OneofString{
			OneofString: req.GetOneofString(),
		}
	}

	return res, nil
}

func (s *server) StreamExample(req *pb.StreamExampleRequest, stream pb.Hello_StreamExampleServer) error {
	fmt.Printf("recieved msg: %s\n", req.Msg)
	time.Sleep(time.Second * 1)
	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("%s stream %d", req.Msg, i+1)
		res := &pb.StreamExampleResponse{Msg: msg}
		if err := stream.Send(res); err != nil {
			log.Fatalf("%v", err)
		}
		time.Sleep(time.Second)
	}
	fmt.Println("end stream")
	return nil
}
