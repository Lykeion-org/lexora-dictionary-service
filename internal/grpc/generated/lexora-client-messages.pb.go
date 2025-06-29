// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: lexora-client-messages.proto

package lexora

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LookUpReferentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	SearchMode    string                 `protobuf:"bytes,2,opt,name=search_mode,json=searchMode,proto3" json:"search_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookUpReferentRequest) Reset() {
	*x = LookUpReferentRequest{}
	mi := &file_lexora_client_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookUpReferentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookUpReferentRequest) ProtoMessage() {}

func (x *LookUpReferentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookUpReferentRequest.ProtoReflect.Descriptor instead.
func (*LookUpReferentRequest) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{0}
}

func (x *LookUpReferentRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *LookUpReferentRequest) GetSearchMode() string {
	if x != nil {
		return x.SearchMode
	}
	return ""
}

type LookUpReferentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Referents     []*Referent            `protobuf:"bytes,1,rep,name=referents,proto3" json:"referents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookUpReferentResponse) Reset() {
	*x = LookUpReferentResponse{}
	mi := &file_lexora_client_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookUpReferentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookUpReferentResponse) ProtoMessage() {}

func (x *LookUpReferentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookUpReferentResponse.ProtoReflect.Descriptor instead.
func (*LookUpReferentResponse) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{1}
}

func (x *LookUpReferentResponse) GetReferents() []*Referent {
	if x != nil {
		return x.Referents
	}
	return nil
}

type LookupSymbolRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	SearchMode    string                 `protobuf:"bytes,2,opt,name=search_mode,json=searchMode,proto3" json:"search_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupSymbolRequest) Reset() {
	*x = LookupSymbolRequest{}
	mi := &file_lexora_client_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupSymbolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSymbolRequest) ProtoMessage() {}

func (x *LookupSymbolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSymbolRequest.ProtoReflect.Descriptor instead.
func (*LookupSymbolRequest) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{2}
}

func (x *LookupSymbolRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *LookupSymbolRequest) GetSearchMode() string {
	if x != nil {
		return x.SearchMode
	}
	return ""
}

type LookupSymbolResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbols       []*Symbol              `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupSymbolResponse) Reset() {
	*x = LookupSymbolResponse{}
	mi := &file_lexora_client_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupSymbolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSymbolResponse) ProtoMessage() {}

func (x *LookupSymbolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSymbolResponse.ProtoReflect.Descriptor instead.
func (*LookupSymbolResponse) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{3}
}

func (x *LookupSymbolResponse) GetSymbols() []*Symbol {
	if x != nil {
		return x.Symbols
	}
	return nil
}

type LookupWordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	SearchMode    string                 `protobuf:"bytes,2,opt,name=search_mode,json=searchMode,proto3" json:"search_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupWordRequest) Reset() {
	*x = LookupWordRequest{}
	mi := &file_lexora_client_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupWordRequest) ProtoMessage() {}

func (x *LookupWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupWordRequest.ProtoReflect.Descriptor instead.
func (*LookupWordRequest) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{4}
}

func (x *LookupWordRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *LookupWordRequest) GetSearchMode() string {
	if x != nil {
		return x.SearchMode
	}
	return ""
}

type LookupWordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Words         []*Word                `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupWordResponse) Reset() {
	*x = LookupWordResponse{}
	mi := &file_lexora_client_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupWordResponse) ProtoMessage() {}

func (x *LookupWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupWordResponse.ProtoReflect.Descriptor instead.
func (*LookupWordResponse) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{5}
}

func (x *LookupWordResponse) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

type AnalyzeTextRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Language      Language               `protobuf:"varint,3,opt,name=language,proto3,enum=lexora.Language" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalyzeTextRequest) Reset() {
	*x = AnalyzeTextRequest{}
	mi := &file_lexora_client_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeTextRequest) ProtoMessage() {}

func (x *AnalyzeTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeTextRequest.ProtoReflect.Descriptor instead.
func (*AnalyzeTextRequest) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{6}
}

func (x *AnalyzeTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AnalyzeTextRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnalyzeTextRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

type AnalyzeTextResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalyzeTextResponse) Reset() {
	*x = AnalyzeTextResponse{}
	mi := &file_lexora_client_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeTextResponse) ProtoMessage() {}

func (x *AnalyzeTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lexora_client_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeTextResponse.ProtoReflect.Descriptor instead.
func (*AnalyzeTextResponse) Descriptor() ([]byte, []int) {
	return file_lexora_client_messages_proto_rawDescGZIP(), []int{7}
}

func (x *AnalyzeTextResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *AnalyzeTextResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnalyzeTextResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_lexora_client_messages_proto protoreflect.FileDescriptor

const file_lexora_client_messages_proto_rawDesc = "" +
	"\n" +
	"\x1clexora-client-messages.proto\x12\x06lexora\x1a\x14language-types.proto\"N\n" +
	"\x15LookUpReferentRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12\x1f\n" +
	"\vsearch_mode\x18\x02 \x01(\tR\n" +
	"searchMode\"H\n" +
	"\x16LookUpReferentResponse\x12.\n" +
	"\treferents\x18\x01 \x03(\v2\x10.lexora.ReferentR\treferents\"L\n" +
	"\x13LookupSymbolRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12\x1f\n" +
	"\vsearch_mode\x18\x02 \x01(\tR\n" +
	"searchMode\"@\n" +
	"\x14LookupSymbolResponse\x12(\n" +
	"\asymbols\x18\x01 \x03(\v2\x0e.lexora.SymbolR\asymbols\"J\n" +
	"\x11LookupWordRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12\x1f\n" +
	"\vsearch_mode\x18\x02 \x01(\tR\n" +
	"searchMode\"8\n" +
	"\x12LookupWordResponse\x12\"\n" +
	"\x05words\x18\x01 \x03(\v2\f.lexora.WordR\x05words\"o\n" +
	"\x12AnalyzeTextRequest\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12,\n" +
	"\blanguage\x18\x03 \x01(\x0e2\x10.lexora.LanguageR\blanguage\"]\n" +
	"\x13AnalyzeTextResponse\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06statusb\x06proto3"

var (
	file_lexora_client_messages_proto_rawDescOnce sync.Once
	file_lexora_client_messages_proto_rawDescData []byte
)

func file_lexora_client_messages_proto_rawDescGZIP() []byte {
	file_lexora_client_messages_proto_rawDescOnce.Do(func() {
		file_lexora_client_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lexora_client_messages_proto_rawDesc), len(file_lexora_client_messages_proto_rawDesc)))
	})
	return file_lexora_client_messages_proto_rawDescData
}

var file_lexora_client_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lexora_client_messages_proto_goTypes = []any{
	(*LookUpReferentRequest)(nil),  // 0: lexora.LookUpReferentRequest
	(*LookUpReferentResponse)(nil), // 1: lexora.LookUpReferentResponse
	(*LookupSymbolRequest)(nil),    // 2: lexora.LookupSymbolRequest
	(*LookupSymbolResponse)(nil),   // 3: lexora.LookupSymbolResponse
	(*LookupWordRequest)(nil),      // 4: lexora.LookupWordRequest
	(*LookupWordResponse)(nil),     // 5: lexora.LookupWordResponse
	(*AnalyzeTextRequest)(nil),     // 6: lexora.AnalyzeTextRequest
	(*AnalyzeTextResponse)(nil),    // 7: lexora.AnalyzeTextResponse
	(*Referent)(nil),               // 8: lexora.Referent
	(*Symbol)(nil),                 // 9: lexora.Symbol
	(*Word)(nil),                   // 10: lexora.Word
	(Language)(0),                  // 11: lexora.Language
}
var file_lexora_client_messages_proto_depIdxs = []int32{
	8,  // 0: lexora.LookUpReferentResponse.referents:type_name -> lexora.Referent
	9,  // 1: lexora.LookupSymbolResponse.symbols:type_name -> lexora.Symbol
	10, // 2: lexora.LookupWordResponse.words:type_name -> lexora.Word
	11, // 3: lexora.AnalyzeTextRequest.language:type_name -> lexora.Language
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_lexora_client_messages_proto_init() }
func file_lexora_client_messages_proto_init() {
	if File_lexora_client_messages_proto != nil {
		return
	}
	file_language_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lexora_client_messages_proto_rawDesc), len(file_lexora_client_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lexora_client_messages_proto_goTypes,
		DependencyIndexes: file_lexora_client_messages_proto_depIdxs,
		MessageInfos:      file_lexora_client_messages_proto_msgTypes,
	}.Build()
	File_lexora_client_messages_proto = out.File
	file_lexora_client_messages_proto_goTypes = nil
	file_lexora_client_messages_proto_depIdxs = nil
}
