NAME = {{name}}

BUILDTIME = $(shell date '+%s')
RELEASES_DIR = releases
BUILD_OPTS = -ldflags "-s -w -X main.VERSION=$(VERSION) -X main.COMMIT=$(COMMIT) -X main.BUILDTIME=$(BUILDTIME)"

clean:
	rm -rf $(RELEASES_DIR)

$(RELEASES_DIR)/$(NAME):
	CGO_ENABLED=0 go build $(BUILD_OPTS) -o $@ cmd/	

$(RELEASES_DIR)/$(NAME)-%:
	GOOS=$* CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_OPTS) -o $@ cmd/

.PHONY: docker
docker
	docker build . -t $(NAME) -t $(NAME):$(VERSION)-$(COMMIT) -t $(NAME):$(VERSION) 