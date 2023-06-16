// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irismod/token/v1/token.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Token defines a standard for the fungible token
type Token struct {
	Symbol        string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Scale         uint32 `protobuf:"varint,3,opt,name=scale,proto3" json:"scale,omitempty"`
	MinUnit       string `protobuf:"bytes,4,opt,name=min_unit,json=minUnit,proto3" json:"min_unit,omitempty"`
	InitialSupply uint64 `protobuf:"varint,5,opt,name=initial_supply,json=initialSupply,proto3" json:"initial_supply,omitempty"`
	MaxSupply     uint64 `protobuf:"varint,6,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	Mintable      bool   `protobuf:"varint,7,opt,name=mintable,proto3" json:"mintable,omitempty"`
	Owner         string `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b3436d30fd508a, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

// Params defines token module's parameters
type Params struct {
	TokenTaxRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=token_tax_rate,json=tokenTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token_tax_rate"`
	IssueTokenBaseFee types.Coin                             `protobuf:"bytes,2,opt,name=issue_token_base_fee,json=issueTokenBaseFee,proto3" json:"issue_token_base_fee"`
	MintTokenFeeRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=mint_token_fee_ratio,json=mintTokenFeeRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mint_token_fee_ratio"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b3436d30fd508a, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Token)(nil), "irismod.token.v1.Token")
	proto.RegisterType((*Params)(nil), "irismod.token.v1.Params")
}

func init() { proto.RegisterFile("irismod/token/v1/token.proto", fileDescriptor_c5b3436d30fd508a) }

var fileDescriptor_c5b3436d30fd508a = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0x6e, 0x13, 0x41,
	0x10, 0xc7, 0x6f, 0xc3, 0xd9, 0x71, 0x16, 0x12, 0x91, 0x95, 0x85, 0x2e, 0x16, 0x9c, 0xad, 0x48,
	0x20, 0x37, 0xdc, 0xe9, 0xa0, 0x4b, 0x69, 0x50, 0x5a, 0xa2, 0xc3, 0x34, 0x34, 0xa7, 0xbd, 0xf3,
	0xc4, 0xac, 0x72, 0xbb, 0x6b, 0xdd, 0xee, 0x19, 0xfb, 0x0d, 0x28, 0xe9, 0x68, 0xf3, 0x38, 0x2e,
	0x53, 0x22, 0x8a, 0x08, 0xec, 0x86, 0xc7, 0x40, 0xfb, 0x61, 0x44, 0x4b, 0xb5, 0x33, 0xff, 0x9d,
	0xfd, 0xcd, 0xec, 0xcc, 0xe0, 0xa7, 0xac, 0x61, 0x8a, 0xcb, 0x59, 0xaa, 0xe5, 0x0d, 0x88, 0x74,
	0x99, 0x39, 0x23, 0x59, 0x34, 0x52, 0x4b, 0xf2, 0xd8, 0xdf, 0x26, 0x4e, 0x5c, 0x66, 0x83, 0xb8,
	0x92, 0x8a, 0x4b, 0x95, 0x96, 0x54, 0x41, 0xba, 0xcc, 0x4a, 0xd0, 0x34, 0x4b, 0x2b, 0xc9, 0xfc,
	0x8b, 0x41, 0x7f, 0x2e, 0xe7, 0xd2, 0x9a, 0xa9, 0xb1, 0x9c, 0x7a, 0xbe, 0x45, 0xb8, 0x33, 0x35,
	0x08, 0xf2, 0x04, 0x77, 0xd5, 0x9a, 0x97, 0xb2, 0x8e, 0xd0, 0x08, 0x8d, 0x8f, 0x72, 0xef, 0x11,
	0x82, 0x43, 0x41, 0x39, 0x44, 0x07, 0x56, 0xb5, 0x36, 0xe9, 0xe3, 0x8e, 0xaa, 0x68, 0x0d, 0xd1,
	0x83, 0x11, 0x1a, 0x1f, 0xe7, 0xce, 0x21, 0x67, 0xb8, 0xc7, 0x99, 0x28, 0x5a, 0xc1, 0x74, 0x14,
	0xda, 0xe8, 0x43, 0xce, 0xc4, 0x07, 0xc1, 0x34, 0x79, 0x8e, 0x4f, 0x98, 0x60, 0x9a, 0xd1, 0xba,
	0x50, 0xed, 0x62, 0x51, 0xaf, 0xa3, 0xce, 0x08, 0x8d, 0xc3, 0xfc, 0xd8, 0xab, 0xef, 0xad, 0x48,
	0x9e, 0x61, 0xcc, 0xe9, 0x6a, 0x1f, 0xd2, 0xb5, 0x21, 0x47, 0x9c, 0xae, 0xfc, 0xf5, 0xc0, 0x26,
	0xd0, 0xb4, 0xac, 0x21, 0x3a, 0x1c, 0xa1, 0x71, 0x2f, 0xff, 0xeb, 0x9b, 0x92, 0xe4, 0x67, 0x01,
	0x4d, 0xd4, 0xb3, 0x99, 0x9d, 0x73, 0x11, 0x7e, 0xb9, 0x1d, 0x06, 0xe7, 0xdf, 0x0e, 0x70, 0xf7,
	0x8a, 0x36, 0x94, 0x2b, 0x32, 0xc5, 0x27, 0xb6, 0x63, 0x85, 0xa6, 0xab, 0xa2, 0xa1, 0x1a, 0xdc,
	0x6f, 0x27, 0xc9, 0xe6, 0x7e, 0x18, 0xfc, 0xb8, 0x1f, 0xbe, 0x98, 0x33, 0xfd, 0xa9, 0x2d, 0x93,
	0x4a, 0xf2, 0xd4, 0x37, 0xd4, 0x1d, 0x2f, 0xd5, 0xec, 0x26, 0xd5, 0xeb, 0x05, 0xa8, 0xe4, 0x2d,
	0x54, 0xf9, 0x23, 0x4b, 0x99, 0xd2, 0x55, 0x4e, 0x35, 0x90, 0x2b, 0xdc, 0x67, 0x4a, 0xb5, 0x50,
	0x38, 0xb6, 0x19, 0x41, 0x71, 0x0d, 0xae, 0x67, 0x0f, 0x5f, 0x9d, 0x25, 0x0e, 0x91, 0x18, 0x3d,
	0xf1, 0xa3, 0x49, 0xde, 0x48, 0x26, 0x26, 0xa1, 0x49, 0x9b, 0x9f, 0xda, 0xc7, 0x76, 0x0c, 0x13,
	0xaa, 0xe0, 0x12, 0x80, 0x14, 0xb8, 0x6f, 0xbe, 0xe6, 0x81, 0xd7, 0x00, 0xa6, 0x58, 0x26, 0x6d,
	0xc3, 0xff, 0xbf, 0xda, 0x53, 0xc3, 0xb2, 0xfc, 0x4b, 0x80, 0xdc, 0x80, 0x2e, 0xc2, 0xdf, 0xb7,
	0x43, 0x34, 0x79, 0xb7, 0xf9, 0x15, 0x07, 0x9b, 0x6d, 0x8c, 0xee, 0xb6, 0x31, 0xfa, 0xb9, 0x8d,
	0xd1, 0xd7, 0x5d, 0x1c, 0xdc, 0xed, 0xe2, 0xe0, 0xfb, 0x2e, 0x0e, 0x3e, 0x66, 0xff, 0xe0, 0xcd,
	0xbe, 0x09, 0xd0, 0xe9, 0x7e, 0x2b, 0xb9, 0x9c, 0xb5, 0x35, 0x28, 0xbf, 0x9d, 0x36, 0x95, 0xd9,
	0xba, 0xae, 0x5d, 0xab, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x64, 0x54, 0x8b, 0x20, 0xbe,
	0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.TokenTaxRate.Equal(that1.TokenTaxRate) {
		return false
	}
	if !this.IssueTokenBaseFee.Equal(&that1.IssueTokenBaseFee) {
		return false
	}
	if !this.MintTokenFeeRatio.Equal(that1.MintTokenFeeRatio) {
		return false
	}
	return true
}
func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.MaxSupply != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x30
	}
	if m.InitialSupply != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.InitialSupply))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MinUnit) > 0 {
		i -= len(m.MinUnit)
		copy(dAtA[i:], m.MinUnit)
		i = encodeVarintToken(dAtA, i, uint64(len(m.MinUnit)))
		i--
		dAtA[i] = 0x22
	}
	if m.Scale != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Scale))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MintTokenFeeRatio.Size()
		i -= size
		if _, err := m.MintTokenFeeRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.IssueTokenBaseFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.TokenTaxRate.Size()
		i -= size
		if _, err := m.TokenTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Scale != 0 {
		n += 1 + sovToken(uint64(m.Scale))
	}
	l = len(m.MinUnit)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.InitialSupply != 0 {
		n += 1 + sovToken(uint64(m.InitialSupply))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovToken(uint64(m.MaxSupply))
	}
	if m.Mintable {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenTaxRate.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.IssueTokenBaseFee.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.MintTokenFeeRatio.Size()
	n += 1 + l + sovToken(uint64(l))
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scale", wireType)
			}
			m.Scale = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Scale |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSupply", wireType)
			}
			m.InitialSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
			m.Mintable = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueTokenBaseFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IssueTokenBaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintTokenFeeRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintTokenFeeRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
