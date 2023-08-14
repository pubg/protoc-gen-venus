package generator

import (
	"fmt"
	"proc-gen-vlossom/generator/vlossom"

	"google.golang.org/protobuf/compiler/protogen"
)

func (g *VlossomGenerator) buildFromScalaField(ctx *HierarchicalContext, field *protogen.Field) (vlossom.Component, error) {
	fd := field.Desc
	fo := GetFieldOptions(fd)

	componentOptions := vlossom.BaseComponentOptions{}
	componentOptions.PropertyName = ctx.PropertiesString()
	componentOptions.Required = !fd.HasOptionalKeyword()
	if fo.GetLabel() == "" {
		componentOptions.Label = fd.TextName()
	} else {
		componentOptions.Label = fo.GetLabel()
	}
	if fo.GetPlaceholder() == "" {
		componentOptions.Placeholder = fd.TextName()
	} else {
		componentOptions.Placeholder = fo.GetPlaceholder()
	}
	if fo.GetLg() == 0 {
		componentOptions.LG = 5
	} else {
		componentOptions.LG = int(fo.GetLg())
	}
	scalaKind := ToScalaKind(fd)
	var componentType ComponentType
	if fo.GetComponent() == ComponentType_Default {
		componentType = scalaKind.defaultComponent
	} else {
		componentType = fo.GetComponent()
	}

	switch scalaKind {
	case NumberKind:
		return g.buildFromNumberField(field, componentType, componentOptions)
	case StringKind:
		return g.buildFromStringField(field, componentType, componentOptions)
	case EnumKind:
		return g.buildFromEnumField(field, componentType, componentOptions)
	case BooleanKind:
		return g.buildFromBooleanField(field, componentType, componentOptions)
	case JsonKind:
		return g.buildFromMapField(field, componentType, componentOptions)
	}
	return nil, fmt.Errorf("unknown scala kind")
}

func (g *VlossomGenerator) buildFromNumberField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := GetFieldOptions(field.Desc)
	switch componentType {
	case ComponentType_Input:
		var max *int
		if fo.GetString_().Max != nil {
			i := int(*fo.GetString_().Max)
			max = &i
		}
		return vlossom.NewInput(componentOptions, max), nil
	}
	return nil, fmt.Errorf("failed buildFromNumberField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromStringField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	switch componentType {
	case ComponentType_Input:
		return vlossom.NewInput(componentOptions, nil), nil
	case ComponentType_Select:
		return vlossom.NewSelect(componentOptions, nil), nil
	case ComponentType_RadioSet:
		return vlossom.NewRadioSet(componentOptions, nil), nil
	case ComponentType_DateRangePicker:
		return vlossom.NewDateRangePicker(componentOptions), nil
	case ComponentType_DateTimePicker:
		return vlossom.NewDateTimePicker(componentOptions), nil
	}
	return nil, fmt.Errorf("failed buildFromStringField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromEnumField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fd := field.Desc
	_ = GetFieldOptions(fd)
	ed := fd.Enum()
	_ = GetEnumOptions(ed)
	values := ed.Values()

	var options []any
	for i := 0; i < values.Len(); i++ {
		vd := values.Get(i)
		vo := GetEnumValueOptions(vd)

		var option any
		if x, ok := vo.GetValue().(*EnumValueOptions_String_); ok {
			option = x.String_
		} else if x, ok := vo.GetValue().(*EnumValueOptions_Integer); ok {
			option = x.Integer
		} else if x, ok := vo.GetValue().(*EnumValueOptions_Float); ok {
			option = x.Float
		} else {
			option = vd.Name()
		}
		options = append(options, option)
	}

	switch componentType {
	case ComponentType_Select:
		return vlossom.NewSelect(componentOptions, options), nil
	case ComponentType_RadioSet:
		return vlossom.NewRadioSet(componentOptions, options), nil
	}
	return nil, fmt.Errorf("failed buildFromEnumField, unknown component type: %s", componentType)
}

// TODO: Toggle 추가해야 함
func (g *VlossomGenerator) buildFromBooleanField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	switch componentType {
	case ComponentType_Checkbox:
		return vlossom.NewCheckbox(componentOptions), nil
	}
	return nil, fmt.Errorf("failed buildFromBooleanField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromMapField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	return vlossom.NewJsonEditor(componentOptions), nil
}
