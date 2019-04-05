// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signedMessage.proto

package oipProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Currently supported message types
type MessageTypes int32

const (
	// Invalid
	MessageTypes_InvalidMessage MessageTypes = 0
	// OIP 0.5 record message
	MessageTypes_OIP05 MessageTypes = 1
	// Historian record
	MessageTypes_Historian MessageTypes = 2
	// Multipart
	MessageTypes_Multipart MessageTypes = 3
)

var MessageTypes_name = map[int32]string{
	0: "InvalidMessage",
	1: "OIP05",
	2: "Historian",
	3: "Multipart",
}
var MessageTypes_value = map[string]int32{
	"InvalidMessage": 0,
	"OIP05":          1,
	"Historian":      2,
	"Multipart":      3,
}

func (x MessageTypes) String() string {
	return proto.EnumName(MessageTypes_name, int32(x))
}
func (MessageTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// Currently supported signature verification means
type SignatureTypes int32

const (
	// Invalid
	SignatureTypes_InvalidSignature SignatureTypes = 0
	// FLO address message signing
	SignatureTypes_Flo SignatureTypes = 1
	// Bitcoin address message signing
	// https://tools.bitcoin.com/verify-message/
	SignatureTypes_Btc SignatureTypes = 2
	// Raven address message signing - ToDo
	SignatureTypes_Rvn SignatureTypes = 3
	// Flo transaction is the signature - ToDo
	SignatureTypes_Tx SignatureTypes = 4
)

var SignatureTypes_name = map[int32]string{
	0: "InvalidSignature",
	1: "Flo",
	2: "Btc",
	3: "Rvn",
	4: "Tx",
}
var SignatureTypes_value = map[string]int32{
	"InvalidSignature": 0,
	"Flo":              1,
	"Btc":              2,
	"Rvn":              3,
	"Tx":               4,
}

func (x SignatureTypes) String() string {
	return proto.EnumName(SignatureTypes_name, int32(x))
}
func (SignatureTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type SignedMessage struct {
	// Raw Data that was signed by this message
	SerializedMessage []byte `protobuf:"bytes,1,opt,name=SerializedMessage,proto3" json:"SerializedMessage,omitempty"`
	// Specifies the type of contained data for further deserialization
	MessageType MessageTypes `protobuf:"varint,2,opt,name=MessageType,enum=oipProto.MessageTypes" json:"MessageType,omitempty"`
	// Identifies signature type for verification
	SignatureType SignatureTypes `protobuf:"varint,3,opt,name=SignatureType,enum=oipProto.SignatureTypes" json:"SignatureType,omitempty"`
	// Public Key used in the signing of original message
	PubKey []byte `protobuf:"bytes,4,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	// Signature of signed message
	Signature []byte `protobuf:"bytes,5,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (m *SignedMessage) Reset()                    { *m = SignedMessage{} }
func (m *SignedMessage) String() string            { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()               {}
func (*SignedMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SignedMessage) GetSerializedMessage() []byte {
	if m != nil {
		return m.SerializedMessage
	}
	return nil
}

func (m *SignedMessage) GetMessageType() MessageTypes {
	if m != nil {
		return m.MessageType
	}
	return MessageTypes_InvalidMessage
}

func (m *SignedMessage) GetSignatureType() SignatureTypes {
	if m != nil {
		return m.SignatureType
	}
	return SignatureTypes_InvalidSignature
}

func (m *SignedMessage) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedMessage)(nil), "oipProto.SignedMessage")
	proto.RegisterEnum("oipProto.MessageTypes", MessageTypes_name, MessageTypes_value)
	proto.RegisterEnum("oipProto.SignatureTypes", SignatureTypes_name, SignatureTypes_value)
}

func init() { proto.RegisterFile("signedMessage.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x49, 0x1b, 0xcd, 0xd8, 0x86, 0x71, 0x94, 0x92, 0x83, 0x87, 0xe2, 0xa9, 0x14,
	0x09, 0xa2, 0x08, 0x9e, 0x3c, 0xf4, 0x20, 0x86, 0x52, 0x0c, 0x69, 0x4f, 0xde, 0xb6, 0xba, 0x94,
	0x85, 0x90, 0x84, 0xdd, 0x4d, 0xb1, 0xfe, 0x67, 0xff, 0x83, 0x24, 0x4d, 0xd3, 0x04, 0x6f, 0x6f,
	0xe6, 0x7b, 0x33, 0xfb, 0x96, 0x81, 0x2b, 0x2d, 0xb7, 0xa9, 0xf8, 0x5a, 0x0a, 0xad, 0xf9, 0x56,
	0x04, 0xb9, 0xca, 0x4c, 0x46, 0xe7, 0x99, 0xcc, 0xa3, 0x52, 0xdd, 0xfe, 0x32, 0x18, 0xad, 0xda,
	0x0e, 0xba, 0x83, 0xcb, 0x95, 0x50, 0x92, 0x27, 0xf2, 0xa7, 0x69, 0xfa, 0x6c, 0xc2, 0xa6, 0xc3,
	0xf8, 0x3f, 0xa0, 0x67, 0xb8, 0xa8, 0xe5, 0x7a, 0x9f, 0x0b, 0xdf, 0x9a, 0xb0, 0xa9, 0xf7, 0x30,
	0x0e, 0x8e, 0xfb, 0x83, 0x16, 0xd4, 0x71, 0xdb, 0x4a, 0x2f, 0x87, 0x87, 0xb9, 0x29, 0xd4, 0x61,
	0xd6, 0xae, 0x66, 0xfd, 0xd3, 0x6c, 0x07, 0xeb, 0xb8, 0x6b, 0xa7, 0x31, 0x38, 0x51, 0xb1, 0x59,
	0x88, 0xbd, 0xdf, 0xaf, 0xc2, 0xd5, 0x15, 0xdd, 0x80, 0xdb, 0x18, 0xfd, 0x41, 0x85, 0x4e, 0x8d,
	0xd9, 0x02, 0x86, 0xed, 0x48, 0x44, 0xe0, 0x85, 0xe9, 0x8e, 0x27, 0xf2, 0xf8, 0x23, 0xec, 0x91,
	0x0b, 0x83, 0xf7, 0x30, 0xba, 0x7f, 0x42, 0x46, 0x23, 0x70, 0xdf, 0xa4, 0x36, 0x99, 0x92, 0x3c,
	0x45, 0xab, 0x2c, 0x97, 0x45, 0x62, 0x64, 0xce, 0x95, 0x41, 0x7b, 0x16, 0x82, 0xd7, 0xcd, 0x48,
	0xd7, 0x80, 0xf5, 0xba, 0x06, 0x60, 0x8f, 0xce, 0xc0, 0x7e, 0x4d, 0x32, 0x64, 0xa5, 0x98, 0x9b,
	0x4f, 0xb4, 0x4a, 0x11, 0xef, 0x52, 0xb4, 0xc9, 0x01, 0x6b, 0xfd, 0x8d, 0xfd, 0x39, 0x7c, 0x34,
	0x37, 0xd9, 0x38, 0xd5, 0x91, 0x1e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44, 0xa0, 0x9b, 0x43,
	0xbb, 0x01, 0x00, 0x00,
}
