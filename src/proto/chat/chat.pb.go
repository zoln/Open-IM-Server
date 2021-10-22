// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/chat.proto

package pbChat // import "./chat"

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

type WSToMsgSvrChatMsg struct {
	SendID               string   `protobuf:"bytes,1,opt,name=SendID" json:"SendID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID" json:"RecvID,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content" json:"Content,omitempty"`
	SendTime             int64    `protobuf:"varint,4,opt,name=SendTime" json:"SendTime,omitempty"`
	MsgFrom              int32    `protobuf:"varint,5,opt,name=MsgFrom" json:"MsgFrom,omitempty"`
	SenderNickName       string   `protobuf:"bytes,6,opt,name=SenderNickName" json:"SenderNickName,omitempty"`
	SenderFaceURL        string   `protobuf:"bytes,7,opt,name=SenderFaceURL" json:"SenderFaceURL,omitempty"`
	ContentType          int32    `protobuf:"varint,8,opt,name=ContentType" json:"ContentType,omitempty"`
	SessionType          int32    `protobuf:"varint,9,opt,name=SessionType" json:"SessionType,omitempty"`
	OperationID          string   `protobuf:"bytes,10,opt,name=OperationID" json:"OperationID,omitempty"`
	MsgID                string   `protobuf:"bytes,11,opt,name=MsgID" json:"MsgID,omitempty"`
	Token                string   `protobuf:"bytes,12,opt,name=Token" json:"Token,omitempty"`
	OfflineInfo          string   `protobuf:"bytes,13,opt,name=OfflineInfo" json:"OfflineInfo,omitempty"`
	Options              string   `protobuf:"bytes,14,opt,name=Options" json:"Options,omitempty"`
	PlatformID           int32    `protobuf:"varint,15,opt,name=PlatformID" json:"PlatformID,omitempty"`
	ForceList            []string `protobuf:"bytes,16,rep,name=ForceList" json:"ForceList,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,17,opt,name=ClientMsgID" json:"ClientMsgID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WSToMsgSvrChatMsg) Reset()         { *m = WSToMsgSvrChatMsg{} }
func (m *WSToMsgSvrChatMsg) String() string { return proto.CompactTextString(m) }
func (*WSToMsgSvrChatMsg) ProtoMessage()    {}
func (*WSToMsgSvrChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{0}
}
func (m *WSToMsgSvrChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Unmarshal(m, b)
}
func (m *WSToMsgSvrChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Marshal(b, m, deterministic)
}
func (dst *WSToMsgSvrChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WSToMsgSvrChatMsg.Merge(dst, src)
}
func (m *WSToMsgSvrChatMsg) XXX_Size() int {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Size(m)
}
func (m *WSToMsgSvrChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WSToMsgSvrChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WSToMsgSvrChatMsg proto.InternalMessageInfo

func (m *WSToMsgSvrChatMsg) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetSenderNickName() string {
	if m != nil {
		return m.SenderNickName
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetSenderFaceURL() string {
	if m != nil {
		return m.SenderFaceURL
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetOfflineInfo() string {
	if m != nil {
		return m.OfflineInfo
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetForceList() []string {
	if m != nil {
		return m.ForceList
	}
	return nil
}

func (m *WSToMsgSvrChatMsg) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

type MsgSvrToPushSvrChatMsg struct {
	SendID               string   `protobuf:"bytes,1,opt,name=SendID" json:"SendID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID" json:"RecvID,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content" json:"Content,omitempty"`
	RecvSeq              int64    `protobuf:"varint,4,opt,name=RecvSeq" json:"RecvSeq,omitempty"`
	SendTime             int64    `protobuf:"varint,5,opt,name=SendTime" json:"SendTime,omitempty"`
	MsgFrom              int32    `protobuf:"varint,6,opt,name=MsgFrom" json:"MsgFrom,omitempty"`
	SenderNickName       string   `protobuf:"bytes,7,opt,name=SenderNickName" json:"SenderNickName,omitempty"`
	SenderFaceURL        string   `protobuf:"bytes,8,opt,name=SenderFaceURL" json:"SenderFaceURL,omitempty"`
	ContentType          int32    `protobuf:"varint,9,opt,name=ContentType" json:"ContentType,omitempty"`
	SessionType          int32    `protobuf:"varint,10,opt,name=SessionType" json:"SessionType,omitempty"`
	OperationID          string   `protobuf:"bytes,11,opt,name=OperationID" json:"OperationID,omitempty"`
	MsgID                string   `protobuf:"bytes,12,opt,name=MsgID" json:"MsgID,omitempty"`
	OfflineInfo          string   `protobuf:"bytes,13,opt,name=OfflineInfo" json:"OfflineInfo,omitempty"`
	Options              string   `protobuf:"bytes,14,opt,name=Options" json:"Options,omitempty"`
	PlatformID           int32    `protobuf:"varint,15,opt,name=PlatformID" json:"PlatformID,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,16,opt,name=ClientMsgID" json:"ClientMsgID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSvrToPushSvrChatMsg) Reset()         { *m = MsgSvrToPushSvrChatMsg{} }
func (m *MsgSvrToPushSvrChatMsg) String() string { return proto.CompactTextString(m) }
func (*MsgSvrToPushSvrChatMsg) ProtoMessage()    {}
func (*MsgSvrToPushSvrChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{1}
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Unmarshal(m, b)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Marshal(b, m, deterministic)
}
func (dst *MsgSvrToPushSvrChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSvrToPushSvrChatMsg.Merge(dst, src)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Size() int {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Size(m)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSvrToPushSvrChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSvrToPushSvrChatMsg proto.InternalMessageInfo

func (m *MsgSvrToPushSvrChatMsg) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetRecvSeq() int64 {
	if m != nil {
		return m.RecvSeq
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetSenderNickName() string {
	if m != nil {
		return m.SenderNickName
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetSenderFaceURL() string {
	if m != nil {
		return m.SenderFaceURL
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetOfflineInfo() string {
	if m != nil {
		return m.OfflineInfo
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

type PullMessageReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	SeqBegin             int64    `protobuf:"varint,2,opt,name=SeqBegin" json:"SeqBegin,omitempty"`
	SeqEnd               int64    `protobuf:"varint,3,opt,name=SeqEnd" json:"SeqEnd,omitempty"`
	OperationID          string   `protobuf:"bytes,4,opt,name=OperationID" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullMessageReq) Reset()         { *m = PullMessageReq{} }
func (m *PullMessageReq) String() string { return proto.CompactTextString(m) }
func (*PullMessageReq) ProtoMessage()    {}
func (*PullMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{2}
}
func (m *PullMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageReq.Unmarshal(m, b)
}
func (m *PullMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageReq.Marshal(b, m, deterministic)
}
func (dst *PullMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageReq.Merge(dst, src)
}
func (m *PullMessageReq) XXX_Size() int {
	return xxx_messageInfo_PullMessageReq.Size(m)
}
func (m *PullMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageReq proto.InternalMessageInfo

func (m *PullMessageReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PullMessageReq) GetSeqBegin() int64 {
	if m != nil {
		return m.SeqBegin
	}
	return 0
}

func (m *PullMessageReq) GetSeqEnd() int64 {
	if m != nil {
		return m.SeqEnd
	}
	return 0
}

func (m *PullMessageReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type PullMessageResp struct {
	ErrCode              int32           `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string          `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	MaxSeq               int64           `protobuf:"varint,3,opt,name=MaxSeq" json:"MaxSeq,omitempty"`
	MinSeq               int64           `protobuf:"varint,4,opt,name=MinSeq" json:"MinSeq,omitempty"`
	SingleUserMsg        []*GatherFormat `protobuf:"bytes,5,rep,name=SingleUserMsg" json:"SingleUserMsg,omitempty"`
	GroupUserMsg         []*GatherFormat `protobuf:"bytes,6,rep,name=GroupUserMsg" json:"GroupUserMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PullMessageResp) Reset()         { *m = PullMessageResp{} }
func (m *PullMessageResp) String() string { return proto.CompactTextString(m) }
func (*PullMessageResp) ProtoMessage()    {}
func (*PullMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{3}
}
func (m *PullMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageResp.Unmarshal(m, b)
}
func (m *PullMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageResp.Marshal(b, m, deterministic)
}
func (dst *PullMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageResp.Merge(dst, src)
}
func (m *PullMessageResp) XXX_Size() int {
	return xxx_messageInfo_PullMessageResp.Size(m)
}
func (m *PullMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageResp proto.InternalMessageInfo

func (m *PullMessageResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *PullMessageResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PullMessageResp) GetMaxSeq() int64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

func (m *PullMessageResp) GetMinSeq() int64 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

func (m *PullMessageResp) GetSingleUserMsg() []*GatherFormat {
	if m != nil {
		return m.SingleUserMsg
	}
	return nil
}

func (m *PullMessageResp) GetGroupUserMsg() []*GatherFormat {
	if m != nil {
		return m.GroupUserMsg
	}
	return nil
}

type PullMessageBySeqListReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	OperationID          string   `protobuf:"bytes,2,opt,name=OperationID" json:"OperationID,omitempty"`
	SeqList              []int64  `protobuf:"varint,3,rep,packed,name=seqList" json:"seqList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullMessageBySeqListReq) Reset()         { *m = PullMessageBySeqListReq{} }
func (m *PullMessageBySeqListReq) String() string { return proto.CompactTextString(m) }
func (*PullMessageBySeqListReq) ProtoMessage()    {}
func (*PullMessageBySeqListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{4}
}
func (m *PullMessageBySeqListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageBySeqListReq.Unmarshal(m, b)
}
func (m *PullMessageBySeqListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageBySeqListReq.Marshal(b, m, deterministic)
}
func (dst *PullMessageBySeqListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageBySeqListReq.Merge(dst, src)
}
func (m *PullMessageBySeqListReq) XXX_Size() int {
	return xxx_messageInfo_PullMessageBySeqListReq.Size(m)
}
func (m *PullMessageBySeqListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageBySeqListReq.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageBySeqListReq proto.InternalMessageInfo

func (m *PullMessageBySeqListReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PullMessageBySeqListReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *PullMessageBySeqListReq) GetSeqList() []int64 {
	if m != nil {
		return m.SeqList
	}
	return nil
}

type GetNewSeqReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	OperationID          string   `protobuf:"bytes,2,opt,name=OperationID" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewSeqReq) Reset()         { *m = GetNewSeqReq{} }
func (m *GetNewSeqReq) String() string { return proto.CompactTextString(m) }
func (*GetNewSeqReq) ProtoMessage()    {}
func (*GetNewSeqReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{5}
}
func (m *GetNewSeqReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewSeqReq.Unmarshal(m, b)
}
func (m *GetNewSeqReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewSeqReq.Marshal(b, m, deterministic)
}
func (dst *GetNewSeqReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewSeqReq.Merge(dst, src)
}
func (m *GetNewSeqReq) XXX_Size() int {
	return xxx_messageInfo_GetNewSeqReq.Size(m)
}
func (m *GetNewSeqReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewSeqReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewSeqReq proto.InternalMessageInfo

func (m *GetNewSeqReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetNewSeqReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type GetNewSeqResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	Seq                  int64    `protobuf:"varint,3,opt,name=Seq" json:"Seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewSeqResp) Reset()         { *m = GetNewSeqResp{} }
func (m *GetNewSeqResp) String() string { return proto.CompactTextString(m) }
func (*GetNewSeqResp) ProtoMessage()    {}
func (*GetNewSeqResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{6}
}
func (m *GetNewSeqResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewSeqResp.Unmarshal(m, b)
}
func (m *GetNewSeqResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewSeqResp.Marshal(b, m, deterministic)
}
func (dst *GetNewSeqResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewSeqResp.Merge(dst, src)
}
func (m *GetNewSeqResp) XXX_Size() int {
	return xxx_messageInfo_GetNewSeqResp.Size(m)
}
func (m *GetNewSeqResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewSeqResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewSeqResp proto.InternalMessageInfo

func (m *GetNewSeqResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GetNewSeqResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetNewSeqResp) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type GatherFormat struct {
	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"id"`
	// @inject_tag: json:"list"
	List                 []*MsgFormat `protobuf:"bytes,2,rep,name=List" json:"list"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GatherFormat) Reset()         { *m = GatherFormat{} }
func (m *GatherFormat) String() string { return proto.CompactTextString(m) }
func (*GatherFormat) ProtoMessage()    {}
func (*GatherFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{7}
}
func (m *GatherFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatherFormat.Unmarshal(m, b)
}
func (m *GatherFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatherFormat.Marshal(b, m, deterministic)
}
func (dst *GatherFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatherFormat.Merge(dst, src)
}
func (m *GatherFormat) XXX_Size() int {
	return xxx_messageInfo_GatherFormat.Size(m)
}
func (m *GatherFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_GatherFormat.DiscardUnknown(m)
}

var xxx_messageInfo_GatherFormat proto.InternalMessageInfo

func (m *GatherFormat) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GatherFormat) GetList() []*MsgFormat {
	if m != nil {
		return m.List
	}
	return nil
}

type MsgFormat struct {
	// @inject_tag: json:"sendID"
	SendID string `protobuf:"bytes,1,opt,name=SendID" json:"sendID"`
	// @inject_tag: json:"recvID"
	RecvID string `protobuf:"bytes,2,opt,name=RecvID" json:"recvID"`
	// @inject_tag: json:"msgFrom"
	MsgFrom int32 `protobuf:"varint,3,opt,name=MsgFrom" json:"msgFrom"`
	// @inject_tag: json:"contentType"
	ContentType int32 `protobuf:"varint,4,opt,name=ContentType" json:"contentType"`
	// @inject_tag: json:"serverMsgID"
	ServerMsgID string `protobuf:"bytes,5,opt,name=ServerMsgID" json:"serverMsgID"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,6,opt,name=Content" json:"content"`
	// @inject_tag: json:"seq"
	Seq int64 `protobuf:"varint,7,opt,name=Seq" json:"seq"`
	// @inject_tag: json:"sendTime"
	SendTime int64 `protobuf:"varint,8,opt,name=SendTime" json:"sendTime"`
	// @inject_tag: json:"senderPlatformID"
	SenderPlatformID int32 `protobuf:"varint,9,opt,name=SenderPlatformID" json:"senderPlatformID"`
	// @inject_tag: json:"senderNickName"
	SenderNickName string `protobuf:"bytes,10,opt,name=SenderNickName" json:"senderNickName"`
	// @inject_tag: json:"senderFaceUrl"
	SenderFaceURL string `protobuf:"bytes,11,opt,name=SenderFaceURL" json:"senderFaceUrl"`
	// @inject_tag: json:"clientMsgID"
	ClientMsgID          string   `protobuf:"bytes,12,opt,name=ClientMsgID" json:"clientMsgID"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgFormat) Reset()         { *m = MsgFormat{} }
func (m *MsgFormat) String() string { return proto.CompactTextString(m) }
func (*MsgFormat) ProtoMessage()    {}
func (*MsgFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{8}
}
func (m *MsgFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgFormat.Unmarshal(m, b)
}
func (m *MsgFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgFormat.Marshal(b, m, deterministic)
}
func (dst *MsgFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFormat.Merge(dst, src)
}
func (m *MsgFormat) XXX_Size() int {
	return xxx_messageInfo_MsgFormat.Size(m)
}
func (m *MsgFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFormat.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFormat proto.InternalMessageInfo

func (m *MsgFormat) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *MsgFormat) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *MsgFormat) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *MsgFormat) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *MsgFormat) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *MsgFormat) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgFormat) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MsgFormat) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *MsgFormat) GetSenderPlatformID() int32 {
	if m != nil {
		return m.SenderPlatformID
	}
	return 0
}

func (m *MsgFormat) GetSenderNickName() string {
	if m != nil {
		return m.SenderNickName
	}
	return ""
}

func (m *MsgFormat) GetSenderFaceURL() string {
	if m != nil {
		return m.SenderFaceURL
	}
	return ""
}

func (m *MsgFormat) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

type UserSendMsgReq struct {
	ReqIdentifier        int32    `protobuf:"varint,1,opt,name=ReqIdentifier" json:"ReqIdentifier,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
	SendID               string   `protobuf:"bytes,3,opt,name=SendID" json:"SendID,omitempty"`
	OperationID          string   `protobuf:"bytes,4,opt,name=OperationID" json:"OperationID,omitempty"`
	SenderNickName       string   `protobuf:"bytes,5,opt,name=SenderNickName" json:"SenderNickName,omitempty"`
	SenderFaceURL        string   `protobuf:"bytes,6,opt,name=SenderFaceURL" json:"SenderFaceURL,omitempty"`
	PlatformID           int32    `protobuf:"varint,7,opt,name=PlatformID" json:"PlatformID,omitempty"`
	SessionType          int32    `protobuf:"varint,8,opt,name=SessionType" json:"SessionType,omitempty"`
	MsgFrom              int32    `protobuf:"varint,9,opt,name=MsgFrom" json:"MsgFrom,omitempty"`
	ContentType          int32    `protobuf:"varint,10,opt,name=ContentType" json:"ContentType,omitempty"`
	RecvID               string   `protobuf:"bytes,11,opt,name=RecvID" json:"RecvID,omitempty"`
	ForceList            []string `protobuf:"bytes,12,rep,name=ForceList" json:"ForceList,omitempty"`
	Content              string   `protobuf:"bytes,13,opt,name=Content" json:"Content,omitempty"`
	Options              string   `protobuf:"bytes,14,opt,name=Options" json:"Options,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,15,opt,name=ClientMsgID" json:"ClientMsgID,omitempty"`
	OffLineInfo          string   `protobuf:"bytes,16,opt,name=OffLineInfo" json:"OffLineInfo,omitempty"`
	Ex                   string   `protobuf:"bytes,17,opt,name=Ex" json:"Ex,omitempty"`
	SendTime             int64    `protobuf:"varint,18,opt,name=sendTime" json:"sendTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSendMsgReq) Reset()         { *m = UserSendMsgReq{} }
func (m *UserSendMsgReq) String() string { return proto.CompactTextString(m) }
func (*UserSendMsgReq) ProtoMessage()    {}
func (*UserSendMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{9}
}
func (m *UserSendMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSendMsgReq.Unmarshal(m, b)
}
func (m *UserSendMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSendMsgReq.Marshal(b, m, deterministic)
}
func (dst *UserSendMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSendMsgReq.Merge(dst, src)
}
func (m *UserSendMsgReq) XXX_Size() int {
	return xxx_messageInfo_UserSendMsgReq.Size(m)
}
func (m *UserSendMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSendMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserSendMsgReq proto.InternalMessageInfo

func (m *UserSendMsgReq) GetReqIdentifier() int32 {
	if m != nil {
		return m.ReqIdentifier
	}
	return 0
}

func (m *UserSendMsgReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserSendMsgReq) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *UserSendMsgReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *UserSendMsgReq) GetSenderNickName() string {
	if m != nil {
		return m.SenderNickName
	}
	return ""
}

func (m *UserSendMsgReq) GetSenderFaceURL() string {
	if m != nil {
		return m.SenderFaceURL
	}
	return ""
}

func (m *UserSendMsgReq) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *UserSendMsgReq) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *UserSendMsgReq) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *UserSendMsgReq) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *UserSendMsgReq) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *UserSendMsgReq) GetForceList() []string {
	if m != nil {
		return m.ForceList
	}
	return nil
}

func (m *UserSendMsgReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UserSendMsgReq) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *UserSendMsgReq) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

func (m *UserSendMsgReq) GetOffLineInfo() string {
	if m != nil {
		return m.OffLineInfo
	}
	return ""
}

func (m *UserSendMsgReq) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

func (m *UserSendMsgReq) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type UserSendMsgResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	ReqIdentifier        int32    `protobuf:"varint,3,opt,name=ReqIdentifier" json:"ReqIdentifier,omitempty"`
	ServerMsgID          string   `protobuf:"bytes,4,opt,name=ServerMsgID" json:"ServerMsgID,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,5,opt,name=ClientMsgID" json:"ClientMsgID,omitempty"`
	SendTime             int64    `protobuf:"varint,6,opt,name=sendTime" json:"sendTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSendMsgResp) Reset()         { *m = UserSendMsgResp{} }
func (m *UserSendMsgResp) String() string { return proto.CompactTextString(m) }
func (*UserSendMsgResp) ProtoMessage()    {}
func (*UserSendMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1c5fb32d3659e154, []int{10}
}
func (m *UserSendMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSendMsgResp.Unmarshal(m, b)
}
func (m *UserSendMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSendMsgResp.Marshal(b, m, deterministic)
}
func (dst *UserSendMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSendMsgResp.Merge(dst, src)
}
func (m *UserSendMsgResp) XXX_Size() int {
	return xxx_messageInfo_UserSendMsgResp.Size(m)
}
func (m *UserSendMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSendMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSendMsgResp proto.InternalMessageInfo

func (m *UserSendMsgResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *UserSendMsgResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserSendMsgResp) GetReqIdentifier() int32 {
	if m != nil {
		return m.ReqIdentifier
	}
	return 0
}

func (m *UserSendMsgResp) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *UserSendMsgResp) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

func (m *UserSendMsgResp) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func init() {
	proto.RegisterType((*WSToMsgSvrChatMsg)(nil), "pbChat.WSToMsgSvrChatMsg")
	proto.RegisterType((*MsgSvrToPushSvrChatMsg)(nil), "pbChat.MsgSvrToPushSvrChatMsg")
	proto.RegisterType((*PullMessageReq)(nil), "pbChat.PullMessageReq")
	proto.RegisterType((*PullMessageResp)(nil), "pbChat.PullMessageResp")
	proto.RegisterType((*PullMessageBySeqListReq)(nil), "pbChat.PullMessageBySeqListReq")
	proto.RegisterType((*GetNewSeqReq)(nil), "pbChat.GetNewSeqReq")
	proto.RegisterType((*GetNewSeqResp)(nil), "pbChat.GetNewSeqResp")
	proto.RegisterType((*GatherFormat)(nil), "pbChat.GatherFormat")
	proto.RegisterType((*MsgFormat)(nil), "pbChat.MsgFormat")
	proto.RegisterType((*UserSendMsgReq)(nil), "pbChat.UserSendMsgReq")
	proto.RegisterType((*UserSendMsgResp)(nil), "pbChat.UserSendMsgResp")
}

func init() { proto.RegisterFile("chat/chat.proto", fileDescriptor_chat_1c5fb32d3659e154) }

var fileDescriptor_chat_1c5fb32d3659e154 = []byte{
	// 978 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x6a, 0xe3, 0x46,
	0x14, 0xc6, 0x92, 0x2d, 0x47, 0xc7, 0xbf, 0x19, 0xd2, 0xec, 0x10, 0x4a, 0x6b, 0xc4, 0xb6, 0x98,
	0x5e, 0xa4, 0xb0, 0xbd, 0x59, 0x5a, 0xd8, 0x8b, 0xc4, 0x4e, 0x2a, 0x88, 0xb3, 0x41, 0xf2, 0x52,
	0xe8, 0x9d, 0xd6, 0x19, 0x2b, 0x62, 0x6d, 0x49, 0xd6, 0x28, 0x69, 0xf6, 0xa6, 0xaf, 0x50, 0xe8,
	0x73, 0xf4, 0x45, 0xfa, 0x28, 0x7d, 0x85, 0x42, 0xa1, 0xcc, 0x8c, 0x64, 0xcd, 0x48, 0x72, 0x12,
	0xb2, 0xec, 0xcd, 0xb2, 0xdf, 0xa7, 0x33, 0x3f, 0xe7, 0x9c, 0x6f, 0xbe, 0x13, 0xc3, 0x60, 0x71,
	0xe3, 0xa5, 0xdf, 0xb3, 0x7f, 0x8e, 0xe3, 0x24, 0x4a, 0x23, 0x64, 0xc4, 0xef, 0x4f, 0x6f, 0xbc,
	0xd4, 0xfa, 0xa3, 0x09, 0xfb, 0xbf, 0xb8, 0xf3, 0x68, 0x46, 0x7d, 0xf7, 0x2e, 0x61, 0xd4, 0x8c,
	0xfa, 0xe8, 0x10, 0x0c, 0x97, 0x84, 0xd7, 0xf6, 0x04, 0x37, 0x46, 0x8d, 0xb1, 0xe9, 0x64, 0x88,
	0xf1, 0x0e, 0x59, 0xdc, 0xd9, 0x13, 0xac, 0x09, 0x5e, 0x20, 0x84, 0xa1, 0x7d, 0x1a, 0x85, 0x29,
	0x09, 0x53, 0xac, 0xf3, 0x0f, 0x39, 0x44, 0x47, 0xb0, 0xc7, 0xd6, 0xce, 0x83, 0x35, 0xc1, 0xcd,
	0x51, 0x63, 0xac, 0x3b, 0x5b, 0xcc, 0x56, 0xcd, 0xa8, 0x7f, 0x96, 0x44, 0x6b, 0xdc, 0x1a, 0x35,
	0xc6, 0x2d, 0x27, 0x87, 0xe8, 0x5b, 0xe8, 0xb3, 0x28, 0x92, 0x5c, 0x06, 0x8b, 0x0f, 0x97, 0xde,
	0x9a, 0x60, 0x83, 0x6f, 0x5b, 0x62, 0xd1, 0x4b, 0xe8, 0x09, 0xe6, 0xcc, 0x5b, 0x90, 0x77, 0xce,
	0x05, 0x6e, 0xf3, 0x30, 0x95, 0x44, 0x23, 0xe8, 0x64, 0xd7, 0x99, 0x7f, 0x8c, 0x09, 0xde, 0xe3,
	0x67, 0xc9, 0x14, 0x8b, 0x70, 0x09, 0xa5, 0x41, 0x14, 0xf2, 0x08, 0x53, 0x44, 0x48, 0x14, 0x8b,
	0x78, 0x1b, 0x93, 0xc4, 0x4b, 0x83, 0x28, 0xb4, 0x27, 0x18, 0xf8, 0x39, 0x32, 0x85, 0x0e, 0xa0,
	0x35, 0xa3, 0xbe, 0x3d, 0xc1, 0x1d, 0xfe, 0x4d, 0x00, 0xc6, 0xce, 0xa3, 0x0f, 0x24, 0xc4, 0x5d,
	0xc1, 0x72, 0xc0, 0x77, 0x5b, 0x2e, 0x57, 0x41, 0x48, 0xec, 0x70, 0x19, 0xe1, 0x5e, 0xb6, 0x5b,
	0x41, 0xb1, 0xda, 0xbc, 0x8d, 0xd9, 0xce, 0x14, 0xf7, 0x45, 0x45, 0x33, 0x88, 0xbe, 0x02, 0xb8,
	0x5a, 0x79, 0xe9, 0x32, 0x4a, 0xd6, 0xf6, 0x04, 0x0f, 0xf8, 0x55, 0x25, 0x06, 0x7d, 0x09, 0xe6,
	0x59, 0x94, 0x2c, 0xc8, 0x45, 0x40, 0x53, 0x3c, 0x1c, 0xe9, 0x63, 0xd3, 0x29, 0x08, 0x5e, 0x8b,
	0x55, 0x40, 0xc2, 0x54, 0xdc, 0x75, 0x5f, 0x9c, 0x2c, 0x51, 0xd6, 0xbf, 0x3a, 0x1c, 0x0a, 0x35,
	0xcc, 0xa3, 0xab, 0x5b, 0x7a, 0xf3, 0x59, 0x64, 0x81, 0xa1, 0xcd, 0x62, 0x5c, 0xb2, 0xc9, 0x54,
	0x91, 0x43, 0x45, 0x30, 0xad, 0xdd, 0x82, 0x31, 0x1e, 0x13, 0x4c, 0xfb, 0x69, 0x82, 0xd9, 0x7b,
	0x82, 0x60, 0xcc, 0x47, 0x05, 0x03, 0x8f, 0x0a, 0xa6, 0xf3, 0x80, 0x60, 0xba, 0xb2, 0x60, 0x3e,
	0xa7, 0x34, 0x4a, 0xcd, 0x1f, 0x56, 0x9b, 0xff, 0x3b, 0xf4, 0xaf, 0x6e, 0x57, 0xab, 0x19, 0xa1,
	0xd4, 0xf3, 0x89, 0x43, 0x36, 0xac, 0xb7, 0xef, 0x28, 0x49, 0x8a, 0x9e, 0x0b, 0x24, 0xfa, 0xb4,
	0x39, 0x21, 0x7e, 0x10, 0xf2, 0xae, 0xf3, 0x3e, 0x09, 0x2c, 0x74, 0xb2, 0x99, 0x86, 0xd7, 0xbc,
	0xed, 0xba, 0x93, 0xa1, 0x72, 0x4d, 0x9a, 0x95, 0x9a, 0x58, 0xff, 0x34, 0x60, 0xa0, 0x5c, 0x80,
	0xc6, 0x2c, 0xdf, 0x69, 0x92, 0x9c, 0x46, 0xd7, 0x84, 0x5f, 0xa1, 0xe5, 0xe4, 0x90, 0x9d, 0x33,
	0x4d, 0x92, 0x19, 0xf5, 0x73, 0xdd, 0x09, 0xc4, 0xf8, 0x99, 0x77, 0xcf, 0xc4, 0x95, 0x9d, 0x2f,
	0x10, 0xe7, 0x83, 0xb0, 0x10, 0x5d, 0x86, 0xd0, 0x8f, 0xd0, 0x73, 0x83, 0xd0, 0x5f, 0x11, 0x96,
	0x1b, 0xdb, 0xae, 0x35, 0xd2, 0xc7, 0x9d, 0x57, 0x07, 0xc7, 0xc2, 0x24, 0x8f, 0xcf, 0xbd, 0xf4,
	0x86, 0x24, 0x67, 0x51, 0xb2, 0xf6, 0x52, 0x47, 0x0d, 0x45, 0xaf, 0xa1, 0x7b, 0x9e, 0x44, 0xb7,
	0x71, 0xbe, 0xd4, 0x78, 0x60, 0xa9, 0x12, 0x69, 0xad, 0xe1, 0x85, 0x94, 0xea, 0xc9, 0x47, 0x97,
	0x6c, 0xd8, 0x13, 0x7d, 0xa8, 0xe8, 0xa5, 0x02, 0x6a, 0x55, 0x51, 0x61, 0x68, 0x53, 0xb1, 0x0f,
	0xd6, 0x47, 0x3a, 0x7b, 0x58, 0x19, 0xb4, 0x7e, 0x86, 0xee, 0x39, 0x49, 0x2f, 0xc9, 0x6f, 0x2e,
	0xd9, 0x7c, 0xd2, 0x19, 0x96, 0x0b, 0x3d, 0x69, 0xa7, 0x67, 0x75, 0x68, 0x08, 0x7a, 0xd1, 0x1e,
	0xf6, 0x5f, 0x6b, 0x0a, 0x5d, 0xb9, 0x56, 0xa8, 0x0f, 0xda, 0xf6, 0x6a, 0x9a, 0x3d, 0x41, 0xdf,
	0x40, 0x93, 0x67, 0xa5, 0xf1, 0xfa, 0xee, 0xe7, 0xf5, 0x65, 0x06, 0x20, 0x8a, 0xcb, 0x3f, 0x5b,
	0xff, 0x69, 0x60, 0x6e, 0xb9, 0xe7, 0x18, 0x56, 0x6e, 0x30, 0xba, 0x6a, 0x30, 0x25, 0x4b, 0x68,
	0xee, 0xb0, 0x84, 0xe4, 0x8e, 0xf7, 0xd6, 0x9e, 0x70, 0xef, 0x32, 0x1d, 0x99, 0x92, 0xed, 0xd0,
	0x50, 0xed, 0x30, 0x2b, 0x47, 0x7b, 0x5b, 0x0e, 0xc5, 0x06, 0xf7, 0x4a, 0x36, 0xf8, 0x1d, 0x0c,
	0x85, 0x5f, 0x49, 0x8f, 0x5d, 0x78, 0x54, 0x85, 0xaf, 0x31, 0x46, 0x78, 0x9a, 0x31, 0x76, 0x76,
	0x19, 0xa3, 0x64, 0x20, 0xdd, 0xaa, 0x81, 0xfc, 0xd5, 0x84, 0x3e, 0x13, 0x12, 0x5b, 0x37, 0xa3,
	0x3e, 0x13, 0xda, 0x4b, 0xe8, 0x39, 0x64, 0x63, 0x5f, 0x93, 0x30, 0x0d, 0x96, 0x01, 0x49, 0x32,
	0x8d, 0xa8, 0x64, 0x31, 0x28, 0x35, 0x79, 0x50, 0x16, 0x0d, 0xd4, 0x95, 0x06, 0x3e, 0xea, 0x24,
	0x35, 0x89, 0xb7, 0x9e, 0x96, 0xb8, 0x51, 0x97, 0xb8, 0xea, 0xac, 0xed, 0x3a, 0x67, 0x95, 0xe7,
	0xc1, 0x5e, 0x75, 0x1e, 0x48, 0xd2, 0x32, 0x1f, 0x94, 0x16, 0x54, 0xa5, 0x55, 0xc8, 0xb5, 0xa3,
	0xc8, 0x55, 0x19, 0xf5, 0xdd, 0xf2, 0xa8, 0x97, 0xe4, 0xd6, 0xab, 0x4c, 0xdf, 0x1d, 0x13, 0xa4,
	0xd4, 0xe0, 0x41, 0xa5, 0xc1, 0xd9, 0x7c, 0xba, 0xc8, 0xe7, 0xd3, 0x70, 0x3b, 0x9f, 0x72, 0x8a,
	0xbd, 0xdc, 0xe9, 0x7d, 0xf6, 0x97, 0x85, 0x36, 0xbd, 0x67, 0x52, 0xa6, 0xb9, 0x94, 0x91, 0x90,
	0x72, 0x8e, 0xad, 0xbf, 0x1b, 0x30, 0x50, 0xe4, 0xf2, 0x2c, 0x37, 0xa9, 0x28, 0x4c, 0xaf, 0x53,
	0x58, 0xe9, 0x81, 0x36, 0xab, 0x0f, 0xb4, 0x94, 0x7d, 0xab, 0x9a, 0xbd, 0x9c, 0x8b, 0xa1, 0xe6,
	0xf2, 0xea, 0x4f, 0x0d, 0x9a, 0xcc, 0x93, 0xd0, 0x6b, 0x30, 0xb7, 0xfe, 0x88, 0x8a, 0x49, 0x20,
	0x99, 0xef, 0xd1, 0x17, 0x35, 0x2c, 0x8d, 0xd1, 0x1b, 0xe8, 0x48, 0x23, 0x01, 0x1d, 0xe6, 0x51,
	0xea, 0x4c, 0x3e, 0x7a, 0x51, 0xcb, 0xd3, 0x18, 0x5d, 0xc1, 0x41, 0xdd, 0x48, 0x41, 0x5f, 0xd7,
	0x2c, 0x90, 0x07, 0xce, 0xee, 0x1d, 0xdf, 0x40, 0x47, 0xea, 0x4f, 0x71, 0x23, 0xf5, 0x8d, 0x17,
	0xeb, 0x4b, 0xcd, 0x3c, 0x19, 0xfc, 0xda, 0x3b, 0xe6, 0xbf, 0x3b, 0x7e, 0x12, 0x01, 0xef, 0x0d,
	0xfe, 0xfb, 0xe3, 0x87, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xc0, 0xa5, 0xf9, 0x92, 0x0c,
	0x00, 0x00,
}
