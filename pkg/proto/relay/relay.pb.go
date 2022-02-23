// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay/relay.proto

package pbRelay // import "./relay"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sdk_ws "Open_IM/pkg/proto/sdk_ws"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type OnlinePushMsgReq struct {
	OperationID          string          `protobuf:"bytes,1,opt,name=OperationID" json:"OperationID,omitempty"`
	MsgData              *sdk_ws.MsgData `protobuf:"bytes,2,opt,name=msgData" json:"msgData,omitempty"`
	PushToUserID         string          `protobuf:"bytes,3,opt,name=pushToUserID" json:"pushToUserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OnlinePushMsgReq) Reset()         { *m = OnlinePushMsgReq{} }
func (m *OnlinePushMsgReq) String() string { return proto.CompactTextString(m) }
func (*OnlinePushMsgReq) ProtoMessage()    {}
func (*OnlinePushMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{0}
}
func (m *OnlinePushMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlinePushMsgReq.Unmarshal(m, b)
}
func (m *OnlinePushMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlinePushMsgReq.Marshal(b, m, deterministic)
}
func (dst *OnlinePushMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlinePushMsgReq.Merge(dst, src)
}
func (m *OnlinePushMsgReq) XXX_Size() int {
	return xxx_messageInfo_OnlinePushMsgReq.Size(m)
}
func (m *OnlinePushMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlinePushMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_OnlinePushMsgReq proto.InternalMessageInfo

func (m *OnlinePushMsgReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *OnlinePushMsgReq) GetMsgData() *sdk_ws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *OnlinePushMsgReq) GetPushToUserID() string {
	if m != nil {
		return m.PushToUserID
	}
	return ""
}

type OnlinePushMsgResp struct {
	Resp                 []*SingleMsgToUser `protobuf:"bytes,1,rep,name=resp" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OnlinePushMsgResp) Reset()         { *m = OnlinePushMsgResp{} }
