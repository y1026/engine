// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	common.proto
	service.proto
	stream.proto

It has these top-level messages:
	Empty
	Transaction
	TxData
	Params
	Envelope
	StreamMessage
	Block
	ConnectionEstablish
	PeerTable
	Peer
	ConsensusMessage
	View
	ElectionMessage
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Transaction_Status int32

const (
	Transaction_UNCONFIRMED Transaction_Status = 0
	Transaction_CONFIRMED   Transaction_Status = 1
	Transaction_UNKNOWN     Transaction_Status = 2
)

var Transaction_Status_name = map[int32]string{
	0: "UNCONFIRMED",
	1: "CONFIRMED",
	2: "UNKNOWN",
}
var Transaction_Status_value = map[string]int32{
	"UNCONFIRMED": 0,
	"CONFIRMED":   1,
	"UNKNOWN":     2,
}

func (x Transaction_Status) String() string {
	return proto.EnumName(Transaction_Status_name, int32(x))
}
func (Transaction_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type TxData_TxDataType int32

const (
	TxData_Invoke TxData_TxDataType = 0
	TxData_Query  TxData_TxDataType = 1
)

var TxData_TxDataType_name = map[int32]string{
	0: "Invoke",
	1: "Query",
}
var TxData_TxDataType_value = map[string]int32{
	"Invoke": 0,
	"Query":  1,
}

func (x TxData_TxDataType) String() string {
	return proto.EnumName(TxData_TxDataType_name, int32(x))
}
func (TxData_TxDataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Transaction struct {
	InvokePeerID      string             `protobuf:"bytes,1,opt,name=InvokePeerID" json:"InvokePeerID,omitempty"`
	TransactionID     string             `protobuf:"bytes,2,opt,name=TransactionID" json:"TransactionID,omitempty"`
	TransactionStatus Transaction_Status `protobuf:"varint,3,opt,name=TransactionStatus,enum=message.Transaction_Status" json:"TransactionStatus,omitempty"`
	TransactionHash   string             `protobuf:"bytes,4,opt,name=TransactionHash" json:"TransactionHash,omitempty"`
	TxData            *TxData            `protobuf:"bytes,5,opt,name=TxData" json:"TxData,omitempty"`
	TransactionType   int32              `protobuf:"varint,6,opt,name=TransactionType" json:"TransactionType,omitempty"`
	TimeStamp         int64              `protobuf:"varint,7,opt,name=TimeStamp" json:"TimeStamp,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Transaction) GetInvokePeerID() string {
	if m != nil {
		return m.InvokePeerID
	}
	return ""
}

func (m *Transaction) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *Transaction) GetTransactionStatus() Transaction_Status {
	if m != nil {
		return m.TransactionStatus
	}
	return Transaction_UNCONFIRMED
}

func (m *Transaction) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *Transaction) GetTxData() *TxData {
	if m != nil {
		return m.TxData
	}
	return nil
}

func (m *Transaction) GetTransactionType() int32 {
	if m != nil {
		return m.TransactionType
	}
	return 0
}

func (m *Transaction) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type TxData struct {
	Jsonrpc    string            `protobuf:"bytes,1,opt,name=Jsonrpc" json:"Jsonrpc,omitempty"`
	Method     TxData_TxDataType `protobuf:"varint,2,opt,name=Method,enum=message.TxData_TxDataType" json:"Method,omitempty"`
	Params     *Params           `protobuf:"bytes,3,opt,name=Params" json:"Params,omitempty"`
	ContractID string            `protobuf:"bytes,4,opt,name=ContractID" json:"ContractID,omitempty"`
}

func (m *TxData) Reset()                    { *m = TxData{} }
func (m *TxData) String() string            { return proto.CompactTextString(m) }
func (*TxData) ProtoMessage()               {}
func (*TxData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TxData) GetJsonrpc() string {
	if m != nil {
		return m.Jsonrpc
	}
	return ""
}

func (m *TxData) GetMethod() TxData_TxDataType {
	if m != nil {
		return m.Method
	}
	return TxData_Invoke
}

func (m *TxData) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TxData) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

type Params struct {
	ParamsType int32    `protobuf:"varint,1,opt,name=ParamsType" json:"ParamsType,omitempty"`
	Function   string   `protobuf:"bytes,2,opt,name=Function" json:"Function,omitempty"`
	Args       []string `protobuf:"bytes,3,rep,name=Args" json:"Args,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Params) GetParamsType() int32 {
	if m != nil {
		return m.ParamsType
	}
	return 0
}

func (m *Params) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Params) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*Transaction)(nil), "message.Transaction")
	proto.RegisterType((*TxData)(nil), "message.TxData")
	proto.RegisterType((*Params)(nil), "message.Params")
	proto.RegisterEnum("message.Transaction_Status", Transaction_Status_name, Transaction_Status_value)
	proto.RegisterEnum("message.TxData_TxDataType", TxData_TxDataType_name, TxData_TxDataType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xeb, 0xdd, 0x6e, 0x42, 0x26, 0x6d, 0x37, 0xcc, 0xc9, 0x2a, 0x08, 0x45, 0x01, 0x89,
	0x9c, 0x72, 0x08, 0xe2, 0x01, 0x50, 0xd3, 0x8a, 0x80, 0x9a, 0x16, 0x37, 0x15, 0x5c, 0x4d, 0xb0,
	0xda, 0x0a, 0x25, 0x8e, 0x1c, 0x2f, 0x62, 0xdf, 0x8f, 0x47, 0xe2, 0x01, 0xd0, 0x3a, 0x2e, 0xf1,
	0xa6, 0xa7, 0xd8, 0x9f, 0xc7, 0xe3, 0x3f, 0x9f, 0x06, 0x8e, 0x1a, 0xd9, 0xb6, 0xb2, 0xcb, 0x7a,
	0x25, 0xb5, 0x44, 0xbf, 0x15, 0xc3, 0xc0, 0xef, 0x44, 0xe2, 0xc3, 0xea, 0xbc, 0xed, 0xf5, 0x36,
	0xf9, 0xbb, 0x80, 0xb0, 0x56, 0xbc, 0x1b, 0x78, 0xa3, 0x1f, 0x64, 0x87, 0x09, 0x1c, 0x95, 0xdd,
	0x2f, 0xf9, 0x53, 0x5c, 0x0b, 0xa1, 0xca, 0x82, 0x92, 0x98, 0xa4, 0x01, 0xdb, 0x63, 0xf8, 0x06,
	0x8e, 0x9d, 0x2b, 0x65, 0x41, 0x17, 0xa6, 0x68, 0x1f, 0x62, 0x09, 0xcf, 0x1d, 0x70, 0xa3, 0xb9,
	0xde, 0x0c, 0x74, 0x19, 0x93, 0xf4, 0x24, 0x7f, 0x91, 0xd9, 0x1c, 0x99, 0x53, 0x91, 0x8d, 0x25,
	0xec, 0xe9, 0x2d, 0x4c, 0x61, 0xed, 0xc0, 0x8f, 0x7c, 0xb8, 0xa7, 0x87, 0xe6, 0xc9, 0x39, 0xc6,
	0xb7, 0xe0, 0xd5, 0xbf, 0x0b, 0xae, 0x39, 0x5d, 0xc5, 0x24, 0x0d, 0xf3, 0xf5, 0xf4, 0x92, 0xc1,
	0xcc, 0x1e, 0xcf, 0x5a, 0xd6, 0xdb, 0x5e, 0x50, 0x2f, 0x26, 0xe9, 0x8a, 0xcd, 0x31, 0xbe, 0x84,
	0xa0, 0x7e, 0x68, 0xc5, 0x8d, 0xe6, 0x6d, 0x4f, 0xfd, 0x98, 0xa4, 0x4b, 0x36, 0x81, 0xe4, 0x3d,
	0x78, 0x36, 0xe4, 0x1a, 0xc2, 0xdb, 0xea, 0xec, 0xaa, 0xba, 0x28, 0xd9, 0xe5, 0x79, 0x11, 0x1d,
	0xe0, 0x31, 0x04, 0xd3, 0x96, 0x60, 0x08, 0xfe, 0x6d, 0xf5, 0xb9, 0xba, 0xfa, 0x5a, 0x45, 0x8b,
	0xe4, 0x0f, 0x79, 0x0c, 0x8a, 0x14, 0xfc, 0x4f, 0x83, 0xec, 0x54, 0xdf, 0x58, 0xd9, 0x8f, 0x5b,
	0xcc, 0xc1, 0xbb, 0x14, 0xfa, 0x5e, 0xfe, 0x30, 0x82, 0x4f, 0xf2, 0xd3, 0xd9, 0xcf, 0xd8, 0xcf,
	0x2e, 0x25, 0xb3, 0x95, 0x3b, 0x01, 0xd7, 0x5c, 0xf1, 0x76, 0x54, 0xed, 0x0a, 0x18, 0x31, 0xb3,
	0xc7, 0xf8, 0x0a, 0xe0, 0x4c, 0x76, 0x5a, 0xf1, 0x46, 0x97, 0x85, 0xd5, 0xe9, 0x90, 0xe4, 0x35,
	0xc0, 0xd4, 0x1e, 0x01, 0xbc, 0x71, 0x04, 0xa2, 0x03, 0x0c, 0x60, 0xf5, 0x65, 0x23, 0xd4, 0x36,
	0x22, 0xc9, 0x37, 0x70, 0xda, 0x8d, 0x2b, 0xa3, 0x92, 0x18, 0x95, 0x0e, 0xc1, 0x53, 0x78, 0x76,
	0xb1, 0xe9, 0x8c, 0x55, 0x3b, 0x2e, 0xff, 0xf7, 0x88, 0x70, 0xf8, 0x41, 0xdd, 0xed, 0x12, 0x2f,
	0xd3, 0x80, 0x99, 0xf5, 0x77, 0xcf, 0x0c, 0xec, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x28, 0x7c, 0xc7, 0xc0, 0x02, 0x00, 0x00,
}