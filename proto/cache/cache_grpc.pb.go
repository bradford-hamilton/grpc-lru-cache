// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: proto/cache/cache.proto

package proto_cache

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

const (
	CacheService_Get_FullMethodName      = "/proto.cache.CacheService/Get"
	CacheService_Set_FullMethodName      = "/proto.cache.CacheService/Set"
	CacheService_GetKeys_FullMethodName  = "/proto.cache.CacheService/GetKeys"
	CacheService_GetFirst_FullMethodName = "/proto.cache.CacheService/GetFirst"
	CacheService_GetLast_FullMethodName  = "/proto.cache.CacheService/GetLast"
	CacheService_Flush_FullMethodName    = "/proto.cache.CacheService/Flush"
	CacheService_Cap_FullMethodName      = "/proto.cache.CacheService/Cap"
	CacheService_Len_FullMethodName      = "/proto.cache.CacheService/Len"
)

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServiceClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
	Set(ctx context.Context, in *Item, opts ...grpc.CallOption) (*SetRes, error)
	GetKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeysRes, error)
	GetFirst(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFirstOrLastRes, error)
	GetLast(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFirstOrLastRes, error)
	Flush(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Cap(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CapRes, error)
	Len(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LenRes, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := c.cc.Invoke(ctx, CacheService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Set(ctx context.Context, in *Item, opts ...grpc.CallOption) (*SetRes, error) {
	out := new(SetRes)
	err := c.cc.Invoke(ctx, CacheService_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeysRes, error) {
	out := new(KeysRes)
	err := c.cc.Invoke(ctx, CacheService_GetKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetFirst(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFirstOrLastRes, error) {
	out := new(GetFirstOrLastRes)
	err := c.cc.Invoke(ctx, CacheService_GetFirst_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetLast(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFirstOrLastRes, error) {
	out := new(GetFirstOrLastRes)
	err := c.cc.Invoke(ctx, CacheService_GetLast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Flush(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CacheService_Flush_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Cap(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CapRes, error) {
	out := new(CapRes)
	err := c.cc.Invoke(ctx, CacheService_Cap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Len(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LenRes, error) {
	out := new(LenRes)
	err := c.cc.Invoke(ctx, CacheService_Len_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations should embed UnimplementedCacheServiceServer
// for forward compatibility
type CacheServiceServer interface {
	Get(context.Context, *GetReq) (*GetRes, error)
	Set(context.Context, *Item) (*SetRes, error)
	GetKeys(context.Context, *Empty) (*KeysRes, error)
	GetFirst(context.Context, *Empty) (*GetFirstOrLastRes, error)
	GetLast(context.Context, *Empty) (*GetFirstOrLastRes, error)
	Flush(context.Context, *Empty) (*Empty, error)
	Cap(context.Context, *Empty) (*CapRes, error)
	Len(context.Context, *Empty) (*LenRes, error)
}

// UnimplementedCacheServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (UnimplementedCacheServiceServer) Get(context.Context, *GetReq) (*GetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServiceServer) Set(context.Context, *Item) (*SetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheServiceServer) GetKeys(context.Context, *Empty) (*KeysRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedCacheServiceServer) GetFirst(context.Context, *Empty) (*GetFirstOrLastRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirst not implemented")
}
func (UnimplementedCacheServiceServer) GetLast(context.Context, *Empty) (*GetFirstOrLastRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLast not implemented")
}
func (UnimplementedCacheServiceServer) Flush(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedCacheServiceServer) Cap(context.Context, *Empty) (*CapRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cap not implemented")
}
func (UnimplementedCacheServiceServer) Len(context.Context, *Empty) (*LenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Len not implemented")
}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Set(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetKeys(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetFirst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetFirst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetFirst_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetFirst(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetLast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetLast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetLast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetLast(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Flush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Flush(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Cap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Cap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Cap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Cap(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Len_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Len(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Len_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Len(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.cache.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheService_Set_Handler,
		},
		{
			MethodName: "GetKeys",
			Handler:    _CacheService_GetKeys_Handler,
		},
		{
			MethodName: "GetFirst",
			Handler:    _CacheService_GetFirst_Handler,
		},
		{
			MethodName: "GetLast",
			Handler:    _CacheService_GetLast_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _CacheService_Flush_Handler,
		},
		{
			MethodName: "Cap",
			Handler:    _CacheService_Cap_Handler,
		},
		{
			MethodName: "Len",
			Handler:    _CacheService_Len_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cache/cache.proto",
}