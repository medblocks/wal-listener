IMAGE_PREFIX ?= theaxer

.GIT_COMMIT=$(shell git rev-parse HEAD)
.GIT_VERSION=$(shell git describe --tags --always --dirty 2>/dev/null)
.GIT_UNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(.GIT_UNTRACKEDCHANGES),)
	.GIT_VERSION := $(.GIT_VERSION)-$(shell date +"%s")
endif


wal-listener: $(shell find ./ -name '*.go') go.mod go.sum
	@echo "+ build $@"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ -v -ldflags "\
		-X github.com/ihippik/wal-listener/config.GitCommit=${.GIT_COMMIT} \
		-X github.com/ihippik/wal-listener/config.Version=${.GIT_VERSION}" \
		./cmd/wal-listener


.PHONY: image
image: wal-listener
	@echo "+ build image"
	docker build \
		-t $(IMAGE_PREFIX)/wal-listener:latest \
		-t $(IMAGE_PREFIX)/wal-listener:${.GIT_VERSION} \
		-f Dockerfile .


.PHONY: push
push: 
	docker push $(IMAGE_PREFIX)/wal-listener:latest
	docker push $(IMAGE_PREFIX)/wal-listener:${.GIT_VERSION}