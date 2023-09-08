package graph

type DescriptorType int

const (
	DescriptorTypePackage DescriptorType = iota
	DescriptorTypeFile
	DescriptorTypeMessage
	DescriptorTypeScalaField
	DescriptorTypeOneOfField
	DescriptorTypeMapField
	DescriptorTypeMessageField
)

type DescriptorReference string

type Descriptor interface {
	Key() DescriptorReference
	Type() DescriptorType
	_descriptor()
}

// baseDescriptor is the base struct for all descriptors
type baseDescriptor struct {
	key   DescriptorReference
	dType DescriptorType
}

func (d *baseDescriptor) Key() DescriptorReference {
	return d.key
}

func (d *baseDescriptor) Type() DescriptorType {
	return d.dType
}

func (d *baseDescriptor) _descriptor() {}

// LazyDescriptor is a wrapper around a descriptor that is lazily loaded from a container
type LazyDescriptor[D any] struct {
	key DescriptorReference

	desc      D
	loaded    bool
	container map[DescriptorReference]Descriptor
}

func NewLazyDescriptor[D any](key DescriptorReference, container map[DescriptorReference]Descriptor) *LazyDescriptor[D] {
	return &LazyDescriptor[D]{key: key, container: container}
}

func (ld *LazyDescriptor[D]) Key() DescriptorReference {
	return ld.key
}

func (ld *LazyDescriptor[D]) Value() D {
	if !ld.loaded {
		ld.desc = ld.container[ld.key].(D)
		ld.loaded = true
	}
	return ld.desc
}
