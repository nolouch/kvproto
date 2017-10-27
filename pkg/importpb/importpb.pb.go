// Code generated by protoc-gen-gogo.
// source: importpb.proto
// DO NOT EDIT!

/*
	Package importpb is a generated protocol buffer package.

	It is generated from these files:
		importpb.proto

	It has these top-level messages:
		SSTMeta
		SSTHandle
		IngestSSTRequest
		IngestSSTResponse
		UploadSSTRequest
		UploadSSTResponse
*/
package importpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	metapb "github.com/pingcap/kvproto/pkg/metapb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SSTMeta struct {
	// The size of the file.
	Len uint64 `protobuf:"varint,1,opt,name=len,proto3" json:"len,omitempty"`
	// The CRC32 checksum of the file.
	Crc32 uint32 `protobuf:"varint,2,opt,name=crc32,proto3" json:"crc32,omitempty"`
	// The handle which identifies the file.
	Handle *SSTHandle `protobuf:"bytes,3,opt,name=handle" json:"handle,omitempty"`
}

func (m *SSTMeta) Reset()                    { *m = SSTMeta{} }
func (m *SSTMeta) String() string            { return proto.CompactTextString(m) }
func (*SSTMeta) ProtoMessage()               {}
func (*SSTMeta) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{0} }

func (m *SSTMeta) GetLen() uint64 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *SSTMeta) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

func (m *SSTMeta) GetHandle() *SSTHandle {
	if m != nil {
		return m.Handle
	}
	return nil
}

