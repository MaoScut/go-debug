package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/MaoScut/go-debug/sql-conn/server/gen/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) ListUser(context.Context, *pb.ListUserReq) (*pb.ListUserRes, error) {
	return &pb.ListUserRes{
		User: []*pb.User{
			{
				Id:   "1",
				Name: "madp",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
