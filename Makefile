export GOPATH=$(PWD)

all:
	@echo "Current GOPATH: " $$GOPATH
	go install -v ytgifcreator

run:
	dev_appserver.py --use_sqlite src/ytgifcreator

travis:
	vendor/google_appengine/dev_appserver.py --skip_sdk_update_check --use_sqlite src/ytgifcreator

rpc:
	@echo "Current GOPATH: " $$GOPATH
	go run src/backend/backend.go -port 8081

deploy:
	appcfg.py update src/ytgifcreator/
