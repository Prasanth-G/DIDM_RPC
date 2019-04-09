// Code generated by protoc-gen-go. DO NOT EDIT.
// source: didm.proto

package didm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DistributedDownloadRequest struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Saveto               string   `protobuf:"bytes,2,opt,name=saveto,proto3" json:"saveto,omitempty"`
	Saveas               string   `protobuf:"bytes,3,opt,name=saveas,proto3" json:"saveas,omitempty"`
	PeerIPAddr           []string `protobuf:"bytes,4,rep,name=peerIPAddr,proto3" json:"peerIPAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedDownloadRequest) Reset()         { *m = DistributedDownloadRequest{} }
func (m *DistributedDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DistributedDownloadRequest) ProtoMessage()    {}
func (*DistributedDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bdc31f56f5ef265, []int{0}
}

func (m *DistributedDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedDownloadRequest.Unmarshal(m, b)
}
func (m *DistributedDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedDownloadRequest.Marshal(b, m, deterministic)
}
func (m *DistributedDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedDownloadRequest.Merge(m, src)
}
func (m *DistributedDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DistributedDownloadRequest.Size(m)
}
func (m *DistributedDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedDownloadRequest proto.InternalMessageInfo

func (m *DistributedDownloadRequest) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *DistributedDownloadRequest) GetSaveto() string {
	if m != nil {
		return m.Saveto
	}
	return ""
}

func (m *DistributedDownloadRequest) GetSaveas() string {
	if m != nil {
		return m.Saveas
	}
	return ""
}

func (m *DistributedDownloadRequest) GetPeerIPAddr() []string {
	if m != nil {
		return m.PeerIPAddr
	}
	return nil
}

type DownloadRequest struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Start                uint64   `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  uint64   `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bdc31f56f5ef265, []int{1}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *DownloadRequest) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DownloadRequest) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

type DownloadResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Types that are valid to be assigned to AccessAnyone:
	//	*DownloadResponse_DDRequest
	//	*DownloadResponse_DRequest
	AccessAnyone         isDownloadResponse_AccessAnyone `protobuf_oneof:"access_anyone"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bdc31f56f5ef265, []int{2}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type isDownloadResponse_AccessAnyone interface {
	isDownloadResponse_AccessAnyone()
}

type DownloadResponse_DDRequest struct {
	DDRequest *DistributedDownloadRequest `protobuf:"bytes,2,opt,name=DDRequest,proto3,oneof"`
}

type DownloadResponse_DRequest struct {
	DRequest *DownloadRequest `protobuf:"bytes,3,opt,name=DRequest,proto3,oneof"`
}

func (*DownloadResponse_DDRequest) isDownloadResponse_AccessAnyone() {}

func (*DownloadResponse_DRequest) isDownloadResponse_AccessAnyone() {}

func (m *DownloadResponse) GetAccessAnyone() isDownloadResponse_AccessAnyone {
	if m != nil {
		return m.AccessAnyone
	}
	return nil
}

func (m *DownloadResponse) GetDDRequest() *DistributedDownloadRequest {
	if x, ok := m.GetAccessAnyone().(*DownloadResponse_DDRequest); ok {
		return x.DDRequest
	}
	return nil
}

func (m *DownloadResponse) GetDRequest() *DownloadRequest {
	if x, ok := m.GetAccessAnyone().(*DownloadResponse_DRequest); ok {
		return x.DRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DownloadResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DownloadResponse_DDRequest)(nil),
		(*DownloadResponse_DRequest)(nil),
	}
}

