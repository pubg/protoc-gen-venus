package generator

import "google.golang.org/protobuf/reflect/protoreflect"

type ScalaKind struct {
	kind             string
	defaultComponent ComponentType
}

var (
	StringKind = ScalaKind{
		kind:             "string",
		defaultComponent: ComponentType_Input,
	}
	NumberKind = ScalaKind{
		kind:             "number",
		defaultComponent: ComponentType_Input,
	}
	BooleanKind = ScalaKind{
		kind:             "boolean",
		defaultComponent: ComponentType_Checkbox,
	}
	EnumKind = ScalaKind{
		kind:             "enum",
		defaultComponent: ComponentType_Select,
	}
)

func ToScalaKind(kind protoreflect.Kind) ScalaKind {
	switch kind {
	case protoreflect.BoolKind:
		return BooleanKind
	case protoreflect.EnumKind:
		return EnumKind
	case protoreflect.Int32Kind:
		fallthrough
	case protoreflect.Sint32Kind:
		fallthrough
	case protoreflect.Uint32Kind:
		fallthrough
	case protoreflect.Int64Kind:
		fallthrough
	case protoreflect.Sint64Kind:
		fallthrough
	case protoreflect.Uint64Kind:
		fallthrough
	case protoreflect.Sfixed32Kind:
		fallthrough
	case protoreflect.Fixed32Kind:
		fallthrough
	case protoreflect.FloatKind:
		fallthrough
	case protoreflect.Sfixed64Kind:
		fallthrough
	case protoreflect.Fixed64Kind:
		fallthrough
	case protoreflect.DoubleKind:
		return NumberKind
	case protoreflect.StringKind:
		fallthrough
	case protoreflect.BytesKind:
		return StringKind
	case protoreflect.MessageKind:
		panic("GroupKind is not supported")
	case protoreflect.GroupKind:
		panic("GroupKind is not supported")
	}
	return StringKind
}
