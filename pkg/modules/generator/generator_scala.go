package generator

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
)

func BuildFromScalaField(field pgs.Field) (component.Component, error) {
	fo := protoptions.GetFieldOptions(field)

	base := ConcreteBaseComponentOptions(field, fo)
	scalaKind := *ToScalaKind(field)
	componentType := getDesiredOrDefaultComponent(fo.GetComponent(), scalaKind.defaultComponent)

	switch scalaKind {
	case NumberKind:
		return buildFromNumberField(field, componentType, base, fo)
	case StringKind:
		return buildFromStringField(field, componentType, base, fo)
	case EnumKind:
		return buildFromEnumField(field, componentType, base, fo)
	case BooleanKind:
		return buildFromBooleanField(field, componentType, base, fo)
	}
	return nil, fmt.Errorf("unknown scala kind")
}

func buildFromNumberField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	switch componentType {
	case protoptions.ComponentType_Input:
		return buildFromInputOptions(fo.GetInput(), protoptions.InputOptions_number.String(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromNumberField, unknown component type: %s", componentType)
}

func buildFromStringField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
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
		return component.NewDateRangePicker(base), nil
	case protoptions.ComponentType_DateTimePicker:
		return component.NewDateTimePicker(base), nil
	case protoptions.ComponentType_MultiString:
		return buildFromMultiStringOptions(fo.GetMultiString(), base), nil
	case protoptions.ComponentType_TextArea:
		return buildFromTextAreaOptions(fo.GetTextArea(), base), nil
	case protoptions.ComponentType_JsonEditor:
		return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromStringField, unknown component type: %s", componentType)
}

func buildFromEnumField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	ft := field.Type()
	ed := ft.Enum()
	values := ed.Values()

	selectOptions := &component.VenusOptions{}
	var labeledOptions []component.LabeledOption
	for i := 0; i < len(values); i++ {
		vd := values[i]
		vo := protoptions.GetEnumValueOptions(vd)

		labeledOption := component.LabeledOption{Label: string(vd.Name())}
		if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_String_); ok {
			labeledOption.Value = x.String_
		} else if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_Integer); ok {
			labeledOption.Value = x.Integer
		} else if x, ok := vo.GetValue().(*protoptions.EnumValueOptions_Float); ok {
			labeledOption.Value = x.Float
		} else {
			labeledOption.Value = vd.Name().String()
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

func buildFromBooleanField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	switch componentType {
	case protoptions.ComponentType_Checkbox:
		return component.NewCheckbox(base), nil
	case protoptions.ComponentType_Switch:
		return component.NewSwitch(base), nil
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

func IsScalaKind(field pgs.Field) bool {
	return ToScalaKind(field) != nil
}

func ToScalaKind(field pgs.Field) *ScalaKind {
	protoType := field.Type().ProtoType()
	if protoType.IsNumeric() {
		return &NumberKind
	}
	if protoType == pgs.BoolT {
		return &BooleanKind
	}
	if protoType == pgs.EnumT {
		return &EnumKind
	}
	if protoType == pgs.StringT || protoType == pgs.BytesT {
		return &StringKind
	}
	// may return null if type is map or group
	return nil
}