func (m *OnlinePushMsgResp) String() string { return proto.CompactTextString(m) }
func (*OnlinePushMsgResp) ProtoMessage()    {}
func (*OnlinePushMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{1}
}
func (m *OnlinePushMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlinePushMsgResp.Unmarshal(m, b)
}
func (m *OnlinePushMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlinePushMsgResp.Marshal(b, m, deterministic)
}
func (dst *OnlinePushMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlinePushMsgResp.Merge(dst, src)
}
func (m *OnlinePushMsgResp) XXX_Size() int {
	return xxx_messageInfo_OnlinePushMsgResp.Size(m)
}
func (m *OnlinePushMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlinePushMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_OnlinePushMsgResp proto.InternalMessageInfo

func (m *OnlinePushMsgResp) GetResp() []*SingleMsgToUser {
	if m != nil {
		return m.Resp
	}
	return nil
}

type SingleMsgToUser struct {
	ResultCode           int64    `protobuf:"varint,1,opt,name=ResultCode" json:"ResultCode,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID" json:"RecvID,omitempty"`
	RecvPlatFormID       int32    `protobuf:"varint,3,opt,name=RecvPlatFormID" json:"RecvPlatFormID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleMsgToUser) Reset()         { *m = SingleMsgToUser{} }
func (m *SingleMsgToUser) String() string { return proto.CompactTextString(m) }
func (*SingleMsgToUser) ProtoMessage()    {}
func (*SingleMsgToUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{2}
}
func (m *SingleMsgToUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleMsgToUser.Unmarshal(m, b)
}
func (m *SingleMsgToUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleMsgToUser.Marshal(b, m, deterministic)
}
func (dst *SingleMsgToUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleMsgToUser.Merge(dst, src)
}
func (m *SingleMsgToUser) XXX_Size() int {
	return xxx_messageInfo_SingleMsgToUser.Size(m)
}
func (m *SingleMsgToUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleMsgToUser.DiscardUnknown(m)
}

var xxx_messageInfo_SingleMsgToUser proto.InternalMessageInfo

func (m *SingleMsgToUser) GetResultCode() int64 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *SingleMsgToUser) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *SingleMsgToUser) GetRecvPlatFormID() int32 {
	if m != nil {
		return m.RecvPlatFormID
	}
	return 0
}

type GetUsersOnlineStatusReq struct {
	UserIDList           []string `protobuf:"bytes,1,rep,name=userIDList" json:"userIDList,omitempty"`
	OperationID          string   `protobuf:"bytes,2,opt,name=operationID" json:"operationID,omitempty"`
	OpUserID             string   `protobuf:"bytes,3,opt,name=opUserID" json:"opUserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersOnlineStatusReq) Reset()         { *m = GetUsersOnlineStatusReq{} }
func (m *GetUsersOnlineStatusReq) String() string { return proto.CompactTextString(m) }
func (*GetUsersOnlineStatusReq) ProtoMessage()    {}
func (*GetUsersOnlineStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{3}
}
func (m *GetUsersOnlineStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersOnlineStatusReq.Unmarshal(m, b)
}
func (m *GetUsersOnlineStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersOnlineStatusReq.Marshal(b, m, deterministic)
}
func (dst *GetUsersOnlineStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersOnlineStatusReq.Merge(dst, src)
}
func (m *GetUsersOnlineStatusReq) XXX_Size() int {
	return xxx_messageInfo_GetUsersOnlineStatusReq.Size(m)
}
func (m *GetUsersOnlineStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersOnlineStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersOnlineStatusReq proto.InternalMessageInfo

func (m *GetUsersOnlineStatusReq) GetUserIDList() []string {
	if m != nil {
		return m.UserIDList
	}
	return nil
}

func (m *GetUsersOnlineStatusReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *GetUsersOnlineStatusReq) GetOpUserID() string {
	if m != nil {
		return m.OpUserID
	}
	return ""
}

type GetUsersOnlineStatusResp struct {
	ErrCode              int32                                     `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg               string                                    `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	SuccessResult        []*GetUsersOnlineStatusResp_SuccessResult `protobuf:"bytes,3,rep,name=successResult" json:"successResult,omitempty"`
	FailedResult         []*GetUsersOnlineStatusResp_FailedDetail  `protobuf:"bytes,4,rep,name=failedResult" json:"failedResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *GetUsersOnlineStatusResp) Reset()         { *m = GetUsersOnlineStatusResp{} }
func (m *GetUsersOnlineStatusResp) String() string { return proto.CompactTextString(m) }
func (*GetUsersOnlineStatusResp) ProtoMessage()    {}
func (*GetUsersOnlineStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{4}
}
func (m *GetUsersOnlineStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersOnlineStatusResp.Unmarshal(m, b)
}
func (m *GetUsersOnlineStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersOnlineStatusResp.Marshal(b, m, deterministic)
}
func (dst *GetUsersOnlineStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersOnlineStatusResp.Merge(dst, src)
}
func (m *GetUsersOnlineStatusResp) XXX_Size() int {
	return xxx_messageInfo_GetUsersOnlineStatusResp.Size(m)
}
func (m *GetUsersOnlineStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersOnlineStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersOnlineStatusResp proto.InternalMessageInfo

func (m *GetUsersOnlineStatusResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GetUsersOnlineStatusResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetUsersOnlineStatusResp) GetSuccessResult() []*GetUsersOnlineStatusResp_SuccessResult {
	if m != nil {
		return m.SuccessResult
	}
	return nil
}

func (m *GetUsersOnlineStatusResp) GetFailedResult() []*GetUsersOnlineStatusResp_FailedDetail {
	if m != nil {
		return m.FailedResult
	}
	return nil
}

type GetUsersOnlineStatusResp_SuccessDetail struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform" json:"platform,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersOnlineStatusResp_SuccessDetail) Reset() {
	*m = GetUsersOnlineStatusResp_SuccessDetail{}
}
func (m *GetUsersOnlineStatusResp_SuccessDetail) String() string { return proto.CompactTextString(m) }
func (*GetUsersOnlineStatusResp_SuccessDetail) ProtoMessage()    {}
func (*GetUsersOnlineStatusResp_SuccessDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{4, 0}
}
func (m *GetUsersOnlineStatusResp_SuccessDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail.Unmarshal(m, b)
}
func (m *GetUsersOnlineStatusResp_SuccessDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail.Marshal(b, m, deterministic)
}
func (dst *GetUsersOnlineStatusResp_SuccessDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail.Merge(dst, src)
}
func (m *GetUsersOnlineStatusResp_SuccessDetail) XXX_Size() int {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail.Size(m)
}
func (m *GetUsersOnlineStatusResp_SuccessDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersOnlineStatusResp_SuccessDetail proto.InternalMessageInfo

func (m *GetUsersOnlineStatusResp_SuccessDetail) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *GetUsersOnlineStatusResp_SuccessDetail) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetUsersOnlineStatusResp_FailedDetail struct {
	UserID               string   `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	ErrCode              int32    `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersOnlineStatusResp_FailedDetail) Reset()         { *m = GetUsersOnlineStatusResp_FailedDetail{} }
func (m *GetUsersOnlineStatusResp_FailedDetail) String() string { return proto.CompactTextString(m) }
func (*GetUsersOnlineStatusResp_FailedDetail) ProtoMessage()    {}
func (*GetUsersOnlineStatusResp_FailedDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{4, 1}
}
func (m *GetUsersOnlineStatusResp_FailedDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail.Unmarshal(m, b)
}
func (m *GetUsersOnlineStatusResp_FailedDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail.Marshal(b, m, deterministic)
}
func (dst *GetUsersOnlineStatusResp_FailedDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail.Merge(dst, src)
}
func (m *GetUsersOnlineStatusResp_FailedDetail) XXX_Size() int {
	return xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail.Size(m)
}
func (m *GetUsersOnlineStatusResp_FailedDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersOnlineStatusResp_FailedDetail proto.InternalMessageInfo

func (m *GetUsersOnlineStatusResp_FailedDetail) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUsersOnlineStatusResp_FailedDetail) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GetUsersOnlineStatusResp_FailedDetail) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type GetUsersOnlineStatusResp_SuccessResult struct {
	UserID               string                                    `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Status               string                                    `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	DetailPlatformStatus []*GetUsersOnlineStatusResp_SuccessDetail `protobuf:"bytes,3,rep,name=detailPlatformStatus" json:"detailPlatformStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *GetUsersOnlineStatusResp_SuccessResult) Reset() {
	*m = GetUsersOnlineStatusResp_SuccessResult{}
}
func (m *GetUsersOnlineStatusResp_SuccessResult) String() string { return proto.CompactTextString(m) }
func (*GetUsersOnlineStatusResp_SuccessResult) ProtoMessage()    {}
func (*GetUsersOnlineStatusResp_SuccessResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_34094e5333f6005a, []int{4, 2}
}
func (m *GetUsersOnlineStatusResp_SuccessResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult.Unmarshal(m, b)
}
func (m *GetUsersOnlineStatusResp_SuccessResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult.Marshal(b, m, deterministic)
}
func (dst *GetUsersOnlineStatusResp_SuccessResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult.Merge(dst, src)
}
func (m *GetUsersOnlineStatusResp_SuccessResult) XXX_Size() int {
	return xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult.Size(m)
}
func (m *GetUsersOnlineStatusResp_SuccessResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersOnlineStatusResp_SuccessResult proto.InternalMessageInfo

func (m *GetUsersOnlineStatusResp_SuccessResult) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUsersOnlineStatusResp_SuccessResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetUsersOnlineStatusResp_SuccessResult) GetDetailPlatformStatus() []*GetUsersOnlineStatusResp_SuccessDetail {
	if m != nil {
		return m.DetailPlatformStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*OnlinePushMsgReq)(nil), "relay.OnlinePushMsgReq")
	proto.RegisterType((*OnlinePushMsgResp)(nil), "relay.OnlinePushMsgResp")
	proto.RegisterType((*SingleMsgToUser)(nil), "relay.SingleMsgToUser")
	proto.RegisterType((*GetUsersOnlineStatusReq)(nil), "relay.GetUsersOnlineStatusReq")
	proto.RegisterType((*GetUsersOnlineStatusResp)(nil), "relay.GetUsersOnlineStatusResp")
	proto.RegisterType((*GetUsersOnlineStatusResp_SuccessDetail)(nil), "relay.GetUsersOnlineStatusResp.SuccessDetail")
	proto.RegisterType((*GetUsersOnlineStatusResp_FailedDetail)(nil), "relay.GetUsersOnlineStatusResp.FailedDetail")
	proto.RegisterType((*GetUsersOnlineStatusResp_SuccessResult)(nil), "relay.GetUsersOnlineStatusResp.SuccessResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OnlineMessageRelayService service

type OnlineMessageRelayServiceClient interface {
	OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error)
	GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error)
}

type onlineMessageRelayServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnlineMessageRelayServiceClient(cc *grpc.ClientConn) OnlineMessageRelayServiceClient {
	return &onlineMessageRelayServiceClient{cc}
}

func (c *onlineMessageRelayServiceClient) OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error) {
	out := new(OnlinePushMsgResp)
	err := grpc.Invoke(ctx, "/relay.OnlineMessageRelayService/OnlinePushMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineMessageRelayServiceClient) GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error) {
	out := new(GetUsersOnlineStatusResp)
	err := grpc.Invoke(ctx, "/relay.OnlineMessageRelayService/GetUsersOnlineStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OnlineMessageRelayService service

type OnlineMessageRelayServiceServer interface {
	OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error)
	GetUsersOnlineStatus(context.Context, *GetUsersOnlineStatusReq) (*GetUsersOnlineStatusResp, error)
}

func RegisterOnlineMessageRelayServiceServer(s *grpc.Server, srv OnlineMessageRelayServiceServer) {
	s.RegisterService(&_OnlineMessageRelayService_serviceDesc, srv)
}

func _OnlineMessageRelayService_OnlinePushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlinePushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineMessageRelayServiceServer).OnlinePushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relay.OnlineMessageRelayService/OnlinePushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineMessageRelayServiceServer).OnlinePushMsg(ctx, req.(*OnlinePushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineMessageRelayService_GetUsersOnlineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersOnlineStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineMessageRelayServiceServer).GetUsersOnlineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relay.OnlineMessageRelayService/GetUsersOnlineStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineMessageRelayServiceServer).GetUsersOnlineStatus(ctx, req.(*GetUsersOnlineStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlineMessageRelayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "relay.OnlineMessageRelayService",
	HandlerType: (*OnlineMessageRelayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlinePushMsg",
			Handler:    _OnlineMessageRelayService_OnlinePushMsg_Handler,
		},
		{
			MethodName: "GetUsersOnlineStatus",
			Handler:    _OnlineMessageRelayService_GetUsersOnlineStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay/relay.proto",
}

func init() { proto.RegisterFile("relay/relay.proto", fileDescriptor_relay_34094e5333f6005a) }

var fileDescriptor_relay_34094e5333f6005a = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0x4d, 0x65, 0x3f, 0xdc, 0x0b, 0xb8, 0x32, 0xd9, 0xec, 0xd6, 0x3e, 0x20, 0xf6, 0xc1, 0x10,
	0xa3, 0x25, 0x41, 0xdf, 0x7c, 0x30, 0xd9, 0x25, 0x6b, 0x48, 0x6c, 0x20, 0x83, 0x46, 0xe3, 0x0b,
	0x99, 0x85, 0xbb, 0xdd, 0x66, 0x0b, 0x1d, 0xe6, 0xb6, 0x10, 0xff, 0x84, 0x3f, 0xc2, 0x3f, 0xa1,
	0x3f, 0xcf, 0x74, 0xa6, 0x60, 0x4b, 0x58, 0x37, 0xfb, 0x42, 0x38, 0x77, 0xee, 0x3d, 0xf7, 0x9c,
	0xd3, 0x76, 0xa0, 0xa1, 0x30, 0x12, 0x3f, 0x3a, 0xfa, 0xd7, 0x93, 0x2a, 0x4e, 0x62, 0xb6, 0xaf,
	0x81, 0xf3, 0x62, 0x20, 0x71, 0x3e, 0xee, 0xfb, 0x1d, 0x79, 0x1b, 0x74, 0xf4, 0x49, 0x87, 0xa6,
	0xb7, 0xe3, 0x15, 0x75, 0x56, 0x64, 0x3a, 0xdd, 0x9f, 0x16, 0x3c, 0x1d, 0xcc, 0xa3, 0x70, 0x8e,
	0xc3, 0x94, 0x6e, 0x7c, 0x0a, 0x38, 0x2e, 0x58, 0x0b, 0xaa, 0x03, 0x89, 0x4a, 0x24, 0x61, 0x3c,
	0xef, 0xf7, 0x6c, 0xab, 0x65, 0xb5, 0x8f, 0x78, 0xb1, 0xc4, 0xde, 0xc1, 0xe1, 0x8c, 0x82, 0x9e,
	0x48, 0x84, 0xfd, 0xa8, 0x65, 0xb5, 0xab, 0x5d, 0xc7, 0x23, 0x54, 0x4b, 0x54, 0x63, 0x21, 0xc3,
	0xb1, 0x14, 0x4a, 0xcc, 0xc8, 0xf3, 0x4d, 0x07, 0x5f, 0xb7, 0x32, 0x17, 0x6a, 0x32, 0xa5, 0x9b,
	0xcf, 0xf1, 0x17, 0x42, 0xd5, 0xef, 0xd9, 0x15, 0x4d, 0x5c, 0xaa, 0xb9, 0x1f, 0xa0, 0xb1, 0xa5,
	0x87, 0x24, 0x7b, 0x05, 0x7b, 0x0a, 0x49, 0xda, 0x56, 0xab, 0xd2, 0xae, 0x76, 0x4f, 0x3d, 0xe3,
	0x75, 0x14, 0xce, 0x83, 0x08, 0x7d, 0x0a, 0xcc, 0x30, 0xd7, 0x3d, 0xee, 0x02, 0x8e, 0xb7, 0x0e,
	0x58, 0x13, 0x80, 0x23, 0xa5, 0x51, 0x72, 0x11, 0x4f, 0x51, 0xdb, 0xa9, 0xf0, 0x42, 0x85, 0x9d,
	0xc2, 0x01, 0xc7, 0xc9, 0xb2, 0xdf, 0xd3, 0x66, 0x8e, 0x78, 0x8e, 0xd8, 0x4b, 0x78, 0x92, 0xfd,
	0x1b, 0x46, 0x22, 0xb9, 0x8c, 0xd5, 0x2c, 0x57, 0xbc, 0xcf, 0xb7, 0xaa, 0xee, 0x0a, 0xce, 0x3e,
	0x62, 0x92, 0xad, 0x22, 0xa3, 0x7d, 0x94, 0x88, 0x24, 0xa5, 0x2c, 0xca, 0x26, 0x40, 0xaa, 0x8d,
	0x7d, 0x0a, 0x29, 0xd1, 0xfa, 0x8f, 0x78, 0xa1, 0x92, 0x45, 0x1d, 0x17, 0xa2, 0x36, 0xfb, 0x8b,
	0x25, 0xe6, 0xc0, 0xe3, 0x58, 0x96, 0x02, 0xdb, 0x60, 0xf7, 0xf7, 0x1e, 0xd8, 0xbb, 0x37, 0x93,
	0x64, 0x36, 0x1c, 0xa2, 0x52, 0x1b, 0xcb, 0xfb, 0x7c, 0x0d, 0x33, 0xbf, 0xa8, 0x94, 0x4f, 0xc1,
	0xda, 0xaf, 0x41, 0x6c, 0x04, 0x75, 0x4a, 0x27, 0x13, 0x24, 0x32, 0xe1, 0xd8, 0x15, 0x9d, 0xf7,
	0x9b, 0x3c, 0xef, 0xbb, 0x36, 0x79, 0xa3, 0xe2, 0x10, 0x2f, 0x73, 0xb0, 0x21, 0xd4, 0xae, 0x45,
	0x18, 0xe1, 0x34, 0xe7, 0xdc, 0xd3, 0x9c, 0xaf, 0xef, 0xe3, 0xbc, 0xd4, 0x33, 0x3d, 0x4c, 0x44,
	0x18, 0xf1, 0x12, 0x83, 0x73, 0x01, 0xf5, 0x7c, 0xa3, 0x39, 0xce, 0x22, 0x92, 0x91, 0x48, 0xae,
	0x63, 0x35, 0xcb, 0x5f, 0xd6, 0x0d, 0xce, 0xbc, 0x92, 0x66, 0x5d, 0x7b, 0x35, 0xc8, 0xf9, 0x06,
	0xb5, 0xe2, 0x8a, 0xac, 0x2f, 0x2d, 0x86, 0x9c, 0xa3, 0x87, 0xa7, 0xe8, 0xfc, 0xb2, 0x36, 0xfa,
	0xf2, 0x08, 0xfe, 0x71, 0x5b, 0x25, 0xee, 0x3b, 0xb4, 0x31, 0x01, 0x27, 0x53, 0xad, 0x6a, 0x98,
	0xbb, 0x30, 0xb9, 0x3c, 0xf0, 0x71, 0xe4, 0xd9, 0xed, 0xa4, 0xea, 0xfe, 0xb1, 0xe0, 0x99, 0x19,
	0xf4, 0x91, 0x48, 0x04, 0xc8, 0x33, 0xce, 0x11, 0xaa, 0x65, 0x38, 0x41, 0x76, 0x0e, 0xf5, 0xd2,
	0x47, 0xc8, 0xce, 0xf2, 0x9d, 0xdb, 0x57, 0x85, 0x63, 0xef, 0x3e, 0x20, 0xc9, 0xbe, 0xc2, 0xc9,
	0x2e, 0x85, 0xac, 0xf9, 0x5f, 0xf9, 0x0b, 0xe7, 0xf9, 0x3d, 0xf6, 0xce, 0x1b, 0xdf, 0x8f, 0x3d,
	0x73, 0xdb, 0xbd, 0x97, 0x57, 0x5a, 0xf6, 0xd5, 0x81, 0xbe, 0xcc, 0xde, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0xdc, 0xcc, 0x70, 0x0b, 0x05, 0x00, 0x00,
}
