// Code generated by protoc-gen-gogo.
// source: ca.proto
// DO NOT EDIT!

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import raftpicker "github.com/docker/swarm-v2/manager/raftpicker"
import codes "google.golang.org/grpc/codes"
import metadata "google.golang.org/grpc/metadata"
import transport "google.golang.org/grpc/transport"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IssuanceStatus_Status int32

const (
	IssuanceStatusUnknown   IssuanceStatus_Status = 0
	IssuanceStatusPreparing IssuanceStatus_Status = 1
	IssuanceStatusReady     IssuanceStatus_Status = 2
	IssuanceStatusComplete  IssuanceStatus_Status = 3
	IssuanceStatusFailed    IssuanceStatus_Status = 4
	IssuanceStatusRejected  IssuanceStatus_Status = 5
)

var IssuanceStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "PREPARING",
	2: "READY",
	3: "COMPLETE",
	4: "FAILED",
	5: "REJECTED",
}
var IssuanceStatus_Status_value = map[string]int32{
	"UNKNOWN":   0,
	"PREPARING": 1,
	"READY":     2,
	"COMPLETE":  3,
	"FAILED":    4,
	"REJECTED":  5,
}

func (x IssuanceStatus_Status) String() string {
	return proto.EnumName(IssuanceStatus_Status_name, int32(x))
}
func (IssuanceStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorCa, []int{0, 0} }

type IssuanceStatus struct {
	Status IssuanceStatus_Status `protobuf:"varint,1,opt,name=status,proto3,enum=docker.cluster.api.IssuanceStatus_Status" json:"status,omitempty"`
	// Err is set if the Certificate Issuance is in an error state.
	//
	// The following states should report a companion error:
	//
	// 	FAILED, REJECTED
	//
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *IssuanceStatus) Reset()                    { *m = IssuanceStatus{} }
func (*IssuanceStatus) ProtoMessage()               {}
func (*IssuanceStatus) Descriptor() ([]byte, []int) { return fileDescriptorCa, []int{0} }

type IssueCertificateRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	CSR  []byte `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *IssueCertificateRequest) Reset()                    { *m = IssueCertificateRequest{} }
func (*IssueCertificateRequest) ProtoMessage()               {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptorCa, []int{1} }

// TODO(diogo): This response should be a random ID that agents
// can then query for status.
type IssueCertificateResponse struct {
	Status           *IssuanceStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CertificateChain []byte          `protobuf:"bytes,2,opt,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
}

func (m *IssueCertificateResponse) Reset()                    { *m = IssueCertificateResponse{} }
func (*IssueCertificateResponse) ProtoMessage()               {}
func (*IssueCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptorCa, []int{2} }

type GetRootCACertificateRequest struct {
}

func (m *GetRootCACertificateRequest) Reset()                    { *m = GetRootCACertificateRequest{} }
func (*GetRootCACertificateRequest) ProtoMessage()               {}
func (*GetRootCACertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptorCa, []int{3} }

