NAME = {{name}}

TAG=$(shell git describe --tags `git rev-list --tags --max-count=1` 2>/dev/null)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse --short HEAD)

# VERSION is the latest tag, or branch name if no tags
VERSION=$(if $(TAG),$(TAG),$(BRANCH))

BUILDTIME = $(shell date '+%s')
RELEASES_DIR = releases
BUILD_OPTS = -ldflags "-s -w -X main.VERSION=$(VERSION) -X main.COMMIT=$(COMMIT) -X main.BUILDTIME=$(BUILDTIME)"

clean:
	rm -rf $(RELEASES_DIR)

$(RELEASES_DIR)/$(NAME):
	CGO_ENABLED=0 go build -o $@ $(BUILD_OPTS) cmd/main.go

$(RELEASES_DIR)/$(NAME)-%:
	GOOS=$* CGO_ENABLED=0 GOARCH=amd64 go build -o $@ $(BUILD_OPTS) cmd/main.go

.PHONY: docker
docker:
	docker build . -t $(NAME) -t $(NAME):$(VERSION)-$(COMMIT) -t $(NAME):$(VERSION) 