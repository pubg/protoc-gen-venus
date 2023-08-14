package generator

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetFileOptions null을 반환할 수 있지만, null 이라도 리시버 호출 됨
func GetFileOptions(desc protoreflect.Descriptor) *FileOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_FileOptions) {
		return nil
	}
	return proto.GetExtension(options, E_FileOptions).(*FileOptions)
}

func GetMessageOptions(desc protoreflect.Descriptor) *MessageOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_MessageOptions) {
		return nil
	}
	return proto.GetExtension(options, E_MessageOptions).(*MessageOptions)
}

func GetFieldOptions(desc protoreflect.Descriptor) *FieldOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_FieldOptions) {
		return nil
	}
	return proto.GetExtension(options, E_FieldOptions).(*FieldOptions)
}

func GetEnumOptions(desc protoreflect.Descriptor) *EnumOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_EnumOptions) {
		return nil
	}
	return proto.GetExtension(options, E_EnumOptions).(*EnumOptions)
}

func GetEnumValueOptions(desc protoreflect.Descriptor) *EnumValueOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_EnumOptions) {
		return nil
	}
	return proto.GetExtension(options, E_EnumValueOptions).(*EnumValueOptions)
}
