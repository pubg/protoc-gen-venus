# protoc-gen-venus

venus의 Dynamic Form을 만드는 proto plugin

## Getting Started
1. protoc-gen-venus을 설치합니다.
```shell
go install github.com/pubg/protoc-gen-venus/cmd/protoc-gen-venus@latest
```
또는 [여기](https://github.com/pubg/protoc-gen-venus/releases)에서 다운로드 받아 사용합니다.
2. proto/options.proto를 Workdir에 복사합니다. 
3. 아래 내용대로 example.proto 파일을 생성합니다. 
```
syntax = "proto3";
package schema;
import "options.proto";

option go_package = "github.com/sample";
option (pubg.venus.file) = {expose: true, entrypointMessage: "Values"};

message Values {
  string sample_input = 1 [(pubg.venus.field) = {component: Input, input: {type: text, max: 10}}];
  MyEnum my_enum = 2;
}

enum MyEnum {
  FOO = 0;
  BAR = 1;
  BAZ = 2;
}

```
4. 다음 명령어를 실행합니다.
```shell
protoc \
  --venus_out=./ \
  -I ./ \
  ./example.proto
```

### 플러그인 옵션은 [Options.md](./Options.md)에서 확인할 수 있습니다.
### 로드맵은 [Roadmap.md](./Roadmap.md)에서 확인할 수 있습니다.

# How run without protoc (Standalone mode)
1. sh examples/debug_input_generate.sh 
2. cat examples/input.dump | go run main.go

protoc 없이 실행하기 위해선 GeneratedRequest(protobuf)를 생성해야 한다.

GeneratedRequest는 바이너리 포맷이라 텍스트로 관리하기 어렵다.

`debug_input_generate.sh` 스크립트를 사용해 example.proto로 부터 GeneratedRequest를 추출한다.

stdin으로 입력은 `cat examples/input.dump |` 또는 Goland에서 `다음 위의 입력 리디렉션`을 활성화하면 된다.

Happy Debugging!

# How run with protoc (Plugin mode)
1. sh examples/generate.sh
