// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plan.proto

package planrpc

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

type CreatePlanParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlanParams) Reset()         { *m = CreatePlanParams{} }
func (m *CreatePlanParams) String() string { return proto.CompactTextString(m) }
func (*CreatePlanParams) ProtoMessage()    {}
func (*CreatePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{0}
}

func (m *CreatePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlanParams.Unmarshal(m, b)
}
func (m *CreatePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlanParams.Marshal(b, m, deterministic)
}
func (m *CreatePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlanParams.Merge(m, src)
}
func (m *CreatePlanParams) XXX_Size() int {
	return xxx_messageInfo_CreatePlanParams.Size(m)
}
func (m *CreatePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlanParams proto.InternalMessageInfo

type Plan struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PriceId              int64    `protobuf:"varint,2,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	TitleI18NId          int64    `protobuf:"varint,3,opt,name=title_i18n_id,json=titleI18nId,proto3" json:"title_i18n_id,omitempty"`
	ProductId            int64    `protobuf:"varint,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	StartAt              string   `protobuf:"bytes,6,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,7,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{1}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Plan) GetPriceId() int64 {
	if m != nil {
		return m.PriceId
	}
	return 0
}

func (m *Plan) GetTitleI18NId() int64 {
	if m != nil {
		return m.TitleI18NId
	}
	return 0
}

func (m *Plan) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Plan) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Plan) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *Plan) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

type UpdatePlanParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlanParams) Reset()         { *m = UpdatePlanParams{} }
func (m *UpdatePlanParams) String() string { return proto.CompactTextString(m) }
func (*UpdatePlanParams) ProtoMessage()    {}
func (*UpdatePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{2}
}

func (m *UpdatePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlanParams.Unmarshal(m, b)
}
func (m *UpdatePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlanParams.Marshal(b, m, deterministic)
}
func (m *UpdatePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlanParams.Merge(m, src)
}
func (m *UpdatePlanParams) XXX_Size() int {
	return xxx_messageInfo_UpdatePlanParams.Size(m)
}
func (m *UpdatePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlanParams proto.InternalMessageInfo

type DeletePlanParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlanParams) Reset()         { *m = DeletePlanParams{} }
func (m *DeletePlanParams) String() string { return proto.CompactTextString(m) }
func (*DeletePlanParams) ProtoMessage()    {}
func (*DeletePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{3}
}

