package component

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Form [][]Component

func (f Form) DeepCopy() Form {
	if f == nil {
		return nil
	}
	dst := Form{}
	for _, row := range f {
		var newRow []Component
		for _, component := range row {
			newRow = append(newRow, component.DeepCopy())
		}
		dst = append(dst, newRow)
	}
	return dst
}

type Component interface {
	_Component()
	SetPropertyName(propertyName string)
	DeepCopy() Component
}

func NewBaseComponent(component string, option BaseComponentOptions) BaseComponent {
	base := BaseComponent{
		Component:    component,
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

func (c *BaseComponent) SetPropertyName(propertyName string) {
	c.PropertyName = propertyName
}

func (c *BaseComponent) DeepCopy() Component {
	dst := &BaseComponent{}
	*dst = *c
	if c.DefaultValue != nil {
		dst.DefaultValue = c.DefaultValue
	}
	dst.Messages = nil
	for _, message := range c.Messages {
		dst.Messages = append(dst.Messages, message.DeepCopy())
	}
	dst.Grid = c.Grid.DeepCopy()
	dst.NoMsg = boolPP(c.NoMsg)
	dst.Disabled = boolPP(c.Disabled)
	dst.Visible = boolPP(c.Visible)
	dst.Readonly = boolPP(c.Readonly)
	dst.NoLabel = boolPP(c.NoLabel)
	dst.NoClear = boolPP(c.NoClear)
	return dst
}

// DeepCopyWithoutPrimitivePointer is a deep copy function for component
// But limitations of gob package, it can't copy primitive pointer
// Example: copy to {min *int = &0} will be {min *int = nil}
// Cause golang assume {*int=&0} >= 0 means default value => pointer's default value is nil => nil
// Thus we need to use DeepCopyWithoutPrimitivePointer carefully
func DeepCopyWithoutPrimitivePointer[T Component](src T, dst T) {
	buffer := &bytes.Buffer{}
	err := gob.NewEncoder(buffer).Encode(src)
	if err != nil {
		panic(fmt.Sprintf("DeepCopyWithoutPrimitivePointer error: %v", err))
	}

	err = gob.NewDecoder(buffer).Decode(dst)
	if err != nil {
		panic(err)
	}
}

type Grid struct {
	Sm       *int `json:"sm,omitempty"`
	Md       *int `json:"md,omitempty"`
	Lg       *int `json:"lg,omitempty"`
	SmOffset *int `json:"smOffset,omitempty"`
	MdOffset *int `json:"mdOffset,omitempty"`
	LgOffset *int `json:"lgOffset,omitempty"`
	Order    *int `json:"order,omitempty"`
}

func (g *Grid) DeepCopy() *Grid {
	if g == nil {
		return nil
	}
	dst := &Grid{}
	dst.Sm = intPP(g.Sm)
	dst.Md = intPP(g.Md)
	dst.Lg = intPP(g.Lg)
	dst.SmOffset = intPP(g.SmOffset)
	dst.MdOffset = intPP(g.MdOffset)
	dst.LgOffset = intPP(g.LgOffset)
	dst.Order = intPP(g.Order)
	return dst
}

type Message struct {
	State string `json:"state"`
	Text  string `json:"text"`
}

func (m Message) DeepCopy() Message {
	return Message{
		State: m.State,
		Text:  m.Text,
	}
}

// VenusOptionsOptions Start
type VenusOptionsOptions interface {
	_VenusOptionsOptions()
	DeepCopy() VenusOptionsOptions
}

// SimpleOptions Start
type SimpleOptions []string

func (o SimpleOptions) _VenusOptionsOptions() {}

func (o SimpleOptions) DeepCopy() VenusOptionsOptions {
	dst := make(SimpleOptions, len(o))
	copy(dst, o)
	return dst
}

// LabeledOption Start
type LabeledOption struct {
	Value any    `json:"value"`
	Label string `json:"label"`
}
type LabeledOptions []LabeledOption

func (o LabeledOptions) _VenusOptionsOptions() {}

func (o LabeledOptions) DeepCopy() VenusOptionsOptions {
	dst := make(LabeledOptions, len(o))
	copy(dst, o)
	return dst
}

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

func (o *VenusOptions) DeepCopy() *VenusOptions {
	dst := &VenusOptions{}
	*dst = *o
	dst.Options = o.Options.DeepCopy()
	return dst
}

func init() {
	gob.Register(SimpleOptions{})
	gob.Register(SimpleOptions{})
	gob.Register(LabeledOption{})
	gob.Register(LabeledOptions{})
}

func intPP(i *int) *int {
	if i == nil {
		return nil
	}
	ii := *i
	return &ii
}

func intP(i int) *int {
	return &i
}

func boolPP(b *bool) *bool {
	if b == nil {
		return nil
	}
	bb := *b
	return &bb
}

func boolP(b bool) *bool {
	return &b
}
