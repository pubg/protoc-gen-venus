package generator

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"github.com/pubg/protoc-gen-venus/generator/venus"
	"github.com/samber/lo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type VenusGenerator struct {
	plugin  *protogen.Plugin
	options *protoptions.PluginOptions
}

func NewVenusGenerator(plugin *protogen.Plugin, options *protoptions.PluginOptions) *VenusGenerator {
	return &VenusGenerator{plugin: plugin, options: options}
}

func (g *VenusGenerator) Run() error {
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

		if len(components) == 0 {
			continue
		}

		buf, err := g.generateToJson(components)
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

func (g *VenusGenerator) buildFromFile(ctx *HierarchicalContext, file *protogen.File) ([]venus.Component, error) {
	fd := file.Desc
	fo := protoptions.GetFileOptions(fd)
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

func (g *VenusGenerator) buildFromMessage(ctx *HierarchicalContext, message *protogen.Message) ([]venus.Component, error) {
	md := message.Desc
	mo := protoptions.GetMessageOptions(md)

	ctx = NewFromHierarchicalContext(ctx)
	if mo != nil {
		ctx.AppendExpose(mo.Expose)
	}

	var resukt []venus.Component
	for _, field := range message.Fields {
		nestedComponents, err := g.buildFromField(ctx, field)
		if err != nil {
			return nil, err
		}
		resukt = append(resukt, nestedComponents...)
	}
	return resukt, nil
}

func (g *VenusGenerator) buildFromField(ctx *HierarchicalContext, field *protogen.Field) ([]venus.Component, error) {
	fd := field.Desc
	fo := protoptions.GetFieldOptions(fd)

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

	// Well-Known Types
	if IsWellKnownKind(fd) {
		if !ctx.Expose() {
			return nil, nil
		}
		component, err := g.buildFromScalaWellKnownField(ctx, field)
		if err != nil {
			return nil, err
		}
		return []venus.Component{component}, nil
	}

	// Repeated Types
	// TODO: Tool Part랑 어떻게 repeated 구현할지 논의해야 함
	_ = fd.Cardinality() == protoreflect.Repeated
	// Nested Types
	if fd.Kind() == protoreflect.GroupKind {
		return g.buildFromMessage(ctx, field.Message)
	} else if fd.Kind() == protoreflect.MessageKind && !fd.IsMap() {
		return g.buildFromMessage(ctx, field.Message)
	} else if fd.ContainingOneof() != nil {
		// TODO: Tool Part랑 어떻게 oneof 구현할지 논의해야 함
	}

	// Scala Types
	if !ctx.Expose() {
		return nil, nil
	}
	component, err := g.buildFromScalaField(ctx, field)
	if err != nil {
		return nil, err
	}
	return []venus.Component{component}, nil
}

func (g *VenusGenerator) generateToJson(components []venus.Component) ([]byte, error) {
	var transformed [][]venus.Component
	for _, component := range components {
		transformed = append(transformed, []venus.Component{component})
	}
	if *g.options.PrettyOutput {
		return json.MarshalIndent(transformed, "", "  ")
	} else {
		return json.Marshal(transformed)
	}
}

func generateFileName(file *protogen.File, suffix string) string {
	origin := file.Proto.GetName()
	ext := filepath.Ext(origin)
	return origin[:len(origin)-len(ext)] + suffix
}
