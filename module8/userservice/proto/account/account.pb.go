// Code generated by protoc-gen-go.
// source: github.com/bketelsen/microclass/module4/4main/proto/account/account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/bketelsen/microclass/module4/4main/proto/account/account.proto

It has these top-level messages:
	User
	Session
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	UpdatePasswordRequest
	UpdatePasswordResponse
	SearchRequest
	SearchResponse
	ReadSessionRequest
	ReadSessionResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type Session struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Expires  int64  `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Session) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Session) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type CreateRequest struct {
	User     *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type UpdatePasswordRequest struct {
	UserId          string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=oldPassword" json:"oldPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword" json:"confirmPassword,omitempty"`
}

func (m *UpdatePasswordRequest) Reset()                    { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()               {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdatePasswordRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdatePasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type UpdatePasswordResponse struct {
}

func (m *UpdatePasswordResponse) Reset()                    { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()               {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type SearchRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Limit    int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SearchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SearchRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ReadSessionRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *ReadSessionRequest) Reset()                    { *m = ReadSessionRequest{} }
func (m *ReadSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionRequest) ProtoMessage()               {}
func (*ReadSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ReadSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type ReadSessionResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *ReadSessionResponse) Reset()                    { *m = ReadSessionResponse{} }
func (m *ReadSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionResponse) ProtoMessage()               {}
func (*ReadSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ReadSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *LoginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LogoutRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *LogoutRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Session)(nil), "Session")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "UpdatePasswordRequest")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "UpdatePasswordResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*ReadSessionRequest)(nil), "ReadSessionRequest")
	proto.RegisterType((*ReadSessionResponse)(nil), "ReadSessionResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "LogoutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.UpdatePassword", in)
	out := new(UpdatePasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.ReadSession", in)
	out := new(ReadSessionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	ReadSession(context.Context, *ReadSessionRequest, *ReadSessionResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Account{hdlr}, opts...))
}

type Account struct {
	AccountHandler
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *Account) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *Account) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *Account) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *Account) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.AccountHandler.Search(ctx, in, out)
}

func (h *Account) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordResponse) error {
	return h.AccountHandler.UpdatePassword(ctx, in, out)
}

func (h *Account) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *Account) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AccountHandler.Logout(ctx, in, out)
}

