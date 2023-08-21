package main

import (
	"flag"

	"github.com/pubg/protoc-gen-vlossom/generator"
	"github.com/pubg/protoc-gen-vlossom/generator/protooptions"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var flags flag.FlagSet
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	conf := &protooptions.PluginOptions{
		ExposeAll:        flags.Bool("expose_all", false, `expose all fields, By default, only fields annotated with 'expose' are exposed.`),
		OutputFileSuffix: flags.String("output_file_suffix", ".vlossom.json", `output file suffix`),
		PrettyOutput:     flags.Bool("pretty_output", true, `pretty format json output`),
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generator.NewVlossomGenerator(plugin, conf).Run()
	})
}
