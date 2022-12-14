// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/admin/admin.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Admin service

func NewAdminEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Admin service

type AdminService interface {
	AdminLogin(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	UserList(ctx context.Context, in *UsersRequest, opts ...client.CallOption) (*UsersResponse, error)
	UserDel(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error)
}

type adminService struct {
	c    client.Client
	name string
}

func NewAdminService(name string, c client.Client) AdminService {
	return &adminService{
		c:    c,
		name: name,
	}
}

func (c *adminService) AdminLogin(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Admin.AdminLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminService) UserList(ctx context.Context, in *UsersRequest, opts ...client.CallOption) (*UsersResponse, error) {
	req := c.c.NewRequest(c.name, "Admin.UserList", in)
	out := new(UsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminService) UserDel(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error) {
	req := c.c.NewRequest(c.name, "Admin.UserDel", in)
	out := new(DelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminHandler interface {
	AdminLogin(context.Context, *LoginRequest, *LoginResponse) error
	UserList(context.Context, *UsersRequest, *UsersResponse) error
	UserDel(context.Context, *DelRequest, *DelResponse) error
}

func RegisterAdminHandler(s server.Server, hdlr AdminHandler, opts ...server.HandlerOption) error {
	type admin interface {
		AdminLogin(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		UserList(ctx context.Context, in *UsersRequest, out *UsersResponse) error
		UserDel(ctx context.Context, in *DelRequest, out *DelResponse) error
	}
	type Admin struct {
		admin
	}
	h := &adminHandler{hdlr}
	return s.Handle(s.NewHandler(&Admin{h}, opts...))
}

type adminHandler struct {
	AdminHandler
}

func (h *adminHandler) AdminLogin(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AdminHandler.AdminLogin(ctx, in, out)
}

func (h *adminHandler) UserList(ctx context.Context, in *UsersRequest, out *UsersResponse) error {
	return h.AdminHandler.UserList(ctx, in, out)
}

func (h *adminHandler) UserDel(ctx context.Context, in *DelRequest, out *DelResponse) error {
	return h.AdminHandler.UserDel(ctx, in, out)
}
