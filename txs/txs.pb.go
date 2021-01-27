// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: txs.proto

package txs

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
	crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_txs_payload "github.com/hyperledger/burrow/txs/payload"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Envelope_EncodingType int32

const (
	Envelope_JSON Envelope_EncodingType = 0
	Envelope_RLP  Envelope_EncodingType = 1
)

var Envelope_EncodingType_name = map[int32]string{
	0: "JSON",
	1: "RLP",
}

var Envelope_EncodingType_value = map[string]int32{
	"JSON": 0,
	"RLP":  1,
}

func (x Envelope_EncodingType) String() string {
	return proto.EnumName(Envelope_EncodingType_name, int32(x))
}

func (Envelope_EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{0, 0}
}

// An envelope contains both the signable Tx and the signatures for each input (in signatories)
type Envelope struct {
	Signatories []Signatory `protobuf:"bytes,1,rep,name=Signatories,proto3" json:"Signatories"`
	// Canonical bytes of the Tx ready to be signed
	Tx                   *Tx                   `protobuf:"bytes,2,opt,name=Tx,proto3,customtype=Tx" json:"Tx,omitempty"`
	Encoding             Envelope_EncodingType `protobuf:"varint,3,opt,name=Encoding,proto3,enum=txs.Envelope_EncodingType" json:"Encoding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Envelope) Reset()      { *m = Envelope{} }
func (*Envelope) ProtoMessage() {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{0}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetSignatories() []Signatory {
	if m != nil {
		return m.Signatories
	}
	return nil
}

func (m *Envelope) GetEncoding() Envelope_EncodingType {
	if m != nil {
		return m.Encoding
	}
	return Envelope_JSON
}

func (*Envelope) XXX_MessageName() string {
	return "txs.Envelope"
}

// Signatory contains signature and one or both of Address and PublicKey to identify the signer
type Signatory struct {
	Address              *github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address,omitempty"`
	PublicKey            *crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Signature            *crypto.Signature                             `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *Signatory) Reset()         { *m = Signatory{} }
func (m *Signatory) String() string { return proto.CompactTextString(m) }
func (*Signatory) ProtoMessage()    {}
func (*Signatory) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{1}
}
func (m *Signatory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signatory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Signatory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signatory.Merge(m, src)
}
func (m *Signatory) XXX_Size() int {
	return m.Size()
}
func (m *Signatory) XXX_DiscardUnknown() {
	xxx_messageInfo_Signatory.DiscardUnknown(m)
}

var xxx_messageInfo_Signatory proto.InternalMessageInfo

func (m *Signatory) GetPublicKey() *crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signatory) GetSignature() *crypto.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (*Signatory) XXX_MessageName() string {
	return "txs.Signatory"
}

// BroadcastTx or Transaction receipt
type Receipt struct {
	// Transaction type
	TxType github_com_hyperledger_burrow_txs_payload.Type `protobuf:"varint,1,opt,name=TxType,proto3,casttype=github.com/hyperledger/burrow/txs/payload.Type" json:"TxType,omitempty"`
	// The hash of the transaction that caused this event to be generated
	TxHash github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,2,opt,name=TxHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"TxHash"`
	// Whether the transaction creates a contract
	CreatesContract bool `protobuf:"varint,3,opt,name=CreatesContract,proto3" json:"CreatesContract,omitempty"`
	// The address of the contract being called
	ContractAddress      github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,4,opt,name=ContractAddress,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"ContractAddress"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{2}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetTxType() github_com_hyperledger_burrow_txs_payload.Type {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *Receipt) GetCreatesContract() bool {
	if m != nil {
		return m.CreatesContract
	}
	return false
}

func (*Receipt) XXX_MessageName() string {
	return "txs.Receipt"
}
func init() {
	proto.RegisterEnum("txs.Envelope_EncodingType", Envelope_EncodingType_name, Envelope_EncodingType_value)
	golang_proto.RegisterEnum("txs.Envelope_EncodingType", Envelope_EncodingType_name, Envelope_EncodingType_value)
	proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	golang_proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	golang_proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	proto.RegisterType((*Receipt)(nil), "txs.Receipt")
	golang_proto.RegisterType((*Receipt)(nil), "txs.Receipt")
}

func init() { proto.RegisterFile("txs.proto", fileDescriptor_372ebcf753025bdc) }
func init() { golang_proto.RegisterFile("txs.proto", fileDescriptor_372ebcf753025bdc) }

var fileDescriptor_372ebcf753025bdc = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x6f, 0xd3, 0x40,
	0x1c, 0xc6, 0x73, 0x8e, 0x95, 0x97, 0x4b, 0x68, 0xcb, 0x09, 0xa1, 0x28, 0x83, 0x9d, 0x66, 0xca,
	0x00, 0x36, 0x0a, 0xd0, 0x01, 0x26, 0x5c, 0x21, 0x55, 0xe1, 0xad, 0xba, 0x7a, 0x62, 0x40, 0xf2,
	0xcb, 0x5f, 0x8e, 0xa5, 0xe0, 0xb3, 0xce, 0x17, 0xb0, 0xbf, 0x09, 0x23, 0x9f, 0x80, 0x1d, 0xb1,
	0x30, 0x66, 0x64, 0x44, 0x19, 0x2c, 0x94, 0x7e, 0x0b, 0x26, 0x74, 0xee, 0xd9, 0x2d, 0x1d, 0x8a,
	0xd8, 0xee, 0xee, 0x79, 0xee, 0xe7, 0xe7, 0xfe, 0x8f, 0x71, 0x5f, 0xe4, 0x99, 0x95, 0x72, 0x26,
	0x18, 0x69, 0x8b, 0x3c, 0x1b, 0xdf, 0x89, 0x58, 0xc4, 0xaa, 0xbd, 0x2d, 0x57, 0x17, 0xd2, 0x78,
	0x18, 0xf0, 0x22, 0x15, 0x6a, 0x37, 0xfd, 0x86, 0x70, 0xef, 0x79, 0xf2, 0x01, 0x56, 0x2c, 0x05,
	0x72, 0x84, 0x07, 0x67, 0x71, 0x94, 0x78, 0x82, 0xf1, 0x18, 0xb2, 0x11, 0x9a, 0xb4, 0x67, 0x83,
	0xf9, 0x9e, 0x25, 0xb1, 0xf5, 0x79, 0xe1, 0xe8, 0x9b, 0xd2, 0x6c, 0xd1, 0xab, 0x46, 0x72, 0x17,
	0x6b, 0x6e, 0x3e, 0xd2, 0x26, 0x68, 0x36, 0x74, 0x3a, 0xdb, 0xd2, 0xd4, 0xdc, 0x9c, 0x6a, 0x6e,
	0x4e, 0x8e, 0x24, 0x3b, 0x60, 0x61, 0x9c, 0x44, 0xa3, 0xf6, 0x04, 0xcd, 0xf6, 0xe6, 0xe3, 0x0a,
	0x56, 0x7f, 0xd0, 0xaa, 0x55, 0xb7, 0x48, 0x81, 0x36, 0xde, 0xe9, 0x21, 0x1e, 0x5e, 0x55, 0x48,
	0x0f, 0xeb, 0x8b, 0xb3, 0x37, 0xaf, 0x0f, 0x5a, 0xa4, 0x8b, 0xdb, 0xf4, 0xe5, 0xe9, 0x01, 0x7a,
	0xa2, 0x7f, 0xfa, 0x6c, 0xb6, 0xa6, 0x5f, 0x11, 0xee, 0x37, 0xc9, 0xc8, 0x02, 0x77, 0x9f, 0x85,
	0x21, 0x87, 0x4c, 0x46, 0x97, 0x59, 0x1e, 0x6c, 0x4b, 0xf3, 0x5e, 0x14, 0x8b, 0xe5, 0xda, 0xb7,
	0x02, 0xf6, 0xde, 0x5e, 0x16, 0x29, 0xf0, 0x15, 0x84, 0x11, 0x70, 0xdb, 0x5f, 0x73, 0xce, 0x3e,
	0xda, 0x6a, 0x18, 0xea, 0x1e, 0xad, 0x01, 0xc4, 0xc6, 0xfd, 0xd3, 0xb5, 0xbf, 0x8a, 0x83, 0x17,
	0x50, 0x54, 0x2f, 0x1b, 0xcc, 0x6f, 0x5b, 0xca, 0xdc, 0x08, 0xf4, 0xd2, 0x23, 0x2f, 0x5c, 0x24,
	0x59, 0x73, 0x18, 0xe9, 0x7f, 0x5f, 0x68, 0x04, 0x7a, 0xe9, 0x99, 0x7e, 0xd1, 0x70, 0x97, 0x42,
	0x00, 0x71, 0x2a, 0xc8, 0x02, 0x77, 0xdc, 0x5c, 0x3e, 0xb5, 0x0a, 0x7e, 0xcb, 0x99, 0xff, 0x2e,
	0x4d, 0xeb, 0xe6, 0xe0, 0x22, 0xcf, 0xec, 0xd4, 0x2b, 0x56, 0xcc, 0x0b, 0xad, 0x6a, 0x7c, 0x8a,
	0x40, 0x5e, 0x49, 0xd6, 0x89, 0x97, 0x2d, 0x55, 0x21, 0x8f, 0x65, 0x5f, 0xdb, 0xd2, 0xbc, 0x7f,
	0x33, 0xcf, 0x8f, 0x13, 0x8f, 0x17, 0xd6, 0x09, 0xe4, 0x4e, 0x21, 0x20, 0xa3, 0x0a, 0x42, 0x66,
	0x78, 0xff, 0x98, 0x83, 0x27, 0x20, 0x3b, 0x66, 0x89, 0xe0, 0x5e, 0x20, 0xaa, 0x2a, 0x7b, 0xf4,
	0xfa, 0x31, 0x79, 0x87, 0xf7, 0xeb, 0x75, 0x5d, 0x83, 0x5e, 0x25, 0x78, 0xa4, 0x12, 0xfc, 0x5f,
	0x15, 0xd7, 0x61, 0xce, 0xd3, 0xcd, 0xce, 0x40, 0x3f, 0x76, 0x06, 0xfa, 0xb9, 0x33, 0xd0, 0xaf,
	0x9d, 0x81, 0xbe, 0x9f, 0x1b, 0x68, 0x73, 0x6e, 0xa0, 0xb7, 0x87, 0xff, 0x1c, 0x95, 0xdf, 0xa9,
	0x7e, 0xf7, 0x87, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x9a, 0xbd, 0xcb, 0x24, 0x03, 0x00,
	0x00,
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Envelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Encoding != 0 {
		i = encodeVarintTxs(dAtA, i, uint64(m.Encoding))
		i--
		dAtA[i] = 0x18
	}
	if m.Tx != nil {
		{
			size := m.Tx.Size()
			i -= size
			if _, err := m.Tx.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTxs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signatories) > 0 {
		for iNdEx := len(m.Signatories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTxs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Signatory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signatory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Signatory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Signature != nil {
		{
			size, err := m.Signature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTxs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTxs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Address != nil {
		{
			size := m.Address.Size()
			i -= size
			if _, err := m.Address.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTxs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size := m.ContractAddress.Size()
		i -= size
		if _, err := m.ContractAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTxs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.CreatesContract {
		i--
		if m.CreatesContract {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TxHash.Size()
		i -= size
		if _, err := m.TxHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTxs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.TxType != 0 {
		i = encodeVarintTxs(dAtA, i, uint64(m.TxType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxs(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatories) > 0 {
		for _, e := range m.Signatories {
			l = e.Size()
			n += 1 + l + sovTxs(uint64(l))
		}
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.Encoding != 0 {
		n += 1 + sovTxs(uint64(m.Encoding))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Signatory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovTxs(uint64(m.TxType))
	}
	l = m.TxHash.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.CreatesContract {
		n += 2
	}
	l = m.ContractAddress.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTxs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxs(x uint64) (n int) {
	return sovTxs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatories = append(m.Signatories, Signatory{})
			if err := m.Signatories[len(m.Signatories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Tx
			m.Tx = &v
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			m.Encoding = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encoding |= Envelope_EncodingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Signatory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Signatory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signatory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.Address = &v
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &crypto.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &crypto.Signature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= github_com_hyperledger_burrow_txs_payload.Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatesContract", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CreatesContract = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTxs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTxs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxs = fmt.Errorf("proto: unexpected end of group")
)
