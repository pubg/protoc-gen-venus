// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: options.proto

package generator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComponentType int32

const (
	ComponentType_Default         ComponentType = 0
	ComponentType_Input           ComponentType = 1
	ComponentType_Select          ComponentType = 2
	ComponentType_Checkbox        ComponentType = 3
	ComponentType_Switch          ComponentType = 4
	ComponentType_RadioSet        ComponentType = 5
	ComponentType_DateRangePicker ComponentType = 6
	ComponentType_DateTimePicker  ComponentType = 7
	ComponentType_JsonEditor      ComponentType = 8
)

// Enum value maps for ComponentType.
var (
	ComponentType_name = map[int32]string{
		0: "Default",
		1: "Input",
		2: "Select",
		3: "Checkbox",
		4: "Switch",
		5: "RadioSet",
		6: "DateRangePicker",
		7: "DateTimePicker",
		8: "JsonEditor",
	}
	ComponentType_value = map[string]int32{
		"Default":         0,
		"Input":           1,
		"Select":          2,
		"Checkbox":        3,
		"Switch":          4,
		"RadioSet":        5,
		"DateRangePicker": 6,
		"DateTimePicker":  7,
		"JsonEditor":      8,
	}
)

func (x ComponentType) Enum() *ComponentType {
	p := new(ComponentType)
	*p = x
	return p
}

func (x ComponentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComponentType) Descriptor() protoreflect.EnumDescriptor {
	return file_options_proto_enumTypes[0].Descriptor()
}

func (ComponentType) Type() protoreflect.EnumType {
	return &file_options_proto_enumTypes[0]
}

func (x ComponentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComponentType.Descriptor instead.
func (ComponentType) EnumDescriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

// Not extendable, just define options
type PluginOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExposeAll        *bool   `protobuf:"varint,1,opt,name=exposeAll,proto3,oneof" json:"exposeAll,omitempty"`
	OutputFileSuffix *string `protobuf:"bytes,2,opt,name=outputFileSuffix,proto3,oneof" json:"outputFileSuffix,omitempty"`
}

func (x *PluginOptions) Reset() {
	*x = PluginOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginOptions) ProtoMessage() {}

func (x *PluginOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginOptions.ProtoReflect.Descriptor instead.
func (*PluginOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *PluginOptions) GetExposeAll() bool {
	if x != nil && x.ExposeAll != nil {
		return *x.ExposeAll
	}
	return false
}

func (x *PluginOptions) GetOutputFileSuffix() string {
	if x != nil && x.OutputFileSuffix != nil {
		return *x.OutputFileSuffix
	}
	return ""
}

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expose            *bool  `protobuf:"varint,1,opt,name=expose,proto3,oneof" json:"expose,omitempty"`
	EntrypointMessage string `protobuf:"bytes,2,opt,name=entrypointMessage,proto3" json:"entrypointMessage,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *FileOptions) GetExpose() bool {
	if x != nil && x.Expose != nil {
		return *x.Expose
	}
	return false
}

func (x *FileOptions) GetEntrypointMessage() string {
	if x != nil {
		return x.EntrypointMessage
	}
	return ""
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expose   *bool  `protobuf:"varint,1,opt,name=expose,proto3,oneof" json:"expose,omitempty"`
	Property string `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{2}
}

func (x *MessageOptions) GetExpose() bool {
	if x != nil && x.Expose != nil {
		return *x.Expose
	}
	return false
}

