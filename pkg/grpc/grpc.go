package grpc

// import (
// 	"fmt"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// type server struct{}

// func Start() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterJWTServiceServer(s, &server{})

// 	// Регистрация reflection service на сервере gRPC
// 	reflection.Register(s)
// 	fmt.Println("gRPC server is running on port :50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
