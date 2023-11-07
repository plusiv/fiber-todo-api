GOCMD=go
SWAGGERCMD=swag
GOBUILD=$(GOCMD) build
BINARY_NAME=server
DOCS_FILE_TO_IGNORE=docs/docs.go
MAIN_FILE=main.go

format:
	@echo "Formatting..."
	$(GOCMD) fmt ./...

swagger:
	@echo "Generating Swagger documentation..."
	$(SWAGGERCMD) fmt ./...
	$(SWAGGERCMD) init

build: swagger
	@echo "Building..."
	$(GOCMD) build -o $(BINARY_NAME) $(MAIN_FILE)

build-prod: format build

run: build
	./$(BINARY_NAME)

watch:
	reflex -G "$(DOCS_FILE_TO_IGNORE)" -s -r "\.go$$" make run

.PHONY: swagger build build-prod run watch format