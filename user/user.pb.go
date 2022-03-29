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
	Roles                []string  `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
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

func (m *Response) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type DeptResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	DeptInfo             string   `protobuf:"bytes,2,opt,name=deptInfo,proto3" json:"deptInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeptResponse) Reset()         { *m = DeptResponse{} }
func (m *DeptResponse) String() string { return proto.CompactTextString(m) }
func (*DeptResponse) ProtoMessage()    {}
func (*DeptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *DeptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeptResponse.Unmarshal(m, b)
}
func (m *DeptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeptResponse.Marshal(b, m, deterministic)
}
func (m *DeptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeptResponse.Merge(m, src)
}
func (m *DeptResponse) XXX_Size() int {
	return xxx_messageInfo_DeptResponse.Size(m)
}
func (m *DeptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeptResponse proto.InternalMessageInfo

func (m *DeptResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DeptResponse) GetDeptInfo() string {
	if m != nil {
		return m.DeptInfo
	}
	return ""
}

type UserInfo struct {
	EmployeeCode         string   `protobuf:"bytes,1,opt,name=EmployeeCode,proto3" json:"EmployeeCode,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	DepartmentId         string   `protobuf:"bytes,3,opt,name=DepartmentId,proto3" json:"DepartmentId,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=UserType,proto3" json:"UserType,omitempty"`
	Disable              bool     `protobuf:"varint,5,opt,name=Disable,proto3" json:"Disable,omitempty"`
	LastLoginTimestamp   int64    `protobuf:"varint,6,opt,name=LastLoginTimestamp,proto3" json:"LastLoginTimestamp,omitempty"`
	CreatedTimestamp     int64    `protobuf:"varint,7,opt,name=CreatedTimestamp,proto3" json:"CreatedTimestamp,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	Alias                string   `protobuf:"bytes,9,opt,name=Alias,proto3" json:"Alias,omitempty"`
	CanChanged           bool     `protobuf:"varint,10,opt,name=CanChanged,proto3" json:"CanChanged,omitempty"`
	Phone                string   `protobuf:"bytes,11,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
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

func (m *UserInfo) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *UserInfo) GetLastLoginTimestamp() int64 {
	if m != nil {
		return m.LastLoginTimestamp
	}
	return 0
}

func (m *UserInfo) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
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

func (m *UserInfo) GetCanChanged() bool {
	if m != nil {
		return m.CanChanged
	}
	return false
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
	proto.RegisterType((*DeptResponse)(nil), "user.DeptResponse")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x65, 0xbb, 0x5f, 0xde, 0x49, 0xa9, 0xaa, 0x11, 0x42, 0xd6, 0x1e, 0xd0, 0x2a, 0xa7, 0xa8,
	0x87, 0x1c, 0xca, 0x11, 0x2e, 0x90, 0xa0, 0x6a, 0xa5, 0x1e, 0x50, 0x5a, 0x0e, 0x70, 0x73, 0x9b,
	0x61, 0x1b, 0x29, 0xb1, 0x8d, 0xed, 0x82, 0x7a, 0xe1, 0x3f, 0xf0, 0x37, 0xf9, 0x15, 0xc8, 0x76,
	0x92, 0x6e, 0x05, 0x62, 0x6f, 0x7e, 0x6f, 0x9e, 0xc7, 0x33, 0xef, 0x25, 0x00, 0xf7, 0x96, 0x4c,
	0xae, 0x8d, 0x72, 0x0a, 0x67, 0xfe, 0x9c, 0x7e, 0x86, 0x65, 0x45, 0xdf, 0xee, 0xc9, 0x3a, 0x44,
	0x98, 0x49, 0xd1, 0x11, 0x9f, 0x6c, 0x26, 0xd9, 0xaa, 0x0a, 0x67, 0x7c, 0x05, 0x40, 0x9d, 0x6e,
	0xd5, 0x03, 0xd1, 0xb6, 0xe6, 0x47, 0xa1, 0xb2, 0xc7, 0xe0, 0x1a, 0x98, 0x16, 0xd6, 0xfe, 0x50,
	0xa6, 0xe6, 0xd3, 0x50, 0x1d, 0x71, 0xfa, 0x13, 0x58, 0x45, 0x56, 0x2b, 0x69, 0x09, 0x4f, 0x61,
	0xda, 0xd9, 0x5d, 0xdf, 0xda, 0x1f, 0x7d, 0xe7, 0x56, 0xed, 0x1a, 0x79, 0xe5, 0x84, 0xa3, 0xd0,
	0x99, 0x55, 0x7b, 0x0c, 0x9e, 0x01, 0xf3, 0x03, 0x6e, 0xe5, 0x57, 0x15, 0x3a, 0x27, 0xe7, 0x27,
	0x79, 0x98, 0xfe, 0x53, 0xcf, 0x56, 0x63, 0x1d, 0x5f, 0xc0, 0xdc, 0xa8, 0x96, 0x2c, 0x9f, 0x6d,
	0xa6, 0xd9, 0xaa, 0x8a, 0x20, 0x7d, 0x0b, 0xc7, 0x25, 0x69, 0xf7, 0x9f, 0x19, 0xd6, 0xc0, 0x6a,
	0xd2, 0x2e, 0xbc, 0x11, 0x77, 0x1b, 0x71, 0xfa, 0xfb, 0x08, 0xd8, 0xf0, 0x14, 0xa6, 0x70, 0xfc,
	0xa1, 0x5f, 0xba, 0x50, 0xf5, 0x60, 0xd1, 0x13, 0x0e, 0x5f, 0xc2, 0x22, 0xe8, 0x07, 0x9b, 0x7a,
	0xe4, 0xef, 0x96, 0xa4, 0x85, 0x71, 0x1d, 0x49, 0xb7, 0x1d, 0x6c, 0x7a, 0xc2, 0xf9, 0x41, 0xbc,
	0xfa, 0xfa, 0x41, 0x13, 0x9f, 0xc5, 0x41, 0x06, 0x8c, 0x1c, 0x96, 0x65, 0x63, 0xc5, 0x4d, 0x4b,
	0x7c, 0x1e, 0x5c, 0x1a, 0x20, 0xe6, 0x80, 0x97, 0xc2, 0xba, 0x4b, 0x6f, 0xda, 0x75, 0xd3, 0x91,
	0x75, 0xa2, 0xd3, 0x7c, 0xb1, 0x99, 0x64, 0xd3, 0xea, 0x1f, 0x15, 0x3c, 0x83, 0xd3, 0xc2, 0x90,
	0x70, 0x54, 0x3f, 0xaa, 0x97, 0x41, 0xfd, 0x17, 0x8f, 0x1b, 0x48, 0x4a, 0xb2, 0xb7, 0xa6, 0xd1,
	0xae, 0x51, 0x92, 0xb3, 0x30, 0xd4, 0x3e, 0xe5, 0x4d, 0x7f, 0xd7, 0x36, 0xc2, 0xf2, 0x55, 0xa8,
	0x45, 0xe0, 0x63, 0x2d, 0x84, 0x2c, 0xee, 0x84, 0xdc, 0x51, 0xcd, 0x21, 0xc6, 0xfa, 0xc8, 0xf8,
	0x5b, 0x1f, 0xef, 0x94, 0x24, 0x9e, 0xc4, 0x5b, 0x01, 0x9c, 0xff, 0x9a, 0x40, 0xe2, 0xd3, 0xbc,
	0x22, 0xf3, 0xbd, 0xb9, 0x25, 0xcc, 0x60, 0x7e, 0x61, 0x88, 0x1c, 0x3e, 0x8f, 0x99, 0xf7, 0x9f,
	0xe8, 0xfa, 0x64, 0x80, 0x31, 0xd2, 0xf4, 0x99, 0x57, 0x86, 0x2d, 0x0f, 0x2b, 0x73, 0x48, 0x2e,
	0xc8, 0x8d, 0x91, 0x1e, 0xd2, 0xbf, 0x67, 0x5f, 0x16, 0xf9, 0x1b, 0x4f, 0xde, 0x2c, 0xc2, 0x0f,
	0xf3, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xeb, 0xbd, 0x6d, 0x3e, 0x03, 0x00, 0x00,
}
