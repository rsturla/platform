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
	methodName := protoreflect.FullName("greeter.v1.GreeterService.SayHello")
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(methodName)
	if err != nil {
		return nil, err
	}
	methodDesc := desc.(protoreflect.MethodDescriptor)

	if humansWhoCanRpc, ok := getHumansWhoCanRpcExtension(methodDesc); ok {
		// Use humansWhoCanRpc as needed
		fmt.Println("Humans who can RPC:", humansWhoCanRpc)
	}

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

func getHumansWhoCanRpcExtension(desc protoreflect.MethodDescriptor) (commonv1.HumansWhoCanRpc, bool) {
	opts := desc.Options()
	if !proto.HasExtension(opts, commonv1.E_HumansWhoCanRpc) {
		return commonv1.HumansWhoCanRpc_HUMANS_WHO_CAN_RPC_UNSPECIFIED, false
	}
	x := proto.GetExtension(opts, commonv1.E_HumansWhoCanRpc)
	return x.(commonv1.HumansWhoCanRpc), true
}
