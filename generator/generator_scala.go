package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"github.com/pubg/protoc-gen-venus/generator/venus"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g *VenusGenerator) buildFromScalaField(ctx *HierarchicalContext, field *protogen.Field) (venus.Component, error) {
	fd := field.Desc
	fo := protoptions.GetFieldOptions(fd)

	base := concreteBaseComponentOptions(ctx, fd, fo)
	scalaKind := ToScalaKind(fd)
	componentType := getDesiredOrDefaultComponent(fo.GetComponent(), scalaKind.defaultComponent)

	switch scalaKind {
	case NumberKind:
		return g.buildFromNumberField(field, componentType, base)
	case StringKind:
		return g.buildFromStringField(field, componentType, base)
	case EnumKind:
		return g.buildFromEnumField(field, componentType, base)
	case BooleanKind:
		return g.buildFromBooleanField(field, componentType, base)
	}
	return nil, fmt.Errorf("unknown scala kind")
}

func (g *VenusGenerator) buildFromNumberField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protoptions.ComponentType_Input:
		return buildFromInputOptions(fo.GetInput(), protoptions.InputOptions_number.String(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromNumberField, unknown component type: %s", componentType)
}

func (g *VenusGenerator) buildFromStringField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protoptions.ComponentType_Input:
		return buildFromInputOptions(fo.GetInput(), protoptions.InputOptions_text.String(), base), nil
	case protoptions.ComponentType_Select:
		if fo.GetSelect() == nil {
			return nil, fmt.Errorf("failed buildFromStringField, select options is nil")
		}
		return buildFromSelectOptions(fo.GetSelect(), convertToVenusOptions(fo.GetSelect().GetOptions()), base), nil
	case protoptions.ComponentType_RadioSet:
		if fo.GetRadioSet() == nil {
			return nil, fmt.Errorf("failed buildFromStringField, select options is nil")
		}
		return buildFromRadioSetOptions(fo.GetRadioSet(), convertToVenusOptions(fo.GetRadioSet().GetOptions()), base), nil
	case protoptions.ComponentType_DateRangePicker:
		return venus.NewDateRangePicker(base), nil
	case protoptions.ComponentType_DateTimePicker:
		return venus.NewDateTimePicker(base), nil
	case protoptions.ComponentType_MultiString:
		return buildFromMultiStringOptions(fo.GetMultiString(), base), nil
	case protoptions.ComponentType_TextArea:
		return buildFromTextAreaOptions(fo.GetTextArea(), base), nil
	case protoptions.ComponentType_JsonEditor:
		return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromStringField, unknown component type: %s", componentType)
}

func (g *VenusGenerator) buildFromEnumField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fd := field.Desc
	fo := protoptions.GetFieldOptions(fd)
	ed := fd.Enum()
	_ = protoptions.GetEnumOptions(ed)
	values := ed.Values()

	selectOptions := &venus.VenusOptions{}
	var labeledOptions []venus.LabeledOption
	for i := 0; i < values.Len(); i++ {
		vd := values.Get(i)
		vo := protoptions.GetEnumValueOptions(vd)

		labeledOption := venus.LabeledOption{Label: string(vd.Name())}
		if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_String_); ok {
			labeledOption.Value = x.String_
		} else if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_Integer); ok {
			labeledOption.Value = x.Integer
		} else if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_Float); ok {
			labeledOption.Value = x.Float
		} else {
			labeledOption.Value = vd.Name()
		}
		labeledOptions = append(labeledOptions, labeledOption)
	}
	selectOptions.SetLabeledOptions(labeledOptions)

	switch componentType {
	case protoptions.ComponentType_Select:
		return buildFromSelectOptions(fo.GetSelect(), selectOptions, base), nil
	case protoptions.ComponentType_RadioSet:
		return buildFromRadioSetOptions(fo.GetRadioSet(), selectOptions, base), nil
	}
	return nil, fmt.Errorf("failed buildFromEnumField, unknown component type: %s", componentType)
}

func (g *VenusGenerator) buildFromBooleanField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	switch componentType {
	case protoptions.ComponentType_Checkbox:
		return venus.NewCheckbox(base), nil
	case protoptions.ComponentType_Switch:
		return venus.NewSwitch(base), nil
	}
	return nil, fmt.Errorf("failed buildFromBooleanField, unknown component type: %s", componentType)
}

type ScalaKind struct {
	kind             string
	defaultComponent protoptions.ComponentType
}

var (
	StringKind = ScalaKind{
		kind:             "string",
		defaultComponent: protoptions.ComponentType_Input,
	}
	NumberKind = ScalaKind{
		kind:             "number",
		defaultComponent: protoptions.ComponentType_Input,
	}
	BooleanKind = ScalaKind{
		kind:             "boolean",
		defaultComponent: protoptions.ComponentType_Checkbox,
	}
	EnumKind = ScalaKind{
		kind:             "enum",
		defaultComponent: protoptions.ComponentType_Select,
	}
)

func ToScalaKind(fieldDescriptor protoreflect.FieldDescriptor) ScalaKind {
	switch fieldDescriptor.Kind() {
	case protoreflect.BoolKind:
		return BooleanKind
	case protoreflect.EnumKind:
		return EnumKind
	case protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.FloatKind,
		protoreflect.Sfixed64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.DoubleKind:
		return NumberKind
	case protoreflect.StringKind,
		protoreflect.BytesKind:
		return StringKind
	case protoreflect.MessageKind:
		panic("MessageKind is not supported")
	case protoreflect.GroupKind:
		panic("GroupKind is not supported")
	}
	panic(fmt.Sprintf("%s is not supported", fieldDescriptor.FullName()))
}
