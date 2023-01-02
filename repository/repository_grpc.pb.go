// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: repository/repository.proto

package repository

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryClient interface {
	Search(ctx context.Context, opts ...grpc.CallOption) (Repository_SearchClient, error)
}

type repositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryClient(cc grpc.ClientConnInterface) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) Search(ctx context.Context, opts ...grpc.CallOption) (Repository_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Repository_ServiceDesc.Streams[0], "/repository.Repository/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositorySearchClient{stream}
	return x, nil
}

type Repository_SearchClient interface {
	Send(*SearchCodeRequest) error
	Recv() (*SearchCodeResponse, error)
	grpc.ClientStream
}

type repositorySearchClient struct {
	grpc.ClientStream
}

func (x *repositorySearchClient) Send(m *SearchCodeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repositorySearchClient) Recv() (*SearchCodeResponse, error) {
	m := new(SearchCodeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RepositoryServer is the server API for Repository service.
// All implementations must embed UnimplementedRepositoryServer
// for forward compatibility
type RepositoryServer interface {
	Search(Repository_SearchServer) error
	mustEmbedUnimplementedRepositoryServer()
}

// UnimplementedRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryServer struct {
}

func (UnimplementedRepositoryServer) Search(Repository_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedRepositoryServer) mustEmbedUnimplementedRepositoryServer() {}

// UnsafeRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServer will
// result in compilation errors.
type UnsafeRepositoryServer interface {
	mustEmbedUnimplementedRepositoryServer()
}

func RegisterRepositoryServer(s grpc.ServiceRegistrar, srv RepositoryServer) {
	s.RegisterService(&Repository_ServiceDesc, srv)
}

func _Repository_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepositoryServer).Search(&repositorySearchServer{stream})
}

type Repository_SearchServer interface {
	Send(*SearchCodeResponse) error
	Recv() (*SearchCodeRequest, error)
	grpc.ServerStream
}

type repositorySearchServer struct {
	grpc.ServerStream
}

func (x *repositorySearchServer) Send(m *SearchCodeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repositorySearchServer) Recv() (*SearchCodeRequest, error) {
	m := new(SearchCodeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Repository_ServiceDesc is the grpc.ServiceDesc for Repository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "repository.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _Repository_Search_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "repository/repository.proto",
}