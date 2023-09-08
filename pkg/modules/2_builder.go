package modules

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/modules/generator"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
	"github.com/pubg/protoc-gen-venus/pkg/venus/graph"
)

type VenusComponentBuilder struct {
	debugger pgs.BuildContext
	ctx      *graph.HierarchicalContext
}

func NewVenusComponentBuilder(debugger pgs.BuildContext) *VenusComponentBuilder {
	return &VenusComponentBuilder{debugger: debugger, ctx: graph.NewHierarchicalContext()}
}

func (b *VenusComponentBuilder) buildFromFile(file *graph.FileDescriptor) []component.Component {
	b.ctx.PushExpose(file.Expose)
	defer b.ctx.PopExpose()

	for _, message := range file.Messages {
		if message.Value().Name == file.EntrypointMessage {
			return b.buildFromMessage(message.Value())
		}
	}
	b.debugger.Debugf("Entrypoint message not found in file: file: %s, entrypoint message: %s", file.Key(), file.EntrypointMessage)
	return nil
}

func (b *VenusComponentBuilder) buildFromMessage(message *graph.MessageDescriptor) []component.Component {
	b.ctx.PushExpose(message.Expose)
	defer b.ctx.PopExpose()

	var components []component.Component
	for _, field := range message.Fields {
		components = append(components, b.buildFromGenericField(field.Value())...)
	}
	return components
}

func (b *VenusComponentBuilder) buildFromGenericField(field graph.Descriptor) []component.Component {
	switch field.Type() {
	case graph.DescriptorTypeScalaField:
		comp := b.buildFromScalaField(field.(*graph.ScalaFieldDescriptor))
		if comp == nil {
			return nil
		}
		return []component.Component{comp}
	case graph.DescriptorTypeOneOfField:
		comp := b.buildFromOneOfField(field.(*graph.OneOfFieldDescriptor))
		if comp == nil {
			return nil
		}
		return []component.Component{comp}
	case graph.DescriptorTypeMapField:

		comp := b.buildFromMapField(field.(*graph.MapFieldDescriptor))
		if comp == nil {
			return nil
		}
		return []component.Component{comp}
	case graph.DescriptorTypeMessageField:
		return b.buildFromMessageField(field.(*graph.MessageFieldDescriptor))
	}

	// Below will exit process
	b.debugger.Failf("Not supported field type: %v, The ProgramContext never reach here", field.Type())
	return nil
}

func (b *VenusComponentBuilder) buildFromScalaField(field *graph.ScalaFieldDescriptor) component.Component {
	b.debugger.Push(field.GetProperty())
	b.ctx.PushProperty(field.GetProperty())
	b.ctx.PushExpose(field.GetExpose())
	defer b.debugger.Pop()
	defer b.ctx.PopProperty()
	defer b.ctx.PopExpose()
	b.debugger.Debugf("Building ScalaField")

	if !b.ctx.Expose() {
		return nil
	}

	comp := field.Component.DeepCopy()
	comp.SetPropertyName(b.ctx.Property())
	return comp
}

func (b *VenusComponentBuilder) buildFromOneOfField(oneOfField *graph.OneOfFieldDescriptor) component.Component {
	b.debugger.Push(oneOfField.GetProperty())
	b.ctx.PushProperty(oneOfField.GetProperty())
	b.ctx.PushExpose(oneOfField.GetExpose())
	defer b.debugger.Pop()
	defer b.ctx.PopProperty()
	defer b.ctx.PopExpose()
	b.debugger.Debugf("Building OneOfField")

	if !b.ctx.Expose() {
		return nil
	}

	concreteForm := func(field graph.Descriptor) component.OneOfForm {
		switch field.Type() {
		case graph.DescriptorTypeScalaField:
			scalaField := field.(*graph.ScalaFieldDescriptor)
			return component.OneOfForm{
				Label: scalaField.GetProperty(),
				Key:   scalaField.GetProperty(),
				Form:  [][]component.Component{{b.buildFromScalaField(scalaField)}},
			}
		case graph.DescriptorTypeMessageField:
			messageField := field.(*graph.MessageFieldDescriptor)
			return component.OneOfForm{
				Label: messageField.GetProperty(),
				Key:   messageField.GetProperty(),
				Form:  [][]component.Component{b.buildFromMessageField(messageField)},
			}
		case graph.DescriptorTypeMapField:
			mapField := field.(*graph.MapFieldDescriptor)
			return component.OneOfForm{
				Label: mapField.GetProperty(),
				Key:   mapField.GetProperty(),
				Form:  [][]component.Component{{b.buildFromMapField(mapField)}},
			}
		}
		// Below will exit process
		b.debugger.Failf("Not supported field type: %v, The ProgramContext never reach here", field.Type())
		return component.OneOfForm{}
	}

	var forms component.OneOfForms
	for _, fieldRef := range oneOfField.Fields {
		forms = append(forms, concreteForm(fieldRef.Value()))
	}
	return generator.BuildFromOneOfField(b.ctx, oneOfField, forms)
}

func (b *VenusComponentBuilder) buildFromMapField(field *graph.MapFieldDescriptor) component.Component {
	b.debugger.Push(field.GetProperty())
	b.ctx.PushProperty(field.GetProperty())
	b.ctx.PushExpose(field.GetExpose())
	defer b.debugger.Pop()
	defer b.ctx.PopProperty()
	defer b.ctx.PopExpose()
	b.debugger.Debugf("Building MapField")

	if !b.ctx.Expose() {
		return nil
	}

	toComponent := func(field graph.Descriptor) []component.Component {
		switch field.Type() {
		case graph.DescriptorTypeScalaField:
			scalaField := field.(*graph.ScalaFieldDescriptor)
			return []component.Component{b.buildFromScalaField(scalaField)}
		case graph.DescriptorTypeMessageField:
			messageField := field.(*graph.MessageFieldDescriptor)
			return b.buildFromMessageField(messageField)
		default:
			// Below will exit process
			b.debugger.Debugf("Not supported field type: %v, The ProgramContext never reach here", field.Type())
			return nil
		}
	}

	keyComponents := b.buildFromScalaField(field.MapKeyField.Value().(*graph.ScalaFieldDescriptor))
	valueComponents := toComponent(field.MapValueField.Value())
	return generator.BuildFromMapField(b.ctx, field, keyComponents, [][]component.Component{valueComponents})
}

func (b *VenusComponentBuilder) buildFromMessageField(field *graph.MessageFieldDescriptor) []component.Component {
	b.debugger.Push(field.GetProperty())
	b.ctx.PushProperty(field.GetProperty())
	b.ctx.PushExpose(field.GetExpose())
	defer b.debugger.Pop()
	defer b.ctx.PopProperty()
	defer b.ctx.PopExpose()
	b.debugger.Debugf("Building MessageField")

	if !b.ctx.Expose() {
		return nil
	}

	return b.buildFromMessage(field.Message.Value())
}