func (m *DeletePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlanParams.Unmarshal(m, b)
}
func (m *DeletePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlanParams.Marshal(b, m, deterministic)
}
func (m *DeletePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlanParams.Merge(m, src)
}
func (m *DeletePlanParams) XXX_Size() int {
	return xxx_messageInfo_DeletePlanParams.Size(m)
}
func (m *DeletePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlanParams proto.InternalMessageInfo

type GetPlanParams struct {
	PlanId               int64    `protobuf:"varint,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlanParams) Reset()         { *m = GetPlanParams{} }
func (m *GetPlanParams) String() string { return proto.CompactTextString(m) }
func (*GetPlanParams) ProtoMessage()    {}
func (*GetPlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{4}
}

func (m *GetPlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlanParams.Unmarshal(m, b)
}
func (m *GetPlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlanParams.Marshal(b, m, deterministic)
}
func (m *GetPlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlanParams.Merge(m, src)
}
func (m *GetPlanParams) XXX_Size() int {
	return xxx_messageInfo_GetPlanParams.Size(m)
}
func (m *GetPlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlanParams proto.InternalMessageInfo

func (m *GetPlanParams) GetPlanId() int64 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

type GetPlansParams struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlansParams) Reset()         { *m = GetPlansParams{} }
func (m *GetPlansParams) String() string { return proto.CompactTextString(m) }
func (*GetPlansParams) ProtoMessage()    {}
func (*GetPlansParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{5}
}

func (m *GetPlansParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlansParams.Unmarshal(m, b)
}
func (m *GetPlansParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlansParams.Marshal(b, m, deterministic)
}
func (m *GetPlansParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlansParams.Merge(m, src)
}
func (m *GetPlansParams) XXX_Size() int {
	return xxx_messageInfo_GetPlansParams.Size(m)
}
func (m *GetPlansParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlansParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlansParams proto.InternalMessageInfo

func (m *GetPlansParams) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type GetPlansResponse struct {
	Plans                []*Plan  `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlansResponse) Reset()         { *m = GetPlansResponse{} }
func (m *GetPlansResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlansResponse) ProtoMessage()    {}
func (*GetPlansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{6}
}

func (m *GetPlansResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlansResponse.Unmarshal(m, b)
}
func (m *GetPlansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlansResponse.Marshal(b, m, deterministic)
}
func (m *GetPlansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlansResponse.Merge(m, src)
}
func (m *GetPlansResponse) XXX_Size() int {
	return xxx_messageInfo_GetPlansResponse.Size(m)
}
func (m *GetPlansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlansResponse proto.InternalMessageInfo

func (m *GetPlansResponse) GetPlans() []*Plan {
	if m != nil {
		return m.Plans
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePlanParams)(nil), "pepeunlimited.billing.CreatePlanParams")
	proto.RegisterType((*Plan)(nil), "pepeunlimited.billing.Plan")
	proto.RegisterType((*UpdatePlanParams)(nil), "pepeunlimited.billing.UpdatePlanParams")
	proto.RegisterType((*DeletePlanParams)(nil), "pepeunlimited.billing.DeletePlanParams")
	proto.RegisterType((*GetPlanParams)(nil), "pepeunlimited.billing.GetPlanParams")
	proto.RegisterType((*GetPlansParams)(nil), "pepeunlimited.billing.GetPlansParams")
	proto.RegisterType((*GetPlansResponse)(nil), "pepeunlimited.billing.GetPlansResponse")
}

func init() { proto.RegisterFile("plan.proto", fileDescriptor_2d655ab2f7683c23) }

var fileDescriptor_2d655ab2f7683c23 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0x49, 0x63, 0x37, 0x2f, 0xb4, 0x04, 0xb1, 0x32, 0xcf, 0x65, 0x23, 0x98, 0x8d, 0xfa,
	0xe4, 0x90, 0xee, 0xd2, 0x6b, 0xf7, 0x83, 0x61, 0xd8, 0xa1, 0x78, 0xec, 0xd2, 0x8b, 0x51, 0xac,
	0xb7, 0x20, 0x50, 0x64, 0x21, 0x2b, 0x1b, 0xfb, 0xbb, 0x76, 0xdd, 0x1f, 0x37, 0x24, 0xbb, 0x35,
	0x36, 0x73, 0xda, 0xe3, 0xf7, 0xe3, 0x7d, 0xca, 0x7b, 0x5f, 0x0c, 0xa0, 0x04, 0x95, 0xa9, 0xd2,
	0x95, 0xa9, 0xc8, 0x85, 0x42, 0x85, 0x07, 0x29, 0xf8, 0x9e, 0x1b, 0x64, 0xe9, 0x96, 0x0b, 0xc1,
	0xe5, 0x2e, 0xba, 0xdc, 0x55, 0xd5, 0x4e, 0xe0, 0xda, 0x99, 0xb6, 0x87, 0x1f, 0x6b, 0xdc, 0x2b,
	0xf3, 0xbb, 0x99, 0x89, 0xde, 0x0c, 0xc5, 0x5f, 0x9a, 0x2a, 0x85, 0xba, 0x6e, 0xf4, 0x98, 0xc0,
	0xf2, 0xa3, 0x46, 0x6a, 0xf0, 0x4e, 0x50, 0x79, 0x47, 0x35, 0xdd, 0xd7, 0xf1, 0x5f, 0x0f, 0x4e,
	0x2c, 0x24, 0xe7, 0x30, 0xe1, 0x2c, 0xf4, 0x56, 0x5e, 0x32, 0xcd, 0x27, 0x9c, 0x91, 0x57, 0x70,
	0xaa, 0x34, 0x2f, 0xb1, 0xe0, 0x2c, 0x9c, 0x38, 0x36, 0x70, 0x38, 0x63, 0x24, 0x86, 0x33, 0xc3,
	0x8d, 0xc0, 0x82, 0x6f, 0x6e, 0xa4, 0xd5, 0xa7, 0x4e, 0x5f, 0x38, 0x32, 0xdb, 0xdc, 0xc8, 0x8c,
	0x91, 0xd7, 0x00, 0x4a, 0x57, 0xec, 0x50, 0x1a, 0x6b, 0x38, 0x71, 0x86, 0x79, 0xcb, 0x64, 0x8c,
	0xbc, 0x80, 0x99, 0x73, 0x87, 0xb3, 0x95, 0x97, 0xcc, 0xf3, 0x06, 0xd8, 0x37, 0x6b, 0x43, 0xb5,
	0x29, 0xa8, 0x09, 0x7d, 0x27, 0x04, 0x0e, 0xdf, 0x1a, 0x72, 0x01, 0x3e, 0x4a, 0x66, 0x85, 0xa0,
	0x99, 0x40, 0xc9, 0x6e, 0x8d, 0x5d, 0xe9, 0xbb, 0x62, 0xfd, 0x95, 0x08, 0x2c, 0x3f, 0xa1, 0xc0,
	0x1e, 0x97, 0xc0, 0xd9, 0x17, 0x34, 0x1d, 0x41, 0x5e, 0x42, 0x60, 0xaf, 0x5d, 0x3c, 0xee, 0xec,
	0x5b, 0x98, 0xb1, 0x78, 0x0d, 0xe7, 0xad, 0xb3, 0x6e, 0xad, 0xfd, 0x55, 0xbc, 0xc1, 0x2a, 0xf1,
	0x67, 0x58, 0x3e, 0x0c, 0xe4, 0x58, 0xab, 0x4a, 0xd6, 0x48, 0x36, 0x30, 0xb3, 0x71, 0x75, 0xe8,
	0xad, 0xa6, 0xc9, 0xe2, 0xfa, 0x32, 0xfd, 0x6f, 0x9b, 0xa9, 0x1d, 0xca, 0x1b, 0xe7, 0xf5, 0x9f,
	0x29, 0x2c, 0x2c, 0xfe, 0x86, 0xfa, 0x27, 0x2f, 0x91, 0xe4, 0x00, 0x5d, 0x59, 0xe4, 0x6a, 0x24,
	0x61, 0xd8, 0x67, 0x74, 0xec, 0x29, 0x9b, 0xd9, 0x5d, 0x6b, 0x34, 0x73, 0x78, 0xd0, 0x27, 0x33,
	0xbb, 0x6b, 0x8f, 0x66, 0x0e, 0x0b, 0x39, 0x9e, 0x79, 0x0f, 0xa7, 0x0f, 0x27, 0x25, 0xef, 0x46,
	0x8c, 0xfd, 0x92, 0xa2, 0xab, 0x27, 0x6c, 0x8f, 0xd5, 0x7c, 0x85, 0xa0, 0xe5, 0xc8, 0xdb, 0xe3,
	0x33, 0xcf, 0xf8, 0xa5, 0x1f, 0xe6, 0xf7, 0xee, 0x6f, 0xa4, 0x55, 0xb9, 0xf5, 0xdd, 0x47, 0xf6,
	0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb0, 0x60, 0x7b, 0xc6, 0x03, 0x00, 0x00,
}
