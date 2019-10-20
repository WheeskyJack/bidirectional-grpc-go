all: proto server client

proto:
	cd ../proto && protoc --go_out=plugins=grpc:. *.proto

server:
	cd ../server/ && go build .

client:
	cd ../client/ && go build .

.PHONY: all proto server client