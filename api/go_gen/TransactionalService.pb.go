// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.2
// source: TransactionalService.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId     string                 `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	InitSystem        string                 `protobuf:"bytes,2,opt,name=initSystem,proto3" json:"initSystem,omitempty"`
	ParticipantSystem string                 `protobuf:"bytes,3,opt,name=participantSystem,proto3" json:"participantSystem,omitempty"`
	TransactionStart  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=transactionStart,proto3" json:"transactionStart,omitempty"`
	Success           bool                   `protobuf:"varint,5,opt,name=success,proto3" json:"success,omitempty"`
	BizType           string                 `protobuf:"bytes,6,opt,name=bizType,proto3" json:"bizType,omitempty"`
	BizId             string                 `protobuf:"bytes,7,opt,name=bizId,proto3" json:"bizId,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TransactionalService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_TransactionalService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_TransactionalService_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *TransactionRequest) GetInitSystem() string {
	if x != nil {
		return x.InitSystem
	}
	return ""
}

func (x *TransactionRequest) GetParticipantSystem() string {
	if x != nil {
		return x.ParticipantSystem
	}
	return ""
}

func (x *TransactionRequest) GetTransactionStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionStart
	}
	return nil
}

func (x *TransactionRequest) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TransactionRequest) GetBizType() string {
	if x != nil {
		return x.BizType
	}
	return ""
}

func (x *TransactionRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionStatus string                 `protobuf:"bytes,1,opt,name=transactionStatus,proto3" json:"transactionStatus,omitempty"`
	TransactionStart  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=transactionStart,proto3" json:"transactionStart,omitempty"`
	TransactionId     string                 `protobuf:"bytes,3,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	InitSystem        string                 `protobuf:"bytes,4,opt,name=initSystem,proto3" json:"initSystem,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TransactionalService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_TransactionalService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_TransactionalService_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionResponse) GetTransactionStatus() string {
	if x != nil {
		return x.TransactionStatus
	}
	return ""
}

func (x *TransactionResponse) GetTransactionStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionStart
	}
	return nil
}

func (x *TransactionResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *TransactionResponse) GetInitSystem() string {
	if x != nil {
		return x.InitSystem
	}
	return ""
}

var File_TransactionalService_proto protoreflect.FileDescriptor

var file_TransactionalService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x12, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x46, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x7a, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0xd1, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e,
	0x69, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x69, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x32, 0xec, 0x04, 0x0a, 0x14, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x73, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12,
	0x2e, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x79, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x2e, 0x69, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x76, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6b, 0x6b, 0x69, 0x65, 0x2d,
	0x69, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x62, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TransactionalService_proto_rawDescOnce sync.Once
	file_TransactionalService_proto_rawDescData = file_TransactionalService_proto_rawDesc
)

func file_TransactionalService_proto_rawDescGZIP() []byte {
	file_TransactionalService_proto_rawDescOnce.Do(func() {
		file_TransactionalService_proto_rawDescData = protoimpl.X.CompressGZIP(file_TransactionalService_proto_rawDescData)
	})
	return file_TransactionalService_proto_rawDescData
}

var file_TransactionalService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TransactionalService_proto_goTypes = []interface{}{
	(*TransactionRequest)(nil),    // 0: id.secondbaser.service.api.TransactionRequest
	(*TransactionResponse)(nil),   // 1: id.secondbaser.service.api.TransactionResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_TransactionalService_proto_depIdxs = []int32{
	2, // 0: id.secondbaser.service.api.TransactionRequest.transactionStart:type_name -> google.protobuf.Timestamp
	2, // 1: id.secondbaser.service.api.TransactionResponse.transactionStart:type_name -> google.protobuf.Timestamp
	0, // 2: id.secondbaser.service.api.TransactionalRequest.StartTransaction:input_type -> id.secondbaser.service.api.TransactionRequest
	0, // 3: id.secondbaser.service.api.TransactionalRequest.RegisterParticipant:input_type -> id.secondbaser.service.api.TransactionRequest
	0, // 4: id.secondbaser.service.api.TransactionalRequest.QueryTransactionStatus:input_type -> id.secondbaser.service.api.TransactionRequest
	0, // 5: id.secondbaser.service.api.TransactionalRequest.CommitTransaction:input_type -> id.secondbaser.service.api.TransactionRequest
	0, // 6: id.secondbaser.service.api.TransactionalRequest.RollbackTransaction:input_type -> id.secondbaser.service.api.TransactionRequest
	1, // 7: id.secondbaser.service.api.TransactionalRequest.StartTransaction:output_type -> id.secondbaser.service.api.TransactionResponse
	1, // 8: id.secondbaser.service.api.TransactionalRequest.RegisterParticipant:output_type -> id.secondbaser.service.api.TransactionResponse
	1, // 9: id.secondbaser.service.api.TransactionalRequest.QueryTransactionStatus:output_type -> id.secondbaser.service.api.TransactionResponse
	1, // 10: id.secondbaser.service.api.TransactionalRequest.CommitTransaction:output_type -> id.secondbaser.service.api.TransactionResponse
	1, // 11: id.secondbaser.service.api.TransactionalRequest.RollbackTransaction:output_type -> id.secondbaser.service.api.TransactionResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TransactionalService_proto_init() }
func file_TransactionalService_proto_init() {
	if File_TransactionalService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TransactionalService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_TransactionalService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_TransactionalService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TransactionalService_proto_goTypes,
		DependencyIndexes: file_TransactionalService_proto_depIdxs,
		MessageInfos:      file_TransactionalService_proto_msgTypes,
	}.Build()
	File_TransactionalService_proto = out.File
	file_TransactionalService_proto_rawDesc = nil
	file_TransactionalService_proto_goTypes = nil
	file_TransactionalService_proto_depIdxs = nil
}
