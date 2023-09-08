package generator

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
)

func BuildFromScalaWellKnownField(field pgs.Field) (component.Component, error) {
	fo := protoptions.GetFieldOptions(field)

	base := ConcreteBaseComponentOptions(field, fo)
	wellKnownKind := *ToWellKnownKind(field)
	componentType := getDesiredOrDefaultComponent(fo.GetComponent(), wellKnownKind.defaultComponent)

	switch wellKnownKind {
	case BooleanRepeatedKind:
		return buildFromBooleanRepeated(field, componentType, base, fo)
	case StringRepeatedKind:
		return buildFromStringRepeated(field, componentType, base, fo)
	case JsonKind:
		return buildFromJsonField(field, componentType, base, fo)
	case AnyKind:
		return buildFromAnyField(field, componentType, base, fo)
	case TimestampKind:
		return buildFromTimestampField(field, componentType, base, fo)
	}
	return nil, fmt.Errorf("unknown well-known kind")
}

func buildFromBooleanRepeated(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	if fo.GetCheckboxSet() == nil || fo.GetCheckboxSet().GetOptions() == nil {
		return nil, fmt.Errorf("failed buildFromBooleanRepeated, select option is required")
	}
	return buildFromCheckboxSetOptions(fo.GetCheckboxSet(), convertToVenusOptions(fo.GetCheckboxSet().GetOptions()), base), nil
}

func buildFromStringRepeated(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	return buildFromMultiStringOptions(fo.GetMultiString(), base), nil
}

func buildFromJsonField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	return buildFromJsonEditorOptions(fo.GetJsonEditor(), base), nil
}

func buildFromAnyField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	return buildFromTextAreaOptions(fo.GetTextArea(), base), nil
}

func buildFromTimestampField(field pgs.Field, componentType protoptions.ComponentType, base component.BaseComponentOptions, fo *protoptions.FieldOptions) (component.Component, error) {
	return component.NewDateTimePicker(base), nil
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

func IsWellKnownKind(field pgs.Field) bool {
	return ToWellKnownKind(field) != nil
}

func ToWellKnownKind(field pgs.Field) *WellKnownKind {
	ft := field.Type()
	if ft.IsRepeated() && ft.ProtoType() == pgs.BoolT {
		return &BooleanRepeatedKind
	}
	if ft.IsRepeated() && ft.ProtoType() == pgs.StringT {
		return &StringRepeatedKind
	}
	if ft.IsMap() {
		return &JsonKind
	}

	// TODO: well-known type 추가해야 함
	if ft.IsEmbed() {
		switch field.Message().FullyQualifiedName() {
		case "google.protobuf.Any":
			return &AnyKind
		case "google.protobuf.Timestamp":
			return &TimestampKind
		case "google.protobuf.Duration":
		case "k8s.io.apimachinery.pkg.apis.util.v1.IntOrString":
			// oneof int or string?
			// input with int or string validation?
			// k8s에서 int or string이 어떤 용례로 쓰이는지 찾아봐야 할듯
		case "k8s.io.api.pkg.core.v1.Volume":
		}
	}
	return nil
}
