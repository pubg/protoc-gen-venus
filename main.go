package main

import (
	"github.com/lyft/protoc-gen-star/v2"
	"github.com/pubg/protoc-gen-venus/pkg/modules"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	feature := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	pgs.Init(
		pgs.DebugEnv("DEBUG"),
		pgs.SupportedFeatures(&feature)).
		RegisterModule(modules.NewVenusModule()).
		Render()
}
