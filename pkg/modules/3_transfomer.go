package modules

import "github.com/pubg/protoc-gen-venus/pkg/venus/component"

type Transformer interface {
	Transform(components []component.Component) component.Form
}

type FlattenTransformer struct {
}

func NewFlattenTransformer() *FlattenTransformer {
	return &FlattenTransformer{}
}

func (t *FlattenTransformer) Transform(components []component.Component) component.Form {
	var form component.Form
	for _, comp := range components {
		form = append(form, []component.Component{comp})
	}
	return form
}

type HierarchicalTransformer struct {
}

func NewHierarchicalTransformer() *HierarchicalTransformer {
	return &HierarchicalTransformer{}
}

func (t *HierarchicalTransformer) Transform(components []component.Component) component.Form {
	panic("not implemented")
}
