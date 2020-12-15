.PHONY: build mod proto deepCopy docker-clean

HARBOR_DOMAIN := $(shell echo ${HARBOR})
PROJECT := lunara-common
REGISTER_IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/lllidan-register:latest"

mod:
	go mod download
	go mod tidy

proto:
	cd scripts && GENS=api bash ./gen.sh

deepCopy:
	go mod vendor
	cd scripts && GENS=deepcopy bash ./gen.sh
	rm -rf vendor

docker-clean:
	echo $(shell docker images | grep lllidan | grep none | awk '{print $3}' | xargs docker image rm -f)

logic:
	go run cmd/logic/main.go -configPath=./cmd/logic/conf.yaml -alsologtostderr=true -v=4