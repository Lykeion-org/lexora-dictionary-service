// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: language-types.proto

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

type WordType int32

const (
	WordType_WORD_TYPE_UNSPECIFIED  WordType = 0
	WordType_WORD_TYPE_NOUN         WordType = 1
	WordType_WORD_TYPE_VERB         WordType = 2
	WordType_WORD_TYPE_ADJECTIVE    WordType = 3
	WordType_WORD_TYPE_ADVERB       WordType = 4
	WordType_WORD_TYPE_PRONOUN      WordType = 5
	WordType_WORD_TYPE_PREPOSITION  WordType = 6
	WordType_WORD_TYPE_CONJUNCTION  WordType = 7
	WordType_WORD_TYPE_INTERJECTION WordType = 8
	WordType_WORD_TYPE_DETERMINER   WordType = 9
)

// Enum value maps for WordType.
var (
	WordType_name = map[int32]string{
		0: "WORD_TYPE_UNSPECIFIED",
		1: "WORD_TYPE_NOUN",
		2: "WORD_TYPE_VERB",
		3: "WORD_TYPE_ADJECTIVE",
		4: "WORD_TYPE_ADVERB",
		5: "WORD_TYPE_PRONOUN",
		6: "WORD_TYPE_PREPOSITION",
		7: "WORD_TYPE_CONJUNCTION",
		8: "WORD_TYPE_INTERJECTION",
		9: "WORD_TYPE_DETERMINER",
	}
	WordType_value = map[string]int32{
		"WORD_TYPE_UNSPECIFIED":  0,
		"WORD_TYPE_NOUN":         1,
		"WORD_TYPE_VERB":         2,
		"WORD_TYPE_ADJECTIVE":    3,
		"WORD_TYPE_ADVERB":       4,
		"WORD_TYPE_PRONOUN":      5,
		"WORD_TYPE_PREPOSITION":  6,
		"WORD_TYPE_CONJUNCTION":  7,
		"WORD_TYPE_INTERJECTION": 8,
		"WORD_TYPE_DETERMINER":   9,
	}
)

func (x WordType) Enum() *WordType {
	p := new(WordType)
	*p = x
	return p
}

func (x WordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordType) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[0].Descriptor()
}

func (WordType) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[0]
}

func (x WordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordType.Descriptor instead.
func (WordType) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{0}
}

type SymbolType int32

const (
	SymbolType_SYMBOL_TYPE_OBJECT   SymbolType = 0
	SymbolType_SYMBOL_TYPE_ACTION   SymbolType = 1
	SymbolType_SYMBOL_TYPE_MODIFIER SymbolType = 2
)

// Enum value maps for SymbolType.
var (
	SymbolType_name = map[int32]string{
		0: "SYMBOL_TYPE_OBJECT",
		1: "SYMBOL_TYPE_ACTION",
		2: "SYMBOL_TYPE_MODIFIER",
	}
	SymbolType_value = map[string]int32{
		"SYMBOL_TYPE_OBJECT":   0,
		"SYMBOL_TYPE_ACTION":   1,
		"SYMBOL_TYPE_MODIFIER": 2,
	}
)

func (x SymbolType) Enum() *SymbolType {
	p := new(SymbolType)
	*p = x
	return p
}

func (x SymbolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymbolType) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[1].Descriptor()
}

func (SymbolType) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[1]
}

func (x SymbolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymbolType.Descriptor instead.
func (SymbolType) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{1}
}

type Language int32

const (
	Language_LANGUAGE_UNSPECIFIED Language = 0
	Language_LANGUAGE_EN          Language = 1
	Language_LANGUAGE_NL          Language = 2
	Language_LANGUAGE_ES          Language = 3
	Language_LANGUAGE_PT          Language = 4
	Language_LANGUAGE_DE          Language = 5
	Language_LANGUAGE_FR          Language = 6
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "LANGUAGE_UNSPECIFIED",
		1: "LANGUAGE_EN",
		2: "LANGUAGE_NL",
		3: "LANGUAGE_ES",
		4: "LANGUAGE_PT",
		5: "LANGUAGE_DE",
		6: "LANGUAGE_FR",
	}
	Language_value = map[string]int32{
		"LANGUAGE_UNSPECIFIED": 0,
		"LANGUAGE_EN":          1,
		"LANGUAGE_NL":          2,
		"LANGUAGE_ES":          3,
		"LANGUAGE_PT":          4,
		"LANGUAGE_DE":          5,
		"LANGUAGE_FR":          6,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[2].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[2]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{2}
}

