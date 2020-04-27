// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk/logical/identity.proto

package logical

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

type Entity struct {
	// ID is the unique identifier for the entity
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the entity
	Name string `sentinel:"" protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Aliases contains thhe alias mappings for the given entity
	Aliases []*Alias `sentinel:"" protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// Metadata represents the custom data tied to this entity
	Metadata map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Disabled is true if the entity is disabled.
	Disabled bool `sentinel:"" protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// NamespaceID is the identifier of the namespace to which this entity
	// belongs to.
	NamespaceID          string   `sentinel:"" protobuf:"bytes,6,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a34d35719c603a1, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Entity) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Entity) GetNamespaceID() string {
	if m != nil {
		return m.NamespaceID
	}
	return ""
}

type Alias struct {
	// MountType is the backend mount's type to which this identity belongs
	MountType string `sentinel:"" protobuf:"bytes,1,opt,name=mount_type,json=mountType,proto3" json:"mount_type,omitempty"`
	// MountAccessor is the identifier of the mount entry to which this
	// identity belongs
	MountAccessor string `sentinel:"" protobuf:"bytes,2,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty"`
	// Name is the identifier of this identity in its authentication source
	Name string `sentinel:"" protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this alias. Fields added
	// to it should have a low rate of change (or no change) because each
	// change incurs a storage write, so quickly-changing fields can have
	// a significant performance impact at scale. See the SDK's
	// "aliasmetadata" package for a helper that eases and standardizes
	// using this safely.
	Metadata map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID is the unique identifier for the alias
	ID string `sentinel:"" protobuf:"bytes,5,opt,name=ID,proto3" json:"ID,omitempty"`
	// NamespaceID is the identifier of the namespace to which this alias
	// belongs.
	NamespaceID          string   `sentinel:"" protobuf:"bytes,6,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a34d35719c603a1, []int{1}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Alias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Alias) GetNamespaceID() string {
	if m != nil {
		return m.NamespaceID
	}
	return ""
}

