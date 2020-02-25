// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type CreateOrderParams struct {
	OrderItems           []*OrderItem `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	UserId               int64        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderParams) Reset()         { *m = CreateOrderParams{} }
func (m *CreateOrderParams) String() string { return proto.CompactTextString(m) }
func (*CreateOrderParams) ProtoMessage()    {}
func (*CreateOrderParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CreateOrderParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderParams.Unmarshal(m, b)
}
func (m *CreateOrderParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderParams.Marshal(b, m, deterministic)
}
func (m *CreateOrderParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderParams.Merge(m, src)
}
func (m *CreateOrderParams) XXX_Size() int {
	return xxx_messageInfo_CreateOrderParams.Size(m)
}
func (m *CreateOrderParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderParams proto.InternalMessageInfo

func (m *CreateOrderParams) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func (m *CreateOrderParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateOrderResponse struct {
	Order                *Order       `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	OrderTxs             []*OrderTx   `protobuf:"bytes,3,rep,name=order_txs,json=orderTxs,proto3" json:"order_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *CreateOrderResponse) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func (m *CreateOrderResponse) GetOrderTxs() []*OrderTx {
	if m != nil {
		return m.OrderTxs
	}
	return nil
}

type Order struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
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

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetOrderParams struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderParams) Reset()         { *m = GetOrderParams{} }
func (m *GetOrderParams) String() string { return proto.CompactTextString(m) }
func (*GetOrderParams) ProtoMessage()    {}
func (*GetOrderParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *GetOrderParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderParams.Unmarshal(m, b)
}
func (m *GetOrderParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderParams.Marshal(b, m, deterministic)
}
func (m *GetOrderParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderParams.Merge(m, src)
}
func (m *GetOrderParams) XXX_Size() int {
	return xxx_messageInfo_GetOrderParams.Size(m)
}
func (m *GetOrderParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderParams proto.InternalMessageInfo

func (m *GetOrderParams) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *GetOrderParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetOrdersParams struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            int64    `protobuf:"varint,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrdersParams) Reset()         { *m = GetOrdersParams{} }
func (m *GetOrdersParams) String() string { return proto.CompactTextString(m) }
func (*GetOrdersParams) ProtoMessage()    {}
func (*GetOrdersParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *GetOrdersParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrdersParams.Unmarshal(m, b)
}
func (m *GetOrdersParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrdersParams.Marshal(b, m, deterministic)
}
func (m *GetOrdersParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrdersParams.Merge(m, src)
}
func (m *GetOrdersParams) XXX_Size() int {
	return xxx_messageInfo_GetOrdersParams.Size(m)
}
func (m *GetOrdersParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrdersParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrdersParams proto.InternalMessageInfo

func (m *GetOrdersParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrdersParams) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetOrdersParams) GetPageToken() int64 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

type GetOrdersResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	NextPageToken        int64    `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrdersResponse) Reset()         { *m = GetOrdersResponse{} }
func (m *GetOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrdersResponse) ProtoMessage()    {}
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *GetOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrdersResponse.Unmarshal(m, b)
}
func (m *GetOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrdersResponse.Marshal(b, m, deterministic)
}
func (m *GetOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrdersResponse.Merge(m, src)
}
func (m *GetOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrdersResponse.Size(m)
}
func (m *GetOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrdersResponse proto.InternalMessageInfo

func (m *GetOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *GetOrdersResponse) GetNextPageToken() int64 {
	if m != nil {
		return m.NextPageToken
	}
	return 0
}

type GetOrderTxsParams struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderTxsParams) Reset()         { *m = GetOrderTxsParams{} }
func (m *GetOrderTxsParams) String() string { return proto.CompactTextString(m) }
func (*GetOrderTxsParams) ProtoMessage()    {}
func (*GetOrderTxsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{6}
}

func (m *GetOrderTxsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderTxsParams.Unmarshal(m, b)
}
func (m *GetOrderTxsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderTxsParams.Marshal(b, m, deterministic)
}
func (m *GetOrderTxsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderTxsParams.Merge(m, src)
}
func (m *GetOrderTxsParams) XXX_Size() int {
	return xxx_messageInfo_GetOrderTxsParams.Size(m)
}
func (m *GetOrderTxsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderTxsParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderTxsParams proto.InternalMessageInfo

func (m *GetOrderTxsParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrderTxsParams) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type OrderTx struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderTx) Reset()         { *m = OrderTx{} }
func (m *OrderTx) String() string { return proto.CompactTextString(m) }
func (*OrderTx) ProtoMessage()    {}
func (*OrderTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{7}
}

func (m *OrderTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderTx.Unmarshal(m, b)
}
func (m *OrderTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderTx.Marshal(b, m, deterministic)
}
func (m *OrderTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderTx.Merge(m, src)
}
func (m *OrderTx) XXX_Size() int {
	return xxx_messageInfo_OrderTx.Size(m)
}
func (m *OrderTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderTx.DiscardUnknown(m)
}

var xxx_messageInfo_OrderTx proto.InternalMessageInfo

func (m *OrderTx) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderTx) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderTx) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderTx) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetOrderTxsResponse struct {
	OrderTxs             []*OrderTx `protobuf:"bytes,1,rep,name=order_txs,json=orderTxs,proto3" json:"order_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetOrderTxsResponse) Reset()         { *m = GetOrderTxsResponse{} }
