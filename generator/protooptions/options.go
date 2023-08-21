package protooptions

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetFileOptions null을 반환할 수 있지만, null 이라도 리시버 호출 됨
func GetFileOptions(desc protoreflect.Descriptor) *FileOptions {
	options := desc.Options()
	if !proto.HasExtension(options, E_File) {
		return nil
	}
	return proto.GetExtension(options, E_File).(*FileOptions)
}

func GetMessageOptions(desc protoreflect.Descriptor) *MessageOptions {
	options := desc.Options()
	if !proto.HasExtension(options, E_Message) {
		return nil
	}
	return proto.GetExtension(options, E_Message).(*MessageOptions)
}

func GetFieldOptions(desc protoreflect.Descriptor) *FieldOptions {
	options := desc.Options()
	if !proto.HasExtension(options, E_Field) {
		return nil
	}
	return proto.GetExtension(options, E_Field).(*FieldOptions)
}

func GetEnumOptions(desc protoreflect.Descriptor) *EnumOptions {
	options := desc.Options()
	if !proto.HasExtension(options, E_Enum) {
		return nil
	}
	return proto.GetExtension(options, E_Enum).(*EnumOptions)
}

func GetEnumValueOptions(desc protoreflect.Descriptor) *EnumValueOptions {
	options := desc.Options()
	if !proto.HasExtension(options, E_EnumValue) {
		return nil
	}
	return proto.GetExtension(options, E_EnumValue).(*EnumValueOptions)
}
