PROTO_FILES = $(wildcard src/protos/*.proto)
PROTOC_GEN_GO=$(shell go list -f '{{.Target}}' github.com/golang/protobuf/protoc-gen-go)
ICEBRIDGE := $(CURDIR)
PROTOS = $(ICEBRIDGE)/src/protos
PBS = src/protos/pbs

all: $(PROTO_FILES:.proto=.pb.go)

%.pb.go: %.proto
	protoc --go_out=$(PBS) \
	--go_opt=paths=source_relative \
	--proto_path=$(PROTOS) \
	--go-grpc_out=$(PBS) \
	--go-grpc_opt=paths=source_relative \
	$(ICEBRIDGE)/src/protos/*.proto
	



create-project:
	go mod init github.com/datastx/IceBridge

tidy:
	go mod tidy

run:
	go run src/main.go -c poop.json -d mydb -p postgres -s snowflake

test:
	go test -covermode=atomic -cover -race ./...