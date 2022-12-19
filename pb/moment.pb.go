// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: moment.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Moment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt    int64    `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	CatId       string   `protobuf:"bytes,3,opt,name=catId,proto3" json:"catId,omitempty"`
	Photos      []string `protobuf:"bytes,4,rep,name=photos,proto3" json:"photos,omitempty"` // 图片url
	Title       string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Text        string   `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	CommunityId string   `protobuf:"bytes,7,opt,name=communityId,proto3" json:"communityId,omitempty"`
	UserId      string   `protobuf:"bytes,8,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Moment) Reset() {
	*x = Moment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moment) ProtoMessage() {}

func (x *Moment) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moment.ProtoReflect.Descriptor instead.
func (*Moment) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{0}
}

func (x *Moment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Moment) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Moment) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *Moment) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Moment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Moment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Moment) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *Moment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SearchMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Skip        int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	CommunityId string `protobuf:"bytes,3,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Keyword     string `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *SearchMomentReq) Reset() {
	*x = SearchMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMomentReq) ProtoMessage() {}

func (x *SearchMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMomentReq.ProtoReflect.Descriptor instead.
func (*SearchMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMomentReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchMomentReq) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *SearchMomentReq) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *SearchMomentReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type SearchMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moments []*Moment `protobuf:"bytes,1,rep,name=moments,proto3" json:"moments,omitempty"`
}

func (x *SearchMomentResp) Reset() {
	*x = SearchMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMomentResp) ProtoMessage() {}

func (x *SearchMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMomentResp.ProtoReflect.Descriptor instead.
func (*SearchMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{2}
}

func (x *SearchMomentResp) GetMoments() []*Moment {
	if x != nil {
		return x.Moments
	}
	return nil
}

type ListMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 获取的数量
	Skip        int64  `protobuf:"varint,3,opt,name=skip,proto3" json:"skip,omitempty"`
}

func (x *ListMomentReq) Reset() {
	*x = ListMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMomentReq) ProtoMessage() {}

func (x *ListMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMomentReq.ProtoReflect.Descriptor instead.
func (*ListMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{3}
}

func (x *ListMomentReq) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *ListMomentReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListMomentReq) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

type ListMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moments []*Moment `protobuf:"bytes,1,rep,name=moments,proto3" json:"moments,omitempty"`
}

func (x *ListMomentResp) Reset() {
	*x = ListMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMomentResp) ProtoMessage() {}

func (x *ListMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMomentResp.ProtoReflect.Descriptor instead.
func (*ListMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{4}
}

func (x *ListMomentResp) GetMoments() []*Moment {
	if x != nil {
		return x.Moments
	}
	return nil
}

type RetrieveMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId,omitempty"`
}

func (x *RetrieveMomentReq) Reset() {
	*x = RetrieveMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveMomentReq) ProtoMessage() {}

func (x *RetrieveMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveMomentReq.ProtoReflect.Descriptor instead.
func (*RetrieveMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{5}
}

func (x *RetrieveMomentReq) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

type RetrieveMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moment *Moment `protobuf:"bytes,1,opt,name=moment,proto3" json:"moment,omitempty"`
}

func (x *RetrieveMomentResp) Reset() {
	*x = RetrieveMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveMomentResp) ProtoMessage() {}

func (x *RetrieveMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveMomentResp.ProtoReflect.Descriptor instead.
func (*RetrieveMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{6}
}

func (x *RetrieveMomentResp) GetMoment() *Moment {
	if x != nil {
		return x.Moment
	}
	return nil
}

type CreateMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moment *Moment `protobuf:"bytes,1,opt,name=moment,proto3" json:"moment,omitempty"`
}

func (x *CreateMomentReq) Reset() {
	*x = CreateMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMomentReq) ProtoMessage() {}

func (x *CreateMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMomentReq.ProtoReflect.Descriptor instead.
func (*CreateMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{7}
}

func (x *CreateMomentReq) GetMoment() *Moment {
	if x != nil {
		return x.Moment
	}
	return nil
}

type CreateMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId,omitempty"`
}

func (x *CreateMomentResp) Reset() {
	*x = CreateMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMomentResp) ProtoMessage() {}

func (x *CreateMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMomentResp.ProtoReflect.Descriptor instead.
func (*CreateMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{8}
}

func (x *CreateMomentResp) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

type UpdateMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moment *Moment `protobuf:"bytes,1,opt,name=moment,proto3" json:"moment,omitempty"`
}

func (x *UpdateMomentReq) Reset() {
	*x = UpdateMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMomentReq) ProtoMessage() {}

func (x *UpdateMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMomentReq.ProtoReflect.Descriptor instead.
func (*UpdateMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateMomentReq) GetMoment() *Moment {
	if x != nil {
		return x.Moment
	}
	return nil
}

type UpdateMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMomentResp) Reset() {
	*x = UpdateMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMomentResp) ProtoMessage() {}

