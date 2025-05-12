ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"


include ./manifest/config/config.local.env


include ./hack/hack.mk

run-local:
	lsof -ti:8002 | xargs kill -9 || true
	air

migrate:
	goose up

rollback:
	goose down

timestr:
	@echo `date +"%Y%m%d%H%M%S"`

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...


ctrl-backend:
	gf gen ctrl -s ./apibackend -d ./internal/controller/backend  -v

ctrl-frontend:
	gf gen ctrl -s ./api/frontend -d ./internal/controller/frontend -m

enum:
	gf gen enums -p internal/packed/boot_enums.go -s .

test: 
	go test -v ./...

# 
jager:
	sh manifest/deploy/dependent/jager.sh

prometheus-grafana:
	docker-compose -f manifest/deploy/dependent/prometheus_grafana/compose.yaml up