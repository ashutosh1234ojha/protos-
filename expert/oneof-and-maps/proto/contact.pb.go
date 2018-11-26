// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oneof-and-maps/proto/contact.proto

package contact

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Contact struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ContactNumber:
	//	*Contact_Mobile
	//	*Contact_Landline
	ContactNumber        isContact_ContactNumber `protobuf_oneof:"contact_number"`
	Email                []string                `protobuf:"bytes,4,rep,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_d24fe2eaa4cbda67, []int{0}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isContact_ContactNumber interface {
	isContact_ContactNumber()
}

type Contact_Mobile struct {
	Mobile int32 `protobuf:"varint,2,opt,name=mobile,proto3,oneof"`
}

type Contact_Landline struct {
	Landline int32 `protobuf:"varint,3,opt,name=landline,proto3,oneof"`
}

func (*Contact_Mobile) isContact_ContactNumber() {}

func (*Contact_Landline) isContact_ContactNumber() {}

func (m *Contact) GetContactNumber() isContact_ContactNumber {
	if m != nil {
		return m.ContactNumber
	}
	return nil
}

func (m *Contact) GetMobile() int32 {
	if x, ok := m.GetContactNumber().(*Contact_Mobile); ok {
		return x.Mobile
	}
	return 0
}

func (m *Contact) GetLandline() int32 {
	if x, ok := m.GetContactNumber().(*Contact_Landline); ok {
		return x.Landline
	}
	return 0
}

func (m *Contact) GetEmail() []string {
	if m != nil {
		return m.Email
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Contact) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Contact_OneofMarshaler, _Contact_OneofUnmarshaler, _Contact_OneofSizer, []interface{}{
		(*Contact_Mobile)(nil),
		(*Contact_Landline)(nil),
	}
}

func _Contact_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Contact)
	// contact_number
	switch x := m.ContactNumber.(type) {
	case *Contact_Mobile:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Mobile))
	case *Contact_Landline:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Landline))
	case nil:
	default:
		return fmt.Errorf("Contact.ContactNumber has unexpected type %T", x)
	}
	return nil
}

func _Contact_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Contact)
	switch tag {
	case 2: // contact_number.mobile
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ContactNumber = &Contact_Mobile{int32(x)}
		return true, err
	case 3: // contact_number.landline
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ContactNumber = &Contact_Landline{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Contact_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Contact)
	// contact_number
	switch x := m.ContactNumber.(type) {
	case *Contact_Mobile:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Mobile))
	case *Contact_Landline:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Landline))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Dictinory struct {
	//fields are second and nano second (UTC)
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	//just like timestamp
	Duration *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	//Note:- Repeted in map is not valid in proto buf
	Dictinory            map[string]*Contact `protobuf:"bytes,3,rep,name=Dictinory,proto3" json:"Dictinory,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Dictinory) Reset()         { *m = Dictinory{} }
func (m *Dictinory) String() string { return proto.CompactTextString(m) }
func (*Dictinory) ProtoMessage()    {}
func (*Dictinory) Descriptor() ([]byte, []int) {
	return fileDescriptor_d24fe2eaa4cbda67, []int{1}
}

func (m *Dictinory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dictinory.Unmarshal(m, b)
}
func (m *Dictinory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dictinory.Marshal(b, m, deterministic)
}
func (m *Dictinory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dictinory.Merge(m, src)
}
func (m *Dictinory) XXX_Size() int {
	return xxx_messageInfo_Dictinory.Size(m)
}
func (m *Dictinory) XXX_DiscardUnknown() {
	xxx_messageInfo_Dictinory.DiscardUnknown(m)
}

var xxx_messageInfo_Dictinory proto.InternalMessageInfo

func (m *Dictinory) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Dictinory) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Dictinory) GetDictinory() map[string]*Contact {
	if m != nil {
		return m.Dictinory
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*Dictinory)(nil), "Dictinory")
	proto.RegisterMapType((map[string]*Contact)(nil), "Dictinory.DictinoryEntry")
}

