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
	// text or number
	Type string `json:"type,omitempty"`
	// When type is text, min length of input text
	// When type is number, min value of input number
	Min *int `json:"min,omitempty"`
	// When type is text, max length of input text
	// When type is number, max value of input number
	Max *int `json:"max,omitempty"`
}

func NewInput(base BaseComponentOptions) *Input {
	return &Input{BaseComponent: NewBaseComponent("vn-input", base)}
}

type MultiString struct {
	BaseComponent `json:",inline"`
	// Delimeter
	Delimeter string `json:"delimeter"`
	// Value max length
	Max *int `json:"max"`
	// Hide chips
	NoChips bool `json:"no-chips"`
	// Show copy button
	Copy bool `json:"copy"`
}

func NewMultiString(base BaseComponentOptions) *MultiString {
	return &MultiString{BaseComponent: NewBaseComponent("vn-multi-string", base), NoChips: true, Copy: true}
}

type Select struct {
	BaseComponent `json:",inline"`
	Options       *VlossomOptions `json:",inline,omitempty"`
	// Apply autocomplete
	Autocomplete bool `json:"autocomplete"`
	// Can choose multiple options
	Multiple bool `json:"multiple"`
	// Can choose multiple options
	SelectAll bool `json:"select-all"`
	// Express chips briefly
	CollapseChips bool `json:"collapse-chips"`
	// Can erase chip(s)
	ClosableChips bool `json:"closable-chips"`
	// Change infinite load count
	InfiniteLoad *int `json:"infinite-load,omitempty"`
}

func NewSelect(base BaseComponentOptions, options *VlossomOptions) *Select {
	return &Select{BaseComponent: NewBaseComponent("vn-select", base), Options: options}
}

type Checkbox struct {
	BaseComponent `json:",inline"`
}

func NewCheckbox(base BaseComponentOptions) *Checkbox {
	base.Required = false
	return &Checkbox{BaseComponent: NewBaseComponent("vn-checkbox", base)}
}

type CheckboxSet struct {
	BaseComponent   `json:",inline"`
	*VlossomOptions `json:",inline"`
	Column          bool `json:"column"`
}

func NewCheckboxSet(base BaseComponentOptions, options *VlossomOptions) *CheckboxSet {
	return &CheckboxSet{BaseComponent: NewBaseComponent("vn-checkbox-set", base), VlossomOptions: options}
}

type Switch struct {
	BaseComponent `json:",inline"`
}

func NewSwitch(base BaseComponentOptions) *Switch {
	// Switch not allow required
	base.Required = false
	return &Switch{BaseComponent: NewBaseComponent("vn-switch", base)}
}

type DateRangePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateRangePicker(base BaseComponentOptions) *DateRangePicker {
	return &DateRangePicker{BaseComponent: NewBaseComponent("vn-date-range-picker", base)}
}

type DateTimePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateTimePicker(base BaseComponentOptions) *DateTimePicker {
	return &DateTimePicker{BaseComponent: NewBaseComponent("vn-date-time-picker", base)}
}

type TextArea struct {
	BaseComponent `json:",inline"`
	Min           *int `json:"min,omitempty"`
	Max           *int `json:"max,omitempty"`
}

func NewTextArea(base BaseComponentOptions) *TextArea {
	return &TextArea{BaseComponent: NewBaseComponent("vn-text-area", base)}
}

type JsonEditor struct {
	BaseComponent `json:",inline"`

	// Example: 400px
	Height string `json:"height,omitempty"`
}

func NewJsonEditor(base BaseComponentOptions) *JsonEditor {
	return &JsonEditor{BaseComponent: NewBaseComponent("vn-json-editor", base)}
}

type Radio struct {
	BaseComponent `json:",inline"`
	RadioLabel    string `json:"radio-label"`
	Name          string `json:"name"`
	Option        any    `json:"option"`
}

func NewRadio(base BaseComponentOptions, name string) *Radio {
	return &Radio{BaseComponent: NewBaseComponent("vn-radio", base), Name: name}
}

type RadioSet struct {
	BaseComponent `json:",inline"`
	Options       *VlossomOptions `json:",inline"`
	Column        bool            `json:"column"`
}

func NewRadioSet(base BaseComponentOptions, options *VlossomOptions) *RadioSet {
	return &RadioSet{BaseComponent: NewBaseComponent("vn-radio-set", base), Options: options}
}
