GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
#GOLINT=golangci-lint run -E godot -E unparam -E unconvert -E golint -E stylecheck -E gocritic
BIN_NAME=veyil

LINTCMD=golangci-lint run
LINTOPTS=-E gocritic \
	 -E godot \
	 -E golint \
	 -E stylecheck \
	 -E unconvert \
	 -E unparam
GOLINT=$(LINTCMD) $(LINTOPTS)


all: test build

build:
	$(GOBUILD)

test:
	$(GOTEST) -v

lint:
	$(GOLINT)

clean:
	$(GOCLEAN)
	rm -f $(BIN_NAME)
