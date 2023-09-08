package graph

type PackageDescriptor struct {
	baseDescriptor

	Files []*LazyDescriptor[*FileDescriptor]
}

func NewPackageDescriptor(key DescriptorReference) *PackageDescriptor {
	return &PackageDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypePackage}}
}

type FileDescriptor struct {
	baseDescriptor

	Expose            *bool
	EntrypointMessage string

	Messages []*LazyDescriptor[*MessageDescriptor]
}

func NewFileDescriptor(key DescriptorReference) *FileDescriptor {
	return &FileDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeFile}}
}

type MessageDescriptor struct {
	baseDescriptor
	Name string

	Expose *bool
	// Fields should be one of ScalaFieldDescriptor, MessageFieldDescriptor, DescriptorTypeMapField, OneOfFieldDescriptor
	Fields []*LazyDescriptor[Descriptor]
}

func NewMessageDescriptor(key DescriptorReference, name string) *MessageDescriptor {
	return &MessageDescriptor{baseDescriptor: baseDescriptor{key: key, dType: DescriptorTypeMessage}, Name: name}
}
