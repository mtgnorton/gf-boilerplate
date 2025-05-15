ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"


include ./manifest/config/config.local.env


include ./hack/hack.mk

# 本地运行,热加载,air 安装 https://github.com/air-verse/air
run-local:
	lsof -ti:8002 | xargs kill -9 || true
	air

# 创建迁移文件
# make create-sql-file create_admin_table,会在manifest/migration下生成20250514095445_create_user_11.sql
create-sql-file:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "请提供文件名称"; \
		exit 1; \
	fi
	@touch manifest/migration/`date +"%Y%m%d%H%M%S"_$(filter-out $@,$(MAKECMDGOALS))`.sql

%:
	@:

# 执行数据库迁移,goose 安装 https://github.com/pressly/goose
migrate:
	goose up

# 回滚数据库迁移
rollback:
	goose down

# 运行代码检查,golangci-lint 安装 https://golangci-lint.run/welcome/install/
lint:
	golangci-lint run ./...

# 运行代码检查并自动修复
lint-fix:
	golangci-lint run --fix ./...

# 生成超管端API控制器代码
ctrl-super:
	gf gen ctrl -s ./apibackend -d ./internal/controller/super  -v

ctrl-backend:
	gf gen ctrl -s ./apibackend -d ./internal/controller/backend  -v



# 生成枚举类型代码
enum:
	gf gen enums -p internal/packed/boot_enums.go -s .

# 运行单元测试
test: 
	go test -v ./...

# 启动jaeger链路追踪
jager:
	sh manifest/deploy/dependent/jager.sh

# 启动prometheus和grafana监控
prometheus-grafana:
	docker-compose -f manifest/deploy/dependent/prometheus_grafana/compose.yaml up