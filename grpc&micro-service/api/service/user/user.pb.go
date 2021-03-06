// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetUserByIdReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIdReq) Reset()         { *m = GetUserByIdReq{} }
func (m *GetUserByIdReq) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdReq) ProtoMessage()    {}
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdReq.Unmarshal(m, b)
}
func (m *GetUserByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdReq.Marshal(b, m, deterministic)
}
func (m *GetUserByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdReq.Merge(m, src)
}
func (m *GetUserByIdReq) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdReq.Size(m)
}
func (m *GetUserByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdReq proto.InternalMessageInfo

func (m *GetUserByIdReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserByIdRep struct {
	Data                 *User    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIdRep) Reset()         { *m = GetUserByIdRep{} }
func (m *GetUserByIdRep) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdRep) ProtoMessage()    {}
func (*GetUserByIdRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserByIdRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdRep.Unmarshal(m, b)
}
func (m *GetUserByIdRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdRep.Marshal(b, m, deterministic)
}
func (m *GetUserByIdRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdRep.Merge(m, src)
}
func (m *GetUserByIdRep) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdRep.Size(m)
}
func (m *GetUserByIdRep) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdRep.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdRep proto.InternalMessageInfo

func (m *GetUserByIdRep) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex                  int32    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserByIdReq)(nil), "user.GetUserByIdReq")
	proto.RegisterType((*GetUserByIdRep)(nil), "user.GetUserByIdRep")
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb8, 0xf8, 0xdc, 0x53,
	0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x9c, 0x2a, 0x3d, 0x53, 0x82, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0x98,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x94, 0x0c, 0xd0, 0x54,
	0x14, 0x08, 0xc9, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x82, 0xd5, 0x70, 0x1b, 0x71, 0xe9, 0x81,
	0x0d, 0x05, 0x29, 0x08, 0x02, 0x8b, 0x2b, 0xb9, 0x70, 0xb1, 0x80, 0x78, 0xe8, 0x26, 0x09, 0x49,
	0x71, 0x71, 0xe4, 0x65, 0x26, 0x67, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0xc5, 0xa9, 0x15, 0x12, 0xcc, 0x60, 0xc5, 0x20, 0xa6,
	0x91, 0x17, 0x17, 0x37, 0xc8, 0x94, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x6b, 0x2e,
	0x6e, 0x24, 0x67, 0x08, 0x89, 0x40, 0x6c, 0x45, 0x75, 0xbb, 0x14, 0x36, 0xd1, 0x02, 0x25, 0x86,
	0x24, 0x36, 0xb0, 0x97, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x40, 0x44, 0x76, 0x00,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRep, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRep, error) {
	out := new(GetUserByIdRep)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdRep, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
