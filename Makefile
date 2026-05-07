.PHONY: all build test coverage vet fmt tidy clean

.DEFAULT_GOAL := all

BINDIR := bin

all: fmt tidy vet test coverage build

fmt:
	@gofmt -s -w .

tidy:
	@go mod tidy

vet:
	@go vet ./...

test:
	@go test ./...

# Self-host: covgate gates its own coverage.
coverage: | $(BINDIR)
	@go test -covermode=set -coverprofile=$(BINDIR)/coverage.tmp.out ./...
	@go run ./cmd/covgate \
		-profile=$(BINDIR)/coverage.tmp.out \
		-out=$(BINDIR)/coverage.out \
		-ignore=.covignore \
		-min=100

build: | $(BINDIR)
	@go build -trimpath -o $(BINDIR)/covgate ./cmd/covgate

$(BINDIR):
	@mkdir -p $(BINDIR)

clean:
	rm -rf $(BINDIR)
