// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: dhctl.proto

package dhctl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DHCTL_Check_FullMethodName     = "/dhctl.DHCTL/Check"
	DHCTL_Bootstrap_FullMethodName = "/dhctl.DHCTL/Bootstrap"
)

// DHCTLClient is the client API for DHCTL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DHCTLClient interface {
	Check(ctx context.Context, opts ...grpc.CallOption) (DHCTL_CheckClient, error)
	Bootstrap(ctx context.Context, opts ...grpc.CallOption) (DHCTL_BootstrapClient, error)
}

type dHCTLClient struct {
	cc grpc.ClientConnInterface
}

func NewDHCTLClient(cc grpc.ClientConnInterface) DHCTLClient {
	return &dHCTLClient{cc}
}

func (c *dHCTLClient) Check(ctx context.Context, opts ...grpc.CallOption) (DHCTL_CheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &DHCTL_ServiceDesc.Streams[0], DHCTL_Check_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dHCTLCheckClient{stream}
	return x, nil
}

type DHCTL_CheckClient interface {
	Send(*CheckRequest) error
	Recv() (*CheckResponse, error)
	grpc.ClientStream
}

type dHCTLCheckClient struct {
	grpc.ClientStream
}

func (x *dHCTLCheckClient) Send(m *CheckRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dHCTLCheckClient) Recv() (*CheckResponse, error) {
	m := new(CheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dHCTLClient) Bootstrap(ctx context.Context, opts ...grpc.CallOption) (DHCTL_BootstrapClient, error) {
	stream, err := c.cc.NewStream(ctx, &DHCTL_ServiceDesc.Streams[1], DHCTL_Bootstrap_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dHCTLBootstrapClient{stream}
	return x, nil
}

type DHCTL_BootstrapClient interface {
	Send(*BootstrapRequest) error
	Recv() (*BootstrapResponse, error)
	grpc.ClientStream
}

type dHCTLBootstrapClient struct {
	grpc.ClientStream
}

func (x *dHCTLBootstrapClient) Send(m *BootstrapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dHCTLBootstrapClient) Recv() (*BootstrapResponse, error) {
	m := new(BootstrapResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DHCTLServer is the server API for DHCTL service.
// All implementations must embed UnimplementedDHCTLServer
// for forward compatibility
type DHCTLServer interface {
	Check(DHCTL_CheckServer) error
	Bootstrap(DHCTL_BootstrapServer) error
	mustEmbedUnimplementedDHCTLServer()
}

// UnimplementedDHCTLServer must be embedded to have forward compatible implementations.
type UnimplementedDHCTLServer struct {
}

func (UnimplementedDHCTLServer) Check(DHCTL_CheckServer) error {
	return status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedDHCTLServer) Bootstrap(DHCTL_BootstrapServer) error {
	return status.Errorf(codes.Unimplemented, "method Bootstrap not implemented")
}
func (UnimplementedDHCTLServer) mustEmbedUnimplementedDHCTLServer() {}

// UnsafeDHCTLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DHCTLServer will
// result in compilation errors.
type UnsafeDHCTLServer interface {
	mustEmbedUnimplementedDHCTLServer()
}

func RegisterDHCTLServer(s grpc.ServiceRegistrar, srv DHCTLServer) {
	s.RegisterService(&DHCTL_ServiceDesc, srv)
}

func _DHCTL_Check_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DHCTLServer).Check(&dHCTLCheckServer{stream})
}

type DHCTL_CheckServer interface {
	Send(*CheckResponse) error
	Recv() (*CheckRequest, error)
	grpc.ServerStream
}

type dHCTLCheckServer struct {
	grpc.ServerStream
}

func (x *dHCTLCheckServer) Send(m *CheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dHCTLCheckServer) Recv() (*CheckRequest, error) {
	m := new(CheckRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DHCTL_Bootstrap_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DHCTLServer).Bootstrap(&dHCTLBootstrapServer{stream})
}

type DHCTL_BootstrapServer interface {
	Send(*BootstrapResponse) error
	Recv() (*BootstrapRequest, error)
	grpc.ServerStream
}

type dHCTLBootstrapServer struct {
	grpc.ServerStream
}

func (x *dHCTLBootstrapServer) Send(m *BootstrapResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dHCTLBootstrapServer) Recv() (*BootstrapRequest, error) {
	m := new(BootstrapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DHCTL_ServiceDesc is the grpc.ServiceDesc for DHCTL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DHCTL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dhctl.DHCTL",
	HandlerType: (*DHCTLServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Check",
			Handler:       _DHCTL_Check_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Bootstrap",
			Handler:       _DHCTL_Bootstrap_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dhctl.proto",
}