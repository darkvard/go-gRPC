package main

import (
	"testing"

	"google.golang.org/grpc"

	pb "github.com/darkvard/go-gRPC/greet/proto"
)

func TestGRPCServerRegistersProtoService(t *testing.T) {
	descriptor := pb.GreetService_ServiceDesc
	info := newGRPCServer([]grpc.ServerOption{}).GetServiceInfo()
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
