// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.27.1
// source: cel/expr/conformance/conformance_service.proto

package conformance

import (
	expr "cel.dev/expr"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type IssueDetails_Severity int32

const (
	IssueDetails_SEVERITY_UNSPECIFIED IssueDetails_Severity = 0
	IssueDetails_DEPRECATION          IssueDetails_Severity = 1
	IssueDetails_WARNING              IssueDetails_Severity = 2
	IssueDetails_ERROR                IssueDetails_Severity = 3
)

// Enum value maps for IssueDetails_Severity.
var (
	IssueDetails_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "DEPRECATION",
		2: "WARNING",
		3: "ERROR",
	}
	IssueDetails_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"DEPRECATION":          1,
		"WARNING":              2,
		"ERROR":                3,
	}
)

func (x IssueDetails_Severity) Enum() *IssueDetails_Severity {
	p := new(IssueDetails_Severity)
	*p = x
	return p
}

func (x IssueDetails_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueDetails_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_cel_expr_conformance_conformance_service_proto_enumTypes[0].Descriptor()
}

func (IssueDetails_Severity) Type() protoreflect.EnumType {
	return &file_cel_expr_conformance_conformance_service_proto_enumTypes[0]
}

func (x IssueDetails_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueDetails_Severity.Descriptor instead.
func (IssueDetails_Severity) EnumDescriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{7, 0}
}

type ParseRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CelSource      string                 `protobuf:"bytes,1,opt,name=cel_source,json=celSource,proto3" json:"cel_source,omitempty"`
	SyntaxVersion  string                 `protobuf:"bytes,2,opt,name=syntax_version,json=syntaxVersion,proto3" json:"syntax_version,omitempty"`
	SourceLocation string                 `protobuf:"bytes,3,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	DisableMacros  bool                   `protobuf:"varint,4,opt,name=disable_macros,json=disableMacros,proto3" json:"disable_macros,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ParseRequest) Reset() {
	*x = ParseRequest{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseRequest) ProtoMessage() {}

func (x *ParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseRequest.ProtoReflect.Descriptor instead.
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{0}
}

func (x *ParseRequest) GetCelSource() string {
	if x != nil {
		return x.CelSource
	}
	return ""
}

func (x *ParseRequest) GetSyntaxVersion() string {
	if x != nil {
		return x.SyntaxVersion
	}
	return ""
}

func (x *ParseRequest) GetSourceLocation() string {
	if x != nil {
		return x.SourceLocation
	}
	return ""
}

func (x *ParseRequest) GetDisableMacros() bool {
	if x != nil {
		return x.DisableMacros
	}
	return false
}

type ParseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParsedExpr    *expr.ParsedExpr       `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	Issues        []*status.Status       `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseResponse) Reset() {
	*x = ParseResponse{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseResponse) ProtoMessage() {}

func (x *ParseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseResponse.ProtoReflect.Descriptor instead.
func (*ParseResponse) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{1}
}

func (x *ParseResponse) GetParsedExpr() *expr.ParsedExpr {
	if x != nil {
		return x.ParsedExpr
	}
	return nil
}

func (x *ParseResponse) GetIssues() []*status.Status {
	if x != nil {
		return x.Issues
	}
	return nil
}

type CheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParsedExpr    *expr.ParsedExpr       `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	TypeEnv       []*expr.Decl           `protobuf:"bytes,2,rep,name=type_env,json=typeEnv,proto3" json:"type_env,omitempty"`
	Container     string                 `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
	NoStdEnv      bool                   `protobuf:"varint,4,opt,name=no_std_env,json=noStdEnv,proto3" json:"no_std_env,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRequest) GetParsedExpr() *expr.ParsedExpr {
	if x != nil {
		return x.ParsedExpr
	}
	return nil
}

func (x *CheckRequest) GetTypeEnv() []*expr.Decl {
	if x != nil {
		return x.TypeEnv
	}
	return nil
}

func (x *CheckRequest) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *CheckRequest) GetNoStdEnv() bool {
	if x != nil {
		return x.NoStdEnv
	}
	return false
}

type CheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CheckedExpr   *expr.CheckedExpr      `protobuf:"bytes,1,opt,name=checked_expr,json=checkedExpr,proto3" json:"checked_expr,omitempty"`
	Issues        []*status.Status       `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckResponse) GetCheckedExpr() *expr.CheckedExpr {
	if x != nil {
		return x.CheckedExpr
	}
	return nil
}

