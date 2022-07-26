package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	helloworldpb "github.com/Akashkumar-Jeyaramans/helloWorld/proto/helloworld"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

// func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("not connected to 8080 %v", err)
	}
	s := grpc.NewServer()
	//Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to register gateway:", err)
	}
	gwmux := runtime.NewServeMux()
	// gwmux := runtime.NewServeMux()
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("failed to reg gateway", err)
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

// package main

// import (
// 	"context"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"

// 	helloworldpb "github.com/Akashkumar-Jeyaramans/helloWorld/proto/helloworld"
// 	// helloworldpb "github.com/Akashkumar-Jeyaramans/helloworld"
// )

// type server struct {
// 	helloworldpb.UnimplementedGreeterServer
// }

// func NewServer() *server {
// 	return &server{}
// }

// func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
// 	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
// }

// func main() {
// 	// Create a listener on TCP port
// 	lis, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatalln("Failed to listen:", err)
// 	}

// 	// Create a gRPC server object
// 	s := grpc.NewServer()
// 	// Attach the Greeter service to the server
// 	helloworldpb.RegisterGreeterServer(s, &server{})
// 	// Serve gRPC Server
// 	log.Println("Serving gRPC on 0.0.0.0:8080")
// 	log.Fatal(s.Serve(lis))
// }
