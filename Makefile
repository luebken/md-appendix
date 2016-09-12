export GOPATH=$(PWD)

run:
	go run src/github.com/luebken/md-appendix/main.go test-files

install:
	go get ./...
	go install github.com/luebken/md-appendix

md-appendix:
	bin/md-appendix
