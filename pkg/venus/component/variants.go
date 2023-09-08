package component

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

	Regex string `json:"regex,omitempty"`
}

func NewInput(base BaseComponentOptions) *Input {
	return &Input{BaseComponent: NewBaseComponent("vn-input", base)}
}

func (c *Input) DeepCopy() Component {
	dst := &Input{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Min = intPP(c.Min)
	dst.Max = intPP(c.Max)
	return dst
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

func (c *MultiString) DeepCopy() Component {
	dst := &MultiString{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Max = intPP(c.Max)
	return dst
}

type Select struct {
	BaseComponent `json:",inline"`
	*VenusOptions `json:",inline"`
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

func NewSelect(base BaseComponentOptions, options *VenusOptions) *Select {
	return &Select{BaseComponent: NewBaseComponent("vn-select", base), VenusOptions: options}
}

func (c *Select) DeepCopy() Component {
	dst := &Select{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.InfiniteLoad = intPP(c.InfiniteLoad)
	return dst
}

type Checkbox struct {
	BaseComponent `json:",inline"`
}

func NewCheckbox(base BaseComponentOptions) *Checkbox {
	base.Required = false
	return &Checkbox{BaseComponent: NewBaseComponent("vn-checkbox", base)}
}

func (c *Checkbox) DeepCopy() Component {
	dst := &Checkbox{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
}

type CheckboxSet struct {
	BaseComponent `json:",inline"`
	*VenusOptions `json:",inline"`
	Column        bool `json:"column"`
}

func NewCheckboxSet(base BaseComponentOptions, options *VenusOptions) *CheckboxSet {
	return &CheckboxSet{BaseComponent: NewBaseComponent("vn-checkbox-set", base), VenusOptions: options}
}

func (c *CheckboxSet) DeepCopy() Component {
	dst := &CheckboxSet{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.VenusOptions = c.VenusOptions.DeepCopy()
	return dst
}

type Switch struct {
	BaseComponent `json:",inline"`
}

func NewSwitch(base BaseComponentOptions) *Switch {
	// Switch not allow required
	base.Required = false
	return &Switch{BaseComponent: NewBaseComponent("vn-switch", base)}
}

func (c *Switch) DeepCopy() Component {
	dst := &Switch{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
}

type DateRangePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateRangePicker(base BaseComponentOptions) *DateRangePicker {
	return &DateRangePicker{BaseComponent: NewBaseComponent("vn-date-range-picker", base)}
}

func (c *DateRangePicker) DeepCopy() Component {
	dst := &DateRangePicker{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
}

type DateTimePicker struct {
	BaseComponent `json:",inline"`
}

func NewDateTimePicker(base BaseComponentOptions) *DateTimePicker {
	return &DateTimePicker{BaseComponent: NewBaseComponent("vn-date-time-picker", base)}
}

func (c *DateTimePicker) DeepCopy() Component {
	dst := &DateTimePicker{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
}

type TextArea struct {
	BaseComponent `json:",inline"`
	Min           *int   `json:"min,omitempty"`
	Max           *int   `json:"max,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

func NewTextArea(base BaseComponentOptions) *TextArea {
	return &TextArea{BaseComponent: NewBaseComponent("vn-text-area", base)}
}

func (c *TextArea) DeepCopy() Component {
	dst := &TextArea{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Min = intPP(c.Min)
	dst.Max = intPP(c.Max)
	return dst
}

type JsonEditor struct {
	BaseComponent `json:",inline"`

	// Example: 400px
	Height string `json:"height,omitempty"`
}

func NewJsonEditor(base BaseComponentOptions) *JsonEditor {
	return &JsonEditor{BaseComponent: NewBaseComponent("vn-json-editor", base)}
}

func (c *JsonEditor) DeepCopy() Component {
	dst := &JsonEditor{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
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

func (c *Radio) DeepCopy() Component {
	dst := &Radio{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	return dst
}

type RadioSet struct {
	BaseComponent `json:",inline"`
	*VenusOptions `json:",inline"`
	Column        bool `json:"column"`
}

func NewRadioSet(base BaseComponentOptions, options *VenusOptions) *RadioSet {
	return &RadioSet{BaseComponent: NewBaseComponent("vn-radio-set", base), VenusOptions: options}
}

func (c *RadioSet) DeepCopy() Component {
	dst := &RadioSet{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.VenusOptions = c.VenusOptions.DeepCopy()
	return dst
}

type OneOfForm struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Form  Form   `json:"settings"`
}
type OneOfForms []OneOfForm

type OneOf struct {
	BaseComponent `json:",inline"`
	Forms         OneOfForms `json:"form"`
}

func NewOneOf(base BaseComponentOptions, forms OneOfForms) *OneOf {
	return &OneOf{BaseComponent: NewBaseComponent("vn-oneof-forms", base), Forms: forms}
}

func (c *OneOf) DeepCopy() Component {
	dst := &OneOf{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Forms = nil
	for _, oneOfForm := range c.Forms {
		dst.Forms = append(dst.Forms, OneOfForm{
			Label: oneOfForm.Label,
			Key:   oneOfForm.Key,
			Form:  oneOfForm.Form.DeepCopy(),
		})
	}
	return dst
}

type Array struct {
	BaseComponent `json:",inline"`
	Min           *int `json:"min,omitempty"`
	Max           *int `json:"max,omitempty"`
	Form          Form `json:"form"`
}

func NewArray(base BaseComponentOptions) *Array {
	return &Array{BaseComponent: NewBaseComponent("vn-array", base)}
}

func (c *Array) DeepCopy() Component {
	dst := &Array{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Min = intPP(c.Min)
	dst.Max = intPP(c.Max)
	dst.Form = c.Form.DeepCopy()
	return dst
}

type Map struct {
	BaseComponent `json:",inline"`

	Key    Component `json:"key"`
	Values Form      `json:"values"`
}

func NewMap(base BaseComponentOptions) *Map {
	return &Map{BaseComponent: NewBaseComponent("vn-map", base)}
}

func (c *Map) DeepCopy() Component {
	dst := &Map{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Key = c.Key.DeepCopy()
	dst.Values = c.Values.DeepCopy()
	return dst
}

type Table struct {
	BaseComponent `json:",inline"`

	Key    Component   `json:"key"`
	Values []Component `json:"values"`
}

func NewTable(base BaseComponentOptions) *Table {
	return &Table{BaseComponent: NewBaseComponent("vn-table", base)}
}

func (c *Table) DeepCopy() Component {
	dst := &Table{}
	*dst = *c
	dst.BaseComponent = *c.BaseComponent.DeepCopy().(*BaseComponent)
	dst.Key = c.Key.DeepCopy()
	dst.Values = nil
	for _, value := range c.Values {
		dst.Values = append(dst.Values, value.DeepCopy())
	}
	return dst
}
