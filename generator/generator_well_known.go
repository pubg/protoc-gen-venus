package generator

import (
	"fmt"

	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"github.com/pubg/protoc-gen-venus/generator/venus"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g *VenusGenerator) buildFromScalaWellKnownField(ctx *HierarchicalContext, field *protogen.Field) (venus.Component, error) {
	fd := field.Desc
	fo := protoptions.GetFieldOptions(fd)

	base := concreteBaseComponentOptions(ctx, fd, fo)
	wellKnownKind := *ToWellKnownKind(fd)
	componentType := getDesiredOrDefaultComponent(fo.GetComponent(), wellKnownKind.defaultComponent)

	switch wellKnownKind {
	case BooleanRepeatedKind:
		return g.buildFromBooleanRepeated(field, componentType, base)
	case StringRepeatedKind:
		return g.buildFromStringRepeated(field, componentType, base)
	case JsonKind:
		return g.buildFromJsonField(field, componentType, base)
	case AnyKind:
		return g.buildFromAnyField(field, componentType, base)
	case TimestampKind:
		return g.buildFromTimestampField(field, componentType, base)
	}
	return nil, fmt.Errorf("unknown well-known kind")
}

func (g *VenusGenerator) buildFromBooleanRepeated(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	if fo.GetCheckboxSet() == nil || fo.GetCheckboxSet().GetOptions() == nil {
		return nil, fmt.Errorf("failed buildFromBooleanRepeated, select option is required")
	}
	return buildFromCheckboxSetOptions(fo.GetCheckboxSet(), convertToVenusOptions(fo.GetCheckboxSet().GetOptions()), base), nil
}

func (g *VenusGenerator) buildFromStringRepeated(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	return buildFromMultiStringOptions(fo.GetMultiString(), base), nil
}

func (g *VenusGenerator) buildFromJsonField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
}

func (g *VenusGenerator) buildFromAnyField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	fo := protoptions.GetFieldOptions(field.Desc)
	return buildFromTextAreaOptions(fo.GetTextArea(), base), nil
}

func (g *VenusGenerator) buildFromTimestampField(field *protogen.Field, componentType protoptions.ComponentType, base venus.BaseComponentOptions) (venus.Component, error) {
	return venus.NewDateTimePicker(base), nil
}

type WellKnownKind struct {
	kind             string
	defaultComponent protoptions.ComponentType
}

var (
	BooleanRepeatedKind = WellKnownKind{
		kind:             "boolean-repeated",
		defaultComponent: protoptions.ComponentType_CheckboxSet,
	}
	StringRepeatedKind = WellKnownKind{
		kind:             "string-repeated",
		defaultComponent: protoptions.ComponentType_MultiString,
	}
	JsonKind = WellKnownKind{
		kind:             "map",
		defaultComponent: protoptions.ComponentType_JsonEditor,
	}
	AnyKind = WellKnownKind{
		kind:             "any",
		defaultComponent: protoptions.ComponentType_TextArea,
	}
	TimestampKind = WellKnownKind{
		kind:             "timestamp",
		defaultComponent: protoptions.ComponentType_DateTimePicker,
	}
)

func IsWellKnownKind(fd protoreflect.FieldDescriptor) bool {
	return ToWellKnownKind(fd) != nil
}

func ToWellKnownKind(fd protoreflect.FieldDescriptor) *WellKnownKind {
	if fd.Cardinality() == protoreflect.Repeated && fd.Kind() == protoreflect.BoolKind {
		return &BooleanRepeatedKind
	}
	if fd.Cardinality() == protoreflect.Repeated && fd.Kind() == protoreflect.StringKind {
		return &StringRepeatedKind
	}
	if fd.IsMap() {
		return &JsonKind
	}

	// TODO: well-known type 추가해야 함
	if fd.Kind() == protoreflect.MessageKind {
		switch fd.Message().FullName() {
		case "google.protobuf.Any":
			return &AnyKind
		case "google.protobuf.Timestamp":
			return &TimestampKind
		case "google.protobuf.Duration":
		case "k8s.io.apimachinery.pkg.apis.util.v1.IntOrString":
		case "k8s.io.api.pkg.core.v1.Volume":
		}
	}
	return nil
}
