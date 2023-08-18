package generator

import (
	"fmt"
	"proc-gen-vlossom/generator/protooptions"
	"proc-gen-vlossom/generator/vlossom"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g *VlossomGenerator) buildFromScalaWellKnownField(ctx *HierarchicalContext, field *protogen.Field) (vlossom.Component, error) {
	fd := field.Desc
	fo := protooptions.GetFieldOptions(fd)

	base := g.concreteBaseComponentOptions(ctx, fd, fo)

	wellKnownKind := *ToWellKnownKind(fd)
	var componentType protooptions.ComponentType
	if fo.GetComponent() != protooptions.ComponentType_Inference {
		componentType = fo.GetComponent()
	} else {
		componentType = wellKnownKind.defaultComponent
	}

	switch wellKnownKind {
	case BooleanRepeatedKind:
		return g.buildFromBooleanRepeated(field, componentType, base)
	case JsonKind:
		return g.buildFromJsonField(field, componentType, base)
	}
	return nil, fmt.Errorf("unknown well-known kind")
}

func (g *VlossomGenerator) buildFromBooleanRepeated(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := protooptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protooptions.ComponentType_CheckboxSet:
		if fo.GetCheckboxSet() == nil {
			return nil, fmt.Errorf("failed buildFromBooleanRepeated, select options is nil")
		}
		return buildFromCheckboxSetOptions(fo.GetCheckboxSet(), convertToVlossomOptions(fo.GetCheckboxSet().GetOptions()), base), nil
	}
	return nil, fmt.Errorf("failed buildFromBooleanRepeated, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromJsonField(field *protogen.Field, componentType protooptions.ComponentType, base vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := protooptions.GetFieldOptions(field.Desc)
	switch componentType {
	case protooptions.ComponentType_JsonEditor:
		return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
	}
	return nil, fmt.Errorf("failed buildFromJsonField, unknown component type: %s", componentType)
}

type WellKnownKind struct {
	kind             string
	defaultComponent protooptions.ComponentType
}

var (
	BooleanRepeatedKind = WellKnownKind{
		kind:             "boolean-repeated",
		defaultComponent: protooptions.ComponentType_CheckboxSet,
	}
	JsonKind = WellKnownKind{
		kind:             "map",
		defaultComponent: protooptions.ComponentType_JsonEditor,
	}
)

func IsWellKnownKind(fd protoreflect.FieldDescriptor) bool {
	return ToWellKnownKind(fd) != nil
}

func ToWellKnownKind(fd protoreflect.FieldDescriptor) *WellKnownKind {
	if fd.Cardinality() == protoreflect.Repeated && fd.Kind() == protoreflect.BoolKind {
		return &BooleanRepeatedKind
	}
	if fd.IsMap() {
		return &JsonKind
	}

	if fd.Kind() == protoreflect.MessageKind {
		switch fd.Message().FullName() {
		case "google.protobuf.Any":
		case "google.protobuf.Timestamp":
		}
	}
	return nil
}