func (h *Account) ReadSession(ctx context.Context, in *ReadSessionRequest, out *ReadSessionResponse) error {
	return h.AccountHandler.ReadSession(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/bketelsen/microclass/module4/4main/proto/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x4d, 0xe2, 0xb8, 0x69, 0xa7, 0xb5, 0x5d, 0x6d, 0xfb, 0xeb, 0xcf, 0x18, 0x10, 0xd1, 0x4a,
	0x48, 0x01, 0xd4, 0xb5, 0xd4, 0xf6, 0x02, 0x37, 0x54, 0x84, 0x54, 0x89, 0x43, 0x95, 0xaa, 0x37,
	0x2e, 0xae, 0xbd, 0x6d, 0x57, 0xd8, 0xde, 0xe0, 0xb5, 0x55, 0x0e, 0x7c, 0x15, 0xbe, 0x09, 0x1f,
	0x0e, 0xed, 0xbf, 0xc4, 0x76, 0x1b, 0xa0, 0x9c, 0xa2, 0x99, 0x79, 0xbb, 0xf3, 0x76, 0xe6, 0x3d,
	0x07, 0xce, 0x6e, 0x58, 0x7d, 0xdb, 0x5c, 0x91, 0x94, 0x17, 0xf1, 0xd5, 0x17, 0x5a, 0xd3, 0x5c,
	0xd0, 0x32, 0x2e, 0x58, 0x5a, 0xf1, 0x34, 0x4f, 0x84, 0x88, 0x0b, 0x9e, 0x35, 0x39, 0x3d, 0x89,
	0x4f, 0x8a, 0x84, 0x95, 0xf1, 0xa2, 0xe2, 0x35, 0x8f, 0x93, 0x34, 0xe5, 0x4d, 0x59, 0xdb, 0x5f,
	0xa2, 0xb2, 0xf8, 0x3b, 0x8c, 0x2f, 0x05, 0xad, 0x90, 0x0f, 0x23, 0x96, 0x85, 0xc3, 0xe9, 0x70,
	0xb6, 0x35, 0x1f, 0xb1, 0x0c, 0x45, 0xb0, 0xd9, 0x08, 0x5a, 0x95, 0x49, 0x41, 0xc3, 0x91, 0xca,
	0x2e, 0x63, 0xb4, 0x0f, 0x2e, 0x2d, 0x12, 0x96, 0x87, 0x8e, 0x2a, 0xe8, 0x00, 0x85, 0x30, 0x49,
	0x2b, 0x9a, 0xd4, 0x34, 0x0b, 0xc7, 0xd3, 0xe1, 0xcc, 0x99, 0xdb, 0x50, 0x56, 0x9a, 0x45, 0xa6,
	0x2a, 0xae, 0xae, 0x98, 0x10, 0x33, 0x98, 0x5c, 0x50, 0x21, 0x18, 0x2f, 0x1f, 0x45, 0xa0, 0xd5,
	0xca, 0xb9, 0xd7, 0x8a, 0x7e, 0x5b, 0xb0, 0x8a, 0x0a, 0x4b, 0xc2, 0x84, 0xf8, 0x23, 0x78, 0xa7,
	0x0a, 0x34, 0xa7, 0x5f, 0x1b, 0x2a, 0x6a, 0xf4, 0x04, 0xc6, 0xf2, 0x42, 0xd5, 0x72, 0xfb, 0xc8,
	0x25, 0x72, 0x0c, 0x73, 0x95, 0x92, 0xbd, 0x17, 0x89, 0x10, 0x77, 0xbc, 0xca, 0x6c, 0x6f, 0x1b,
	0xe3, 0x5d, 0xf0, 0xed, 0x3d, 0x62, 0xc1, 0x4b, 0x41, 0xf1, 0x0b, 0xf0, 0x3e, 0xd0, 0x9c, 0xae,
	0x6e, 0xee, 0x3d, 0x45, 0x1e, 0xb1, 0x00, 0x73, 0xe4, 0x39, 0x6c, 0xcf, 0x69, 0x92, 0xad, 0x3b,
	0xf0, 0x0a, 0x76, 0x74, 0x59, 0xc3, 0x7f, 0x43, 0x15, 0xbf, 0x06, 0xef, 0x52, 0x0d, 0xf3, 0xcf,
	0xcf, 0x92, 0x3c, 0x2c, 0xd6, 0xf0, 0xf8, 0x31, 0x84, 0xff, 0x74, 0xea, 0xdc, 0xbc, 0xcf, 0x5e,
	0x73, 0x00, 0x1b, 0xf2, 0xcc, 0x99, 0xa5, 0x65, 0x22, 0x34, 0x85, 0x6d, 0x9e, 0x67, 0xe7, 0xdd,
	0xe9, 0xb4, 0x53, 0x12, 0x51, 0xd2, 0xbb, 0x25, 0x42, 0x6b, 0xa4, 0x9d, 0x42, 0x33, 0x08, 0x52,
	0x5e, 0x5e, 0xb3, 0xaa, 0x58, 0xa2, 0xc6, 0x0a, 0xd5, 0x4f, 0xe3, 0x10, 0x0e, 0xfa, 0xf4, 0x0c,
	0x73, 0x0e, 0xde, 0x05, 0x4d, 0xaa, 0xf4, 0xd6, 0x12, 0x6e, 0xeb, 0x65, 0xb8, 0x4e, 0xb0, 0xa3,
	0xb6, 0x60, 0xf7, 0xc1, 0xcd, 0x59, 0xc1, 0x6a, 0xa3, 0x21, 0x1d, 0xc8, 0x87, 0xf3, 0xeb, 0x6b,
	0x41, 0x6b, 0x23, 0x20, 0x13, 0xe1, 0x43, 0xf0, 0x6d, 0x43, 0xb3, 0x95, 0xa7, 0xe0, 0xca, 0x0e,
	0x22, 0x1c, 0x4e, 0x9d, 0xd5, 0xa8, 0x75, 0x0e, 0x1f, 0x01, 0x92, 0x2b, 0x34, 0xea, 0xb6, 0x24,
	0x9f, 0xc1, 0x96, 0xd0, 0x99, 0xe5, 0x60, 0x57, 0x09, 0xfc, 0x16, 0xf6, 0x3a, 0x67, 0x4c, 0x1f,
	0x0c, 0x13, 0x83, 0x31, 0x4b, 0xdd, 0x24, 0x16, 0x62, 0x0b, 0xf8, 0x33, 0xec, 0x7c, 0xe2, 0x37,
	0xac, 0xfc, 0xf7, 0x69, 0xb4, 0x35, 0xef, 0xf4, 0x34, 0x7f, 0x0c, 0x9e, 0xb9, 0xfd, 0x11, 0x94,
	0x0e, 0xd5, 0x21, 0xde, 0xd4, 0x7f, 0xf7, 0xf8, 0x5d, 0xf0, 0x2d, 0x5c, 0x37, 0x39, 0xfa, 0xe9,
	0xc0, 0xe4, 0xbd, 0xfe, 0x58, 0xa1, 0x37, 0xb0, 0xa1, 0x5d, 0x87, 0x7c, 0xd2, 0xb1, 0x71, 0x14,
	0x90, 0x9e, 0x1d, 0x07, 0xe8, 0x25, 0x8c, 0xe5, 0x1c, 0xd1, 0x0e, 0x69, 0x99, 0x2c, 0xf2, 0x48,
	0xdb, 0x53, 0x78, 0x20, 0xef, 0xd4, 0xe2, 0x42, 0x3e, 0xe9, 0x78, 0x28, 0x0a, 0x48, 0xcf, 0x27,
	0x0a, 0xac, 0x3d, 0x8c, 0x7c, 0xd2, 0x71, 0x7b, 0x14, 0x90, 0x9e, 0xb9, 0x15, 0x58, 0x6b, 0x05,
	0xf9, 0xa4, 0xa3, 0xd2, 0x28, 0x20, 0x5d, 0x11, 0xe1, 0x01, 0x3a, 0xb5, 0xae, 0x5c, 0xfa, 0xe3,
	0x80, 0x3c, 0xe8, 0xc9, 0xe8, 0x7f, 0xb2, 0xc6, 0x0c, 0x03, 0x34, 0x03, 0x57, 0x6d, 0x08, 0x79,
	0xa4, 0xad, 0x83, 0xc8, 0x27, 0x9d, 0xc5, 0x69, 0x6e, 0x7a, 0xce, 0x48, 0xd5, 0x56, 0xfb, 0x89,
	0x02, 0xd2, 0x5d, 0x00, 0x1e, 0xa0, 0x77, 0xfa, 0x3b, 0x65, 0xbf, 0xd1, 0x7b, 0xe4, 0xbe, 0xa6,
	0xa3, 0x7d, 0xf2, 0x80, 0x68, 0xf1, 0xe0, 0x6a, 0x43, 0xfd, 0xc1, 0x1c, 0xff, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x43, 0x2d, 0xf5, 0xeb, 0xad, 0x06, 0x00, 0x00,
}