package main

import (
	"context"
	"log"
	"time"

	pb "readtime/example"
	"google.golang.org/grpc"
)

func main() {
	// 연결 설정 (읽기 타임아웃 10초)
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewExampleServiceClient(conn)

	// 서버에 요청 보내기 (context에 타임아웃 설정)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

