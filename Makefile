export GOPATH=$(PWD)

run:
	go run src/github.com/luebken/md-appendix/main.go test-files
	#go run src/github.com/luebken/md-appendix/main.go /Users/mdl/workspace/github/luebken/container-patterns


install:
	go get ./...
	go install github.com/luebken/md-appendix

md-appendix:
	bin/md-appendix
