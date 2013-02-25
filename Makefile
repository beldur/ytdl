export GOPATH=$(PWD)

all:
	@echo "Current GOPATH: " $$GOPATH
	go install -v ytgifcreator

run:
	dev_appserver.py --use_sqlite src/ytgifcreator

rpc:
	go run src/backend/backend.go -port 8081

deploy:
	appcfg.py update src/ytgifcreator/
