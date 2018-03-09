// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/logging/audit_data.proto

/*
Package logging is a generated protocol buffer package.

It is generated from these files:
	google/iam/v1/logging/audit_data.proto

It has these top-level messages:
	AuditData
*/
package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_iam_v1 "google.golang.org/genproto/googleapis/iam/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Audit log information specific to Cloud IAM. This message is serialized
// as an `Any` type in the `ServiceData` message of an
// `AuditLog` message.
type AuditData struct {
	// Policy delta between the original policy and the newly set policy.
	PolicyDelta *google_iam_v1.PolicyDelta `protobuf:"bytes,2,opt,name=policy_delta,json=policyDelta" json:"policy_delta,omitempty"`
}

func (m *AuditData) Reset()                    { *m = AuditData{} }
func (m *AuditData) String() string            { return proto.CompactTextString(m) }
func (*AuditData) ProtoMessage()               {}
func (*AuditData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuditData) GetPolicyDelta() *google_iam_v1.PolicyDelta {
	if m != nil {
		return m.PolicyDelta
	}
	return nil
}

func init() {
	proto.RegisterType((*AuditData)(nil), "google.iam.v1.logging.AuditData")
}

func init() { proto.RegisterFile("google/iam/v1/logging/audit_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0x2b, 0x04, 0x73, 0x62, 0x71, 0x20, 0x68, 0xb4, 0x10, 0x0b, 0xb1, 0x9a, 0xb0,
	0x5a, 0xaa, 0x85, 0xe7, 0x81, 0x28, 0x16, 0x8b, 0x85, 0x85, 0xcd, 0x31, 0x5e, 0x96, 0x61, 0x20,
	0xc9, 0x84, 0xbb, 0xdc, 0x82, 0x8f, 0xe0, 0xab, 0xf8, 0x94, 0xb2, 0x9b, 0xa0, 0xac, 0x58, 0x85,
	0xf0, 0x7f, 0xff, 0x7c, 0xc3, 0xa8, 0x73, 0x12, 0x21, 0xd7, 0x1a, 0x46, 0x6f, 0xba, 0xda, 0x38,
	0x21, 0xe2, 0x40, 0x06, 0xb7, 0x96, 0xd3, 0xd2, 0x62, 0x42, 0x88, 0x6b, 0x49, 0x32, 0x3b, 0xc8,
	0x1c, 0x30, 0x7a, 0xe8, 0x6a, 0x28, 0x9c, 0x3e, 0x29, 0x75, 0x8c, 0x6c, 0x30, 0x04, 0x49, 0x98,
	0x58, 0xc2, 0x26, 0x97, 0xb4, 0x1e, 0x0f, 0x8f, 0xe2, 0x78, 0xf5, 0x91, 0xb3, 0xb3, 0x27, 0xb5,
	0x7b, 0xd7, 0x4b, 0x16, 0x98, 0x70, 0x76, 0xab, 0xf6, 0x72, 0xb8, 0xb4, 0xad, 0x4b, 0x78, 0x38,
	0x39, 0xad, 0x2e, 0xa6, 0x97, 0x1a, 0xc6, 0xd2, 0x66, 0x40, 0x16, 0x3d, 0xf1, 0x32, 0x8d, 0xbf,
	0x9f, 0xf9, 0x67, 0xa5, 0x8e, 0x56, 0xe2, 0xe1, 0xdf, 0x1d, 0xe7, 0xfb, 0x3f, 0x9e, 0xa6, 0x37,
	0x37, 0xd5, 0xdb, 0x4d, 0x01, 0x49, 0x1c, 0x06, 0x02, 0x59, 0x93, 0xa1, 0x36, 0x0c, 0x7b, 0x99,
	0x1c, 0x61, 0xe4, 0xcd, 0x9f, 0x9b, 0x5c, 0x97, 0xf7, 0x6b, 0x72, 0xfc, 0x90, 0xeb, 0xf7, 0x4e,
	0xb6, 0x16, 0x1e, 0xd1, 0xc3, 0x6b, 0x0d, 0xcf, 0x39, 0x7d, 0xdf, 0x19, 0xc6, 0x5c, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x29, 0xf1, 0xcb, 0x3a, 0x59, 0x01, 0x00, 0x00,
}
