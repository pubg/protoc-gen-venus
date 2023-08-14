package generator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

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
	JsonKind = ScalaKind{
		kind:             "map",
		defaultComponent: ComponentType_JsonEditor,
	}
)

func ToScalaKind(fieldDescriptor protoreflect.FieldDescriptor) ScalaKind {
	switch fieldDescriptor.Kind() {
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
		if fieldDescriptor.IsMap() {
			return JsonKind
		}
		panic("MessageKind is not supported")
	case protoreflect.GroupKind:
		panic("GroupKind is not supported")
	}
	panic(fmt.Sprintf("%s is not supported", fieldDescriptor.FullName()))
}

type HierarchicalContext struct {
	exposes    []bool
	properties []string
}

func (c *HierarchicalContext) AppendExpose(expose *bool) {
	if expose != nil {
		c.exposes = append(c.exposes, *expose)
	}
}

func (c *HierarchicalContext) Expose() bool {
	if len(c.exposes) == 0 {
		return false
	}
	for _, expose := range c.exposes {
		if !expose {
			return false
		}
	}
	return true
}

func (c *HierarchicalContext) AppendProperty(property string) {
	c.properties = append(c.properties, property)
}

func (c *HierarchicalContext) AppendPropertyName(name protoreflect.Name) {
	c.properties = append(c.properties, string(name))
}

func (c *HierarchicalContext) PropertiesString() string {
	return strings.Join(c.properties, ".")
}

func NewFromHierarchicalContext(src *HierarchicalContext) *HierarchicalContext {
	dst := &HierarchicalContext{
		exposes:    make([]bool, len(src.exposes)),
		properties: make([]string, len(src.properties)),
	}
	copy(dst.exposes, src.exposes)
	copy(dst.properties, src.properties)
	return dst
}

func NewHierarchicalContext() *HierarchicalContext {
	return &HierarchicalContext{}
}
