// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: file_service.proto

package file

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

// FileStreamClient is the client API for FileStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStreamClient interface {
	SendFileName(ctx context.Context, in *FileName, opts ...grpc.CallOption) (FileStream_SendFileNameClient, error)
	ListFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FileList, error)
	GetFileInfo(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileInfo, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileStream_UploadFileClient, error)
}

type fileStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStreamClient(cc grpc.ClientConnInterface) FileStreamClient {
	return &fileStreamClient{cc}
}

func (c *fileStreamClient) SendFileName(ctx context.Context, in *FileName, opts ...grpc.CallOption) (FileStream_SendFileNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileStream_ServiceDesc.Streams[0], "/file.FileStream/SendFileName", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamSendFileNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileStream_SendFileNameClient interface {
	Recv() (*FileByte, error)
	grpc.ClientStream
}

type fileStreamSendFileNameClient struct {
	grpc.ClientStream
}

func (x *fileStreamSendFileNameClient) Recv() (*FileByte, error) {
	m := new(FileByte)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileStreamClient) ListFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/file.FileStream/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStreamClient) GetFileInfo(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileInfo, error) {
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, "/file.FileStream/GetFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStreamClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileStream_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileStream_ServiceDesc.Streams[1], "/file.FileStream/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamUploadFileClient{stream}
	return x, nil
}

type FileStream_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type fileStreamUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileStreamUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileStreamUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStreamServer is the server API for FileStream service.
// All implementations must embed UnimplementedFileStreamServer
// for forward compatibility
type FileStreamServer interface {
	SendFileName(*FileName, FileStream_SendFileNameServer) error
	ListFiles(context.Context, *Empty) (*FileList, error)
	GetFileInfo(context.Context, *FileName) (*FileInfo, error)
	UploadFile(FileStream_UploadFileServer) error
	mustEmbedUnimplementedFileStreamServer()
}

// UnimplementedFileStreamServer must be embedded to have forward compatible implementations.
type UnimplementedFileStreamServer struct {
}

func (UnimplementedFileStreamServer) SendFileName(*FileName, FileStream_SendFileNameServer) error {
	return status.Errorf(codes.Unimplemented, "method SendFileName not implemented")
}
func (UnimplementedFileStreamServer) ListFiles(context.Context, *Empty) (*FileList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileStreamServer) GetFileInfo(context.Context, *FileName) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (UnimplementedFileStreamServer) UploadFile(FileStream_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileStreamServer) mustEmbedUnimplementedFileStreamServer() {}

// UnsafeFileStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStreamServer will
// result in compilation errors.
type UnsafeFileStreamServer interface {
	mustEmbedUnimplementedFileStreamServer()
}

func RegisterFileStreamServer(s grpc.ServiceRegistrar, srv FileStreamServer) {
	s.RegisterService(&FileStream_ServiceDesc, srv)
}

func _FileStream_SendFileName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileStreamServer).SendFileName(m, &fileStreamSendFileNameServer{stream})
}

type FileStream_SendFileNameServer interface {
	Send(*FileByte) error
	grpc.ServerStream
}

type fileStreamSendFileNameServer struct {
	grpc.ServerStream
}

func (x *fileStreamSendFileNameServer) Send(m *FileByte) error {
	return x.ServerStream.SendMsg(m)
}

func _FileStream_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStreamServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileStream/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStreamServer).ListFiles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStream_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStreamServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileStream/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStreamServer).GetFileInfo(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStream_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileStreamServer).UploadFile(&fileStreamUploadFileServer{stream})
}

type FileStream_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileStreamUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileStreamUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileStreamUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStream_ServiceDesc is the grpc.ServiceDesc for FileStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileStream",
	HandlerType: (*FileStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileStream_ListFiles_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _FileStream_GetFileInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendFileName",
			Handler:       _FileStream_SendFileName_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _FileStream_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file_service.proto",
}