type WordTense int32

const (
	WordTense_WORD_TENSE_UNSPECFIED WordTense = 0
	WordTense_WORD_TENSE_PAST       WordTense = 1
	WordTense_WORD_TENSE_PRESENT    WordTense = 2
	WordTense_WORD_TENSE_FUTURE     WordTense = 3
)

// Enum value maps for WordTense.
var (
	WordTense_name = map[int32]string{
		0: "WORD_TENSE_UNSPECFIED",
		1: "WORD_TENSE_PAST",
		2: "WORD_TENSE_PRESENT",
		3: "WORD_TENSE_FUTURE",
	}
	WordTense_value = map[string]int32{
		"WORD_TENSE_UNSPECFIED": 0,
		"WORD_TENSE_PAST":       1,
		"WORD_TENSE_PRESENT":    2,
		"WORD_TENSE_FUTURE":     3,
	}
)

func (x WordTense) Enum() *WordTense {
	p := new(WordTense)
	*p = x
	return p
}

func (x WordTense) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordTense) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[3].Descriptor()
}

func (WordTense) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[3]
}

func (x WordTense) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordTense.Descriptor instead.
func (WordTense) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{3}
}

type WordGender int32

const (
	WordGender_WORD_GENDER_UNSPECIFIED WordGender = 0
	WordGender_WORD_GENDER_MALE        WordGender = 1
	WordGender_WORD_GENDER_FEMALE      WordGender = 2
	WordGender_WORD_GENDER_NEUTER      WordGender = 3
)

// Enum value maps for WordGender.
var (
	WordGender_name = map[int32]string{
		0: "WORD_GENDER_UNSPECIFIED",
		1: "WORD_GENDER_MALE",
		2: "WORD_GENDER_FEMALE",
		3: "WORD_GENDER_NEUTER",
	}
	WordGender_value = map[string]int32{
		"WORD_GENDER_UNSPECIFIED": 0,
		"WORD_GENDER_MALE":        1,
		"WORD_GENDER_FEMALE":      2,
		"WORD_GENDER_NEUTER":      3,
	}
)

func (x WordGender) Enum() *WordGender {
	p := new(WordGender)
	*p = x
	return p
}

func (x WordGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordGender) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[4].Descriptor()
}

func (WordGender) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[4]
}

func (x WordGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordGender.Descriptor instead.
func (WordGender) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{4}
}

type WordNumber int32

const (
	WordNumber_WORD_NUMBER_UNSPECIFIED WordNumber = 0
	WordNumber_WORD_NUMBER_SINGULAR    WordNumber = 1
	WordNumber_WORD_NUMBER_PLURAL      WordNumber = 2
)

// Enum value maps for WordNumber.
var (
	WordNumber_name = map[int32]string{
		0: "WORD_NUMBER_UNSPECIFIED",
		1: "WORD_NUMBER_SINGULAR",
		2: "WORD_NUMBER_PLURAL",
	}
	WordNumber_value = map[string]int32{
		"WORD_NUMBER_UNSPECIFIED": 0,
		"WORD_NUMBER_SINGULAR":    1,
		"WORD_NUMBER_PLURAL":      2,
	}
)

func (x WordNumber) Enum() *WordNumber {
	p := new(WordNumber)
	*p = x
	return p
}

func (x WordNumber) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordNumber) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[5].Descriptor()
}

func (WordNumber) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[5]
}

func (x WordNumber) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordNumber.Descriptor instead.
func (WordNumber) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{5}
}

type WordMood int32

const (
	WordMood_WORD_MOOD_UNSPECIFIED WordMood = 0
)

// Enum value maps for WordMood.
var (
	WordMood_name = map[int32]string{
		0: "WORD_MOOD_UNSPECIFIED",
	}
	WordMood_value = map[string]int32{
		"WORD_MOOD_UNSPECIFIED": 0,
	}
)

