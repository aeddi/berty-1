// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bertymessenger.proto

package bertymessenger

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ParseDeepLink_Kind int32

const (
	ParseDeepLink_UnknownKind ParseDeepLink_Kind = 0
	ParseDeepLink_BertyID     ParseDeepLink_Kind = 1
)

var ParseDeepLink_Kind_name = map[int32]string{
	0: "UnknownKind",
	1: "BertyID",
}

var ParseDeepLink_Kind_value = map[string]int32{
	"UnknownKind": 0,
	"BertyID":     1,
}

func (x ParseDeepLink_Kind) String() string {
	return proto.EnumName(ParseDeepLink_Kind_name, int32(x))
}

func (ParseDeepLink_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{2, 0}
}

type InstanceShareableBertyID struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceShareableBertyID) Reset()         { *m = InstanceShareableBertyID{} }
func (m *InstanceShareableBertyID) String() string { return proto.CompactTextString(m) }
func (*InstanceShareableBertyID) ProtoMessage()    {}
func (*InstanceShareableBertyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{0}
}
func (m *InstanceShareableBertyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceShareableBertyID.Unmarshal(m, b)
}
func (m *InstanceShareableBertyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceShareableBertyID.Marshal(b, m, deterministic)
}
func (m *InstanceShareableBertyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceShareableBertyID.Merge(m, src)
}
func (m *InstanceShareableBertyID) XXX_Size() int {
	return xxx_messageInfo_InstanceShareableBertyID.Size(m)
}
func (m *InstanceShareableBertyID) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceShareableBertyID.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceShareableBertyID proto.InternalMessageInfo

type InstanceShareableBertyID_Request struct {
	// reset will regenerate a new link
	Reset_               bool     `protobuf:"varint,1,opt,name=reset,proto3" json:"reset,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceShareableBertyID_Request) Reset()         { *m = InstanceShareableBertyID_Request{} }
func (m *InstanceShareableBertyID_Request) String() string { return proto.CompactTextString(m) }
func (*InstanceShareableBertyID_Request) ProtoMessage()    {}
func (*InstanceShareableBertyID_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{0, 0}
}
func (m *InstanceShareableBertyID_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceShareableBertyID_Request.Unmarshal(m, b)
}
func (m *InstanceShareableBertyID_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceShareableBertyID_Request.Marshal(b, m, deterministic)
}
func (m *InstanceShareableBertyID_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceShareableBertyID_Request.Merge(m, src)
}
func (m *InstanceShareableBertyID_Request) XXX_Size() int {
	return xxx_messageInfo_InstanceShareableBertyID_Request.Size(m)
}
func (m *InstanceShareableBertyID_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceShareableBertyID_Request.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceShareableBertyID_Request proto.InternalMessageInfo

func (m *InstanceShareableBertyID_Request) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

func (m *InstanceShareableBertyID_Request) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type InstanceShareableBertyID_Reply struct {
	BertyID              *BertyID `protobuf:"bytes,1,opt,name=berty_id,json=bertyId,proto3" json:"berty_id,omitempty"`
	BertyIDPayload       string   `protobuf:"bytes,2,opt,name=berty_id_payload,json=bertyIdPayload,proto3" json:"berty_id_payload,omitempty"`
	DeepLink             string   `protobuf:"bytes,3,opt,name=deep_link,json=deepLink,proto3" json:"deep_link,omitempty"`
	HTMLURL              string   `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceShareableBertyID_Reply) Reset()         { *m = InstanceShareableBertyID_Reply{} }
