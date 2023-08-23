package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pubg/protoc-gen-venus/generator"
	"github.com/pubg/protoc-gen-venus/generator/protoptions"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

type Testcase struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	RequestFile        string `json:"requestFile"`
	ExpectResultFile   string `json:"expectResultFile"`
	ExpectResultIsNull bool   `json:"expectResultIsNull"`
	//ExpectResultIsError bool   `json:"expectResultIsError"`
}

func TestPlugin(t *testing.T) {
	dirs, err := os.ReadDir("../../testdata/cases")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range dirs {
		testcase, testRequest, expectedResult, err := readTestCase("../../testdata/cases", dir)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(testcase.Name, func(t *testing.T) {
			response, err := toGenerateResponse(testRequest, &protoptions.PluginOptions{
				ExposeAll:        &[]bool{false}[0],
				OutputFileSuffix: &[]string{".venus.json"}[0],
				PrettyOutput:     &[]bool{false}[0],
			})
			if err != nil {
				t.Fatal(err)
			}

			if testcase.ExpectResultIsNull {
				if len(response.File) != 0 {
					t.Errorf("response should be empty, but got %d", len(response.File))
				}
				return
			}

			if len(response.File) != 1 {
				t.Errorf("response should contain 1 file, but got %d", len(response.File))
				return
			}

			actuals, err := toComparableComponent([]byte(*response.File[0].Content))
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			for index, expect := range expectedResult {
				require.Equal(t, expect, actuals[index], "not equals at index %s: %s", testcase.Name, testcase.Description)
			}
		})
	}
}

func readTestCase(parentDir string, dir os.DirEntry) (*Testcase, *pluginpb.CodeGeneratorRequest, []any, error) {
	if !dir.IsDir() {
		return nil, nil, nil, fmt.Errorf("dir %s is not directory", dir.Name())
	}

	path := filepath.Join(parentDir, dir.Name())
	testcase, err := readTestCase0(filepath.Join(path, "test.yaml"))
	if err != nil {
		return nil, nil, nil, err
	}

	request, err := readGeneratorRequest(filepath.Join(path, testcase.RequestFile))
	if err != nil {
		return nil, nil, nil, err
	}

	if testcase.ExpectResultIsNull {
		return testcase, request, nil, nil
	}

	result, err := readVenusResult(filepath.Join(path, testcase.ExpectResultFile))
	if err != nil {
		return nil, nil, nil, err
	}
	return testcase, request, result, nil
}

func readTestCase0(path string) (*Testcase, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	testcase := &Testcase{}
	if err := yaml.Unmarshal(buf, testcase); err != nil {
		return nil, err
	}
	return testcase, nil
}

func readGeneratorRequest(path string) (*pluginpb.CodeGeneratorRequest, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		return nil, err
	}
	return req, nil
}

func readVenusResult(path string) ([]any, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return toComparableComponent(buf)
}

func toGenerateResponse(request *pluginpb.CodeGeneratorRequest, options *protoptions.PluginOptions) (*pluginpb.CodeGeneratorResponse, error) {
	opts := protogen.Options{}

	plugin, err := opts.New(request)
	if err != nil {
		return nil, err
	}

	err = generator.NewVenusGenerator(plugin, options).Run()
	if err != nil {
		return nil, err
	}
	return plugin.Response(), nil
}

func toComparableComponent(buf []byte) ([]any, error) {
	containers := [][]any{}
	err := json.Unmarshal(buf, &containers)
	if err != nil {
		return nil, err
	}

	return lo.FlatMap[[]any, any](containers, func(container []any, _ int) []any { return container }), nil
}