// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: options/collection_permissions.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	options/collection_permissions.proto

It has these top-level messages:
	CollectionPermissions
	FilterOptions
	FieldPermissions
	MethodOptions
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CollectionPermissions_FilterType int32

const (
	CollectionPermissions_DEFAULT CollectionPermissions_FilterType = 0
	CollectionPermissions_STRING  CollectionPermissions_FilterType = 1
	CollectionPermissions_NUMBER  CollectionPermissions_FilterType = 2
)

var CollectionPermissions_FilterType_name = map[int32]string{
	0: "DEFAULT",
	1: "STRING",
	2: "NUMBER",
}
var CollectionPermissions_FilterType_value = map[string]int32{
	"DEFAULT": 0,
	"STRING":  1,
	"NUMBER":  2,
}

func (x CollectionPermissions_FilterType) Enum() *CollectionPermissions_FilterType {
	p := new(CollectionPermissions_FilterType)
	*p = x
	return p
}
func (x CollectionPermissions_FilterType) String() string {
	return proto.EnumName(CollectionPermissions_FilterType_name, int32(x))
}
func (x *CollectionPermissions_FilterType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CollectionPermissions_FilterType_value, data, "CollectionPermissions_FilterType")
	if err != nil {
		return err
	}
	*x = CollectionPermissions_FilterType(value)
	return nil
}
func (CollectionPermissions_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{0, 0}
}

type CollectionPermissions struct {
	DisableSorting     *bool                             `protobuf:"varint,1,opt,name=disable_sorting,json=disableSorting" json:"disable_sorting,omitempty"`
	Filters            *FilterOptions                    `protobuf:"bytes,2,opt,name=filters" json:"filters,omitempty"`
	FilterType         *CollectionPermissions_FilterType `protobuf:"varint,3,opt,name=filter_type,json=filterType,enum=atlas.query.CollectionPermissions_FilterType" json:"filter_type,omitempty"`
	EnableNestedFields *bool                             `protobuf:"varint,4,opt,name=enable_nested_fields,json=enableNestedFields" json:"enable_nested_fields,omitempty"`
	XXX_unrecognized   []byte                            `json:"-"`
}

func (m *CollectionPermissions) Reset()         { *m = CollectionPermissions{} }
func (m *CollectionPermissions) String() string { return proto.CompactTextString(m) }
func (*CollectionPermissions) ProtoMessage()    {}
func (*CollectionPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{0}
}

func (m *CollectionPermissions) GetDisableSorting() bool {
	if m != nil && m.DisableSorting != nil {
		return *m.DisableSorting
	}
	return false
}

func (m *CollectionPermissions) GetFilters() *FilterOptions {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *CollectionPermissions) GetFilterType() CollectionPermissions_FilterType {
	if m != nil && m.FilterType != nil {
		return *m.FilterType
	}
	return CollectionPermissions_DEFAULT
}

func (m *CollectionPermissions) GetEnableNestedFields() bool {
	if m != nil && m.EnableNestedFields != nil {
		return *m.EnableNestedFields
	}
	return false
}