func (m *GetOrderTxsResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderTxsResponse) ProtoMessage()    {}
func (*GetOrderTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{8}
}

func (m *GetOrderTxsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderTxsResponse.Unmarshal(m, b)
}
func (m *GetOrderTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderTxsResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderTxsResponse.Merge(m, src)
}
func (m *GetOrderTxsResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderTxsResponse.Size(m)
}
func (m *GetOrderTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderTxsResponse proto.InternalMessageInfo

func (m *GetOrderTxsResponse) GetOrderTxs() []*OrderTx {
	if m != nil {
		return m.OrderTxs
	}
	return nil
}

type GetOrderItemsParams struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderItemsParams) Reset()         { *m = GetOrderItemsParams{} }
func (m *GetOrderItemsParams) String() string { return proto.CompactTextString(m) }
func (*GetOrderItemsParams) ProtoMessage()    {}
func (*GetOrderItemsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{9}
}

func (m *GetOrderItemsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderItemsParams.Unmarshal(m, b)
}
func (m *GetOrderItemsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderItemsParams.Marshal(b, m, deterministic)
}
func (m *GetOrderItemsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderItemsParams.Merge(m, src)
}
func (m *GetOrderItemsParams) XXX_Size() int {
	return xxx_messageInfo_GetOrderItemsParams.Size(m)
}
func (m *GetOrderItemsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderItemsParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderItemsParams proto.InternalMessageInfo

func (m *GetOrderItemsParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrderItemsParams) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type OrderItem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PriceId              int64    `protobuf:"varint,2,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	Quantity             uint32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	OrderId              int64    `protobuf:"varint,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PlanId               int64    `protobuf:"varint,5,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{10}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderItem) GetPriceId() int64 {
	if m != nil {
		return m.PriceId
	}
	return 0
}

func (m *OrderItem) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderItem) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderItem) GetPlanId() int64 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