func (x WordMood) Enum() *WordMood {
	p := new(WordMood)
	*p = x
	return p
}

func (x WordMood) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordMood) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[6].Descriptor()
}

func (WordMood) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[6]
}

func (x WordMood) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordMood.Descriptor instead.
func (WordMood) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{6}
}

type WordAspect int32

const (
	WordAspect_WORD_ASPECT_UNSPECIFIED WordAspect = 0
)

// Enum value maps for WordAspect.
var (
	WordAspect_name = map[int32]string{
		0: "WORD_ASPECT_UNSPECIFIED",
	}
	WordAspect_value = map[string]int32{
		"WORD_ASPECT_UNSPECIFIED": 0,
	}
)

func (x WordAspect) Enum() *WordAspect {
	p := new(WordAspect)
	*p = x
	return p
}

func (x WordAspect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordAspect) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[7].Descriptor()
}

func (WordAspect) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[7]
}

func (x WordAspect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordAspect.Descriptor instead.
func (WordAspect) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{7}
}

type WordPerson int32

const (
	WordPerson_WORD_PERSON_UNSPECIFIED WordPerson = 0
	WordPerson_WORD_PERSON_FIRST       WordPerson = 1
	WordPerson_WORD_PERSON_SECOND      WordPerson = 2
	WordPerson_WORD_PERSON_THIRD       WordPerson = 3
)

// Enum value maps for WordPerson.
var (
	WordPerson_name = map[int32]string{
		0: "WORD_PERSON_UNSPECIFIED",
		1: "WORD_PERSON_FIRST",
		2: "WORD_PERSON_SECOND",
		3: "WORD_PERSON_THIRD",
	}
	WordPerson_value = map[string]int32{
		"WORD_PERSON_UNSPECIFIED": 0,
		"WORD_PERSON_FIRST":       1,
		"WORD_PERSON_SECOND":      2,
		"WORD_PERSON_THIRD":       3,
	}
)

func (x WordPerson) Enum() *WordPerson {
	p := new(WordPerson)
	*p = x
	return p
}

func (x WordPerson) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordPerson) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[8].Descriptor()
}

func (WordPerson) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[8]
}

func (x WordPerson) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordPerson.Descriptor instead.
func (WordPerson) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{8}
}

type WordCase int32

const (
	WordCase_WORD_CASE_UNSPECIFIED  WordCase = 0
	WordCase_WORD_CASE_NOMINATIVE   WordCase = 1
	WordCase_WORD_CASE_GENITIVE     WordCase = 2
	WordCase_WORD_CASE_DATIVE       WordCase = 3
	WordCase_WORD_CASE_ACCUSATIVE   WordCase = 4
	WordCase_WORD_CASE_INSTRUMENTAL WordCase = 5
	WordCase_WORD_CASE_LOCATIVE     WordCase = 6
	WordCase_WORD_CASE_VOCATIVE     WordCase = 7
	WordCase_WORD_CASE_ABLATIVE     WordCase = 8
)

// Enum value maps for WordCase.
var (
	WordCase_name = map[int32]string{
		0: "WORD_CASE_UNSPECIFIED",
		1: "WORD_CASE_NOMINATIVE",
		2: "WORD_CASE_GENITIVE",
		3: "WORD_CASE_DATIVE",
		4: "WORD_CASE_ACCUSATIVE",
		5: "WORD_CASE_INSTRUMENTAL",
		6: "WORD_CASE_LOCATIVE",
		7: "WORD_CASE_VOCATIVE",
		8: "WORD_CASE_ABLATIVE",
	}
	WordCase_value = map[string]int32{
		"WORD_CASE_UNSPECIFIED":  0,
		"WORD_CASE_NOMINATIVE":   1,
		"WORD_CASE_GENITIVE":     2,
		"WORD_CASE_DATIVE":       3,
		"WORD_CASE_ACCUSATIVE":   4,
		"WORD_CASE_INSTRUMENTAL": 5,
		"WORD_CASE_LOCATIVE":     6,
		"WORD_CASE_VOCATIVE":     7,
		"WORD_CASE_ABLATIVE":     8,
	}
)

func (x WordCase) Enum() *WordCase {
	p := new(WordCase)
	*p = x
	return p
}