func (x *MessageOptions) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expose      *bool                       `protobuf:"varint,1,opt,name=expose,proto3,oneof" json:"expose,omitempty"`
	Property    string                      `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
	Lg          int32                       `protobuf:"varint,3,opt,name=lg,proto3" json:"lg,omitempty"`
	Label       string                      `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Placeholder string                      `protobuf:"bytes,5,opt,name=placeholder,proto3" json:"placeholder,omitempty"`
	Component   ComponentType               `protobuf:"varint,6,opt,name=component,proto3,enum=protoc.gen.vlossom.ComponentType" json:"component,omitempty"`
	String_     *FieldOptions_StringOptions `protobuf:"bytes,10,opt,name=string,proto3" json:"string,omitempty"`
	Number      *FieldOptions_NumberOptions `protobuf:"bytes,11,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{3}
}

func (x *FieldOptions) GetExpose() bool {
	if x != nil && x.Expose != nil {
		return *x.Expose
	}
	return false
}

func (x *FieldOptions) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *FieldOptions) GetLg() int32 {
	if x != nil {
		return x.Lg
	}
	return 0
}

func (x *FieldOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *FieldOptions) GetPlaceholder() string {
	if x != nil {
		return x.Placeholder
	}
	return ""
}

func (x *FieldOptions) GetComponent() ComponentType {
	if x != nil {
		return x.Component
	}
	return ComponentType_Default
}

func (x *FieldOptions) GetString_() *FieldOptions_StringOptions {
	if x != nil {
		return x.String_
	}
	return nil
}

func (x *FieldOptions) GetNumber() *FieldOptions_NumberOptions {
	if x != nil {
		return x.Number
	}
	return nil
}

type EnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnumOptions) Reset() {
	*x = EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions) ProtoMessage() {}

func (x *EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumOptions.ProtoReflect.Descriptor instead.
func (*EnumOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{4}
}

type EnumValueOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*EnumValueOptions_String_
	//	*EnumValueOptions_Integer
	//	*EnumValueOptions_Float
	Value isEnumValueOptions_Value `protobuf_oneof:"Value"`
}

func (x *EnumValueOptions) Reset() {
	*x = EnumValueOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumValueOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumValueOptions) ProtoMessage() {}

func (x *EnumValueOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumValueOptions.ProtoReflect.Descriptor instead.
func (*EnumValueOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{5}
}

func (m *EnumValueOptions) GetValue() isEnumValueOptions_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *EnumValueOptions) GetString_() string {
	if x, ok := x.GetValue().(*EnumValueOptions_String_); ok {
		return x.String_
	}
	return ""
}

func (x *EnumValueOptions) GetInteger() int64 {
	if x, ok := x.GetValue().(*EnumValueOptions_Integer); ok {
		return x.Integer
	}
	return 0
}

func (x *EnumValueOptions) GetFloat() float64 {
	if x, ok := x.GetValue().(*EnumValueOptions_Float); ok {
		return x.Float
	}
	return 0
}

type isEnumValueOptions_Value interface {
	isEnumValueOptions_Value()
}

type EnumValueOptions_String_ struct {
	String_ string `protobuf:"bytes,1,opt,name=string,proto3,oneof"`
}

type EnumValueOptions_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,proto3,oneof"`
}

type EnumValueOptions_Float struct {
	Float float64 `protobuf:"fixed64,3,opt,name=float,proto3,oneof"`
}

func (*EnumValueOptions_String_) isEnumValueOptions_Value() {}

func (*EnumValueOptions_Integer) isEnumValueOptions_Value() {}

func (*EnumValueOptions_Float) isEnumValueOptions_Value() {}

type FieldOptions_StringOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min   *int32  `protobuf:"varint,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max   *int32  `protobuf:"varint,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
	Regex *string `protobuf:"bytes,3,opt,name=regex,proto3,oneof" json:"regex,omitempty"`
}

func (x *FieldOptions_StringOptions) Reset() {
	*x = FieldOptions_StringOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions_StringOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions_StringOptions) ProtoMessage() {}

func (x *FieldOptions_StringOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions_StringOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions_StringOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FieldOptions_StringOptions) GetMin() int32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *FieldOptions_StringOptions) GetMax() int32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *FieldOptions_StringOptions) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

type FieldOptions_NumberOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slide bool   `protobuf:"varint,1,opt,name=slide,proto3" json:"slide,omitempty"`
	Min   *int32 `protobuf:"varint,2,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max   *int32 `protobuf:"varint,3,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *FieldOptions_NumberOptions) Reset() {
	*x = FieldOptions_NumberOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions_NumberOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions_NumberOptions) ProtoMessage() {}

func (x *FieldOptions_NumberOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions_NumberOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions_NumberOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{3, 1}
}

func (x *FieldOptions_NumberOptions) GetSlide() bool {
	if x != nil {
		return x.Slide
	}
	return false
}

