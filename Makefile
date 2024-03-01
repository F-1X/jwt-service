# JWT
include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-jwt-api



generate-jwt-api:
	mkdir -p pkg/jwt
	protoc --proto_path api/jwt \
	--go_out=pkg/jwt --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/jwt --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/jwt/jwt.proto

# go 
go-build:
	go build -o bin/jwt cmd/jwt/* 

go-run: go-build
	./bin/jwt

# docker
# create-volume:
#     sudo docker volume create mongodb_data

# docker-up:
#     sudo docker-compose up

# docker-down:
#     sudo docker-compose down

# docker-rm:
#     sudo docker-compose rm 

# docker-restart:
# 	sudo docker-compose down -v 
# 	sudo docker-compose up

