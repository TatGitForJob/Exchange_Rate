// Code generated manually to keep protobuf stubs in-repo. DO NOT EDIT.

package ratesv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const RatesService_GetRates_FullMethodName = "/rates.v1.RatesService/GetRates"

// RatesServiceClient is the client API for RatesService service.
type RatesServiceClient interface {
	GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error)
}

type ratesServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewRatesServiceClient creates a typed client for RatesService.
func NewRatesServiceClient(cc grpc.ClientConnInterface) RatesServiceClient {
	return &ratesServiceClient{cc}
}

func (c *ratesServiceClient) GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error) {
	out := new(GetRatesResponse)
	err := c.cc.Invoke(ctx, RatesService_GetRates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatesServiceServer is the server API for RatesService service.
type RatesServiceServer interface {
	GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error)
	mustEmbedUnimplementedRatesServiceServer()
}

// UnimplementedRatesServiceServer should be embedded for forward compatibility.
type UnimplementedRatesServiceServer struct{}

func (UnimplementedRatesServiceServer) GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}

func (UnimplementedRatesServiceServer) mustEmbedUnimplementedRatesServiceServer() {}

// UnsafeRatesServiceServer may be embedded to opt out of forward compatibility.
type UnsafeRatesServiceServer interface {
	mustEmbedUnimplementedRatesServiceServer()
}

// RegisterRatesServiceServer registers the gRPC service implementation.
func RegisterRatesServiceServer(s grpc.ServiceRegistrar, srv RatesServiceServer) {
	s.RegisterService(&RatesService_ServiceDesc, srv)
}

func _RatesService_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatesServiceServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatesService_GetRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatesServiceServer).GetRates(ctx, req.(*GetRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatesService_ServiceDesc is the grpc.ServiceDesc for RatesService service.
var RatesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rates.v1.RatesService",
	HandlerType: (*RatesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _RatesService_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rates/v1/rates.proto",
}
