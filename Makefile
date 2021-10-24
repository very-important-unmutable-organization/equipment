all: build down up

pull:
	docker-compose pull

push:
	docker-compose push

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

TOOLS_MOD_DIR = ./src/internal/tools

.PHONY: install-tools lint goimports fmt
install-tools:
	cd $(TOOLS_MOD_DIR) && go install golang.org/x/tools/cmd/goimports
	cd $(TOOLS_MOD_DIR) && go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-tools
	cd src && golangci-lint run --allow-parallel-runners ./...
