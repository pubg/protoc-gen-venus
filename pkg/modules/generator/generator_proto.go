package generator

import (
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
	"github.com/pubg/protoc-gen-venus/pkg/venus/graph"
)

func BuildFromOneOfField(ctx *graph.HierarchicalContext, oneOf *graph.OneOfFieldDescriptor, forms component.OneOfForms) component.Component {
	comp := component.NewOneOf(oneOf.BaseComponentOptions, forms)
	comp.SetPropertyName(ctx.Property())
	return comp
}

func BuildFromMapField(ctx *graph.HierarchicalContext, mapField *graph.MapFieldDescriptor, key component.Component, values component.Form) component.Component {
	comp := component.NewMap(mapField.BaseComponentOptions)
	comp.SetPropertyName(ctx.Property())
	comp.Key = key
	comp.Values = values
	return comp
}
