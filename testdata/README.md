# TestCase

Integration Test를 위한 폴더


## Test Case 구조
```dtd
.
├── cases
│   ├── basic
│   ├── deny-all-expose
│   ├── ... 테스트 케이스들
│   └── message-level-inherit-expose
└── hack
    ├── dump_test_cases.sh: 테스트를 위한 protoc request를 dump하는 스크립트 
    ├── generate_single_test_result.sh: 단일 테스트 케이스 결과를 만드는 스크립트 
    ├── generate_test_result.sh: 모든 테스트 케이스 결과를 만드는 스크립트
    └── redirect: dump_test_cases.sh에서 사용하는 redirect 파일
```

## 단일 Test Case 구조
```dtd
.
├── entry.proto: 테스트 케이스 입력 데이터
├── entry.venus.json: 예상되는 출력 값
└── request.pb.bin: protoc 의존성을 제거하기 위해 미리 dump해 둔 파일, 실제 테스트때 입력값으로 사용 
```

## 테스트 케이스 생성 방법
1. `cases` 폴더에 원하는 이름으로 새로운 폴더를 생성한다.
2. `hack/dump_test_cases.sh`를 실행한다. (workdir 위치는 상관없다)
3. `cmd/protoc-gen-venus/main_test.go`의 TestCases에 새로운 테스트 케이스를 추가한다.

## 테스트 실행하는 방법
1. repository root에서 `go test ./...`를 실행한다.
