package modules

import (
	"encoding/json"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/pubg/protoc-gen-venus/pkg/protoptions"
	"github.com/pubg/protoc-gen-venus/pkg/venus/component"
	"github.com/pubg/protoc-gen-venus/pkg/venus/graph"
	"google.golang.org/protobuf/encoding/protojson"
)

type VenusModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	po  protoptions.PluginOptions
}

func NewVenusModule() *VenusModule {
	return &VenusModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *VenusModule) Name() string {
	return "VenusModule"
}

func (m *VenusModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())

	if _, found := m.Parameters()["expose_all"]; found {
		exposeAll, _ := m.Parameters().Bool("expose_all")
		m.po.ExposeAll = &exposeAll
	}
	m.po.OutputFileSuffix = m.Parameters().StrDefault("output_file_suffix", ".venus.json")
	m.po.PrettyOutput, _ = m.Parameters().BoolDefault("pretty_output", true)
	m.Debugf("Parsed PluginOptions: %s", protojson.Format(&m.po))

}

func (m *VenusModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	// Phase 1
	// Convert Proto Descriptor to Venus Descriptor
	visitor := NewProtoVisitor(m)
	for _, pkg := range packages {
		m.Push(pkg.ProtoName().String())
		m.CheckErr(pgs.Walk(visitor, pkg), "failed to walk package")
		m.Pop()
	}

	// Phase 2
	// Generate Actual Venus Components
	builder := NewVenusComponentBuilder(m)
	builder.ctx.PushExpose(m.po.ExposeAll)
	componentsMap := map[string][]component.Component{}
	for name, file := range targets {
		fd := visitor.container[fqdnToReference(file)].(*graph.FileDescriptor)
		result := builder.buildFromFile(fd)
		if result != nil {
			componentsMap[name] = result
		}
	}

	// Phase 3
	// re-ordering and transform Venus Components
	transfomer := NewFlattenTransformer()
	formsMap := map[string]component.Form{}
	for name, components := range componentsMap {
		formsMap[name] = transfomer.Transform(components)
	}

	// Phase 4
	// Write to protoc result
	for name, form := range formsMap {
		content, err := m.generateFile(form)
		if err != nil {
			return nil
		}

		// example.proto -> example.venus.json
		fileName := targets[name].InputPath().SetExt(m.po.OutputFileSuffix).String()
		m.AddGeneratorFile(fileName, string(content))
	}
	return m.Artifacts()
}

func (m *VenusModule) generateFile(form component.Form) ([]byte, error) {
	if m.po.PrettyOutput {
		return json.MarshalIndent(form, "", "  ")
	} else {
		return json.Marshal(form)
	}
}
