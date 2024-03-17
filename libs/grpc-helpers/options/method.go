package options

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetMethodValue retrieves the value of a specified extension for a given method.
func GetMethodValue(ctx context.Context, extension protoreflect.ExtensionType) (interface{}, error) {
	methodName, err := getMethodName(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get method name: %v", err)
	}

	descriptor, err := getDescriptorByName(methodName)
	if err != nil {
		return nil, fmt.Errorf("failed to get method descriptor by name: %v", err)
	}

	value, ok := getMethodExtensionValue(descriptor.(protoreflect.MethodDescriptor), extension)
	if !ok {
		return nil, fmt.Errorf("failed to get method extension value")
	}

	return value, nil
}

// getMethodName retrieves the fully qualified name of the method from the context.
func getMethodName(ctx context.Context) (string, error) {
	procedure, ok := grpc.Method(ctx)
	if !ok {
		return "", fmt.Errorf("could not get method from context")
	}
	return strings.ReplaceAll(strings.Trim(procedure, "/"), "/", "."), nil
}

// getExtensionValue retrieves the value of the specified extension for a method.
func getMethodExtensionValue(desc protoreflect.Descriptor, extension protoreflect.ExtensionType) (interface{}, bool) {
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
