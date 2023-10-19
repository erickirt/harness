// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: merge.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MergeRequest_MergeMethod int32

const (
	MergeRequest_merge  MergeRequest_MergeMethod = 0
	MergeRequest_squash MergeRequest_MergeMethod = 1
	MergeRequest_rebase MergeRequest_MergeMethod = 2
)

// Enum value maps for MergeRequest_MergeMethod.
var (
	MergeRequest_MergeMethod_name = map[int32]string{
		0: "merge",
		1: "squash",
		2: "rebase",
	}
	MergeRequest_MergeMethod_value = map[string]int32{
		"merge":  0,
		"squash": 1,
		"rebase": 2,
	}
)

func (x MergeRequest_MergeMethod) Enum() *MergeRequest_MergeMethod {
	p := new(MergeRequest_MergeMethod)
	*p = x
	return p
}

func (x MergeRequest_MergeMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MergeRequest_MergeMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_merge_proto_enumTypes[0].Descriptor()
}

func (MergeRequest_MergeMethod) Type() protoreflect.EnumType {
	return &file_merge_proto_enumTypes[0]
}

func (x MergeRequest_MergeMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MergeRequest_MergeMethod.Descriptor instead.
func (MergeRequest_MergeMethod) EnumDescriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{0, 0}
}

type MergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *WriteRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// head_branch is the source branch we want to merge
	HeadBranch string `protobuf:"bytes,2,opt,name=head_branch,json=headBranch,proto3" json:"head_branch,omitempty"`
	// base_branch is the branch into which the given commit shall be merged and whose
	// reference is going to be updated.
	BaseBranch string `protobuf:"bytes,3,opt,name=base_branch,json=baseBranch,proto3" json:"base_branch,omitempty"`
	// title is the title to use for the merge commit.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// message is the message to use for the merge commit.
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	// author is the person who originally wrote the code
	Author *Identity `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	// authorDate is the date when the code was written
	AuthorDate int64 `protobuf:"varint,7,opt,name=authorDate,proto3" json:"authorDate,omitempty"`
	// committer is the person who last applied the patch
	Committer *Identity `protobuf:"bytes,8,opt,name=committer,proto3" json:"committer,omitempty"`
	// committer is the date when the code was applied
	CommitterDate int64 `protobuf:"varint,9,opt,name=committerDate,proto3" json:"committerDate,omitempty"`
	// ref_type is an otional value and is used to generate the full
	// reference in which the merge result is stored.
	RefType RefType `protobuf:"varint,10,opt,name=ref_type,json=refType,proto3,enum=rpc.RefType" json:"ref_type,omitempty"`
	// ref_name is an otional value and is used to generate the full
	// reference in which the merge result is stored.
	RefName string `protobuf:"bytes,11,opt,name=ref_name,json=refName,proto3" json:"ref_name,omitempty"`
	// head_expected_sha is commit sha on the head branch, if head_expected_sha is older
	// than the head_branch latest sha then merge will fail.
	HeadExpectedSha string `protobuf:"bytes,12,opt,name=head_expected_sha,json=headExpectedSha,proto3" json:"head_expected_sha,omitempty"`
	// force merge
	Force bool `protobuf:"varint,13,opt,name=force,proto3" json:"force,omitempty"`
	// delete branch after merge
	DeleteHeadBranch bool `protobuf:"varint,14,opt,name=delete_head_branch,json=deleteHeadBranch,proto3" json:"delete_head_branch,omitempty"`
	// merging method
	Method MergeRequest_MergeMethod `protobuf:"varint,15,opt,name=method,proto3,enum=rpc.MergeRequest_MergeMethod" json:"method,omitempty"`
}

func (x *MergeRequest) Reset() {
	*x = MergeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeRequest) ProtoMessage() {}

func (x *MergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeRequest.ProtoReflect.Descriptor instead.
func (*MergeRequest) Descriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{0}
}

func (x *MergeRequest) GetBase() *WriteRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *MergeRequest) GetHeadBranch() string {
	if x != nil {
		return x.HeadBranch
	}
	return ""
}

func (x *MergeRequest) GetBaseBranch() string {
	if x != nil {
		return x.BaseBranch
	}
	return ""
}

func (x *MergeRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MergeRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MergeRequest) GetAuthor() *Identity {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *MergeRequest) GetAuthorDate() int64 {
	if x != nil {
		return x.AuthorDate
	}
	return 0
}

func (x *MergeRequest) GetCommitter() *Identity {
	if x != nil {
		return x.Committer
	}
	return nil
}

func (x *MergeRequest) GetCommitterDate() int64 {
	if x != nil {
		return x.CommitterDate
	}
	return 0
}

func (x *MergeRequest) GetRefType() RefType {
	if x != nil {
		return x.RefType
	}
	return RefType_Undefined
}

func (x *MergeRequest) GetRefName() string {
	if x != nil {
		return x.RefName
	}
	return ""
}

func (x *MergeRequest) GetHeadExpectedSha() string {
	if x != nil {
		return x.HeadExpectedSha
	}
	return ""
}

func (x *MergeRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *MergeRequest) GetDeleteHeadBranch() bool {
	if x != nil {
		return x.DeleteHeadBranch
	}
	return false
}

func (x *MergeRequest) GetMethod() MergeRequest_MergeMethod {
	if x != nil {
		return x.Method
	}
	return MergeRequest_merge
}

type MergeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base_sha is the sha of the latest commit on the base branch that was used for merging.
	BaseSha string `protobuf:"bytes,1,opt,name=base_sha,json=baseSha,proto3" json:"base_sha,omitempty"`
	// head_sha is the sha of the latest commit on the head branch that was used for merging.
	HeadSha string `protobuf:"bytes,2,opt,name=head_sha,json=headSha,proto3" json:"head_sha,omitempty"`
	// merge_base_sha is the sha of the merge base of the head_sha and base_sha
	MergeBaseSha string `protobuf:"bytes,3,opt,name=merge_base_sha,json=mergeBaseSha,proto3" json:"merge_base_sha,omitempty"`
	// merge_sha is the sha of the commit after merging head_sha with base_sha.
	MergeSha string `protobuf:"bytes,4,opt,name=merge_sha,json=mergeSha,proto3" json:"merge_sha,omitempty"`
}

func (x *MergeResponse) Reset() {
	*x = MergeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeResponse) ProtoMessage() {}

func (x *MergeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeResponse.ProtoReflect.Descriptor instead.
func (*MergeResponse) Descriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{1}
}

func (x *MergeResponse) GetBaseSha() string {
	if x != nil {
		return x.BaseSha
	}
	return ""
}

func (x *MergeResponse) GetHeadSha() string {
	if x != nil {
		return x.HeadSha
	}
	return ""
}

func (x *MergeResponse) GetMergeBaseSha() string {
	if x != nil {
		return x.MergeBaseSha
	}
	return ""
}

func (x *MergeResponse) GetMergeSha() string {
	if x != nil {
		return x.MergeSha
	}
	return ""
}

// MergeConflictError is an error returned in the case when merging two commits
// fails due to a merge conflict.
type MergeConflictError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConflictingFiles is the set of files which have been conflicting.
	ConflictingFiles []string `protobuf:"bytes,1,rep,name=conflicting_files,json=conflictingFiles,proto3" json:"conflicting_files,omitempty"`
}

func (x *MergeConflictError) Reset() {
	*x = MergeConflictError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeConflictError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeConflictError) ProtoMessage() {}

func (x *MergeConflictError) ProtoReflect() protoreflect.Message {
	mi := &file_merge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeConflictError.ProtoReflect.Descriptor instead.
func (*MergeConflictError) Descriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{2}
}

func (x *MergeConflictError) GetConflictingFiles() []string {
	if x != nil {
		return x.ConflictingFiles
	}
	return nil
}

var File_merge_proto protoreflect.FileDescriptor

var file_merge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xde, 0x04, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64,
	0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68,
	0x65, 0x61, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x61, 0x73, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x72, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x65, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x65, 0x61, 0x64,
	0x5f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x53, 0x68, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x30, 0x0a, 0x0b, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x09,
	0x0a, 0x05, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x71, 0x75,
	0x61, 0x73, 0x68, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x10,
	0x02, 0x22, 0x88, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x53, 0x68, 0x61, 0x12, 0x19,
	0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x53, 0x68, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x53, 0x68, 0x61, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x53, 0x68, 0x61, 0x22, 0x41, 0x0a, 0x12,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x32,
	0x40, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x67, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_merge_proto_rawDescOnce sync.Once
	file_merge_proto_rawDescData = file_merge_proto_rawDesc
)

func file_merge_proto_rawDescGZIP() []byte {
	file_merge_proto_rawDescOnce.Do(func() {
		file_merge_proto_rawDescData = protoimpl.X.CompressGZIP(file_merge_proto_rawDescData)
	})
	return file_merge_proto_rawDescData
}

var file_merge_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_merge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_merge_proto_goTypes = []interface{}{
	(MergeRequest_MergeMethod)(0), // 0: rpc.MergeRequest.MergeMethod
	(*MergeRequest)(nil),          // 1: rpc.MergeRequest
	(*MergeResponse)(nil),         // 2: rpc.MergeResponse
	(*MergeConflictError)(nil),    // 3: rpc.MergeConflictError
	(*WriteRequest)(nil),          // 4: rpc.WriteRequest
	(*Identity)(nil),              // 5: rpc.Identity
	(RefType)(0),                  // 6: rpc.RefType
}
var file_merge_proto_depIdxs = []int32{
	4, // 0: rpc.MergeRequest.base:type_name -> rpc.WriteRequest
	5, // 1: rpc.MergeRequest.author:type_name -> rpc.Identity
	5, // 2: rpc.MergeRequest.committer:type_name -> rpc.Identity
	6, // 3: rpc.MergeRequest.ref_type:type_name -> rpc.RefType
	0, // 4: rpc.MergeRequest.method:type_name -> rpc.MergeRequest.MergeMethod
	1, // 5: rpc.MergeService.Merge:input_type -> rpc.MergeRequest
	2, // 6: rpc.MergeService.Merge:output_type -> rpc.MergeResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_merge_proto_init() }
func file_merge_proto_init() {
	if File_merge_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_merge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeRequest); i {
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
		file_merge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeResponse); i {
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
		file_merge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeConflictError); i {
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
			RawDescriptor: file_merge_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merge_proto_goTypes,
		DependencyIndexes: file_merge_proto_depIdxs,
		EnumInfos:         file_merge_proto_enumTypes,
		MessageInfos:      file_merge_proto_msgTypes,
	}.Build()
	File_merge_proto = out.File
	file_merge_proto_rawDesc = nil
	file_merge_proto_goTypes = nil
	file_merge_proto_depIdxs = nil
}
