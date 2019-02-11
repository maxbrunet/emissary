// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/month_of_year.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumerates months of the year, e.g., "January".
type MonthOfYearEnum_MonthOfYear int32

const (
	// Not specified.
	MonthOfYearEnum_UNSPECIFIED MonthOfYearEnum_MonthOfYear = 0
	// The value is unknown in this version.
	MonthOfYearEnum_UNKNOWN MonthOfYearEnum_MonthOfYear = 1
	// January.
	MonthOfYearEnum_JANUARY MonthOfYearEnum_MonthOfYear = 2
	// February.
	MonthOfYearEnum_FEBRUARY MonthOfYearEnum_MonthOfYear = 3
	// March.
	MonthOfYearEnum_MARCH MonthOfYearEnum_MonthOfYear = 4
	// April.
	MonthOfYearEnum_APRIL MonthOfYearEnum_MonthOfYear = 5
	// May.
	MonthOfYearEnum_MAY MonthOfYearEnum_MonthOfYear = 6
	// June.
	MonthOfYearEnum_JUNE MonthOfYearEnum_MonthOfYear = 7
	// July.
	MonthOfYearEnum_JULY MonthOfYearEnum_MonthOfYear = 8
	// August.
	MonthOfYearEnum_AUGUST MonthOfYearEnum_MonthOfYear = 9
	// September.
	MonthOfYearEnum_SEPTEMBER MonthOfYearEnum_MonthOfYear = 10
	// October.
	MonthOfYearEnum_OCTOBER MonthOfYearEnum_MonthOfYear = 11
	// November.
	MonthOfYearEnum_NOVEMBER MonthOfYearEnum_MonthOfYear = 12
	// December.
	MonthOfYearEnum_DECEMBER MonthOfYearEnum_MonthOfYear = 13
)

var MonthOfYearEnum_MonthOfYear_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "JANUARY",
	3:  "FEBRUARY",
	4:  "MARCH",
	5:  "APRIL",
	6:  "MAY",
	7:  "JUNE",
	8:  "JULY",
	9:  "AUGUST",
	10: "SEPTEMBER",
	11: "OCTOBER",
	12: "NOVEMBER",
	13: "DECEMBER",
}
var MonthOfYearEnum_MonthOfYear_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"JANUARY":     2,
	"FEBRUARY":    3,
	"MARCH":       4,
	"APRIL":       5,
	"MAY":         6,
	"JUNE":        7,
	"JULY":        8,
	"AUGUST":      9,
	"SEPTEMBER":   10,
	"OCTOBER":     11,
	"NOVEMBER":    12,
	"DECEMBER":    13,
}

func (x MonthOfYearEnum_MonthOfYear) String() string {
	return proto.EnumName(MonthOfYearEnum_MonthOfYear_name, int32(x))
}
func (MonthOfYearEnum_MonthOfYear) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_month_of_year_f94758a12a1b8695, []int{0, 0}
}

// Container for enumeration of months of the year, e.g., "January".
type MonthOfYearEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonthOfYearEnum) Reset()         { *m = MonthOfYearEnum{} }
func (m *MonthOfYearEnum) String() string { return proto.CompactTextString(m) }
func (*MonthOfYearEnum) ProtoMessage()    {}
func (*MonthOfYearEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_month_of_year_f94758a12a1b8695, []int{0}
}
func (m *MonthOfYearEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonthOfYearEnum.Unmarshal(m, b)
}
func (m *MonthOfYearEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonthOfYearEnum.Marshal(b, m, deterministic)
}
func (dst *MonthOfYearEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonthOfYearEnum.Merge(dst, src)
}
func (m *MonthOfYearEnum) XXX_Size() int {
	return xxx_messageInfo_MonthOfYearEnum.Size(m)
}
func (m *MonthOfYearEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MonthOfYearEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MonthOfYearEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MonthOfYearEnum)(nil), "google.ads.googleads.v0.enums.MonthOfYearEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.MonthOfYearEnum_MonthOfYear", MonthOfYearEnum_MonthOfYear_name, MonthOfYearEnum_MonthOfYear_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/month_of_year.proto", fileDescriptor_month_of_year_f94758a12a1b8695)
}

