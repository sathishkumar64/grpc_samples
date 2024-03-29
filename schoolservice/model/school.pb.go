// Code generated by protoc-gen-go. DO NOT EDIT.
// source: school.proto

package model

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

type School struct {
	SchoolId             string   `protobuf:"bytes,1,opt,name=schoolId,proto3" json:"schoolId,omitempty"`
	SchoolName           string   `protobuf:"bytes,2,opt,name=schoolName,proto3" json:"schoolName,omitempty"`
	EduMode              string   `protobuf:"bytes,3,opt,name=eduMode,proto3" json:"eduMode,omitempty"`
	Address              *Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Rating               float32  `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *School) Reset()         { *m = School{} }
func (m *School) String() string { return proto.CompactTextString(m) }
func (*School) ProtoMessage()    {}
func (*School) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{0}
}

func (m *School) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_School.Unmarshal(m, b)
}
func (m *School) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_School.Marshal(b, m, deterministic)
}
func (m *School) XXX_Merge(src proto.Message) {
	xxx_messageInfo_School.Merge(m, src)
}
func (m *School) XXX_Size() int {
	return xxx_messageInfo_School.Size(m)
}
func (m *School) XXX_DiscardUnknown() {
	xxx_messageInfo_School.DiscardUnknown(m)
}

var xxx_messageInfo_School proto.InternalMessageInfo

func (m *School) GetSchoolId() string {
	if m != nil {
		return m.SchoolId
	}
	return ""
}

func (m *School) GetSchoolName() string {
	if m != nil {
		return m.SchoolName
	}
	return ""
}

func (m *School) GetEduMode() string {
	if m != nil {
		return m.EduMode
	}
	return ""
}

func (m *School) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *School) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type Address struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type CreateSchoolReq struct {
	School               *School  `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSchoolReq) Reset()         { *m = CreateSchoolReq{} }
func (m *CreateSchoolReq) String() string { return proto.CompactTextString(m) }
func (*CreateSchoolReq) ProtoMessage()    {}
func (*CreateSchoolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{2}
}

