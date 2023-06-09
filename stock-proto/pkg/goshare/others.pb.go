// Code generated by protoc-gen-go. DO NOT EDIT.
// source: others.proto

package goshare

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

type NetInAmountDetail struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Percentage           float64  `protobuf:"fixed64,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetInAmountDetail) Reset()         { *m = NetInAmountDetail{} }
func (m *NetInAmountDetail) String() string { return proto.CompactTextString(m) }
func (*NetInAmountDetail) ProtoMessage()    {}
func (*NetInAmountDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_9597193cc8d6d418, []int{0}
}

func (m *NetInAmountDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetInAmountDetail.Unmarshal(m, b)
}
func (m *NetInAmountDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetInAmountDetail.Marshal(b, m, deterministic)
}
func (m *NetInAmountDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetInAmountDetail.Merge(m, src)
}
func (m *NetInAmountDetail) XXX_Size() int {
	return xxx_messageInfo_NetInAmountDetail.Size(m)
}
func (m *NetInAmountDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_NetInAmountDetail.DiscardUnknown(m)
}

var xxx_messageInfo_NetInAmountDetail proto.InternalMessageInfo

func (m *NetInAmountDetail) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *NetInAmountDetail) GetPercentage() float64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type RealtimeMoneyTrendItem struct {
	Symbol               string             `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64            `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	UpdownPercentage     float64            `protobuf:"fixed64,4,opt,name=updown_percentage,json=updownPercentage,proto3" json:"updown_percentage,omitempty"`
	Time                 int64              `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"`
	SuperSuperBigOrder   *NetInAmountDetail `protobuf:"bytes,7,opt,name=super_super_big_order,json=superSuperBigOrder,proto3" json:"super_super_big_order,omitempty"`
	SuperBigOrder        *NetInAmountDetail `protobuf:"bytes,8,opt,name=super_big_order,json=superBigOrder,proto3" json:"super_big_order,omitempty"`
	BigOrder             *NetInAmountDetail `protobuf:"bytes,9,opt,name=big_order,json=bigOrder,proto3" json:"big_order,omitempty"`
	MiddleOrder          *NetInAmountDetail `protobuf:"bytes,10,opt,name=middle_order,json=middleOrder,proto3" json:"middle_order,omitempty"`
	SmallOrder           *NetInAmountDetail `protobuf:"bytes,11,opt,name=small_order,json=smallOrder,proto3" json:"small_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RealtimeMoneyTrendItem) Reset()         { *m = RealtimeMoneyTrendItem{} }
func (m *RealtimeMoneyTrendItem) String() string { return proto.CompactTextString(m) }
func (*RealtimeMoneyTrendItem) ProtoMessage()    {}
func (*RealtimeMoneyTrendItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9597193cc8d6d418, []int{1}
}

func (m *RealtimeMoneyTrendItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Unmarshal(m, b)
}
func (m *RealtimeMoneyTrendItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Marshal(b, m, deterministic)
}
func (m *RealtimeMoneyTrendItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealtimeMoneyTrendItem.Merge(m, src)
}
func (m *RealtimeMoneyTrendItem) XXX_Size() int {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Size(m)
}
func (m *RealtimeMoneyTrendItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RealtimeMoneyTrendItem.DiscardUnknown(m)
}

var xxx_messageInfo_RealtimeMoneyTrendItem proto.InternalMessageInfo

func (m *RealtimeMoneyTrendItem) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *RealtimeMoneyTrendItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RealtimeMoneyTrendItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetUpdownPercentage() float64 {
	if m != nil {
		return m.UpdownPercentage
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetSuperSuperBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.SuperSuperBigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetSuperBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.SuperBigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.BigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetMiddleOrder() *NetInAmountDetail {
	if m != nil {
		return m.MiddleOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetSmallOrder() *NetInAmountDetail {
	if m != nil {
		return m.SmallOrder
	}
	return nil
}

type RealtimeMoneyTrendItemList struct {
	List                 []*RealtimeMoneyTrendItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RealtimeMoneyTrendItemList) Reset()         { *m = RealtimeMoneyTrendItemList{} }
func (m *RealtimeMoneyTrendItemList) String() string { return proto.CompactTextString(m) }
func (*RealtimeMoneyTrendItemList) ProtoMessage()    {}
func (*RealtimeMoneyTrendItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9597193cc8d6d418, []int{2}
}

func (m *RealtimeMoneyTrendItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Unmarshal(m, b)
}
func (m *RealtimeMoneyTrendItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Marshal(b, m, deterministic)
}
func (m *RealtimeMoneyTrendItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealtimeMoneyTrendItemList.Merge(m, src)
}
func (m *RealtimeMoneyTrendItemList) XXX_Size() int {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Size(m)
}
func (m *RealtimeMoneyTrendItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_RealtimeMoneyTrendItemList.DiscardUnknown(m)
}

var xxx_messageInfo_RealtimeMoneyTrendItemList proto.InternalMessageInfo

func (m *RealtimeMoneyTrendItemList) GetList() []*RealtimeMoneyTrendItem {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*NetInAmountDetail)(nil), "goshare.NetInAmountDetail")
	proto.RegisterType((*RealtimeMoneyTrendItem)(nil), "goshare.RealtimeMoneyTrendItem")
	proto.RegisterType((*RealtimeMoneyTrendItemList)(nil), "goshare.RealtimeMoneyTrendItemList")
}

func init() { proto.RegisterFile("others.proto", fileDescriptor_9597193cc8d6d418) }

var fileDescriptor_9597193cc8d6d418 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0xa9, 0x9d, 0xbb, 0xb7, 0xa7, 0x57, 0x74, 0x41, 0x47, 0xd9, 0x83, 0x8e, 0x3d, 0x15,
	0xc4, 0x16, 0x36, 0xd1, 0x07, 0xf1, 0xc1, 0xe1, 0x83, 0x43, 0xe7, 0x8f, 0xe8, 0x93, 0x2f, 0x23,
	0x6d, 0x0f, 0x5d, 0x58, 0x93, 0x94, 0x24, 0x55, 0xf6, 0xf7, 0xfa, 0x8f, 0x48, 0xd3, 0x3a, 0xa7,
	0x77, 0xd0, 0x97, 0x70, 0xce, 0xc9, 0xf7, 0xf3, 0x69, 0x48, 0x03, 0x37, 0xca, 0xee, 0x51, 0x9b,
	0xa4, 0xd6, 0xca, 0x2a, 0x72, 0x55, 0x2a, 0xb3, 0x67, 0x1a, 0x17, 0xef, 0x61, 0xf2, 0x11, 0xed,
	0x46, 0xbe, 0x11, 0xaa, 0x91, 0xf6, 0x2d, 0x5a, 0xc6, 0x2b, 0x32, 0x85, 0x31, 0x73, 0x7d, 0xe4,
	0xcd, 0xbd, 0xd8, 0xa3, 0x7d, 0x47, 0x1e, 0x03, 0xd4, 0xa8, 0x73, 0x94, 0x96, 0x95, 0x18, 0xdd,
	0x71, 0x7b, 0x67, 0x93, 0xc5, 0x2f, 0x1f, 0xa6, 0x14, 0x59, 0x65, 0xb9, 0xc0, 0xad, 0x92, 0x78,
	0xfc, 0xa6, 0x51, 0x16, 0x1b, 0x8b, 0xa2, 0x55, 0x9a, 0xa3, 0xc8, 0x54, 0xe5, 0x94, 0x01, 0xed,
	0x3b, 0x42, 0x60, 0x24, 0x99, 0xe8, 0x64, 0x01, 0x75, 0x35, 0x79, 0x08, 0x77, 0x6b, 0xcd, 0x73,
	0x8c, 0x7c, 0xf7, 0x85, 0xae, 0x21, 0x4f, 0x61, 0xd2, 0xd4, 0x85, 0xfa, 0x29, 0x77, 0x67, 0x67,
	0x18, 0xb9, 0xc4, 0x83, 0x6e, 0xe3, 0xf3, 0x69, 0xde, 0x6a, 0xdb, 0x43, 0x44, 0xe3, 0xb9, 0x17,
	0xfb, 0xd4, 0xd5, 0x64, 0x0b, 0x8f, 0x4c, 0x53, 0xa3, 0xde, 0x75, 0x6b, 0xc6, 0xcb, 0x9d, 0xd2,
	0x05, 0xea, 0xe8, 0x6a, 0xee, 0xc5, 0xe1, 0x72, 0x96, 0xf4, 0x77, 0x92, 0xdc, 0xba, 0x10, 0x4a,
	0x1c, 0xf2, 0xb5, 0x5d, 0xd6, 0xbc, 0xfc, 0xd4, 0x52, 0x64, 0x0d, 0xf7, 0xff, 0x17, 0x5d, 0x0f,
	0x8a, 0xee, 0x99, 0x7f, 0x1c, 0x2f, 0x21, 0xf8, 0x4b, 0x07, 0x83, 0xf4, 0x75, 0xf6, 0x07, 0x7c,
	0x0d, 0x37, 0x82, 0x17, 0x45, 0x85, 0x3d, 0x0b, 0x83, 0x6c, 0xd8, 0xe5, 0x3b, 0xfc, 0x15, 0x84,
	0x46, 0xb0, 0xaa, 0xea, 0xe9, 0x70, 0x90, 0x06, 0x17, 0x77, 0xf0, 0xe2, 0x0b, 0xcc, 0x2e, 0xff,
	0xe4, 0x0f, 0xdc, 0x58, 0xb2, 0x82, 0x51, 0xc5, 0x4d, 0xfb, 0x72, 0xfc, 0x38, 0x5c, 0x3e, 0x39,
	0x39, 0x2f, 0x23, 0xd4, 0x85, 0xd7, 0x2f, 0xbe, 0x3f, 0x2f, 0xb9, 0xdd, 0x37, 0x59, 0x92, 0x2b,
	0x91, 0x52, 0xf6, 0x03, 0xe5, 0xbb, 0x46, 0xa5, 0xa5, 0x7a, 0xe6, 0x9e, 0x6b, 0x6a, 0xac, 0xca,
	0x0f, 0x7d, 0x5d, 0x1f, 0xca, 0xb4, 0x97, 0x66, 0x63, 0x37, 0x5a, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x77, 0x93, 0x40, 0x2a, 0xdd, 0x02, 0x00, 0x00,
}
