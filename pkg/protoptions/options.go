package protoptions

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func GetFileOptions(file pgs.File) *FileOptions {
	options := &FileOptions{}
	if ok, _ := file.Extension(E_File, options); ok {
		return options
	}
	return nil
}

func GetMessageOptions(message pgs.Message) *MessageOptions {
	options := &MessageOptions{}
	if ok, _ := message.Extension(E_Message, options); ok {
		return options
	}
	return nil
}

func GetFieldOptions(field pgs.Field) *FieldOptions {
	options := &FieldOptions{}
	if ok, _ := field.Extension(E_Field, options); ok {
		return options
	}
	return nil
}

func GetEnumValueOptions(enumValue pgs.EnumValue) *EnumValueOptions {
	options := &EnumValueOptions{}
	if ok, _ := enumValue.Extension(E_EnumValue, options); ok {
		return options
	}
	return nil
}

func GetOneOfOptions(oneOf pgs.OneOf) *OneOfOptions {
	options := &OneOfOptions{}
	if ok, _ := oneOf.Extension(E_OneOf, options); ok {
		return options
	}
	return nil

}