func (m *CreateSchoolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchoolReq.Unmarshal(m, b)
}
func (m *CreateSchoolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchoolReq.Marshal(b, m, deterministic)
}
func (m *CreateSchoolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchoolReq.Merge(m, src)
}
func (m *CreateSchoolReq) XXX_Size() int {
	return xxx_messageInfo_CreateSchoolReq.Size(m)
}
func (m *CreateSchoolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchoolReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchoolReq proto.InternalMessageInfo

func (m *CreateSchoolReq) GetSchool() *School {
	if m != nil {
		return m.School
	}
	return nil
}

type CreateSchoolRes struct {
	School               *School  `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSchoolRes) Reset()         { *m = CreateSchoolRes{} }
func (m *CreateSchoolRes) String() string { return proto.CompactTextString(m) }
func (*CreateSchoolRes) ProtoMessage()    {}
func (*CreateSchoolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{3}
}

func (m *CreateSchoolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchoolRes.Unmarshal(m, b)
}
func (m *CreateSchoolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchoolRes.Marshal(b, m, deterministic)
}
func (m *CreateSchoolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchoolRes.Merge(m, src)
}
func (m *CreateSchoolRes) XXX_Size() int {
	return xxx_messageInfo_CreateSchoolRes.Size(m)
}
func (m *CreateSchoolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchoolRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchoolRes proto.InternalMessageInfo

func (m *CreateSchoolRes) GetSchool() *School {
	if m != nil {
		return m.School
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{4}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type ListSchoolRes struct {
	School               *School  `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSchoolRes) Reset()         { *m = ListSchoolRes{} }
func (m *ListSchoolRes) String() string { return proto.CompactTextString(m) }
func (*ListSchoolRes) ProtoMessage()    {}
func (*ListSchoolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb068902851add0, []int{5}
}

func (m *ListSchoolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSchoolRes.Unmarshal(m, b)
}
func (m *ListSchoolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSchoolRes.Marshal(b, m, deterministic)
}
func (m *ListSchoolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSchoolRes.Merge(m, src)
}
func (m *ListSchoolRes) XXX_Size() int {
	return xxx_messageInfo_ListSchoolRes.Size(m)
}
func (m *ListSchoolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSchoolRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListSchoolRes proto.InternalMessageInfo

func (m *ListSchoolRes) GetSchool() *School {
	if m != nil {
		return m.School
	}
	return nil
}

func init() {
	proto.RegisterType((*School)(nil), "model.School")
	proto.RegisterType((*Address)(nil), "model.Address")
	proto.RegisterType((*CreateSchoolReq)(nil), "model.CreateSchoolReq")
	proto.RegisterType((*CreateSchoolRes)(nil), "model.CreateSchoolRes")
	proto.RegisterType((*Void)(nil), "model.Void")
	proto.RegisterType((*ListSchoolRes)(nil), "model.ListSchoolRes")
}

func init() { proto.RegisterFile("school.proto", fileDescriptor_fdb068902851add0) }

var fileDescriptor_fdb068902851add0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x25, 0x6b, 0x3f, 0x74, 0x6a, 0x15, 0x86, 0x45, 0xc2, 0x1e, 0xa4, 0x14, 0x84, 0x9e, 0x8a,
	0x56, 0x10, 0xaf, 0xea, 0x49, 0x70, 0x3d, 0x74, 0xc1, 0x7b, 0x6c, 0x06, 0x0d, 0xec, 0x1a, 0x4d,
	0xa2, 0xe0, 0xbf, 0xf1, 0xa7, 0xca, 0x26, 0xe9, 0xaa, 0x78, 0xd1, 0x5b, 0xde, 0x7b, 0xf3, 0xe6,
	0xbd, 0x21, 0xb0, 0x6b, 0x87, 0x47, 0xad, 0x97, 0xed, 0xb3, 0xd1, 0x4e, 0x63, 0xba, 0xd2, 0x92,
	0x96, 0xf5, 0x07, 0x83, 0x6c, 0xe1, 0x79, 0x9c, 0xc1, 0x76, 0x98, 0xb8, 0x96, 0x9c, 0x55, 0xac,
	0xd9, 0xe9, 0x37, 0x18, 0x0f, 0x01, 0xc2, 0xfb, 0x56, 0xac, 0x88, 0x4f, 0xbc, 0xfa, 0x8d, 0x41,
	0x0e, 0x39, 0xc9, 0xd7, 0xb9, 0x96, 0xc4, 0xb7, 0xbc, 0x38, 0x42, 0x6c, 0x20, 0x17, 0x52, 0x1a,
	0xb2, 0x96, 0x27, 0x15, 0x6b, 0x8a, 0x6e, 0xaf, 0xf5, 0xc9, 0xed, 0x45, 0x60, 0xfb, 0x51, 0xc6,
	0x03, 0xc8, 0x8c, 0x70, 0xea, 0xe9, 0x81, 0xa7, 0x15, 0x6b, 0x26, 0x7d, 0x44, 0xf5, 0x1c, 0xf2,
	0x38, 0xbb, 0x8e, 0x19, 0x97, 0x85, 0x86, 0x1b, 0xf3, 0x14, 0x52, 0xeb, 0x84, 0x1b, 0xbb, 0x05,
	0x80, 0x08, 0xc9, 0xa0, 0xdc, 0x7b, 0xec, 0xe4, 0xdf, 0xf5, 0x39, 0xec, 0x5f, 0x19, 0x12, 0x8e,
	0xc2, 0xd9, 0x3d, 0xbd, 0xe0, 0x11, 0x64, 0xe1, 0x16, 0xbf, 0xb5, 0xe8, 0xca, 0x58, 0x31, 0x4e,
	0x44, 0xf1, 0xb7, 0xd3, 0xfe, 0xd5, 0x99, 0x41, 0x72, 0xa7, 0x95, 0xac, 0xcf, 0xa0, 0xbc, 0x51,
	0xd6, 0xfd, 0xd7, 0xdf, 0x5d, 0x42, 0x19, 0x98, 0x05, 0x99, 0x37, 0x35, 0x10, 0x9e, 0x00, 0x7c,
	0x2d, 0xc2, 0x22, 0xba, 0xd6, 0x19, 0xb3, 0x69, 0x04, 0x3f, 0x82, 0x8e, 0xd9, 0x7d, 0xe6, 0xff,
	0xfd, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf9, 0x13, 0x91, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchoolServiceClient is the client API for SchoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchoolServiceClient interface {
	ListSchool(ctx context.Context, in *Void, opts ...grpc.CallOption) (SchoolService_ListSchoolClient, error)
}

type schoolServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchoolServiceClient(cc *grpc.ClientConn) SchoolServiceClient {
	return &schoolServiceClient{cc}
}

func (c *schoolServiceClient) ListSchool(ctx context.Context, in *Void, opts ...grpc.CallOption) (SchoolService_ListSchoolClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SchoolService_serviceDesc.Streams[0], "/model.SchoolService/ListSchool", opts...)
	if err != nil {
		return nil, err
	}
	x := &schoolServiceListSchoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SchoolService_ListSchoolClient interface {
	Recv() (*ListSchoolRes, error)
	grpc.ClientStream
}

type schoolServiceListSchoolClient struct {
	grpc.ClientStream
}

func (x *schoolServiceListSchoolClient) Recv() (*ListSchoolRes, error) {
	m := new(ListSchoolRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SchoolServiceServer is the server API for SchoolService service.
type SchoolServiceServer interface {
	ListSchool(*Void, SchoolService_ListSchoolServer) error
}

// UnimplementedSchoolServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSchoolServiceServer struct {
}

func (*UnimplementedSchoolServiceServer) ListSchool(req *Void, srv SchoolService_ListSchoolServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSchool not implemented")
}

func RegisterSchoolServiceServer(s *grpc.Server, srv SchoolServiceServer) {
	s.RegisterService(&_SchoolService_serviceDesc, srv)
}

func _SchoolService_ListSchool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchoolServiceServer).ListSchool(m, &schoolServiceListSchoolServer{stream})
}

type SchoolService_ListSchoolServer interface {
	Send(*ListSchoolRes) error
	grpc.ServerStream
}

type schoolServiceListSchoolServer struct {
	grpc.ServerStream
}

func (x *schoolServiceListSchoolServer) Send(m *ListSchoolRes) error {
	return x.ServerStream.SendMsg(m)
}

var _SchoolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.SchoolService",
	HandlerType: (*SchoolServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSchool",
			Handler:       _SchoolService_ListSchool_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "school.proto",
}
