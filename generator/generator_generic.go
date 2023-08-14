package generator

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"proc-gen-vlossom/generator/vlossom"

	"github.com/samber/lo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type VlossomGenerator struct {
	plugin  *protogen.Plugin
	options *PluginOptions
}

func NewVlossomGenerator(plugin *protogen.Plugin, options *PluginOptions) *VlossomGenerator {
	return &VlossomGenerator{plugin: plugin, options: options}
}

func (g *VlossomGenerator) Run() error {
	ctx := NewHierarchicalContext()
	if *g.options.ExposeAll {
		ctx.AppendExpose(g.options.ExposeAll)
	}

	for _, file := range g.plugin.Files {
		if !file.Generate {
			continue
		}

		components, err := g.buildFromFile(ctx, file)
		if err != nil {
			return err
		}

		buf, err := generateToJson(components)
		if err != nil {
			return err
		}

		outputFile := g.plugin.NewGeneratedFile(generateFileName(file, *g.options.OutputFileSuffix), "")
		_, err = outputFile.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VlossomGenerator) buildFromFile(ctx *HierarchicalContext, file *protogen.File) ([]vlossom.Component, error) {
	fd := file.Desc
	fo := GetFileOptions(fd)
	if fo == nil {
		return nil, fmt.Errorf("FileOption must be set, File: %s", fd.Path())
	}

	ctx = NewFromHierarchicalContext(ctx)
	ctx.AppendExpose(fo.Expose)

	message, found := lo.Find[*protogen.Message](file.Messages, func(message *protogen.Message) bool {
		return fo.GetEntrypointMessage() == string(message.Desc.Name())
	})
	if !found {
		return nil, fmt.Errorf("cannot find matched entrypoint message, File: %s, EntrypointMessage: %s", fd.Path(), fo.GetEntrypointMessage())
	}
	return g.buildFromMessage(ctx, message)
}

func (g *VlossomGenerator) buildFromMessage(ctx *HierarchicalContext, message *protogen.Message) ([]vlossom.Component, error) {
	md := message.Desc
	mo := GetMessageOptions(md)

	ctx = NewFromHierarchicalContext(ctx)
	if mo != nil {
		ctx.AppendExpose(mo.Expose)
	}

	var resukt []vlossom.Component
	for _, field := range message.Fields {
		nestedComponents, err := g.buildFromField(ctx, field)
		if err != nil {
			return nil, err
		}
		resukt = append(resukt, nestedComponents...)
	}
	return resukt, nil
}

func (g *VlossomGenerator) buildFromField(ctx *HierarchicalContext, field *protogen.Field) ([]vlossom.Component, error) {
	fd := field.Desc
	fo := GetFieldOptions(fd)

	ctx = NewFromHierarchicalContext(ctx)
	if fo != nil {
		ctx.AppendExpose(fo.Expose)
	}
	if fo.GetProperty() != "" {
		ctx.AppendProperty(fo.GetProperty())
	} else if fd.HasJSONName() {
		ctx.AppendProperty(fd.JSONName())
	} else {
		ctx.AppendPropertyName(fd.Name())
	}

	// TODO: Tool Part랑 어떻게 repeated 구현할지 논의해야 함
	_ = fd.Cardinality() == protoreflect.Repeated

	// Nested Type
	if fd.Kind() == protoreflect.GroupKind {
		return g.buildFromMessage(ctx, field.Message)
	} else if fd.Kind() == protoreflect.MessageKind && !fd.IsMap() {
		return g.buildFromMessage(ctx, field.Message)
	} else if fd.ContainingOneof() != nil {
		// TODO: Tool Part랑 어떻게 oneof 구현할지 논의해야 함
	}

	// Scala Type
	if !ctx.Expose() {
		return nil, nil
	}
	component, err := g.buildFromScalaField(ctx, field)
	if err != nil {
		return nil, err
	}
	return []vlossom.Component{component}, nil
}

func generateToJson(components []vlossom.Component) ([]byte, error) {
	var transformed [][]vlossom.Component
	for _, component := range components {
		transformed = append(transformed, []vlossom.Component{component})
	}
	return json.Marshal(transformed)
}

func generateFileName(file *protogen.File, suffix string) string {
	origin := file.Proto.GetName()
	ext := filepath.Ext(origin)
	return origin[:len(origin)-len(ext)] + suffix
}
