package vlossom

type Component interface {
	_Component()
}

func NewBaseComponent(component string, option BaseComponentOptions) BaseComponent {
	return BaseComponent{
		Component:    component,
		PropertyName: option.PropertyName,
		Label:        option.Label,
		Placeholder:  option.Placeholder,
		Required:     option.Required,
		Grid: Grid{
			LG: option.LG,
		},
	}
}

type BaseComponentOptions struct {
	PropertyName string
	Label        string
	Placeholder  string
	Required     bool
	LG           int
}

type BaseComponent struct {
	// Form Component Type
	// Example: input, select, text
	Component string `json:"component"`

	// Property Key
	// Example: metadata, workload.image
	PropertyName string `json:"propertyName"`

	// Displayed Component Name
	// Example: ID, First Name, School, Phone
	Label string `json:"label"`

	// Hint of component
	// Example: 000-0000-0000
	Placeholder string `json:"placeholder"`
	Required    bool   `json:"required"`
	Grid        Grid   `json:"grid"`
}

type Grid struct {
	LG int `json:"lg"`
}

func (c *BaseComponent) _Component() {
}
