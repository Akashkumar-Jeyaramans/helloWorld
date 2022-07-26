// package main

// import (
// 	"context"
// 	helloworldpb "helloWorld/proto/helloworld/"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"
// )

// type server struct {
// 	helloworldpb.UnimplementedGreeterServer
// }

// func NewServer() *server {
// 	return &server{}
// }

// // func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
// func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
// 	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.fatalln("not connected to 8080 %v", err)
// 	}
// 	s := grpc.NewServer()
// 	//Attach the Greeter service to the server
// 	helloworldpb.RegisterGreeterServer(s, &server{})
// 	// Serve gRPC Server
// 	log.Println("Serving gRPC on 0.0.0.0:8080")
// 	log.Fatal(s.Serve(lis))
// }

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworldpb "helloWorld/proto/helloworld"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}