type SSTHandle struct {
	// The UUID of the file, to distinguish files with the same data.
	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The CF that this file will be ingested to.
	Cfname string `protobuf:"bytes,2,opt,name=cfname,proto3" json:"cfname,omitempty"`
	// The region that this file will be ingested to.
	RegionId uint64 `protobuf:"varint,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// The epoch of the region when this file is uploaded.
	RegionEpoch *metapb.RegionEpoch `protobuf:"bytes,4,opt,name=region_epoch,json=regionEpoch" json:"region_epoch,omitempty"`
}

func (m *SSTHandle) Reset()                    { *m = SSTHandle{} }
func (m *SSTHandle) String() string            { return proto.CompactTextString(m) }
func (*SSTHandle) ProtoMessage()               {}
func (*SSTHandle) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{1} }

func (m *SSTHandle) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *SSTHandle) GetCfname() string {
	if m != nil {
		return m.Cfname
	}
	return ""
}

func (m *SSTHandle) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *SSTHandle) GetRegionEpoch() *metapb.RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

type IngestSSTRequest struct {
	Context *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	// The handles of the files that will be ingested.
	// Only files of the same column family can be ingested atomically for now.
	Handles []*SSTHandle `protobuf:"bytes,2,rep,name=handles" json:"handles,omitempty"`
}

func (m *IngestSSTRequest) Reset()                    { *m = IngestSSTRequest{} }
func (m *IngestSSTRequest) String() string            { return proto.CompactTextString(m) }
func (*IngestSSTRequest) ProtoMessage()               {}
func (*IngestSSTRequest) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{2} }

func (m *IngestSSTRequest) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *IngestSSTRequest) GetHandles() []*SSTHandle {
	if m != nil {
		return m.Handles
	}
	return nil
}

type IngestSSTResponse struct {
	Error *errorpb.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *IngestSSTResponse) Reset()                    { *m = IngestSSTResponse{} }
func (m *IngestSSTResponse) String() string            { return proto.CompactTextString(m) }
func (*IngestSSTResponse) ProtoMessage()               {}
func (*IngestSSTResponse) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{3} }

func (m *IngestSSTResponse) GetError() *errorpb.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type UploadSSTRequest struct {
	Meta *SSTMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Data []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UploadSSTRequest) Reset()                    { *m = UploadSSTRequest{} }
func (m *UploadSSTRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadSSTRequest) ProtoMessage()               {}
func (*UploadSSTRequest) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{4} }

func (m *UploadSSTRequest) GetMeta() *SSTMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UploadSSTRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadSSTResponse struct {
}

func (m *UploadSSTResponse) Reset()                    { *m = UploadSSTResponse{} }
func (m *UploadSSTResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadSSTResponse) ProtoMessage()               {}
func (*UploadSSTResponse) Descriptor() ([]byte, []int) { return fileDescriptorImportpb, []int{5} }

func init() {
	proto.RegisterType((*SSTMeta)(nil), "importpb.SSTMeta")
	proto.RegisterType((*SSTHandle)(nil), "importpb.SSTHandle")
	proto.RegisterType((*IngestSSTRequest)(nil), "importpb.IngestSSTRequest")
	proto.RegisterType((*IngestSSTResponse)(nil), "importpb.IngestSSTResponse")
	proto.RegisterType((*UploadSSTRequest)(nil), "importpb.UploadSSTRequest")
	proto.RegisterType((*UploadSSTResponse)(nil), "importpb.UploadSSTResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Import service

type ImportClient interface {
	IngestSST(ctx context.Context, in *IngestSSTRequest, opts ...grpc.CallOption) (*IngestSSTResponse, error)
	UploadSST(ctx context.Context, opts ...grpc.CallOption) (Import_UploadSSTClient, error)
}

type importClient struct {
	cc *grpc.ClientConn
}

func NewImportClient(cc *grpc.ClientConn) ImportClient {
	return &importClient{cc}
}

func (c *importClient) IngestSST(ctx context.Context, in *IngestSSTRequest, opts ...grpc.CallOption) (*IngestSSTResponse, error) {
	out := new(IngestSSTResponse)
	err := grpc.Invoke(ctx, "/importpb.Import/IngestSST", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importClient) UploadSST(ctx context.Context, opts ...grpc.CallOption) (Import_UploadSSTClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Import_serviceDesc.Streams[0], c.cc, "/importpb.Import/UploadSST", opts...)
	if err != nil {
		return nil, err
	}
	x := &importUploadSSTClient{stream}
	return x, nil
}

type Import_UploadSSTClient interface {
	Send(*UploadSSTRequest) error
	CloseAndRecv() (*UploadSSTResponse, error)
	grpc.ClientStream
}

type importUploadSSTClient struct {
	grpc.ClientStream
}

func (x *importUploadSSTClient) Send(m *UploadSSTRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *importUploadSSTClient) CloseAndRecv() (*UploadSSTResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadSSTResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Import service

type ImportServer interface {
	IngestSST(context.Context, *IngestSSTRequest) (*IngestSSTResponse, error)
	UploadSST(Import_UploadSSTServer) error
}

func RegisterImportServer(s *grpc.Server, srv ImportServer) {
	s.RegisterService(&_Import_serviceDesc, srv)
}

func _Import_IngestSST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestSSTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).IngestSST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/importpb.Import/IngestSST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).IngestSST(ctx, req.(*IngestSSTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Import_UploadSST_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImportServer).UploadSST(&importUploadSSTServer{stream})
}

type Import_UploadSSTServer interface {
	SendAndClose(*UploadSSTResponse) error
	Recv() (*UploadSSTRequest, error)
	grpc.ServerStream
}

type importUploadSSTServer struct {
	grpc.ServerStream
}

func (x *importUploadSSTServer) SendAndClose(m *UploadSSTResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *importUploadSSTServer) Recv() (*UploadSSTRequest, error) {
	m := new(UploadSSTRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Import_serviceDesc = grpc.ServiceDesc{
	ServiceName: "importpb.Import",
	HandlerType: (*ImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestSST",
			Handler:    _Import_IngestSST_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadSST",
			Handler:       _Import_UploadSST_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "importpb.proto",
}

func (m *SSTMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Len != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Len))
	}
	if m.Crc32 != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Crc32))
	}
	if m.Handle != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Handle.Size()))
		n1, err := m.Handle.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SSTHandle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTHandle) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	if len(m.Cfname) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(len(m.Cfname)))
		i += copy(dAtA[i:], m.Cfname)
	}
	if m.RegionId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.RegionId))
	}
	if m.RegionEpoch != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.RegionEpoch.Size()))
		n2, err := m.RegionEpoch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *IngestSSTRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngestSSTRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Context.Size()))
		n3, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Handles) > 0 {
		for _, msg := range m.Handles {
			dAtA[i] = 0x12
			i++
			i = encodeVarintImportpb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *IngestSSTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngestSSTResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Error.Size()))
		n4, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *UploadSSTRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSSTRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Meta != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(m.Meta.Size()))
		n5, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImportpb(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *UploadSSTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSSTResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Importpb(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Importpb(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintImportpb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SSTMeta) Size() (n int) {
	var l int
	_ = l
	if m.Len != 0 {
		n += 1 + sovImportpb(uint64(m.Len))
	}
	if m.Crc32 != 0 {
		n += 1 + sovImportpb(uint64(m.Crc32))
	}
	if m.Handle != nil {
		l = m.Handle.Size()
		n += 1 + l + sovImportpb(uint64(l))
	}
	return n
}

func (m *SSTHandle) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovImportpb(uint64(l))
	}
	l = len(m.Cfname)
	if l > 0 {
		n += 1 + l + sovImportpb(uint64(l))
	}
	if m.RegionId != 0 {
		n += 1 + sovImportpb(uint64(m.RegionId))
	}
	if m.RegionEpoch != nil {
		l = m.RegionEpoch.Size()
		n += 1 + l + sovImportpb(uint64(l))
	}
	return n
}

func (m *IngestSSTRequest) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovImportpb(uint64(l))
	}
	if len(m.Handles) > 0 {
		for _, e := range m.Handles {
			l = e.Size()
			n += 1 + l + sovImportpb(uint64(l))
		}
	}
	return n
}

func (m *IngestSSTResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovImportpb(uint64(l))
	}
	return n
}

func (m *UploadSSTRequest) Size() (n int) {
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovImportpb(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovImportpb(uint64(l))
	}
	return n
}

func (m *UploadSSTResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovImportpb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImportpb(x uint64) (n int) {
	return sovImportpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SSTMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SSTMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Len", wireType)
			}
			m.Len = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Len |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc32", wireType)
			}
			m.Crc32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Crc32 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Handle == nil {
				m.Handle = &SSTHandle{}
			}
			if err := m.Handle.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func (m *SSTHandle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SSTHandle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTHandle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid[:0], dAtA[iNdEx:postIndex]...)
			if m.Uuid == nil {
				m.Uuid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cfname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cfname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionEpoch == nil {
				m.RegionEpoch = &metapb.RegionEpoch{}
			}
			if err := m.RegionEpoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func (m *IngestSSTRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IngestSSTRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngestSSTRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handles = append(m.Handles, &SSTHandle{})
			if err := m.Handles[len(m.Handles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func (m *IngestSSTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IngestSSTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngestSSTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &errorpb.Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func (m *UploadSSTRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadSSTRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSSTRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &SSTMeta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImportpb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func (m *UploadSSTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadSSTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSSTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipImportpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportpb
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
func skipImportpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImportpb
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
					return 0, ErrIntOverflowImportpb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowImportpb
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthImportpb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImportpb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipImportpb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthImportpb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImportpb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("importpb.proto", fileDescriptorImportpb) }

var fileDescriptorImportpb = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0xad, 0xd7, 0xac, 0x5d, 0x6e, 0xbb, 0xa9, 0xf5, 0x26, 0x14, 0x75, 0x52, 0x55, 0x45, 0x20,
	0x55, 0x43, 0x04, 0x29, 0x93, 0x90, 0x78, 0x05, 0x0d, 0xad, 0x0f, 0x7b, 0x71, 0xca, 0x2b, 0x28,
	0x73, 0x4c, 0x16, 0xb5, 0xb1, 0x8d, 0xe3, 0x56, 0x7c, 0x04, 0x1f, 0x80, 0xf8, 0x22, 0x1e, 0xf9,
	0x04, 0x54, 0x7e, 0x04, 0xc5, 0x4e, 0xd2, 0x50, 0xd8, 0x93, 0xef, 0xb9, 0xc7, 0x39, 0xe7, 0xdc,
	0xeb, 0xc0, 0x59, 0x96, 0x4b, 0xa1, 0xb4, 0xbc, 0x0f, 0xa4, 0x12, 0x5a, 0xe0, 0x93, 0x1a, 0x4f,
	0x86, 0x39, 0xd3, 0x71, 0xdd, 0x9f, 0x9c, 0xae, 0xb6, 0x4a, 0xd2, 0x3d, 0x64, 0x4a, 0x09, 0xd5,
	0xc0, 0x8b, 0x54, 0xa4, 0xc2, 0x94, 0x2f, 0xcb, 0xca, 0x76, 0xfd, 0x0f, 0xd0, 0x8f, 0xa2, 0xe5,
	0x1d, 0xd3, 0x31, 0x1e, 0x41, 0x77, 0xcd, 0xb8, 0x87, 0x66, 0x68, 0xee, 0x90, 0xb2, 0xc4, 0x17,
	0x70, 0x4c, 0x15, 0xbd, 0x0e, 0xbd, 0xa3, 0x19, 0x9a, 0x9f, 0x12, 0x0b, 0xf0, 0x73, 0xe8, 0x3d,
	0xc4, 0x3c, 0x59, 0x33, 0xaf, 0x3b, 0x43, 0xf3, 0x41, 0x78, 0x1e, 0x34, 0xf9, 0xa2, 0x68, 0x79,
	0x6b, 0x28, 0x52, 0x5d, 0xf1, 0xbf, 0x22, 0x70, 0x9b, 0x2e, 0xc6, 0xe0, 0x6c, 0x36, 0x59, 0x62,
	0x3c, 0x86, 0xc4, 0xd4, 0xf8, 0x09, 0xf4, 0xe8, 0x27, 0x1e, 0xe7, 0xcc, 0xb8, 0xb8, 0xa4, 0x42,
	0xf8, 0x12, 0x5c, 0xc5, 0xd2, 0x4c, 0xf0, 0x8f, 0x59, 0x62, 0x9c, 0x1c, 0x72, 0x62, 0x1b, 0x8b,
	0x04, 0xbf, 0x82, 0x61, 0x45, 0x32, 0x29, 0xe8, 0x83, 0xe7, 0x54, 0x49, 0xaa, 0x7d, 0x10, 0xc3,
	0xdd, 0x94, 0x14, 0x19, 0xa8, 0x3d, 0xf0, 0x73, 0x18, 0x2d, 0x78, 0xca, 0x0a, 0x1d, 0x45, 0x4b,
	0xc2, 0x3e, 0x6f, 0x58, 0xa1, 0xf1, 0x15, 0xf4, 0xa9, 0xe0, 0x9a, 0x7d, 0xd1, 0x26, 0xd7, 0x20,
	0x1c, 0x05, 0xf5, 0x22, 0xdf, 0xda, 0x3e, 0xa9, 0x2f, 0xe0, 0x17, 0xd0, 0xb7, 0x83, 0x15, 0xde,
	0xd1, 0xac, 0xfb, 0xd8, 0xf0, 0xf5, 0x1d, 0xff, 0x35, 0x8c, 0x5b, 0x76, 0x85, 0x14, 0xbc, 0x60,
	0xf8, 0x29, 0x1c, 0x9b, 0x97, 0xa9, 0xdc, 0xce, 0x82, 0xfa, 0x9d, 0x6e, 0xca, 0x93, 0x58, 0xd2,
	0xbf, 0x83, 0xd1, 0x7b, 0xb9, 0x16, 0x71, 0xd2, 0x4a, 0xfa, 0x0c, 0x9c, 0x72, 0xc0, 0xea, 0xc3,
	0xf1, 0x5f, 0xd6, 0xe5, 0x13, 0x12, 0x43, 0x97, 0x5b, 0x4e, 0x62, 0x1d, 0x9b, 0x7d, 0x0e, 0x89,
	0xa9, 0xfd, 0x73, 0x18, 0xb7, 0xe4, 0x6c, 0x92, 0xf0, 0x3b, 0x82, 0xde, 0xc2, 0x68, 0xe0, 0x77,
	0xe0, 0x36, 0x49, 0xf1, 0x64, 0xaf, 0x7c, 0xb8, 0xad, 0xc9, 0xe5, 0x7f, 0x39, 0x2b, 0xe8, 0x77,
	0xf0, 0x2d, 0xb8, 0x8d, 0x4f, 0x5b, 0xe7, 0x70, 0x96, 0xb6, 0xce, 0x3f, 0xc1, 0xfc, 0xce, 0x1c,
	0xbd, 0xb9, 0xfa, 0xb1, 0x9b, 0xa2, 0x9f, 0xbb, 0x29, 0xfa, 0xb5, 0x9b, 0xa2, 0x6f, 0xbf, 0xa7,
	0x1d, 0xf0, 0xa8, 0xc8, 0x03, 0x99, 0xf1, 0x94, 0xc6, 0x32, 0xd0, 0xd9, 0x6a, 0x1b, 0xac, 0xb6,
	0xe6, 0x2f, 0xbe, 0xef, 0x99, 0xe3, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x2d, 0x88,
	0xe6, 0x2a, 0x03, 0x00, 0x00,
}