func (x *FieldOptions_NumberOptions) GetMin() int32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *FieldOptions_NumberOptions) GetMax() int32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         11241,
		Name:          "protoc.gen.vlossom.file_options",
		Tag:           "bytes,11241,opt,name=file_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         11242,
		Name:          "protoc.gen.vlossom.field_options",
		Tag:           "bytes,11242,opt,name=field_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         11243,
		Name:          "protoc.gen.vlossom.message_options",
		Tag:           "bytes,11243,opt,name=message_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*EnumOptions)(nil),
		Field:         11244,
		Name:          "protoc.gen.vlossom.enum_options",
		Tag:           "bytes,11244,opt,name=enum_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*EnumValueOptions)(nil),
		Field:         11245,
		Name:          "protoc.gen.vlossom.enum_value_options",
		Tag:           "bytes,11245,opt,name=enum_value_options",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional protoc.gen.vlossom.FileOptions file_options = 11241;
	E_FileOptions = &file_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional protoc.gen.vlossom.FieldOptions field_options = 11242;
	E_FieldOptions = &file_options_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional protoc.gen.vlossom.MessageOptions message_options = 11243;
	E_MessageOptions = &file_options_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional protoc.gen.vlossom.EnumOptions enum_options = 11244;
	E_EnumOptions = &file_options_proto_extTypes[3]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional protoc.gen.vlossom.EnumValueOptions enum_value_options = 11245;
	E_EnumValueOptions = &file_options_proto_extTypes[4]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x73,
	0x65, 0x41, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x63,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x22, 0xc4, 0x04, 0x0a, 0x0c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x6c, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f,
	0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x72, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x1a, 0x63, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x6d,
	0x61, 0x78, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65,
	0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x69, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a,
	0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x42, 0x07, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x94, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x53, 0x65, 0x74, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x06, 0x12, 0x12, 0x0a,
	0x0e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x10,
	0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x10,
	0x08, 0x3a, 0x61, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe9, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x65, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xea, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x6d, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xeb, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x61, 0x0a, 0x0c, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f,
	0x73, 0x73, 0x6f, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x76, 0x0a,
	0x12, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xed, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x6c, 0x6f, 0x73, 0x73,
	0x6f, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x10, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x76, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_options_proto_goTypes = []interface{}{
	(ComponentType)(0),                    // 0: protoc.gen.vlossom.ComponentType
	(*PluginOptions)(nil),                 // 1: protoc.gen.vlossom.PluginOptions
	(*FileOptions)(nil),                   // 2: protoc.gen.vlossom.FileOptions
	(*MessageOptions)(nil),                // 3: protoc.gen.vlossom.MessageOptions
	(*FieldOptions)(nil),                  // 4: protoc.gen.vlossom.FieldOptions
	(*EnumOptions)(nil),                   // 5: protoc.gen.vlossom.EnumOptions
	(*EnumValueOptions)(nil),              // 6: protoc.gen.vlossom.EnumValueOptions
	(*FieldOptions_StringOptions)(nil),    // 7: protoc.gen.vlossom.FieldOptions.StringOptions
	(*FieldOptions_NumberOptions)(nil),    // 8: protoc.gen.vlossom.FieldOptions.NumberOptions
	(*descriptorpb.FileOptions)(nil),      // 9: google.protobuf.FileOptions
	(*descriptorpb.FieldOptions)(nil),     // 10: google.protobuf.FieldOptions
	(*descriptorpb.MessageOptions)(nil),   // 11: google.protobuf.MessageOptions
	(*descriptorpb.EnumOptions)(nil),      // 12: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 13: google.protobuf.EnumValueOptions
}
var file_options_proto_depIdxs = []int32{
	0,  // 0: protoc.gen.vlossom.FieldOptions.component:type_name -> protoc.gen.vlossom.ComponentType
	7,  // 1: protoc.gen.vlossom.FieldOptions.string:type_name -> protoc.gen.vlossom.FieldOptions.StringOptions
	8,  // 2: protoc.gen.vlossom.FieldOptions.number:type_name -> protoc.gen.vlossom.FieldOptions.NumberOptions
	9,  // 3: protoc.gen.vlossom.file_options:extendee -> google.protobuf.FileOptions
	10, // 4: protoc.gen.vlossom.field_options:extendee -> google.protobuf.FieldOptions
	11, // 5: protoc.gen.vlossom.message_options:extendee -> google.protobuf.MessageOptions
	12, // 6: protoc.gen.vlossom.enum_options:extendee -> google.protobuf.EnumOptions
	13, // 7: protoc.gen.vlossom.enum_value_options:extendee -> google.protobuf.EnumValueOptions
	2,  // 8: protoc.gen.vlossom.file_options:type_name -> protoc.gen.vlossom.FileOptions
	4,  // 9: protoc.gen.vlossom.field_options:type_name -> protoc.gen.vlossom.FieldOptions
	3,  // 10: protoc.gen.vlossom.message_options:type_name -> protoc.gen.vlossom.MessageOptions
	5,  // 11: protoc.gen.vlossom.enum_options:type_name -> protoc.gen.vlossom.EnumOptions
	6,  // 12: protoc.gen.vlossom.enum_value_options:type_name -> protoc.gen.vlossom.EnumValueOptions
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	8,  // [8:13] is the sub-list for extension type_name
	3,  // [3:8] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginOptions); i {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
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
		file_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
		file_options_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumOptions); i {
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
		file_options_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumValueOptions); i {
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
		file_options_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions_StringOptions); i {
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
		file_options_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions_NumberOptions); i {
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
	file_options_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EnumValueOptions_String_)(nil),
		(*EnumValueOptions_Integer)(nil),
		(*EnumValueOptions_Float)(nil),
	}
	file_options_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		EnumInfos:         file_options_proto_enumTypes,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
