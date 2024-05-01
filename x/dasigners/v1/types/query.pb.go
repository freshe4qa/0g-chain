// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryEpochNumberRequest struct {
}

func (m *QueryEpochNumberRequest) Reset()         { *m = QueryEpochNumberRequest{} }
func (m *QueryEpochNumberRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEpochNumberRequest) ProtoMessage()    {}
func (*QueryEpochNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_991a610b84b5964c, []int{0}
}
func (m *QueryEpochNumberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochNumberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochNumberRequest.Merge(m, src)
}
func (m *QueryEpochNumberRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochNumberRequest proto.InternalMessageInfo

type QueryEpochNumberResponse struct {
	EpochNumber uint64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *QueryEpochNumberResponse) Reset()         { *m = QueryEpochNumberResponse{} }
func (m *QueryEpochNumberResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEpochNumberResponse) ProtoMessage()    {}
func (*QueryEpochNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_991a610b84b5964c, []int{1}
}
func (m *QueryEpochNumberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochNumberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochNumberResponse.Merge(m, src)
}
func (m *QueryEpochNumberResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochNumberResponse proto.InternalMessageInfo

type QueryEpochSignerSetRequest struct {
	EpochNumber uint64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *QueryEpochSignerSetRequest) Reset()         { *m = QueryEpochSignerSetRequest{} }
func (m *QueryEpochSignerSetRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEpochSignerSetRequest) ProtoMessage()    {}
func (*QueryEpochSignerSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_991a610b84b5964c, []int{2}
}
func (m *QueryEpochSignerSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochSignerSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochSignerSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochSignerSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochSignerSetRequest.Merge(m, src)
}
func (m *QueryEpochSignerSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochSignerSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochSignerSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochSignerSetRequest proto.InternalMessageInfo

type QueryEpochSignerSetResponse struct {
	Signers []*Signer `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *QueryEpochSignerSetResponse) Reset()         { *m = QueryEpochSignerSetResponse{} }
func (m *QueryEpochSignerSetResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEpochSignerSetResponse) ProtoMessage()    {}
func (*QueryEpochSignerSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_991a610b84b5964c, []int{3}
}
func (m *QueryEpochSignerSetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochSignerSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochSignerSetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochSignerSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochSignerSetResponse.Merge(m, src)
}
func (m *QueryEpochSignerSetResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochSignerSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochSignerSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochSignerSetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryEpochNumberRequest)(nil), "zgc.dasigners.v1.QueryEpochNumberRequest")
	proto.RegisterType((*QueryEpochNumberResponse)(nil), "zgc.dasigners.v1.QueryEpochNumberResponse")
	proto.RegisterType((*QueryEpochSignerSetRequest)(nil), "zgc.dasigners.v1.QueryEpochSignerSetRequest")
	proto.RegisterType((*QueryEpochSignerSetResponse)(nil), "zgc.dasigners.v1.QueryEpochSignerSetResponse")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/query.proto", fileDescriptor_991a610b84b5964c) }

var fileDescriptor_991a610b84b5964c = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0xef, 0xd2, 0x30,
	0x1c, 0x5d, 0xf1, 0x57, 0x52, 0x8c, 0x31, 0x8d, 0x89, 0x63, 0x92, 0x89, 0x8b, 0x31, 0x48, 0xdc,
	0x0a, 0x78, 0x36, 0x26, 0x26, 0x1e, 0x35, 0x01, 0x6e, 0x5e, 0x48, 0x37, 0x6b, 0x59, 0xc2, 0xda,
	0xb1, 0x76, 0x44, 0x38, 0xfa, 0x17, 0x90, 0x78, 0xf6, 0x1f, 0xf1, 0x2f, 0xe0, 0x48, 0xe2, 0xc5,
	0xa3, 0x82, 0x7f, 0x88, 0xa1, 0x2d, 0x22, 0xf0, 0xfd, 0x12, 0x6e, 0xeb, 0x7b, 0xef, 0xf3, 0xfa,
	0xde, 0xa7, 0x83, 0xf5, 0x39, 0x4b, 0xf0, 0x47, 0x22, 0x53, 0xc6, 0x69, 0x21, 0xf1, 0xb4, 0x83,
	0x27, 0x25, 0x2d, 0x66, 0x51, 0x5e, 0x08, 0x25, 0xd0, 0xfd, 0x39, 0x4b, 0xa2, 0x7f, 0x6c, 0x34,
	0xed, 0x78, 0xb5, 0x44, 0xc8, 0x4c, 0xc8, 0xa1, 0xe6, 0xb1, 0x39, 0x18, 0xb1, 0xf7, 0x80, 0x09,
	0x26, 0x0c, 0xbe, 0xfd, 0xb2, 0x68, 0x9d, 0x09, 0xc1, 0xc6, 0x14, 0x93, 0x3c, 0xc5, 0x84, 0x73,
	0xa1, 0x88, 0x4a, 0x05, 0xdf, 0xcd, 0xd4, 0x2c, 0xab, 0x4f, 0x71, 0xf9, 0x09, 0x13, 0x6e, 0xef,
	0xf6, 0x1e, 0x1f, 0x53, 0x2a, 0xcd, 0xa8, 0x54, 0x24, 0xcb, 0xad, 0xa0, 0x71, 0x12, 0x7d, 0x9f,
	0x54, 0x2b, 0x82, 0x1a, 0x7c, 0xd8, 0xdb, 0xb6, 0x79, 0x9b, 0x8b, 0x64, 0xf4, 0xbe, 0xcc, 0x62,
	0x5a, 0xf4, 0xe9, 0xa4, 0xa4, 0x52, 0x05, 0xaf, 0xa0, 0x7b, 0x4a, 0xc9, 0x5c, 0x70, 0x49, 0xd1,
	0x13, 0x78, 0x97, 0x6e, 0xe1, 0x21, 0xd7, 0xb8, 0x0b, 0x1a, 0xa0, 0x79, 0xb3, 0x5f, 0xa5, 0x7b,
	0x69, 0xf0, 0x1a, 0x7a, 0xfb, 0xf1, 0x81, 0xbe, 0x74, 0x40, 0x95, 0x35, 0xbf, 0xc4, 0xa0, 0x07,
	0x1f, 0x5d, 0x69, 0x60, 0x23, 0x74, 0xe1, 0x1d, 0x5b, 0xc5, 0x05, 0x8d, 0x1b, 0xcd, 0x6a, 0xd7,
	0x8d, 0x8e, 0x9f, 0x22, 0x32, 0x53, 0xfd, 0x9d, 0xb0, 0xfb, 0xbd, 0x02, 0x6f, 0x69, 0x4f, 0xb4,
	0x00, 0xb0, 0xfa, 0x5f, 0x31, 0xf4, 0xfc, 0x74, 0xf8, 0x9a, 0xbd, 0x78, 0xad, 0x4b, 0xa4, 0x26,
	0x64, 0xd0, 0xfa, 0xf2, 0xe3, 0xcf, 0xd7, 0xca, 0x53, 0x14, 0xe0, 0x36, 0x4b, 0x46, 0x24, 0xe5,
	0x87, 0xaf, 0xa1, 0xeb, 0x86, 0x66, 0x05, 0xe8, 0x1b, 0x80, 0xf7, 0x0e, 0xbb, 0xa2, 0x17, 0xe7,
	0xae, 0x3a, 0xde, 0xa9, 0x17, 0x5e, 0xa8, 0xb6, 0xd9, 0x22, 0x9d, 0xad, 0x89, 0x9e, 0x9d, 0xcb,
	0x66, 0x80, 0x50, 0x52, 0xf5, 0xe6, 0xdd, 0xf2, 0xb7, 0xef, 0x2c, 0xd7, 0x3e, 0x58, 0xad, 0x7d,
	0xf0, 0x6b, 0xed, 0x83, 0xc5, 0xc6, 0x77, 0x56, 0x1b, 0xdf, 0xf9, 0xb9, 0xf1, 0x9d, 0x0f, 0x98,
	0xa5, 0x6a, 0x54, 0xc6, 0x51, 0x22, 0x32, 0xdc, 0x66, 0x63, 0x12, 0x4b, 0xdc, 0x66, 0xa1, 0xf1,
	0xfd, 0x7c, 0xe8, 0xac, 0x66, 0x39, 0x95, 0xf1, 0x6d, 0xfd, 0x03, 0xbe, 0xfc, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0x31, 0x5d, 0xa6, 0x5f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	EpochNumber(ctx context.Context, in *QueryEpochNumberRequest, opts ...grpc.CallOption) (*QueryEpochNumberResponse, error)
	EpochSignerSet(ctx context.Context, in *QueryEpochSignerSetRequest, opts ...grpc.CallOption) (*QueryEpochSignerSetResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EpochNumber(ctx context.Context, in *QueryEpochNumberRequest, opts ...grpc.CallOption) (*QueryEpochNumberResponse, error) {
	out := new(QueryEpochNumberResponse)
	err := c.cc.Invoke(ctx, "/zgc.dasigners.v1.Query/EpochNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EpochSignerSet(ctx context.Context, in *QueryEpochSignerSetRequest, opts ...grpc.CallOption) (*QueryEpochSignerSetResponse, error) {
	out := new(QueryEpochSignerSetResponse)
	err := c.cc.Invoke(ctx, "/zgc.dasigners.v1.Query/EpochSignerSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	EpochNumber(context.Context, *QueryEpochNumberRequest) (*QueryEpochNumberResponse, error)
	EpochSignerSet(context.Context, *QueryEpochSignerSetRequest) (*QueryEpochSignerSetResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EpochNumber(ctx context.Context, req *QueryEpochNumberRequest) (*QueryEpochNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochNumber not implemented")
}
func (*UnimplementedQueryServer) EpochSignerSet(ctx context.Context, req *QueryEpochSignerSetRequest) (*QueryEpochSignerSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochSignerSet not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EpochNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.dasigners.v1.Query/EpochNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochNumber(ctx, req.(*QueryEpochNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EpochSignerSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochSignerSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochSignerSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.dasigners.v1.Query/EpochSignerSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochSignerSet(ctx, req.(*QueryEpochSignerSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zgc.dasigners.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EpochNumber",
			Handler:    _Query_EpochNumber_Handler,
		},
		{
			MethodName: "EpochSignerSet",
			Handler:    _Query_EpochSignerSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zgc/dasigners/v1/query.proto",
}

func (m *QueryEpochNumberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochNumberRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochNumberRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEpochNumberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochNumberResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochNumberResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEpochSignerSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochSignerSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochSignerSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEpochSignerSetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochSignerSetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochSignerSetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEpochNumberRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEpochNumberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovQuery(uint64(m.EpochNumber))
	}
	return n
}

func (m *QueryEpochSignerSetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovQuery(uint64(m.EpochNumber))
	}
	return n
}

func (m *QueryEpochSignerSetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEpochNumberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochNumberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochNumberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochNumberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochNumberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochNumberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochSignerSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochSignerSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochSignerSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochSignerSetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochSignerSetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochSignerSetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &Signer{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
