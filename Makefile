export GOPATH=$(PWD)

all:
	@echo "Current GOPATH: " $$GOPATH
	go install -v ytgifcreator

run:
	dev_appserver.py --use_sqlite src/ytgifcreator
