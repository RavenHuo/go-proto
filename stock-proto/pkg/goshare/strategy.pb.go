// Code generated by protoc-gen-go. DO NOT EDIT.
// source: strategy.proto

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

// Signal 交易信号
type Signal struct {
	OpenTime               int64    `protobuf:"varint,1,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	OpenPrice              float64  `protobuf:"fixed64,2,opt,name=open_price,json=openPrice,proto3" json:"open_price,omitempty"`
	CloseTime              int64    `protobuf:"varint,3,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	ClosePrice             float64  `protobuf:"fixed64,4,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"`
	Volume                 int32    `protobuf:"varint,5,opt,name=volume,proto3" json:"volume,omitempty"`
	Direction              int32    `protobuf:"varint,6,opt,name=direction,proto3" json:"direction,omitempty"`
	Profit                 float64  `protobuf:"fixed64,7,opt,name=profit,proto3" json:"profit,omitempty"`
	AccumulativeProfit     float64  `protobuf:"fixed64,8,opt,name=accumulative_profit,json=accumulativeProfit,proto3" json:"accumulative_profit,omitempty"`
	Commission             float64  `protobuf:"fixed64,9,opt,name=commission,proto3" json:"commission,omitempty"`
	AccumulativeCommission float64  `protobuf:"fixed64,10,opt,name=accumulative_commission,json=accumulativeCommission,proto3" json:"accumulative_commission,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Signal) Reset()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) ProtoMessage()    {}
func (*Signal) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec5ce6dd46feab, []int{0}
}

func (m *Signal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signal.Unmarshal(m, b)
}
func (m *Signal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signal.Marshal(b, m, deterministic)
}
func (m *Signal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signal.Merge(m, src)
}
func (m *Signal) XXX_Size() int {
	return xxx_messageInfo_Signal.Size(m)
}
func (m *Signal) XXX_DiscardUnknown() {
	xxx_messageInfo_Signal.DiscardUnknown(m)
}

var xxx_messageInfo_Signal proto.InternalMessageInfo

func (m *Signal) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *Signal) GetOpenPrice() float64 {
	if m != nil {
		return m.OpenPrice
	}
	return 0
}

func (m *Signal) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *Signal) GetClosePrice() float64 {
	if m != nil {
		return m.ClosePrice
	}
	return 0
}

func (m *Signal) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Signal) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *Signal) GetProfit() float64 {
	if m != nil {
		return m.Profit
	}
	return 0
}

func (m *Signal) GetAccumulativeProfit() float64 {
	if m != nil {
		return m.AccumulativeProfit
	}
	return 0
}

func (m *Signal) GetCommission() float64 {
	if m != nil {
		return m.Commission
	}
	return 0
}

func (m *Signal) GetAccumulativeCommission() float64 {
	if m != nil {
		return m.AccumulativeCommission
	}
	return 0
}

// strategy signal
type AnalyzeResult struct {
	Profit               float64  `protobuf:"fixed64,1,opt,name=profit,proto3" json:"profit,omitempty"`
	MaxDrawdown          float64  `protobuf:"fixed64,2,opt,name=max_drawdown,json=maxDrawdown,proto3" json:"max_drawdown,omitempty"`
	TakeTimes            int32    `protobuf:"varint,3,opt,name=take_times,json=takeTimes,proto3" json:"take_times,omitempty"`
	LossTimes            int32    `protobuf:"varint,4,opt,name=loss_times,json=lossTimes,proto3" json:"loss_times,omitempty"`
	TradeTimes           int32    `protobuf:"varint,5,opt,name=trade_times,json=tradeTimes,proto3" json:"trade_times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalyzeResult) Reset()         { *m = AnalyzeResult{} }
func (m *AnalyzeResult) String() string { return proto.CompactTextString(m) }
func (*AnalyzeResult) ProtoMessage()    {}
func (*AnalyzeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec5ce6dd46feab, []int{1}
}

func (m *AnalyzeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeResult.Unmarshal(m, b)
}
func (m *AnalyzeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeResult.Marshal(b, m, deterministic)
}
func (m *AnalyzeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeResult.Merge(m, src)
}
func (m *AnalyzeResult) XXX_Size() int {
	return xxx_messageInfo_AnalyzeResult.Size(m)
}
func (m *AnalyzeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeResult.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeResult proto.InternalMessageInfo

func (m *AnalyzeResult) GetProfit() float64 {
	if m != nil {
		return m.Profit
	}
	return 0
}

func (m *AnalyzeResult) GetMaxDrawdown() float64 {
	if m != nil {
		return m.MaxDrawdown
	}
	return 0
}

func (m *AnalyzeResult) GetTakeTimes() int32 {
	if m != nil {
		return m.TakeTimes
	}
	return 0
}

func (m *AnalyzeResult) GetLossTimes() int32 {
	if m != nil {
		return m.LossTimes
	}
	return 0
}

func (m *AnalyzeResult) GetTradeTimes() int32 {
	if m != nil {
		return m.TradeTimes
	}
	return 0
}

// SignalList list
type SignalList struct {
	Exchange             string         `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol               string         `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Multiple             float64        `protobuf:"fixed64,3,opt,name=multiple,proto3" json:"multiple,omitempty"`
	PriceTick            float64        `protobuf:"fixed64,4,opt,name=price_tick,json=priceTick,proto3" json:"price_tick,omitempty"`
	List                 []*Signal      `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	Result               *AnalyzeResult `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignalList) Reset()         { *m = SignalList{} }
func (m *SignalList) String() string { return proto.CompactTextString(m) }
func (*SignalList) ProtoMessage()    {}
func (*SignalList) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec5ce6dd46feab, []int{2}
}

func (m *SignalList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalList.Unmarshal(m, b)
}
func (m *SignalList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalList.Marshal(b, m, deterministic)
}
func (m *SignalList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalList.Merge(m, src)
}
func (m *SignalList) XXX_Size() int {
	return xxx_messageInfo_SignalList.Size(m)
}
func (m *SignalList) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalList.DiscardUnknown(m)
}