func (x WordCase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordCase) Descriptor() protoreflect.EnumDescriptor {
	return file_language_types_proto_enumTypes[9].Descriptor()
}

func (WordCase) Type() protoreflect.EnumType {
	return &file_language_types_proto_enumTypes[9]
}

func (x WordCase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordCase.Descriptor instead.
func (WordCase) EnumDescriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{9}
}

type Word struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Word          string                 `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	SoundSource   string                 `protobuf:"bytes,3,opt,name=sound_source,json=soundSource,proto3" json:"sound_source,omitempty"`
	Ipa           string                 `protobuf:"bytes,4,opt,name=ipa,proto3" json:"ipa,omitempty"`
	WordType      WordType               `protobuf:"varint,5,opt,name=word_type,json=wordType,proto3,enum=lexora.WordType" json:"word_type,omitempty"`
	Attributes    *WordAttributes        `protobuf:"bytes,6,opt,name=attributes,proto3" json:"attributes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Word) Reset() {
	*x = Word{}
	mi := &file_language_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{0}
}

func (x *Word) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Word) GetSoundSource() string {
	if x != nil {
		return x.SoundSource
	}
	return ""
}

func (x *Word) GetIpa() string {
	if x != nil {
		return x.Ipa
	}
	return ""
}

func (x *Word) GetWordType() WordType {
	if x != nil {
		return x.WordType
	}
	return WordType_WORD_TYPE_UNSPECIFIED
}

func (x *Word) GetAttributes() *WordAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Symbol struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Language      Language               `protobuf:"varint,2,opt,name=language,proto3,enum=lexora.Language" json:"language,omitempty"`
	SymbolType    SymbolType             `protobuf:"varint,3,opt,name=symbol_type,json=symbolType,proto3,enum=lexora.SymbolType" json:"symbol_type,omitempty"`
	Lemma         *Word                  `protobuf:"bytes,4,opt,name=lemma,proto3" json:"lemma,omitempty"`
	Words         []*Word                `protobuf:"bytes,5,rep,name=words,proto3" json:"words,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Symbol) Reset() {
	*x = Symbol{}
	mi := &file_language_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Symbol) ProtoMessage() {}

func (x *Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Symbol.ProtoReflect.Descriptor instead.
func (*Symbol) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{1}
}

func (x *Symbol) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Symbol) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

func (x *Symbol) GetSymbolType() SymbolType {
	if x != nil {
		return x.SymbolType
	}
	return SymbolType_SYMBOL_TYPE_OBJECT
}

func (x *Symbol) GetLemma() *Word {
	if x != nil {
		return x.Lemma
	}
	return nil
}

func (x *Symbol) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

type Referent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EnReference   string                 `protobuf:"bytes,2,opt,name=en_reference,json=enReference,proto3" json:"en_reference,omitempty"`
	ImageSource   string                 `protobuf:"bytes,3,opt,name=image_source,json=imageSource,proto3" json:"image_source,omitempty"`
	Symbols       []*Symbol              `protobuf:"bytes,4,rep,name=symbols,proto3" json:"symbols,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Referent) Reset() {
	*x = Referent{}
	mi := &file_language_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Referent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Referent) ProtoMessage() {}

func (x *Referent) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Referent.ProtoReflect.Descriptor instead.
func (*Referent) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{2}
}

func (x *Referent) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Referent) GetEnReference() string {
	if x != nil {
		return x.EnReference
	}
	return ""
}

func (x *Referent) GetImageSource() string {
	if x != nil {
		return x.ImageSource
	}
	return ""
}

func (x *Referent) GetSymbols() []*Symbol {
	if x != nil {
		return x.Symbols
	}
	return nil
}

type ReferentAttributes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReferentAttributes) Reset() {
	*x = ReferentAttributes{}
	mi := &file_language_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReferentAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferentAttributes) ProtoMessage() {}

func (x *ReferentAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferentAttributes.ProtoReflect.Descriptor instead.
func (*ReferentAttributes) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{3}
}

