package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strings"

	greeterv1 "github.com/rsturla/platform-contracts/gen/go/greeter/v1"
	"google.golang.org/grpc"
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

	// Add reflection to the server
	reflection.Register(server)

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
	humansWhoCanRpcFullName := protoreflect.FullName("common.v1.humans_who_can_rpc")
	humansWhoCanRpcDesc, err := protoregistry.GlobalFiles.FindDescriptorByName(humansWhoCanRpcFullName)
	if err != nil {
		log.Printf("Failed to get descriptor by name: %v", err)
		return
	}

	// Convert the descriptor to ExtensionType
	extensionType, ok := humansWhoCanRpcDesc.(protoreflect.ExtensionType)
	if !ok {
		// THIS IS ALWAYS HIT!
		log.Printf("Failed to convert descriptor to ExtensionType. Descriptor: %v", humansWhoCanRpcDesc)
		return
	}

	humansWhoCanRpc, ok := getMethodExtensionValue(methodDesc, extensionType)
	if !ok {
		log.Printf("Failed to get method extension")
		return
	}

	fmt.Println("humansWhoCanRpc: ", humansWhoCanRpc)
}

func getMethodName(ctx context.Context) (string, error) {
	procedure, ok := grpc.Method(ctx)
	if !ok {
		return "", fmt.Errorf("could not get method from context")
	}
	return strings.ReplaceAll(strings.Trim(procedure, "/"), "/", "."), nil
}

func getMethodExtensionValue(desc protoreflect.MethodDescriptor, extension protoreflect.ExtensionType) (interface{}, bool) {
	opts := desc.Options()
	if opts == nil {
		return nil, false
	}

	extensionField := proto.GetExtension(opts, extension)
	if extensionField == nil {
		return nil, false
	}

	return extensionField, true
}

func getDescriptorByName(name string) (protoreflect.Descriptor, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}
	return desc, nil
}
