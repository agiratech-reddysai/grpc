// Code generated by protoc-gen-go.
// source: customer.proto
// DO NOT EDIT!

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	CustomerRequest
	CustomerResponse
	CustomerFilter
*/
package customer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Request message for creating a new customer
type CustomerRequest struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone     string                     `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Addresses []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *CustomerRequest) Reset()                    { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()               {}
func (*CustomerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CustomerRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type CustomerRequest_Address struct {
	Street            string `protobuf:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `protobuf:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `protobuf:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *CustomerRequest_Address) Reset()                    { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()               {}
func (*CustomerRequest_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CustomerRequest_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CustomerRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CustomerRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CustomerRequest_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *CustomerRequest_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type CustomerResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *CustomerResponse) Reset()                    { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()               {}
func (*CustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CustomerResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CustomerFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *CustomerFilter) Reset()                    { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string            { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()               {}
func (*CustomerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerRequest)(nil), "customer.CustomerRequest")
	proto.RegisterType((*CustomerRequest_Address)(nil), "customer.CustomerRequest.Address")
	proto.RegisterType((*CustomerResponse)(nil), "customer.CustomerResponse")
	proto.RegisterType((*CustomerFilter)(nil), "customer.CustomerFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Customer service

type CustomerClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Customer_serviceDesc.Streams[0], c.cc, "/customer.Customer/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetCustomersClient interface {
	Recv() (*CustomerRequest, error)
	grpc.ClientStream
}

type customerGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerGetCustomersClient) Recv() (*CustomerRequest, error) {
	m := new(CustomerRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := grpc.Invoke(ctx, "/customer.Customer/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
	// Create a new Customer - A simple RPC
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).GetCustomers(m, &customerGetCustomersServer{stream})
}

type Customer_GetCustomersServer interface {
	Send(*CustomerRequest) error
	grpc.ServerStream
}

type customerGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerGetCustomersServer) Send(m *CustomerRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Customer_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "customer.proto",
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4a, 0xf3, 0x40,
	0x10, 0xc0, 0x9b, 0xf4, 0xff, 0x7c, 0x1f, 0xb5, 0x0e, 0x22, 0x6b, 0x9e, 0x6a, 0x9e, 0x8a, 0x48,
	0x91, 0xfa, 0x2a, 0x88, 0x08, 0x16, 0x5f, 0xe3, 0x09, 0x62, 0x32, 0xd8, 0xc5, 0x36, 0x89, 0x3b,
	0x5b, 0xa4, 0x5e, 0xc1, 0x3b, 0x78, 0x06, 0x8f, 0x68, 0x76, 0x93, 0x6d, 0xc1, 0xd8, 0xb7, 0xf9,
	0xcd, 0xce, 0xcc, 0xfe, 0x76, 0x58, 0x18, 0x25, 0x1b, 0xd6, 0xf9, 0x9a, 0xd4, 0xac, 0x50, 0xb9,
	0xce, 0x71, 0xe0, 0x38, 0xfc, 0xf6, 0xe1, 0xe8, 0xbe, 0x86, 0x88, 0xde, 0x36, 0xc4, 0x1a, 0x47,
	0xe0, 0xcb, 0x54, 0x78, 0x13, 0x6f, 0xda, 0x8d, 0xca, 0x08, 0x11, 0x3a, 0x59, 0xbc, 0x26, 0xe1,
	0x97, 0x99, 0x61, 0x64, 0x63, 0x3c, 0x81, 0x2e, 0xad, 0x63, 0xb9, 0x12, 0x6d, 0x9b, 0xac, 0xc0,
	0x64, 0x8b, 0x65, 0x9e, 0x91, 0xe8, 0x54, 0x59, 0x0b, 0x78, 0x0b, 0xc3, 0x38, 0x4d, 0x15, 0x31,
	0x13, 0x8b, 0xee, 0xa4, 0x3d, 0xfd, 0x37, 0x3f, 0x9f, 0xed, 0x8c, 0x7e, 0xdd, 0x3e, 0xbb, 0xab,
	0x4a, 0xa3, 0x7d, 0x4f, 0xf0, 0xe9, 0x41, 0xbf, 0x4e, 0xe3, 0x29, 0xf4, 0x58, 0x2b, 0x22, 0x6d,
	0x05, 0x87, 0x51, 0x4d, 0x46, 0x32, 0x91, 0x7a, 0xeb, 0x24, 0x4d, 0x6c, 0x74, 0x58, 0xc7, 0x9a,
	0x9c, 0xa4, 0x05, 0x1c, 0x43, 0xfb, 0x43, 0x16, 0xb5, 0xa2, 0x09, 0xf1, 0x12, 0x8e, 0x25, 0x3f,
	0x2d, 0x65, 0x51, 0xc8, 0xec, 0xa5, 0xbe, 0xa8, 0x14, 0xf5, 0xa6, 0x83, 0xa8, 0x79, 0x10, 0xde,
	0xc0, 0x78, 0xef, 0xcc, 0x45, 0x9e, 0x31, 0x35, 0x56, 0x26, 0xa0, 0xcf, 0x9b, 0x24, 0x31, 0x73,
	0x7c, 0x3b, 0xc7, 0x61, 0x78, 0x01, 0x23, 0xd7, 0xfd, 0x20, 0x57, 0x9a, 0x94, 0xa9, 0x7d, 0xa5,
	0xed, 0x7b, 0xae, 0xd2, 0xfa, 0x49, 0x0e, 0xe7, 0x5f, 0x1e, 0x0c, 0x5c, 0x31, 0x2e, 0xe0, 0xff,
	0x82, 0xb4, 0x43, 0x46, 0xd1, 0x5c, 0x61, 0x35, 0x30, 0x38, 0x3b, 0xb8, 0xdc, 0xb0, 0x75, 0xe5,
	0xe1, 0x63, 0x69, 0xa0, 0xa8, 0xdc, 0xc4, 0x6e, 0xf4, 0xe1, 0x86, 0x20, 0xf8, 0xeb, 0xa8, 0x7a,
	0x74, 0xd8, 0x7a, 0xee, 0xd9, 0xef, 0x74, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xde, 0x91, 0xd3,
	0x62, 0x60, 0x02, 0x00, 0x00,
}
