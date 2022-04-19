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
	OperationState       bool      `protobuf:"varint,5,opt,name=operationState,proto3" json:"operationState,omitempty"`
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

func (m *Response) GetOperationState() bool {
	if m != nil {
		return m.OperationState
	}
	return false
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

type UserListRequest struct {
	// 是否是查询操作
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DeptID string `protobuf:"bytes,2,opt,name=deptID,proto3" json:"deptID,omitempty"`
	// 排序方式 +id -id
	Sort string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	// 第几页
	Page int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	// 每页项数
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

func (m *UserListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserListRequest) GetDeptID() string {
	if m != nil {
		return m.DeptID
	}
	return ""
}

func (m *UserListRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *UserListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type UserListResponse struct {
	UserInfoList         []*UserInfo `protobuf:"bytes,1,rep,name=userInfoList,proto3" json:"userInfoList,omitempty"`
	Total                int32       `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUserInfoList() []*UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

func (m *UserListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type OperationUserRequest struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OperationUserRequest) Reset()         { *m = OperationUserRequest{} }
func (m *OperationUserRequest) String() string { return proto.CompactTextString(m) }
func (*OperationUserRequest) ProtoMessage()    {}
func (*OperationUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *OperationUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationUserRequest.Unmarshal(m, b)
}
func (m *OperationUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationUserRequest.Marshal(b, m, deterministic)
}
func (m *OperationUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationUserRequest.Merge(m, src)
}
func (m *OperationUserRequest) XXX_Size() int {
	return xxx_messageInfo_OperationUserRequest.Size(m)
}
func (m *OperationUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperationUserRequest proto.InternalMessageInfo

func (m *OperationUserRequest) GetUserInfo() *UserInfo {
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
	return fileDescriptor_116e343673f7ffaf, []int{6}
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

type RoleControlRequest struct {
	EmployeeCode         string   `protobuf:"bytes,1,opt,name=employeeCode,proto3" json:"employeeCode,omitempty"`
	RoleIdList           []string `protobuf:"bytes,2,rep,name=roleIdList,proto3" json:"roleIdList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleControlRequest) Reset()         { *m = RoleControlRequest{} }
func (m *RoleControlRequest) String() string { return proto.CompactTextString(m) }
func (*RoleControlRequest) ProtoMessage()    {}
func (*RoleControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *RoleControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleControlRequest.Unmarshal(m, b)
}
func (m *RoleControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleControlRequest.Marshal(b, m, deterministic)
}
func (m *RoleControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleControlRequest.Merge(m, src)
}
func (m *RoleControlRequest) XXX_Size() int {
	return xxx_messageInfo_RoleControlRequest.Size(m)
}
func (m *RoleControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleControlRequest proto.InternalMessageInfo

func (m *RoleControlRequest) GetEmployeeCode() string {
	if m != nil {
		return m.EmployeeCode
	}
	return ""
}

func (m *RoleControlRequest) GetRoleIdList() []string {
	if m != nil {
		return m.RoleIdList
	}
	return nil
}

type DeptRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Is                   string   `protobuf:"bytes,2,opt,name=is,proto3" json:"is,omitempty"`
	ParentId             string   `protobuf:"bytes,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeptRequest) Reset()         { *m = DeptRequest{} }
func (m *DeptRequest) String() string { return proto.CompactTextString(m) }
func (*DeptRequest) ProtoMessage()    {}
func (*DeptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *DeptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeptRequest.Unmarshal(m, b)
}
func (m *DeptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeptRequest.Marshal(b, m, deterministic)
}
func (m *DeptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeptRequest.Merge(m, src)
}
func (m *DeptRequest) XXX_Size() int {
	return xxx_messageInfo_DeptRequest.Size(m)
}
func (m *DeptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeptRequest proto.InternalMessageInfo

func (m *DeptRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeptRequest) GetIs() string {
	if m != nil {
		return m.Is
	}
	return ""
}

func (m *DeptRequest) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*DeptResponse)(nil), "user.DeptResponse")
	proto.RegisterType((*UserListRequest)(nil), "user.UserListRequest")
	proto.RegisterType((*UserListResponse)(nil), "user.UserListResponse")
	proto.RegisterType((*OperationUserRequest)(nil), "user.OperationUserRequest")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*RoleControlRequest)(nil), "user.RoleControlRequest")
	proto.RegisterType((*DeptRequest)(nil), "user.DeptRequest")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0x8e, 0x6c, 0xcb, 0x91, 0xc7, 0x3e, 0x4e, 0xce, 0x92, 0x13, 0x84, 0x2f, 0x0e, 0x46, 0x17,
	0xc5, 0xe4, 0xc2, 0x50, 0x87, 0xd2, 0x42, 0x4b, 0xa1, 0xb1, 0x42, 0x30, 0xa4, 0xb4, 0x28, 0x09,
	0xb4, 0xa5, 0x37, 0x9b, 0x68, 0xea, 0x88, 0x4a, 0xda, 0xed, 0xee, 0x26, 0x25, 0xf4, 0x7d, 0xfa,
	0x40, 0x7d, 0x8c, 0x3e, 0x45, 0xd9, 0x5d, 0x49, 0x96, 0x13, 0x63, 0x93, 0xde, 0xed, 0xf7, 0xed,
	0xfc, 0xed, 0xcc, 0x37, 0x12, 0xc0, 0x8d, 0x44, 0x31, 0xe6, 0x82, 0x29, 0x46, 0x5a, 0xfa, 0x1c,
	0x7c, 0x84, 0xed, 0x08, 0xbf, 0xdd, 0xa0, 0x54, 0x84, 0x40, 0x2b, 0xa7, 0x19, 0xfa, 0xce, 0xd0,
	0x19, 0x75, 0x22, 0x73, 0x26, 0xff, 0x03, 0x60, 0xc6, 0x53, 0x76, 0x87, 0x38, 0x8b, 0xfd, 0x86,
	0xb9, 0xa9, 0x31, 0x64, 0x00, 0x1e, 0xa7, 0x52, 0x7e, 0x67, 0x22, 0xf6, 0x9b, 0xe6, 0xb6, 0xc2,
	0xc1, 0x4f, 0x07, 0xbc, 0x08, 0x25, 0x67, 0xb9, 0x44, 0xb2, 0x0b, 0xcd, 0x4c, 0xce, 0x8b, 0xd8,
	0xfa, 0xa8, 0x43, 0xa7, 0x6c, 0x9e, 0xe4, 0x67, 0x8a, 0x2a, 0x34, 0xa1, 0xbd, 0xa8, 0xc6, 0x90,
	0x03, 0xf0, 0x74, 0x85, 0xb3, 0xfc, 0x0b, 0x33, 0xa1, 0xbb, 0x93, 0xfe, 0xd8, 0x94, 0x7f, 0x51,
	0xb0, 0x51, 0x75, 0x4f, 0xf6, 0xc0, 0x15, 0x2c, 0x45, 0xe9, 0xb7, 0x86, 0xcd, 0x51, 0x27, 0xb2,
	0x80, 0x3c, 0x81, 0x3e, 0xe3, 0x28, 0xa8, 0x4a, 0x58, 0x91, 0xc5, 0x35, 0x59, 0xee, 0xb1, 0xc1,
	0x2b, 0xe8, 0x85, 0xc8, 0xd5, 0x9a, 0x5a, 0x07, 0xe0, 0xc5, 0xc8, 0x95, 0xa9, 0xc5, 0x36, 0xa1,
	0xc2, 0xc1, 0x0f, 0xd8, 0xd1, 0x15, 0x9d, 0x26, 0x52, 0xad, 0xeb, 0xe4, 0x3e, 0xb4, 0x8d, 0x4b,
	0x58, 0x04, 0x28, 0x90, 0xb6, 0x95, 0x4c, 0xa8, 0xa2, 0x7b, 0xe6, 0xac, 0x39, 0x4e, 0xe7, 0xe8,
	0xb7, 0x86, 0xce, 0xc8, 0x8d, 0xcc, 0x59, 0x3f, 0x31, 0x4d, 0xb2, 0x44, 0x99, 0x37, 0xb8, 0x91,
	0x05, 0xc1, 0x67, 0xd8, 0x5d, 0x24, 0x2f, 0xca, 0x9f, 0x40, 0xaf, 0x6c, 0x8c, 0xe6, 0x7d, 0x67,
	0xd8, 0x5c, 0xd1, 0xbc, 0x25, 0x1b, 0x1d, 0xfd, 0x9c, 0x29, 0x9a, 0x9a, 0xe2, 0xdc, 0xc8, 0x82,
	0xe0, 0x08, 0xf6, 0xde, 0x95, 0xad, 0xd2, 0x8e, 0xe5, 0xfb, 0xea, 0xa3, 0x71, 0xd6, 0x8f, 0x26,
	0xf8, 0xdd, 0x00, 0xaf, 0xa4, 0x49, 0x00, 0xbd, 0xe3, 0x42, 0x3c, 0x53, 0x16, 0x97, 0x0d, 0x5a,
	0xe2, 0x74, 0xa3, 0x8c, 0x7d, 0x29, 0xb7, 0x02, 0x69, 0xdf, 0x10, 0x39, 0x15, 0x2a, 0xc3, 0x5c,
	0xcd, 0x4a, 0xb9, 0x2d, 0x71, 0x7a, 0x4e, 0xda, 0xfa, 0xfc, 0x8e, 0xdb, 0xe6, 0x75, 0xa2, 0x0a,
	0x13, 0x1f, 0xb6, 0xc3, 0x44, 0xd2, 0xcb, 0xb4, 0x94, 0x41, 0x09, 0xc9, 0x18, 0xc8, 0x29, 0x95,
	0xea, 0x54, 0x6b, 0xef, 0x3c, 0xc9, 0x50, 0x2a, 0x9a, 0x71, 0xbf, 0x3d, 0x74, 0x46, 0xcd, 0x68,
	0xc5, 0x0d, 0x39, 0x80, 0xdd, 0xa9, 0x40, 0xaa, 0x30, 0x5e, 0x58, 0x6f, 0x1b, 0xeb, 0x07, 0x3c,
	0x19, 0x42, 0x37, 0x44, 0x79, 0x25, 0x12, 0xae, 0x9b, 0xe8, 0x7b, 0xa6, 0xa8, 0x3a, 0xa5, 0x5b,
	0xff, 0x26, 0x4d, 0xa8, 0xf4, 0x3b, 0xe6, 0xce, 0x02, 0xbd, 0x1d, 0x53, 0x9a, 0x4f, 0xaf, 0x69,
	0x3e, 0xc7, 0xd8, 0x07, 0xbb, 0x1d, 0x0b, 0x46, 0x7b, 0xbd, 0xbf, 0x66, 0x39, 0xfa, 0x5d, 0xeb,
	0x65, 0x40, 0xf0, 0x01, 0x48, 0xc4, 0x52, 0x9c, 0xb2, 0x5c, 0x09, 0x96, 0x96, 0xe3, 0x0a, 0xa0,
	0x87, 0x2b, 0xba, 0x5e, 0xe7, 0x74, 0x3e, 0xbd, 0x34, 0xb3, 0xd8, 0x48, 0xa6, 0x61, 0xd6, 0xa8,
	0xc6, 0x04, 0x6f, 0xf5, 0x3b, 0xf8, 0x5a, 0x85, 0xf7, 0xa1, 0x91, 0xc8, 0x62, 0x68, 0x8d, 0x44,
	0xda, 0x6f, 0x83, 0xa8, 0x0f, 0xab, 0xc2, 0x93, 0x5f, 0x2e, 0x74, 0xb5, 0x44, 0xce, 0x50, 0xdc,
	0x26, 0x57, 0x48, 0x46, 0xe0, 0x9e, 0x08, 0x44, 0x45, 0xfe, 0xb1, 0x42, 0x2a, 0xf2, 0x0c, 0xfa,
	0x25, 0xb4, 0xda, 0x0e, 0xb6, 0xb4, 0xa5, 0x19, 0xc7, 0x66, 0xcb, 0x31, 0x74, 0x4f, 0x50, 0x55,
	0xda, 0xdb, 0x68, 0x7f, 0x04, 0x3b, 0x35, 0x7b, 0xb3, 0x16, 0xff, 0x2d, 0x64, 0x5d, 0xdb, 0xef,
	0xc1, 0xfe, 0x7d, 0xba, 0x8a, 0x31, 0x31, 0x39, 0xc3, 0xe2, 0xdb, 0x70, 0x3f, 0x27, 0xb1, 0xb0,
	0xfe, 0xb1, 0x09, 0xb6, 0xc8, 0x33, 0x20, 0x56, 0x36, 0x4b, 0x6b, 0xb0, 0xb1, 0xdc, 0xd7, 0xd0,
	0x5f, 0x76, 0x23, 0x03, 0x6b, 0xb3, 0x6a, 0x65, 0x57, 0xfb, 0x5f, 0xf0, 0xf8, 0xef, 0xfd, 0x9f,
	0x42, 0x3f, 0xc4, 0x14, 0x6b, 0xfe, 0x1b, 0x4b, 0x7e, 0x01, 0x10, 0xe1, 0x2d, 0xfb, 0x8a, 0x5a,
	0xa4, 0xc4, 0x2f, 0xee, 0x1f, 0x08, 0x76, 0x85, 0xe7, 0x73, 0xe8, 0x9c, 0x08, 0x9a, 0xab, 0x47,
	0x3b, 0x1e, 0x82, 0x77, 0x1c, 0x27, 0x66, 0x22, 0xe4, 0xdf, 0x7a, 0xfb, 0xd7, 0x4f, 0x04, 0xec,
	0xd3, 0x1e, 0xe5, 0x76, 0xe4, 0x7d, 0x6a, 0x8f, 0x5f, 0xea, 0x8b, 0xcb, 0xb6, 0xf9, 0xc5, 0x1e,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x72, 0x60, 0xc0, 0x98, 0x70, 0x07, 0x00, 0x00,
}
