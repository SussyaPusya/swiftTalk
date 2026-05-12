
PROTO_PATH := api
OUTPUT_PATH := api
BINARY_NAME := auth_service

# Цель по умолчанию
all: build

# Генерация кода из .proto файлов
generate:
	protoc --proto_path=$(PROTO_PATH) --go_out=$(OUTPUT_PATH) --go_opt=paths=source_relative --go-grpc_out=$(OUTPUT_PATH) --go-grpc_opt=paths=source_relative $(PROTO_PATH)/auth_service.proto

ping:
	ls


# Сборка бинарника
build: generate
	go build -o $(BINARY_NAME) cmd/main.go

# Очистка сгенерированных файлов и бинарника
clean:
	rm -rf $(OUTPUT_PATH) $(BINARY_NAME)

# Установка необходимых инструментов
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

.PHONY: all generate build clean install-tools ping
