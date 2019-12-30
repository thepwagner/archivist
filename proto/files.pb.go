// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/files.proto

package archivist

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

// Blob is an archived collection of bytes.
type Blob struct {
	BlobId               []byte          `protobuf:"bytes,1,opt,name=blob_id,json=blobId,proto3" json:"blob_id,omitempty"`
	Size                 uint64          `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Integrity            *Blob_Integrity `protobuf:"bytes,3,opt,name=integrity,proto3" json:"integrity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_132a9b3180f1f907, []int{0}
}

func (m *Blob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob.Unmarshal(m, b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return xxx_messageInfo_Blob.Size(m)
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetBlobId() []byte {
	if m != nil {
		return m.BlobId
	}
	return nil
}

func (m *Blob) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Blob) GetIntegrity() *Blob_Integrity {
	if m != nil {
		return m.Integrity
	}
	return nil
}

type Blob_Integrity struct {
	Sha512               []byte   `protobuf:"bytes,1,opt,name=sha512,proto3" json:"sha512,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blob_Integrity) Reset()         { *m = Blob_Integrity{} }
func (m *Blob_Integrity) String() string { return proto.CompactTextString(m) }
func (*Blob_Integrity) ProtoMessage()    {}
func (*Blob_Integrity) Descriptor() ([]byte, []int) {
	return fileDescriptor_132a9b3180f1f907, []int{0, 0}
}

func (m *Blob_Integrity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob_Integrity.Unmarshal(m, b)
}
func (m *Blob_Integrity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob_Integrity.Marshal(b, m, deterministic)
}
func (m *Blob_Integrity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob_Integrity.Merge(m, src)
}
func (m *Blob_Integrity) XXX_Size() int {
	return xxx_messageInfo_Blob_Integrity.Size(m)
}
func (m *Blob_Integrity) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob_Integrity.DiscardUnknown(m)
}

var xxx_messageInfo_Blob_Integrity proto.InternalMessageInfo

func (m *Blob_Integrity) GetSha512() []byte {
	if m != nil {
		return m.Sha512
	}
	return nil
}

// BlobIDs are a collection of Blobs by reference.
type BlobIDs struct {
	BlobIds              [][]byte `protobuf:"bytes,1,rep,name=blob_ids,json=blobIds,proto3" json:"blob_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlobIDs) Reset()         { *m = BlobIDs{} }
func (m *BlobIDs) String() string { return proto.CompactTextString(m) }
func (*BlobIDs) ProtoMessage()    {}
func (*BlobIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_132a9b3180f1f907, []int{1}
}

func (m *BlobIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlobIDs.Unmarshal(m, b)
}
func (m *BlobIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlobIDs.Marshal(b, m, deterministic)
}
func (m *BlobIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlobIDs.Merge(m, src)
}
func (m *BlobIDs) XXX_Size() int {
	return xxx_messageInfo_BlobIDs.Size(m)
}
func (m *BlobIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlobIDs.DiscardUnknown(m)
}

var xxx_messageInfo_BlobIDs proto.InternalMessageInfo

func (m *BlobIDs) GetBlobIds() [][]byte {
	if m != nil {
		return m.BlobIds
	}
	return nil
}

// Filenames is an index of filenames to Blobs.
type Filenames struct {
	Index                map[string]*BlobIDs `protobuf:"bytes,1,rep,name=index,proto3" json:"index,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filenames) Reset()         { *m = Filenames{} }
func (m *Filenames) String() string { return proto.CompactTextString(m) }
func (*Filenames) ProtoMessage()    {}
func (*Filenames) Descriptor() ([]byte, []int) {
	return fileDescriptor_132a9b3180f1f907, []int{2}
}

func (m *Filenames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filenames.Unmarshal(m, b)
}
func (m *Filenames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filenames.Marshal(b, m, deterministic)
}
func (m *Filenames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filenames.Merge(m, src)
}
func (m *Filenames) XXX_Size() int {
	return xxx_messageInfo_Filenames.Size(m)
}
func (m *Filenames) XXX_DiscardUnknown() {
	xxx_messageInfo_Filenames.DiscardUnknown(m)
}

var xxx_messageInfo_Filenames proto.InternalMessageInfo

func (m *Filenames) GetIndex() map[string]*BlobIDs {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterType((*Blob)(nil), "archivist.Blob")
	proto.RegisterType((*Blob_Integrity)(nil), "archivist.Blob.Integrity")
	proto.RegisterType((*BlobIDs)(nil), "archivist.BlobIDs")
	proto.RegisterType((*Filenames)(nil), "archivist.Filenames")
	proto.RegisterMapType((map[string]*BlobIDs)(nil), "archivist.Filenames.IndexEntry")
}

func init() { proto.RegisterFile("proto/files.proto", fileDescriptor_132a9b3180f1f907) }

var fileDescriptor_132a9b3180f1f907 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0xb3, 0x85, 0xc2, 0x9f, 0xa1, 0x87, 0xbf, 0x73, 0x50, 0xda, 0x4b, 0x11, 0x3d, 0x70,
	0x5a, 0x22, 0xa6, 0xd1, 0x78, 0x6c, 0xaa, 0x09, 0x89, 0xa7, 0x3d, 0x7a, 0x31, 0x50, 0x56, 0xd8,
	0x48, 0xa1, 0x61, 0xb7, 0x55, 0xfc, 0x14, 0xfa, 0x8d, 0x0d, 0x60, 0xa9, 0xf1, 0xf6, 0x66, 0xf2,
	0xdb, 0xb7, 0xef, 0x65, 0xe0, 0x64, 0x5b, 0x57, 0xaa, 0x0a, 0x5e, 0x44, 0xc1, 0x25, 0xed, 0x34,
	0x5a, 0x71, 0xbd, 0xce, 0xc5, 0x5e, 0x48, 0xe5, 0x7d, 0x11, 0xd0, 0x97, 0x45, 0x95, 0xe0, 0x19,
	0x98, 0x49, 0x51, 0x25, 0xcf, 0x22, 0x75, 0x88, 0x4b, 0xfc, 0x09, 0x33, 0xda, 0x31, 0x4a, 0x11,
	0x41, 0x97, 0xe2, 0x83, 0x3b, 0x23, 0x97, 0xf8, 0x3a, 0xeb, 0x34, 0xde, 0x80, 0x25, 0x4a, 0xc5,
	0xb3, 0x5a, 0xa8, 0xc6, 0xd1, 0x5c, 0xe2, 0xdb, 0xe1, 0x94, 0x0e, 0xa6, 0xb4, 0x35, 0xa4, 0xd1,
	0x01, 0x60, 0x47, 0x76, 0x76, 0x01, 0xd6, 0xb0, 0xc7, 0x53, 0x30, 0x64, 0x1e, 0x2f, 0xae, 0xc2,
	0xc3, 0x8f, 0xfd, 0xe4, 0x5d, 0x82, 0xd9, 0x3a, 0x44, 0x2b, 0x89, 0x53, 0xf8, 0xf7, 0x93, 0x4a,
	0x3a, 0xc4, 0xd5, 0xfc, 0x09, 0x33, 0xfb, 0x58, 0xd2, 0xfb, 0x24, 0x60, 0x3d, 0x88, 0x82, 0x97,
	0xf1, 0x86, 0x4b, 0x5c, 0xc0, 0x58, 0x94, 0x29, 0x7f, 0xef, 0x28, 0x3b, 0x9c, 0xff, 0x4a, 0x33,
	0x40, 0x34, 0x6a, 0x89, 0xfb, 0x52, 0xd5, 0x0d, 0xeb, 0xe9, 0xd9, 0x23, 0xc0, 0x71, 0x89, 0xff,
	0x41, 0x7b, 0xe5, 0x4d, 0x97, 0xc6, 0x62, 0xad, 0x44, 0x1f, 0xc6, 0xfb, 0xb8, 0xd8, 0xf5, 0xed,
	0xed, 0x10, 0xff, 0x94, 0x8c, 0x56, 0x92, 0xf5, 0xc0, 0xdd, 0xe8, 0x96, 0x2c, 0xcf, 0x9f, 0xe6,
	0x99, 0x50, 0xf9, 0x2e, 0xa1, 0xeb, 0x6a, 0x13, 0xa8, 0x9c, 0x6f, 0xdf, 0xe2, 0xac, 0xe4, 0x75,
	0x30, 0xbc, 0x4a, 0x8c, 0xee, 0x02, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x6f, 0xed,
	0xee, 0x96, 0x01, 0x00, 0x00,
}
