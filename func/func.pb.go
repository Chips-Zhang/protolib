// Code generated by protoc-gen-go. DO NOT EDIT.
// source: func.proto

package _func

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
	ControlUnitId        string   `protobuf:"bytes,1,opt,name=controlUnitId,proto3" json:"controlUnitId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{0}
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

func (m *Request) GetControlUnitId() string {
	if m != nil {
		return m.ControlUnitId
	}
	return ""
}

type FetchListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AdminId              string   `protobuf:"bytes,2,opt,name=adminId,proto3" json:"adminId,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchListRequest) Reset()         { *m = FetchListRequest{} }
func (m *FetchListRequest) String() string { return proto.CompactTextString(m) }
func (*FetchListRequest) ProtoMessage()    {}
func (*FetchListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{1}
}

func (m *FetchListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchListRequest.Unmarshal(m, b)
}
func (m *FetchListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchListRequest.Marshal(b, m, deterministic)
}
func (m *FetchListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchListRequest.Merge(m, src)
}
func (m *FetchListRequest) XXX_Size() int {
	return xxx_messageInfo_FetchListRequest.Size(m)
}
func (m *FetchListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchListRequest proto.InternalMessageInfo

func (m *FetchListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FetchListRequest) GetAdminId() string {
	if m != nil {
		return m.AdminId
	}
	return ""
}

func (m *FetchListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FetchListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type FetchListResponse struct {
	ControlUnitList      []*ControlUnitInfo `protobuf:"bytes,1,rep,name=controlUnitList,proto3" json:"controlUnitList,omitempty"`
	Total                int32              `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FetchListResponse) Reset()         { *m = FetchListResponse{} }
func (m *FetchListResponse) String() string { return proto.CompactTextString(m) }
func (*FetchListResponse) ProtoMessage()    {}
func (*FetchListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{2}
}

func (m *FetchListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchListResponse.Unmarshal(m, b)
}
func (m *FetchListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchListResponse.Marshal(b, m, deterministic)
}
func (m *FetchListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchListResponse.Merge(m, src)
}
func (m *FetchListResponse) XXX_Size() int {
	return xxx_messageInfo_FetchListResponse.Size(m)
}
func (m *FetchListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchListResponse proto.InternalMessageInfo

func (m *FetchListResponse) GetControlUnitList() []*ControlUnitInfo {
	if m != nil {
		return m.ControlUnitList
	}
	return nil
}

func (m *FetchListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ControlUnitInfo struct {
	ControlUnitId        string   `protobuf:"bytes,1,opt,name=controlUnitId,proto3" json:"controlUnitId,omitempty"`
	ControlUnitName      string   `protobuf:"bytes,2,opt,name=controlUnitName,proto3" json:"controlUnitName,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ControlUnitLevel     int32    `protobuf:"varint,4,opt,name=controlUnitLevel,proto3" json:"controlUnitLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlUnitInfo) Reset()         { *m = ControlUnitInfo{} }
func (m *ControlUnitInfo) String() string { return proto.CompactTextString(m) }
func (*ControlUnitInfo) ProtoMessage()    {}
func (*ControlUnitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{3}
}

func (m *ControlUnitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlUnitInfo.Unmarshal(m, b)
}
func (m *ControlUnitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlUnitInfo.Marshal(b, m, deterministic)
}
func (m *ControlUnitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlUnitInfo.Merge(m, src)
}
func (m *ControlUnitInfo) XXX_Size() int {
	return xxx_messageInfo_ControlUnitInfo.Size(m)
}
func (m *ControlUnitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlUnitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ControlUnitInfo proto.InternalMessageInfo

func (m *ControlUnitInfo) GetControlUnitId() string {
	if m != nil {
		return m.ControlUnitId
	}
	return ""
}

func (m *ControlUnitInfo) GetControlUnitName() string {
	if m != nil {
		return m.ControlUnitName
	}
	return ""
}

func (m *ControlUnitInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ControlUnitInfo) GetControlUnitLevel() int32 {
	if m != nil {
		return m.ControlUnitLevel
	}
	return 0
}

type AddControlUnitRequest struct {
	ControlUnitId        string   `protobuf:"bytes,1,opt,name=controlUnitId,proto3" json:"controlUnitId,omitempty"`
	ControlUnitName      string   `protobuf:"bytes,2,opt,name=controlUnitName,proto3" json:"controlUnitName,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ControlUnitLevel     int32    `protobuf:"varint,4,opt,name=controlUnitLevel,proto3" json:"controlUnitLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddControlUnitRequest) Reset()         { *m = AddControlUnitRequest{} }
func (m *AddControlUnitRequest) String() string { return proto.CompactTextString(m) }
func (*AddControlUnitRequest) ProtoMessage()    {}
func (*AddControlUnitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{4}
}

func (m *AddControlUnitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddControlUnitRequest.Unmarshal(m, b)
}
func (m *AddControlUnitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddControlUnitRequest.Marshal(b, m, deterministic)
}
func (m *AddControlUnitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddControlUnitRequest.Merge(m, src)
}
func (m *AddControlUnitRequest) XXX_Size() int {
	return xxx_messageInfo_AddControlUnitRequest.Size(m)
}
func (m *AddControlUnitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddControlUnitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddControlUnitRequest proto.InternalMessageInfo

func (m *AddControlUnitRequest) GetControlUnitId() string {
	if m != nil {
		return m.ControlUnitId
	}
	return ""
}

func (m *AddControlUnitRequest) GetControlUnitName() string {
	if m != nil {
		return m.ControlUnitName
	}
	return ""
}

func (m *AddControlUnitRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddControlUnitRequest) GetControlUnitLevel() int32 {
	if m != nil {
		return m.ControlUnitLevel
	}
	return 0
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1bdb44c2d3501c, []int{5}
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

func (m *Response) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "role.Request")
	proto.RegisterType((*FetchListRequest)(nil), "role.FetchListRequest")
	proto.RegisterType((*FetchListResponse)(nil), "role.FetchListResponse")
	proto.RegisterType((*ControlUnitInfo)(nil), "role.ControlUnitInfo")
	proto.RegisterType((*AddControlUnitRequest)(nil), "role.AddControlUnitRequest")
	proto.RegisterType((*Response)(nil), "role.Response")
}

func init() { proto.RegisterFile("func.proto", fileDescriptor_6b1bdb44c2d3501c) }

var fileDescriptor_6b1bdb44c2d3501c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0x25, 0xfd, 0xee, 0x2d, 0xfd, 0xb2, 0x5a, 0x88, 0xca, 0x12, 0x45, 0x0c, 0x11, 0x43, 0x2a,
	0x15, 0x36, 0x86, 0x0a, 0x8a, 0x8a, 0x22, 0x21, 0x86, 0x54, 0x2c, 0x6c, 0x69, 0x72, 0x5b, 0x5c,
	0x25, 0x76, 0x48, 0x9c, 0xce, 0xfc, 0x23, 0x26, 0xfe, 0x1f, 0xb2, 0x93, 0x22, 0x37, 0xe5, 0xe9,
	0xbd, 0x37, 0xbd, 0xed, 0xfa, 0xf8, 0x9e, 0xdc, 0xe3, 0x73, 0x6e, 0x00, 0x0e, 0x05, 0x0b, 0xdd,
	0x34, 0xe3, 0x82, 0x93, 0x56, 0xc6, 0x63, 0xb4, 0x97, 0xd0, 0xf5, 0xf1, 0x67, 0x81, 0xb9, 0x20,
	0xaf, 0x61, 0x18, 0x72, 0x26, 0x32, 0x1e, 0x7f, 0x63, 0x54, 0x78, 0x91, 0x69, 0x58, 0x86, 0xd3,
	0xf7, 0xaf, 0x41, 0xfb, 0x04, 0x93, 0x2d, 0x8a, 0xf0, 0xc7, 0x17, 0x9a, 0x8b, 0x0b, 0x93, 0x40,
	0x8b, 0x05, 0x09, 0x56, 0x04, 0x55, 0x13, 0x13, 0xba, 0x41, 0x94, 0x50, 0xe6, 0x45, 0x66, 0x43,
	0xc1, 0x97, 0xa3, 0xec, 0x4e, 0x83, 0x23, 0x9a, 0x4d, 0xcb, 0x70, 0xda, 0xbe, 0xaa, 0xc9, 0x0c,
	0xda, 0x31, 0x4d, 0xa8, 0x30, 0x5b, 0x0a, 0x2c, 0x0f, 0xf6, 0x09, 0xa6, 0xda, 0xac, 0x3c, 0xe5,
	0x2c, 0x47, 0xb2, 0x86, 0xb1, 0xa6, 0x48, 0x5e, 0x99, 0x86, 0xd5, 0x74, 0x06, 0xab, 0xb9, 0x2b,
	0x5f, 0xe4, 0x6e, 0x34, 0xb9, 0xec, 0xc0, 0xfd, 0x7a, 0xb7, 0x9c, 0x25, 0xb8, 0x08, 0x62, 0xa5,
	0xab, 0xed, 0x97, 0x07, 0xfb, 0xb7, 0x01, 0xe3, 0x1a, 0xf5, 0x61, 0x8e, 0x10, 0xe7, 0x4a, 0xd0,
	0x57, 0x69, 0x44, 0xf9, 0xe2, 0x3a, 0x4c, 0x2c, 0x18, 0x44, 0x98, 0x87, 0x19, 0x4d, 0x05, 0xe5,
	0x4c, 0x19, 0xd0, 0xf7, 0x75, 0x88, 0xbc, 0x81, 0x89, 0x2e, 0x17, 0xcf, 0x18, 0x57, 0x96, 0xdc,
	0xe0, 0xf6, 0x1f, 0x03, 0xe6, 0x1f, 0xa2, 0x48, 0x13, 0xfd, 0xa8, 0x24, 0x9f, 0x4c, 0xb7, 0x0b,
	0xbd, 0x7f, 0x61, 0x4e, 0xa0, 0x99, 0xe4, 0xc7, 0x4a, 0x9f, 0x2c, 0x25, 0xc2, 0x8a, 0xa4, 0xca,
	0x46, 0x96, 0xab, 0x5f, 0x0d, 0x18, 0xc8, 0xbd, 0xdd, 0x61, 0x76, 0xa6, 0x21, 0x12, 0x0f, 0x66,
	0x6a, 0x2b, 0x36, 0xb5, 0x5c, 0x5f, 0x94, 0xf9, 0xd7, 0xb7, 0x73, 0xf1, 0xf2, 0x06, 0x2f, 0x87,
	0xdb, 0xcf, 0xc8, 0x1a, 0x46, 0xd7, 0x0e, 0x92, 0x57, 0x65, 0xf3, 0x7f, 0x7d, 0x5d, 0x8c, 0xca,
	0x4b, 0xed, 0x03, 0x4b, 0x78, 0xfe, 0x19, 0xc5, 0xae, 0xd8, 0x6f, 0x0b, 0x16, 0x7a, 0x11, 0x19,
	0x5e, 0x3a, 0xee, 0x22, 0xbc, 0x83, 0xe9, 0x27, 0x8c, 0x51, 0xa0, 0x3e, 0xf4, 0x3e, 0xd6, 0xc7,
	0xde, 0xf7, 0x8e, 0xfb, 0x5e, 0x7a, 0xb0, 0xef, 0xa8, 0x9f, 0xf7, 0xed, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x67, 0x3a, 0xb4, 0x26, 0xca, 0x03, 0x00, 0x00,
}
