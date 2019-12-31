// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/index.proto

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

// Index stores data for later retrieval.
type Index struct {
	Drives               []*Drive        `protobuf:"bytes,1,rep,name=drives,proto3" json:"drives,omitempty"`
	Blobs                []*Blob         `protobuf:"bytes,2,rep,name=blobs,proto3" json:"blobs,omitempty"`
	BlobFilenames        map[string]*IDs `protobuf:"bytes,3,rep,name=blob_filenames,json=blobFilenames,proto3" json:"blob_filenames,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dde39a7f93d1aab, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetDrives() []*Drive {
	if m != nil {
		return m.Drives
	}
	return nil
}

func (m *Index) GetBlobs() []*Blob {
	if m != nil {
		return m.Blobs
	}
	return nil
}

func (m *Index) GetBlobFilenames() map[string]*IDs {
	if m != nil {
		return m.BlobFilenames
	}
	return nil
}

func init() {
	proto.RegisterType((*Index)(nil), "archivist.v1.Index")
	proto.RegisterMapType((map[string]*IDs)(nil), "archivist.v1.Index.BlobFilenamesEntry")
}

func init() { proto.RegisterFile("proto/index.proto", fileDescriptor_2dde39a7f93d1aab) }

var fileDescriptor_2dde39a7f93d1aab = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0xd0, 0x03, 0xb3, 0x85, 0x78, 0x12, 0x8b, 0x92, 0x33,
	0x32, 0xcb, 0x32, 0x8b, 0x4b, 0xf4, 0xca, 0x0c, 0xa5, 0xa0, 0x0a, 0xd2, 0x32, 0x73, 0x52, 0x8b,
	0x21, 0x0a, 0x94, 0x9a, 0x98, 0xb8, 0x58, 0x3d, 0x41, 0x1a, 0x84, 0xb4, 0xb9, 0xd8, 0x52, 0x8a,
	0x32, 0xcb, 0x52, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x90, 0xf5, 0xea,
	0xb9, 0x80, 0xe4, 0x82, 0xa0, 0x4a, 0x84, 0x34, 0xb8, 0x58, 0x93, 0x72, 0xf2, 0x93, 0x8a, 0x25,
	0x98, 0xc0, 0x6a, 0x85, 0x50, 0xd5, 0x3a, 0xe5, 0xe4, 0x27, 0x05, 0x41, 0x14, 0x08, 0xf9, 0x72,
	0xf1, 0x81, 0x18, 0xf1, 0x20, 0x4b, 0xf3, 0x12, 0x73, 0x53, 0x8b, 0x25, 0x98, 0xc1, 0x5a, 0xd4,
	0x50, 0xb5, 0x80, 0xdd, 0x00, 0xd6, 0xe8, 0x06, 0x53, 0xe8, 0x9a, 0x57, 0x52, 0x54, 0x19, 0xc4,
	0x9b, 0x84, 0x2c, 0x26, 0x15, 0xcc, 0x25, 0x84, 0xa9, 0x48, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5,
	0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x52, 0xe7, 0x62, 0x2d, 0x4b, 0xcc,
	0x29, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0x44, 0xb3, 0xcd, 0xa5, 0x38, 0x08,
	0x22, 0x6f, 0xc5, 0x64, 0xc1, 0xe8, 0xa4, 0x18, 0x25, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x92, 0x91, 0x5a, 0x50, 0x9e, 0x98, 0x9e, 0x97, 0x5a, 0xa4, 0x0f,
	0xd7, 0x94, 0xc4, 0x06, 0x0e, 0x2e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xdc, 0xc5,
	0x8f, 0x64, 0x01, 0x00, 0x00,
}
