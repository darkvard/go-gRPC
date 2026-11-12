package main

import (
	"testing"

	pb "github.com/darkvard/go-gRPC/calculator/proto"
)

func TestGRPCServerRegistersProtoService(t *testing.T) {
	descriptor := pb.CalculatorService_ServiceDesc
	info := newGRPCServer().GetServiceInfo()
	service, ok := info[descriptor.ServiceName]
	if !ok {
		t.Fatalf("service %q is not registered", descriptor.ServiceName)
	}

	registeredMethods := make(map[string]struct{}, len(service.Methods))
	for _, method := range service.Methods {
		registeredMethods[method.Name] = struct{}{}
	}

	for _, method := range descriptor.Methods {
		if _, ok := registeredMethods[method.MethodName]; !ok {
			t.Errorf("method %q is not registered for service %q", method.MethodName, descriptor.ServiceName)
		}
	}
}
