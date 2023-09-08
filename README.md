# protoc-gen-venus

protoc-gen-venus는 protobuf 파일을 기반으로 `Venus Dynamic Form` 을 생성하는 플러그인입니다.

# For Plugin Users

# Getting Started
### protoc-gen-venus 설치
1. install via go get
```shell
go install github.com/pubg/protoc-gen-venus/cmd/protoc-gen-venus@latest
``` 
2. install via release

[여기](https://github.com/pubg/protoc-gen-venus/releases)에서 다운로드 받아 사용합니다.

### venus.proto 복사
`proto/venus.proto` 를 Working Directory로 복사합니다.

### example.proto 생성 
```
syntax = "proto3";
package schema;
import "venus.proto";

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

### protoc 실행
```shell
protoc \
  --venus_out=./ \
  -I ./ \
  ./example.proto
```

[Options](./Options.md)

[Roadmap](./Roadmap.md)

---------------

# For Plugin Developers 

[Workflow](./Workflow.md)

# How to run to normal way (Plugin mode)
```shell
sh examples/generate.sh
```

# How to run without protoc (Standalone mode)
1. sh examples/debug_input_generate.sh
2. cat examples/input.dump | go run main.go

Go 프로그램 디버그를 위해 protoc 의존성을 제거해야 할 때가 있다. 그러나 protoc가 플러그인의 생명주기를 관리하기 때문에, protoc를 사용하지 않고도 동작하는 방법이 필요하다.

1. `debug_input_generate.sh` 스크립트를 사용해 example.proto로 부터 GeneratedRequest를 추출한다.
2. stdin으로 입력은 `cat examples/input.dump |` 또는 Goland에서 `다음 위의 입력 리디렉션`을 활성화하면 된다.
3. Happy Debugging!
