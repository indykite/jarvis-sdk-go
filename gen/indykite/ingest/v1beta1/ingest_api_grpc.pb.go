// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: indykite/ingest/v1beta1/ingest_api.proto

package ingestv1beta1

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

// IngestAPIClient is the client API for IngestAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngestAPIClient interface {
	StreamRecords(ctx context.Context, opts ...grpc.CallOption) (IngestAPI_StreamRecordsClient, error)
}

type ingestAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewIngestAPIClient(cc grpc.ClientConnInterface) IngestAPIClient {
	return &ingestAPIClient{cc}
}

func (c *ingestAPIClient) StreamRecords(ctx context.Context, opts ...grpc.CallOption) (IngestAPI_StreamRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IngestAPI_ServiceDesc.Streams[0], "/indykite.ingest.v1beta1.IngestAPI/StreamRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingestAPIStreamRecordsClient{stream}
	return x, nil
}

type IngestAPI_StreamRecordsClient interface {
	Send(*StreamRecordsRequest) error
	Recv() (*StreamRecordsResponse, error)
	grpc.ClientStream
}

type ingestAPIStreamRecordsClient struct {
	grpc.ClientStream
}

func (x *ingestAPIStreamRecordsClient) Send(m *StreamRecordsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingestAPIStreamRecordsClient) Recv() (*StreamRecordsResponse, error) {
	m := new(StreamRecordsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IngestAPIServer is the server API for IngestAPI service.
// All implementations should embed UnimplementedIngestAPIServer
// for forward compatibility
type IngestAPIServer interface {
	StreamRecords(IngestAPI_StreamRecordsServer) error
}

// UnimplementedIngestAPIServer should be embedded to have forward compatible implementations.
type UnimplementedIngestAPIServer struct {
}

func (UnimplementedIngestAPIServer) StreamRecords(IngestAPI_StreamRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecords not implemented")
}

// UnsafeIngestAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngestAPIServer will
// result in compilation errors.
type UnsafeIngestAPIServer interface {
	mustEmbedUnimplementedIngestAPIServer()
}

func RegisterIngestAPIServer(s grpc.ServiceRegistrar, srv IngestAPIServer) {
	s.RegisterService(&IngestAPI_ServiceDesc, srv)
}

func _IngestAPI_StreamRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestAPIServer).StreamRecords(&ingestAPIStreamRecordsServer{stream})
}

type IngestAPI_StreamRecordsServer interface {
	Send(*StreamRecordsResponse) error
	Recv() (*StreamRecordsRequest, error)
	grpc.ServerStream
}

type ingestAPIStreamRecordsServer struct {
	grpc.ServerStream
}

func (x *ingestAPIStreamRecordsServer) Send(m *StreamRecordsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingestAPIStreamRecordsServer) Recv() (*StreamRecordsRequest, error) {
	m := new(StreamRecordsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IngestAPI_ServiceDesc is the grpc.ServiceDesc for IngestAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngestAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indykite.ingest.v1beta1.IngestAPI",
	HandlerType: (*IngestAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRecords",
			Handler:       _IngestAPI_StreamRecords_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "indykite/ingest/v1beta1/ingest_api.proto",
}
