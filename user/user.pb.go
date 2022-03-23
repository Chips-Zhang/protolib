// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	EmployeeId           string   `protobuf:"bytes,2,opt,name=employeeId,proto3" json:"employeeId,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetEmployeeId() string {
	if m != nil {
		return m.EmployeeId
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Msg                  string    `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	LoginState           bool      `protobuf:"varint,2,opt,name=loginState,proto3" json:"loginState,omitempty"`
	UserInfo             *UserInfo `protobuf:"bytes,3,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetLoginState() bool {
	if m != nil {
		return m.LoginState
	}
	return false
}

func (m *Response) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type UserInfo struct {
	EmployeeCode         string   `protobuf:"bytes,1,opt,name=EmployeeCode,proto3" json:"EmployeeCode,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	DepartmentId         string   `protobuf:"bytes,3,opt,name=DepartmentId,proto3" json:"DepartmentId,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=UserType,proto3" json:"UserType,omitempty"`
	Disable              int32    `protobuf:"varint,5,opt,name=Disable,proto3" json:"Disable,omitempty"`
	LastLoginTime        uint64   `protobuf:"varint,6,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,7,opt,name=CreatedTime,proto3" json:"CreatedTime,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	Alias                string   `protobuf:"bytes,9,opt,name=Alias,proto3" json:"Alias,omitempty"`
	CanChanged           int32    `protobuf:"varint,10,opt,name=CanChanged,proto3" json:"CanChanged,omitempty"`
	Phone                string   `protobuf:"bytes,11,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetEmployeeCode() string {
	if m != nil {
		return m.EmployeeCode
	}
	return ""
}

func (m *UserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserInfo) GetDepartmentId() string {
	if m != nil {
		return m.DepartmentId
	}
	return ""
}

func (m *UserInfo) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *UserInfo) GetDisable() int32 {
	if m != nil {
		return m.Disable
	}
	return 0
}

func (m *UserInfo) GetLastLoginTime() uint64 {
	if m != nil {
		return m.LastLoginTime
	}
	return 0
}

func (m *UserInfo) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *UserInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UserInfo) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *UserInfo) GetCanChanged() int32 {
	if m != nil {
		return m.CanChanged
	}
	return 0
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0xae, 0xda, 0x30,
	0x14, 0x86, 0x1b, 0x48, 0x42, 0x38, 0x29, 0xa8, 0xb2, 0xaa, 0xca, 0x62, 0xa8, 0xa2, 0xa8, 0x43,
	0xd4, 0x21, 0x03, 0x1d, 0x3b, 0xb5, 0xa1, 0x42, 0x48, 0x0c, 0x55, 0xa0, 0x43, 0xbb, 0x19, 0x72,
	0x0a, 0x91, 0x12, 0x3b, 0xb5, 0x4d, 0xaf, 0x78, 0x8c, 0xfb, 0x72, 0xf7, 0x79, 0xae, 0xec, 0x24,
	0x28, 0x4c, 0x6c, 0xfe, 0xbf, 0xf3, 0xfb, 0xe8, 0xfc, 0x3e, 0x06, 0xb8, 0x28, 0x94, 0x69, 0x23,
	0x85, 0x16, 0xc4, 0x35, 0xe7, 0xf8, 0x37, 0x4c, 0x72, 0xfc, 0x77, 0x41, 0xa5, 0x09, 0x01, 0x97,
	0xb3, 0x1a, 0xa9, 0x13, 0x39, 0xc9, 0x34, 0xb7, 0x67, 0xf2, 0x11, 0x00, 0xeb, 0xa6, 0x12, 0x57,
	0xc4, 0x4d, 0x41, 0x47, 0xb6, 0x32, 0x20, 0x64, 0x01, 0x41, 0xc3, 0x94, 0x7a, 0x12, 0xb2, 0xa0,
	0x63, 0x5b, 0xbd, 0xe9, 0xf8, 0x0c, 0x41, 0x8e, 0xaa, 0x11, 0x5c, 0x21, 0x79, 0x07, 0xe3, 0x5a,
	0x9d, 0xba, 0xd6, 0xe6, 0x68, 0x3a, 0x57, 0xe2, 0x54, 0xf2, 0x9d, 0x66, 0x1a, 0x6d, 0xe7, 0x20,
	0x1f, 0x10, 0xf2, 0x19, 0x02, 0x33, 0xe0, 0x86, 0xff, 0x15, 0xb6, 0x73, 0xb8, 0x9c, 0xa7, 0x76,
	0xfa, 0x5f, 0x1d, 0xcd, 0x6f, 0xf5, 0xf8, 0x65, 0x04, 0x41, 0x8f, 0x49, 0x0c, 0x6f, 0x7f, 0x74,
	0x03, 0x66, 0xa2, 0xe8, 0xe3, 0xdc, 0x31, 0xf2, 0x01, 0x7c, 0xeb, 0xef, 0x23, 0x75, 0xca, 0xdc,
	0x5d, 0x61, 0xc3, 0xa4, 0xae, 0x91, 0xeb, 0x4d, 0x1f, 0xe9, 0x8e, 0x99, 0xc8, 0xc6, 0xbd, 0xbf,
	0x36, 0x48, 0xdd, 0x36, 0x72, 0xaf, 0x09, 0x85, 0xc9, 0xaa, 0x54, 0xec, 0x50, 0x21, 0xf5, 0x22,
	0x27, 0xf1, 0xf2, 0x5e, 0x92, 0x4f, 0x30, 0xdb, 0x32, 0xa5, 0xb7, 0x26, 0xe0, 0xbe, 0xac, 0x91,
	0xfa, 0x91, 0x93, 0xb8, 0xf9, 0x3d, 0x24, 0x11, 0x84, 0x99, 0x44, 0xa6, 0xb1, 0xb0, 0x9e, 0x89,
	0xf5, 0x0c, 0x91, 0x71, 0xac, 0x50, 0x1d, 0x65, 0xd9, 0xe8, 0x52, 0x70, 0x1a, 0xd8, 0x01, 0x86,
	0x88, 0xbc, 0x07, 0xef, 0x5b, 0x55, 0x32, 0x45, 0xa7, 0xb6, 0xd6, 0x0a, 0xf3, 0xdc, 0x19, 0xe3,
	0xd9, 0x99, 0xf1, 0x13, 0x16, 0x14, 0xec, 0x70, 0x03, 0x62, 0x6e, 0xfd, 0x3c, 0x0b, 0x8e, 0x34,
	0x6c, 0x6f, 0x59, 0xb1, 0x7c, 0x76, 0x20, 0x34, 0xaf, 0xbc, 0x43, 0xf9, 0xbf, 0x3c, 0x22, 0x49,
	0xc0, 0x5b, 0x4b, 0x44, 0x4d, 0x66, 0xed, 0x2e, 0xba, 0xaf, 0xb3, 0x98, 0xf7, 0xb2, 0x5d, 0x77,
	0xfc, 0xc6, 0x38, 0x6d, 0xac, 0xc7, 0xce, 0x14, 0xc2, 0x35, 0xea, 0xdb, 0xfa, 0x1e, 0xf9, 0xbf,
	0x07, 0x7f, 0xfc, 0xf4, 0xab, 0x81, 0x07, 0xdf, 0x7e, 0xe4, 0x2f, 0xaf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5e, 0x22, 0x09, 0x31, 0xd6, 0x02, 0x00, 0x00,
}
