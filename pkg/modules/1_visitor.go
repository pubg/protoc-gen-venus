package modules

import (
	"errors"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/modules/generator"
	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/graph"
)

type ProtoVisitor struct {
	pgs.Visitor

	debugger pgs.DebuggerCommon

	// container stores all the container that have been visited
	// Key: fully qualified name of the container
	// Value: incomplete generated venus components
	container map[graph.DescriptorReference]graph.Descriptor
}

func NewProtoVisitor(debugger pgs.DebuggerCommon) *ProtoVisitor {
	v := &ProtoVisitor{
		debugger:  debugger,
		container: make(map[graph.DescriptorReference]graph.Descriptor),
	}
	v.Visitor = pgs.PassThroughVisitor(v)
	return v
}

func (vv *ProtoVisitor) VisitPackage(pkg pgs.Package) (v pgs.Visitor, err error) {
	pd := graph.NewPackageDescriptor(protoNameToReference(pkg))
	vv.container[pd.Key()] = pd

	for _, file := range pkg.Files() {
		ref := fqdnToReference(file)
		desc := graph.NewLazyDescriptor[*graph.FileDescriptor](ref, vv.container)
		pd.Files = append(pd.Files, desc)
	}
	return vv, nil
}

func (vv *ProtoVisitor) VisitFile(file pgs.File) (v pgs.Visitor, err error) {
	fd := graph.NewFileDescriptor(fqdnToReference(file))
	vv.container[fd.Key()] = fd

	for _, message := range file.Messages() {
		ref := fqdnToReference(message)
		desc := graph.NewLazyDescriptor[*graph.MessageDescriptor](ref, vv.container)
		fd.Messages = append(fd.Messages, desc)
	}
	fo := protoptions.GetFileOptions(file)
	if fo == nil {
		return vv, nil
	}

	fd.Expose = fo.Expose
	fd.EntrypointMessage = fo.EntrypointMessage
	return vv, nil
}

func (vv *ProtoVisitor) VisitMessage(message pgs.Message) (pgs.Visitor, error) {
	md := graph.NewMessageDescriptor(fqdnToReference(message), message.Name().String())
	vv.container[md.Key()] = md

	for _, field := range message.Fields() {
		// Message Field
		// OneOfBlock Field
		// Scala Field
		if !field.InRealOneOf() {
			ref := fqdnToReference(field)
			desc := graph.NewLazyDescriptor[graph.Descriptor](ref, vv.container)
			md.Fields = append(md.Fields, desc)
		}
	}

	mo := protoptions.GetMessageOptions(message)
	if mo == nil {
		return vv, nil
	}

	md.Expose = mo.Expose
	return vv, nil
}

func (vv *ProtoVisitor) VisitField(field pgs.Field) (pgs.Visitor, error) {
	fo := protoptions.GetFieldOptions(field)

	// Phase 1: Well Known Type
	if generator.IsWellKnownKind(field) {
		// Initialize ScalaField Descriptor
		fd := graph.NewScalaFieldDescriptor(fqdnToReference(field))
		vv.container[fd.Key()] = fd
		vv.setFieldCommonFromFieldOptions(fd, fo, field)

		componentBuilder, err := generator.BuildFromScalaWellKnownField(field)
		if err != nil {
			return nil, err
		}
		fd.Component = componentBuilder
		return vv, nil
	}

	// Phase 2: Proto Specific Types, e.g. Map, Message
	if field.Type().ProtoType() == pgs.MessageT {
		// Initialize MessageField Descriptor
		fd := graph.NewMessageFieldDescriptor(fqdnToReference(field))
		vv.container[fd.Key()] = fd
		vv.setFieldCommonFromFieldOptions(fd, fo, field)

		fd.Message = graph.NewLazyDescriptor[*graph.MessageDescriptor](graph.DescriptorReference(field.Descriptor().GetTypeName()), vv.container)
		return vv, nil
	}

	// Phase 3: Scala Types, e.g. Int, String, Enum
	if generator.IsScalaKind(field) {
		// Initialize ScalaField Descriptor
		fd := graph.NewScalaFieldDescriptor(fqdnToReference(field))
		vv.container[fd.Key()] = fd
		vv.setFieldCommonFromFieldOptions(fd, fo, field)

		incompleteComponent, err := generator.BuildFromScalaField(field)
		if err != nil {
			return nil, err
		}
		fd.Component = incompleteComponent
		return vv, nil
	}

	// Below will exit process
	vv.debugger.Failf("Not supported field type: %v, The ProgramContext never reach here", field.Type())
	return nil, errors.New("failed to visit field")
}

func (vv *ProtoVisitor) setFieldCommonFromFieldOptions(accessor graph.FieldCommonAccessor, fo *protoptions.FieldOptions, field pgs.Field) {
	if fo != nil {
		accessor.SetExpose(fo.Expose)
	}
	if fo != nil && fo.GetProperty() != "" {
		accessor.SetProperty(fo.GetProperty())
	} else if field.Descriptor().GetJsonName() != "" {
		accessor.SetProperty(field.Descriptor().GetJsonName())
	} else {
		accessor.SetProperty(field.Name().String())
	}
}

func (vv *ProtoVisitor) VisitOneOf(oneOf pgs.OneOf) (pgs.Visitor, error) {
	od := graph.NewOneOfFieldDescriptor(fqdnToReference(oneOf))
	vv.container[od.Key()] = od

	for _, field := range oneOf.Fields() {
		ref := fqdnToReference(field)
		if field.Type().IsMap() {
			desc := graph.NewLazyDescriptor[*graph.MapFieldDescriptor](ref, vv.container)
			graph.AddFieldToOneOf[*graph.MapFieldDescriptor](od, desc)
		} else if field.Type().ProtoType() == pgs.MessageT {
			desc := graph.NewLazyDescriptor[*graph.MessageFieldDescriptor](ref, vv.container)
			graph.AddFieldToOneOf[*graph.MessageFieldDescriptor](od, desc)
		} else {
			desc := graph.NewLazyDescriptor[*graph.ScalaFieldDescriptor](ref, vv.container)
			graph.AddFieldToOneOf[*graph.ScalaFieldDescriptor](od, desc)
		}
	}

	oo := protoptions.GetOneOfOptions(oneOf)
	if oo != nil {
		od.SetExpose(oo.Expose)
	}
	if oo != nil && oo.GetProperty() != "" {
		od.SetProperty(oo.GetProperty())
	} else {
		od.SetProperty(oneOf.Name().String())
	}

	od.BaseComponentOptions = generator.ConcreteBaseComponentOptionsFromOneOf(oneOf, oo)
	return vv, nil
}

type FullyQualifiedNameResolver interface {
	FullyQualifiedName() string
}

func fqdnToReference(nameResolver FullyQualifiedNameResolver) graph.DescriptorReference {
	return graph.DescriptorReference(nameResolver.FullyQualifiedName())
}

type ProtoNameResolver interface {
	ProtoName() pgs.Name
}

func protoNameToReference(nameResolver ProtoNameResolver) graph.DescriptorReference {
	return graph.DescriptorReference(nameResolver.ProtoName().String())
}