type FilterOptions struct {
	// Types that are valid to be assigned to Operators:
	//	*FilterOptions_Allow
	//	*FilterOptions_Deny
	Operators        isFilterOptions_Operators `protobuf_oneof:"operators"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *FilterOptions) Reset()         { *m = FilterOptions{} }
func (m *FilterOptions) String() string { return proto.CompactTextString(m) }
func (*FilterOptions) ProtoMessage()    {}
func (*FilterOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{1}
}

type isFilterOptions_Operators interface {
	isFilterOptions_Operators()
}

type FilterOptions_Allow struct {
	Allow string `protobuf:"bytes,1,opt,name=allow,oneof"`
}
type FilterOptions_Deny struct {
	Deny string `protobuf:"bytes,2,opt,name=deny,oneof"`
}

func (*FilterOptions_Allow) isFilterOptions_Operators() {}
func (*FilterOptions_Deny) isFilterOptions_Operators()  {}

func (m *FilterOptions) GetOperators() isFilterOptions_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *FilterOptions) GetAllow() string {
	if x, ok := m.GetOperators().(*FilterOptions_Allow); ok {
		return x.Allow
	}
	return ""
}

func (m *FilterOptions) GetDeny() string {
	if x, ok := m.GetOperators().(*FilterOptions_Deny); ok {
		return x.Deny
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FilterOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FilterOptions_OneofMarshaler, _FilterOptions_OneofUnmarshaler, _FilterOptions_OneofSizer, []interface{}{
		(*FilterOptions_Allow)(nil),
		(*FilterOptions_Deny)(nil),
	}
}

func _FilterOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FilterOptions)
	// operators
	switch x := m.Operators.(type) {
	case *FilterOptions_Allow:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Allow)
	case *FilterOptions_Deny:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Deny)
	case nil:
	default:
		return fmt.Errorf("FilterOptions.Operators has unexpected type %T", x)
	}
	return nil
}

func _FilterOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FilterOptions)
	switch tag {
	case 1: // operators.allow
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operators = &FilterOptions_Allow{x}
		return true, err
	case 2: // operators.deny
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operators = &FilterOptions_Deny{x}
		return true, err
	default:
		return false, nil
	}
}

func _FilterOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FilterOptions)
	// operators
	switch x := m.Operators.(type) {
	case *FilterOptions_Allow:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Allow)))
		n += len(x.Allow)
	case *FilterOptions_Deny:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Deny)))
		n += len(x.Deny)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FieldPermissions struct {
	Name             *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	DisableSorting   *bool          `protobuf:"varint,2,opt,name=disable_sorting,json=disableSorting" json:"disable_sorting,omitempty"`
	Filters          *FilterOptions `protobuf:"bytes,3,opt,name=filters" json:"filters,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FieldPermissions) Reset()         { *m = FieldPermissions{} }
func (m *FieldPermissions) String() string { return proto.CompactTextString(m) }
func (*FieldPermissions) ProtoMessage()    {}
func (*FieldPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{2}
}

func (m *FieldPermissions) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FieldPermissions) GetDisableSorting() bool {
	if m != nil && m.DisableSorting != nil {
		return *m.DisableSorting
	}
	return false
}

func (m *FieldPermissions) GetFilters() *FilterOptions {
	if m != nil {
		return m.Filters
	}
	return nil
}

type MethodOptions struct {
	ForMessage         *string             `protobuf:"bytes,1,opt,name=for_message,json=forMessage" json:"for_message,omitempty"`
	Fields             []*FieldPermissions `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	AllowMissingFields *bool               `protobuf:"varint,3,opt,name=allow_missing_fields,json=allowMissingFields" json:"allow_missing_fields,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{3}
}

func (m *MethodOptions) GetForMessage() string {
	if m != nil && m.ForMessage != nil {
		return *m.ForMessage
	}
	return ""
}

func (m *MethodOptions) GetFields() []*FieldPermissions {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *MethodOptions) GetAllowMissingFields() bool {
	if m != nil && m.AllowMissingFields != nil {
		return *m.AllowMissingFields
	}
	return false
}

var E_Permissions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*CollectionPermissions)(nil),
	Field:         52121,
	Name:          "atlas.query.permissions",
	Tag:           "bytes,52121,opt,name=permissions",
	Filename:      "options/collection_permissions.proto",
}

var E_MethodPermissions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodOptions)(nil),
	Field:         52121,
	Name:          "atlas.query.method_permissions",
	Tag:           "bytes,52121,opt,name=method_permissions,json=methodPermissions",
	Filename:      "options/collection_permissions.proto",
}