func (x *CheckResponse) GetIssues() []*status.Status {
	if x != nil {
		return x.Issues
	}
	return nil
}

type EvalRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to ExprKind:
	//
	//	*EvalRequest_ParsedExpr
	//	*EvalRequest_CheckedExpr
	ExprKind      isEvalRequest_ExprKind     `protobuf_oneof:"expr_kind"`
	Bindings      map[string]*expr.ExprValue `protobuf:"bytes,3,rep,name=bindings,proto3" json:"bindings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Container     string                     `protobuf:"bytes,4,opt,name=container,proto3" json:"container,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvalRequest) Reset() {
	*x = EvalRequest{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalRequest) ProtoMessage() {}

func (x *EvalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalRequest.ProtoReflect.Descriptor instead.
func (*EvalRequest) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{4}
}

func (x *EvalRequest) GetExprKind() isEvalRequest_ExprKind {
	if x != nil {
		return x.ExprKind
	}
	return nil
}

func (x *EvalRequest) GetParsedExpr() *expr.ParsedExpr {
	if x != nil {
		if x, ok := x.ExprKind.(*EvalRequest_ParsedExpr); ok {
			return x.ParsedExpr
		}
	}
	return nil
}

func (x *EvalRequest) GetCheckedExpr() *expr.CheckedExpr {
	if x != nil {
		if x, ok := x.ExprKind.(*EvalRequest_CheckedExpr); ok {
			return x.CheckedExpr
		}
	}
	return nil
}

func (x *EvalRequest) GetBindings() map[string]*expr.ExprValue {
	if x != nil {
		return x.Bindings
	}
	return nil
}

func (x *EvalRequest) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

type isEvalRequest_ExprKind interface {
	isEvalRequest_ExprKind()
}

type EvalRequest_ParsedExpr struct {
	ParsedExpr *expr.ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3,oneof"`
}

type EvalRequest_CheckedExpr struct {
	CheckedExpr *expr.CheckedExpr `protobuf:"bytes,2,opt,name=checked_expr,json=checkedExpr,proto3,oneof"`
}

func (*EvalRequest_ParsedExpr) isEvalRequest_ExprKind() {}

func (*EvalRequest_CheckedExpr) isEvalRequest_ExprKind() {}

type EvalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *expr.ExprValue        `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Issues        []*status.Status       `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvalResponse) Reset() {
	*x = EvalResponse{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalResponse) ProtoMessage() {}

func (x *EvalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalResponse.ProtoReflect.Descriptor instead.
func (*EvalResponse) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{5}
}

func (x *EvalResponse) GetResult() *expr.ExprValue {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *EvalResponse) GetIssues() []*status.Status {
	if x != nil {
		return x.Issues
	}
	return nil
}

type SourcePosition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Location      string                 `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Line          int32                  `protobuf:"varint,3,opt,name=line,proto3" json:"line,omitempty"`
	Column        int32                  `protobuf:"varint,4,opt,name=column,proto3" json:"column,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SourcePosition) Reset() {
	*x = SourcePosition{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourcePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePosition) ProtoMessage() {}

func (x *SourcePosition) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePosition.ProtoReflect.Descriptor instead.
func (*SourcePosition) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{6}
}

func (x *SourcePosition) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SourcePosition) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SourcePosition) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourcePosition) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

type IssueDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Severity      IssueDetails_Severity  `protobuf:"varint,1,opt,name=severity,proto3,enum=cel.expr.conformance.IssueDetails_Severity" json:"severity,omitempty"`
	Position      *SourcePosition        `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Id            int64                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueDetails) Reset() {
	*x = IssueDetails{}
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueDetails) ProtoMessage() {}

func (x *IssueDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_conformance_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueDetails.ProtoReflect.Descriptor instead.
func (*IssueDetails) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_conformance_service_proto_rawDescGZIP(), []int{7}
}

func (x *IssueDetails) GetSeverity() IssueDetails_Severity {
	if x != nil {
		return x.Severity
	}
	return IssueDetails_SEVERITY_UNSPECIFIED
}

func (x *IssueDetails) GetPosition() *SourcePosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *IssueDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_cel_expr_conformance_conformance_service_proto protoreflect.FileDescriptor

var file_cel_expr_conformance_conformance_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x73, 0x79,
	0x6e, 0x74, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x6c, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d,
	0x61, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x22, 0x72, 0x0a, 0x0d, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x45, 0x78,
	0x70, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x22, 0xac,
	0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65,
	0x6e, 0x76, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x76, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x0a, 0x6e, 0x6f, 0x5f, 0x73, 0x74, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x6f, 0x53, 0x74, 0x64, 0x45, 0x6e, 0x76, 0x22, 0x75, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x22, 0xcc, 0x02, 0x0a, 0x0b, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c, 0x2e,
	0x65, 0x78, 0x70, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x48,
	0x00, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x12, 0x3a, 0x0a,
	0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x12, 0x4b, 0x0a, 0x08, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x65,
	0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x1a, 0x50, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x72, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x22, 0x67, 0x0a, 0x0c, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x2a, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0xf8,
	0x01, 0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x47, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x65, 0x6c,
	0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x08, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x32, 0x8d, 0x02, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x52, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x65, 0x6c, 0x2e,
	0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x22, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x04, 0x45, 0x76, 0x61, 0x6c,
	0x12, 0x21, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x52, 0x0a, 0x18, 0x64, 0x65, 0x76,
	0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x18, 0x63, 0x65, 0x6c, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cel_expr_conformance_conformance_service_proto_rawDescOnce sync.Once
	file_cel_expr_conformance_conformance_service_proto_rawDescData = file_cel_expr_conformance_conformance_service_proto_rawDesc
)

func file_cel_expr_conformance_conformance_service_proto_rawDescGZIP() []byte {
	file_cel_expr_conformance_conformance_service_proto_rawDescOnce.Do(func() {
		file_cel_expr_conformance_conformance_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cel_expr_conformance_conformance_service_proto_rawDescData)
	})
	return file_cel_expr_conformance_conformance_service_proto_rawDescData
}

var file_cel_expr_conformance_conformance_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cel_expr_conformance_conformance_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cel_expr_conformance_conformance_service_proto_goTypes = []any{
	(IssueDetails_Severity)(0), // 0: cel.expr.conformance.IssueDetails.Severity
	(*ParseRequest)(nil),       // 1: cel.expr.conformance.ParseRequest
	(*ParseResponse)(nil),      // 2: cel.expr.conformance.ParseResponse
	(*CheckRequest)(nil),       // 3: cel.expr.conformance.CheckRequest
	(*CheckResponse)(nil),      // 4: cel.expr.conformance.CheckResponse
	(*EvalRequest)(nil),        // 5: cel.expr.conformance.EvalRequest
	(*EvalResponse)(nil),       // 6: cel.expr.conformance.EvalResponse
	(*SourcePosition)(nil),     // 7: cel.expr.conformance.SourcePosition
	(*IssueDetails)(nil),       // 8: cel.expr.conformance.IssueDetails
	nil,                        // 9: cel.expr.conformance.EvalRequest.BindingsEntry
	(*expr.ParsedExpr)(nil),    // 10: cel.expr.ParsedExpr
	(*status.Status)(nil),      // 11: google.rpc.Status
	(*expr.Decl)(nil),          // 12: cel.expr.Decl
	(*expr.CheckedExpr)(nil),   // 13: cel.expr.CheckedExpr
	(*expr.ExprValue)(nil),     // 14: cel.expr.ExprValue
}
var file_cel_expr_conformance_conformance_service_proto_depIdxs = []int32{
	10, // 0: cel.expr.conformance.ParseResponse.parsed_expr:type_name -> cel.expr.ParsedExpr
	11, // 1: cel.expr.conformance.ParseResponse.issues:type_name -> google.rpc.Status
	10, // 2: cel.expr.conformance.CheckRequest.parsed_expr:type_name -> cel.expr.ParsedExpr
	12, // 3: cel.expr.conformance.CheckRequest.type_env:type_name -> cel.expr.Decl
	13, // 4: cel.expr.conformance.CheckResponse.checked_expr:type_name -> cel.expr.CheckedExpr
	11, // 5: cel.expr.conformance.CheckResponse.issues:type_name -> google.rpc.Status
	10, // 6: cel.expr.conformance.EvalRequest.parsed_expr:type_name -> cel.expr.ParsedExpr
	13, // 7: cel.expr.conformance.EvalRequest.checked_expr:type_name -> cel.expr.CheckedExpr
	9,  // 8: cel.expr.conformance.EvalRequest.bindings:type_name -> cel.expr.conformance.EvalRequest.BindingsEntry
	14, // 9: cel.expr.conformance.EvalResponse.result:type_name -> cel.expr.ExprValue
	11, // 10: cel.expr.conformance.EvalResponse.issues:type_name -> google.rpc.Status
	0,  // 11: cel.expr.conformance.IssueDetails.severity:type_name -> cel.expr.conformance.IssueDetails.Severity
	7,  // 12: cel.expr.conformance.IssueDetails.position:type_name -> cel.expr.conformance.SourcePosition
	14, // 13: cel.expr.conformance.EvalRequest.BindingsEntry.value:type_name -> cel.expr.ExprValue
	1,  // 14: cel.expr.conformance.ConformanceService.Parse:input_type -> cel.expr.conformance.ParseRequest
	3,  // 15: cel.expr.conformance.ConformanceService.Check:input_type -> cel.expr.conformance.CheckRequest
	5,  // 16: cel.expr.conformance.ConformanceService.Eval:input_type -> cel.expr.conformance.EvalRequest
	2,  // 17: cel.expr.conformance.ConformanceService.Parse:output_type -> cel.expr.conformance.ParseResponse
	4,  // 18: cel.expr.conformance.ConformanceService.Check:output_type -> cel.expr.conformance.CheckResponse
	6,  // 19: cel.expr.conformance.ConformanceService.Eval:output_type -> cel.expr.conformance.EvalResponse
	17, // [17:20] is the sub-list for method output_type
	14, // [14:17] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_cel_expr_conformance_conformance_service_proto_init() }
func file_cel_expr_conformance_conformance_service_proto_init() {
	if File_cel_expr_conformance_conformance_service_proto != nil {
		return
	}
	file_cel_expr_conformance_conformance_service_proto_msgTypes[4].OneofWrappers = []any{
		(*EvalRequest_ParsedExpr)(nil),
		(*EvalRequest_CheckedExpr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cel_expr_conformance_conformance_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cel_expr_conformance_conformance_service_proto_goTypes,
		DependencyIndexes: file_cel_expr_conformance_conformance_service_proto_depIdxs,
		EnumInfos:         file_cel_expr_conformance_conformance_service_proto_enumTypes,
		MessageInfos:      file_cel_expr_conformance_conformance_service_proto_msgTypes,
	}.Build()
	File_cel_expr_conformance_conformance_service_proto = out.File
	file_cel_expr_conformance_conformance_service_proto_rawDesc = nil
	file_cel_expr_conformance_conformance_service_proto_goTypes = nil
	file_cel_expr_conformance_conformance_service_proto_depIdxs = nil
}
