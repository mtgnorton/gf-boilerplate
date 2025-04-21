ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING=root:secret@tcp(127.0.0.1:3308)/gf-boilerplate?charset=utf8mb4
export GOOSE_MIGRATION_DIR=./manifest/migration


include ./hack/hack.mk

run-local:
	gf run main.go

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
	gf gen ctrl -s ./apibackend -d ./internal/controller/backend -m -v

ctrl-frontend:
	gf gen ctrl -s ./api/frontend -d ./internal/controller/frontend -m


