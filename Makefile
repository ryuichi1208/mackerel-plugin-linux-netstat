BINDIR := bin
NAME := mackerel-plugin-linux-tcp-memory
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)' -s -w

.PHONY:
deps:
	go mod tidy

.PHONY: devel-deps
devel-deps: deps
	go install golang.org/x/lint/golint@latest
	go install github.com/x-motemen/gobump/cmd/gobump@latest
	go install github.com/Songmu/make2help/cmd/make2help@latest


.PHONY: lint
lint:
	#golint .
	go vet ./...

.PHONY: build
build: lint
	go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(NAME)

.PHONY: clean
clean:
	rm -f ./$(BINDIR)/$(NAME)