func init() {
	proto.RegisterType((*CollectionPermissions)(nil), "atlas.query.CollectionPermissions")
	proto.RegisterType((*FilterOptions)(nil), "atlas.query.FilterOptions")
	proto.RegisterType((*FieldPermissions)(nil), "atlas.query.FieldPermissions")
	proto.RegisterType((*MethodOptions)(nil), "atlas.query.MethodOptions")
	proto.RegisterEnum("atlas.query.CollectionPermissions_FilterType", CollectionPermissions_FilterType_name, CollectionPermissions_FilterType_value)
	proto.RegisterExtension(E_Permissions)
	proto.RegisterExtension(E_MethodPermissions)
}

func init() {
	proto.RegisterFile("options/collection_permissions.proto", fileDescriptorCollectionPermissions)
}

var fileDescriptorCollectionPermissions = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x9d, 0xd2, 0x92, 0xb1, 0x5a, 0xc2, 0xaa, 0x20, 0xab, 0x52, 0xc1, 0xb2, 0x90, 0xc8,
	0x25, 0x0e, 0x44, 0x70, 0x29, 0x27, 0x52, 0x12, 0x7e, 0x44, 0x02, 0xda, 0xa6, 0x17, 0x2e, 0x96,
	0x13, 0x8f, 0x5d, 0x0b, 0x67, 0xd7, 0xec, 0x6e, 0x04, 0x79, 0x00, 0x9e, 0x01, 0xf5, 0x19, 0x78,
	0x49, 0xe4, 0x5d, 0x27, 0xb5, 0xa1, 0x52, 0x39, 0x65, 0x3c, 0xdf, 0xe4, 0x9b, 0x6f, 0xbe, 0x99,
	0x85, 0x27, 0xbc, 0x50, 0x19, 0x67, 0xb2, 0xbf, 0xe0, 0x79, 0x8e, 0x8b, 0x32, 0x0e, 0x0b, 0x14,
	0xcb, 0x4c, 0xca, 0x32, 0x1d, 0x14, 0x82, 0x2b, 0x4e, 0x9c, 0x48, 0xe5, 0x91, 0x0c, 0xbe, 0xad,
	0x50, 0xac, 0x8f, 0xbd, 0x94, 0xf3, 0x34, 0xc7, 0xbe, 0x86, 0xe6, 0xab, 0xa4, 0x1f, 0xa3, 0x5c,
	0x88, 0xac, 0x50, 0x5c, 0x98, 0x72, 0xff, 0xb7, 0x0d, 0x0f, 0xce, 0xb6, 0x7c, 0x9f, 0xaf, 0xe9,
	0xc8, 0x53, 0xb8, 0x17, 0x67, 0x32, 0x9a, 0xe7, 0x18, 0x4a, 0x2e, 0x54, 0xc6, 0x52, 0xd7, 0xf2,
	0xac, 0xee, 0x5d, 0x7a, 0x58, 0xa5, 0xcf, 0x4d, 0x96, 0xbc, 0x80, 0xfd, 0x24, 0xcb, 0x15, 0x0a,
	0xe9, 0xda, 0x9e, 0xd5, 0x75, 0x06, 0xc7, 0x41, 0x4d, 0x43, 0x30, 0xd6, 0xd8, 0x27, 0xa3, 0x9d,
	0x6e, 0x4a, 0xc9, 0x14, 0x1c, 0x13, 0x86, 0x6a, 0x5d, 0xa0, 0xdb, 0xf2, 0xac, 0xee, 0xe1, 0xa0,
	0xd7, 0xf8, 0xe7, 0x8d, 0xba, 0x2a, 0xbe, 0xd9, 0xba, 0x40, 0x0a, 0xc9, 0x36, 0x26, 0xcf, 0xe0,
	0x08, 0x99, 0x56, 0xcb, 0x50, 0x2a, 0x8c, 0xc3, 0x24, 0xc3, 0x3c, 0x96, 0xee, 0xae, 0xd6, 0x4c,
	0x0c, 0x36, 0xd5, 0xd0, 0x58, 0x23, 0xfe, 0x73, 0x80, 0x6b, 0x2e, 0xe2, 0xc0, 0xfe, 0x9b, 0xd1,
	0xf8, 0xf5, 0xc5, 0xc7, 0x59, 0x67, 0x87, 0x00, 0xec, 0x9d, 0xcf, 0xe8, 0xfb, 0xe9, 0xdb, 0x8e,
	0x55, 0xc6, 0xd3, 0x8b, 0xc9, 0x70, 0x44, 0x3b, 0xb6, 0xff, 0x01, 0x0e, 0x1a, 0xe3, 0x90, 0x87,
	0x70, 0x27, 0xca, 0x73, 0xfe, 0x5d, 0x5b, 0xd3, 0x7e, 0xb7, 0x43, 0xcd, 0x27, 0x39, 0x82, 0xdd,
	0x18, 0xd9, 0x5a, 0x1b, 0x52, 0xa6, 0xf5, 0xd7, 0xd0, 0x81, 0x36, 0x2f, 0x50, 0x44, 0x8a, 0x0b,
	0xe9, 0xff, 0xb4, 0xa0, 0xa3, 0x95, 0xd4, 0x4d, 0x27, 0xb0, 0xcb, 0xa2, 0x25, 0xba, 0x96, 0x67,
	0x77, 0xdb, 0x54, 0xc7, 0x37, 0x2d, 0xc2, 0xbe, 0x6d, 0x11, 0xad, 0xff, 0x5e, 0x84, 0x7f, 0x65,
	0xc1, 0xc1, 0x04, 0xd5, 0x25, 0x8f, 0x37, 0x43, 0x3d, 0x06, 0x27, 0xe1, 0x22, 0x5c, 0xa2, 0x94,
	0x51, 0x8a, 0x66, 0x34, 0x0a, 0x09, 0x17, 0x13, 0x93, 0x21, 0x2f, 0x61, 0xaf, 0x72, 0xd7, 0xf6,
	0x5a, 0x5d, 0x67, 0x70, 0xf2, 0x57, 0x9f, 0xe6, 0x50, 0xb4, 0x2a, 0x2e, 0x57, 0xa4, 0xdd, 0x09,
	0x35, 0xc2, 0xd2, 0xcd, 0x8a, 0x5a, 0x66, 0x45, 0x1a, 0x9b, 0x18, 0xc8, 0xac, 0xe8, 0x14, 0xc1,
	0xa9, 0x5d, 0x38, 0x39, 0x09, 0xcc, 0x3d, 0x07, 0x9b, 0x7b, 0x36, 0xbd, 0x2a, 0xdd, 0xee, 0xd5,
	0x2f, 0x33, 0xb6, 0x7f, 0xfb, 0x15, 0xd1, 0x3a, 0xef, 0xe9, 0x57, 0x20, 0x4b, 0xed, 0x40, 0xfd,
	0x3d, 0x91, 0x47, 0xff, 0x74, 0x6b, 0xd8, 0xb4, 0x6d, 0xd7, 0x74, 0xb9, 0x51, 0x43, 0xef, 0x1b,
	0xde, 0x5a, 0xe7, 0xe1, 0xe8, 0xcb, 0x59, 0x9a, 0xa9, 0xcb, 0xd5, 0x3c, 0x58, 0xf0, 0x65, 0x3f,
	0x63, 0x09, 0x9f, 0xe7, 0xfc, 0x07, 0x2f, 0x90, 0x99, 0x67, 0xba, 0xe8, 0xa5, 0xc8, 0x7a, 0x9a,
	0xb2, 0xa7, 0x29, 0x7b, 0xa5, 0xa8, 0x7e, 0xf5, 0xf8, 0x5f, 0x55, 0xbf, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x66, 0x78, 0x42, 0x0e, 0x04, 0x00, 0x00,
}
