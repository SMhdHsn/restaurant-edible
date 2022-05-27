// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/edible/inventory/request.proto

package eipb

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

// The request schema for calling Recycle on inventory.
type InventoryRecycleRequest struct {
	RecycleFinished      bool     `protobuf:"varint,1,opt,name=recycle_finished,json=recycleFinished,proto3" json:"recycle_finished,omitempty"`
	RecycleExpired       bool     `protobuf:"varint,2,opt,name=recycle_expired,json=recycleExpired,proto3" json:"recycle_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InventoryRecycleRequest) Reset()         { *m = InventoryRecycleRequest{} }
func (m *InventoryRecycleRequest) String() string { return proto.CompactTextString(m) }
func (*InventoryRecycleRequest) ProtoMessage()    {}
func (*InventoryRecycleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5483dda930989e31, []int{0}
}

func (m *InventoryRecycleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryRecycleRequest.Unmarshal(m, b)
}
func (m *InventoryRecycleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryRecycleRequest.Marshal(b, m, deterministic)
}
func (m *InventoryRecycleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryRecycleRequest.Merge(m, src)
}
func (m *InventoryRecycleRequest) XXX_Size() int {
	return xxx_messageInfo_InventoryRecycleRequest.Size(m)
}
func (m *InventoryRecycleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryRecycleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryRecycleRequest proto.InternalMessageInfo

func (m *InventoryRecycleRequest) GetRecycleFinished() bool {
	if m != nil {
		return m.RecycleFinished
	}
	return false
}

func (m *InventoryRecycleRequest) GetRecycleExpired() bool {
	if m != nil {
		return m.RecycleExpired
	}
	return false
}

// The request schema for calling Use on inventory.
type InventoryUseRequest struct {
	FoodId               uint32   `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InventoryUseRequest) Reset()         { *m = InventoryUseRequest{} }
func (m *InventoryUseRequest) String() string { return proto.CompactTextString(m) }
func (*InventoryUseRequest) ProtoMessage()    {}
func (*InventoryUseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5483dda930989e31, []int{1}
}

func (m *InventoryUseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryUseRequest.Unmarshal(m, b)
}
func (m *InventoryUseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryUseRequest.Marshal(b, m, deterministic)
}
func (m *InventoryUseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryUseRequest.Merge(m, src)
}
func (m *InventoryUseRequest) XXX_Size() int {
	return xxx_messageInfo_InventoryUseRequest.Size(m)
}
func (m *InventoryUseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryUseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryUseRequest proto.InternalMessageInfo

func (m *InventoryUseRequest) GetFoodId() uint32 {
	if m != nil {
		return m.FoodId
	}
	return 0
}

// The request schema for calling Buy on inventory.
type InventoryBuyRequest struct {
	Amount               uint32   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	ExpiresAt            uint32   `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InventoryBuyRequest) Reset()         { *m = InventoryBuyRequest{} }
func (m *InventoryBuyRequest) String() string { return proto.CompactTextString(m) }
func (*InventoryBuyRequest) ProtoMessage()    {}
func (*InventoryBuyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5483dda930989e31, []int{2}
}

func (m *InventoryBuyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryBuyRequest.Unmarshal(m, b)
}
func (m *InventoryBuyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryBuyRequest.Marshal(b, m, deterministic)
}
func (m *InventoryBuyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryBuyRequest.Merge(m, src)
}
func (m *InventoryBuyRequest) XXX_Size() int {
	return xxx_messageInfo_InventoryBuyRequest.Size(m)
}
func (m *InventoryBuyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryBuyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryBuyRequest proto.InternalMessageInfo

func (m *InventoryBuyRequest) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InventoryBuyRequest) GetExpiresAt() uint32 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*InventoryRecycleRequest)(nil), "edible.inventory.request.InventoryRecycleRequest")
	proto.RegisterType((*InventoryUseRequest)(nil), "edible.inventory.request.InventoryUseRequest")
	proto.RegisterType((*InventoryBuyRequest)(nil), "edible.inventory.request.InventoryBuyRequest")
}

func init() {
	proto.RegisterFile("protos/edible/inventory/request.proto", fileDescriptor_5483dda930989e31)
}

var fileDescriptor_5483dda930989e31 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0x47, 0xa9, 0x43, 0xd5, 0x0b, 0x55, 0x89, 0x60, 0xbb, 0x08, 0x52, 0x10, 0x75, 0x49, 0x07,
	0x3f, 0x81, 0x82, 0x42, 0xc1, 0x29, 0xe0, 0x1c, 0xda, 0xe6, 0x16, 0x03, 0x6d, 0x52, 0x93, 0x54,
	0xec, 0xb7, 0x7f, 0xbc, 0xfc, 0x29, 0x6f, 0xbc, 0xe7, 0xfc, 0xe0, 0x84, 0xc0, 0xe3, 0x62, 0xb4,
	0xd3, 0xb6, 0x41, 0x21, 0xfb, 0x09, 0x1b, 0xa9, 0xfe, 0x50, 0x39, 0x6d, 0xb6, 0xc6, 0xe0, 0xef,
	0x8a, 0xd6, 0x51, 0xef, 0x49, 0x15, 0x3c, 0xdd, 0x3d, 0x8d, 0xbe, 0x9e, 0xa1, 0x6c, 0x13, 0x64,
	0x38, 0x6c, 0xc3, 0x84, 0x2c, 0x28, 0xf2, 0x02, 0x37, 0x26, 0x10, 0x3e, 0x4a, 0x25, 0xed, 0x0f,
	0x8a, 0x2a, 0x7b, 0xc8, 0x9e, 0x2f, 0xd8, 0x75, 0xe4, 0x9f, 0x11, 0x93, 0x27, 0x48, 0x88, 0xe3,
	0xff, 0x22, 0x0d, 0x8a, 0xea, 0xcc, 0x2f, 0xaf, 0x22, 0xfe, 0x08, 0xb4, 0xa6, 0x70, 0xbb, 0xe7,
	0xbe, 0xed, 0x9e, 0x2a, 0xe1, 0x7c, 0xd4, 0x5a, 0x70, 0x19, 0x0a, 0x05, 0xcb, 0x8f, 0x67, 0x2b,
	0xea, 0xaf, 0x93, 0xfd, 0xfb, 0xba, 0xa5, 0xfd, 0x1d, 0xe4, 0xdd, 0xac, 0x57, 0xe5, 0xd2, 0x3c,
	0x5c, 0xe4, 0x1e, 0x20, 0xf4, 0x2d, 0xef, 0x9c, 0x7f, 0x42, 0xc1, 0x2e, 0x23, 0x79, 0x73, 0x7d,
	0xee, 0x7f, 0xe3, 0xf5, 0x10, 0x00, 0x00, 0xff, 0xff, 0x13, 0x1e, 0x6a, 0x7e, 0x36, 0x01, 0x00,
	0x00,
}