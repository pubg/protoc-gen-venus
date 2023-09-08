package graph

import "github.com/pubg/protoc-gen-venus/pkg/venus/component"

type OneOfFieldDescriptor struct {
	baseDescriptor

	fieldCommon
	BaseComponentOptions component.BaseComponentOptions
	// Fields should be one of ScalaFieldDescriptor, MessageFieldDescriptor, MapFieldDescriptor
	Fields []*LazyDescriptor[Descriptor]
}

func NewOneOfFieldDescriptor(key DescriptorReference) *OneOfFieldDescriptor {
	return &OneOfFieldDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeOneOfField}}
}

type AddFieldToOneOfConstraints interface {
	*ScalaFieldDescriptor | *MessageFieldDescriptor | *MapFieldDescriptor
}

func AddFieldToOneOf[D AddFieldToOneOfConstraints](oneOf *OneOfFieldDescriptor, field *LazyDescriptor[D]) {
	oneOf.Fields = append(oneOf.Fields, NewLazyDescriptor[Descriptor](field.Key(), nil))
}

type MapFieldDescriptor struct {
	baseDescriptor

	fieldCommon
	BaseComponentOptions component.BaseComponentOptions
	// KeyField should be one of ScalaFieldDescriptor, MessageFieldDescriptor
	MapKeyField *LazyDescriptor[Descriptor]
	// ValueField should be one of ScalaFieldDescriptor, MessageFieldDescriptor
	MapValueField *LazyDescriptor[Descriptor]
}

func NewMapFieldDescriptor(key DescriptorReference) *MapFieldDescriptor {
	return &MapFieldDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeMapField}}
}

type MessageFieldDescriptor struct {
	baseDescriptor

	fieldCommon
	Repeated bool
	Message  *LazyDescriptor[*MessageDescriptor]
}

func NewMessageFieldDescriptor(key DescriptorReference) *MessageFieldDescriptor {
	return &MessageFieldDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeMessageField}}
}

type ScalaFieldDescriptor struct {
	baseDescriptor

	fieldCommon
	Repeated  bool
	Component component.Component
}

func NewScalaFieldDescriptor(key DescriptorReference) *ScalaFieldDescriptor {
	return &ScalaFieldDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeScalaField}}
}

type FieldCommonAccessor interface {
	GetExpose() *bool
	SetExpose(expose *bool)
	GetProperty() string
	SetProperty(property string)
}

// FieldCommon is a common struct for all field descriptors
type fieldCommon struct {
	expose   *bool
	property string
}

func (f *fieldCommon) GetExpose() *bool {
	return f.expose
}

func (f *fieldCommon) SetExpose(expose *bool) {
	f.expose = expose
}

func (f *fieldCommon) GetProperty() string {
	return f.property
}

func (f *fieldCommon) SetProperty(property string) {
	f.property = property
}
