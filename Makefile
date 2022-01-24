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

psql:
	docker exec -it equipment__db psql -U postgres

TOOLS_MOD_DIR = ./src/internal/tools

.PHONY: install-tools lint goimports fmt
install-tools:
	cd $(TOOLS_MOD_DIR) && go install golang.org/x/tools/cmd/goimports
	cd $(TOOLS_MOD_DIR) && go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-tools
	cd src && golangci-lint run --allow-parallel-runners ./...

fmt:
	cd src && go fmt ./...

goimports: install-tools
	cd src && goimports -w  -local github.com/very-important-unmutable-organization/equipment ./
