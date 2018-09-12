ASSET_FILES = $(shell find ./api -type f)
ASSET_DIRS = $(shell find ./api -type d)
INTERNAL_SRC = $(shell find ./cmd ./internal -type f -name '*.go')
TAG ?= dev

build: internal/assets.go gcp

.PHONY: clean
clean:
	rm -vf gcp

internal/assets.go: $(ASSET_FILES)
	$(GOPATH)/bin/go-bindata -o ./internal/assets.go -pkg internal -prefix api/ $(ASSET_DIRS)

gcp: ./main.go $(INTERNAL_SRC) $(ASSET_FILES)
	go build -o gcp ./main.go

.PHONY: docker
docker:
	docker build --tag gitzup/gcp:$(TAG) --file ./build/Dockerfile .

.PHONY: docker
push-docker: docker
	docker push gitzup/gcp:$(TAG)
	docker tag gitzup/gcp:$(TAG) gitzup/gcp:latest
	docker push gitzup/gcp:latest
