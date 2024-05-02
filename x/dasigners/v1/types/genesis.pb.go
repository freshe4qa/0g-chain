// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Params struct {
	QuorumSize    uint64 `protobuf:"varint,1,opt,name=quorum_size,json=quorumSize,proto3" json:"quorum_size,omitempty"`
	TokensPerVote string `protobuf:"bytes,2,opt,name=tokens_per_vote,json=tokensPerVote,proto3" json:"tokens_per_vote,omitempty"`
	MaxVotes      uint64 `protobuf:"varint,3,opt,name=max_votes,json=maxVotes,proto3" json:"max_votes,omitempty"`
	EpochBlocks   uint64 `protobuf:"varint,4,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{0}
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

func (m *Params) GetQuorumSize() uint64 {
	if m != nil {
		return m.QuorumSize
	}
	return 0
}

func (m *Params) GetTokensPerVote() string {
	if m != nil {
		return m.TokensPerVote
	}
	return ""
}

func (m *Params) GetMaxVotes() uint64 {
	if m != nil {
		return m.MaxVotes
	}
	return 0
}

func (m *Params) GetEpochBlocks() uint64 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

// GenesisState defines the dasigners module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// params epoch_number the epoch number
	EpochNumber uint64 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	// signers defines all signers information
	Signers []*Signer `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
	// signers_by_epoch defines chosen signers by epoch
	SignersByEpoch []*EpochSignerSet `protobuf:"bytes,4,rep,name=signers_by_epoch,json=signersByEpoch,proto3" json:"signers_by_epoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *GenesisState) GetSigners() []*Signer {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *GenesisState) GetSignersByEpoch() []*EpochSignerSet {
	if m != nil {
		return m.SignersByEpoch
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "zgc.dasigners.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "zgc.dasigners.v1.GenesisState")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/genesis.proto", fileDescriptor_896efa766aaca3be) }

var fileDescriptor_896efa766aaca3be = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x6b, 0x1a, 0x15, 0xe6, 0x0e, 0x98, 0x2c, 0x0e, 0xd9, 0x90, 0xd2, 0xb0, 0x03, 0xda,
	0x85, 0x78, 0x1b, 0x12, 0x0f, 0x10, 0x09, 0x21, 0x38, 0xa0, 0x29, 0x91, 0x38, 0x70, 0x89, 0x9c,
	0xf0, 0xc7, 0x8d, 0x56, 0xc7, 0x21, 0x76, 0xaa, 0x26, 0x4f, 0x01, 0x6f, 0xb5, 0xe3, 0x8e, 0x9c,
	0x10, 0x6a, 0x4f, 0xbc, 0x05, 0xea, 0xdf, 0x19, 0x13, 0xeb, 0x6e, 0x7f, 0x7f, 0xbf, 0xcf, 0x9f,
	0x3f, 0xdb, 0x34, 0xe8, 0x65, 0xc1, 0xbf, 0x08, 0x53, 0xca, 0x0a, 0x1a, 0xc3, 0x97, 0x67, 0x5c,
	0x42, 0x05, 0xa6, 0x34, 0x51, 0xdd, 0x68, 0xab, 0xd9, 0x41, 0x2f, 0x8b, 0xe8, 0x1f, 0x8f, 0x96,
	0x67, 0x47, 0x87, 0x85, 0x36, 0x4a, 0x9b, 0x0c, 0x39, 0x77, 0x0b, 0x67, 0x3e, 0x7a, 0x26, 0xb5,
	0xd4, 0x4e, 0xdf, 0x4e, 0x83, 0x7a, 0x28, 0xb5, 0x96, 0x0b, 0xe0, 0xb8, 0xca, 0xdb, 0xaf, 0x5c,
	0x54, 0xdd, 0x80, 0x66, 0x77, 0x91, 0x2d, 0x15, 0x18, 0x2b, 0x54, 0x3d, 0x18, 0xc2, 0x9d, 0x7a,
	0xb7, 0x5d, 0xd0, 0x71, 0xfc, 0x83, 0xd0, 0xc9, 0x85, 0x68, 0x84, 0x32, 0x6c, 0x46, 0xa7, 0xdf,
	0x5a, 0xdd, 0xb4, 0x2a, 0x33, 0x65, 0x0f, 0x3e, 0x09, 0xc9, 0x89, 0x97, 0x50, 0x27, 0xa5, 0x65,
	0x0f, 0xec, 0x25, 0x7d, 0x6a, 0xf5, 0x25, 0x54, 0x26, 0xab, 0xa1, 0xc9, 0x96, 0xda, 0x82, 0xff,
	0x20, 0x24, 0x27, 0x7b, 0xc9, 0x63, 0x27, 0x5f, 0x40, 0xf3, 0x49, 0x5b, 0x60, 0xcf, 0xe9, 0x9e,
	0x12, 0x2b, 0x34, 0x18, 0x7f, 0x8c, 0x31, 0x8f, 0x94, 0x58, 0x6d, 0x99, 0x61, 0x2f, 0xe8, 0x3e,
	0xd4, 0xba, 0x98, 0x67, 0xf9, 0x42, 0x17, 0x97, 0xc6, 0xf7, 0x90, 0x4f, 0x51, 0x8b, 0x51, 0x3a,
	0xfe, 0x43, 0xe8, 0xfe, 0x3b, 0xf7, 0x8c, 0xa9, 0x15, 0x16, 0xd8, 0x1b, 0x3a, 0xa9, 0xb1, 0x23,
	0x96, 0x9a, 0x9e, 0xfb, 0xd1, 0xdd, 0x67, 0x8d, 0xdc, 0x1d, 0x62, 0xef, 0xea, 0xd7, 0x6c, 0x94,
	0x0c, 0xee, 0xdb, 0xb3, 0xaa, 0x56, 0xe5, 0xd0, 0x60, 0xdb, 0x9b, 0xb3, 0x3e, 0xa2, 0xc4, 0xce,
	0xe9, 0xc3, 0x21, 0xc5, 0x1f, 0x87, 0xe3, 0xfb, 0xb3, 0x53, 0x1c, 0x93, 0x1b, 0x23, 0xfb, 0x40,
	0x0f, 0x86, 0x31, 0xcb, 0xbb, 0x0c, 0xd3, 0x7c, 0x0f, 0x37, 0x87, 0xbb, 0x9b, 0xdf, 0x6e, 0xb1,
	0x4b, 0x48, 0xc1, 0x26, 0x4f, 0x06, 0x14, 0x77, 0x08, 0xe2, 0xf7, 0x57, 0xeb, 0x80, 0x5c, 0xaf,
	0x03, 0xf2, 0x7b, 0x1d, 0x90, 0xef, 0x9b, 0x60, 0x74, 0xbd, 0x09, 0x46, 0x3f, 0x37, 0xc1, 0xe8,
	0x33, 0x97, 0xa5, 0x9d, 0xb7, 0x79, 0x54, 0x68, 0xc5, 0x4f, 0xe5, 0x42, 0xe4, 0x86, 0x9f, 0xca,
	0x57, 0xc5, 0x5c, 0x94, 0x15, 0x5f, 0xfd, 0xff, 0xa9, 0xb6, 0xab, 0xc1, 0xe4, 0x13, 0xfc, 0xd1,
	0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x84, 0xf4, 0xab, 0x94, 0x02, 0x00, 0x00,
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
	if m.EpochBlocks != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxVotes != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxVotes))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TokensPerVote) > 0 {
		i -= len(m.TokensPerVote)
		copy(dAtA[i:], m.TokensPerVote)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokensPerVote)))
		i--
		dAtA[i] = 0x12
	}
	if m.QuorumSize != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.QuorumSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignersByEpoch) > 0 {
		for iNdEx := len(m.SignersByEpoch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignersByEpoch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EpochNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuorumSize != 0 {
		n += 1 + sovGenesis(uint64(m.QuorumSize))
	}
	l = len(m.TokensPerVote)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxVotes != 0 {
		n += 1 + sovGenesis(uint64(m.MaxVotes))
	}
	if m.EpochBlocks != 0 {
		n += 1 + sovGenesis(uint64(m.EpochBlocks))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovGenesis(uint64(m.EpochNumber))
	}
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SignersByEpoch) > 0 {
		for _, e := range m.SignersByEpoch {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumSize", wireType)
			}
			m.QuorumSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuorumSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerVote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokensPerVote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotes", wireType)
			}
			m.MaxVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &Signer{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignersByEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignersByEpoch = append(m.SignersByEpoch, &EpochSignerSet{})
			if err := m.SignersByEpoch[len(m.SignersByEpoch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)