var fileDescriptor_month_of_year_f94758a12a1b8695 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0xb5, 0xe5, 0xfb, 0xf8, 0x99, 0x42, 0x98, 0xcc, 0x9e, 0x05, 0x3c, 0xc0, 0xb4, 0xc6, 0xdd,
	0xb8, 0x9a, 0x96, 0x01, 0x41, 0xfa, 0x93, 0x42, 0x6b, 0x6a, 0x9a, 0x90, 0x6a, 0x4b, 0x35, 0xa1,
	0x1d, 0xd2, 0x01, 0x12, 0x5f, 0xc7, 0xa5, 0xef, 0xe0, 0x0b, 0xb8, 0xf3, 0x21, 0xdc, 0xf8, 0x14,
	0x66, 0x3a, 0x42, 0xd8, 0xe8, 0xa6, 0x39, 0xe7, 0x9e, 0x73, 0x7a, 0xe7, 0x1e, 0xf0, 0xb2, 0xe4,
	0xbc, 0x3c, 0x16, 0x66, 0x96, 0x0b, 0x53, 0x41, 0x89, 0xae, 0x96, 0x59, 0xd4, 0x97, 0x4a, 0x98,
	0x15, 0xaf, 0xcf, 0x9f, 0xf6, 0xfc, 0xb0, 0xff, 0x52, 0x64, 0x0d, 0x3e, 0x35, 0xfc, 0xcc, 0xd1,
	0x44, 0xf9, 0x70, 0x96, 0x0b, 0x7c, 0x8f, 0xe0, 0xab, 0x85, 0xdb, 0xc8, 0xec, 0xa7, 0x06, 0xc6,
	0xae, 0x8c, 0xf9, 0x87, 0xa4, 0xc8, 0x1a, 0x56, 0x5f, 0xaa, 0xd9, 0x77, 0x0d, 0x18, 0x0f, 0x33,
	0x34, 0x06, 0x46, 0xe4, 0x6d, 0x03, 0xe6, 0xac, 0x16, 0x2b, 0x36, 0x87, 0x2f, 0x90, 0x01, 0x7a,
	0x91, 0xf7, 0xd6, 0xf3, 0xdf, 0x79, 0x50, 0x93, 0x64, 0x4d, 0xbd, 0x88, 0x86, 0x09, 0xd4, 0xd1,
	0x10, 0xf4, 0x17, 0xcc, 0x0e, 0x5b, 0xd6, 0x41, 0x03, 0xf0, 0xec, 0xd2, 0xd0, 0x79, 0x03, 0x9f,
	0x24, 0xa4, 0x41, 0xb8, 0xda, 0xc0, 0x67, 0xd4, 0x03, 0x1d, 0x97, 0x26, 0xb0, 0x8b, 0xfa, 0xe0,
	0x69, 0x1d, 0x79, 0x0c, 0xf6, 0x14, 0xda, 0x24, 0xb0, 0x8f, 0x00, 0xe8, 0xd2, 0x68, 0x19, 0x6d,
	0x77, 0x70, 0x80, 0x46, 0x60, 0xb0, 0x65, 0xc1, 0x8e, 0xb9, 0x36, 0x0b, 0x21, 0x90, 0x8b, 0x7c,
	0x67, 0xe7, 0x4b, 0x62, 0xc8, 0x45, 0x9e, 0x1f, 0x2b, 0x69, 0x28, 0xd9, 0x9c, 0x39, 0x8a, 0x8d,
	0xec, 0x5f, 0x1a, 0x98, 0x7e, 0xe4, 0x15, 0xfe, 0xef, 0xe5, 0x36, 0x7c, 0x38, 0x31, 0x90, 0x55,
	0x05, 0xda, 0x7b, 0xfb, 0x6f, 0xa4, 0xe4, 0xc7, 0xac, 0x2e, 0x31, 0x6f, 0x4a, 0xb3, 0x2c, 0xea,
	0xb6, 0xc8, 0x5b, 0xdf, 0xa7, 0xcf, 0xe2, 0x1f, 0xf5, 0xbf, 0x6e, 0xbf, 0x5f, 0xf5, 0xce, 0x92,
	0xd2, 0x6f, 0xfa, 0x64, 0xa9, 0x7e, 0x45, 0x73, 0x81, 0x15, 0x94, 0x28, 0xb6, 0xb0, 0xac, 0x58,
	0xfc, 0xb8, 0xe9, 0x29, 0xcd, 0x45, 0x7a, 0xd7, 0xd3, 0xd8, 0x4a, 0x5b, 0xfd, 0xb7, 0x3e, 0x55,
	0x43, 0x42, 0x68, 0x2e, 0x08, 0xb9, 0x3b, 0x08, 0x89, 0x2d, 0x42, 0x5a, 0xcf, 0x87, 0x6e, 0xfb,
	0xb0, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x40, 0x36, 0x45, 0x16, 0x02, 0x00, 0x00,
}
