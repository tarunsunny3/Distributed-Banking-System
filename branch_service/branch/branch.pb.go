// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: branch.proto

package branch

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{0}
}

func (x *Branch) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Branch) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type BranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BranchRequest) Reset() {
	*x = BranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchRequest) ProtoMessage() {}

func (x *BranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchRequest.ProtoReflect.Descriptor instead.
func (*BranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{1}
}

func (x *BranchRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	WriteEventID int32   `protobuf:"varint,2,opt,name=writeEventID,proto3" json:"writeEventID,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WithdrawRequest) GetWriteEventID() int32 {
	if x != nil {
		return x.WriteEventID
	}
	return 0
}

type WithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewBalance float32 `protobuf:"fixed32,1,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"`
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{3}
}

func (x *WithdrawResponse) GetNewBalance() float32 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

type QueryBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LastWriteEventID int32 `protobuf:"varint,2,opt,name=lastWriteEventID,proto3" json:"lastWriteEventID,omitempty"`
}

func (x *QueryBalanceRequest) Reset() {
	*x = QueryBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBalanceRequest) ProtoMessage() {}

func (x *QueryBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBalanceRequest.ProtoReflect.Descriptor instead.
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{4}
}

func (x *QueryBalanceRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *QueryBalanceRequest) GetLastWriteEventID() int32 {
	if x != nil {
		return x.LastWriteEventID
	}
	return 0
}

type QueryBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *QueryBalanceResponse) Reset() {
	*x = QueryBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBalanceResponse) ProtoMessage() {}

func (x *QueryBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBalanceResponse.ProtoReflect.Descriptor instead.
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{5}
}

func (x *QueryBalanceResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	WriteEventID int32   `protobuf:"varint,2,opt,name=writeEventID,proto3" json:"writeEventID,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{6}
}

func (x *DepositRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DepositRequest) GetWriteEventID() int32 {
	if x != nil {
		return x.WriteEventID
	}
	return 0
}

type DepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewBalance float32 `protobuf:"fixed32,1,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"`
}

func (x *DepositResponse) Reset() {
	*x = DepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositResponse) ProtoMessage() {}

func (x *DepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositResponse.ProtoReflect.Descriptor instead.
func (*DepositResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{7}
}

func (x *DepositResponse) GetNewBalance() float32 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

type PropagateWithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance      float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
	WriteEventID int32   `protobuf:"varint,2,opt,name=writeEventID,proto3" json:"writeEventID,omitempty"`
}

func (x *PropagateWithdrawRequest) Reset() {
	*x = PropagateWithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagateWithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagateWithdrawRequest) ProtoMessage() {}

func (x *PropagateWithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagateWithdrawRequest.ProtoReflect.Descriptor instead.
func (*PropagateWithdrawRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{8}
}

func (x *PropagateWithdrawRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *PropagateWithdrawRequest) GetWriteEventID() int32 {
	if x != nil {
		return x.WriteEventID
	}
	return 0
}

type PropagateWithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PropagateWithdrawResponse) Reset() {
	*x = PropagateWithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagateWithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagateWithdrawResponse) ProtoMessage() {}

func (x *PropagateWithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagateWithdrawResponse.ProtoReflect.Descriptor instead.
func (*PropagateWithdrawResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{9}
}

func (x *PropagateWithdrawResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PropagateDepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance      float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
	WriteEventID int32   `protobuf:"varint,2,opt,name=writeEventID,proto3" json:"writeEventID,omitempty"`
}

func (x *PropagateDepositRequest) Reset() {
	*x = PropagateDepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagateDepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagateDepositRequest) ProtoMessage() {}

func (x *PropagateDepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagateDepositRequest.ProtoReflect.Descriptor instead.
func (*PropagateDepositRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{10}
}

func (x *PropagateDepositRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *PropagateDepositRequest) GetWriteEventID() int32 {
	if x != nil {
		return x.WriteEventID
	}
	return 0
}

type PropagateDepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PropagateDepositResponse) Reset() {
	*x = PropagateDepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagateDepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagateDepositResponse) ProtoMessage() {}

func (x *PropagateDepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagateDepositResponse.ProtoReflect.Descriptor instead.
func (*PropagateDepositResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{11}
}

func (x *PropagateDepositResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_branch_proto protoreflect.FileDescriptor

var file_branch_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x32, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x62, 0x0a,
	0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x30, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x32, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61,
	0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x35, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x34, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xf2, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x62, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branch_proto_rawDescOnce sync.Once
	file_branch_proto_rawDescData = file_branch_proto_rawDesc
)

func file_branch_proto_rawDescGZIP() []byte {
	file_branch_proto_rawDescOnce.Do(func() {
		file_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_branch_proto_rawDescData)
	})
	return file_branch_proto_rawDescData
}

var file_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_branch_proto_goTypes = []interface{}{
	(*Branch)(nil),                    // 0: main.Branch
	(*BranchRequest)(nil),             // 1: main.BranchRequest
	(*WithdrawRequest)(nil),           // 2: main.WithdrawRequest
	(*WithdrawResponse)(nil),          // 3: main.WithdrawResponse
	(*QueryBalanceRequest)(nil),       // 4: main.QueryBalanceRequest
	(*QueryBalanceResponse)(nil),      // 5: main.QueryBalanceResponse
	(*DepositRequest)(nil),            // 6: main.DepositRequest
	(*DepositResponse)(nil),           // 7: main.DepositResponse
	(*PropagateWithdrawRequest)(nil),  // 8: main.PropagateWithdrawRequest
	(*PropagateWithdrawResponse)(nil), // 9: main.PropagateWithdrawResponse
	(*PropagateDepositRequest)(nil),   // 10: main.PropagateDepositRequest
	(*PropagateDepositResponse)(nil),  // 11: main.PropagateDepositResponse
}
var file_branch_proto_depIdxs = []int32{
	2,  // 0: main.BranchService.Withdraw:input_type -> main.WithdrawRequest
	4,  // 1: main.BranchService.QueryBalance:input_type -> main.QueryBalanceRequest
	6,  // 2: main.BranchService.Deposit:input_type -> main.DepositRequest
	8,  // 3: main.BranchService.PropagateWithdraw:input_type -> main.PropagateWithdrawRequest
	10, // 4: main.BranchService.PropagateDeposit:input_type -> main.PropagateDepositRequest
	3,  // 5: main.BranchService.Withdraw:output_type -> main.WithdrawResponse
	5,  // 6: main.BranchService.QueryBalance:output_type -> main.QueryBalanceResponse
	7,  // 7: main.BranchService.Deposit:output_type -> main.DepositResponse
	9,  // 8: main.BranchService.PropagateWithdraw:output_type -> main.PropagateWithdrawResponse
	11, // 9: main.BranchService.PropagateDeposit:output_type -> main.PropagateDepositResponse
	5,  // [5:10] is the sub-list for method output_type
	0,  // [0:5] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_branch_proto_init() }
func file_branch_proto_init() {
	if File_branch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagateWithdrawRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagateWithdrawResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagateDepositRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_branch_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagateDepositResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branch_proto_goTypes,
		DependencyIndexes: file_branch_proto_depIdxs,
		MessageInfos:      file_branch_proto_msgTypes,
	}.Build()
	File_branch_proto = out.File
	file_branch_proto_rawDesc = nil
	file_branch_proto_goTypes = nil
	file_branch_proto_depIdxs = nil
}
