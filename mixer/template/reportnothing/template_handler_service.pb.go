// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/reportnothing/template_handler_service.proto

/*
Package reportnothing is a generated protocol buffer package.

The `reportnothing` template represents an empty block of data, which can useful
in different testing scenarios.

ReportNothing represents an empty block of data that is used for Report-capable
adapters which don't require any parameters. This is primarily intended for testing
scenarios.

Example config:
```yaml
apiVersion: "config.istio.io/v1alpha2"
kind: reportnothing
metadata:
  name: reportrequest
  namespace: istio-system
spec:
```

It is generated from these files:
	mixer/template/reportnothing/template_handler_service.proto

It has these top-level messages:
	HandleReportNothingRequest
	InstanceMsg
	Type
	InstanceParam
*/
package reportnothing

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import google_protobuf1 "github.com/gogo/protobuf/types"
import istio_mixer_adapter_model_v1beta11 "istio.io/api/mixer/adapter/model/v1beta1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for HandleReportNothing method.
type HandleReportNothingRequest struct {
	// 'reportnothing' instances.
	Instances []*InstanceMsg `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *google_protobuf1.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleReportNothingRequest) Reset()      { *m = HandleReportNothingRequest{} }
func (*HandleReportNothingRequest) ProtoMessage() {}
func (*HandleReportNothingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{0}
}

// Contains instance payload for 'reportnothing' template. This is passed to infrastructure backends during request-time
// through HandleReportNothingService.HandleReportNothing.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{1}
}

// Contains inferred type information about specific instance of 'reportnothing' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateHandlerService, []int{2} }

// Represents instance configuration schema for 'reportnothing' template.
type InstanceParam struct {
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{3}
}

func init() {
	proto.RegisterType((*HandleReportNothingRequest)(nil), "reportnothing.HandleReportNothingRequest")
	proto.RegisterType((*InstanceMsg)(nil), "reportnothing.InstanceMsg")
	proto.RegisterType((*Type)(nil), "reportnothing.Type")
	proto.RegisterType((*InstanceParam)(nil), "reportnothing.InstanceParam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HandleReportNothingService service

type HandleReportNothingServiceClient interface {
	// HandleReportNothing is called by Mixer at request-time to deliver 'reportnothing' instances to the backend.
	HandleReportNothing(ctx context.Context, in *HandleReportNothingRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.ReportResult, error)
}

type handleReportNothingServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleReportNothingServiceClient(cc *grpc.ClientConn) HandleReportNothingServiceClient {
	return &handleReportNothingServiceClient{cc}
}

func (c *handleReportNothingServiceClient) HandleReportNothing(ctx context.Context, in *HandleReportNothingRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.ReportResult, error) {
	out := new(istio_mixer_adapter_model_v1beta11.ReportResult)
	err := grpc.Invoke(ctx, "/reportnothing.HandleReportNothingService/HandleReportNothing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HandleReportNothingService service

type HandleReportNothingServiceServer interface {
	// HandleReportNothing is called by Mixer at request-time to deliver 'reportnothing' instances to the backend.
	HandleReportNothing(context.Context, *HandleReportNothingRequest) (*istio_mixer_adapter_model_v1beta11.ReportResult, error)
}

func RegisterHandleReportNothingServiceServer(s *grpc.Server, srv HandleReportNothingServiceServer) {
	s.RegisterService(&_HandleReportNothingService_serviceDesc, srv)
}

func _HandleReportNothingService_HandleReportNothing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleReportNothingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleReportNothingServiceServer).HandleReportNothing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reportnothing.HandleReportNothingService/HandleReportNothing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleReportNothingServiceServer).HandleReportNothing(ctx, req.(*HandleReportNothingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleReportNothingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reportnothing.HandleReportNothingService",
	HandlerType: (*HandleReportNothingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleReportNothing",
			Handler:    _HandleReportNothingService_HandleReportNothing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/template/reportnothing/template_handler_service.proto",
}

func (m *HandleReportNothingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleReportNothingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateHandlerService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.AdapterConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.AdapterConfig.Size()))
		n1, err := m.AdapterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.DedupId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i += copy(dAtA[i:], m.DedupId)
	}
	return i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0xe4
		i++
		dAtA[i] = 0x93
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HandleReportNothingRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
			l = e.Size()
			n += 1 + l + sovTemplateHandlerService(uint64(l))
		}
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleReportNothingRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleReportNothingRequest{`,
		`Instances:` + strings.Replace(fmt.Sprintf("%v", this.Instances), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "google_protobuf1.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleReportNothingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandleReportNothingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleReportNothingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instances = append(m.Instances, &InstanceMsg{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &google_protobuf1.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
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
				next, err := skipTemplateHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/template/reportnothing/template_handler_service.proto", fileDescriptorTemplateHandlerService)
}

var fileDescriptorTemplateHandlerService = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6e, 0xd4, 0x30,
	0x1c, 0xc6, 0xe3, 0xb6, 0x2a, 0xd4, 0xa7, 0x03, 0x29, 0x14, 0x29, 0xcd, 0x60, 0x9d, 0x22, 0x86,
	0x20, 0x21, 0x5b, 0x3d, 0x24, 0x84, 0xd4, 0x09, 0x58, 0xe8, 0x00, 0x42, 0x81, 0x3d, 0xf2, 0x5d,
	0xfe, 0x4d, 0x2d, 0x25, 0x76, 0x6a, 0x3b, 0x55, 0x6f, 0xe3, 0x0d, 0x40, 0xe2, 0x25, 0x98, 0xe0,
	0x05, 0x78, 0x80, 0x8a, 0xa9, 0x62, 0x62, 0x24, 0x81, 0x81, 0xb1, 0x23, 0x23, 0xc2, 0x4e, 0x81,
	0x43, 0x47, 0xb7, 0xfc, 0xfd, 0xfd, 0xe2, 0xff, 0xf7, 0xf9, 0xc3, 0x7b, 0xb5, 0x38, 0x01, 0xcd,
	0x2c, 0xd4, 0x4d, 0xc5, 0x2d, 0x30, 0x0d, 0x8d, 0xd2, 0x56, 0x2a, 0x7b, 0x28, 0x64, 0xf9, 0xfb,
	0x38, 0x3f, 0xe4, 0xb2, 0xa8, 0x40, 0xe7, 0x06, 0xf4, 0xb1, 0x98, 0x03, 0x6d, 0xb4, 0xb2, 0x2a,
	0x1c, 0x2f, 0xd1, 0xf1, 0x76, 0xa9, 0x4a, 0xe5, 0x14, 0xf6, 0xeb, 0xcb, 0x43, 0xf1, 0x1d, 0xbf,
	0x81, 0x17, 0xbc, 0xb1, 0xa0, 0x59, 0xad, 0x0a, 0xa8, 0xd8, 0xf1, 0xee, 0x0c, 0x2c, 0xdf, 0x65,
	0x70, 0x62, 0x41, 0x1a, 0xa1, 0xa4, 0x19, 0xe8, 0x9d, 0x52, 0xa9, 0xb2, 0x02, 0xe6, 0xa6, 0x59,
	0x7b, 0xc0, 0xb8, 0x5c, 0x0c, 0x52, 0x7a, 0xd9, 0x45, 0xde, 0x89, 0x27, 0x93, 0x77, 0x08, 0xc7,
	0x8f, 0x9d, 0xe3, 0xcc, 0x1d, 0x3f, 0xf5, 0x06, 0x33, 0x38, 0x6a, 0xc1, 0xd8, 0xf0, 0x3e, 0xde,
	0x12, 0xd2, 0x58, 0x2e, 0xe7, 0x60, 0x22, 0x34, 0x59, 0x4f, 0x47, 0xd3, 0x98, 0x2e, 0x45, 0xa1,
	0xfb, 0x83, 0xfe, 0xc4, 0x94, 0xd9, 0x1f, 0x38, 0xdc, 0xc3, 0xd7, 0x86, 0xf5, 0xf9, 0x5c, 0xc9,
	0x03, 0x51, 0x46, 0x6b, 0x13, 0x94, 0x8e, 0xa6, 0xdb, 0xd4, 0xdb, 0xa6, 0x17, 0xb6, 0xe9, 0x03,
	0xb9, 0xc8, 0xc6, 0x03, 0xfb, 0xc8, 0xa1, 0xe1, 0x0e, 0xbe, 0x5a, 0x40, 0xd1, 0x36, 0xb9, 0x28,
	0xa2, 0xf5, 0x09, 0x4a, 0xb7, 0xb2, 0x2b, 0x6e, 0xde, 0x2f, 0x92, 0x5b, 0x78, 0xf4, 0xd7, 0xc6,
	0xf0, 0x26, 0xde, 0x90, 0xbc, 0x86, 0xe8, 0xfd, 0xc7, 0x0f, 0x89, 0x03, 0xdd, 0x98, 0x6c, 0xe2,
	0x8d, 0x17, 0x8b, 0x06, 0x92, 0xeb, 0x78, 0x7c, 0x41, 0x3f, 0xe3, 0x9a, 0xd7, 0xd3, 0x57, 0xab,
	0xf3, 0x3e, 0xf7, 0x65, 0x85, 0x47, 0xf8, 0xc6, 0x0a, 0x35, 0xbc, 0xfd, 0x4f, 0xe6, 0xff, 0xbf,
	0x58, 0xcc, 0xa8, 0x30, 0x56, 0x28, 0xea, 0x1a, 0xa0, 0x43, 0x2c, 0xea, 0x1a, 0xa0, 0x43, 0x03,
	0xd4, 0xff, 0x98, 0x81, 0x69, 0x2b, 0xfb, 0xf0, 0xde, 0x69, 0x47, 0x82, 0xb3, 0x8e, 0x04, 0x9f,
	0x3b, 0x12, 0x9c, 0x77, 0x24, 0x78, 0xd9, 0x13, 0xf4, 0xb6, 0x27, 0xc1, 0x69, 0x4f, 0xd0, 0x59,
	0x4f, 0xd0, 0x97, 0x9e, 0xa0, 0xef, 0x3d, 0x09, 0xce, 0x7b, 0x82, 0x5e, 0x7f, 0x25, 0xc1, 0x8f,
	0x4f, 0xdf, 0xde, 0xac, 0xa1, 0xd9, 0xa6, 0x7b, 0xc0, 0xbb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x53, 0x5b, 0x59, 0x97, 0x02, 0x00, 0x00,
}
