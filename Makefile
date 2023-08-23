test:
	go test ./... -v

install:
	go install cmd/protoc-gen-venus/main.go
