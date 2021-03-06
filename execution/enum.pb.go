// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execution/enum.proto

package execution

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowExecutionStatus int32

const (
	WorkflowExecutionStatus_Unknown        WorkflowExecutionStatus = 0
	WorkflowExecutionStatus_Running        WorkflowExecutionStatus = 1
	WorkflowExecutionStatus_Completed      WorkflowExecutionStatus = 2
	WorkflowExecutionStatus_Failed         WorkflowExecutionStatus = 3
	WorkflowExecutionStatus_Canceled       WorkflowExecutionStatus = 4
	WorkflowExecutionStatus_Terminated     WorkflowExecutionStatus = 5
	WorkflowExecutionStatus_ContinuedAsNew WorkflowExecutionStatus = 6
	WorkflowExecutionStatus_TimedOut       WorkflowExecutionStatus = 7
)

var WorkflowExecutionStatus_name = map[int32]string{
	0: "Unknown",
	1: "Running",
	2: "Completed",
	3: "Failed",
	4: "Canceled",
	5: "Terminated",
	6: "ContinuedAsNew",
	7: "TimedOut",
}

var WorkflowExecutionStatus_value = map[string]int32{
	"Unknown":        0,
	"Running":        1,
	"Completed":      2,
	"Failed":         3,
	"Canceled":       4,
	"Terminated":     5,
	"ContinuedAsNew": 6,
	"TimedOut":       7,
}

func (x WorkflowExecutionStatus) String() string {
	return proto.EnumName(WorkflowExecutionStatus_name, int32(x))
}

func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2297211ce2f9cb26, []int{0}
}

type PendingActivityState int32

const (
	PendingActivityState_Scheduled       PendingActivityState = 0
	PendingActivityState_Started         PendingActivityState = 1
	PendingActivityState_CancelRequested PendingActivityState = 2
)

var PendingActivityState_name = map[int32]string{
	0: "Scheduled",
	1: "Started",
	2: "CancelRequested",
}

var PendingActivityState_value = map[string]int32{
	"Scheduled":       0,
	"Started":         1,
	"CancelRequested": 2,
}

func (x PendingActivityState) String() string {
	return proto.EnumName(PendingActivityState_name, int32(x))
}

func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2297211ce2f9cb26, []int{1}
}

func init() {
	proto.RegisterEnum("execution.WorkflowExecutionStatus", WorkflowExecutionStatus_name, WorkflowExecutionStatus_value)
	proto.RegisterEnum("execution.PendingActivityState", PendingActivityState_name, PendingActivityState_value)
}

func init() { proto.RegisterFile("execution/enum.proto", fileDescriptor_2297211ce2f9cb26) }

var fileDescriptor_2297211ce2f9cb26 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xc1, 0x4e, 0x32, 0x31,
	0x18, 0xdc, 0xfe, 0xbf, 0x82, 0x7c, 0x2a, 0x36, 0x95, 0xc4, 0x83, 0x49, 0xaf, 0x26, 0x24, 0xc2,
	0xc1, 0x27, 0x40, 0xa2, 0xde, 0x94, 0x00, 0xc6, 0xc4, 0x93, 0x75, 0xfb, 0x89, 0x0d, 0xbb, 0x5f,
	0x71, 0x69, 0x5d, 0x7d, 0x09, 0xe3, 0x63, 0x79, 0xe4, 0xe8, 0xd1, 0xec, 0xbe, 0x88, 0xe9, 0x12,
	0x36, 0xde, 0x3a, 0xd3, 0xf9, 0x66, 0x26, 0x03, 0x1d, 0x7c, 0xc3, 0xd8, 0x3b, 0x63, 0xa9, 0x8f,
	0xe4, 0xd3, 0xde, 0x22, 0xb3, 0xce, 0x8a, 0x56, 0xcd, 0x76, 0x3f, 0x18, 0x1c, 0xdd, 0xd9, 0x6c,
	0xfe, 0x94, 0xd8, 0xfc, 0x62, 0xc3, 0x4e, 0x9c, 0x72, 0x7e, 0x29, 0x76, 0xa1, 0x79, 0x4b, 0x73,
	0xb2, 0x39, 0xf1, 0x28, 0x80, 0xb1, 0x27, 0x32, 0x34, 0xe3, 0x4c, 0xec, 0x43, 0x6b, 0x68, 0xd3,
	0x45, 0x82, 0x0e, 0x35, 0xff, 0x27, 0x00, 0x1a, 0x97, 0xca, 0x24, 0xa8, 0xf9, 0x7f, 0xb1, 0x07,
	0x3b, 0x43, 0x45, 0x31, 0x06, 0xb4, 0x25, 0xda, 0x00, 0x53, 0xcc, 0x52, 0x43, 0x2a, 0x28, 0xb7,
	0x85, 0x80, 0xf6, 0xd0, 0x92, 0x33, 0xe4, 0x51, 0x0f, 0x96, 0xd7, 0x98, 0xf3, 0x46, 0xb8, 0x98,
	0x9a, 0x14, 0xf5, 0x8d, 0x77, 0xbc, 0xd9, 0xbd, 0x82, 0xce, 0x08, 0x49, 0x1b, 0x9a, 0x0d, 0x62,
	0x67, 0x5e, 0x8d, 0x7b, 0x0f, 0x6d, 0x30, 0x44, 0x4e, 0xe2, 0x67, 0xd4, 0x3e, 0x18, 0x57, 0x75,
	0x26, 0x4e, 0x65, 0xc1, 0x95, 0x89, 0x43, 0x38, 0x58, 0x67, 0x8e, 0xf1, 0xc5, 0xe3, 0xb2, 0x2a,
	0x75, 0xfe, 0xf0, 0x55, 0x48, 0xb6, 0x2a, 0x24, 0xfb, 0x29, 0x24, 0xfb, 0x2c, 0x65, 0xb4, 0x2a,
	0x65, 0xf4, 0x5d, 0xca, 0x08, 0x8e, 0x8d, 0xed, 0x39, 0x4c, 0x17, 0x36, 0x53, 0xc9, 0x7a, 0x91,
	0x5e, 0x3d, 0xc8, 0x88, 0xdd, 0x9f, 0xcc, 0xfe, 0x7c, 0x1b, 0xdb, 0xdf, 0xbc, 0x4f, 0x2b, 0x69,
	0xbf, 0x96, 0x3e, 0x36, 0x2a, 0xe2, 0xec, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x2c, 0x39, 0xb7,
	0x65, 0x01, 0x00, 0x00,
}