func init() { proto.RegisterFile("oneof-and-maps/proto/contact.proto", fileDescriptor_d24fe2eaa4cbda67) }

var fileDescriptor_d24fe2eaa4cbda67 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x34, 0x6d, 0xdf, 0x66, 0x0a, 0xa5, 0x2c, 0x1e, 0xd6, 0x22, 0x35, 0xf4, 0xd4,
	0x4b, 0x13, 0xa8, 0x88, 0xc5, 0xa3, 0x56, 0xf1, 0xbc, 0x78, 0x97, 0x4d, 0xb2, 0x2d, 0x8b, 0xfb,
	0x11, 0xd2, 0x8d, 0x90, 0x83, 0x7f, 0xba, 0x20, 0xee, 0x6e, 0x12, 0x3f, 0x6e, 0x33, 0xf3, 0xfc,
	0x66, 0xf7, 0x79, 0x06, 0x56, 0x5a, 0x31, 0x7d, 0xd8, 0x50, 0x55, 0x6c, 0x24, 0x2d, 0x4f, 0x69,
	0x59, 0x69, 0xa3, 0xd3, 0x5c, 0x2b, 0x43, 0x73, 0x93, 0xd8, 0x6e, 0x71, 0x79, 0xd4, 0xfa, 0x28,
	0x98, 0xd3, 0xb2, 0xfa, 0x90, 0x1a, 0x2e, 0xd9, 0xc9, 0x50, 0x59, 0x7a, 0x60, 0xf9, 0x1b, 0x28,
	0xea, 0x8a, 0x1a, 0xae, 0x95, 0xd3, 0x57, 0xef, 0xf0, 0xff, 0xde, 0xbd, 0x88, 0x10, 0x0c, 0x15,
	0x95, 0x0c, 0x07, 0x71, 0xb0, 0x8e, 0x88, 0xad, 0x11, 0x86, 0xb1, 0xd4, 0x19, 0x17, 0x0c, 0x0f,
	0xe2, 0x60, 0x3d, 0x7a, 0xfa, 0x47, 0x7c, 0x8f, 0x2e, 0x60, 0x22, 0xa8, 0x2a, 0x04, 0x57, 0x0c,
	0x87, 0x5e, 0xeb, 0x26, 0xe8, 0x0c, 0x46, 0x4c, 0x52, 0x2e, 0xf0, 0x30, 0x0e, 0xd7, 0x11, 0x71,
	0xcd, 0xdd, 0x1c, 0x66, 0xde, 0xfe, 0x8b, 0xaa, 0x65, 0xc6, 0xaa, 0xd5, 0x47, 0x00, 0xd1, 0x9e,
	0xe7, 0x86, 0x2b, 0x5d, 0x35, 0x68, 0x07, 0x51, 0xe7, 0xdf, 0xda, 0x98, 0x6e, 0x17, 0x89, 0x0b,
	0x90, 0xb4, 0x01, 0x92, 0xe7, 0x96, 0x20, 0x3d, 0x8c, 0xae, 0x61, 0xd2, 0x06, 0xb3, 0x4e, 0xa7,
	0xdb, 0xf3, 0x3f, 0x8b, 0x7b, 0x0f, 0x90, 0x0e, 0x45, 0x37, 0xdf, 0x7e, 0xc7, 0x61, 0x1c, 0xda,
	0xbd, 0x6e, 0xd2, 0x57, 0x0f, 0xca, 0x54, 0x0d, 0xe9, 0xd9, 0xc5, 0x23, 0xcc, 0x7e, 0x8a, 0x68,
	0x0e, 0xe1, 0x2b, 0x6b, 0xfc, 0xf1, 0xbe, 0x4a, 0xb4, 0x84, 0xd1, 0x1b, 0x15, 0x35, 0xf3, 0x86,
	0x26, 0x89, 0x3f, 0x34, 0x71, 0xe3, 0xdb, 0xc1, 0x2e, 0xc8, 0xc6, 0xd6, 0xdd, 0xd5, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x46, 0xe0, 0x8a, 0x8b, 0xec, 0x01, 0x00, 0x00,
}