var xxx_messageInfo_SignalList proto.InternalMessageInfo

func (m *SignalList) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *SignalList) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SignalList) GetMultiple() float64 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *SignalList) GetPriceTick() float64 {
	if m != nil {
		return m.PriceTick
	}
	return 0
}

func (m *SignalList) GetList() []*Signal {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SignalList) GetResult() *AnalyzeResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Signal)(nil), "goshare.Signal")
	proto.RegisterType((*AnalyzeResult)(nil), "goshare.AnalyzeResult")
	proto.RegisterType((*SignalList)(nil), "goshare.SignalList")
}

func init() { proto.RegisterFile("strategy.proto", fileDescriptor_46ec5ce6dd46feab) }

var fileDescriptor_46ec5ce6dd46feab = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x55, 0xa6, 0x8f, 0x69, 0x6e, 0x78, 0x48, 0x46, 0x2a, 0x11, 0xaf, 0x29, 0x65, 0xd3, 0x0d,
	0x89, 0x34, 0x20, 0x58, 0xf3, 0x58, 0xb0, 0x60, 0x31, 0x32, 0xb3, 0x62, 0x53, 0xb9, 0xae, 0x49,
	0xad, 0xd8, 0x71, 0x14, 0x3b, 0x9d, 0x96, 0x3f, 0xe2, 0x6f, 0xf8, 0x00, 0x3e, 0x06, 0xf9, 0xda,
	0x84, 0x76, 0x97, 0x7b, 0x1e, 0xf6, 0xcd, 0xc9, 0x09, 0x3c, 0xb0, 0xae, 0x63, 0x4e, 0x54, 0xc7,
	0xa2, 0xed, 0x8c, 0x33, 0xe4, 0xb2, 0x32, 0x76, 0xc7, 0x3a, 0xb1, 0xfc, 0x73, 0x01, 0xd3, 0x6f,
	0xb2, 0x6a, 0x98, 0x22, 0x4f, 0x21, 0x35, 0xad, 0x68, 0xd6, 0x4e, 0x6a, 0x91, 0x27, 0x8b, 0x64,
	0x35, 0xa2, 0x33, 0x0f, 0xdc, 0x4a, 0x2d, 0xc8, 0x73, 0x00, 0x24, 0xdb, 0x4e, 0x72, 0x91, 0x5f,
	0x2c, 0x92, 0x55, 0x42, 0x51, 0x7e, 0xe3, 0x01, 0x4f, 0x73, 0x65, 0xac, 0x08, 0xe6, 0x11, 0x9a,
	0x53, 0x44, 0xd0, 0x7d, 0x05, 0x59, 0xa0, 0x83, 0x7d, 0x8c, 0xf6, 0xe0, 0x08, 0xfe, 0x39, 0x4c,
	0xf7, 0x46, 0xf5, 0x5a, 0xe4, 0x93, 0x45, 0xb2, 0x9a, 0xd0, 0x38, 0x91, 0x67, 0x90, 0x6e, 0x65,
	0x27, 0xb8, 0x93, 0xa6, 0xc9, 0xa7, 0x48, 0xfd, 0x07, 0xbc, 0xab, 0xed, 0xcc, 0x0f, 0xe9, 0xf2,
	0x4b, 0x3c, 0x31, 0x4e, 0xa4, 0x84, 0x47, 0x8c, 0xf3, 0x5e, 0xf7, 0x8a, 0x39, 0xb9, 0xf7, 0xb7,
	0xa2, 0x68, 0x86, 0x22, 0x72, 0x4a, 0xdd, 0x04, 0xc3, 0x0b, 0x00, 0x6e, 0xb4, 0x96, 0xd6, 0xfa,
	0x7b, 0xd2, 0xb8, 0xde, 0x80, 0x90, 0xf7, 0xf0, 0xf8, 0xec, 0xc0, 0x13, 0x31, 0xa0, 0x78, 0x7e,
	0x4a, 0x7f, 0x1a, 0xd8, 0xe5, 0xaf, 0x04, 0xee, 0x7f, 0x68, 0x98, 0x3a, 0xfe, 0x14, 0x54, 0xd8,
	0x5e, 0xb9, 0x93, 0x9d, 0x93, 0xb3, 0x9d, 0x5f, 0xc2, 0x3d, 0xcd, 0x0e, 0xeb, 0x6d, 0xc7, 0xee,
	0xb6, 0xe6, 0xae, 0x89, 0x11, 0x67, 0x9a, 0x1d, 0x3e, 0x47, 0xc8, 0x87, 0xec, 0x58, 0x1d, 0x32,
	0xb6, 0x18, 0xf2, 0x84, 0xa6, 0x1e, 0xf1, 0x19, 0x5b, 0x4f, 0x2b, 0x63, 0x6d, 0xa4, 0xc7, 0x81,
	0xf6, 0x48, 0xa0, 0xaf, 0x20, 0x73, 0x1d, 0xdb, 0xfe, 0xb3, 0x87, 0x9c, 0x01, 0x21, 0x14, 0x2c,
	0x7f, 0x27, 0x00, 0xa1, 0x0a, 0x5f, 0xa5, 0x75, 0xe4, 0x09, 0xcc, 0xc4, 0x81, 0xef, 0x58, 0x53,
	0x85, 0x36, 0xa4, 0x74, 0x98, 0xfd, 0x4b, 0xd8, 0xa3, 0xde, 0x18, 0x85, 0x6b, 0xa6, 0x34, 0x4e,
	0xde, 0xa3, 0x7b, 0xe5, 0x64, 0xab, 0x42, 0x09, 0x12, 0x3a, 0xcc, 0x7e, 0x3d, 0xfc, 0xfa, 0x6b,
	0x27, 0x79, 0x1d, 0x2b, 0x90, 0x22, 0x72, 0x2b, 0x79, 0x4d, 0x5e, 0xc1, 0x58, 0x49, 0xeb, 0xf2,
	0xc9, 0x62, 0xb4, 0xca, 0xae, 0x1f, 0x16, 0xb1, 0xa0, 0x45, 0xd8, 0x88, 0x22, 0x49, 0x0a, 0x98,
	0x76, 0x18, 0x23, 0x76, 0x21, 0xbb, 0x9e, 0x0f, 0xb2, 0xb3, 0x90, 0x69, 0x54, 0x7d, 0x7c, 0xf7,
	0xfd, 0x6d, 0x25, 0xdd, 0xae, 0xdf, 0x14, 0xdc, 0xe8, 0x92, 0xb2, 0xbd, 0x68, 0xbe, 0xf4, 0xa6,
	0xac, 0xcc, 0x6b, 0xfc, 0x0d, 0x4a, 0xeb, 0x0c, 0xaf, 0xe3, 0x73, 0x5b, 0x57, 0x65, 0x3c, 0x6d,
	0x33, 0x45, 0xe8, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xa7, 0x29, 0xae, 0x37, 0x03,
	0x00, 0x00,
}