type SymbolAttributes struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	LanguageVariation int64                  `protobuf:"varint,1,opt,name=language_variation,json=languageVariation,proto3" json:"language_variation,omitempty"`
	OriginLanguage    string                 `protobuf:"bytes,2,opt,name=origin_language,json=originLanguage,proto3" json:"origin_language,omitempty"`
	AncestorWord      string                 `protobuf:"bytes,3,opt,name=ancestor_word,json=ancestorWord,proto3" json:"ancestor_word,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *SymbolAttributes) Reset() {
	*x = SymbolAttributes{}
	mi := &file_language_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SymbolAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolAttributes) ProtoMessage() {}

func (x *SymbolAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolAttributes.ProtoReflect.Descriptor instead.
func (*SymbolAttributes) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{4}
}

func (x *SymbolAttributes) GetLanguageVariation() int64 {
	if x != nil {
		return x.LanguageVariation
	}
	return 0
}

func (x *SymbolAttributes) GetOriginLanguage() string {
	if x != nil {
		return x.OriginLanguage
	}
	return ""
}

func (x *SymbolAttributes) GetAncestorWord() string {
	if x != nil {
		return x.AncestorWord
	}
	return ""
}

type WordAttributes struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Gender         WordGender             `protobuf:"varint,1,opt,name=gender,proto3,enum=lexora.WordGender" json:"gender,omitempty"`
	Number         WordNumber             `protobuf:"varint,2,opt,name=number,proto3,enum=lexora.WordNumber" json:"number,omitempty"`
	Unique         bool                   `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
	Diminutive     bool                   `protobuf:"varint,4,opt,name=diminutive,proto3" json:"diminutive,omitempty"`
	Case           WordCase               `protobuf:"varint,5,opt,name=case,proto3,enum=lexora.WordCase" json:"case,omitempty"`
	Mood           WordMood               `protobuf:"varint,6,opt,name=mood,proto3,enum=lexora.WordMood" json:"mood,omitempty"`
	Tense          WordTense              `protobuf:"varint,7,opt,name=tense,proto3,enum=lexora.WordTense" json:"tense,omitempty"`
	Aspect         WordAspect             `protobuf:"varint,8,opt,name=aspect,proto3,enum=lexora.WordAspect" json:"aspect,omitempty"`
	Person         WordPerson             `protobuf:"varint,9,opt,name=person,proto3,enum=lexora.WordPerson" json:"person,omitempty"`
	DirectObject   bool                   `protobuf:"varint,10,opt,name=direct_object,json=directObject,proto3" json:"direct_object,omitempty"`
	IndirectObject bool                   `protobuf:"varint,11,opt,name=indirect_object,json=indirectObject,proto3" json:"indirect_object,omitempty"`
	Valency        int32                  `protobuf:"varint,12,opt,name=valency,proto3" json:"valency,omitempty"`
	Reflexive      bool                   `protobuf:"varint,13,opt,name=reflexive,proto3" json:"reflexive,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WordAttributes) Reset() {
	*x = WordAttributes{}
	mi := &file_language_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WordAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordAttributes) ProtoMessage() {}

func (x *WordAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_language_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordAttributes.ProtoReflect.Descriptor instead.
func (*WordAttributes) Descriptor() ([]byte, []int) {
	return file_language_types_proto_rawDescGZIP(), []int{5}
}

func (x *WordAttributes) GetGender() WordGender {
	if x != nil {
		return x.Gender
	}
	return WordGender_WORD_GENDER_UNSPECIFIED
}

func (x *WordAttributes) GetNumber() WordNumber {
	if x != nil {
		return x.Number
	}
	return WordNumber_WORD_NUMBER_UNSPECIFIED
}

func (x *WordAttributes) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *WordAttributes) GetDiminutive() bool {
	if x != nil {
		return x.Diminutive
	}
	return false
}

func (x *WordAttributes) GetCase() WordCase {
	if x != nil {
		return x.Case
	}
	return WordCase_WORD_CASE_UNSPECIFIED
}

func (x *WordAttributes) GetMood() WordMood {
	if x != nil {
		return x.Mood
	}
	return WordMood_WORD_MOOD_UNSPECIFIED
}

func (x *WordAttributes) GetTense() WordTense {
	if x != nil {
		return x.Tense
	}
	return WordTense_WORD_TENSE_UNSPECFIED
}

func (x *WordAttributes) GetAspect() WordAspect {
	if x != nil {
		return x.Aspect
	}
	return WordAspect_WORD_ASPECT_UNSPECIFIED
}

func (x *WordAttributes) GetPerson() WordPerson {
	if x != nil {
		return x.Person
	}
	return WordPerson_WORD_PERSON_UNSPECIFIED
}

func (x *WordAttributes) GetDirectObject() bool {
	if x != nil {
		return x.DirectObject
	}
	return false
}

func (x *WordAttributes) GetIndirectObject() bool {
	if x != nil {
		return x.IndirectObject
	}
	return false
}

func (x *WordAttributes) GetValency() int32 {
	if x != nil {
		return x.Valency
	}
	return 0
}

func (x *WordAttributes) GetReflexive() bool {
	if x != nil {
		return x.Reflexive
	}
	return false
}

var File_language_types_proto protoreflect.FileDescriptor

const file_language_types_proto_rawDesc = "" +
	"\n" +
	"\x14language-types.proto\x12\x06lexora\"\xc8\x01\n" +
	"\x04Word\x12\x10\n" +
	"\x03uid\x18\x01 \x01(\tR\x03uid\x12\x12\n" +
	"\x04word\x18\x02 \x01(\tR\x04word\x12!\n" +
	"\fsound_source\x18\x03 \x01(\tR\vsoundSource\x12\x10\n" +
	"\x03ipa\x18\x04 \x01(\tR\x03ipa\x12-\n" +
	"\tword_type\x18\x05 \x01(\x0e2\x10.lexora.WordTypeR\bwordType\x126\n" +
	"\n" +
	"attributes\x18\x06 \x01(\v2\x16.lexora.WordAttributesR\n" +
	"attributes\"\xc5\x01\n" +
	"\x06Symbol\x12\x10\n" +
	"\x03uid\x18\x01 \x01(\tR\x03uid\x12,\n" +
	"\blanguage\x18\x02 \x01(\x0e2\x10.lexora.LanguageR\blanguage\x123\n" +
	"\vsymbol_type\x18\x03 \x01(\x0e2\x12.lexora.SymbolTypeR\n" +
	"symbolType\x12\"\n" +
	"\x05lemma\x18\x04 \x01(\v2\f.lexora.WordR\x05lemma\x12\"\n" +
	"\x05words\x18\x05 \x03(\v2\f.lexora.WordR\x05words\"\x8c\x01\n" +
	"\bReferent\x12\x10\n" +
	"\x03uid\x18\x01 \x01(\tR\x03uid\x12!\n" +
	"\fen_reference\x18\x02 \x01(\tR\venReference\x12!\n" +
	"\fimage_source\x18\x03 \x01(\tR\vimageSource\x12(\n" +
	"\asymbols\x18\x04 \x03(\v2\x0e.lexora.SymbolR\asymbols\"\x14\n" +
	"\x12ReferentAttributes\"\x8f\x01\n" +
	"\x10SymbolAttributes\x12-\n" +
	"\x12language_variation\x18\x01 \x01(\x03R\x11languageVariation\x12'\n" +
	"\x0forigin_language\x18\x02 \x01(\tR\x0eoriginLanguage\x12#\n" +
	"\rancestor_word\x18\x03 \x01(\tR\fancestorWord\"\xf3\x03\n" +
	"\x0eWordAttributes\x12*\n" +
	"\x06gender\x18\x01 \x01(\x0e2\x12.lexora.WordGenderR\x06gender\x12*\n" +
	"\x06number\x18\x02 \x01(\x0e2\x12.lexora.WordNumberR\x06number\x12\x16\n" +
	"\x06unique\x18\x03 \x01(\bR\x06unique\x12\x1e\n" +
	"\n" +
	"diminutive\x18\x04 \x01(\bR\n" +
	"diminutive\x12$\n" +
	"\x04case\x18\x05 \x01(\x0e2\x10.lexora.WordCaseR\x04case\x12$\n" +
	"\x04mood\x18\x06 \x01(\x0e2\x10.lexora.WordMoodR\x04mood\x12'\n" +
	"\x05tense\x18\a \x01(\x0e2\x11.lexora.WordTenseR\x05tense\x12*\n" +
	"\x06aspect\x18\b \x01(\x0e2\x12.lexora.WordAspectR\x06aspect\x12*\n" +
	"\x06person\x18\t \x01(\x0e2\x12.lexora.WordPersonR\x06person\x12#\n" +
	"\rdirect_object\x18\n" +
	" \x01(\bR\fdirectObject\x12'\n" +
	"\x0findirect_object\x18\v \x01(\bR\x0eindirectObject\x12\x18\n" +
	"\avalency\x18\f \x01(\x05R\avalency\x12\x1c\n" +
	"\treflexive\x18\r \x01(\bR\treflexive*\xff\x01\n" +
	"\bWordType\x12\x19\n" +
	"\x15WORD_TYPE_UNSPECIFIED\x10\x00\x12\x12\n" +
	"\x0eWORD_TYPE_NOUN\x10\x01\x12\x12\n" +
	"\x0eWORD_TYPE_VERB\x10\x02\x12\x17\n" +
	"\x13WORD_TYPE_ADJECTIVE\x10\x03\x12\x14\n" +
	"\x10WORD_TYPE_ADVERB\x10\x04\x12\x15\n" +
	"\x11WORD_TYPE_PRONOUN\x10\x05\x12\x19\n" +
	"\x15WORD_TYPE_PREPOSITION\x10\x06\x12\x19\n" +
	"\x15WORD_TYPE_CONJUNCTION\x10\a\x12\x1a\n" +
	"\x16WORD_TYPE_INTERJECTION\x10\b\x12\x18\n" +
	"\x14WORD_TYPE_DETERMINER\x10\t*V\n" +
	"\n" +
	"SymbolType\x12\x16\n" +
	"\x12SYMBOL_TYPE_OBJECT\x10\x00\x12\x16\n" +
	"\x12SYMBOL_TYPE_ACTION\x10\x01\x12\x18\n" +
	"\x14SYMBOL_TYPE_MODIFIER\x10\x02*\x8a\x01\n" +
	"\bLanguage\x12\x18\n" +
	"\x14LANGUAGE_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vLANGUAGE_EN\x10\x01\x12\x0f\n" +
	"\vLANGUAGE_NL\x10\x02\x12\x0f\n" +
	"\vLANGUAGE_ES\x10\x03\x12\x0f\n" +
	"\vLANGUAGE_PT\x10\x04\x12\x0f\n" +
	"\vLANGUAGE_DE\x10\x05\x12\x0f\n" +
	"\vLANGUAGE_FR\x10\x06*j\n" +
	"\tWordTense\x12\x19\n" +
	"\x15WORD_TENSE_UNSPECFIED\x10\x00\x12\x13\n" +
	"\x0fWORD_TENSE_PAST\x10\x01\x12\x16\n" +
	"\x12WORD_TENSE_PRESENT\x10\x02\x12\x15\n" +
	"\x11WORD_TENSE_FUTURE\x10\x03*o\n" +
	"\n" +
	"WordGender\x12\x1b\n" +
	"\x17WORD_GENDER_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10WORD_GENDER_MALE\x10\x01\x12\x16\n" +
	"\x12WORD_GENDER_FEMALE\x10\x02\x12\x16\n" +
	"\x12WORD_GENDER_NEUTER\x10\x03*[\n" +
	"\n" +
	"WordNumber\x12\x1b\n" +
	"\x17WORD_NUMBER_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14WORD_NUMBER_SINGULAR\x10\x01\x12\x16\n" +
	"\x12WORD_NUMBER_PLURAL\x10\x02*%\n" +
	"\bWordMood\x12\x19\n" +
	"\x15WORD_MOOD_UNSPECIFIED\x10\x00*)\n" +
	"\n" +
	"WordAspect\x12\x1b\n" +
	"\x17WORD_ASPECT_UNSPECIFIED\x10\x00*o\n" +
	"\n" +
	"WordPerson\x12\x1b\n" +
	"\x17WORD_PERSON_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11WORD_PERSON_FIRST\x10\x01\x12\x16\n" +
	"\x12WORD_PERSON_SECOND\x10\x02\x12\x15\n" +
	"\x11WORD_PERSON_THIRD\x10\x03*\xeb\x01\n" +
	"\bWordCase\x12\x19\n" +
	"\x15WORD_CASE_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14WORD_CASE_NOMINATIVE\x10\x01\x12\x16\n" +
	"\x12WORD_CASE_GENITIVE\x10\x02\x12\x14\n" +
	"\x10WORD_CASE_DATIVE\x10\x03\x12\x18\n" +
	"\x14WORD_CASE_ACCUSATIVE\x10\x04\x12\x1a\n" +
	"\x16WORD_CASE_INSTRUMENTAL\x10\x05\x12\x16\n" +
	"\x12WORD_CASE_LOCATIVE\x10\x06\x12\x16\n" +
	"\x12WORD_CASE_VOCATIVE\x10\a\x12\x16\n" +
	"\x12WORD_CASE_ABLATIVE\x10\bb\x06proto3"

var (
	file_language_types_proto_rawDescOnce sync.Once
	file_language_types_proto_rawDescData []byte
)

func file_language_types_proto_rawDescGZIP() []byte {
	file_language_types_proto_rawDescOnce.Do(func() {
		file_language_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_language_types_proto_rawDesc), len(file_language_types_proto_rawDesc)))
	})
	return file_language_types_proto_rawDescData
}

var file_language_types_proto_enumTypes = make([]protoimpl.EnumInfo, 10)
var file_language_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_language_types_proto_goTypes = []any{
	(WordType)(0),              // 0: lexora.WordType
	(SymbolType)(0),            // 1: lexora.SymbolType
	(Language)(0),              // 2: lexora.Language
	(WordTense)(0),             // 3: lexora.WordTense
	(WordGender)(0),            // 4: lexora.WordGender
	(WordNumber)(0),            // 5: lexora.WordNumber
	(WordMood)(0),              // 6: lexora.WordMood
	(WordAspect)(0),            // 7: lexora.WordAspect
	(WordPerson)(0),            // 8: lexora.WordPerson
	(WordCase)(0),              // 9: lexora.WordCase
	(*Word)(nil),               // 10: lexora.Word
	(*Symbol)(nil),             // 11: lexora.Symbol
	(*Referent)(nil),           // 12: lexora.Referent
	(*ReferentAttributes)(nil), // 13: lexora.ReferentAttributes
	(*SymbolAttributes)(nil),   // 14: lexora.SymbolAttributes
	(*WordAttributes)(nil),     // 15: lexora.WordAttributes
}
var file_language_types_proto_depIdxs = []int32{
	0,  // 0: lexora.Word.word_type:type_name -> lexora.WordType
	15, // 1: lexora.Word.attributes:type_name -> lexora.WordAttributes
	2,  // 2: lexora.Symbol.language:type_name -> lexora.Language
	1,  // 3: lexora.Symbol.symbol_type:type_name -> lexora.SymbolType
	10, // 4: lexora.Symbol.lemma:type_name -> lexora.Word
	10, // 5: lexora.Symbol.words:type_name -> lexora.Word
	11, // 6: lexora.Referent.symbols:type_name -> lexora.Symbol
	4,  // 7: lexora.WordAttributes.gender:type_name -> lexora.WordGender
	5,  // 8: lexora.WordAttributes.number:type_name -> lexora.WordNumber
	9,  // 9: lexora.WordAttributes.case:type_name -> lexora.WordCase
	6,  // 10: lexora.WordAttributes.mood:type_name -> lexora.WordMood
	3,  // 11: lexora.WordAttributes.tense:type_name -> lexora.WordTense
	7,  // 12: lexora.WordAttributes.aspect:type_name -> lexora.WordAspect
	8,  // 13: lexora.WordAttributes.person:type_name -> lexora.WordPerson
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_language_types_proto_init() }
func file_language_types_proto_init() {
	if File_language_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_language_types_proto_rawDesc), len(file_language_types_proto_rawDesc)),
			NumEnums:      10,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_language_types_proto_goTypes,
		DependencyIndexes: file_language_types_proto_depIdxs,
		EnumInfos:         file_language_types_proto_enumTypes,
		MessageInfos:      file_language_types_proto_msgTypes,
	}.Build()
	File_language_types_proto = out.File
	file_language_types_proto_goTypes = nil
	file_language_types_proto_depIdxs = nil
}
