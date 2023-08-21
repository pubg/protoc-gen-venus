test:
	go test ./... -v

install:
	go install cmd/protoc-gen-vlossom/main.go
