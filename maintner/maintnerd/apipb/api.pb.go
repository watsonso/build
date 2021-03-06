// Code generated by protoc-gen-go4grpc; DO NOT EDIT
// source: api.proto

/*
Package apipb is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	HasAncestorRequest
	HasAncestorResponse
	GetRefRequest
	GetRefResponse
	GoFindTryWorkRequest
	GoFindTryWorkResponse
	GerritTryWorkItem
	ListGoReleasesRequest
	ListGoReleasesResponse
	GoRelease
*/
package apipb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "grpc.go4.org"
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

type HasAncestorRequest struct {
	Commit   string `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Ancestor string `protobuf:"bytes,2,opt,name=ancestor" json:"ancestor,omitempty"`
}

func (m *HasAncestorRequest) Reset()                    { *m = HasAncestorRequest{} }
func (m *HasAncestorRequest) String() string            { return proto.CompactTextString(m) }
func (*HasAncestorRequest) ProtoMessage()               {}
func (*HasAncestorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HasAncestorRequest) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *HasAncestorRequest) GetAncestor() string {
	if m != nil {
		return m.Ancestor
	}
	return ""
}

type HasAncestorResponse struct {
	// has_ancestor is whether ancestor appears in commit's history.
	HasAncestor bool `protobuf:"varint,1,opt,name=has_ancestor,json=hasAncestor" json:"has_ancestor,omitempty"`
	// unknown_commit is true if the provided commit was unknown.
	UnknownCommit bool `protobuf:"varint,2,opt,name=unknown_commit,json=unknownCommit" json:"unknown_commit,omitempty"`
}

func (m *HasAncestorResponse) Reset()                    { *m = HasAncestorResponse{} }
func (m *HasAncestorResponse) String() string            { return proto.CompactTextString(m) }
func (*HasAncestorResponse) ProtoMessage()               {}
func (*HasAncestorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HasAncestorResponse) GetHasAncestor() bool {
	if m != nil {
		return m.HasAncestor
	}
	return false
}

func (m *HasAncestorResponse) GetUnknownCommit() bool {
	if m != nil {
		return m.UnknownCommit
	}
	return false
}

type GetRefRequest struct {
	Ref string `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	// Either gerrit_server & gerrit_project must be specified, or
	// github. Currently only Gerrit is supported.
	GerritServer  string `protobuf:"bytes,2,opt,name=gerrit_server,json=gerritServer" json:"gerrit_server,omitempty"`
	GerritProject string `protobuf:"bytes,3,opt,name=gerrit_project,json=gerritProject" json:"gerrit_project,omitempty"`
}

func (m *GetRefRequest) Reset()                    { *m = GetRefRequest{} }
func (m *GetRefRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRefRequest) ProtoMessage()               {}
func (*GetRefRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRefRequest) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GetRefRequest) GetGerritServer() string {
	if m != nil {
		return m.GerritServer
	}
	return ""
}

func (m *GetRefRequest) GetGerritProject() string {
	if m != nil {
		return m.GerritProject
	}
	return ""
}

type GetRefResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetRefResponse) Reset()                    { *m = GetRefResponse{} }
func (m *GetRefResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRefResponse) ProtoMessage()               {}
func (*GetRefResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRefResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GoFindTryWorkRequest struct {
	// for_staging says whether this is a trybot request for the staging
	// cluster. When using staging, the comment "Run-StagingTryBot"
	// is used instead of label:Run-TryBot=1.
	ForStaging bool `protobuf:"varint,1,opt,name=for_staging,json=forStaging" json:"for_staging,omitempty"`
}

func (m *GoFindTryWorkRequest) Reset()                    { *m = GoFindTryWorkRequest{} }
func (m *GoFindTryWorkRequest) String() string            { return proto.CompactTextString(m) }
func (*GoFindTryWorkRequest) ProtoMessage()               {}
func (*GoFindTryWorkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GoFindTryWorkRequest) GetForStaging() bool {
	if m != nil {
		return m.ForStaging
	}
	return false
}

type GoFindTryWorkResponse struct {
	// waiting are the Gerrit CLs wanting a trybot run and not yet with results.
	// These might already be running.
	Waiting []*GerritTryWorkItem `protobuf:"bytes,1,rep,name=waiting" json:"waiting,omitempty"`
}

func (m *GoFindTryWorkResponse) Reset()                    { *m = GoFindTryWorkResponse{} }
func (m *GoFindTryWorkResponse) String() string            { return proto.CompactTextString(m) }
func (*GoFindTryWorkResponse) ProtoMessage()               {}
func (*GoFindTryWorkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GoFindTryWorkResponse) GetWaiting() []*GerritTryWorkItem {
	if m != nil {
		return m.Waiting
	}
	return nil
}

type GerritTryWorkItem struct {
	Project  string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Branch   string `protobuf:"bytes,2,opt,name=branch" json:"branch,omitempty"`
	ChangeId string `protobuf:"bytes,3,opt,name=change_id,json=changeId" json:"change_id,omitempty"`
	Commit   string `protobuf:"bytes,4,opt,name=commit" json:"commit,omitempty"`
	// go_commit is set for subrepos and is the Go commit(s) to test against.
	// go_branch is a branch name of go_commit, for showing to users when
	// a try set fails.
	GoCommit []string `protobuf:"bytes,5,rep,name=go_commit,json=goCommit" json:"go_commit,omitempty"`
	GoBranch []string `protobuf:"bytes,6,rep,name=go_branch,json=goBranch" json:"go_branch,omitempty"`
}

func (m *GerritTryWorkItem) Reset()                    { *m = GerritTryWorkItem{} }
func (m *GerritTryWorkItem) String() string            { return proto.CompactTextString(m) }
func (*GerritTryWorkItem) ProtoMessage()               {}
func (*GerritTryWorkItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GerritTryWorkItem) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritTryWorkItem) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *GerritTryWorkItem) GetChangeId() string {
	if m != nil {
		return m.ChangeId
	}
	return ""
}

func (m *GerritTryWorkItem) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *GerritTryWorkItem) GetGoCommit() []string {
	if m != nil {
		return m.GoCommit
	}
	return nil
}

func (m *GerritTryWorkItem) GetGoBranch() []string {
	if m != nil {
		return m.GoBranch
	}
	return nil
}

// By default, ListGoReleases returns only the latest patches
// of releases that are considered supported per policy.
type ListGoReleasesRequest struct {
}

func (m *ListGoReleasesRequest) Reset()                    { *m = ListGoReleasesRequest{} }
func (m *ListGoReleasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGoReleasesRequest) ProtoMessage()               {}
func (*ListGoReleasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListGoReleasesResponse struct {
	// Releases are Go releases, sorted with latest release first.
	Releases []*GoRelease `protobuf:"bytes,1,rep,name=releases" json:"releases,omitempty"`
}

func (m *ListGoReleasesResponse) Reset()                    { *m = ListGoReleasesResponse{} }
func (m *ListGoReleasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGoReleasesResponse) ProtoMessage()               {}
func (*ListGoReleasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListGoReleasesResponse) GetReleases() []*GoRelease {
	if m != nil {
		return m.Releases
	}
	return nil
}

type GoRelease struct {
	Major     int32  `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor     int32  `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch     int32  `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	TagName   string `protobuf:"bytes,4,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	TagCommit string `protobuf:"bytes,5,opt,name=tag_commit,json=tagCommit" json:"tag_commit,omitempty"`
	// Release branch information for this major-minor version pair.
	BranchName   string `protobuf:"bytes,6,opt,name=branch_name,json=branchName" json:"branch_name,omitempty"`
	BranchCommit string `protobuf:"bytes,7,opt,name=branch_commit,json=branchCommit" json:"branch_commit,omitempty"`
}

func (m *GoRelease) Reset()                    { *m = GoRelease{} }
func (m *GoRelease) String() string            { return proto.CompactTextString(m) }
func (*GoRelease) ProtoMessage()               {}
func (*GoRelease) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GoRelease) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *GoRelease) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *GoRelease) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *GoRelease) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func (m *GoRelease) GetTagCommit() string {
	if m != nil {
		return m.TagCommit
	}
	return ""
}

func (m *GoRelease) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *GoRelease) GetBranchCommit() string {
	if m != nil {
		return m.BranchCommit
	}
	return ""
}

