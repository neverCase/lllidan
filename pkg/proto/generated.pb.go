/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/nevercase/lllidan/pkg/proto/generated.proto

package proto

import (
	fmt "fmt"

	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Gateway) Reset()      { *m = Gateway{} }
func (*Gateway) ProtoMessage() {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_52097ae5c98ea68f, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return m.Size()
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *GatewayList) Reset()      { *m = GatewayList{} }
func (*GatewayList) ProtoMessage() {}
func (*GatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_52097ae5c98ea68f, []int{1}
}
func (m *GatewayList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayList.Merge(m, src)
}
func (m *GatewayList) XXX_Size() int {
	return m.Size()
}
func (m *GatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayList proto.InternalMessageInfo

func (m *Request) Reset()      { *m = Request{} }
func (*Request) ProtoMessage() {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_52097ae5c98ea68f, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Response) Reset()      { *m = Response{} }
func (*Response) ProtoMessage() {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_52097ae5c98ea68f, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Gateway)(nil), "github.com.nevercase.lllidan.pkg.proto.Gateway")
	proto.RegisterType((*GatewayList)(nil), "github.com.nevercase.lllidan.pkg.proto.GatewayList")
	proto.RegisterType((*Request)(nil), "github.com.nevercase.lllidan.pkg.proto.Request")
	proto.RegisterType((*Response)(nil), "github.com.nevercase.lllidan.pkg.proto.Response")
}

func init() {
	proto.RegisterFile("github.com/nevercase/lllidan/pkg/proto/generated.proto", fileDescriptor_52097ae5c98ea68f)
}

var fileDescriptor_52097ae5c98ea68f = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x86, 0xe3, 0x5c, 0xa2, 0x5c, 0x7d, 0x87, 0x90, 0x32, 0x45, 0x27, 0xe4, 0x8b, 0x32, 0xa0,
	0x22, 0x44, 0x22, 0x31, 0x30, 0x43, 0x40, 0x42, 0x95, 0x40, 0x3a, 0x19, 0x26, 0x36, 0x5f, 0xfc,
	0x61, 0xac, 0x6b, 0x63, 0x13, 0xbb, 0x45, 0x6c, 0xfc, 0x04, 0xfe, 0x0d, 0x7f, 0xa1, 0x63, 0xc7,
	0x4e, 0x15, 0x0d, 0xff, 0x82, 0x09, 0xd5, 0x31, 0x25, 0x48, 0x0c, 0x0c, 0x37, 0x7e, 0x7e, 0xbe,
	0xf7, 0x7d, 0xfd, 0x7e, 0xf8, 0x89, 0x90, 0xf6, 0xc3, 0xf2, 0xba, 0x6c, 0xd4, 0xa2, 0x6a, 0x61,
	0x05, 0x5d, 0xc3, 0x0c, 0x54, 0xf3, 0xf9, 0x5c, 0x72, 0xd6, 0x56, 0xfa, 0x46, 0x54, 0xba, 0x53,
	0x56, 0x55, 0x02, 0x5a, 0xe8, 0x98, 0x05, 0x5e, 0xba, 0x39, 0xbd, 0xff, 0x47, 0x57, 0x1e, 0x75,
	0xa5, 0xd7, 0x95, 0xfa, 0x46, 0x0c, 0x7b, 0x17, 0x8f, 0x46, 0xfe, 0x42, 0x09, 0x35, 0xd8, 0x5d,
	0x2f, 0xdf, 0xbb, 0xc9, 0x7b, 0x2b, 0xa1, 0x86, 0xf5, 0xa2, 0xc1, 0xc9, 0x4b, 0x66, 0xe1, 0x13,
	0xfb, 0x9c, 0x5e, 0xe0, 0x50, 0xf2, 0x0c, 0xe5, 0x68, 0x1a, 0xd7, 0x78, 0xbd, 0xbb, 0x0c, 0xfa,
	0xdd, 0x65, 0x38, 0xe3, 0x34, 0x94, 0xdc, 0x31, 0x9d, 0x85, 0x39, 0x9a, 0x4e, 0x46, 0x4c, 0xd3,
	0x50, 0xea, 0x34, 0xc7, 0x91, 0x56, 0x9d, 0xcd, 0x4e, 0x9c, 0xf2, 0xdc, 0xd3, 0xe8, 0x4a, 0x75,
	0x96, 0x3a, 0x52, 0x34, 0xf8, 0xcc, 0x87, 0xbc, 0x92, 0xc6, 0xa6, 0x6f, 0x71, 0x2c, 0x2d, 0x2c,
	0x4c, 0x86, 0xf2, 0x93, 0xe9, 0xd9, 0xe3, 0xaa, 0xfc, 0xbf, 0x6a, 0xa5, 0xf7, 0xa8, 0xef, 0xf8,
	0x88, 0x78, 0x76, 0x70, 0xa1, 0x83, 0x59, 0x21, 0x71, 0x42, 0xe1, 0xe3, 0x12, 0x8c, 0x4d, 0x9f,
	0x62, 0x6c, 0xa0, 0x5b, 0xc9, 0x06, 0x9e, 0x69, 0xe9, 0x1a, 0x4d, 0xea, 0xdc, 0x8b, 0xf0, 0x1b,
	0x4f, 0xae, 0x66, 0x3f, 0xff, 0x9a, 0xe8, 0x48, 0x93, 0xde, 0xc3, 0x11, 0x67, 0x96, 0xb9, 0xc6,
	0xe7, 0xf5, 0xe9, 0xa1, 0xcf, 0x0b, 0x66, 0x19, 0x75, 0xaf, 0xc5, 0x37, 0x84, 0x4f, 0x29, 0x18,
	0xad, 0x5a, 0x03, 0xb7, 0x10, 0x96, 0xe3, 0xa8, 0x51, 0x1c, 0x5c, 0xd8, 0xe8, 0x80, 0xcf, 0x15,
	0x07, 0xea, 0x48, 0xfa, 0x00, 0x27, 0x0b, 0x30, 0x86, 0x09, 0x70, 0x57, 0x9e, 0xd4, 0x77, 0xfd,
	0x52, 0xf2, 0x7a, 0x78, 0xa6, 0xbf, 0xf9, 0xf1, 0xe7, 0xd1, 0xbf, 0x7e, 0x5e, 0x3f, 0x5c, 0xef,
	0x49, 0xb0, 0xd9, 0x93, 0x60, 0xbb, 0x27, 0xc1, 0x97, 0x9e, 0xa0, 0x75, 0x4f, 0xd0, 0xa6, 0x27,
	0x68, 0xdb, 0x13, 0xf4, 0xbd, 0x27, 0xe8, 0xeb, 0x0f, 0x12, 0xbc, 0x8b, 0xdd, 0xbd, 0x7f, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x34, 0xf8, 0x09, 0x14, 0xab, 0x02, 0x00, 0x00,
}