type Group struct {
	// ID is the unique identifier for the group
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the group
	Name string `sentinel:"" protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this group
	Metadata map[string]string `sentinel:"" protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NamespaceID is the identifier of the namespace to which this group
	// belongs to.
	NamespaceID          string   `sentinel:"" protobuf:"bytes,4,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a34d35719c603a1, []int{2}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Group) GetNamespaceID() string {
	if m != nil {
		return m.NamespaceID
	}
	return ""
}

func init() {
	proto.RegisterType((*Entity)(nil), "logical.Entity")
	proto.RegisterMapType((map[string]string)(nil), "logical.Entity.MetadataEntry")
	proto.RegisterType((*Alias)(nil), "logical.Alias")
	proto.RegisterMapType((map[string]string)(nil), "logical.Alias.MetadataEntry")
	proto.RegisterType((*Group)(nil), "logical.Group")
	proto.RegisterMapType((map[string]string)(nil), "logical.Group.MetadataEntry")
}

func init() {
	proto.RegisterFile("sdk/logical/identity.proto", fileDescriptor_4a34d35719c603a1)
}

var fileDescriptor_4a34d35719c603a1 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xd2, 0xf4, 0xcf, 0xf4, 0x6d, 0x79, 0x59, 0x3c, 0x84, 0x62, 0xa1, 0x16, 0x94,
	0x9c, 0x12, 0xd0, 0x4b, 0xd5, 0x53, 0x45, 0x91, 0x1e, 0xbc, 0x04, 0x4f, 0x5e, 0xca, 0x36, 0xbb,
	0xb4, 0x4b, 0x93, 0x6e, 0xc8, 0x6e, 0x0a, 0xf9, 0x0c, 0xe2, 0xc7, 0xf2, 0x7b, 0x99, 0x6e, 0xb6,
	0x21, 0xad, 0x14, 0x3c, 0xe8, 0x6d, 0xe6, 0x99, 0xc9, 0xec, 0xfe, 0x9e, 0xcd, 0xc0, 0x40, 0x90,
	0xb5, 0x1f, 0xf1, 0x25, 0x0b, 0x71, 0xe4, 0x33, 0x42, 0x37, 0x92, 0xc9, 0xdc, 0x4b, 0x52, 0x2e,
	0x39, 0x6a, 0x69, 0x7d, 0xfc, 0x61, 0x42, 0xf3, 0x49, 0x55, 0x50, 0x1f, 0xcc, 0xd9, 0xa3, 0x63,
	0x8c, 0x0c, 0xb7, 0x13, 0x14, 0x11, 0x42, 0xd0, 0xd8, 0xe0, 0x98, 0x3a, 0xa6, 0x52, 0x54, 0x8c,
	0x5c, 0x68, 0xe1, 0x88, 0x61, 0x41, 0x85, 0x63, 0x8d, 0x2c, 0xb7, 0x7b, 0xdd, 0xf7, 0xf4, 0x24,
	0x6f, 0xba, 0xd3, 0x83, 0x7d, 0x19, 0xdd, 0x42, 0x3b, 0xa6, 0x12, 0x13, 0x2c, 0xb1, 0xd3, 0x50,
	0xad, 0xc3, 0xaa, 0xb5, 0x3c, 0xd0, 0x7b, 0xd1, 0xf5, 0x22, 0x4d, 0xf3, 0xa0, 0x6a, 0x47, 0x03,
	0x68, 0x13, 0x26, 0xf0, 0x22, 0xa2, 0xc4, 0xb1, 0x8b, 0xc3, 0xdb, 0x41, 0x95, 0xa3, 0x0b, 0xf8,
	0xb7, 0xbb, 0x88, 0x48, 0x70, 0x48, 0xe7, 0x8c, 0x38, 0x4d, 0x75, 0xb9, 0x6e, 0xa5, 0xcd, 0xc8,
	0xe0, 0x1e, 0x7a, 0x07, 0x93, 0xd1, 0x7f, 0xb0, 0xd6, 0x34, 0xd7, 0x64, 0xbb, 0x10, 0x9d, 0x81,
	0xbd, 0xc5, 0x51, 0xb6, 0x67, 0x2b, 0x93, 0x3b, 0x73, 0x62, 0x8c, 0xdf, 0x4d, 0xb0, 0x15, 0x09,
	0x1a, 0x02, 0xc4, 0x3c, 0xdb, 0xc8, 0xb9, 0xcc, 0x13, 0xaa, 0x3f, 0xee, 0x28, 0xe5, 0xb5, 0x10,
	0xd0, 0x25, 0xf4, 0xcb, 0x32, 0x0e, 0x43, 0x2a, 0x04, 0x4f, 0xf5, 0xac, 0x9e, 0x52, 0xa7, 0x5a,
	0xac, 0x4c, 0xb4, 0x6a, 0x26, 0x4e, 0xbe, 0x59, 0x73, 0x7e, 0xe8, 0xe2, 0x49, 0x67, 0xca, 0x27,
	0xb2, 0xab, 0x27, 0xfa, 0x6b, 0x37, 0x3e, 0x0d, 0xb0, 0x9f, 0x53, 0x9e, 0x25, 0x3f, 0xfa, 0x39,
	0xea, 0x5c, 0xd6, 0x11, 0x97, 0x9a, 0x72, 0x92, 0xeb, 0x98, 0xa3, 0xf1, 0xbb, 0x1c, 0x0f, 0xee,
	0xdb, 0xd5, 0x92, 0xc9, 0x55, 0xb6, 0xf0, 0x42, 0x1e, 0xfb, 0x2b, 0x2c, 0x56, 0x2c, 0xe4, 0x69,
	0xe2, 0x6f, 0x71, 0x16, 0x49, 0xbf, 0xb6, 0x27, 0x8b, 0xa6, 0xda, 0x8f, 0x9b, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0xe2, 0x28, 0xc0, 0x3d, 0x03, 0x00, 0x00,
}
