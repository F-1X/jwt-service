install-protobuf:
	apt install -y protobuf-compiler

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate:
	make generate-jwt-api

generate-jwt-api-protoc:
	mkdir -p pkg/api/jwt
	protoc --proto_path api/jwt \
	--go_out=pkg/api/jwt --go_opt=paths=source_relative api/jwt/jwt.proto \
	 --plugin=protoc-gen-go=bin/protoc-gen-go

generate-jwt-api:
	mkdir -p pkg/jwtGRPC
	protoc --proto_path api/jwt \
	--go_out=pkg/jwtGRPC --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/jwtGRPC --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/jwt/jwt.proto

# go 
go-build:
	@go build -o bin/jwt cmd/jwt/* 

go-run: go-build
	@./bin/jwt


docker-up:
	sudo docker-compose up

# docker
# create-volume:
#     sudo docker volume create mongodb_data


# docker-down:
#     sudo docker-compose down

# docker-rm:
#     sudo docker-compose rm 

docker-restart:
	sudo docker-compose down -v 
	sudo docker-compose up

