ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING=root:secret@tcp(127.0.0.1:3308)/gf-boilerplate?charset=utf8mb4
export GOOSE_MIGRATION_DIR=./manifest/migration


# include ./hack/hack.mk

goose-up:
	goose up

goose-down:
	goose down

timestr:
	@echo `date +"%Y%m%d%H%M%S"`

lint:
	golangci-lint run main.go

lint-fix:
	golangci-lint run --fix main.go