package main

import (
	"context"
	"fmt"
	commonv1 "github.com/rsturla/platform-contracts/gen/go/common/v1"
	greeterv1 "github.com/rsturla/platform-contracts/gen/go/greeter/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"net"
	"strings"
)

type greeterService struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	greeterService := &greeterService{}

	s := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(s, greeterService)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *greeterService) SayHello(ctx context.Context, req *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	methodName, err := getMethodName(ctx)
	if err != nil {
		return nil, err
	}

	desc, err := getDescriptorByName(methodName)
	methodDesc := desc.(protoreflect.MethodDescriptor)

	if humansWhoCanRpc, ok := getMethodExtension(methodDesc, commonv1.E_HumansWhoCanRpc); ok {
		fmt.Println("Humans who can RPC:", humansWhoCanRpc)
	}

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
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

func getMethodExtensionByName(methodName string, extensionName string) (interface{}, bool) {
	desc, err := getDescriptorByName(methodName)
	if err != nil {
		return nil, false
	}
	methodDesc := desc.(protoreflect.MethodDescriptor)
	extensionType, err := getDescriptorByName(extensionName)
	if err != nil {
		return nil, false
	}
	return getMethodExtension(methodDesc, extensionType.(protoreflect.ExtensionType))
}

func getDescriptorByName(name string) (protoreflect.Descriptor, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}
	return desc, nil
}
