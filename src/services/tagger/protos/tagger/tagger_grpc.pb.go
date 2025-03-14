// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tagger

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TaggerClient is the client API for Tagger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaggerClient interface {
	// sends back llm'd tags
	TagText(ctx context.Context, in *UnstructuredText, opts ...grpc.CallOption) (*TagReply, error)
}

type taggerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaggerClient(cc grpc.ClientConnInterface) TaggerClient {
	return &taggerClient{cc}
}

func (c *taggerClient) TagText(ctx context.Context, in *UnstructuredText, opts ...grpc.CallOption) (*TagReply, error) {
	out := new(TagReply)
	err := c.cc.Invoke(ctx, "/tagger.Tagger/TagText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaggerServer is the server API for Tagger service.
// All implementations must embed UnimplementedTaggerServer
// for forward compatibility
type TaggerServer interface {
	// sends back llm'd tags
	TagText(context.Context, *UnstructuredText) (*TagReply, error)
	mustEmbedUnimplementedTaggerServer()
}

// UnimplementedTaggerServer must be embedded to have forward compatible implementations.
type UnimplementedTaggerServer struct {
}

func (UnimplementedTaggerServer) TagText(context.Context, *UnstructuredText) (*TagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagText not implemented")
}
func (UnimplementedTaggerServer) mustEmbedUnimplementedTaggerServer() {}

// UnsafeTaggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaggerServer will
// result in compilation errors.
type UnsafeTaggerServer interface {
	mustEmbedUnimplementedTaggerServer()
}

func RegisterTaggerServer(s *grpc.Server, srv TaggerServer) {
	s.RegisterService(&_Tagger_serviceDesc, srv)
}

func _Tagger_TagText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnstructuredText)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaggerServer).TagText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tagger.Tagger/TagText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaggerServer).TagText(ctx, req.(*UnstructuredText))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tagger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tagger.Tagger",
	HandlerType: (*TaggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TagText",
			Handler:    _Tagger_TagText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tagger/tagger.proto",
}
