.PHONY: build mod proto deepCopy docker-clean

HARBOR_DOMAIN := $(shell echo ${HARBOR})
PROJECT := lunara-common
LOGIC_IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/lllidan-logic:latest"
GATEWAY_IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/lllidan-gateway:latest"

build:
	-i docker image rm $(LOGIC_IMAGE)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lllidan-logic cmd/logic/main.go
	cp cmd/logic/Dockerfile . && docker build -t $(LOGIC_IMAGE) .
	rm -f Dockerfile && rm -f lllidan-logic
	docker push $(LOGIC_IMAGE)

build-ga:
	-i docker image rm $(GATEWAY_IMAGE)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lllidan-gateway cmd/gateway/main.go
	cp cmd/gateway/Dockerfile . && docker build -t $(GATEWAY_IMAGE) .
	rm -f Dockerfile && rm -f lllidan-gateway
	docker push $(GATEWAY_IMAGE)

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

gateway:
	go run cmd/gateway/main.go -configPath=./cmd/gateway/conf.yaml -alsologtostderr=true -v=4