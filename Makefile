.PHONY: build mod proto docker-clean

HARBOR_DOMAIN := $(shell echo ${HARBOR})
PROJECT := lunara-common
REGISTER_IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/register:latest"

mod:
	go mod download
	go mod tidy

proto:
	cd scripts && GENS=api bash ./gen.sh

docker-clean:
	echo $(shell docker images | grep none | awk '{print $3}' | xargs docker image rm -f)
