package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"github.com/pubg/protoc-gen-venus/generator/venus"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func intP(i *int32) *int {
	if i == nil {
		return nil
	}

	ii := int(*i)
	return &ii
}

func concreteBaseComponentOptions(ctx *HierarchicalContext, fd protoreflect.FieldDescriptor, fo *protoptions.FieldOptions) venus.BaseComponentOptions {
	base := venus.BaseComponentOptions{}
	base.PropertyName = ctx.PropertiesString()
	base.Required = !fd.HasOptionalKeyword()

	// Fill Label
	if fo.GetLabel() == "" {
		base.Label = fd.TextName()
	} else {
		base.Label = fo.GetLabel()
	}

	// Fille Placeholder
	if fo.GetPlaceholder() == "" {
		base.Placeholder = fd.TextName()
	} else {
		base.Placeholder = fo.GetPlaceholder()
	}

	// Fill State, State는 inference를 제공하지 않는다.
	if fo.GetState() != protoptions.State_unspecified {
		base.State = fo.GetState().String()
	}

	// Fill Messages
	for _, message := range fo.GetMessages() {
		vMessage := venus.Message{}
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
		base.Grid = &venus.Grid{
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
