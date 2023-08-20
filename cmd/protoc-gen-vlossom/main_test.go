package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pubg/protoc-gen-vlossom/generator"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

type Testcase struct {
	Name        string
	Description string

	Path        string
	ResultFile  string
	RequestFile string
}

const reqFile = "request.pb.bin"
const resFile = "entry.vlossom.json"

var testcases = []Testcase{
	{
		Name:        "Basic Generate",
		Description: "Just a basic test to make sure the plugin works",
		Path:        "../../testdata/cases/basic",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "Deny All Expose",
		Description: "Plugin should not expose any fields",
		Path:        "../../testdata/cases/deny-all-expose",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "Field Level Deny Expose",
		Description: "Some field are deny to expose by filed options, result should not contain those fields",
		Path:        "../../testdata/cases/field-level-deny-expose",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "File Level Inherit Expose",
		Description: "FileOptions.expose enabled, result should contain all fields in file",
		Path:        "../../testdata/cases/file-level-inherit-expose",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "Intermediate Deny Expose",
		Description: "Some field are deny to expose by field options, result should not contain those fields",
		Path:        "../../testdata/cases/intermediate-deny-expose",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "Options EnumValue",
		Description: "EnumValueOptions should work as expected",
		Path:        "../../testdata/cases/options-enumvalue",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "JsonName Generic",
		Description: "Can override property as protobuf's `jsonnanme` option",
		Path:        "../../testdata/cases/json-name-generic",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
	{
		Name:        "Message Level Inherit Expose",
		Description: "MessageOptions.expose enabled, result should contain all fields in message",
		Path:        "../../testdata/cases/message-level-inherit-expose",
		ResultFile:  resFile,
		RequestFile: reqFile,
	},
}

func TestPlugin(t *testing.T) {
	for _, testcase := range testcases {
		testcase := testcase
		expectedResult, err := os.ReadFile(filepath.Join(testcase.Path, testcase.ResultFile))
		if err != nil {
			t.Fatal(err)
		}

		requestBuf, err := os.ReadFile(filepath.Join(testcase.Path, testcase.RequestFile))
		if err != nil {
			t.Fatal(err)
		}
		req := &pluginpb.CodeGeneratorRequest{}
		if err := proto.Unmarshal(requestBuf, req); err != nil {
			t.Fatal(err)
		}

		t.Run(testcase.Name, func(t *testing.T) {
			response, err := run(req, &generator.PluginOptions{
				ExposeAll:        &[]bool{false}[0],
				OutputFileSuffix: &[]string{".vlossom.json"}[0],
			})
			if err != nil {
				t.Fatal(err)
			}

			if len(response.File) != 1 {
				t.Fatalf("response should contain 1 file, but got %d", len(response.File))
			}

			expects, err := toRawComponents(expectedResult)
			if err != nil {
				t.Fatal(err)
			}
			actuals, err := toRawComponents([]byte(*response.File[0].Content))
			if err != nil {
				t.Fatal(err)
			}

			for index, expect := range expects {
				require.Equal(t, expect, actuals[index], "not equals at index %s: %s", testcase.Name, testcase.Description)
			}
		})
	}
}

func run(request *pluginpb.CodeGeneratorRequest, options *generator.PluginOptions) (*pluginpb.CodeGeneratorResponse, error) {
	opts := protogen.Options{}

	plugin, err := opts.New(request)
	if err != nil {
		return nil, err
	}

	err = generator.NewVlossomGenerator(plugin, options).Run()
	if err != nil {
		return nil, err
	}
	return plugin.Response(), nil
}

func toRawComponents(buf []byte) ([]any, error) {
	containers := [][]any{}
	err := json.Unmarshal(buf, &containers)
	if err != nil {
		return nil, err
	}

	return lo.FlatMap[[]any, any](containers, func(container []any, _ int) []any { return container }), nil
}