func (m *InstanceShareableBertyID_Reply) String() string { return proto.CompactTextString(m) }
func (*InstanceShareableBertyID_Reply) ProtoMessage()    {}
func (*InstanceShareableBertyID_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{0, 1}
}
func (m *InstanceShareableBertyID_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceShareableBertyID_Reply.Unmarshal(m, b)
}
func (m *InstanceShareableBertyID_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceShareableBertyID_Reply.Marshal(b, m, deterministic)
}
func (m *InstanceShareableBertyID_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceShareableBertyID_Reply.Merge(m, src)
}
func (m *InstanceShareableBertyID_Reply) XXX_Size() int {
	return xxx_messageInfo_InstanceShareableBertyID_Reply.Size(m)
}
func (m *InstanceShareableBertyID_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceShareableBertyID_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceShareableBertyID_Reply proto.InternalMessageInfo

func (m *InstanceShareableBertyID_Reply) GetBertyID() *BertyID {
	if m != nil {
		return m.BertyID
	}
	return nil
}

func (m *InstanceShareableBertyID_Reply) GetBertyIDPayload() string {
	if m != nil {
		return m.BertyIDPayload
	}
	return ""
}

func (m *InstanceShareableBertyID_Reply) GetDeepLink() string {
	if m != nil {
		return m.DeepLink
	}
	return ""
}

func (m *InstanceShareableBertyID_Reply) GetHTMLURL() string {
	if m != nil {
		return m.HTMLURL
	}
	return ""
}

type DevShareInstanceBertyID struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevShareInstanceBertyID) Reset()         { *m = DevShareInstanceBertyID{} }
func (m *DevShareInstanceBertyID) String() string { return proto.CompactTextString(m) }
func (*DevShareInstanceBertyID) ProtoMessage()    {}
func (*DevShareInstanceBertyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{1}
}
func (m *DevShareInstanceBertyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevShareInstanceBertyID.Unmarshal(m, b)
}
func (m *DevShareInstanceBertyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevShareInstanceBertyID.Marshal(b, m, deterministic)
}
func (m *DevShareInstanceBertyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevShareInstanceBertyID.Merge(m, src)
}
func (m *DevShareInstanceBertyID) XXX_Size() int {
	return xxx_messageInfo_DevShareInstanceBertyID.Size(m)
}
func (m *DevShareInstanceBertyID) XXX_DiscardUnknown() {
	xxx_messageInfo_DevShareInstanceBertyID.DiscardUnknown(m)
}

var xxx_messageInfo_DevShareInstanceBertyID proto.InternalMessageInfo

