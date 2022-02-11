/*
Copyright The TestGrid Authors.

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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: custom_evaluator.proto

package custom_evaluator

import (
	fmt "fmt"
	test_status "github.com/GoogleCloudPlatform/testgrid/pb/test_status"
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

type Comparison_Operator int32

const (
	// Unknown. May assume OP_EQ for legacy purposes, but should warn.
	Comparison_OP_UNKNOWN Comparison_Operator = 0
	// Equals operator.
	Comparison_OP_EQ Comparison_Operator = 1
	// Not equals operator.
	Comparison_OP_NE Comparison_Operator = 2
	// Comparison value less than TestResult's value
	Comparison_OP_LT Comparison_Operator = 3
	// Comparison value less than or equal TestResult's value
	Comparison_OP_LE Comparison_Operator = 4
	// Comparison value greater than TestResult's value
	Comparison_OP_GT Comparison_Operator = 5
	// Comparison value greater than or equal TestResult's value
	Comparison_OP_GE Comparison_Operator = 6
	// Regex match of Comparison.value string with the TestResult's evaluation
	// value string.
	Comparison_OP_REGEX Comparison_Operator = 7
	// Checks to see if the evaluation value string starts with the
	// Comparison.value string
	Comparison_OP_STARTS_WITH Comparison_Operator = 8
	// Checks to see if the evaluation value string is contained within the
	// Comparison.value string
	Comparison_OP_CONTAINS Comparison_Operator = 9
)

var Comparison_Operator_name = map[int32]string{
	0: "OP_UNKNOWN",
	1: "OP_EQ",
	2: "OP_NE",
	3: "OP_LT",
	4: "OP_LE",
	5: "OP_GT",
	6: "OP_GE",
	7: "OP_REGEX",
	8: "OP_STARTS_WITH",
	9: "OP_CONTAINS",
}

var Comparison_Operator_value = map[string]int32{
	"OP_UNKNOWN":     0,
	"OP_EQ":          1,
	"OP_NE":          2,
	"OP_LT":          3,
	"OP_LE":          4,
	"OP_GT":          5,
	"OP_GE":          6,
	"OP_REGEX":       7,
	"OP_STARTS_WITH": 8,
	"OP_CONTAINS":    9,
}

func (x Comparison_Operator) String() string {
	return proto.EnumName(Comparison_Operator_name, int32(x))
}

func (Comparison_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14164f833d03200a, []int{3, 0}
}

// A collection of Rule objects. Used to define many rules.
type RuleSet struct {
	Rules                []*Rule  `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleSet) Reset()         { *m = RuleSet{} }
func (m *RuleSet) String() string { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()    {}
func (*RuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_14164f833d03200a, []int{0}
}

func (m *RuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleSet.Unmarshal(m, b)
}
func (m *RuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleSet.Marshal(b, m, deterministic)
}
func (m *RuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleSet.Merge(m, src)
}
func (m *RuleSet) XXX_Size() int {
	return xxx_messageInfo_RuleSet.Size(m)
}
func (m *RuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_RuleSet proto.InternalMessageInfo

func (m *RuleSet) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A single rule that describes how to evaluate a test_cases_pb2.TestResult
type Rule struct {
	// Multiple comparisons to run against a result. EVERY TestResultComparison
	// has to succeed for this Rule to succeed.
	TestResultComparisons []*TestResultComparison `protobuf:"bytes,1,rep,name=test_result_comparisons,json=testResultComparisons,proto3" json:"test_result_comparisons,omitempty"`
	// Required: The TestStatus to return if the comparison succeeds.
	ComputedStatus       test_status.TestStatus `protobuf:"varint,3,opt,name=computed_status,json=computedStatus,proto3,enum=TestStatus" json:"computed_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_14164f833d03200a, []int{1}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetTestResultComparisons() []*TestResultComparison {
	if m != nil {
		return m.TestResultComparisons
	}
	return nil
}

func (m *Rule) GetComputedStatus() test_status.TestStatus {
	if m != nil {
		return m.ComputedStatus
	}
	return test_status.TestStatus_NO_RESULT
}

// Describes how to get information the TestResult proto and how to compare the
// value against the comparison value.
type TestResultComparison struct {
	// Required: This is the comparison that will be used as
	Comparison *Comparison `protobuf:"bytes,1,opt,name=comparison,proto3" json:"comparison,omitempty"`
	// Types that are valid to be assigned to TestResultInfo:
	//	*TestResultComparison_PropertyKey
	//	*TestResultComparison_TestResultField
	//	*TestResultComparison_TestResultErrorField
	TestResultInfo       isTestResultComparison_TestResultInfo `protobuf_oneof:"test_result_info"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *TestResultComparison) Reset()         { *m = TestResultComparison{} }
func (m *TestResultComparison) String() string { return proto.CompactTextString(m) }
func (*TestResultComparison) ProtoMessage()    {}
func (*TestResultComparison) Descriptor() ([]byte, []int) {
	return fileDescriptor_14164f833d03200a, []int{2}
}

func (m *TestResultComparison) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultComparison.Unmarshal(m, b)
}
func (m *TestResultComparison) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultComparison.Marshal(b, m, deterministic)
}
func (m *TestResultComparison) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultComparison.Merge(m, src)
}
func (m *TestResultComparison) XXX_Size() int {
	return xxx_messageInfo_TestResultComparison.Size(m)
}
func (m *TestResultComparison) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultComparison.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultComparison proto.InternalMessageInfo

func (m *TestResultComparison) GetComparison() *Comparison {
	if m != nil {
		return m.Comparison
	}
	return nil
}

type isTestResultComparison_TestResultInfo interface {
	isTestResultComparison_TestResultInfo()
}

type TestResultComparison_PropertyKey struct {
	PropertyKey string `protobuf:"bytes,2,opt,name=property_key,json=propertyKey,proto3,oneof"`
}

type TestResultComparison_TestResultField struct {
	TestResultField string `protobuf:"bytes,3,opt,name=test_result_field,json=testResultField,proto3,oneof"`
}

type TestResultComparison_TestResultErrorField struct {
	TestResultErrorField string `protobuf:"bytes,4,opt,name=test_result_error_field,json=testResultErrorField,proto3,oneof"`
}

func (*TestResultComparison_PropertyKey) isTestResultComparison_TestResultInfo() {}

func (*TestResultComparison_TestResultField) isTestResultComparison_TestResultInfo() {}

func (*TestResultComparison_TestResultErrorField) isTestResultComparison_TestResultInfo() {}

func (m *TestResultComparison) GetTestResultInfo() isTestResultComparison_TestResultInfo {
	if m != nil {
		return m.TestResultInfo
	}
	return nil
}

func (m *TestResultComparison) GetPropertyKey() string {
	if x, ok := m.GetTestResultInfo().(*TestResultComparison_PropertyKey); ok {
		return x.PropertyKey
	}
	return ""
}

func (m *TestResultComparison) GetTestResultField() string {
	if x, ok := m.GetTestResultInfo().(*TestResultComparison_TestResultField); ok {
		return x.TestResultField
	}
	return ""
}

func (m *TestResultComparison) GetTestResultErrorField() string {
	if x, ok := m.GetTestResultInfo().(*TestResultComparison_TestResultErrorField); ok {
		return x.TestResultErrorField
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestResultComparison) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestResultComparison_PropertyKey)(nil),
		(*TestResultComparison_TestResultField)(nil),
		(*TestResultComparison_TestResultErrorField)(nil),
	}
}

// The method of comparison used for evaluation. Describes how to compare two
// values.
type Comparison struct {
	// Required: Defines how to compare two attributes.
	// When the TestResult value is numerical, numerical_value will be used to
	// compare. When the TestResult value is a string, string_value will be used.
	Op Comparison_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=Comparison_Operator" json:"op,omitempty"`
	// Types that are valid to be assigned to ComparisonValue:
	//	*Comparison_StringValue
	//	*Comparison_NumericalValue
	ComparisonValue      isComparison_ComparisonValue `protobuf_oneof:"comparison_value"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Comparison) Reset()         { *m = Comparison{} }
func (m *Comparison) String() string { return proto.CompactTextString(m) }
func (*Comparison) ProtoMessage()    {}
func (*Comparison) Descriptor() ([]byte, []int) {
	return fileDescriptor_14164f833d03200a, []int{3}
}

func (m *Comparison) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comparison.Unmarshal(m, b)
}
func (m *Comparison) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comparison.Marshal(b, m, deterministic)
}
func (m *Comparison) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comparison.Merge(m, src)
}
func (m *Comparison) XXX_Size() int {
	return xxx_messageInfo_Comparison.Size(m)
}
func (m *Comparison) XXX_DiscardUnknown() {
	xxx_messageInfo_Comparison.DiscardUnknown(m)
}

var xxx_messageInfo_Comparison proto.InternalMessageInfo

func (m *Comparison) GetOp() Comparison_Operator {
	if m != nil {
		return m.Op
	}
	return Comparison_OP_UNKNOWN
}

type isComparison_ComparisonValue interface {
	isComparison_ComparisonValue()
}

type Comparison_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Comparison_NumericalValue struct {
	NumericalValue float64 `protobuf:"fixed64,3,opt,name=numerical_value,json=numericalValue,proto3,oneof"`
}

func (*Comparison_StringValue) isComparison_ComparisonValue() {}

func (*Comparison_NumericalValue) isComparison_ComparisonValue() {}

func (m *Comparison) GetComparisonValue() isComparison_ComparisonValue {
	if m != nil {
		return m.ComparisonValue
	}
	return nil
}

func (m *Comparison) GetStringValue() string {
	if x, ok := m.GetComparisonValue().(*Comparison_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Comparison) GetNumericalValue() float64 {
	if x, ok := m.GetComparisonValue().(*Comparison_NumericalValue); ok {
		return x.NumericalValue
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Comparison) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Comparison_StringValue)(nil),
		(*Comparison_NumericalValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("Comparison_Operator", Comparison_Operator_name, Comparison_Operator_value)
	proto.RegisterType((*RuleSet)(nil), "RuleSet")
	proto.RegisterType((*Rule)(nil), "Rule")
	proto.RegisterType((*TestResultComparison)(nil), "TestResultComparison")
	proto.RegisterType((*Comparison)(nil), "Comparison")
}

func init() { proto.RegisterFile("custom_evaluator.proto", fileDescriptor_14164f833d03200a) }

var fileDescriptor_14164f833d03200a = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x9d, 0xfb, 0xb3, 0xb5, 0x5f, 0xa6, 0xd4, 0x58, 0x1d, 0x54, 0x70, 0x13, 0x05, 0x84, 0x8a,
	0x40, 0x41, 0x2a, 0x48, 0x5c, 0x8f, 0x29, 0xac, 0xd3, 0x20, 0x09, 0x4e, 0x60, 0xdc, 0x59, 0x59,
	0xe7, 0xa1, 0x88, 0xb4, 0x8e, 0x6c, 0x07, 0xa9, 0xcf, 0xc0, 0x05, 0xaf, 0xc1, 0x7b, 0xf1, 0x22,
	0xc8, 0xf9, 0x59, 0x82, 0xb4, 0xbb, 0xe3, 0x73, 0xbe, 0x63, 0x9f, 0x63, 0x1b, 0x1e, 0x6e, 0x4a,
	0xa5, 0xc5, 0x96, 0xf1, 0x9f, 0x69, 0x5e, 0xa6, 0x5a, 0x48, 0xaf, 0x90, 0x42, 0x8b, 0xc7, 0x4e,
	0x71, 0xfd, 0x5a, 0x73, 0xa5, 0x99, 0xd2, 0xa9, 0x2e, 0x55, 0x1f, 0xd7, 0x13, 0xee, 0x73, 0x38,
	0xa2, 0x65, 0xce, 0x63, 0xae, 0xc9, 0x13, 0x18, 0xcb, 0x32, 0xe7, 0x6a, 0x81, 0x9c, 0xe1, 0xd2,
	0x5a, 0x8d, 0x3d, 0x23, 0xd0, 0x9a, 0x73, 0x7f, 0x21, 0x18, 0x99, 0x35, 0xf9, 0x04, 0x8f, 0xaa,
	0x5d, 0x24, 0x57, 0x65, 0xae, 0xd9, 0x46, 0x6c, 0x8b, 0x54, 0x66, 0x4a, 0xec, 0x5a, 0xdf, 0x89,
	0x97, 0x70, 0xa5, 0x69, 0x25, 0x9f, 0xdd, 0xa9, 0xf4, 0x44, 0xdf, 0xc3, 0x2a, 0xf2, 0x16, 0x66,
	0x66, 0x8b, 0x52, 0xf3, 0x9b, 0x26, 0xd8, 0x62, 0xe8, 0xa0, 0xa5, 0xbd, 0xb2, 0xaa, 0x6d, 0xe2,
	0x8a, 0xa2, 0x76, 0x3b, 0x53, 0xaf, 0xdd, 0xbf, 0x08, 0xe6, 0xf7, 0x9d, 0x42, 0x5e, 0x02, 0x74,
	0x89, 0x16, 0xc8, 0x41, 0x4b, 0x6b, 0x65, 0x79, 0xbd, 0x18, 0x3d, 0x99, 0x3c, 0x85, 0xe3, 0x42,
	0x8a, 0x82, 0x4b, 0xbd, 0x67, 0x3f, 0xf8, 0x7e, 0x31, 0x70, 0xd0, 0x72, 0xba, 0x3e, 0xa0, 0x56,
	0xcb, 0x5e, 0xf2, 0x3d, 0x79, 0x05, 0x0f, 0xfa, 0x7d, 0x6f, 0x33, 0x9e, 0xdf, 0x54, 0x11, 0xcd,
	0xe4, 0xac, 0x2b, 0xf5, 0xc1, 0x08, 0xe4, 0xdd, 0xff, 0xb7, 0xc3, 0xa5, 0x14, 0xb2, 0xf1, 0x8c,
	0x1a, 0xcf, 0xbc, 0xf3, 0xf8, 0x46, 0xae, 0x8c, 0xef, 0x09, 0xe0, 0xbe, 0x31, 0xdb, 0xdd, 0x0a,
	0xf7, 0xcf, 0x00, 0xa0, 0xd7, 0xed, 0x19, 0x0c, 0x44, 0x51, 0x75, 0xb2, 0x57, 0xf3, 0x5e, 0x27,
	0x2f, 0x2c, 0xb8, 0x34, 0x8f, 0x4e, 0x07, 0xa2, 0x30, 0xa5, 0x94, 0x96, 0xd9, 0xee, 0x3b, 0x33,
	0x7f, 0x81, 0x77, 0xa5, 0x6a, 0xf6, 0xab, 0x21, 0xc9, 0x0b, 0x98, 0xed, 0xca, 0x2d, 0x97, 0xd9,
	0x26, 0xcd, 0x9b, 0x39, 0x53, 0x09, 0xad, 0x0f, 0xa8, 0x7d, 0x27, 0x54, 0xa3, 0xee, 0x6f, 0x04,
	0x93, 0xf6, 0x00, 0x62, 0x03, 0x84, 0x11, 0xfb, 0x12, 0x5c, 0x06, 0xe1, 0x55, 0x80, 0x0f, 0xc8,
	0x14, 0xc6, 0x61, 0xc4, 0xfc, 0xcf, 0x18, 0x35, 0x30, 0xf0, 0xf1, 0xa0, 0x81, 0x1f, 0x13, 0x3c,
	0x6c, 0xa1, 0x8f, 0x47, 0x0d, 0x3c, 0x4f, 0xf0, 0xb8, 0x85, 0x3e, 0x3e, 0x24, 0xc7, 0x30, 0x09,
	0x23, 0x46, 0xfd, 0x73, 0xff, 0x1b, 0x3e, 0x22, 0x04, 0xec, 0x30, 0x62, 0x71, 0x72, 0x4a, 0x93,
	0x98, 0x5d, 0x5d, 0x24, 0x6b, 0x3c, 0x21, 0x33, 0xb0, 0xc2, 0x88, 0x9d, 0x85, 0x41, 0x72, 0x7a,
	0x11, 0xc4, 0x78, 0x6a, 0xae, 0xaa, 0x7b, 0xc4, 0x3a, 0xfd, 0xf5, 0x61, 0xf5, 0x9b, 0xdf, 0xfc,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x06, 0x22, 0x01, 0xaf, 0x09, 0x03, 0x00, 0x00,
}
