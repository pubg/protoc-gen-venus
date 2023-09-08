package graph

import (
	"strings"
)

// HierarchicalContext is a context for hierarchical and inheritable fields.
// It is used to determine whether a field should be exposed or not and to build a hierarchical property name.
// Or you can add additional inheritable properties.
type HierarchicalContext struct {
	exposes    []*bool
	properties []string
}

func (c *HierarchicalContext) PushProperty(property string) {
	c.properties = append(c.properties, property)
}

func (c *HierarchicalContext) PopProperty() {
	if len(c.properties) > 0 {
		c.properties = c.properties[:len(c.properties)-1]
	}
}

func (c *HierarchicalContext) Property() string {
	return strings.Join(c.properties, ".")
}

func (c *HierarchicalContext) PushExpose(exposes *bool) {
	c.exposes = append(c.exposes, exposes)
}

func (c *HierarchicalContext) PopExpose() {
	if len(c.exposes) > 0 {
		c.exposes = c.exposes[:len(c.exposes)-1]
	}
}

// Expose returns that all context's expose is true or not.
// nil is inherited from parent value.
// 1. [null, null, null, false] => false
// 2. [null, false, true, true] => false
// 3. [true, true, false, true] => false
// 4. [true] => true
// 5. [] => false
// 6. [true, true, true, true] => true
// 7. [true, true, null, null] => true
func (c *HierarchicalContext) Expose() bool {
	if len(c.exposes) == 0 {
		return false
	}
	hasTrue := false
	for _, expose := range c.exposes {
		if expose == nil {
			continue
		}
		if *expose {
			hasTrue = true
		}
		if !*expose {
			return false
		}
	}
	return hasTrue
}

func NewFromHierarchicalContext(src *HierarchicalContext) *HierarchicalContext {
	if src == nil {
		return NewHierarchicalContext()
	}

	dst := &HierarchicalContext{
		exposes:    make([]*bool, len(src.exposes)),
		properties: make([]string, len(src.properties)),
	}

	for _, expose := range src.exposes {
		dst.exposes = append(dst.exposes, expose)
	}
	copy(dst.properties, src.properties)
	return dst
}

func NewHierarchicalContext() *HierarchicalContext {
	return &HierarchicalContext{}
}