type GetRootCACertificateResponse struct {
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (m *GetRootCACertificateResponse) Reset()                    { *m = GetRootCACertificateResponse{} }
func (*GetRootCACertificateResponse) ProtoMessage()               {}
func (*GetRootCACertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptorCa, []int{4} }

func init() {
	proto.RegisterType((*IssuanceStatus)(nil), "docker.cluster.api.IssuanceStatus")
	proto.RegisterType((*IssueCertificateRequest)(nil), "docker.cluster.api.IssueCertificateRequest")
	proto.RegisterType((*IssueCertificateResponse)(nil), "docker.cluster.api.IssueCertificateResponse")
	proto.RegisterType((*GetRootCACertificateRequest)(nil), "docker.cluster.api.GetRootCACertificateRequest")
	proto.RegisterType((*GetRootCACertificateResponse)(nil), "docker.cluster.api.GetRootCACertificateResponse")
	proto.RegisterEnum("docker.cluster.api.IssuanceStatus_Status", IssuanceStatus_Status_name, IssuanceStatus_Status_value)
}

func (m *IssuanceStatus) Copy() *IssuanceStatus {
	if m == nil {
		return nil
	}

	o := &IssuanceStatus{
		Status: m.Status,
		Err:    m.Err,
	}

	return o
}

func (m *IssueCertificateRequest) Copy() *IssueCertificateRequest {
	if m == nil {
		return nil
	}

	o := &IssueCertificateRequest{
		Role: m.Role,
		CSR:  m.CSR,
	}

	return o
}

func (m *IssueCertificateResponse) Copy() *IssueCertificateResponse {
	if m == nil {
		return nil
	}

	o := &IssueCertificateResponse{
		Status:           m.Status.Copy(),
		CertificateChain: m.CertificateChain,
	}

	return o
}

func (m *GetRootCACertificateRequest) Copy() *GetRootCACertificateRequest {
	if m == nil {
		return nil
	}

	o := &GetRootCACertificateRequest{}

	return o
}

func (m *GetRootCACertificateResponse) Copy() *GetRootCACertificateResponse {
	if m == nil {
		return nil
	}

	o := &GetRootCACertificateResponse{
		Certificate: m.Certificate,
	}

	return o
}

func (this *IssuanceStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&api.IssuanceStatus{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Err: "+fmt.Sprintf("%#v", this.Err)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IssueCertificateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&api.IssueCertificateRequest{")
	s = append(s, "Role: "+fmt.Sprintf("%#v", this.Role)+",\n")
	s = append(s, "CSR: "+fmt.Sprintf("%#v", this.CSR)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IssueCertificateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&api.IssueCertificateResponse{")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	}
	s = append(s, "CertificateChain: "+fmt.Sprintf("%#v", this.CertificateChain)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetRootCACertificateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&api.GetRootCACertificateRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetRootCACertificateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api.GetRootCACertificateResponse{")
	s = append(s, "Certificate: "+fmt.Sprintf("%#v", this.Certificate)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCa(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringCa(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for CA service

type CAClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error)
	GetRootCACertificate(ctx context.Context, in *GetRootCACertificateRequest, opts ...grpc.CallOption) (*GetRootCACertificateResponse, error)
}

type cAClient struct {
	cc *grpc.ClientConn
}

func NewCAClient(cc *grpc.ClientConn) CAClient {
	return &cAClient{cc}
}

func (c *cAClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error) {
	out := new(IssueCertificateResponse)
	err := grpc.Invoke(ctx, "/docker.cluster.api.CA/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAClient) GetRootCACertificate(ctx context.Context, in *GetRootCACertificateRequest, opts ...grpc.CallOption) (*GetRootCACertificateResponse, error) {
	out := new(GetRootCACertificateResponse)
	err := grpc.Invoke(ctx, "/docker.cluster.api.CA/GetRootCACertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CA service

type CAServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*IssueCertificateResponse, error)
	GetRootCACertificate(context.Context, *GetRootCACertificateRequest) (*GetRootCACertificateResponse, error)
}

func RegisterCAServer(s *grpc.Server, srv CAServer) {
	s.RegisterService(&_CA_serviceDesc, srv)
}

func _CA_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CAServer).IssueCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CA_GetRootCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRootCACertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CAServer).GetRootCACertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "docker.cluster.api.CA",
	HandlerType: (*CAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CA_IssueCertificate_Handler,
		},
		{
			MethodName: "GetRootCACertificate",
			Handler:    _CA_GetRootCACertificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *IssuanceStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IssuanceStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCa(data, i, uint64(m.Status))
	}
	if len(m.Err) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintCa(data, i, uint64(len(m.Err)))
		i += copy(data[i:], m.Err)
	}
	return i, nil
}

func (m *IssueCertificateRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IssueCertificateRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCa(data, i, uint64(len(m.Role)))
		i += copy(data[i:], m.Role)
	}
	if m.CSR != nil {
		if len(m.CSR) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintCa(data, i, uint64(len(m.CSR)))
			i += copy(data[i:], m.CSR)
		}
	}
	return i, nil
}

func (m *IssueCertificateResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IssueCertificateResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		data[i] = 0xa
		i++
		i = encodeVarintCa(data, i, uint64(m.Status.Size()))
		n1, err := m.Status.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CertificateChain != nil {
		if len(m.CertificateChain) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintCa(data, i, uint64(len(m.CertificateChain)))
			i += copy(data[i:], m.CertificateChain)
		}
	}
	return i, nil
}

func (m *GetRootCACertificateRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetRootCACertificateRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetRootCACertificateResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetRootCACertificateResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Certificate != nil {
		if len(m.Certificate) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintCa(data, i, uint64(len(m.Certificate)))
			i += copy(data[i:], m.Certificate)
		}
	}
	return i, nil
}

func encodeFixed64Ca(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Ca(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCa(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

type raftProxyCAServer struct {
	local   CAServer
	conn    *grpc.ClientConn
	cluster raftpicker.RaftCluster
}

func NewRaftProxyCAServer(local CAServer, conn *grpc.ClientConn, cluster raftpicker.RaftCluster) CAServer {
	return &raftProxyCAServer{
		local:   local,
		cluster: cluster,
		conn:    conn,
	}
}

func (p *raftProxyCAServer) IssueCertificate(ctx context.Context, r *IssueCertificateRequest) (*IssueCertificateResponse, error) {

	if p.cluster.IsLeader() {
		return p.local.IssueCertificate(ctx, r)
	}
	var addr string
	s, ok := transport.StreamFromContext(ctx)
	if ok {
		addr = s.ServerTransport().RemoteAddr().String()
	}
	md, ok := metadata.FromContext(ctx)
	if ok && len(md["redirect"]) != 0 {
		return nil, grpc.Errorf(codes.ResourceExhausted, "more than one redirect to leader from: %s", md["redirect"])
	}
	if !ok {
		md = metadata.New(map[string]string{})
	}
	md["redirect"] = append(md["redirect"], addr)
	ctx = metadata.NewContext(ctx, md)

	return NewCAClient(p.conn).IssueCertificate(ctx, r)
}

func (p *raftProxyCAServer) GetRootCACertificate(ctx context.Context, r *GetRootCACertificateRequest) (*GetRootCACertificateResponse, error) {

	if p.cluster.IsLeader() {
		return p.local.GetRootCACertificate(ctx, r)
	}
	if err := p.initConn(); err != nil {
		return nil, err
	}

	var addr string
	s, ok := transport.StreamFromContext(ctx)
	if ok {
		addr = s.ServerTransport().RemoteAddr().String()
	}
	md, ok := metadata.FromContext(ctx)
	if ok && len(md["redirect"]) != 0 {
		return nil, grpc.Errorf(codes.ResourceExhausted, "more than one redirect to leader from: %s", md["redirect"])
	}
	if !ok {
		md = metadata.New(map[string]string{})
	}
	md["redirect"] = append(md["redirect"], addr)
	ctx = metadata.NewContext(ctx, md)

	return NewCAClient(p.conn).GetRootCACertificate(ctx, r)
}

func (m *IssuanceStatus) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovCa(uint64(m.Status))
	}
	l = len(m.Err)
	if l > 0 {
		n += 1 + l + sovCa(uint64(l))
	}
	return n
}

func (m *IssueCertificateRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovCa(uint64(l))
	}
	if m.CSR != nil {
		l = len(m.CSR)
		if l > 0 {
			n += 1 + l + sovCa(uint64(l))
		}
	}
	return n
}

func (m *IssueCertificateResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovCa(uint64(l))
	}
	if m.CertificateChain != nil {
		l = len(m.CertificateChain)
		if l > 0 {
			n += 1 + l + sovCa(uint64(l))
		}
	}
	return n
}

