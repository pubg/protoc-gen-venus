package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"github.com/pubg/protoc-gen-venus/generator/venus"
)

func buildFromInputOptions(option *protoptions.InputOptions, defaultType string, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewInput(base)
	component.Type = defaultType
	if option == nil {
		return component
	}

	if option.GetType() != protoptions.InputOptions_inference {
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

func buildFromMultiStringOptions(option *protoptions.MultiStringOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewMultiString(base)
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

func buildFromSelectOptions(option *protoptions.SelectOptions, selectOptions *venus.VenusOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewSelect(base, selectOptions)
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

func buildFromCheckboxSetOptions(option *protoptions.CheckboxSetOptions, selectOptions *venus.VenusOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewCheckboxSet(base, selectOptions)
	if option == nil {
		return component
	}

	component.Column = option.GetColumn()
	return component
}

func buildFromTextAreaOptions(option *protoptions.TextAreaOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewTextArea(base)
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

func buildFromJsonEditorOptions(option *protoptions.JsonEditorOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewJsonEditor(base)
	if option == nil {
		return component
	}

	if option.Height != nil {
		component.Height = fmt.Sprintf("%d", option.GetHeight())
	}
	return component
}

func buildFromRadioOptions(option *protoptions.RadioOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewRadio(base, "radio")
	if option == nil {
		return component
	}

	component.RadioLabel = option.GetRadioLabel()
	if option.GetName() != "" {
		component.Name = option.GetName()
	}
	return component
}

func buildFromRadioSetOptions(option *protoptions.RadioSetOptions, selectOptions *venus.VenusOptions, base venus.BaseComponentOptions) venus.Component {
	component := venus.NewRadioSet(base, selectOptions)
	if option == nil {
		return component
	}

	component.Column = option.GetColumn()
	return component
}

func convertToVenusOptions(protoOptions *protoptions.VenusOptions) *venus.VenusOptions {
	if protoOptions == nil {
		return nil
	}
	vo := venus.NewVenusOptions()
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
		var venusOptions []venus.LabeledOption
		for _, labeledOption := range protoOptions.LabeledOptions {
			venusLabeledOption := venus.LabeledOption{Label: labeledOption.GetLabel()}
			if x, ok := labeledOption.GetValue().(*protoptions.VenusOptions_LabeledOption_String_); ok {
				venusLabeledOption.Value = x.String_
			} else if x, ok := labeledOption.GetValue().(*protoptions.VenusOptions_LabeledOption_Integer); ok {
				venusLabeledOption.Value = x.Integer
			} else if x, ok := labeledOption.GetValue().(*protoptions.VenusOptions_LabeledOption_Float); ok {
				venusLabeledOption.Value = x.Float
			}
			venusOptions = append(venusOptions, venusLabeledOption)
		}
		vo.SetLabeledOptions(venusOptions)
	}
	if protoOptions.GetSimpleOptions() != nil {
		vo.SetSimpleOptions(protoOptions.GetSimpleOptions())
	}

	return vo
}
