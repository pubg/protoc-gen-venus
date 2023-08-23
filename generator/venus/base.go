package venus

type Component interface {
	_Component()
}

func NewBaseComponent(component string, option BaseComponentOptions) BaseComponent {
	base := BaseComponent{
		Component:    component,
		PropertyName: option.PropertyName,
		Label:        option.Label,
		Placeholder:  option.Placeholder,
		Required:     option.Required,
		State:        option.State,
		DefaultValue: option.DefaultValue,
		Messages:     option.Messages,
		Width:        option.Width,
		Grid:         option.Grid,
	}
	return base
}

type BaseComponentOptions struct {
	PropertyName string
	Label        string
	Placeholder  string
	Required     bool
	State        string
	DefaultValue any
	Messages     []Message
	Width        string
	Grid         *Grid
}

type BaseComponent struct {
	Component string `json:"component"`

	PropertyName string `json:"propertyName"`

	Label string `json:"label"`

	Placeholder string `json:"placeholder"`

	Required bool `json:"required"`

	State string `json:"state,omitempty"`

	DefaultValue any `json:"defaultValue,omitempty"`

	Messages []Message `json:"messages,omitempty"`

	// Width of component conflict with grid
	Width string `json:"width,omitempty"`
	// Grid of component conflict with width
	Grid *Grid `json:"grid,omitempty"`

	/**
	 * Not useful options
	 */
	// No Message? IDK
	NoMsg *bool `json:"no-msg,omitempty"`
	// Disable component
	Disabled *bool `json:"disabled,omitempty"`
	// Component is visible or not
	Visible *bool `json:"visible,omitempty"`
	// Cannot modify input form
	Readonly *bool `json:"readonly,omitempty"`
	// No Label? IDK, It doesn't work to my expected
	NoLabel *bool `json:"no-label,omitempty"`
	// No Clear? IDK
	NoClear *bool `json:"no-clear,omitempty"`
}

func (c *BaseComponent) _Component() {}

type Grid struct {
	Sm       *int `json:"sm,omitempty"`
	Md       *int `json:"md,omitempty"`
	Lg       *int `json:"lg,omitempty"`
	SmOffset *int `json:"smOffset,omitempty"`
	MdOffset *int `json:"mdOffset,omitempty"`
	LgOffset *int `json:"lgOffset,omitempty"`
	Order    *int `json:"order,omitempty"`
}

type Message struct {
	State string `json:"state"`
	Text  string `json:"text"`
}

type VenusOptionsOptions interface {
	_VenusOptionsOptions()
}

type SimpleOptions []string

func (o SimpleOptions) _VenusOptionsOptions() {}

type LabeledOptions []LabeledOption

func (o LabeledOptions) _VenusOptionsOptions() {}

type VenusOptions struct {
	Options     VenusOptionsOptions `json:"options,omitempty"`
	OptionLabel string              `json:"option-label,omitempty"`
	OptionValue string              `json:"option-value,omitempty"`
}

func NewVenusOptions() *VenusOptions {
	return &VenusOptions{}
}

func (o *VenusOptions) SetLabeledOptions(options LabeledOptions) {
	o.Options = options
	o.OptionValue = "value"
	o.OptionLabel = "label"
}

func (o *VenusOptions) SetSimpleOptions(options SimpleOptions) {
	o.Options = options
	o.OptionValue = ""
	o.OptionLabel = ""
}

type LabeledOption struct {
	Value any    `json:"value"`
	Label string `json:"label"`
}
