// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.17.0
// source: user/v1/filter.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserGetUserByPhone = "/user.v1.User/GetUserByPhone"
const OperationUserLogin = "/user.v1.User/Login"
const OperationUserRegister = "/user.v1.User/Register"
const OperationUserUpdateUser = "/user.v1.User/UpdateUser"

type UserHTTPServer interface {
	GetUserByPhone(context.Context, *GetUserByPhoneRequest) (*UserReply, error)
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user", _User_Register0_HTTP_Handler(srv))
	r.POST("/v1/user/login", _User_Login0_HTTP_Handler(srv))
	r.GET("/v1/user/{phone}", _User_GetUserByPhone0_HTTP_Handler(srv))
	r.PUT("/v1/user", _User_UpdateUser0_HTTP_Handler(srv))
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserByPhone0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserByPhoneRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserByPhone)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserByPhone(ctx, req.(*GetUserByPhoneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	GetUserByPhone(ctx context.Context, req *GetUserByPhoneRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetUserByPhone(ctx context.Context, in *GetUserByPhoneRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/user/{phone}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserByPhone))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