type DevShareInstanceBertyID_Request struct {
	// reset will regenerate a new link
	Reset_               bool     `protobuf:"varint,1,opt,name=reset,proto3" json:"reset,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevShareInstanceBertyID_Request) Reset()         { *m = DevShareInstanceBertyID_Request{} }
func (m *DevShareInstanceBertyID_Request) String() string { return proto.CompactTextString(m) }
func (*DevShareInstanceBertyID_Request) ProtoMessage()    {}
func (*DevShareInstanceBertyID_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{1, 0}
}
func (m *DevShareInstanceBertyID_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevShareInstanceBertyID_Request.Unmarshal(m, b)
}
func (m *DevShareInstanceBertyID_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevShareInstanceBertyID_Request.Marshal(b, m, deterministic)
}
func (m *DevShareInstanceBertyID_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevShareInstanceBertyID_Request.Merge(m, src)
}
func (m *DevShareInstanceBertyID_Request) XXX_Size() int {
	return xxx_messageInfo_DevShareInstanceBertyID_Request.Size(m)
}
func (m *DevShareInstanceBertyID_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_DevShareInstanceBertyID_Request.DiscardUnknown(m)
}

var xxx_messageInfo_DevShareInstanceBertyID_Request proto.InternalMessageInfo

func (m *DevShareInstanceBertyID_Request) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

func (m *DevShareInstanceBertyID_Request) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type DevShareInstanceBertyID_Reply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevShareInstanceBertyID_Reply) Reset()         { *m = DevShareInstanceBertyID_Reply{} }
func (m *DevShareInstanceBertyID_Reply) String() string { return proto.CompactTextString(m) }
func (*DevShareInstanceBertyID_Reply) ProtoMessage()    {}
func (*DevShareInstanceBertyID_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{1, 1}
}
func (m *DevShareInstanceBertyID_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevShareInstanceBertyID_Reply.Unmarshal(m, b)
}
func (m *DevShareInstanceBertyID_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevShareInstanceBertyID_Reply.Marshal(b, m, deterministic)
}
func (m *DevShareInstanceBertyID_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevShareInstanceBertyID_Reply.Merge(m, src)
}
func (m *DevShareInstanceBertyID_Reply) XXX_Size() int {
	return xxx_messageInfo_DevShareInstanceBertyID_Reply.Size(m)
}
func (m *DevShareInstanceBertyID_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_DevShareInstanceBertyID_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_DevShareInstanceBertyID_Reply proto.InternalMessageInfo

type ParseDeepLink struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseDeepLink) Reset()         { *m = ParseDeepLink{} }
func (m *ParseDeepLink) String() string { return proto.CompactTextString(m) }
func (*ParseDeepLink) ProtoMessage()    {}
func (*ParseDeepLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{2}
}
func (m *ParseDeepLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseDeepLink.Unmarshal(m, b)
}
func (m *ParseDeepLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseDeepLink.Marshal(b, m, deterministic)
}
func (m *ParseDeepLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseDeepLink.Merge(m, src)
}
func (m *ParseDeepLink) XXX_Size() int {
	return xxx_messageInfo_ParseDeepLink.Size(m)
}
func (m *ParseDeepLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseDeepLink.DiscardUnknown(m)
}

var xxx_messageInfo_ParseDeepLink proto.InternalMessageInfo

type ParseDeepLink_Request struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseDeepLink_Request) Reset()         { *m = ParseDeepLink_Request{} }
func (m *ParseDeepLink_Request) String() string { return proto.CompactTextString(m) }
func (*ParseDeepLink_Request) ProtoMessage()    {}
func (*ParseDeepLink_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{2, 0}
}
func (m *ParseDeepLink_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseDeepLink_Request.Unmarshal(m, b)
}
func (m *ParseDeepLink_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseDeepLink_Request.Marshal(b, m, deterministic)
}
func (m *ParseDeepLink_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseDeepLink_Request.Merge(m, src)
}
func (m *ParseDeepLink_Request) XXX_Size() int {
	return xxx_messageInfo_ParseDeepLink_Request.Size(m)
}
func (m *ParseDeepLink_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseDeepLink_Request.DiscardUnknown(m)
}

var xxx_messageInfo_ParseDeepLink_Request proto.InternalMessageInfo

func (m *ParseDeepLink_Request) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type ParseDeepLink_Reply struct {
	Kind                 ParseDeepLink_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=berty.messenger.ParseDeepLink_Kind" json:"kind,omitempty"`
	BertyID              *BertyID           `protobuf:"bytes,3,opt,name=berty_id,json=bertyId,proto3" json:"berty_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ParseDeepLink_Reply) Reset()         { *m = ParseDeepLink_Reply{} }
func (m *ParseDeepLink_Reply) String() string { return proto.CompactTextString(m) }
func (*ParseDeepLink_Reply) ProtoMessage()    {}
func (*ParseDeepLink_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{2, 1}
}
func (m *ParseDeepLink_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseDeepLink_Reply.Unmarshal(m, b)
}
func (m *ParseDeepLink_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseDeepLink_Reply.Marshal(b, m, deterministic)
}
func (m *ParseDeepLink_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseDeepLink_Reply.Merge(m, src)
}
func (m *ParseDeepLink_Reply) XXX_Size() int {
	return xxx_messageInfo_ParseDeepLink_Reply.Size(m)
}
func (m *ParseDeepLink_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseDeepLink_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_ParseDeepLink_Reply proto.InternalMessageInfo

func (m *ParseDeepLink_Reply) GetKind() ParseDeepLink_Kind {
	if m != nil {
		return m.Kind
	}
	return ParseDeepLink_UnknownKind
}

func (m *ParseDeepLink_Reply) GetBertyID() *BertyID {
	if m != nil {
		return m.BertyID
	}
	return nil
}

type SendContactRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendContactRequest) Reset()         { *m = SendContactRequest{} }
func (m *SendContactRequest) String() string { return proto.CompactTextString(m) }
func (*SendContactRequest) ProtoMessage()    {}
func (*SendContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{3}
}
func (m *SendContactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendContactRequest.Unmarshal(m, b)
}
func (m *SendContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendContactRequest.Marshal(b, m, deterministic)
}
func (m *SendContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendContactRequest.Merge(m, src)
}
func (m *SendContactRequest) XXX_Size() int {
	return xxx_messageInfo_SendContactRequest.Size(m)
}
func (m *SendContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendContactRequest proto.InternalMessageInfo

type SendContactRequest_Request struct {
	BertyID              *BertyID `protobuf:"bytes,1,opt,name=berty_id,json=bertyId,proto3" json:"berty_id,omitempty"`
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	OwnMetadata          []byte   `protobuf:"bytes,3,opt,name=own_metadata,json=ownMetadata,proto3" json:"own_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendContactRequest_Request) Reset()         { *m = SendContactRequest_Request{} }
func (m *SendContactRequest_Request) String() string { return proto.CompactTextString(m) }
func (*SendContactRequest_Request) ProtoMessage()    {}
func (*SendContactRequest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{3, 0}
}
func (m *SendContactRequest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendContactRequest_Request.Unmarshal(m, b)
}
func (m *SendContactRequest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendContactRequest_Request.Marshal(b, m, deterministic)
}
func (m *SendContactRequest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendContactRequest_Request.Merge(m, src)
}
func (m *SendContactRequest_Request) XXX_Size() int {
	return xxx_messageInfo_SendContactRequest_Request.Size(m)
}
func (m *SendContactRequest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_SendContactRequest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_SendContactRequest_Request proto.InternalMessageInfo

func (m *SendContactRequest_Request) GetBertyID() *BertyID {
	if m != nil {
		return m.BertyID
	}
	return nil
}

func (m *SendContactRequest_Request) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SendContactRequest_Request) GetOwnMetadata() []byte {
	if m != nil {
		return m.OwnMetadata
	}
	return nil
}

type SendContactRequest_Reply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendContactRequest_Reply) Reset()         { *m = SendContactRequest_Reply{} }
func (m *SendContactRequest_Reply) String() string { return proto.CompactTextString(m) }
func (*SendContactRequest_Reply) ProtoMessage()    {}
func (*SendContactRequest_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{3, 1}
}
func (m *SendContactRequest_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendContactRequest_Reply.Unmarshal(m, b)
}
func (m *SendContactRequest_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendContactRequest_Reply.Marshal(b, m, deterministic)
}
func (m *SendContactRequest_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendContactRequest_Reply.Merge(m, src)
}
func (m *SendContactRequest_Reply) XXX_Size() int {
	return xxx_messageInfo_SendContactRequest_Reply.Size(m)
}
func (m *SendContactRequest_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendContactRequest_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_SendContactRequest_Reply proto.InternalMessageInfo

type BertyID struct {
	PublicRendezvousSeed []byte   `protobuf:"bytes,1,opt,name=public_rendezvous_seed,json=publicRendezvousSeed,proto3" json:"public_rendezvous_seed,omitempty"`
	AccountPK            []byte   `protobuf:"bytes,2,opt,name=account_pk,json=accountPk,proto3" json:"account_pk,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BertyID) Reset()         { *m = BertyID{} }
func (m *BertyID) String() string { return proto.CompactTextString(m) }
func (*BertyID) ProtoMessage()    {}
func (*BertyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3bf21e238da6aa, []int{4}
}
func (m *BertyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BertyID.Unmarshal(m, b)
}
func (m *BertyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BertyID.Marshal(b, m, deterministic)
}
func (m *BertyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BertyID.Merge(m, src)
}
func (m *BertyID) XXX_Size() int {
	return xxx_messageInfo_BertyID.Size(m)
}
func (m *BertyID) XXX_DiscardUnknown() {
	xxx_messageInfo_BertyID.DiscardUnknown(m)
}

var xxx_messageInfo_BertyID proto.InternalMessageInfo

func (m *BertyID) GetPublicRendezvousSeed() []byte {
	if m != nil {
		return m.PublicRendezvousSeed
	}
	return nil
}

func (m *BertyID) GetAccountPK() []byte {
	if m != nil {
		return m.AccountPK
	}
	return nil
}

func (m *BertyID) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterEnum("berty.messenger.ParseDeepLink_Kind", ParseDeepLink_Kind_name, ParseDeepLink_Kind_value)
	proto.RegisterType((*InstanceShareableBertyID)(nil), "berty.messenger.InstanceShareableBertyID")
	proto.RegisterType((*InstanceShareableBertyID_Request)(nil), "berty.messenger.InstanceShareableBertyID.Request")
	proto.RegisterType((*InstanceShareableBertyID_Reply)(nil), "berty.messenger.InstanceShareableBertyID.Reply")
	proto.RegisterType((*DevShareInstanceBertyID)(nil), "berty.messenger.DevShareInstanceBertyID")
	proto.RegisterType((*DevShareInstanceBertyID_Request)(nil), "berty.messenger.DevShareInstanceBertyID.Request")
	proto.RegisterType((*DevShareInstanceBertyID_Reply)(nil), "berty.messenger.DevShareInstanceBertyID.Reply")
	proto.RegisterType((*ParseDeepLink)(nil), "berty.messenger.ParseDeepLink")
	proto.RegisterType((*ParseDeepLink_Request)(nil), "berty.messenger.ParseDeepLink.Request")
	proto.RegisterType((*ParseDeepLink_Reply)(nil), "berty.messenger.ParseDeepLink.Reply")
	proto.RegisterType((*SendContactRequest)(nil), "berty.messenger.SendContactRequest")
	proto.RegisterType((*SendContactRequest_Request)(nil), "berty.messenger.SendContactRequest.Request")
	proto.RegisterType((*SendContactRequest_Reply)(nil), "berty.messenger.SendContactRequest.Reply")
	proto.RegisterType((*BertyID)(nil), "berty.messenger.BertyID")
}

func init() { proto.RegisterFile("bertymessenger.proto", fileDescriptor_fd3bf21e238da6aa) }

var fileDescriptor_fd3bf21e238da6aa = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x4d, 0x4a, 0x92, 0x4d, 0xda, 0x46, 0xab, 0x0a, 0x22, 0x4b, 0x28, 0x25, 0x54, 0x55,
	0x2b, 0xc0, 0x81, 0x82, 0xc4, 0x85, 0x03, 0x98, 0x1e, 0xa8, 0xda, 0xa2, 0x68, 0x43, 0x2f, 0x48,
	0xc8, 0x5a, 0xdb, 0xd3, 0xc4, 0xb2, 0xbd, 0x6b, 0xec, 0x75, 0xab, 0x80, 0xc4, 0x81, 0x23, 0x27,
	0x9e, 0x83, 0x57, 0xe0, 0x49, 0x40, 0x28, 0x87, 0x3c, 0x09, 0xf2, 0xfa, 0xa7, 0x4a, 0x4d, 0x54,
	0x0a, 0xdc, 0x76, 0x66, 0xbe, 0xf9, 0x66, 0xf6, 0x9b, 0x59, 0x1b, 0xad, 0x9b, 0x10, 0x8a, 0x89,
	0x0f, 0x51, 0x04, 0x6c, 0x04, 0xa1, 0x16, 0x84, 0x5c, 0x70, 0xbc, 0x26, 0xbd, 0x5a, 0xe1, 0x56,
	0xef, 0x8f, 0x1c, 0x31, 0x8e, 0x4d, 0xcd, 0xe2, 0x7e, 0x7f, 0xc4, 0x47, 0xbc, 0x2f, 0x71, 0x66,
	0x7c, 0x22, 0x2d, 0x69, 0xc8, 0x53, 0x9a, 0xdf, 0xfb, 0xb6, 0x84, 0x3a, 0xfb, 0x2c, 0x12, 0x94,
	0x59, 0x30, 0x1c, 0xd3, 0x10, 0xa8, 0xe9, 0x81, 0x9e, 0x70, 0xee, 0xef, 0xa9, 0x3a, 0xaa, 0x11,
	0x78, 0x17, 0x43, 0x24, 0xf0, 0x3a, 0x5a, 0x0e, 0x21, 0x02, 0xd1, 0x51, 0x36, 0x94, 0xed, 0x3a,
	0x49, 0x0d, 0x7c, 0x1b, 0xb5, 0x6c, 0x27, 0x0a, 0x3c, 0x3a, 0x31, 0x18, 0xf5, 0xa1, 0xb3, 0xb4,
	0xa1, 0x6c, 0x37, 0x48, 0x33, 0xf3, 0xbd, 0xa2, 0x3e, 0xa8, 0x3f, 0x14, 0xb4, 0x4c, 0x20, 0xf0,
	0x26, 0xf8, 0x19, 0xaa, 0xcb, 0x66, 0x0d, 0xc7, 0x96, 0x2c, 0xcd, 0xdd, 0x8e, 0x76, 0xa1, 0x7b,
	0x2d, 0xab, 0xac, 0x37, 0x67, 0xd3, 0x6e, 0x2d, 0x33, 0x48, 0x4d, 0xa2, 0xf6, 0x6d, 0xfc, 0x14,
	0xb5, 0x73, 0x06, 0x23, 0xa0, 0x13, 0x8f, 0x53, 0x3b, 0x2d, 0xa9, 0xe3, 0xd9, 0xb4, 0xbb, 0x9a,
	0xe1, 0x07, 0x69, 0x84, 0xac, 0x66, 0x69, 0x99, 0x8d, 0x77, 0x50, 0xc3, 0x06, 0x08, 0x0c, 0xcf,
	0x61, 0x6e, 0xa7, 0x22, 0xd3, 0x5a, 0xb3, 0x69, 0xb7, 0xbe, 0x07, 0x10, 0x1c, 0x3a, 0xcc, 0x25,
	0x75, 0x3b, 0x3b, 0xe1, 0x2d, 0x54, 0x1f, 0x0b, 0xdf, 0x33, 0xe2, 0xd0, 0xeb, 0x54, 0x25, 0x52,
	0x36, 0xf4, 0xf2, 0xf5, 0xd1, 0xe1, 0x31, 0x39, 0x24, 0xb5, 0x24, 0x78, 0x1c, 0x7a, 0xbd, 0x13,
	0x74, 0x73, 0x0f, 0x4e, 0xa5, 0x6e, 0xb9, 0x88, 0xff, 0x53, 0xbb, 0x5a, 0x26, 0x5d, 0xef, 0xbb,
	0x82, 0x56, 0x06, 0x34, 0x8c, 0x20, 0xef, 0x55, 0xbd, 0x75, 0x4e, 0x8f, 0x51, 0x55, 0x5e, 0x49,
	0x91, 0x04, 0xf2, 0xac, 0x7e, 0x2a, 0x54, 0x7f, 0x82, 0xaa, 0xae, 0xc3, 0x52, 0xc5, 0x57, 0x77,
	0xef, 0x94, 0x14, 0x9f, 0xa3, 0xd5, 0x0e, 0x1c, 0x66, 0x13, 0x99, 0x30, 0x37, 0xae, 0xca, 0xdf,
	0x8c, 0xab, 0xb7, 0x89, 0xaa, 0x09, 0x1f, 0x5e, 0x43, 0xcd, 0x63, 0xe6, 0x32, 0x7e, 0xc6, 0x12,
	0xb3, 0x7d, 0x0d, 0x37, 0x51, 0x0e, 0x6e, 0x2b, 0xbd, 0xaf, 0x0a, 0xc2, 0x43, 0x60, 0xf6, 0x0b,
	0xce, 0x04, 0xb5, 0x44, 0x76, 0x2b, 0xf5, 0xb3, 0x72, 0x7e, 0xc3, 0x7f, 0xdf, 0x1c, 0x15, 0xd5,
	0x7d, 0x10, 0xd4, 0xa6, 0x82, 0x4a, 0xa1, 0x5b, 0xa4, 0xb0, 0x93, 0x41, 0xf0, 0x33, 0x66, 0x14,
	0xf1, 0x8a, 0x8c, 0x37, 0xf9, 0x19, 0x3b, 0xca, 0x5c, 0xe7, 0x83, 0xf8, 0xa2, 0x14, 0xad, 0xe3,
	0xc7, 0xe8, 0x46, 0x10, 0x9b, 0x9e, 0x63, 0x19, 0x21, 0x30, 0x1b, 0xde, 0x9f, 0xf2, 0x38, 0x32,
	0x22, 0x80, 0xb4, 0xc7, 0x16, 0x59, 0x4f, 0xa3, 0xa4, 0x08, 0x0e, 0x01, 0x6c, 0x7c, 0x0f, 0x21,
	0x6a, 0x59, 0x3c, 0x66, 0xc2, 0x08, 0xdc, 0xb4, 0x17, 0x7d, 0x65, 0x36, 0xed, 0x36, 0x9e, 0xa7,
	0xde, 0xc1, 0x01, 0x69, 0x64, 0x80, 0x81, 0x5b, 0x5a, 0x92, 0x4a, 0x69, 0x49, 0x76, 0x7f, 0x56,
	0x50, 0xfb, 0x28, 0x97, 0x61, 0x08, 0xe1, 0xa9, 0x63, 0x01, 0xfe, 0xb8, 0xf8, 0x55, 0xe3, 0x87,
	0x25, 0xed, 0x16, 0x41, 0xb5, 0x7c, 0x18, 0xfd, 0xab, 0xa4, 0x24, 0x5b, 0xf7, 0x61, 0xe1, 0xc3,
	0xc0, 0x0f, 0x4a, 0x5c, 0x0b, 0x90, 0x45, 0x75, 0xed, 0x0a, 0x19, 0x49, 0xf1, 0xb7, 0x17, 0x1e,
	0x0b, 0xde, 0xba, 0x64, 0xeb, 0xf3, 0x42, 0x9b, 0x97, 0xe2, 0x12, 0x7a, 0xef, 0x77, 0xfb, 0x8a,
	0xef, 0x96, 0x72, 0xcb, 0xa0, 0xa2, 0xd0, 0xce, 0x9f, 0x81, 0x03, 0x6f, 0xa2, 0x6f, 0xbf, 0xd9,
	0x4a, 0xb1, 0x02, 0xac, 0x71, 0x5f, 0x1e, 0xfb, 0xc9, 0x57, 0xdd, 0x1d, 0xf5, 0xe7, 0x7f, 0x08,
	0xe6, 0x75, 0xf9, 0x45, 0x7f, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x24, 0x87, 0x3e, 0x29,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessengerServiceClient is the client API for MessengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessengerServiceClient interface {
	// InstanceShareableBertyID returns a Berty ID that can be shared as a string, QR code or deep link.
	InstanceShareableBertyID(ctx context.Context, in *InstanceShareableBertyID_Request, opts ...grpc.CallOption) (*InstanceShareableBertyID_Reply, error)
	// DevShareInstanceBertyID shares your Berty ID on a dev channel.
	// TODO: remove for public.
	DevShareInstanceBertyID(ctx context.Context, in *DevShareInstanceBertyID_Request, opts ...grpc.CallOption) (*DevShareInstanceBertyID_Reply, error)
	// ParseDeepLink parses a link in the form of berty://xxx or https://berty.tech/id# and returns a structure
	// that can be used to display information.
	// This action is read-only.
	ParseDeepLink(ctx context.Context, in *ParseDeepLink_Request, opts ...grpc.CallOption) (*ParseDeepLink_Reply, error)
	// SendContactRequest takes the payload received from ParseDeepLink and send a contact request using the Berty Protocol.
	SendContactRequest(ctx context.Context, in *SendContactRequest_Request, opts ...grpc.CallOption) (*SendContactRequest_Reply, error)
}

type messengerServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessengerServiceClient(cc *grpc.ClientConn) MessengerServiceClient {
	return &messengerServiceClient{cc}
}

func (c *messengerServiceClient) InstanceShareableBertyID(ctx context.Context, in *InstanceShareableBertyID_Request, opts ...grpc.CallOption) (*InstanceShareableBertyID_Reply, error) {
	out := new(InstanceShareableBertyID_Reply)
	err := c.cc.Invoke(ctx, "/berty.messenger.MessengerService/InstanceShareableBertyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) DevShareInstanceBertyID(ctx context.Context, in *DevShareInstanceBertyID_Request, opts ...grpc.CallOption) (*DevShareInstanceBertyID_Reply, error) {
	out := new(DevShareInstanceBertyID_Reply)
	err := c.cc.Invoke(ctx, "/berty.messenger.MessengerService/DevShareInstanceBertyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) ParseDeepLink(ctx context.Context, in *ParseDeepLink_Request, opts ...grpc.CallOption) (*ParseDeepLink_Reply, error) {
	out := new(ParseDeepLink_Reply)
	err := c.cc.Invoke(ctx, "/berty.messenger.MessengerService/ParseDeepLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) SendContactRequest(ctx context.Context, in *SendContactRequest_Request, opts ...grpc.CallOption) (*SendContactRequest_Reply, error) {
	out := new(SendContactRequest_Reply)
	err := c.cc.Invoke(ctx, "/berty.messenger.MessengerService/SendContactRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServiceServer is the server API for MessengerService service.
type MessengerServiceServer interface {
	// InstanceShareableBertyID returns a Berty ID that can be shared as a string, QR code or deep link.
	InstanceShareableBertyID(context.Context, *InstanceShareableBertyID_Request) (*InstanceShareableBertyID_Reply, error)
	// DevShareInstanceBertyID shares your Berty ID on a dev channel.
	// TODO: remove for public.
	DevShareInstanceBertyID(context.Context, *DevShareInstanceBertyID_Request) (*DevShareInstanceBertyID_Reply, error)
	// ParseDeepLink parses a link in the form of berty://xxx or https://berty.tech/id# and returns a structure
	// that can be used to display information.
	// This action is read-only.
	ParseDeepLink(context.Context, *ParseDeepLink_Request) (*ParseDeepLink_Reply, error)
	// SendContactRequest takes the payload received from ParseDeepLink and send a contact request using the Berty Protocol.
	SendContactRequest(context.Context, *SendContactRequest_Request) (*SendContactRequest_Reply, error)
}

// UnimplementedMessengerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessengerServiceServer struct {
}

func (*UnimplementedMessengerServiceServer) InstanceShareableBertyID(ctx context.Context, req *InstanceShareableBertyID_Request) (*InstanceShareableBertyID_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstanceShareableBertyID not implemented")
}
func (*UnimplementedMessengerServiceServer) DevShareInstanceBertyID(ctx context.Context, req *DevShareInstanceBertyID_Request) (*DevShareInstanceBertyID_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevShareInstanceBertyID not implemented")
}
func (*UnimplementedMessengerServiceServer) ParseDeepLink(ctx context.Context, req *ParseDeepLink_Request) (*ParseDeepLink_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseDeepLink not implemented")
}
func (*UnimplementedMessengerServiceServer) SendContactRequest(ctx context.Context, req *SendContactRequest_Request) (*SendContactRequest_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendContactRequest not implemented")
}

func RegisterMessengerServiceServer(s *grpc.Server, srv MessengerServiceServer) {
	s.RegisterService(&_MessengerService_serviceDesc, srv)
}

func _MessengerService_InstanceShareableBertyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceShareableBertyID_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).InstanceShareableBertyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.messenger.MessengerService/InstanceShareableBertyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).InstanceShareableBertyID(ctx, req.(*InstanceShareableBertyID_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_DevShareInstanceBertyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevShareInstanceBertyID_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).DevShareInstanceBertyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.messenger.MessengerService/DevShareInstanceBertyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).DevShareInstanceBertyID(ctx, req.(*DevShareInstanceBertyID_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_ParseDeepLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseDeepLink_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).ParseDeepLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.messenger.MessengerService/ParseDeepLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).ParseDeepLink(ctx, req.(*ParseDeepLink_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_SendContactRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendContactRequest_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).SendContactRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.messenger.MessengerService/SendContactRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).SendContactRequest(ctx, req.(*SendContactRequest_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessengerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "berty.messenger.MessengerService",
	HandlerType: (*MessengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstanceShareableBertyID",
			Handler:    _MessengerService_InstanceShareableBertyID_Handler,
		},
		{
			MethodName: "DevShareInstanceBertyID",
			Handler:    _MessengerService_DevShareInstanceBertyID_Handler,
		},
		{
			MethodName: "ParseDeepLink",
			Handler:    _MessengerService_ParseDeepLink_Handler,
		},
		{
			MethodName: "SendContactRequest",
			Handler:    _MessengerService_SendContactRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bertymessenger.proto",
}
