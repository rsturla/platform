package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	commonv1 "github.com/rsturla/platform-contracts/gen/go/common/v1"
	greeterv1 "github.com/rsturla/platform-contracts/gen/go/greeter/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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

	greeterService := &greeterService{}
	server := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(server, greeterService)

	log.Printf("Server listening on port %s\n", serverPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *greeterService) SayHello(ctx context.Context, req *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	handleExtensions(ctx)

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

func handleExtensions(ctx context.Context) {
	methodName, err := getMethodName(ctx)
	if err != nil {
		log.Printf("Failed to get method name: %v", err)
		return
	}

	desc, err := getDescriptorByName(methodName)
	if err != nil {
		log.Printf("Failed to get descriptor by name: %v", err)
		return
	}

	methodDesc := desc.(protoreflect.MethodDescriptor)

	if humansWhoCanRpc, ok := getMethodExtension(methodDesc, commonv1.E_HumansWhoCanRpc); ok {
		log.Printf("Humans who can RPC: %v\n", humansWhoCanRpc)
	}
}

func getMethodName(ctx context.Context) (string, error) {
	procedure, ok := grpc.Method(ctx)
	if !ok {
		return "", fmt.Errorf("could not get method from context")
	}
	return strings.ReplaceAll(strings.Trim(procedure, "/"), "/", "."), nil
}

func getMethodExtension(desc protoreflect.MethodDescriptor, extensionType protoreflect.ExtensionType) (interface{}, bool) {
	opts := desc.Options()
	if !proto.HasExtension(opts, extensionType) {
		return nil, false
	}
	x := proto.GetExtension(opts, extensionType)
	return x, true
}

func getDescriptorByName(name string) (protoreflect.Descriptor, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}
	return desc, nil
}
