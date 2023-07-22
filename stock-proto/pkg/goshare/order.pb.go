// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

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

// 投机套保标志
type HedgeType int32

const (
	// 投机
	HedgeType_SPECULATION HedgeType = 0
	// 套利
	HedgeType_ARBITRAGE HedgeType = 1
	// 套保
	HedgeType_HEDGE HedgeType = 2
	// 投机2(特别标志)
	HedgeType_SPECULATION2 HedgeType = 3
)

var HedgeType_name = map[int32]string{
	0: "SPECULATION",
	1: "ARBITRAGE",
	2: "HEDGE",
	3: "SPECULATION2",
}

var HedgeType_value = map[string]int32{
	"SPECULATION":  0,
	"ARBITRAGE":    1,
	"HEDGE":        2,
	"SPECULATION2": 3,
}

func (x HedgeType) String() string {
	return proto.EnumName(HedgeType_name, int32(x))
}

func (HedgeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

type OrderStatus int32

const (
	// 完全成交
	OrderStatus_DONE OrderStatus = 0
	// 部分成交还在队列中
	OrderStatus_PENDING_WITH_PARTIAL_DONE OrderStatus = 1
	// 部分成交且已撤单
	OrderStatus_CANCELED_WITH_PARTIAL_DONE OrderStatus = 2
	// 队列中
	OrderStatus_PENDING OrderStatus = 3
	// 已经撤单
	OrderStatus_CANCELED OrderStatus = 4
	// 未知
	OrderStatus_UNKOWN OrderStatus = 5
)

var OrderStatus_name = map[int32]string{
	0: "DONE",
	1: "PENDING_WITH_PARTIAL_DONE",
	2: "CANCELED_WITH_PARTIAL_DONE",
	3: "PENDING",
	4: "CANCELED",
	5: "UNKOWN",
}

var OrderStatus_value = map[string]int32{
	"DONE":                       0,
	"PENDING_WITH_PARTIAL_DONE":  1,
	"CANCELED_WITH_PARTIAL_DONE": 2,
	"PENDING":                    3,
	"CANCELED":                   4,
	"UNKOWN":                     5,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

type OrderPriceType int32

const (
	// 限价
	OrderPriceType_LIMIT_PRICE OrderPriceType = 0
	// 市价
	OrderPriceType_MARKET_PRICE OrderPriceType = 1
	// 最优价
	OrderPriceType_BEST_PRICE OrderPriceType = 2
	// 最新价
	OrderPriceType_LAST_PRICE OrderPriceType = 3
)

var OrderPriceType_name = map[int32]string{
	0: "LIMIT_PRICE",
	1: "MARKET_PRICE",
	2: "BEST_PRICE",
	3: "LAST_PRICE",
}

var OrderPriceType_value = map[string]int32{
	"LIMIT_PRICE":  0,
	"MARKET_PRICE": 1,
	"BEST_PRICE":   2,
	"LAST_PRICE":   3,
}

func (x OrderPriceType) String() string {
	return proto.EnumName(OrderPriceType_name, int32(x))
}

func (OrderPriceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

type OffsetFlag int32

const (
	// 开仓
	OffsetFlag_OPEN OffsetFlag = 0
	// 平仓
	OffsetFlag_CLOSE OffsetFlag = 1
	// 强平
	OffsetFlag_FORCE_CLOSE OffsetFlag = 2
	// 平今
	OffsetFlag_CLOSE_TODAY OffsetFlag = 3
	// 平昨
	OffsetFlag_CLOSE_YESTERDAY OffsetFlag = 4
)

var OffsetFlag_name = map[int32]string{
	0: "OPEN",
	1: "CLOSE",
	2: "FORCE_CLOSE",
	3: "CLOSE_TODAY",
	4: "CLOSE_YESTERDAY",
}

var OffsetFlag_value = map[string]int32{
	"OPEN":            0,
	"CLOSE":           1,
	"FORCE_CLOSE":     2,
	"CLOSE_TODAY":     3,
	"CLOSE_YESTERDAY": 4,
}

func (x OffsetFlag) String() string {
	return proto.EnumName(OffsetFlag_name, int32(x))
}

func (OffsetFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

// 报单方向
type DirectionType int32

const (
	DirectionType_LONG  DirectionType = 0
	DirectionType_SHORT DirectionType = 1
	DirectionType_NET   DirectionType = 2
)

var DirectionType_name = map[int32]string{
	0: "LONG",
	1: "SHORT",
	2: "NET",
}

var DirectionType_value = map[string]int32{
	"LONG":  0,
	"SHORT": 1,
	"NET":   2,
}

func (x DirectionType) String() string {
	return proto.EnumName(DirectionType_name, int32(x))
}

func (DirectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

// 组合方向
type CombDirectionType int32

const (
	CombDirectionType_COMB   CombDirectionType = 0
	CombDirectionType_UNCOMB CombDirectionType = 1
)

var CombDirectionType_name = map[int32]string{
	0: "COMB",
	1: "UNCOMB",
}

var CombDirectionType_value = map[string]int32{
	"COMB":   0,
	"UNCOMB": 1,
}

func (x CombDirectionType) String() string {
	return proto.EnumName(CombDirectionType_name, int32(x))
}

func (CombDirectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

// 强平类型
type ForceCloseReason int32

const (
	// 资金不足
	ForceCloseReason_NOT_ENOUGH_MARGIN ForceCloseReason = 0
	// 客户超仓
	ForceCloseReason_client_over_position_limit ForceCloseReason = 1
	// 会员超仓
	ForceCloseReason_broker_over_poisition_limit ForceCloseReason = 2
	// 违规
	ForceCloseReason_violation ForceCloseReason = 3
	// 其它
	ForceCloseReason_other ForceCloseReason = 4
	// 自然人临近交割
	ForceCloseReason_person_deliver ForceCloseReason = 5
)

var ForceCloseReason_name = map[int32]string{
	0: "NOT_ENOUGH_MARGIN",
	1: "client_over_position_limit",
	2: "broker_over_poisition_limit",
	3: "violation",
	4: "other",
	5: "person_deliver",
}

var ForceCloseReason_value = map[string]int32{
	"NOT_ENOUGH_MARGIN":           0,
	"client_over_position_limit":  1,
	"broker_over_poisition_limit": 2,
	"violation":                   3,
	"other":                       4,
	"person_deliver":              5,
}

func (x ForceCloseReason) String() string {
	return proto.EnumName(ForceCloseReason_name, int32(x))
}

func (ForceCloseReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{6}
}

// 报单来源
type OrderSourceType int32

const (
	// 客户
	OrderSourceType_CLIENT OrderSourceType = 0
	// 管理员
	OrderSourceType_ADMINISTRATOR OrderSourceType = 1
	// 风控
	OrderSourceType_RISK_SRV OrderSourceType = 2
	// 手机端
	OrderSourceType_MOBILE OrderSourceType = 3
	// PC端
	OrderSourceType_PC OrderSourceType = 4
	// 结算衍生
	OrderSourceType_SETTLEMENT_DERIVED OrderSourceType = 5
	// 套利单衍生
	OrderSourceType_COMBINATION_DERIVED OrderSourceType = 6
	// 多路径平仓衍生
	OrderSourceType_MULTI_ROUTE_CLOSE_DERIVED OrderSourceType = 7
	// 分红送股
	OrderSourceType_PLACEMENT_SHARE OrderSourceType = 8
	// 回购操作
	OrderSourceType_REPURCHASE OrderSourceType = 9
	// CTP上传
	OrderSourceType_CTP_UPLOAD OrderSourceType = 10
)

var OrderSourceType_name = map[int32]string{
	0:  "CLIENT",
	1:  "ADMINISTRATOR",
	2:  "RISK_SRV",
	3:  "MOBILE",
	4:  "PC",
	5:  "SETTLEMENT_DERIVED",
	6:  "COMBINATION_DERIVED",
	7:  "MULTI_ROUTE_CLOSE_DERIVED",
	8:  "PLACEMENT_SHARE",
	9:  "REPURCHASE",
	10: "CTP_UPLOAD",
}

var OrderSourceType_value = map[string]int32{
	"CLIENT":                    0,
	"ADMINISTRATOR":             1,
	"RISK_SRV":                  2,
	"MOBILE":                    3,
	"PC":                        4,
	"SETTLEMENT_DERIVED":        5,
	"COMBINATION_DERIVED":       6,
	"MULTI_ROUTE_CLOSE_DERIVED": 7,
	"PLACEMENT_SHARE":           8,
	"REPURCHASE":                9,
	"CTP_UPLOAD":                10,
}

func (x OrderSourceType) String() string {
	return proto.EnumName(OrderSourceType_name, int32(x))
}

func (OrderSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{7}
}

// 委托
type Order struct {
	UserId               string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaId                 string        `protobuf:"bytes,2,opt,name=ta_id,json=taId,proto3" json:"ta_id,omitempty"`
	BuId                 string        `protobuf:"bytes,3,opt,name=bu_id,json=buId,proto3" json:"bu_id,omitempty"`
	Exchange             string        `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol               string        `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Product              string        `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
	OrderRef             string        `protobuf:"bytes,7,opt,name=order_ref,json=orderRef,proto3" json:"order_ref,omitempty"`
	Direction            DirectionType `protobuf:"varint,8,opt,name=direction,proto3,enum=goshare.DirectionType" json:"direction,omitempty"`
	Offset               int32         `protobuf:"varint,9,opt,name=offset,proto3" json:"offset,omitempty"`
	Price                float64       `protobuf:"fixed64,10,opt,name=price,proto3" json:"price,omitempty"`
	Volume               int32         `protobuf:"varint,11,opt,name=volume,proto3" json:"volume,omitempty"`
	VolumeTraded         int32         `protobuf:"varint,12,opt,name=volume_traded,json=volumeTraded,proto3" json:"volume_traded,omitempty"`
	VolumeCanceled       int32         `protobuf:"varint,13,opt,name=volume_canceled,json=volumeCanceled,proto3" json:"volume_canceled,omitempty"`
	Status               int32         `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	PriceType            int32         `protobuf:"varint,15,opt,name=price_type,json=priceType,proto3" json:"price_type,omitempty"`
	FrozenCommission     float64       `protobuf:"fixed64,16,opt,name=frozen_commission,json=frozenCommission,proto3" json:"frozen_commission,omitempty"`
	SendOrderTime        int64         `protobuf:"varint,17,opt,name=send_order_time,json=sendOrderTime,proto3" json:"send_order_time,omitempty"`
	SendOrderTradingDay  int32         `protobuf:"varint,18,opt,name=send_order_trading_day,json=sendOrderTradingDay,proto3" json:"send_order_trading_day,omitempty"`
	StatusMsg            string        `protobuf:"bytes,19,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	ForceCloseType       int32         `protobuf:"varint,20,opt,name=force_close_type,json=forceCloseType,proto3" json:"force_close_type,omitempty"`
	OrderSourceType      int32         `protobuf:"varint,21,opt,name=order_source_type,json=orderSourceType,proto3" json:"order_source_type,omitempty"`
	HedgeType            int32         `protobuf:"varint,22,opt,name=hedge_type,json=hedgeType,proto3" json:"hedge_type,omitempty"`
	UserProductInfo      string        `protobuf:"bytes,23,opt,name=user_product_info,json=userProductInfo,proto3" json:"user_product_info,omitempty"`
	ExchangeOrderId      string        `protobuf:"bytes,24,opt,name=exchange_order_id,json=exchangeOrderId,proto3" json:"exchange_order_id,omitempty"`
	TimeConditionType    int32         `protobuf:"varint,25,opt,name=time_condition_type,json=timeConditionType,proto3" json:"time_condition_type,omitempty"`
	IsPending            bool          `protobuf:"varint,26,opt,name=is_pending,json=isPending,proto3" json:"is_pending,omitempty"`
	Commission           float64       `protobuf:"fixed64,27,opt,name=commission,proto3" json:"commission,omitempty"`
	ProductType          string        `protobuf:"bytes,28,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	TradedAmount         float64       `protobuf:"fixed64,29,opt,name=traded_amount,json=tradedAmount,proto3" json:"traded_amount,omitempty"`
	Multiple             int32         `protobuf:"varint,30,opt,name=multiple,proto3" json:"multiple,omitempty"`
	PriceTick            float64       `protobuf:"fixed64,31,opt,name=price_tick,json=priceTick,proto3" json:"price_tick,omitempty"`
	LastPrice            float64       `protobuf:"fixed64,32,opt,name=last_price,json=lastPrice,proto3" json:"last_price,omitempty"`
	FrozenMargin         float64       `protobuf:"fixed64,33,opt,name=frozen_margin,json=frozenMargin,proto3" json:"frozen_margin,omitempty"`
	SymbolName           string        `protobuf:"bytes,34,opt,name=symbol_name,json=symbolName,proto3" json:"symbol_name,omitempty"`
	Reason               string        `protobuf:"bytes,35,opt,name=reason,proto3" json:"reason,omitempty"`
	StampTax             float64       `protobuf:"fixed64,36,opt,name=stamp_tax,json=stampTax,proto3" json:"stamp_tax,omitempty"`
	TransferFee          float64       `protobuf:"fixed64,37,opt,name=transfer_fee,json=transferFee,proto3" json:"transfer_fee,omitempty"`
	UserName             string        `protobuf:"bytes,38,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Ip                   string        `protobuf:"bytes,39,opt,name=ip,proto3" json:"ip,omitempty"`
	Mac                  string        `protobuf:"bytes,40,opt,name=mac,proto3" json:"mac,omitempty"`
	Branch               string        `protobuf:"bytes,41,opt,name=branch,proto3" json:"branch,omitempty"`
	BranchName           string        `protobuf:"bytes,42,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	BuType               int32         `protobuf:"varint,43,opt,name=bu_type,json=buType,proto3" json:"bu_type,omitempty"`
	Debug                string        `protobuf:"bytes,44,opt,name=debug,proto3" json:"debug,omitempty"`
	SessionId            int32         `protobuf:"varint,45,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ClientTag            int64         `protobuf:"varint,46,opt,name=client_tag,json=clientTag,proto3" json:"client_tag,omitempty"`
	FrontId              int32         `protobuf:"varint,47,opt,name=front_id,json=frontId,proto3" json:"front_id,omitempty"`
	RequestId            int32         `protobuf:"varint,48,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	LocalOrderSeq        int32         `protobuf:"varint,49,opt,name=local_order_seq,json=localOrderSeq,proto3" json:"local_order_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetTaId() string {
	if m != nil {
		return m.TaId
	}
	return ""
}

func (m *Order) GetBuId() string {
	if m != nil {
		return m.BuId
	}
	return ""
}

func (m *Order) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Order) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Order) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Order) GetOrderRef() string {
	if m != nil {
		return m.OrderRef
	}
	return ""
}

func (m *Order) GetDirection() DirectionType {
	if m != nil {
		return m.Direction
	}
	return DirectionType_LONG
}

func (m *Order) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Order) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Order) GetVolumeTraded() int32 {
	if m != nil {
		return m.VolumeTraded
	}
	return 0
}

func (m *Order) GetVolumeCanceled() int32 {
	if m != nil {
		return m.VolumeCanceled
	}
	return 0
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetPriceType() int32 {
	if m != nil {
		return m.PriceType
	}
	return 0
}

func (m *Order) GetFrozenCommission() float64 {
	if m != nil {
		return m.FrozenCommission
	}
	return 0
}

func (m *Order) GetSendOrderTime() int64 {
	if m != nil {
		return m.SendOrderTime
	}
	return 0
}

func (m *Order) GetSendOrderTradingDay() int32 {
	if m != nil {
		return m.SendOrderTradingDay
	}
	return 0
}

func (m *Order) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *Order) GetForceCloseType() int32 {
	if m != nil {
		return m.ForceCloseType
	}
	return 0
}

func (m *Order) GetOrderSourceType() int32 {
	if m != nil {
		return m.OrderSourceType
	}
	return 0
}

func (m *Order) GetHedgeType() int32 {
	if m != nil {
		return m.HedgeType
	}
	return 0
}

func (m *Order) GetUserProductInfo() string {
	if m != nil {
		return m.UserProductInfo
	}
	return ""
}

func (m *Order) GetExchangeOrderId() string {
	if m != nil {
		return m.ExchangeOrderId
	}
	return ""
}

func (m *Order) GetTimeConditionType() int32 {
	if m != nil {
		return m.TimeConditionType
	}
	return 0
}

func (m *Order) GetIsPending() bool {
	if m != nil {
		return m.IsPending
	}
	return false
}

func (m *Order) GetCommission() float64 {
	if m != nil {
		return m.Commission
	}
	return 0
}

func (m *Order) GetProductType() string {
	if m != nil {
		return m.ProductType
	}
	return ""
}

func (m *Order) GetTradedAmount() float64 {
	if m != nil {
		return m.TradedAmount
	}
	return 0
}

func (m *Order) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *Order) GetPriceTick() float64 {
	if m != nil {
		return m.PriceTick
	}
	return 0
}

func (m *Order) GetLastPrice() float64 {
	if m != nil {
		return m.LastPrice
	}
	return 0
}

func (m *Order) GetFrozenMargin() float64 {
	if m != nil {
		return m.FrozenMargin
	}
	return 0
}

func (m *Order) GetSymbolName() string {
	if m != nil {
		return m.SymbolName
	}
	return ""
}

func (m *Order) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Order) GetStampTax() float64 {
	if m != nil {
		return m.StampTax
	}
	return 0
}

