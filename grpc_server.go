package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "example.proto" // example.proto 파일이 있는 경로
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	// 읽기 타임아웃 설정 (예: 5초)
	err := ctx.Err()
	if err != nil {
		log.Printf("Context error before deadline: %v", err)
	}
	deadline, ok := ctx.Deadline()
	if ok {
		log.Printf("Deadline: %v", deadline)
	}
	if deadline.IsZero() {
		newDeadline := time.Now().Add(5 * time.Second)
		log.Printf("Setting deadline to: %v", newDeadline)
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, newDeadline)
		defer cancel()
	}

	// 시간이 걸리는 작업 (예: 외부 API 호출, DB 조회 등)
	//time.Sleep(3 * time.Second) // 예시

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