func (m *GetRootCACertificateRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetRootCACertificateResponse) Size() (n int) {
	var l int
	_ = l
	if m.Certificate != nil {
		l = len(m.Certificate)
		if l > 0 {
			n += 1 + l + sovCa(uint64(l))
		}
	}
	return n
}

func sovCa(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCa(x uint64) (n int) {
	return sovCa(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *IssuanceStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IssuanceStatus{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Err:` + fmt.Sprintf("%v", this.Err) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IssueCertificateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IssueCertificateRequest{`,
		`Role:` + fmt.Sprintf("%v", this.Role) + `,`,
		`CSR:` + fmt.Sprintf("%v", this.CSR) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IssueCertificateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IssueCertificateResponse{`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "IssuanceStatus", "IssuanceStatus", 1) + `,`,
		`CertificateChain:` + fmt.Sprintf("%v", this.CertificateChain) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetRootCACertificateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRootCACertificateRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetRootCACertificateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRootCACertificateResponse{`,
		`Certificate:` + fmt.Sprintf("%v", this.Certificate) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCa(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *IssuanceStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IssuanceStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssuanceStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Status |= (IssuanceStatus_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Err = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
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
func (m *IssueCertificateRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IssueCertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueCertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CSR", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CSR = append(m.CSR[:0], data[iNdEx:postIndex]...)
			if m.CSR == nil {
				m.CSR = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
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
func (m *IssueCertificateResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IssueCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &IssuanceStatus{}
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateChain", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificateChain = append(m.CertificateChain[:0], data[iNdEx:postIndex]...)
			if m.CertificateChain == nil {
				m.CertificateChain = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
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
func (m *GetRootCACertificateRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRootCACertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRootCACertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCa(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
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
func (m *GetRootCACertificateResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRootCACertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRootCACertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = append(m.Certificate[:0], data[iNdEx:postIndex]...)
			if m.Certificate == nil {
				m.Certificate = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
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
func skipCa(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCa
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCa
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCa
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCa
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCa
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCa(data[start:])
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
	ErrInvalidLengthCa = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCa   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCa = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x27, 0x4d, 0x9b, 0x69, 0x55, 0x99, 0x25, 0xd0, 0x74, 0x5b, 0x42, 0x64, 0x21,
	0x54, 0x28, 0x72, 0x51, 0xb8, 0x71, 0xc2, 0x71, 0xdd, 0x12, 0x68, 0xd3, 0x68, 0xdb, 0x0a, 0x71,
	0x42, 0xc6, 0x19, 0x82, 0x69, 0xf0, 0x1a, 0x7b, 0x23, 0x84, 0xb8, 0x54, 0xe2, 0x82, 0x78, 0x07,
	0x4e, 0xbc, 0x4c, 0x8f, 0x1c, 0x39, 0x21, 0xda, 0x27, 0x40, 0x88, 0x07, 0x60, 0x6d, 0x07, 0x88,
	0x89, 0x81, 0x1c, 0x46, 0x3b, 0xda, 0xf9, 0xf6, 0xf7, 0x3f, 0x33, 0x32, 0xcc, 0xb9, 0x8e, 0x11,
	0x84, 0x5c, 0x70, 0x42, 0x7a, 0xdc, 0x3d, 0xc2, 0xd0, 0x70, 0x07, 0xc3, 0x48, 0xc8, 0xd3, 0x09,
	0x3c, 0x5a, 0xed, 0xf3, 0x3e, 0x4f, 0xca, 0x1b, 0x71, 0x96, 0x92, 0xfa, 0x77, 0x15, 0x16, 0xdb,
	0x51, 0x34, 0x74, 0x7c, 0x17, 0xf7, 0x85, 0x23, 0x86, 0x11, 0xd9, 0x85, 0x72, 0x94, 0x64, 0x35,
	0xa5, 0xa1, 0xac, 0x2d, 0x36, 0xaf, 0x19, 0x93, 0x6a, 0x46, 0xf6, 0x8d, 0x91, 0x1e, 0x2d, 0x38,
	0xfb, 0x7c, 0xb9, 0x9c, 0xe6, 0x6c, 0x24, 0x42, 0x34, 0x28, 0x62, 0x18, 0xd6, 0x54, 0xa9, 0x55,
	0x61, 0x71, 0xaa, 0x1f, 0xab, 0x30, 0x82, 0xc8, 0x55, 0x98, 0x3d, 0xec, 0xdc, 0xef, 0xec, 0x3d,
	0xe8, 0x68, 0x05, 0xba, 0xfc, 0xee, 0x7d, 0xe3, 0x42, 0x56, 0xf8, 0xd0, 0x3f, 0xf2, 0xf9, 0x4b,
	0x9f, 0x5c, 0x87, 0x4a, 0x97, 0xd9, 0x5d, 0x93, 0xb5, 0x3b, 0xdb, 0x9a, 0x42, 0x57, 0x24, 0xb9,
	0x94, 0x25, 0xbb, 0x21, 0x06, 0x4e, 0xe8, 0xf9, 0x7d, 0xa2, 0xc3, 0x0c, 0xb3, 0xcd, 0xcd, 0x87,
	0x9a, 0x4a, 0x97, 0x24, 0x77, 0x3e, 0xcb, 0x31, 0x74, 0x7a, 0xaf, 0xc8, 0x1a, 0xcc, 0x59, 0x7b,
	0xbb, 0xdd, 0x1d, 0xfb, 0xc0, 0xd6, 0x8a, 0x94, 0x4a, 0xec, 0x62, 0x16, 0xb3, 0xf8, 0xf3, 0x60,
	0x80, 0x02, 0xc9, 0x15, 0x28, 0x6f, 0x99, 0xed, 0x1d, 0x7b, 0x53, 0x2b, 0xd1, 0x9a, 0xe4, 0xaa,
	0x59, 0x6e, 0xcb, 0xf1, 0x06, 0xd8, 0x8b, 0xf5, 0x98, 0x7d, 0xcf, 0xb6, 0x0e, 0x24, 0x37, 0x93,
	0xa7, 0xc7, 0xf0, 0x19, 0xba, 0x02, 0x7b, 0xb4, 0xf4, 0xf6, 0x43, 0xbd, 0xa0, 0xdf, 0x85, 0xc4,
	0x3e, 0x5a, 0x18, 0x0a, 0xef, 0x89, 0xe7, 0x3a, 0x02, 0x19, 0xbe, 0x18, 0x62, 0x24, 0x08, 0x81,
	0x52, 0xc8, 0x07, 0x98, 0x0c, 0xbf, 0xc2, 0x92, 0x9c, 0x2c, 0x43, 0xd1, 0x8d, 0xd2, 0x19, 0x2e,
	0xb4, 0x66, 0xe5, 0x90, 0x8b, 0xd6, 0x3e, 0x63, 0xf1, 0x9d, 0xfe, 0x46, 0x81, 0xda, 0xa4, 0x54,
	0x14, 0x70, 0x3f, 0x42, 0x72, 0x3b, 0xb3, 0xca, 0xf9, 0xa6, 0xfe, 0xff, 0x55, 0xfe, 0xda, 0xdb,
	0x3a, 0x9c, 0x73, 0x7f, 0x4b, 0x3e, 0x72, 0x9f, 0x3a, 0x9e, 0x9f, 0x3a, 0x60, 0xda, 0x58, 0xc1,
	0x8a, 0xef, 0xf5, 0x4b, 0xb0, 0xb2, 0x8d, 0x82, 0x71, 0x2e, 0x2c, 0x73, 0xb2, 0x27, 0xfd, 0x0e,
	0xac, 0xe6, 0x97, 0x47, 0x3e, 0x1b, 0x30, 0x3f, 0x26, 0x99, 0x98, 0x5d, 0x60, 0xe3, 0x57, 0xcd,
	0x6f, 0x0a, 0xa8, 0x96, 0x49, 0x38, 0x68, 0x7f, 0x36, 0x4b, 0xd6, 0xff, 0xd6, 0x54, 0xce, 0x74,
	0xe9, 0x8d, 0xe9, 0xe0, 0xd4, 0x97, 0x5e, 0x20, 0xaf, 0xa1, 0x9a, 0xe7, 0x9c, 0x6c, 0xe4, 0xe9,
	0xfc, 0x63, 0x04, 0xf4, 0xe6, 0xf4, 0x0f, 0x7e, 0x7e, 0xbc, 0xb5, 0x7a, 0x72, 0x5a, 0x2f, 0x7c,
	0x92, 0xf1, 0xf5, 0xb4, 0xae, 0x1c, 0x9f, 0xd5, 0x95, 0x13, 0x19, 0x1f, 0x65, 0x7c, 0x91, 0xf1,
	0xb8, 0x9c, 0xfc, 0xc1, 0xb7, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xad, 0xeb, 0x3b, 0xf7,
	0x03, 0x00, 0x00,
}
