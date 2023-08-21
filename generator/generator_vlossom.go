package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-vlossom/generator/protooptions"
	"github.com/pubg/protoc-gen-vlossom/generator/vlossom"
)

func buildFromInputOptions(option *protooptions.InputOptions, defaultType string, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewInput(base)
	component.Type = defaultType
	if option == nil {
		return component
	}

	if option.GetType() != protooptions.InputOptions_inference {
		component.Type = option.GetType().String()
	}
	if option.Min != nil {
		component.Min = intP(option.Min)
	}
	if option.Max != nil {
		component.Max = intP(option.Max)
	}
	return component
}

func buildFromMultiStringOptions(option *protooptions.MultiStringOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewMultiString(base)
	if option == nil {
		return component
	}

	component.Delimeter = option.GetDelimeter()
	if option.Max != nil {
		component.Max = intP(option.Max)
	}
	component.NoChips = option.GetNoChips()
	component.Copy = option.GetCopy()

	return component
}

func buildFromSelectOptions(option *protooptions.SelectOptions, selectOptions *vlossom.VlossomOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewSelect(base, selectOptions)
	if option == nil {
		return component
	}

	component.Autocomplete = option.GetAutocomplete()
	component.Multiple = option.GetMultiple()
	component.SelectAll = option.GetSelectAll()
	component.CollapseChips = option.GetCollapseChips()
	component.ClosableChips = option.GetClosableChips()
	if option.InfiniteLoad != nil {
		component.InfiniteLoad = intP(option.InfiniteLoad)
	}
	return component
}

func buildFromCheckboxSetOptions(option *protooptions.CheckboxSetOptions, selectOptions *vlossom.VlossomOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewCheckboxSet(base, selectOptions)
	if option == nil {
		return component
	}

	component.Column = option.GetColumn()
	return component
}

func buildFromTextAreaOptions(option *protooptions.TextAreaOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewTextArea(base)
	if option == nil {
		return component
	}

	if option.Min != nil {
		component.Min = intP(option.Min)
	}
	if option.Max != nil {
		component.Max = intP(option.Max)
	}
	return component
}

func buildFromJsonEditorOptions(option *protooptions.JsonEditorOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewJsonEditor(base)
	if option == nil {
		return component
	}

	if option.Height != nil {
		component.Height = fmt.Sprintf("%d", option.GetHeight())
	}
	return component
}

func buildFromRadioOptions(option *protooptions.RadioOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewRadio(base, "radio")
	if option == nil {
		return component
	}

	component.RadioLabel = option.GetRadioLabel()
	if option.GetName() != "" {
		component.Name = option.GetName()
	}
	return component
}

func buildFromRadioSetOptions(option *protooptions.RadioSetOptions, selectOptions *vlossom.VlossomOptions, base vlossom.BaseComponentOptions) vlossom.Component {
	component := vlossom.NewRadioSet(base, selectOptions)
	if option == nil {
		return component
	}

	component.Column = option.GetColumn()
	return component
}

func convertToVlossomOptions(protoOptions *protooptions.VlossomOptions) *vlossom.VlossomOptions {
	if protoOptions == nil {
		return nil
	}
	vo := vlossom.NewVlossomOptions()
	if protoOptions.GetLabeledOptions() != nil {
		if protoOptions.GetOptionLabel() == "" {
			vo.OptionLabel = "label"
		} else {
			vo.OptionLabel = protoOptions.GetOptionLabel()
		}
		if protoOptions.GetOptionValue() == "" {
			vo.OptionValue = "value"
		} else {
			vo.OptionValue = protoOptions.GetOptionValue()
		}
		var vlossomOptions []vlossom.LabeledOption
		for _, labeledOption := range protoOptions.LabeledOptions {
			vlossomLabeledOption := vlossom.LabeledOption{Label: labeledOption.GetLabel()}
			if x, ok := labeledOption.GetValue().(*protooptions.VlossomOptions_LabeledOption_String_); ok {
				vlossomLabeledOption.Value = x.String_
			} else if x, ok := labeledOption.GetValue().(*protooptions.VlossomOptions_LabeledOption_Integer); ok {
				vlossomLabeledOption.Value = x.Integer
			} else if x, ok := labeledOption.GetValue().(*protooptions.VlossomOptions_LabeledOption_Float); ok {
				vlossomLabeledOption.Value = x.Float
			}
			vlossomOptions = append(vlossomOptions, vlossomLabeledOption)
		}
		vo.SetLabeledOptions(vlossomOptions)
	}
	if protoOptions.GetSimpleOptions() != nil {
		vo.SetSimpleOptions(protoOptions.GetSimpleOptions())
	}

	return vo
}
