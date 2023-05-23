package main

import (
	"context"
	"log"
	"net"

	pb "github.com/quangtm-hblab/core-server/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("received: %v %v", in.GetNum1(), in.GetNum2())
	result := in.Num1 + in.GetNum2()
	return &pb.SumResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("err when create listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