func (m *Gateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gateway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gateway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintGenerated(dAtA, i, uint64(m.Port))
	i--
	dAtA[i] = 0x18
	i -= len(m.Ip)
	copy(dAtA[i:], m.Ip)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Ip)))
	i--
	dAtA[i] = 0x12
	i = encodeVarintGenerated(dAtA, i, uint64(m.Id))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *GatewayList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GatewayList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	i -= len(m.ServiceAPI)
	copy(dAtA[i:], m.ServiceAPI)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ServiceAPI)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	i -= len(m.Message)
	copy(dAtA[i:], m.Message)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Message)))
	i--
	dAtA[i] = 0x1a
	i = encodeVarintGenerated(dAtA, i, uint64(m.Code))
	i--
	dAtA[i] = 0x10
	i -= len(m.ServiceAPI)
	copy(dAtA[i:], m.ServiceAPI)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ServiceAPI)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Gateway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Id))
	l = len(m.Ip)
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Port))
	return n
}

func (m *GatewayList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceAPI)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceAPI)
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Code))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Gateway) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Gateway{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Ip:` + fmt.Sprintf("%v", this.Ip) + `,`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GatewayList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Gateway{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Gateway", "Gateway", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&GatewayList{`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Request{`,
		`ServiceAPI:` + fmt.Sprintf("%v", this.ServiceAPI) + `,`,
		`Data:` + valueToStringGenerated(this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`ServiceAPI:` + fmt.Sprintf("%v", this.ServiceAPI) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Data:` + valueToStringGenerated(this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Gateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Gateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *GatewayList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: GatewayList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Gateway{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAPI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAPI = ServiceAPI(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
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
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAPI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAPI = ServiceAPI(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
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
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
				}
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)
