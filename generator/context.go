package generator

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// HierarchicalContext is a context for hierarchical and inheritable fields.
// It is used to determine whether a field should be exposed or not and to build a hierarchical property name.
// Or you can add additional inheritable properties.
type HierarchicalContext struct {
	exposes    []bool
	properties []string
}

func (c *HierarchicalContext) AppendExpose(expose *bool) {
	if expose != nil {
		c.exposes = append(c.exposes, *expose)
	}
}

func (c *HierarchicalContext) Expose() bool {
	if len(c.exposes) == 0 {
		return false
	}
	for _, expose := range c.exposes {
		if !expose {
			return false
		}
	}
	return true
}

func (c *HierarchicalContext) AppendProperty(property string) {
	c.properties = append(c.properties, property)
}

func (c *HierarchicalContext) AppendPropertyName(name protoreflect.Name) {
	c.properties = append(c.properties, string(name))
}

func (c *HierarchicalContext) PropertiesString() string {
	return strings.Join(c.properties, ".")
}

func NewFromHierarchicalContext(src *HierarchicalContext) *HierarchicalContext {
	dst := &HierarchicalContext{
		exposes:    make([]bool, len(src.exposes)),
		properties: make([]string, len(src.properties)),
	}
	copy(dst.exposes, src.exposes)
	copy(dst.properties, src.properties)
	return dst
}

func NewHierarchicalContext() *HierarchicalContext {
	return &HierarchicalContext{}
}
