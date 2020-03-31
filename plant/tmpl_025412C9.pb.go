// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmpl_025412C9.proto

/*
Package templates is a generated protocol buffer package.

It is generated from these files:
	tmpl_025412C9.proto

It has these top-level messages:
	Tmpl_025412C9
*/
package templates

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

// Checklist plant
type Tmpl_025412C9 struct {
	Symbol                   string `protobuf:"bytes,1,opt,name=Symbol" json:"Symbol,omitempty"`
	SynonymSymbol            string `protobuf:"bytes,2,opt,name=SynonymSymbol" json:"SynonymSymbol,omitempty"`
	ScientificNameWithAuthor string `protobuf:"bytes,3,opt,name=ScientificNameWithAuthor" json:"ScientificNameWithAuthor,omitempty"`
	CommonName               string `protobuf:"bytes,4,opt,name=CommonName" json:"CommonName,omitempty"`
	Family                   string `protobuf:"bytes,5,opt,name=Family" json:"Family,omitempty"`
}

func (m *Tmpl_025412C9) Reset()                    { *m = Tmpl_025412C9{} }
func (m *Tmpl_025412C9) String() string            { return proto.CompactTextString(m) }
func (*Tmpl_025412C9) ProtoMessage()               {}
func (*Tmpl_025412C9) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tmpl_025412C9) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Tmpl_025412C9) GetSynonymSymbol() string {
	if m != nil {
		return m.SynonymSymbol
	}
	return ""
}

func (m *Tmpl_025412C9) GetScientificNameWithAuthor() string {
	if m != nil {
		return m.ScientificNameWithAuthor
	}
	return ""
}

func (m *Tmpl_025412C9) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *Tmpl_025412C9) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func init() {
	proto.RegisterType((*Tmpl_025412C9)(nil), "oipProto.templates.tmpl_025412C9")
}

func init() { proto.RegisterFile("tmpl_025412C9.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x2d, 0xc8,
	0x89, 0x37, 0x30, 0x32, 0x35, 0x31, 0x34, 0x72, 0xb6, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xca, 0xcf, 0x2c, 0x08, 0x00, 0xb1, 0xf4, 0x4a, 0x52, 0x73, 0x0b, 0x72, 0x12, 0x4b, 0x52,
	0x8b, 0x95, 0x0e, 0x32, 0x72, 0xf1, 0xa2, 0xa8, 0x15, 0x12, 0xe3, 0x62, 0x0b, 0xae, 0xcc, 0x4d,
	0xca, 0xcf, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x54, 0xb8, 0x78, 0x83,
	0x2b, 0xf3, 0xf2, 0xf3, 0x2a, 0x73, 0xa1, 0xd2, 0x4c, 0x60, 0x69, 0x54, 0x41, 0x21, 0x2b, 0x2e,
	0x89, 0xe0, 0xe4, 0xcc, 0xd4, 0xbc, 0x92, 0xcc, 0xb4, 0xcc, 0x64, 0xbf, 0xc4, 0xdc, 0xd4, 0xf0,
	0xcc, 0x92, 0x0c, 0xc7, 0xd2, 0x92, 0x8c, 0xfc, 0x22, 0x09, 0x66, 0xb0, 0x06, 0x9c, 0xf2, 0x42,
	0x72, 0x5c, 0x5c, 0xce, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x20, 0x71, 0x09, 0x16, 0xb0, 0x6a, 0x24,
	0x11, 0x90, 0xcb, 0xdc, 0x12, 0x73, 0x33, 0x73, 0x2a, 0x25, 0x58, 0x21, 0x2e, 0x83, 0xf0, 0x9c,
	0xb8, 0xa3, 0x38, 0xe1, 0x1e, 0x4a, 0x62, 0x03, 0xfb, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x69, 0x1e, 0x43, 0x46, 0x02, 0x01, 0x00, 0x00,
}