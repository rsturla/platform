package options

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// getDescriptorByName retrieves the descriptor for a given method name.
func getDescriptorByName(name string) (protoreflect.Descriptor, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, fmt.Errorf("failed to find descriptor by name: %v", err)
	}
	return desc, nil
}
