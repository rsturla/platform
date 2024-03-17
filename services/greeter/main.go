package main

import (
	"context"
	"fmt"
	commonpb "github.com/rsturla/platform-contracts/gen/go/common/v1"
	greeterv1 "github.com/rsturla/platform-contracts/gen/go/greeter/v1"
	"github.com/rsturla/platform/libs/grpc-helpers/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	serverPort = ":8080"
)

type greeterService struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func main() {
	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	greeterSvc := &greeterService{}
	server := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(server, greeterSvc)

	// Add reflection to the server
	reflection.Register(server)

	log.Printf("Server listening on port %s\n", serverPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *greeterService) SayHello(ctx context.Context, request *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	humansWhoCanRpc, err := options.GetMethodValue(ctx, commonpb.E_HumansWhoCanRpc)
	if err != nil {
		log.Printf("Failed to handle extensions: %v", err)
	}

	handler, err := options.GetMethodValue(ctx, commonpb.E_Handler)
	if err != nil {
		log.Printf("Failed to handle extensions: %v", err)
	}

	fmt.Println("handler: ", handler)
	fmt.Println("handler.path: ", handler.(*commonpb.Handler).Path)
	fmt.Println("humansWhoCanRpc: ", humansWhoCanRpc)

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}
