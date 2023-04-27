// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package stock_finance_system

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type BuyStockRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	IsVirtual            bool     `protobuf:"varint,5,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyStockRequest) Reset()         { *m = BuyStockRequest{} }
func (m *BuyStockRequest) String() string { return proto.CompactTextString(m) }
func (*BuyStockRequest) ProtoMessage()    {}
func (*BuyStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *BuyStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyStockRequest.Unmarshal(m, b)
}
func (m *BuyStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyStockRequest.Marshal(b, m, deterministic)
}
func (m *BuyStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyStockRequest.Merge(m, src)
}
func (m *BuyStockRequest) XXX_Size() int {
	return xxx_messageInfo_BuyStockRequest.Size(m)
}
func (m *BuyStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyStockRequest proto.InternalMessageInfo

func (m *BuyStockRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BuyStockRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BuyStockRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *BuyStockRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BuyStockRequest) GetIsVirtual() bool {
	if m != nil {
		return m.IsVirtual
	}
	return false
}

type BuyStockResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance              float32  `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyStockResponse) Reset()         { *m = BuyStockResponse{} }
func (m *BuyStockResponse) String() string { return proto.CompactTextString(m) }
func (*BuyStockResponse) ProtoMessage()    {}
func (*BuyStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *BuyStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyStockResponse.Unmarshal(m, b)
}
func (m *BuyStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyStockResponse.Marshal(b, m, deterministic)
}
func (m *BuyStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyStockResponse.Merge(m, src)
}
func (m *BuyStockResponse) XXX_Size() int {
	return xxx_messageInfo_BuyStockResponse.Size(m)
}
func (m *BuyStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuyStockResponse proto.InternalMessageInfo

func (m *BuyStockResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuyStockResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BuyStockResponse) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type SellStockRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	IsVirtual            bool     `protobuf:"varint,5,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellStockRequest) Reset()         { *m = SellStockRequest{} }
func (m *SellStockRequest) String() string { return proto.CompactTextString(m) }
func (*SellStockRequest) ProtoMessage()    {}
func (*SellStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *SellStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellStockRequest.Unmarshal(m, b)
}
func (m *SellStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellStockRequest.Marshal(b, m, deterministic)
}
func (m *SellStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellStockRequest.Merge(m, src)
}
func (m *SellStockRequest) XXX_Size() int {
	return xxx_messageInfo_SellStockRequest.Size(m)
}
func (m *SellStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellStockRequest proto.InternalMessageInfo

func (m *SellStockRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellStockRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SellStockRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *SellStockRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SellStockRequest) GetIsVirtual() bool {
	if m != nil {
		return m.IsVirtual
	}
	return false
}

type SellStockResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance              float32  `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellStockResponse) Reset()         { *m = SellStockResponse{} }
func (m *SellStockResponse) String() string { return proto.CompactTextString(m) }
func (*SellStockResponse) ProtoMessage()    {}
func (*SellStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *SellStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellStockResponse.Unmarshal(m, b)
}
func (m *SellStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellStockResponse.Marshal(b, m, deterministic)
}
func (m *SellStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellStockResponse.Merge(m, src)
}
func (m *SellStockResponse) XXX_Size() int {
	return xxx_messageInfo_SellStockResponse.Size(m)
}
func (m *SellStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellStockResponse proto.InternalMessageInfo

func (m *SellStockResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SellStockResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellStockResponse) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type NewAccountReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance              float32  `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAccountReq) Reset()         { *m = NewAccountReq{} }
func (m *NewAccountReq) String() string { return proto.CompactTextString(m) }
func (*NewAccountReq) ProtoMessage()    {}
func (*NewAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *NewAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountReq.Unmarshal(m, b)
}
func (m *NewAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountReq.Marshal(b, m, deterministic)
}
func (m *NewAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountReq.Merge(m, src)
}
func (m *NewAccountReq) XXX_Size() int {
	return xxx_messageInfo_NewAccountReq.Size(m)
}
func (m *NewAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountReq proto.InternalMessageInfo

func (m *NewAccountReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewAccountReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *NewAccountReq) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type NewAccountResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAccountResp) Reset()         { *m = NewAccountResp{} }
func (m *NewAccountResp) String() string { return proto.CompactTextString(m) }
func (*NewAccountResp) ProtoMessage()    {}
func (*NewAccountResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *NewAccountResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountResp.Unmarshal(m, b)
}
func (m *NewAccountResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountResp.Marshal(b, m, deterministic)
}
func (m *NewAccountResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountResp.Merge(m, src)
}
func (m *NewAccountResp) XXX_Size() int {
	return xxx_messageInfo_NewAccountResp.Size(m)
}
func (m *NewAccountResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountResp.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountResp proto.InternalMessageInfo

func (m *NewAccountResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NewAccountResp) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*BuyStockRequest)(nil), "goshare.spider.BuyStockRequest")
	proto.RegisterType((*BuyStockResponse)(nil), "goshare.spider.BuyStockResponse")
	proto.RegisterType((*SellStockRequest)(nil), "goshare.spider.SellStockRequest")
	proto.RegisterType((*SellStockResponse)(nil), "goshare.spider.SellStockResponse")
	proto.RegisterType((*NewAccountReq)(nil), "goshare.spider.NewAccountReq")
	proto.RegisterType((*NewAccountResp)(nil), "goshare.spider.NewAccountResp")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x94, 0x4d, 0x93, 0x26, 0x4f, 0x6a, 0x28, 0x2b, 0x04, 0xc6, 0x52, 0xc1, 0xe4, 0xe4, 0x0b,
	0x8e, 0x04, 0xbf, 0xa0, 0xad, 0x84, 0x40, 0x08, 0x84, 0xd6, 0x22, 0x07, 0x2e, 0xd6, 0xda, 0x7e,
	0xb8, 0xab, 0xda, 0x5e, 0xc7, 0x6f, 0x37, 0xc8, 0x3f, 0x81, 0x33, 0xfc, 0x60, 0xe4, 0xaf, 0x36,
	0x09, 0x6a, 0x0f, 0x88, 0x03, 0xb7, 0x1d, 0xef, 0xbc, 0xf1, 0xbc, 0xd1, 0x2c, 0x9c, 0x10, 0xd6,
	0x5b, 0x99, 0x60, 0x50, 0xd5, 0x4a, 0x2b, 0xb6, 0xc8, 0x14, 0x5d, 0x89, 0x1a, 0x03, 0xaa, 0x64,
	0x8a, 0xf5, 0xf2, 0xa7, 0x05, 0x0f, 0x2f, 0x4c, 0x13, 0x6a, 0x95, 0x5c, 0x73, 0xdc, 0x18, 0x24,
	0xcd, 0x9e, 0xc2, 0xb1, 0x21, 0xac, 0x23, 0x99, 0x3a, 0x96, 0x67, 0xf9, 0x13, 0x3e, 0x6d, 0xe1,
	0xfb, 0x94, 0x3d, 0x81, 0x29, 0x35, 0x45, 0xac, 0x72, 0xc7, 0xf6, 0x2c, 0x7f, 0xce, 0x07, 0xc4,
	0x5c, 0x98, 0x6d, 0x8c, 0x28, 0xb5, 0xd4, 0x8d, 0xf3, 0xa0, 0x9b, 0xb8, 0xc1, 0xec, 0x31, 0x4c,
	0xaa, 0x5a, 0x26, 0xe8, 0x1c, 0x79, 0x96, 0x6f, 0xf3, 0x1e, 0xb0, 0x33, 0x00, 0x49, 0xd1, 0x56,
	0xd6, 0xda, 0x88, 0xdc, 0x99, 0x78, 0x96, 0x3f, 0xe3, 0x73, 0x49, 0xeb, 0xfe, 0xc3, 0xf2, 0x0b,
	0x9c, 0xde, 0x9a, 0xa2, 0x4a, 0x95, 0x84, 0x6c, 0x01, 0xf6, 0x60, 0x68, 0xce, 0x6d, 0x99, 0xee,
	0xba, 0xb4, 0xf7, 0x5c, 0x3a, 0x70, 0x1c, 0x8b, 0x5c, 0x94, 0x09, 0x76, 0x66, 0x6c, 0x3e, 0xc2,
	0xe5, 0x2f, 0x0b, 0x4e, 0x43, 0xcc, 0xf3, 0xff, 0x6c, 0xdb, 0x35, 0x3c, 0xda, 0x71, 0xf5, 0xef,
	0xd6, 0x5d, 0xc3, 0xc9, 0x27, 0xfc, 0x7e, 0x9e, 0x24, 0xca, 0x94, 0x9a, 0xe3, 0x86, 0x31, 0x38,
	0x2a, 0x45, 0x81, 0x83, 0x6a, 0x77, 0xfe, 0x1b, 0xdd, 0x4b, 0x58, 0xec, 0xea, 0x52, 0xd5, 0x72,
	0x0b, 0x24, 0x12, 0xd9, 0xa8, 0x3d, 0xc2, 0x3b, 0xe5, 0x5f, 0xff, 0xb0, 0xe1, 0x59, 0xb7, 0xf1,
	0x5b, 0x59, 0xb6, 0xaa, 0x61, 0x43, 0x1a, 0x8b, 0xb0, 0x2f, 0x2b, 0xfb, 0x08, 0xb3, 0xb1, 0x00,
	0xec, 0x45, 0xb0, 0xdf, 0xd9, 0xe0, 0xa0, 0xaf, 0xae, 0x77, 0x37, 0x61, 0x08, 0xf3, 0x33, 0xcc,
	0x6f, 0x12, 0x66, 0x7f, 0xd0, 0x0f, 0x2b, 0xe1, 0xbe, 0xbc, 0x87, 0x31, 0x28, 0x7e, 0x00, 0xb8,
	0xcd, 0x80, 0x9d, 0x1d, 0x0e, 0xec, 0xe5, 0xee, 0x3e, 0xbf, 0xef, 0x9a, 0xaa, 0x8b, 0xcb, 0xaf,
	0xe7, 0x99, 0xd4, 0x57, 0x26, 0x0e, 0x12, 0x55, 0xac, 0xb8, 0xd8, 0x62, 0xf9, 0xce, 0xa8, 0x55,
	0xa6, 0x5e, 0x75, 0x8f, 0x76, 0x45, 0xed, 0xaf, 0x87, 0x73, 0x75, 0x9d, 0xf5, 0x38, 0xfa, 0xd6,
	0x07, 0x17, 0x51, 0x97, 0x5c, 0x3c, 0xed, 0xee, 0xdf, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0x6d, 0xe3, 0x54, 0xf1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockFinanceSystemServiceClient is the client API for StockFinanceSystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockFinanceSystemServiceClient interface {
	BuyStock(ctx context.Context, in *BuyStockRequest, opts ...grpc.CallOption) (*BuyStockResponse, error)
	SellStock(ctx context.Context, in *SellStockRequest, opts ...grpc.CallOption) (*SellStockResponse, error)
	NewAccount(ctx context.Context, in *NewAccountReq, opts ...grpc.CallOption) (*NewAccountResp, error)
}

type stockFinanceSystemServiceClient struct {
	cc *grpc.ClientConn
}

func NewStockFinanceSystemServiceClient(cc *grpc.ClientConn) StockFinanceSystemServiceClient {
	return &stockFinanceSystemServiceClient{cc}
}

func (c *stockFinanceSystemServiceClient) BuyStock(ctx context.Context, in *BuyStockRequest, opts ...grpc.CallOption) (*BuyStockResponse, error) {
	out := new(BuyStockResponse)
	err := c.cc.Invoke(ctx, "/goshare.spider.StockFinanceSystemService/BuyStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockFinanceSystemServiceClient) SellStock(ctx context.Context, in *SellStockRequest, opts ...grpc.CallOption) (*SellStockResponse, error) {
	out := new(SellStockResponse)
	err := c.cc.Invoke(ctx, "/goshare.spider.StockFinanceSystemService/SellStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockFinanceSystemServiceClient) NewAccount(ctx context.Context, in *NewAccountReq, opts ...grpc.CallOption) (*NewAccountResp, error) {
	out := new(NewAccountResp)
	err := c.cc.Invoke(ctx, "/goshare.spider.StockFinanceSystemService/NewAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockFinanceSystemServiceServer is the server API for StockFinanceSystemService service.
type StockFinanceSystemServiceServer interface {
	BuyStock(context.Context, *BuyStockRequest) (*BuyStockResponse, error)
	SellStock(context.Context, *SellStockRequest) (*SellStockResponse, error)
	NewAccount(context.Context, *NewAccountReq) (*NewAccountResp, error)
}

func RegisterStockFinanceSystemServiceServer(s *grpc.Server, srv StockFinanceSystemServiceServer) {
	s.RegisterService(&_StockFinanceSystemService_serviceDesc, srv)
}

func _StockFinanceSystemService_BuyStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockFinanceSystemServiceServer).BuyStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goshare.spider.StockFinanceSystemService/BuyStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockFinanceSystemServiceServer).BuyStock(ctx, req.(*BuyStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockFinanceSystemService_SellStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockFinanceSystemServiceServer).SellStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goshare.spider.StockFinanceSystemService/SellStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockFinanceSystemServiceServer).SellStock(ctx, req.(*SellStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockFinanceSystemService_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockFinanceSystemServiceServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goshare.spider.StockFinanceSystemService/NewAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockFinanceSystemServiceServer).NewAccount(ctx, req.(*NewAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockFinanceSystemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goshare.spider.StockFinanceSystemService",
	HandlerType: (*StockFinanceSystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyStock",
			Handler:    _StockFinanceSystemService_BuyStock_Handler,
		},
		{
			MethodName: "SellStock",
			Handler:    _StockFinanceSystemService_SellStock_Handler,
		},
		{
			MethodName: "NewAccount",
			Handler:    _StockFinanceSystemService_NewAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
