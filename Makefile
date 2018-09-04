ASSET_FILES = $(shell find ./api -type f)
ASSET_DIRS = $(shell find ./api -type d)
INTERNAL_SRC = $(shell find ./internal -type f -name '*.go')
TAG ?= dev

build: internal/assets.go resource

.PHONY: clean
clean:
	rm -vf agent

internal/assets.go: $(ASSET_FILES)
	$(GOPATH)/bin/go-bindata -o ./internal/assets.go -pkg internal -prefix api/ $(ASSET_DIRS)

resource: ./cmd/resource.go $(INTERNAL_SRC) $(ASSET_FILES)
	go build ./cmd/resource.go

.PHONY: docker
docker:
	docker build --tag gitzup/gcp-base:$(TAG) --file ./build/Dockerfile .
	docker tag gitzup/gcp-base:$(TAG) gitzup/gcp-project:$(TAG)

.PHONY: docker
push-docker: docker
	docker push gitzup/gcp-base:$(TAG)
	docker push gitzup/gcp-project:$(TAG)
	docker tag gitzup/gcp-base:$(TAG) gitzup/gcp-base:latest
	docker tag gitzup/gcp-project:$(TAG) gitzup/gcp-project:latest
	docker push gitzup/gcp-base:latest
	docker push gitzup/gcp-project:latest
