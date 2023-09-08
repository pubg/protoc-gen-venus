package generator

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
)

func intP(i *int32) *int {
	if i == nil {
		return nil
	}

	ii := int(*i)
	return &ii
}

func ConcreteBaseComponentOptionsFromOneOf(oneOf pgs.OneOf, oo *protoptions.OneOfOptions) component.BaseComponentOptions {
	base := component.BaseComponentOptions{}
	base.Required = true

	// Fill Label
	if oo.GetLabel() == "" {
		base.Label = oneOf.Name().String()
	} else {
		base.Label = oo.GetLabel()
	}
	return base
}

func ConcreteBaseComponentOptions(field pgs.Field, fo *protoptions.FieldOptions) component.BaseComponentOptions {
	base := component.BaseComponentOptions{}
	base.Required = !field.HasOptionalKeyword()

	// Fill Label
	if fo.GetLabel() == "" {
		base.Label = field.Name().String()
	} else {
		base.Label = fo.GetLabel()
	}

	// Fille Placeholder
	if fo.GetPlaceholder() == "" {
		base.Placeholder = field.Name().String()
	} else {
		base.Placeholder = fo.GetPlaceholder()
	}

	// Fill State, State는 inference를 제공하지 않는다.
	if fo.GetState() != protoptions.State_unspecified {
		base.State = fo.GetState().String()
	}

	// Fill Messages
	for _, message := range fo.GetMessages() {
		vMessage := component.Message{}
		vMessage.Text = message.GetText()
		if message.GetState() != protoptions.State_unspecified {
			vMessage.State = message.GetState().String()
		}
		base.Messages = append(base.Messages, vMessage)
	}

	// Fill Grid or Width(rem)
	switch size := fo.GetSize().(type) {
	case *protoptions.FieldOptions_Grid:
		grid := size.Grid
		base.Grid = &component.Grid{
			Sm:       intP(grid.Md),
			Md:       intP(grid.Md),
			Lg:       intP(grid.Lg),
			SmOffset: intP(grid.MdOffset),
			MdOffset: intP(grid.MdOffset),
			LgOffset: intP(grid.LgOffset),
			Order:    intP(grid.Order),
		}
	case *protoptions.FieldOptions_Rem:
		base.Width = fmt.Sprintf("%drem", size.Rem)
	}

	// Fill DefaultValue
	switch defaultValue := fo.GetDefaultValue().(type) {
	case *protoptions.FieldOptions_DefaultString:
		base.DefaultValue = defaultValue.DefaultString
	case *protoptions.FieldOptions_DefaultInteger:
		base.DefaultValue = defaultValue.DefaultInteger
	case *protoptions.FieldOptions_DefaultFloat:
		base.DefaultValue = defaultValue.DefaultFloat
	}
	return base
}

func getDesiredOrDefaultComponent(desired protoptions.ComponentType, _default protoptions.ComponentType) protoptions.ComponentType {
	if desired == protoptions.ComponentType_Inference {
		return _default
	}
	return desired
}
