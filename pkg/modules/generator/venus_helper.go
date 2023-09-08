package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
)

func buildFromInputOptions(option *protoptions.InputOptions, defaultType string, base component.BaseComponentOptions) component.Component {
	comp := component.NewInput(base)
	comp.Type = defaultType
	if option == nil {
		return comp
	}

	if option.GetType() != protoptions.InputOptions_inference {
		comp.Type = option.GetType().String()
	}
	if option.Min != nil {
		comp.Min = intP(option.Min)
	}
	if option.Max != nil {
		comp.Max = intP(option.Max)
	}
	if option.Regex != "" {
		comp.Regex = option.Regex
	}
	return comp
}

func buildFromMultiStringOptions(option *protoptions.MultiStringOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewMultiString(base)
	if option == nil {
		return comp
	}

	comp.Delimeter = option.GetDelimeter()
	if option.Max != nil {
		comp.Max = intP(option.Max)
	}
	comp.NoChips = option.GetNoChips()
	comp.Copy = option.GetCopy()

	return comp
}

func buildFromSelectOptions(option *protoptions.SelectOptions, selectOptions *component.VenusOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewSelect(base, selectOptions)
	if option == nil {
		return comp
	}

	comp.Autocomplete = option.GetAutocomplete()
	comp.Multiple = option.GetMultiple()
	comp.SelectAll = option.GetSelectAll()
	comp.CollapseChips = option.GetCollapseChips()
	comp.ClosableChips = option.GetClosableChips()
	if option.InfiniteLoad != nil {
		comp.InfiniteLoad = intP(option.InfiniteLoad)
	}
	return comp
}

func buildFromCheckboxSetOptions(option *protoptions.CheckboxSetOptions, selectOptions *component.VenusOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewCheckboxSet(base, selectOptions)
	if option == nil {
		return comp
	}

	comp.Column = option.GetColumn()
	return comp
}

func buildFromTextAreaOptions(option *protoptions.TextAreaOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewTextArea(base)
	if option == nil {
		return comp
	}

	if option.Min != nil {
		comp.Min = intP(option.Min)
	}
	if option.Max != nil {
		comp.Max = intP(option.Max)
	}
	if option.Regex != "" {
		comp.Regex = option.Regex
	}
	return comp
}

func buildFromJsonEditorOptions(option *protoptions.JsonEditorOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewJsonEditor(base)
	if option == nil {
		return comp
	}

	if option.Height != nil {
		comp.Height = fmt.Sprintf("%d", option.GetHeight())
	}
	return comp
}

func buildFromRadioOptions(option *protoptions.RadioOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewRadio(base, "radio")
	if option == nil {
		return comp
	}

	comp.RadioLabel = option.GetRadioLabel()
	if option.GetName() != "" {
		comp.Name = option.GetName()
	}
	return comp
}

func buildFromRadioSetOptions(option *protoptions.RadioSetOptions, selectOptions *component.VenusOptions, base component.BaseComponentOptions) component.Component {
	comp := component.NewRadioSet(base, selectOptions)
	if option == nil {
		return comp
	}

	comp.Column = option.GetColumn()
	return comp
}

func buildFromArrayOptions(option *protoptions.ArrayOptions, form component.Form, base component.BaseComponentOptions) component.Component {
	comp := component.NewArray(base)
	comp.Form = form
	if option == nil {
		return comp
	}

	if option.Min != nil {
		comp.Min = intP(option.Min)
	}
	if option.Max != nil {
		comp.Max = intP(option.Max)
	}
	return comp
}

func buildFromMapParameters(key component.Component, values component.Form, base component.BaseComponentOptions) component.Component {
	comp := component.NewMap(base)
	comp.Key = key
	comp.Values = values
	return comp
}

func buildFromTableParameters(key component.Component, values []component.Component, base component.BaseComponentOptions) component.Component {
	comp := component.NewTable(base)
	comp.Key = key
	comp.Values = values
	return comp
}

func convertToVenusOptions(protoOptions *protoptions.VenusOptions) *component.VenusOptions {
	if protoOptions == nil {
		return nil
	}
	vo := component.NewVenusOptions()
	if protoOptions.GetLabeledOptions() != nil {
		var venusOptions []component.LabeledOption
		for _, labeledOption := range protoOptions.LabeledOptions {
			venusLabeledOption := component.LabeledOption{Label: labeledOption.GetLabel()}
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
	} else if protoOptions.GetSimpleOptions() != nil {
		vo.SetSimpleOptions(protoOptions.GetSimpleOptions())
	}

	return vo
}