func (m *Order) GetTransferFee() float64 {
	if m != nil {
		return m.TransferFee
	}
	return 0
}

func (m *Order) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Order) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Order) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Order) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Order) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Order) GetBuType() int32 {
	if m != nil {
		return m.BuType
	}
	return 0
}

func (m *Order) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

func (m *Order) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Order) GetClientTag() int64 {
	if m != nil {
		return m.ClientTag
	}
	return 0
}

func (m *Order) GetFrontId() int32 {
	if m != nil {
		return m.FrontId
	}
	return 0
}

func (m *Order) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Order) GetLocalOrderSeq() int32 {
	if m != nil {
		return m.LocalOrderSeq
	}
	return 0
}

type OrderList struct {
	List                 []*Order `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderList) Reset()         { *m = OrderList{} }
func (m *OrderList) String() string { return proto.CompactTextString(m) }
func (*OrderList) ProtoMessage()    {}
func (*OrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *OrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderList.Unmarshal(m, b)
}
func (m *OrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderList.Marshal(b, m, deterministic)
}
func (m *OrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderList.Merge(m, src)
}
func (m *OrderList) XXX_Size() int {
	return xxx_messageInfo_OrderList.Size(m)
}
func (m *OrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderList proto.InternalMessageInfo

func (m *OrderList) GetList() []*Order {
	if m != nil {
		return m.List
	}
	return nil
}

type JointOrder struct {
	LocalOrder           *Order   `protobuf:"bytes,1,opt,name=local_order,json=localOrder,proto3" json:"local_order,omitempty"`
	ExchangeOrder        *Order   `protobuf:"bytes,2,opt,name=exchange_order,json=exchangeOrder,proto3" json:"exchange_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JointOrder) Reset()         { *m = JointOrder{} }
func (m *JointOrder) String() string { return proto.CompactTextString(m) }
func (*JointOrder) ProtoMessage()    {}
func (*JointOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *JointOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JointOrder.Unmarshal(m, b)
}
func (m *JointOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JointOrder.Marshal(b, m, deterministic)
}
func (m *JointOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JointOrder.Merge(m, src)
}
func (m *JointOrder) XXX_Size() int {
	return xxx_messageInfo_JointOrder.Size(m)
}
func (m *JointOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_JointOrder.DiscardUnknown(m)
}

var xxx_messageInfo_JointOrder proto.InternalMessageInfo

func (m *JointOrder) GetLocalOrder() *Order {
	if m != nil {
		return m.LocalOrder
	}
	return nil
}

func (m *JointOrder) GetExchangeOrder() *Order {
	if m != nil {
		return m.ExchangeOrder
	}
	return nil
}

func init() {
	proto.RegisterEnum("goshare.HedgeType", HedgeType_name, HedgeType_value)
	proto.RegisterEnum("goshare.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("goshare.OrderPriceType", OrderPriceType_name, OrderPriceType_value)
	proto.RegisterEnum("goshare.OffsetFlag", OffsetFlag_name, OffsetFlag_value)
	proto.RegisterEnum("goshare.DirectionType", DirectionType_name, DirectionType_value)
	proto.RegisterEnum("goshare.CombDirectionType", CombDirectionType_name, CombDirectionType_value)
	proto.RegisterEnum("goshare.ForceCloseReason", ForceCloseReason_name, ForceCloseReason_value)
	proto.RegisterEnum("goshare.OrderSourceType", OrderSourceType_name, OrderSourceType_value)
	proto.RegisterType((*Order)(nil), "goshare.Order")
	proto.RegisterType((*OrderList)(nil), "goshare.OrderList")
	proto.RegisterType((*JointOrder)(nil), "goshare.JointOrder")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 1467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x96, 0x5f, 0x73, 0xdb, 0xb8,
	0x11, 0xc0, 0x43, 0xc9, 0xfa, 0xb7, 0xfa, 0x07, 0xc3, 0x39, 0x07, 0x49, 0x9a, 0x44, 0xe7, 0xb4,
	0x77, 0x3a, 0x5d, 0x63, 0xb7, 0xb9, 0x6b, 0xdf, 0x69, 0x8a, 0xb1, 0xd8, 0x48, 0xa4, 0x86, 0xa2,
	0xef, 0x9a, 0xbe, 0x70, 0x28, 0x12, 0x92, 0x39, 0x16, 0x09, 0x85, 0xa4, 0x3c, 0x71, 0x67, 0xfa,
	0x3d, 0xfa, 0xdc, 0x2f, 0xd5, 0xaf, 0xd3, 0xc1, 0x82, 0x52, 0x9c, 0xe6, 0xde, 0xb8, 0xbf, 0x5d,
	0xec, 0x1f, 0x60, 0x17, 0x20, 0xb4, 0x45, 0x16, 0xf1, 0xec, 0x7c, 0x9b, 0x89, 0x42, 0xd0, 0xc6,
	0x5a, 0xe4, 0x37, 0x41, 0xc6, 0xcf, 0xfe, 0xd3, 0x81, 0x9a, 0x23, 0x15, 0xf4, 0x09, 0x34, 0x76,
	0x39, 0xcf, 0xfc, 0x38, 0x62, 0xda, 0x40, 0x1b, 0xb6, 0xdc, 0xba, 0x14, 0xad, 0x88, 0x9e, 0x40,
	0xad, 0x08, 0x24, 0xae, 0x20, 0x3e, 0x2a, 0x02, 0x05, 0x97, 0x3b, 0x09, 0xab, 0x0a, 0x2e, 0x77,
	0x56, 0x44, 0x9f, 0x41, 0x93, 0x7f, 0x0a, 0x6f, 0x82, 0x74, 0xcd, 0xd9, 0x11, 0xf2, 0x83, 0x4c,
	0x4f, 0xa1, 0x9e, 0xdf, 0x27, 0x4b, 0xb1, 0x61, 0x35, 0xe5, 0x5d, 0x49, 0x94, 0x41, 0x63, 0x9b,
	0x89, 0x68, 0x17, 0x16, 0xac, 0x8e, 0x8a, 0xbd, 0x48, 0x9f, 0x43, 0x0b, 0x53, 0xf6, 0x33, 0xbe,
	0x62, 0x0d, 0xe5, 0x0e, 0x81, 0xcb, 0x57, 0xf4, 0x67, 0x68, 0x45, 0x71, 0xc6, 0xc3, 0x22, 0x16,
	0x29, 0x6b, 0x0e, 0xb4, 0x61, 0xef, 0xed, 0xe9, 0x79, 0x59, 0xd4, 0xf9, 0x78, 0xaf, 0xf1, 0xee,
	0xb7, 0xdc, 0xfd, 0x6c, 0x28, 0x93, 0x10, 0xab, 0x55, 0xce, 0x0b, 0xd6, 0x1a, 0x68, 0xc3, 0x9a,
	0x5b, 0x4a, 0xf4, 0x31, 0xd4, 0xb6, 0x59, 0x1c, 0x72, 0x06, 0x03, 0x6d, 0xa8, 0xb9, 0x4a, 0x90,
	0xd6, 0x77, 0x62, 0xb3, 0x4b, 0x38, 0x6b, 0x2b, 0x6b, 0x25, 0xd1, 0xd7, 0xd0, 0x55, 0x5f, 0x7e,
	0x91, 0x05, 0x11, 0x8f, 0x58, 0x07, 0xd5, 0x1d, 0x05, 0x3d, 0x64, 0xf4, 0x7b, 0xe8, 0x97, 0x46,
	0x61, 0x90, 0x86, 0x7c, 0xc3, 0x23, 0xd6, 0x45, 0xb3, 0x9e, 0xc2, 0x46, 0x49, 0x71, 0x63, 0x8a,
	0xa0, 0xd8, 0xe5, 0xac, 0xa7, 0xa2, 0x28, 0x89, 0xbe, 0x00, 0xc0, 0x34, 0xfc, 0xe2, 0x7e, 0xcb,
	0x59, 0x1f, 0x75, 0x2d, 0x24, 0xb2, 0x2a, 0xfa, 0x23, 0x1c, 0xaf, 0x32, 0xf1, 0x4f, 0x9e, 0xfa,
	0xa1, 0x48, 0x92, 0x38, 0xcf, 0xe5, 0x46, 0x10, 0x4c, 0x9f, 0x28, 0x85, 0x71, 0xe0, 0xf4, 0x3b,
	0xe8, 0xe7, 0x3c, 0x8d, 0x7c, 0xb5, 0x9f, 0x45, 0x9c, 0x70, 0x76, 0x3c, 0xd0, 0x86, 0x55, 0xb7,
	0x2b, 0x31, 0x9e, 0xbf, 0x17, 0x27, 0x9c, 0xfe, 0x04, 0xa7, 0x0f, 0xed, 0xb2, 0x20, 0x8a, 0xd3,
	0xb5, 0x1f, 0x05, 0xf7, 0x8c, 0x62, 0xfc, 0x93, 0xcf, 0xe6, 0x4a, 0x37, 0x0e, 0xee, 0x65, 0xa2,
	0x2a, 0x65, 0x3f, 0xc9, 0xd7, 0xec, 0x04, 0x0f, 0xaa, 0xa5, 0xc8, 0x2c, 0x5f, 0xd3, 0x21, 0x90,
	0x95, 0xc8, 0x42, 0xee, 0x87, 0x1b, 0x91, 0x97, 0xd5, 0x3c, 0x56, 0x3b, 0x81, 0xdc, 0x90, 0x18,
	0x4b, 0x1a, 0xc1, 0xb1, 0x0a, 0x9c, 0x8b, 0x5d, 0xb6, 0x2f, 0xfc, 0x1b, 0x34, 0xed, 0xa3, 0x62,
	0x81, 0x1c, 0x6d, 0x5f, 0x00, 0xdc, 0xf0, 0x68, 0x5d, 0x1a, 0x9d, 0xaa, 0xdd, 0x41, 0xb2, 0x77,
	0x85, 0xcd, 0x5c, 0xf6, 0x92, 0x1f, 0xa7, 0x2b, 0xc1, 0x9e, 0x60, 0x6a, 0x7d, 0xa9, 0x98, 0x2b,
	0x6e, 0xa5, 0x2b, 0x21, 0x6d, 0xf7, 0x5d, 0x5a, 0x16, 0x1e, 0x47, 0x8c, 0x29, 0xdb, 0xbd, 0x02,
	0x6b, 0xb6, 0x22, 0x7a, 0x0e, 0x27, 0x72, 0xf7, 0xfc, 0x50, 0xa4, 0x51, 0x2c, 0x5b, 0x4a, 0xc5,
	0x7f, 0x8a, 0xf1, 0x8f, 0xa5, 0xca, 0xd8, 0x6b, 0xf6, 0x69, 0xc6, 0xb9, 0xbf, 0xe5, 0xa9, 0xdc,
	0x2c, 0xf6, 0x6c, 0xa0, 0x0d, 0x9b, 0x6e, 0x2b, 0xce, 0xe7, 0x0a, 0xd0, 0x97, 0x00, 0x0f, 0x4e,
	0xef, 0x39, 0x9e, 0xde, 0x03, 0x42, 0xbf, 0x85, 0xce, 0xbe, 0x02, 0x8c, 0xf3, 0x3b, 0xcc, 0xaa,
	0x5d, 0x32, 0x8c, 0xf0, 0x1a, 0xba, 0xaa, 0x0b, 0xfd, 0x20, 0x11, 0xbb, 0xb4, 0x60, 0x2f, 0xd0,
	0x4b, 0x47, 0x41, 0x1d, 0x99, 0x1c, 0xcc, 0x64, 0xb7, 0x29, 0xe2, 0xed, 0x86, 0xb3, 0x97, 0x98,
	0xeb, 0x41, 0x7e, 0xd0, 0x67, 0x71, 0x78, 0xcb, 0x5e, 0xe1, 0xea, 0xb2, 0xcf, 0xe2, 0xf0, 0x56,
	0xaa, 0x37, 0x41, 0x5e, 0xf8, 0x6a, 0x3e, 0x06, 0x4a, 0x2d, 0xc9, 0x1c, 0x67, 0xe4, 0x35, 0x74,
	0xcb, 0x36, 0x4c, 0x82, 0x6c, 0x1d, 0xa7, 0xec, 0x5b, 0x15, 0x5e, 0xc1, 0x19, 0x32, 0xfa, 0x0a,
	0xda, 0x6a, 0xda, 0xfd, 0x34, 0x48, 0x38, 0x3b, 0xc3, 0x2a, 0x40, 0x21, 0x3b, 0x48, 0x70, 0xd2,
	0x32, 0x1e, 0xe4, 0x22, 0x65, 0xaf, 0xd5, 0xe5, 0xa0, 0x24, 0x79, 0x05, 0xe4, 0x45, 0x90, 0x6c,
	0xfd, 0x22, 0xf8, 0xc4, 0x7e, 0x8f, 0x9e, 0x9b, 0x08, 0xbc, 0xe0, 0x93, 0xdc, 0x9c, 0x22, 0x0b,
	0xd2, 0x7c, 0xc5, 0x33, 0x7f, 0xc5, 0x39, 0xfb, 0x03, 0xea, 0xdb, 0x7b, 0xf6, 0x8e, 0x73, 0xb9,
	0x1e, 0xdb, 0x00, 0xc3, 0x7e, 0xa7, 0xae, 0x10, 0x09, 0x30, 0x68, 0x0f, 0x2a, 0xf1, 0x96, 0x7d,
	0x8f, 0xb4, 0x12, 0x6f, 0x29, 0x81, 0x6a, 0x12, 0x84, 0x6c, 0x88, 0x40, 0x7e, 0xca, 0xb4, 0x96,
	0x59, 0x90, 0x86, 0x37, 0xec, 0x07, 0x95, 0x96, 0x92, 0x64, 0x3d, 0xea, 0x4b, 0x39, 0x1e, 0xa9,
	0x7a, 0x14, 0x42, 0xd7, 0x4f, 0xa0, 0xb1, 0xdc, 0xa9, 0x23, 0xfb, 0x51, 0x0d, 0xf5, 0x72, 0x87,
	0xa7, 0xf5, 0x18, 0x6a, 0x11, 0x5f, 0xee, 0xd6, 0xec, 0x8f, 0xb8, 0x46, 0x09, 0x38, 0x41, 0x1c,
	0x4f, 0x5c, 0xb6, 0xde, 0x1b, 0xd5, 0xcc, 0x25, 0xb1, 0x22, 0xa9, 0x0e, 0x37, 0x31, 0x4f, 0x0b,
	0xbf, 0x08, 0xd6, 0xec, 0x1c, 0x07, 0xb7, 0xa5, 0x88, 0x17, 0xac, 0xe9, 0x53, 0x68, 0xae, 0x32,
	0x91, 0x16, 0x72, 0xed, 0x05, 0xae, 0x6d, 0xa0, 0xac, 0x56, 0x66, 0xfc, 0xe3, 0x8e, 0xe7, 0xa8,
	0xfc, 0x93, 0x72, 0x5c, 0x12, 0x2b, 0x92, 0xd7, 0xc2, 0x46, 0x84, 0xc1, 0xa6, 0x6c, 0xfb, 0x9c,
	0x7f, 0x64, 0x7f, 0x46, 0x9b, 0x2e, 0x62, 0x6c, 0xfa, 0x05, 0xff, 0x78, 0x76, 0x01, 0x2d, 0xfc,
	0x9e, 0xc6, 0x79, 0x41, 0xcf, 0xe0, 0x68, 0x13, 0xe7, 0x05, 0xd3, 0x06, 0xd5, 0x61, 0xfb, 0x6d,
	0xef, 0x70, 0xe9, 0xa2, 0x85, 0x8b, 0xba, 0xb3, 0x02, 0xe0, 0x6f, 0x22, 0x4e, 0x0b, 0xf5, 0xb2,
	0x5c, 0x40, 0xfb, 0x41, 0x18, 0x7c, 0x5d, 0xbe, 0x5e, 0x08, 0x9f, 0x43, 0xd2, 0xbf, 0x40, 0xef,
	0xcb, 0x89, 0xc4, 0xa7, 0xe7, 0xeb, 0x35, 0xdd, 0x2f, 0xc6, 0x73, 0x34, 0x81, 0xd6, 0xe4, 0x70,
	0x03, 0xf4, 0xa1, 0xbd, 0x98, 0x9b, 0xc6, 0xf5, 0x54, 0xf7, 0x2c, 0xc7, 0x26, 0x8f, 0x68, 0x17,
	0x5a, 0xba, 0x7b, 0x69, 0x79, 0xae, 0x7e, 0x65, 0x12, 0x8d, 0xb6, 0xa0, 0x36, 0x31, 0xc7, 0x57,
	0x26, 0xa9, 0x50, 0x02, 0x9d, 0x07, 0xa6, 0x6f, 0x49, 0x75, 0xf4, 0x2f, 0x68, 0xab, 0xe2, 0xd5,
	0x55, 0xdc, 0x84, 0xa3, 0xb1, 0x63, 0x9b, 0xe4, 0x11, 0x7d, 0x01, 0x4f, 0xe7, 0xa6, 0x3d, 0xb6,
	0xec, 0x2b, 0xff, 0x57, 0xcb, 0x9b, 0xf8, 0x73, 0xdd, 0xf5, 0x2c, 0x7d, 0xea, 0xa3, 0x5a, 0xa3,
	0x2f, 0xe1, 0x99, 0xa1, 0xdb, 0x86, 0x39, 0x35, 0xc7, 0xbf, 0xa1, 0xaf, 0xd0, 0x36, 0x34, 0xca,
	0xe5, 0xa4, 0x4a, 0x3b, 0xd0, 0xdc, 0x1b, 0x93, 0x23, 0x0a, 0x50, 0xbf, 0xb6, 0xdf, 0x3b, 0xbf,
	0xda, 0xa4, 0x36, 0x5a, 0x40, 0x0f, 0xc3, 0xcf, 0x0f, 0xb7, 0x7d, 0x1f, 0xda, 0x53, 0x6b, 0x66,
	0x79, 0xfe, 0xdc, 0xb5, 0x0c, 0x99, 0x08, 0x81, 0xce, 0x4c, 0x77, 0xdf, 0x9b, 0x7b, 0xa2, 0xd1,
	0x1e, 0xc0, 0xa5, 0xb9, 0xd8, 0xcb, 0x15, 0x29, 0x4f, 0xf5, 0x83, 0x5c, 0x1d, 0xfd, 0x1d, 0xc0,
	0xc1, 0xd7, 0xee, 0xdd, 0x26, 0x58, 0xcb, 0x92, 0x9c, 0xb9, 0x29, 0xf7, 0xa5, 0x05, 0x35, 0x63,
	0xea, 0x2c, 0xa4, 0x8b, 0x3e, 0xb4, 0xdf, 0x39, 0xae, 0x61, 0xfa, 0x0a, 0x54, 0x24, 0xc0, 0x4f,
	0xdf, 0x73, 0xc6, 0xfa, 0x07, 0x52, 0xa5, 0x27, 0xd0, 0x57, 0xe0, 0x83, 0xb9, 0xf0, 0x4c, 0x57,
	0xc2, 0xa3, 0xd1, 0x1b, 0xe8, 0x7e, 0xf1, 0xe2, 0x4a, 0xe7, 0x53, 0xc7, 0xbe, 0x52, 0xce, 0x17,
	0x13, 0xc7, 0xf5, 0x88, 0x46, 0x1b, 0x50, 0xb5, 0x4d, 0x8f, 0x54, 0x46, 0x3f, 0xc0, 0xb1, 0x21,
	0x92, 0xe5, 0x57, 0x4b, 0x0c, 0x67, 0x76, 0x49, 0x1e, 0xa9, 0x8d, 0xc0, 0x6f, 0x6d, 0xf4, 0x6f,
	0x0d, 0xc8, 0xbb, 0xc3, 0x23, 0xe1, 0xaa, 0x4b, 0xe1, 0x1b, 0x38, 0xb6, 0x1d, 0xcf, 0x37, 0x6d,
	0xe7, 0xfa, 0x6a, 0xe2, 0xcf, 0x74, 0xf7, 0xca, 0x92, 0x75, 0xbc, 0x84, 0x67, 0xe5, 0x94, 0x88,
	0x3b, 0x79, 0xf3, 0x8b, 0x5c, 0x5d, 0xd0, 0x9b, 0x38, 0x89, 0x0b, 0xa2, 0xd1, 0x57, 0xf0, 0x7c,
	0x99, 0x89, 0x5b, 0x9e, 0xed, 0xf5, 0xf1, 0x17, 0x06, 0x15, 0xd9, 0x20, 0x77, 0xb1, 0xd8, 0x04,
	0x12, 0x92, 0xaa, 0x4c, 0x5d, 0x14, 0x37, 0x3c, 0x23, 0x47, 0x94, 0x42, 0x6f, 0xcb, 0xb3, 0x5c,
	0xa4, 0x7e, 0xc4, 0x37, 0xf1, 0x1d, 0xcf, 0x48, 0x6d, 0xf4, 0x5f, 0x0d, 0xfa, 0xce, 0xff, 0x3d,
	0x4a, 0x00, 0x75, 0x63, 0x6a, 0x99, 0xb6, 0x47, 0x1e, 0xd1, 0x63, 0xe8, 0xea, 0xe3, 0x99, 0x65,
	0x5b, 0x0b, 0xcf, 0xd5, 0x3d, 0xc7, 0x25, 0x9a, 0x3c, 0x70, 0xd7, 0x5a, 0xbc, 0xf7, 0x17, 0xee,
	0x2f, 0xa4, 0x22, 0x8d, 0x67, 0xce, 0xa5, 0x35, 0x35, 0x49, 0x95, 0xd6, 0xa1, 0x32, 0x37, 0xc8,
	0x11, 0x3d, 0x05, 0xba, 0x30, 0x3d, 0x6f, 0x6a, 0xce, 0x4c, 0xdb, 0xf3, 0xc7, 0xa6, 0x6b, 0xfd,
	0x62, 0x8e, 0x49, 0x8d, 0x3e, 0x81, 0x13, 0xb9, 0x23, 0x96, 0x8d, 0x1d, 0x7a, 0x50, 0xd4, 0x65,
	0x3f, 0xce, 0xae, 0xa7, 0x9e, 0xe5, 0xbb, 0xce, 0xb5, 0x57, 0x9e, 0xdb, 0x41, 0xdd, 0x90, 0xc7,
	0x35, 0x9f, 0xea, 0x86, 0x72, 0xb7, 0x98, 0xe8, 0xae, 0x49, 0x9a, 0xb2, 0x31, 0x5c, 0x73, 0x7e,
	0xed, 0x1a, 0x13, 0x7d, 0x61, 0x92, 0x96, 0x94, 0x0d, 0x6f, 0xee, 0x5f, 0xcf, 0xa7, 0x8e, 0x3e,
	0x26, 0x70, 0xf9, 0xd7, 0x7f, 0xfc, 0xbc, 0x8e, 0x8b, 0x9b, 0xdd, 0xf2, 0x3c, 0x14, 0xc9, 0x85,
	0x1b, 0xdc, 0xf1, 0x74, 0xb2, 0x13, 0x17, 0x6b, 0xf1, 0x06, 0xff, 0x1d, 0x2f, 0xf2, 0x42, 0x84,
	0xb7, 0xe5, 0xf7, 0xf6, 0x76, 0x7d, 0x51, 0xce, 0xe4, 0xb2, 0x8e, 0xe8, 0xa7, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x01, 0xc3, 0x1b, 0x69, 0x0a, 0x00, 0x00,
}
