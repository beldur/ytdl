export GOPATH=$(PWD)

all:
	@echo "Current GOPATH: " $$GOPATH
	go build -v ytdl