func (x *UpdateMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMomentResp.ProtoReflect.Descriptor instead.
func (*UpdateMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{10}
}

type DeleteMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId,omitempty"`
}

func (x *DeleteMomentReq) Reset() {
	*x = DeleteMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentReq) ProtoMessage() {}

func (x *DeleteMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentReq.ProtoReflect.Descriptor instead.
func (*DeleteMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteMomentReq) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

type DeleteMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMomentResp) Reset() {
	*x = DeleteMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentResp) ProtoMessage() {}

func (x *DeleteMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentResp.ProtoReflect.Descriptor instead.
func (*DeleteMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{12}
}

var File_moment_proto protoreflect.FileDescriptor

var file_moment_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x77, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x07,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5b, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x6b, 0x69, 0x70, 0x22, 0x3a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x2f, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x39,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x06,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0x9e, 0x03, 0x0a, 0x0a,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moment_proto_rawDescOnce sync.Once
	file_moment_proto_rawDescData = file_moment_proto_rawDesc
)

func file_moment_proto_rawDescGZIP() []byte {
	file_moment_proto_rawDescOnce.Do(func() {
		file_moment_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_proto_rawDescData)
	})
	return file_moment_proto_rawDescData
}

var file_moment_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_moment_proto_goTypes = []interface{}{
	(*Moment)(nil),             // 0: moment.Moment
	(*SearchMomentReq)(nil),    // 1: moment.SearchMomentReq
	(*SearchMomentResp)(nil),   // 2: moment.SearchMomentResp
	(*ListMomentReq)(nil),      // 3: moment.ListMomentReq
	(*ListMomentResp)(nil),     // 4: moment.ListMomentResp
	(*RetrieveMomentReq)(nil),  // 5: moment.RetrieveMomentReq
	(*RetrieveMomentResp)(nil), // 6: moment.RetrieveMomentResp
	(*CreateMomentReq)(nil),    // 7: moment.CreateMomentReq
	(*CreateMomentResp)(nil),   // 8: moment.CreateMomentResp
	(*UpdateMomentReq)(nil),    // 9: moment.UpdateMomentReq
	(*UpdateMomentResp)(nil),   // 10: moment.UpdateMomentResp
	(*DeleteMomentReq)(nil),    // 11: moment.DeleteMomentReq
	(*DeleteMomentResp)(nil),   // 12: moment.DeleteMomentResp
}
var file_moment_proto_depIdxs = []int32{
	0,  // 0: moment.SearchMomentResp.moments:type_name -> moment.Moment
	0,  // 1: moment.ListMomentResp.moments:type_name -> moment.Moment
	0,  // 2: moment.RetrieveMomentResp.moment:type_name -> moment.Moment
	0,  // 3: moment.CreateMomentReq.moment:type_name -> moment.Moment
	0,  // 4: moment.UpdateMomentReq.moment:type_name -> moment.Moment
	1,  // 5: moment.moment_rpc.SearchMoment:input_type -> moment.SearchMomentReq
	3,  // 6: moment.moment_rpc.ListMoment:input_type -> moment.ListMomentReq
	5,  // 7: moment.moment_rpc.RetrieveMoment:input_type -> moment.RetrieveMomentReq
	7,  // 8: moment.moment_rpc.CreateMoment:input_type -> moment.CreateMomentReq
	9,  // 9: moment.moment_rpc.UpdateMoment:input_type -> moment.UpdateMomentReq
	11, // 10: moment.moment_rpc.DeleteMoment:input_type -> moment.DeleteMomentReq
	2,  // 11: moment.moment_rpc.SearchMoment:output_type -> moment.SearchMomentResp
	4,  // 12: moment.moment_rpc.ListMoment:output_type -> moment.ListMomentResp
	6,  // 13: moment.moment_rpc.RetrieveMoment:output_type -> moment.RetrieveMomentResp
	8,  // 14: moment.moment_rpc.CreateMoment:output_type -> moment.CreateMomentResp
	10, // 15: moment.moment_rpc.UpdateMoment:output_type -> moment.UpdateMomentResp
	12, // 16: moment.moment_rpc.DeleteMoment:output_type -> moment.DeleteMomentResp
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_moment_proto_init() }
func file_moment_proto_init() {
	if File_moment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_moment_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moment_proto_goTypes,
		DependencyIndexes: file_moment_proto_depIdxs,
		MessageInfos:      file_moment_proto_msgTypes,
	}.Build()
	File_moment_proto = out.File
	file_moment_proto_rawDesc = nil
	file_moment_proto_goTypes = nil
	file_moment_proto_depIdxs = nil
}
