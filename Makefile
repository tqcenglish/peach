.PHONY: build run

all: build run

build:
	env GOOS=darwin GOARCH=amd64 go build -o k-peach;
	cp  k-peach /Users/tqcenglish/Applications/bin/
run:
	./k-peach web

bindata:
	go-bindata -o=pkg/bindata/bindata.go -ignore="\\.DS_Store|README|config.codekit|.less" -pkg=bindata templates/... conf/... public/...

release:
	env GOOS=darwin GOARCH=amd64 go build -o k-peach; tar czf darwin_amd64.tar.gz peach
	env GOOS=linux GOARCH=amd64 go build -o k-peach; tar czf linux_amd64.tar.gz peach
	env GOOS=linux GOARCH=386 go build -o k-peach; tar czf linux_386.tar.gz peach
	env GOOS=linux GOARCH=arm go build -o k-peach; tar czf linux_arm.tar.gz peach
	env GOOS=windows GOARCH=amd64 go build -o k-peach; tar czf windows_amd64.tar.gz peach
	env GOOS=windows GOARCH=386 go build -o k-peach; tar czf windows_386.tar.gz peach
