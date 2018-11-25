// Code generated by protoc-gen-go. DO NOT EDIT.
// source: readwritefileusingproto/proto/userdata.proto

package userdata

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserData struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32               `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	PhoneNumber          []string            `protobuf:"bytes,3,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Birthday             *Date               `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	UserAddress          []*UserData_Address `protobuf:"bytes,5,rep,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4f88cece6fedf0, []int{0}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserData) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserData) GetPhoneNumber() []string {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *UserData) GetBirthday() *Date {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *UserData) GetUserAddress() []*UserData_Address {
	if m != nil {
		return m.UserAddress
	}
	return nil
}

type UserData_Address struct {
	AddressLine1         string   `protobuf:"bytes,1,opt,name=address_line1,json=addressLine1,proto3" json:"address_line1,omitempty"`
	AddressLine2         string   `protobuf:"bytes,2,opt,name=address_line2,json=addressLine2,proto3" json:"address_line2,omitempty"`
	AddressLine3         string   `protobuf:"bytes,3,opt,name=address_line3,json=addressLine3,proto3" json:"address_line3,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData_Address) Reset()         { *m = UserData_Address{} }
func (m *UserData_Address) String() string { return proto.CompactTextString(m) }
func (*UserData_Address) ProtoMessage()    {}
func (*UserData_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4f88cece6fedf0, []int{0, 0}
}

func (m *UserData_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData_Address.Unmarshal(m, b)
}
func (m *UserData_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData_Address.Marshal(b, m, deterministic)
}
func (m *UserData_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData_Address.Merge(m, src)
}
func (m *UserData_Address) XXX_Size() int {
	return xxx_messageInfo_UserData_Address.Size(m)
}
func (m *UserData_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData_Address.DiscardUnknown(m)
}

var xxx_messageInfo_UserData_Address proto.InternalMessageInfo

func (m *UserData_Address) GetAddressLine1() string {
	if m != nil {
		return m.AddressLine1
	}
	return ""
}

func (m *UserData_Address) GetAddressLine2() string {
	if m != nil {
		return m.AddressLine2
	}
	return ""
}

func (m *UserData_Address) GetAddressLine3() string {
	if m != nil {
		return m.AddressLine3
	}
	return ""
}

func (m *UserData_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserData_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type Date struct {
	Day                  int32    `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month                int32    `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year                 int32    `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4f88cece6fedf0, []int{1}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func init() {
	proto.RegisterType((*UserData)(nil), "UserData")
	proto.RegisterType((*UserData_Address)(nil), "UserData.Address")
	proto.RegisterType((*Date)(nil), "Date")
}

func init() {
	proto.RegisterFile("readwritefileusingproto/proto/userdata.proto", fileDescriptor_9d4f88cece6fedf0)
}

var fileDescriptor_9d4f88cece6fedf0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x25, 0xeb, 0xe2, 0xb6, 0xaf, 0x13, 0x34, 0x78, 0x08, 0x9e, 0xba, 0x79, 0xe9, 0x41, 0x2a,
	0x76, 0xfe, 0x01, 0x65, 0x47, 0xf1, 0x10, 0xf0, 0x5c, 0xd2, 0xf5, 0x73, 0x0d, 0xac, 0xe9, 0x48,
	0x53, 0xa4, 0x7f, 0xc8, 0x7f, 0xe4, 0xff, 0x91, 0xa4, 0x99, 0xe0, 0xbc, 0x84, 0xf7, 0x5e, 0x5e,
	0x92, 0xef, 0xbd, 0xc0, 0xbd, 0x41, 0x59, 0x7d, 0x1a, 0x65, 0xf1, 0x43, 0x1d, 0xb0, 0xef, 0x94,
	0xde, 0x1f, 0x4d, 0x6b, 0xdb, 0x87, 0x71, 0xed, 0x3b, 0x34, 0x95, 0xb4, 0x32, 0xf3, 0x74, 0xfd,
	0x3d, 0x81, 0xf9, 0x7b, 0x87, 0x66, 0x2b, 0xad, 0x64, 0x0c, 0xa6, 0x5a, 0x36, 0xc8, 0x49, 0x42,
	0xd2, 0x85, 0xf0, 0x98, 0x5d, 0x41, 0x24, 0xf7, 0xc8, 0x27, 0x09, 0x49, 0xa9, 0x70, 0x90, 0xad,
	0x60, 0x79, 0xac, 0x5b, 0x8d, 0x85, 0xee, 0x9b, 0x12, 0x0d, 0x8f, 0x92, 0x28, 0x5d, 0x88, 0xd8,
	0x6b, 0x6f, 0x5e, 0x62, 0x2b, 0x98, 0x97, 0xca, 0xd8, 0xba, 0x92, 0x03, 0x9f, 0x26, 0x24, 0x8d,
	0x73, 0x9a, 0x6d, 0xa5, 0x45, 0xf1, 0x2b, 0xb3, 0x27, 0x58, 0xba, 0x51, 0x0a, 0x59, 0x55, 0x06,
	0xbb, 0x8e, 0xd3, 0x24, 0x4a, 0xe3, 0xfc, 0x3a, 0x3b, 0x0d, 0x93, 0x3d, 0x8f, 0x1b, 0x22, 0x76,
	0xb6, 0x40, 0x6e, 0xbf, 0x08, 0xcc, 0x02, 0x66, 0x77, 0x70, 0x19, 0x0e, 0x17, 0x07, 0xa5, 0xf1,
	0x31, 0x8c, 0xbd, 0x0c, 0xe2, 0xab, 0xd3, 0xce, 0x4d, 0xb9, 0x0f, 0xf2, 0xd7, 0x94, 0x9f, 0x9b,
	0x36, 0x3c, 0xfa, 0x67, 0xda, 0xb8, 0x72, 0x76, 0xca, 0x8e, 0x79, 0x16, 0xc2, 0x63, 0xc6, 0x61,
	0xb6, 0x6b, 0x7b, 0x6d, 0xcd, 0xc0, 0xa9, 0x97, 0x4f, 0x74, 0xfd, 0x02, 0x53, 0x17, 0xd8, 0xd5,
	0xe7, 0x4a, 0x20, 0x63, 0x7d, 0x2e, 0xf8, 0x0d, 0xd0, 0xa6, 0xd5, 0xb6, 0x0e, 0x95, 0x8e, 0xc4,
	0xdd, 0x3e, 0xa0, 0x34, 0xfe, 0x65, 0x2a, 0x3c, 0x2e, 0x2f, 0xfc, 0x17, 0x6d, 0x7e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x33, 0x8f, 0x2a, 0xd2, 0x01, 0x00, 0x00,
}
