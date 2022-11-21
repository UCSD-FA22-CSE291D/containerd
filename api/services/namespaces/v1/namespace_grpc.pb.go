// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: github.com/containerd/containerd/api/services/namespaces/v1/namespace.proto

package namespaces

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NamespacesClient is the client API for Namespaces service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespacesClient interface {
	Get(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
	List(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	Create(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	Update(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error)
	Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type namespacesClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespacesClient(cc grpc.ClientConnInterface) NamespacesClient {
	return &namespacesClient{cc}
}

func (c *namespacesClient) Get(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.namespaces.v1.Namespaces/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) List(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.namespaces.v1.Namespaces/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Create(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.namespaces.v1.Namespaces/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Update(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error) {
	out := new(UpdateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.namespaces.v1.Namespaces/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.namespaces.v1.Namespaces/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespacesServer is the server API for Namespaces service.
// All implementations must embed UnimplementedNamespacesServer
// for forward compatibility
type NamespacesServer interface {
	Get(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	List(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error)
	Delete(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNamespacesServer()
}

// UnimplementedNamespacesServer must be embedded to have forward compatible implementations.
type UnimplementedNamespacesServer struct {
}

func (UnimplementedNamespacesServer) Get(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNamespacesServer) List(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNamespacesServer) Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNamespacesServer) Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNamespacesServer) Delete(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNamespacesServer) mustEmbedUnimplementedNamespacesServer() {}

// UnsafeNamespacesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespacesServer will
// result in compilation errors.
type UnsafeNamespacesServer interface {
	mustEmbedUnimplementedNamespacesServer()
}

func RegisterNamespacesServer(s grpc.ServiceRegistrar, srv NamespacesServer) {
	s.RegisterService(&Namespaces_ServiceDesc, srv)
}

func _Namespaces_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.namespaces.v1.Namespaces/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Get(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.namespaces.v1.Namespaces/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).List(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.namespaces.v1.Namespaces/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Create(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.namespaces.v1.Namespaces/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Update(ctx, req.(*UpdateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.namespaces.v1.Namespaces/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Delete(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Namespaces_ServiceDesc is the grpc.ServiceDesc for Namespaces service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Namespaces_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.namespaces.v1.Namespaces",
	HandlerType: (*NamespacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Namespaces_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Namespaces_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Namespaces_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Namespaces_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Namespaces_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/namespaces/v1/namespace.proto",
}
