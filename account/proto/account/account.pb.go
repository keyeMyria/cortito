// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors.js if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirmation string   `protobuf:"bytes,5,opt,name=password_confirmation,json=passwordConfirmation,proto3" json:"password_confirmation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{0}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRequest) GetPasswordConfirmation() string {
	if m != nil {
		return m.PasswordConfirmation
	}
	return ""
}

type UserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserExistResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserExistResponse) Reset()         { *m = UserExistResponse{} }
func (m *UserExistResponse) String() string { return proto.CompactTextString(m) }
func (*UserExistResponse) ProtoMessage()    {}
func (*UserExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{2}
}
func (m *UserExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserExistResponse.Unmarshal(m, b)
}
func (m *UserExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserExistResponse.Marshal(b, m, deterministic)
}
func (dst *UserExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExistResponse.Merge(dst, src)
}
func (m *UserExistResponse) XXX_Size() int {
	return xxx_messageInfo_UserExistResponse.Size(m)
}
func (m *UserExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserExistResponse proto.InternalMessageInfo

func (m *UserExistResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type AuthRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{3}
}
func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (dst *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(dst, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *UserResponse `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{4}
}
func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (dst *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(dst, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetUser() *UserResponse {
	if m != nil {
		return m.User
	}
	return nil
}

type VerifyTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenRequest) Reset()         { *m = VerifyTokenRequest{} }
func (m *VerifyTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenRequest) ProtoMessage()    {}
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{5}
}
func (m *VerifyTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenRequest.Unmarshal(m, b)
}
func (m *VerifyTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenRequest.Marshal(b, m, deterministic)
}
func (dst *VerifyTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenRequest.Merge(dst, src)
}
func (m *VerifyTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenRequest.Size(m)
}
func (m *VerifyTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenRequest proto.InternalMessageInfo

func (m *VerifyTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *UserResponse `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VerifyTokenResponse) Reset()         { *m = VerifyTokenResponse{} }
func (m *VerifyTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenResponse) ProtoMessage()    {}
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{6}
}
func (m *VerifyTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenResponse.Unmarshal(m, b)
}
func (m *VerifyTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenResponse.Marshal(b, m, deterministic)
}
func (dst *VerifyTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenResponse.Merge(dst, src)
}
func (m *VerifyTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenResponse.Size(m)
}
func (m *VerifyTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenResponse proto.InternalMessageInfo

func (m *VerifyTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyTokenResponse) GetUser() *UserResponse {
	if m != nil {
		return m.User
	}
	return nil
}

type RevokeTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeTokenRequest) Reset()         { *m = RevokeTokenRequest{} }
func (m *RevokeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeTokenRequest) ProtoMessage()    {}
func (*RevokeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{7}
}
func (m *RevokeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeTokenRequest.Unmarshal(m, b)
}
func (m *RevokeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeTokenRequest.Marshal(b, m, deterministic)
}
func (dst *RevokeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeTokenRequest.Merge(dst, src)
}
func (m *RevokeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeTokenRequest.Size(m)
}
func (m *RevokeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeTokenRequest proto.InternalMessageInfo

func (m *RevokeTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RevokeTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeTokenResponse) Reset()         { *m = RevokeTokenResponse{} }
func (m *RevokeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeTokenResponse) ProtoMessage()    {}
func (*RevokeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_48d38f28d85196d7, []int{8}
}
func (m *RevokeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeTokenResponse.Unmarshal(m, b)
}
func (m *RevokeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeTokenResponse.Marshal(b, m, deterministic)
}
func (dst *RevokeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeTokenResponse.Merge(dst, src)
}
func (m *RevokeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeTokenResponse.Size(m)
}
func (m *RevokeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserRequest)(nil), "account.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "account.UserResponse")
	proto.RegisterType((*UserExistResponse)(nil), "account.UserExistResponse")
	proto.RegisterType((*AuthRequest)(nil), "account.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "account.AuthResponse")
	proto.RegisterType((*VerifyTokenRequest)(nil), "account.VerifyTokenRequest")
	proto.RegisterType((*VerifyTokenResponse)(nil), "account.VerifyTokenResponse")
	proto.RegisterType((*RevokeTokenRequest)(nil), "account.RevokeTokenRequest")
	proto.RegisterType((*RevokeTokenResponse)(nil), "account.RevokeTokenResponse")
}

func init() {
	proto.RegisterFile("proto/account/account.proto", fileDescriptor_account_48d38f28d85196d7)
}

var fileDescriptor_account_48d38f28d85196d7 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0x19, 0x55, 0x72, 0xed, 0x23, 0x53, 0xe8, 0x58, 0x2e, 0x42, 0x6e, 0xc1, 0x68, 0xe5,
	0xb6, 0xe0, 0x82, 0xbd, 0x28, 0xed, 0x22, 0xc1, 0x98, 0x84, 0xec, 0x02, 0x22, 0xf6, 0xd6, 0x28,
	0xd2, 0x98, 0x0c, 0x8e, 0x35, 0x8a, 0x66, 0x94, 0xcb, 0x0b, 0xe4, 0x19, 0xf2, 0x64, 0x79, 0x93,
	0xec, 0x83, 0x46, 0x97, 0x8c, 0x62, 0x1b, 0x62, 0x92, 0x95, 0x75, 0x6e, 0xbf, 0x3e, 0x9f, 0xff,
	0x20, 0xe8, 0xc5, 0x09, 0x13, 0xec, 0x8f, 0x1f, 0x04, 0x2c, 0x8d, 0x44, 0xf9, 0x3b, 0x94, 0x59,
	0xfc, 0xb9, 0x08, 0xdd, 0x07, 0x04, 0xe6, 0x8c, 0x93, 0xc4, 0x23, 0x57, 0x29, 0xe1, 0x02, 0x7f,
	0x01, 0x8d, 0x86, 0x36, 0xea, 0xa3, 0x81, 0xe1, 0x69, 0x34, 0xc4, 0x18, 0xf4, 0xc8, 0x5f, 0x13,
	0x5b, 0xeb, 0xa3, 0x41, 0xcb, 0x93, 0xcf, 0xd8, 0x02, 0x83, 0xac, 0x7d, 0x7a, 0x69, 0x7f, 0x92,
	0xc9, 0x3c, 0xc0, 0x0e, 0x34, 0x63, 0x9f, 0xf3, 0x1b, 0x96, 0x84, 0xb6, 0x2e, 0x0b, 0x55, 0x8c,
	0xc7, 0xd0, 0x2d, 0x9f, 0x17, 0x01, 0x8b, 0x96, 0x34, 0x59, 0xfb, 0x82, 0xb2, 0xc8, 0x36, 0x64,
	0xa3, 0x55, 0x16, 0xa7, 0x4a, 0xcd, 0xbd, 0x47, 0xd0, 0xce, 0xd1, 0x78, 0xcc, 0x22, 0x4e, 0xde,
	0xc1, 0xf6, 0x03, 0x20, 0x48, 0x88, 0x2f, 0x48, 0xb8, 0xf0, 0x45, 0x41, 0xd7, 0x2a, 0x32, 0x13,
	0x91, 0x95, 0xd3, 0x38, 0x2c, 0xcb, 0x39, 0x53, 0xab, 0xc8, 0x4c, 0x84, 0xfb, 0x1b, 0xbe, 0x66,
	0x1c, 0x47, 0xb7, 0x94, 0x8b, 0x0a, 0xe6, 0x1b, 0x34, 0x48, 0x96, 0xe0, 0x12, 0xa8, 0xe9, 0x15,
	0x91, 0x7b, 0x08, 0xe6, 0x24, 0x15, 0x17, 0xe5, 0x3e, 0x2b, 0x1e, 0xb4, 0x6b, 0x57, 0x5a, 0x7d,
	0x57, 0xee, 0x29, 0xb4, 0x73, 0x81, 0xe2, 0x45, 0x16, 0x18, 0x82, 0xad, 0x48, 0x54, 0x2a, 0xc8,
	0x00, 0xff, 0x04, 0x3d, 0xe5, 0x24, 0x91, 0xd3, 0xe6, 0xa8, 0x3b, 0x2c, 0xed, 0x55, 0x17, 0xe6,
	0xc9, 0x16, 0xf7, 0x17, 0xe0, 0x39, 0x49, 0xe8, 0xf2, 0xee, 0x2c, 0x9b, 0x54, 0xc0, 0x36, 0x65,
	0xdd, 0x39, 0x74, 0x6a, 0xbd, 0x1f, 0xc8, 0xe0, 0x91, 0x6b, 0xb6, 0x22, 0x6f, 0x60, 0xe8, 0x42,
	0xa7, 0xd6, 0x9b, 0x0b, 0x8d, 0x9e, 0x10, 0xe8, 0x99, 0x32, 0xfe, 0x07, 0x30, 0x95, 0xd6, 0xc9,
	0xc8, 0x7a, 0xf5, 0x5a, 0xa9, 0xec, 0x6c, 0x87, 0xc9, 0x46, 0x67, 0xd2, 0xd6, 0xfd, 0x47, 0xff,
	0x42, 0xf3, 0x98, 0x46, 0xe1, 0xfe, 0x83, 0x07, 0x00, 0xd5, 0xf5, 0xf0, 0x1d, 0xa3, 0x4e, 0x2d,
	0x5b, 0x3b, 0xb4, 0xd1, 0x23, 0x02, 0x3d, 0x3b, 0x08, 0xfc, 0x1f, 0xcc, 0xfc, 0x7f, 0xcb, 0xbd,
	0x28, 0x4a, 0xca, 0xbd, 0x29, 0x10, 0xb5, 0x23, 0x3a, 0x01, 0x53, 0xf1, 0x15, 0xf7, 0xaa, 0xae,
	0xcd, 0xcb, 0x70, 0xbe, 0x6f, 0x2f, 0xbe, 0x28, 0x29, 0xee, 0x28, 0x4a, 0x9b, 0xfe, 0x2a, 0x4a,
	0x5b, 0x0c, 0x3d, 0x6f, 0xc8, 0x4f, 0xd1, 0xf8, 0x39, 0x00, 0x00, 0xff, 0xff, 0x05, 0x40, 0xd5,
	0xa5, 0xa9, 0x04, 0x00, 0x00,
}
