package generator

import (
	"fmt"
	"proc-gen-vlossom/generator/protooptions"
	"proc-gen-vlossom/generator/vlossom"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g *VlossomGenerator) concreteBaseComponentOptions(ctx *HierarchicalContext, fd protoreflect.FieldDescriptor, fo *protooptions.FieldOptions) vlossom.BaseComponentOptions {
	base := vlossom.BaseComponentOptions{}
	base.PropertyName = ctx.PropertiesString()
	base.Required = !fd.HasOptionalKeyword()
	if fo.GetLabel() == "" {
		base.Label = fd.TextName()
	} else {
		base.Label = fo.GetLabel()
	}
	if fo.GetPlaceholder() == "" {
		base.Placeholder = fd.TextName()
	} else {
		base.Placeholder = fo.GetPlaceholder()
	}
	if fo.GetRem() != 0 {
		base.Width = fmt.Sprintf("%drem", fo.GetRem())
	} else if fo.GetLg() != 0 {
		base.LG = int(fo.GetLg())
	}
	// State는 inference를 제공하지 않는다.
	if fo.GetState() != protooptions.State_unspecified {
		base.State = fo.GetState().String()
	}
	return base
}

func (g *VlossomGenerator) buildFromScalaField(ctx *HierarchicalContext, field *protogen.Field) (vlossom.Component, error) {
	fd := field.Desc
	fo := protooptions.GetFieldOptions(fd)

	base := g.concreteBaseComponentOptions(ctx, fd, fo)

	scalaKind := ToScalaKind(fd)
	var componentType protooptions.ComponentType
	if fo.GetComponent() != protooptions.ComponentType_Inference {
		componentType = fo.GetComponent()
	} else {
		componentType = scalaKind.defaultComponent
	}

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

func (g *VlossomGenerator) buildFromNumberField(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := protooptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protooptions.ComponentType_Input:
		return buildFromInputOptions(fo.GetInput(), protooptions.InputOptions_number.String(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromNumberField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromStringField(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := protooptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protooptions.ComponentType_Input:
		return buildFromInputOptions(fo.GetInput(), protooptions.InputOptions_text.String(), base), nil
	case protooptions.ComponentType_Select:
		if fo.GetSelect() == nil {
			return nil, fmt.Errorf("failed buildFromStringField, select options is nil")
		}
		return buildFromSelectOptions(fo.GetSelect(), convertToVlossomOptions(fo.GetSelect().GetOptions()), base), nil
	case protooptions.ComponentType_RadioSet:
		if fo.GetRadioSet() == nil {
			return nil, fmt.Errorf("failed buildFromStringField, select options is nil")
		}
		return buildFromRadioSetOptions(fo.GetRadioSet(), convertToVlossomOptions(fo.GetRadioSet().GetOptions()), base), nil
	case protooptions.ComponentType_DateRangePicker:
		return vlossom.NewDateRangePicker(base), nil
	case protooptions.ComponentType_DateTimePicker:
		return vlossom.NewDateTimePicker(base), nil
	case protooptions.ComponentType_MultiString:
		return buildFromMultiStringOptions(fo.GetMultiString(), base), nil
	case protooptions.ComponentType_TextArea:
		return buildFromTextAreaOptions(fo.GetTextArea(), base), nil
	case protooptions.ComponentType_JsonEditor:
		return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromStringField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromEnumField(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fd := field.Desc
	fo := protooptions.GetFieldOptions(fd)
	ed := fd.Enum()
	_ = protooptions.GetEnumOptions(ed)
	values := ed.Values()

	selectOptions := &vlossom.VlossomOptions{}
	var labeledOptions []vlossom.LabeledOption
	for i := 0; i < values.Len(); i++ {
		vd := values.Get(i)
		vo := protooptions.GetEnumValueOptions(vd)

		labeledOption := vlossom.LabeledOption{Label: string(vd.Name())}
		if x, ok := vo.GetValue().(*protooptions.EnumValueOptions_String_); ok {
			labeledOption.Value = x.String_
		} else if x, ok := vo.GetValue().(*protooptions.EnumValueOptions_Integer); ok {
			labeledOption.Value = x.Integer
		} else if x, ok := vo.GetValue().(*protooptions.EnumValueOptions_Float); ok {
			labeledOption.Value = x.Float
		} else {
			labeledOption.Value = vd.Name()
		}
		labeledOptions = append(labeledOptions, labeledOption)
	}
	selectOptions.SetLabeledOptions(labeledOptions)

	switch componentType {
	case protooptions.ComponentType_Select:
		return buildFromSelectOptions(fo.GetSelect(), selectOptions, base), nil
	case protooptions.ComponentType_RadioSet:
		return buildFromRadioSetOptions(fo.GetRadioSet(), selectOptions, base), nil
	}
	return nil, fmt.Errorf("failed buildFromEnumField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromBooleanField(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	switch componentType {
	case protooptions.ComponentType_Checkbox:
		return vlossom.NewCheckbox(base), nil
	case protooptions.ComponentType_Switch:
		return vlossom.NewSwitch(base), nil
	}
	return nil, fmt.Errorf("failed buildFromBooleanField, unknown component type: %s", componentType)
}

type ScalaKind struct {
	kind             string
	defaultComponent protooptions.ComponentType
}

var (
	StringKind = ScalaKind{
		kind:             "string",
		defaultComponent: protooptions.ComponentType_Input,
	}
	NumberKind = ScalaKind{
		kind:             "number",
		defaultComponent: protooptions.ComponentType_Input,
	}
	BooleanKind = ScalaKind{
		kind:             "boolean",
		defaultComponent: protooptions.ComponentType_Checkbox,
	}
	EnumKind = ScalaKind{
		kind:             "enum",
		defaultComponent: protooptions.ComponentType_Select,
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
