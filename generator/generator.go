package generator

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"proc-gen-vlossom/generator/vlossom"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type VlossomGenerator struct {
	plugin  *protogen.Plugin
	options *PluginOptions
}

func NewVlossomGenerator(plugin *protogen.Plugin, options *PluginOptions) *VlossomGenerator {
	return &VlossomGenerator{
		plugin:  plugin,
		options: options,
	}
}

func (g *VlossomGenerator) Run() error {
	exposePlugin := g.options.GetExposeAll()
	for _, file := range g.plugin.Files {
		if !file.Generate {
			continue
		}

		outputFile := g.plugin.NewGeneratedFile(generateFileName(file), "")
		components, err := g.buildFormFromFile(file, exposePlugin)
		if err != nil {
			return err
		}
		buf, err := g.generate(components)
		if err != nil {
			return err
		}
		_, err = outputFile.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateFileName(file *protogen.File) string {
	origin := file.Proto.GetName()
	ext := filepath.Ext(origin)
	return origin[:len(origin)-len(ext)] + ".vlossom.json"
}

func (g *VlossomGenerator) buildFormFromFile(file *protogen.File, parentExpose bool) ([]vlossom.Component, error) {
	fd := file.Desc
	fo := GetFileOptions(fd)
	exposeFile := parentExpose || fo.GetExpose()

	var components []vlossom.Component
	for _, message := range file.Messages {
		nestedComponents, err := g.buildFormFromMessage(message, exposeFile)
		if err != nil {
			return nil, err
		}
		components = append(components, nestedComponents...)
	}
	return components, nil
}

func (g *VlossomGenerator) buildFormFromMessage(message *protogen.Message, parentExpose bool) ([]vlossom.Component, error) {
	md := message.Desc
	mo := GetMessageOptions(md)
	exposeMessage := parentExpose || mo.GetExpose()

	var components []vlossom.Component
	for _, field := range message.Fields {
		fd := field.Desc
		fo := GetFieldOptions(fd)
		exposeField := exposeMessage || fo.GetExpose()

		// TODO: Tool Part랑 어떻게 repeated 구현할지 논의해야 함
		_ = fd.Cardinality() == protoreflect.Repeated

		if fd.Kind() == protoreflect.GroupKind {
			nestedComponents, err := g.buildFormFromMessage(field.Message, exposeField)
			if err != nil {
				return nil, err
			}
			components = append(components, nestedComponents...)

		} else if fd.Kind() == protoreflect.MessageKind && fd.IsMap() {
			// map은 repeated entry<key, value>와 동일하다.
			// map 일 때 repeated 적용 필요없음
			component, err := g.buildFromMapField(field, vlossom.BaseComponentOptions{})
			if err != nil {
				return nil, err
			}
			components = append(components, component)

		} else if fd.Kind() == protoreflect.MessageKind && !fd.IsMap() {
			nestedComponents, err := g.buildFormFromMessage(field.Message, exposeField)
			if err != nil {
				return nil, err
			}
			components = append(components, nestedComponents...)

		} else if fd.ContainingOneof() != nil {
			// TODO: Tool Part랑 어떻게 oneof 구현할지 논의해야 함
		} else {
			if !(exposeField) {
				continue
			}

			component, err := g.buildFromScalaField(field)
			if err != nil {
				return nil, err
			}
			components = append(components, component)
		}
	}
	return components, nil
}

func (g *VlossomGenerator) buildFromScalaField(field *protogen.Field) (vlossom.Component, error) {
	fd := field.Desc
	fo := GetFieldOptions(fd)

	componentOptions := vlossom.BaseComponentOptions{
		PropertyName: string(fd.FullName().Name()),
		Required:     !fd.HasOptionalKeyword(),
	}

	if fo.GetLabel() == "" {
		componentOptions.Label = fd.TextName()
	} else {
		componentOptions.Label = fo.GetLabel()
	}
	if fo.GetPlaceholder() == "" {
		componentOptions.Placeholder = fd.TextName()
	} else {
		componentOptions.Placeholder = fo.GetPlaceholder()
	}
	if fo.GetLg() == 0 {
		componentOptions.LG = 5
	} else {
		componentOptions.LG = int(fo.GetLg())
	}

	scalaKind := ToScalaKind(fd.Kind())
	var componentType ComponentType
	if fo.GetComponent() == ComponentType_Default {
		componentType = scalaKind.defaultComponent
	} else {
		componentType = fo.GetComponent()
	}

	switch scalaKind {
	case NumberKind:
		return g.buildFromNumberField(field, componentType, componentOptions)
	case StringKind:
		return g.buildFromStringField(field, componentType, componentOptions)
	case EnumKind:
		return g.buildFromEnumField(field, componentType, componentOptions)
	case BooleanKind:
		return g.buildFromBooleanField(field, componentType, componentOptions)
	}
	return nil, fmt.Errorf("unknown scala kind")
}

func (g *VlossomGenerator) buildFromNumberField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fo := GetFieldOptions(field.Desc)
	switch componentType {
	case ComponentType_Input:
		var max *int
		if fo.GetString_().Max != nil {
			i := int(*fo.GetString_().Max)
			max = &i
		}
		return vlossom.NewInput(componentOptions, max), nil
	}
	return nil, fmt.Errorf("failed buildFromNumberField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromStringField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	switch componentType {
	case ComponentType_Input:
		return vlossom.NewInput(componentOptions, nil), nil
	case ComponentType_Select:
		return vlossom.NewSelect(componentOptions, nil), nil
	case ComponentType_RadioSet:
		return vlossom.NewRadioSet(componentOptions, nil), nil
	case ComponentType_DateRangePicker:
		return vlossom.NewDateRangePicker(componentOptions), nil
	case ComponentType_DateTimePicker:
		return vlossom.NewDateTimePicker(componentOptions), nil
	}
	return nil, fmt.Errorf("failed buildFromStringField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromEnumField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	fd := field.Desc
	_ = GetFieldOptions(fd)
	ed := fd.Enum()
	_ = GetEnumOptions(ed)
	values := ed.Values()

	var options []any
	for i := 0; i < values.Len(); i++ {
		vd := values.Get(i)
		vo := GetEnumValueOptions(vd)

		var option any
		if x, ok := vo.GetValue().(*EnumValueOptions_String_); ok {
			option = x.String_
		} else if x, ok := vo.GetValue().(*EnumValueOptions_Integer); ok {
			option = x.Integer
		} else if x, ok := vo.GetValue().(*EnumValueOptions_Float); ok {
			option = x.Float
		} else {
			option = vd.Name()
		}
		options = append(options, option)
	}

	switch componentType {
	case ComponentType_Select:
		return vlossom.NewSelect(componentOptions, options), nil
	case ComponentType_RadioSet:
		return vlossom.NewRadioSet(componentOptions, options), nil
	}
	return nil, fmt.Errorf("failed buildFromEnumField, unknown component type: %s", componentType)
}

// TODO: Toggle 추가해야 함
func (g *VlossomGenerator) buildFromBooleanField(field *protogen.Field, componentType ComponentType, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	switch componentType {
	case ComponentType_Checkbox:
		return vlossom.NewCheckbox(componentOptions), nil
	}
	return nil, fmt.Errorf("failed buildFromBooleanField, unknown component type: %s", componentType)
}

func (g *VlossomGenerator) buildFromMapField(field *protogen.Field, componentOptions vlossom.BaseComponentOptions) (vlossom.Component, error) {
	return vlossom.NewJsonEditor(componentOptions), nil
}

func (g *VlossomGenerator) generate(components []vlossom.Component) ([]byte, error) {
	var transformed [][]vlossom.Component
	for _, component := range components {
		transformed = append(transformed, []vlossom.Component{component})
	}
	return json.Marshal(transformed)
}

// GetFileOptions null을 반환할 수 있지만, null 이라도 리시버 호출 됨
func GetFileOptions(desc protoreflect.Descriptor) *FileOptions {
	options := desc.Options()
	if options != nil {
		return nil
	}
	if !proto.HasExtension(options, E_FileOptions) {
		return nil
	}
	return proto.GetExtension(options, E_FileOptions).(*FileOptions)
}

func GetMessageOptions(desc protoreflect.Descriptor) *MessageOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_MessageOptions) {
		return nil
	}
	return proto.GetExtension(options, E_MessageOptions).(*MessageOptions)
}

func GetFieldOptions(desc protoreflect.Descriptor) *FieldOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_FieldOptions) {
		return nil
	}
	return proto.GetExtension(options, E_FieldOptions).(*FieldOptions)
}

func GetEnumOptions(desc protoreflect.Descriptor) *EnumOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_EnumOptions) {
		return nil
	}
	return proto.GetExtension(options, E_EnumOptions).(*EnumOptions)
}

func GetEnumValueOptions(desc protoreflect.Descriptor) *EnumValueOptions {
	options := desc.Options()
	if options == nil {
		return nil
	}
	if !proto.HasExtension(options, E_EnumOptions) {
		return nil
	}
	return proto.GetExtension(options, E_EnumValueOptions).(*EnumValueOptions)
}
