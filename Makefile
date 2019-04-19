.PHONY: build run

all: bindata build

build:
	env GOOS=darwin GOARCH=amd64 go build -o k-peach -ldflags "-X main.GitCommit=$(shell git rev-list -1 HEAD)"; 
	cp  k-peach /Users/tqcenglish/Applications/bin/
run:
	cd my.peach && ../k-peach web

bindata:
	go-bindata -o=pkg/bindata/bindata.go -ignore="\\.DS_Store|config.codekit|.less" -pkg=bindata templates/... conf/... public/... docs/...

release:
	gox -osarch="linux/arm linux/amd64" -ldflags "-X main.GitCommit=$(shell git rev-list -1 HEAD)";