type GetOrderItemsResponse struct {
	OrderItems           []*OrderItem `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOrderItemsResponse) Reset()         { *m = GetOrderItemsResponse{} }
func (m *GetOrderItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderItemsResponse) ProtoMessage()    {}
func (*GetOrderItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{11}
}

func (m *GetOrderItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderItemsResponse.Unmarshal(m, b)
}
func (m *GetOrderItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderItemsResponse.Merge(m, src)
}
func (m *GetOrderItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderItemsResponse.Size(m)
}
func (m *GetOrderItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderItemsResponse proto.InternalMessageInfo

func (m *GetOrderItemsResponse) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrderParams)(nil), "pepeunlimited.billing.CreateOrderParams")
	proto.RegisterType((*CreateOrderResponse)(nil), "pepeunlimited.billing.CreateOrderResponse")
	proto.RegisterType((*Order)(nil), "pepeunlimited.billing.Order")
	proto.RegisterType((*GetOrderParams)(nil), "pepeunlimited.billing.GetOrderParams")
	proto.RegisterType((*GetOrdersParams)(nil), "pepeunlimited.billing.GetOrdersParams")
	proto.RegisterType((*GetOrdersResponse)(nil), "pepeunlimited.billing.GetOrdersResponse")
	proto.RegisterType((*GetOrderTxsParams)(nil), "pepeunlimited.billing.GetOrderTxsParams")
	proto.RegisterType((*OrderTx)(nil), "pepeunlimited.billing.OrderTx")
	proto.RegisterType((*GetOrderTxsResponse)(nil), "pepeunlimited.billing.GetOrderTxsResponse")
	proto.RegisterType((*GetOrderItemsParams)(nil), "pepeunlimited.billing.GetOrderItemsParams")
	proto.RegisterType((*OrderItem)(nil), "pepeunlimited.billing.OrderItem")
	proto.RegisterType((*GetOrderItemsResponse)(nil), "pepeunlimited.billing.GetOrderItemsResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0x95, 0x71, 0xf8, 0xf1, 0xf0, 0x91, 0x28, 0x1b, 0xe5, 0x2b, 0x75, 0xda, 0x08, 0x59, 0x6a,
	0x84, 0xaa, 0x8a, 0x48, 0xb4, 0x77, 0xbd, 0x4a, 0x5b, 0x29, 0xe2, 0x8a, 0xc8, 0xe1, 0x2a, 0x52,
	0x85, 0x0c, 0x9e, 0xa0, 0x55, 0xc0, 0x76, 0xbc, 0x4b, 0x4b, 0xf2, 0x00, 0x7d, 0xb7, 0x3e, 0x4f,
	0x5f, 0xa0, 0xda, 0xb5, 0xbd, 0x59, 0x03, 0xc6, 0x6d, 0x73, 0xe7, 0x9d, 0x9f, 0x73, 0x66, 0xce,
	0xcc, 0xae, 0xa1, 0x19, 0xc6, 0x3e, 0xc6, 0xbd, 0x28, 0x0e, 0x79, 0x48, 0x8e, 0x23, 0x8c, 0x70,
	0x19, 0xcc, 0xe9, 0x82, 0x72, 0xf4, 0x7b, 0x13, 0x3a, 0x9f, 0xd3, 0x60, 0x66, 0x9f, 0xcc, 0xc2,
	0x70, 0x36, 0xc7, 0x73, 0x19, 0x34, 0x59, 0xde, 0x9e, 0xe3, 0x22, 0xe2, 0x0f, 0x49, 0x8e, 0x7d,
	0xba, 0xee, 0xfc, 0x1e, 0x7b, 0x51, 0x84, 0x31, 0x4b, 0xfc, 0x4e, 0x08, 0x87, 0x9f, 0x63, 0xf4,
	0x38, 0x0e, 0x05, 0xd1, 0x95, 0x17, 0x7b, 0x0b, 0x46, 0x2e, 0x52, 0xde, 0x31, 0xe5, 0xb8, 0x60,
	0x6d, 0xa3, 0x63, 0x76, 0x9b, 0xfd, 0x4e, 0x6f, 0x2b, 0x7d, 0x4f, 0x26, 0x0e, 0x38, 0x2e, 0x5c,
	0x08, 0xb3, 0x4f, 0x46, 0x5e, 0x40, 0x7d, 0xc9, 0x04, 0x82, 0xdf, 0xae, 0x74, 0x8c, 0xae, 0xe9,
	0xd6, 0xc4, 0x71, 0xe0, 0x3b, 0x3f, 0x0d, 0x38, 0xd2, 0x18, 0x5d, 0x64, 0x51, 0x18, 0x30, 0x24,
	0x7d, 0xa8, 0xca, 0xf4, 0xb6, 0xd1, 0x31, 0xba, 0xcd, 0xfe, 0xab, 0x5d, 0x6c, 0x6e, 0x12, 0xba,
	0x5e, 0x67, 0xe5, 0x1f, 0xea, 0xfc, 0x08, 0x56, 0x02, 0xc1, 0x57, 0xac, 0x6d, 0x4a, 0x80, 0xd3,
	0x5d, 0x00, 0xa3, 0x95, 0xdb, 0x08, 0x93, 0x0f, 0xe6, 0x0c, 0xa1, 0x2a, 0x8d, 0x64, 0x1f, 0x2a,
	0xd4, 0x97, 0x95, 0x9b, 0x6e, 0x85, 0xfa, 0xe4, 0x35, 0xc0, 0x54, 0xf6, 0xe8, 0x8f, 0x3d, 0x2e,
	0x05, 0xb0, 0x5c, 0x2b, 0xb5, 0x5c, 0x70, 0x5d, 0x1c, 0x33, 0x27, 0xce, 0x17, 0xd8, 0xbf, 0x44,
	0xae, 0x8f, 0xe2, 0x25, 0x34, 0xd2, 0x16, 0x33, 0xfc, 0x7a, 0x52, 0xbd, 0x5f, 0x2c, 0xf1, 0x2d,
	0x1c, 0x64, 0x28, 0x2c, 0x85, 0xd1, 0x62, 0x0d, 0x3d, 0x96, 0x9c, 0x80, 0x15, 0x79, 0x33, 0x1c,
	0x33, 0xfa, 0x88, 0x12, 0xa6, 0xea, 0x36, 0x84, 0xe1, 0x9a, 0x3e, 0xa2, 0x68, 0x43, 0x3a, 0x79,
	0x78, 0x87, 0x41, 0x5a, 0xaa, 0x0c, 0x1f, 0x09, 0x83, 0x73, 0x0f, 0x87, 0x8a, 0x47, 0xcd, 0xf1,
	0x03, 0xd4, 0x64, 0x81, 0xd9, 0xda, 0xec, 0x1e, 0x64, 0x1a, 0x4b, 0xce, 0xe0, 0x20, 0xc0, 0x15,
	0x1f, 0x6b, 0x74, 0x49, 0x4f, 0x2d, 0x61, 0xbe, 0x52, 0x94, 0x97, 0x4f, 0x94, 0xa3, 0x55, 0x69,
	0x73, 0xba, 0x78, 0x95, 0x9c, 0x78, 0xce, 0x1d, 0xd4, 0x53, 0x94, 0x8d, 0xe1, 0x15, 0x67, 0xad,
	0xcd, 0xd5, 0x5c, 0x9f, 0xeb, 0xff, 0x50, 0x63, 0xdc, 0xe3, 0x4b, 0xd6, 0xde, 0x93, 0xae, 0xf4,
	0xe4, 0xb8, 0x70, 0xa4, 0x55, 0xad, 0xa4, 0xca, 0xed, 0x9e, 0xf1, 0x97, 0xbb, 0x37, 0x78, 0xc2,
	0x94, 0x9b, 0xfc, 0x0c, 0x2d, 0x7e, 0x18, 0x60, 0x29, 0xa0, 0x6d, 0x72, 0x44, 0x31, 0x9d, 0xa2,
	0x96, 0x28, 0xcf, 0x03, 0x9f, 0xd8, 0xd0, 0xb8, 0x5f, 0x7a, 0x01, 0xa7, 0xfc, 0x41, 0x8a, 0xd1,
	0x72, 0xd5, 0x39, 0xc7, 0xb7, 0xb7, 0xb1, 0xb8, 0xd1, 0xdc, 0x0b, 0x84, 0xa7, 0x9a, 0xd4, 0x28,
	0x8e, 0x03, 0xdf, 0xb9, 0x81, 0xe3, 0x5c, 0x4f, 0x4a, 0xa9, 0xe7, 0x3f, 0x48, 0xfd, 0x5f, 0x26,
	0xfc, 0x27, 0x3d, 0xd7, 0x18, 0x7f, 0xa3, 0x53, 0x24, 0x53, 0x68, 0x6a, 0xef, 0x10, 0xe9, 0x16,
	0xa0, 0x6d, 0xbc, 0x8e, 0xf6, 0xdb, 0xf2, 0x48, 0x55, 0xf8, 0x57, 0xb0, 0xd4, 0x15, 0x21, 0x67,
	0x05, 0x89, 0x6b, 0x97, 0xd5, 0xee, 0x96, 0xc5, 0x29, 0xf8, 0x21, 0x34, 0x32, 0x23, 0x79, 0x53,
	0x92, 0x95, 0x82, 0xef, 0xbc, 0x8f, 0x42, 0x14, 0x6d, 0x53, 0x49, 0x59, 0x25, 0xea, 0x0e, 0x16,
	0x8a, 0xb2, 0x6d, 0xef, 0x29, 0xb4, 0x72, 0x63, 0x26, 0x65, 0xc9, 0xda, 0x82, 0xdb, 0xef, 0xfe,
	0x24, 0x36, 0xa3, 0xfa, 0x54, 0xbf, 0x49, 0x7e, 0x15, 0x93, 0x9a, 0xfc, 0xdd, 0xbd, 0xff, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x8f, 0x00, 0xc4, 0x51, 0x07, 0x00, 0x00,
}