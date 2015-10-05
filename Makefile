VERSION := $(shell git describe --long --tags)
TAG := $(shell git describe --tags)
INSTALL := go install -ldflags "-X main.version=$(VERSION)" ./...
NAME := git-get
BINARY := $(NAME)-$(shell uname -s)-$(shell uname -m)
CHECKSUM := $(BINARY).sha256

clean:
	rm -f $(NAME) $(BINARY) $(CHECKSUM)

install:
	$(INSTALL)

$(NAME):
	GOBIN=$(CURDIR) $(INSTALL)

binary: $(BINARY)
$(BINARY): $(NAME)
	mv $< $@

checksum: $(CHECKSUM)
$(CHECKSUM): $(BINARY)
	shasum -p -a 256 $< > $@

release: $(BINARY) $(CHECKSUM)
	hub release create -a $(BINARY) -m $(TAG) -a $(CHECKSUM) $(TAG)

test:
	go test -v ./...