func init() {
	proto.RegisterType((*DistributedDownloadRequest)(nil), "DistributedDownloadRequest")
	proto.RegisterType((*DownloadRequest)(nil), "DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "DownloadResponse")
}

func init() { proto.RegisterFile("didm.proto", fileDescriptor_9bdc31f56f5ef265) }

var fileDescriptor_9bdc31f56f5ef265 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0x4c, 0x2c, 0xed, 0xa8, 0x34, 0x0e, 0x22, 0xa1, 0x82, 0x94, 0x9c, 0x7a, 0x0a, 0x58,
	0x8f, 0x9e, 0x94, 0x20, 0x7a, 0x10, 0x64, 0xfd, 0x00, 0xd9, 0x66, 0xe7, 0x10, 0xac, 0xbb, 0x71,
	0x67, 0x5b, 0xf1, 0xe6, 0x07, 0xf8, 0x19, 0x7e, 0xa8, 0x74, 0x6d, 0xb3, 0x62, 0x55, 0xbc, 0xbd,
	0x79, 0x3b, 0xef, 0xbd, 0x99, 0x65, 0x00, 0x54, 0xad, 0x1e, 0x8b, 0xc6, 0x1a, 0x67, 0xf2, 0xd7,
	0x08, 0x86, 0x65, 0xcd, 0xce, 0xd6, 0xd3, 0xb9, 0x23, 0x55, 0x9a, 0x67, 0x3d, 0x33, 0x52, 0x09,
	0x7a, 0x9a, 0x13, 0x3b, 0x44, 0x48, 0x66, 0xb5, 0x7e, 0xc8, 0xa2, 0x51, 0x34, 0xee, 0x0b, 0x8f,
	0xf1, 0x10, 0xba, 0x2c, 0x17, 0xe4, 0x4c, 0xb6, 0xe5, 0xd9, 0x55, 0xb5, 0xe6, 0x25, 0x67, 0x71,
	0xe0, 0x25, 0xe3, 0x31, 0x40, 0x43, 0x64, 0xaf, 0x6f, 0xcf, 0x95, 0xb2, 0x59, 0x32, 0x8a, 0xc7,
	0x7d, 0xf1, 0x85, 0xc9, 0x6f, 0x60, 0xf0, 0x9f, 0xd8, 0x03, 0xd8, 0x66, 0x27, 0xad, 0xf3, 0xa9,
	0x89, 0xf8, 0x2c, 0x30, 0x85, 0x98, 0xb4, 0xf2, 0x89, 0x89, 0x58, 0xc2, 0xfc, 0x3d, 0x82, 0x34,
	0xf8, 0x71, 0x63, 0x34, 0xd3, 0xd2, 0x50, 0x49, 0x27, 0xbd, 0xe1, 0xae, 0xf0, 0x18, 0xcf, 0xa0,
	0x5f, 0x96, 0xab, 0x44, 0x6f, 0xba, 0x33, 0x39, 0x2a, 0x7e, 0xff, 0x8b, 0xab, 0x8e, 0x08, 0xfd,
	0x58, 0x40, 0xaf, 0xd5, 0xc6, 0x5e, 0x9b, 0x16, 0x9b, 0x82, 0xb6, 0xe7, 0x62, 0x00, 0x7b, 0xb2,
	0xaa, 0x88, 0xf9, 0x5e, 0xea, 0x17, 0xa3, 0x69, 0xf2, 0x16, 0x85, 0xb5, 0xef, 0xc8, 0x2e, 0xea,
	0x8a, 0xf0, 0x12, 0x30, 0xe4, 0xaf, 0x1f, 0xf1, 0xaf, 0xa1, 0x86, 0xfb, 0xc5, 0xf7, 0x5d, 0xf3,
	0x0e, 0x9e, 0x40, 0xaf, 0x55, 0x6f, 0x8c, 0xf5, 0xa3, 0x64, 0xda, 0xf5, 0xe7, 0x70, 0xfa, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x0d, 0x1b, 0x8e, 0xe1, 0x1c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloadServiceClient interface {
	DistributeDownload(ctx context.Context, in *DistributedDownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type downloadServiceClient struct {
	cc *grpc.ClientConn
}

func NewDownloadServiceClient(cc *grpc.ClientConn) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) DistributeDownload(ctx context.Context, in *DistributedDownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/DownloadService/DistributeDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/DownloadService/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadServiceServer is the server API for DownloadService service.
type DownloadServiceServer interface {
	DistributeDownload(context.Context, *DistributedDownloadRequest) (*DownloadResponse, error)
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
}

// UnimplementedDownloadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDownloadServiceServer struct {
}

func (*UnimplementedDownloadServiceServer) DistributeDownload(ctx context.Context, req *DistributedDownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeDownload not implemented")
}
func (*UnimplementedDownloadServiceServer) Download(ctx context.Context, req *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterDownloadServiceServer(s *grpc.Server, srv DownloadServiceServer) {
	s.RegisterService(&_DownloadService_serviceDesc, srv)
}

func _DownloadService_DistributeDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributedDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).DistributeDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DownloadService/DistributeDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).DistributeDownload(ctx, req.(*DistributedDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DownloadService/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownloadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeDownload",
			Handler:    _DownloadService_DistributeDownload_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _DownloadService_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "didm.proto",
}