func init() {
	proto.RegisterType((*HasAncestorRequest)(nil), "apipb.HasAncestorRequest")
	proto.RegisterType((*HasAncestorResponse)(nil), "apipb.HasAncestorResponse")
	proto.RegisterType((*GetRefRequest)(nil), "apipb.GetRefRequest")
	proto.RegisterType((*GetRefResponse)(nil), "apipb.GetRefResponse")
	proto.RegisterType((*GoFindTryWorkRequest)(nil), "apipb.GoFindTryWorkRequest")
	proto.RegisterType((*GoFindTryWorkResponse)(nil), "apipb.GoFindTryWorkResponse")
	proto.RegisterType((*GerritTryWorkItem)(nil), "apipb.GerritTryWorkItem")
	proto.RegisterType((*ListGoReleasesRequest)(nil), "apipb.ListGoReleasesRequest")
	proto.RegisterType((*ListGoReleasesResponse)(nil), "apipb.ListGoReleasesResponse")
	proto.RegisterType((*GoRelease)(nil), "apipb.GoRelease")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MaintnerService service

type MaintnerServiceClient interface {
	// HasAncestor reports whether one commit contains another commit
	// in its git history.
	HasAncestor(ctx context.Context, in *HasAncestorRequest, opts ...grpc.CallOption) (*HasAncestorResponse, error)
	// GetRef returns information about a git ref.
	GetRef(ctx context.Context, in *GetRefRequest, opts ...grpc.CallOption) (*GetRefResponse, error)
	// GoFindTryWork finds trybot work for the coordinator to build & test.
	GoFindTryWork(ctx context.Context, in *GoFindTryWorkRequest, opts ...grpc.CallOption) (*GoFindTryWorkResponse, error)
	// ListGoReleases lists Go releases. A release is considered to exist for
	// each git tag named "goX", "goX.Y", or "goX.Y.Z", as long as it has a
	// corresponding "release-branch.goX" or "release-branch.goX.Y" release branch.
	ListGoReleases(ctx context.Context, in *ListGoReleasesRequest, opts ...grpc.CallOption) (*ListGoReleasesResponse, error)
}

type maintnerServiceClient struct {
	cc *grpc.ClientConn
}

func NewMaintnerServiceClient(cc *grpc.ClientConn) MaintnerServiceClient {
	return &maintnerServiceClient{cc}
}

func (c *maintnerServiceClient) HasAncestor(ctx context.Context, in *HasAncestorRequest, opts ...grpc.CallOption) (*HasAncestorResponse, error) {
	out := new(HasAncestorResponse)
	err := grpc.Invoke(ctx, "/apipb.MaintnerService/HasAncestor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) GetRef(ctx context.Context, in *GetRefRequest, opts ...grpc.CallOption) (*GetRefResponse, error) {
	out := new(GetRefResponse)
	err := grpc.Invoke(ctx, "/apipb.MaintnerService/GetRef", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) GoFindTryWork(ctx context.Context, in *GoFindTryWorkRequest, opts ...grpc.CallOption) (*GoFindTryWorkResponse, error) {
	out := new(GoFindTryWorkResponse)
	err := grpc.Invoke(ctx, "/apipb.MaintnerService/GoFindTryWork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) ListGoReleases(ctx context.Context, in *ListGoReleasesRequest, opts ...grpc.CallOption) (*ListGoReleasesResponse, error) {
	out := new(ListGoReleasesResponse)
	err := grpc.Invoke(ctx, "/apipb.MaintnerService/ListGoReleases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MaintnerService service

type MaintnerServiceServer interface {
	// HasAncestor reports whether one commit contains another commit
	// in its git history.
	HasAncestor(context.Context, *HasAncestorRequest) (*HasAncestorResponse, error)
	// GetRef returns information about a git ref.
	GetRef(context.Context, *GetRefRequest) (*GetRefResponse, error)
	// GoFindTryWork finds trybot work for the coordinator to build & test.
	GoFindTryWork(context.Context, *GoFindTryWorkRequest) (*GoFindTryWorkResponse, error)
	// ListGoReleases lists Go releases. A release is considered to exist for
	// each git tag named "goX", "goX.Y", or "goX.Y.Z", as long as it has a
	// corresponding "release-branch.goX" or "release-branch.goX.Y" release branch.
	ListGoReleases(context.Context, *ListGoReleasesRequest) (*ListGoReleasesResponse, error)
}

func RegisterMaintnerServiceServer(s *grpc.Server, srv MaintnerServiceServer) {
	s.RegisterService(&_MaintnerService_serviceDesc, srv)
}

func _MaintnerService_HasAncestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAncestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).HasAncestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.MaintnerService/HasAncestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).HasAncestor(ctx, req.(*HasAncestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_GetRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).GetRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.MaintnerService/GetRef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).GetRef(ctx, req.(*GetRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_GoFindTryWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoFindTryWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).GoFindTryWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.MaintnerService/GoFindTryWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).GoFindTryWork(ctx, req.(*GoFindTryWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_ListGoReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).ListGoReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.MaintnerService/ListGoReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).ListGoReleases(ctx, req.(*ListGoReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MaintnerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.MaintnerService",
	HandlerType: (*MaintnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasAncestor",
			Handler:    _MaintnerService_HasAncestor_Handler,
		},
		{
			MethodName: "GetRef",
			Handler:    _MaintnerService_GetRef_Handler,
		},
		{
			MethodName: "GoFindTryWork",
			Handler:    _MaintnerService_GoFindTryWork_Handler,
		},
		{
			MethodName: "ListGoReleases",
			Handler:    _MaintnerService_ListGoReleases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0x55, 0x9a, 0x2f, 0x3f, 0x73, 0xd3, 0xf4, 0x2b, 0xa6, 0x2d, 0xe9, 0x94, 0xaa, 0x65, 0x2a,
	0x50, 0x17, 0x28, 0x8b, 0x20, 0xc4, 0x1a, 0x8a, 0x9a, 0x16, 0x28, 0x42, 0x53, 0x24, 0x96, 0x23,
	0x67, 0xe2, 0x4c, 0xdc, 0x76, 0xec, 0xc1, 0xe3, 0xb4, 0xe2, 0x91, 0xd8, 0xf3, 0x1a, 0xbc, 0x13,
	0xb2, 0x7d, 0x3d, 0xe4, 0x8f, 0x5d, 0xee, 0x39, 0xe7, 0xfe, 0xcc, 0xb9, 0xd7, 0x81, 0x80, 0x16,
	0xbc, 0x5f, 0x28, 0xa9, 0x25, 0x69, 0xd0, 0x82, 0x17, 0xa3, 0xe8, 0x02, 0xc8, 0x05, 0x2d, 0xdf,
	0x8a, 0x94, 0x95, 0x5a, 0xaa, 0x98, 0x7d, 0x9f, 0xb1, 0x52, 0x93, 0x3d, 0x68, 0xa6, 0x32, 0xcf,
	0xb9, 0xee, 0xd5, 0x8e, 0x6b, 0xa7, 0x41, 0x8c, 0x11, 0x09, 0xa1, 0x4d, 0x51, 0xda, 0xdb, 0xb0,
	0x4c, 0x15, 0x47, 0x09, 0x3c, 0x5e, 0xa8, 0x54, 0x16, 0x52, 0x94, 0x8c, 0x3c, 0x83, 0xcd, 0x29,
	0x2d, 0x93, 0x2a, 0xcd, 0x14, 0x6c, 0xc7, 0x9d, 0xe9, 0x5f, 0x29, 0x79, 0x0e, 0x5b, 0x33, 0x71,
	0x2b, 0xe4, 0x83, 0x48, 0xb0, 0xeb, 0x86, 0x15, 0x75, 0x11, 0x3d, 0xb3, 0x60, 0x94, 0x43, 0x77,
	0xc8, 0x74, 0xcc, 0x26, 0x7e, 0xca, 0x6d, 0xa8, 0x2b, 0x36, 0xc1, 0x11, 0xcd, 0x4f, 0x72, 0x02,
	0xdd, 0x8c, 0x29, 0xc5, 0x75, 0x52, 0x32, 0x75, 0xcf, 0xfc, 0x90, 0x9b, 0x0e, 0xbc, 0xb6, 0x98,
	0x69, 0x87, 0xa2, 0x42, 0xc9, 0x1b, 0x96, 0xea, 0x5e, 0xdd, 0xaa, 0x30, 0xf5, 0x8b, 0x03, 0xa3,
	0x17, 0xb0, 0xe5, 0xdb, 0xe1, 0xa7, 0xec, 0x40, 0xe3, 0x9e, 0xde, 0xcd, 0x18, 0x76, 0x74, 0x41,
	0xf4, 0x06, 0x76, 0x86, 0xf2, 0x9c, 0x8b, 0xf1, 0x57, 0xf5, 0xe3, 0x9b, 0x54, 0xb7, 0x7e, 0xba,
	0x23, 0xe8, 0x4c, 0xa4, 0x4a, 0x4a, 0x4d, 0x33, 0x2e, 0x32, 0xfc, 0x6e, 0x98, 0x48, 0x75, 0xed,
	0x90, 0xe8, 0x23, 0xec, 0x2e, 0x25, 0x62, 0x9f, 0x01, 0xb4, 0x1e, 0x28, 0xd7, 0x2e, 0xab, 0x7e,
	0xda, 0x19, 0xf4, 0xfa, 0x76, 0x59, 0xfd, 0xa1, 0x1d, 0x10, 0xe5, 0x97, 0x9a, 0xe5, 0xb1, 0x17,
	0x46, 0xbf, 0x6a, 0xf0, 0x68, 0x85, 0x26, 0x3d, 0x68, 0xf9, 0x6f, 0x74, 0x33, 0xfb, 0xd0, 0x6c,
	0x78, 0xa4, 0xa8, 0x48, 0xa7, 0x68, 0x11, 0x46, 0xe4, 0x00, 0x82, 0x74, 0x4a, 0x45, 0xc6, 0x12,
	0x3e, 0x46, 0x5f, 0xda, 0x0e, 0xb8, 0x1c, 0xcf, 0x9d, 0xc5, 0x7f, 0x0b, 0x67, 0x71, 0x00, 0x41,
	0x26, 0xfd, 0xee, 0x1a, 0xc7, 0x75, 0x93, 0x94, 0xc9, 0xb3, 0x79, 0x12, 0x9b, 0x35, 0x3d, 0xf9,
	0xce, 0xc6, 0xd1, 0x13, 0xd8, 0xfd, 0xc4, 0x4b, 0x3d, 0x94, 0x31, 0xbb, 0x63, 0xb4, 0x64, 0x25,
	0xba, 0x17, 0x9d, 0xc3, 0xde, 0x32, 0x81, 0xee, 0xbc, 0x84, 0xb6, 0x42, 0x0c, 0xed, 0xd9, 0xf6,
	0xf6, 0x78, 0x71, 0x5c, 0x29, 0xa2, 0xdf, 0x35, 0x08, 0x2a, 0xdc, 0x6c, 0x30, 0xa7, 0x37, 0x78,
	0x85, 0x8d, 0xd8, 0x05, 0x16, 0xe5, 0x02, 0x4f, 0xda, 0xa0, 0x26, 0x30, 0x68, 0x41, 0x75, 0x3a,
	0xb5, 0x2e, 0x34, 0x62, 0x17, 0x90, 0x7d, 0x68, 0x6b, 0x9a, 0x25, 0x82, 0xe6, 0x0c, 0x4d, 0x68,
	0x69, 0x9a, 0x7d, 0xa6, 0x39, 0x23, 0x87, 0x00, 0x86, 0xaa, 0x6c, 0x30, 0x64, 0xa0, 0x69, 0x86,
	0x3e, 0x1c, 0x41, 0xc7, 0x99, 0xe0, 0x92, 0x9b, 0x96, 0x07, 0x07, 0xd9, 0xfc, 0x13, 0xe8, 0xa2,
	0x00, 0x4b, 0xb4, 0xdc, 0xf1, 0x3a, 0xd0, 0x55, 0x19, 0xfc, 0xdc, 0x80, 0xff, 0xaf, 0x28, 0x17,
	0x5a, 0x30, 0x65, 0xee, 0x99, 0xa7, 0x8c, 0xbc, 0x87, 0xce, 0xdc, 0xcb, 0x23, 0xfb, 0x68, 0xc7,
	0xea, 0xbb, 0x0e, 0xc3, 0x75, 0x14, 0xfa, 0xfa, 0x1a, 0x9a, 0xee, 0xde, 0xc9, 0x4e, 0x75, 0x6e,
	0x73, 0xaf, 0x2d, 0xdc, 0x5d, 0x42, 0x31, 0xed, 0x03, 0x74, 0x17, 0xae, 0x98, 0x1c, 0x54, 0xdb,
	0x58, 0x7d, 0x14, 0xe1, 0xd3, 0xf5, 0x24, 0xd6, 0xba, 0x82, 0xad, 0xc5, 0xa5, 0x13, 0xaf, 0x5f,
	0x7b, 0x24, 0xe1, 0xe1, 0x3f, 0x58, 0x57, 0x6e, 0xd4, 0xb4, 0xff, 0x74, 0xaf, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x13, 0xc3, 0xc8, 0xa7, 0xf6, 0x04, 0x00, 0x00,
}
