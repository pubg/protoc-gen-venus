package vlossom

import "encoding/json"

type Form struct {
	Rows []Row
}

func (f *Form) Generate() ([]byte, error) {
	return json.Marshal(f.Rows)
}

type Row []Component

type Input struct {
	BaseComponent `json:",inline"`
	Max           *int `json:"max,omitempty"`
}

func NewInput(options BaseComponentOptions, max *int) *Input {
	return &Input{BaseComponent: NewBaseComponent("vn-input", options), Max: max}
}

type Select struct {
	BaseComponent `json:",inline"`
	Options       []any `json:"options,omitempty"`
}

func NewSelect(options BaseComponentOptions, selects []any) *Select {
	return &Select{BaseComponent: NewBaseComponent("vn-select", options), Options: selects}
}

type Checkbox struct {
	BaseComponent `json:",inline"`
}

func NewCheckbox(options BaseComponentOptions) *Checkbox {
	return &Checkbox{BaseComponent: NewBaseComponent("vn-checkbox", options)}
}

type Switch struct {
	BaseComponent `json:",inline"`
}

func NewSwitch(options BaseComponentOptions) *Switch {
	// Switch not allow required
	options.Required = false
	return &Switch{BaseComponent: NewBaseComponent("vn-switch", options)}
}

type RadioSet struct {
	BaseComponent `json:",inline"`
	Options       []any `json:"options,omitempty"`
}

func NewRadioSet(options BaseComponentOptions, selects []any) *RadioSet {
	return &RadioSet{BaseComponent: NewBaseComponent("vn-radio-set", options), Options: selects}
}

type DateRangePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateRangePicker(options BaseComponentOptions) *DateRangePicker {
	return &DateRangePicker{BaseComponent: NewBaseComponent("vn-date-range-picker", options)}
}

type DateTimePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateTimePicker(options BaseComponentOptions) *DateTimePicker {
	return &DateTimePicker{BaseComponent: NewBaseComponent("vn-date-time-picker", options)}
}

type JsonEditor struct {
	BaseComponent `json:",inline"`
}

func NewJsonEditor(options BaseComponentOptions) *JsonEditor {
	return &JsonEditor{BaseComponent: NewBaseComponent("vn-json-editor", options)}
}
