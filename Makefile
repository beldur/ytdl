export GOPATH := $(CURDIR)

all:
	@echo $$GOPATH
	go build -v ytdl
