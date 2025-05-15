## 流水线检查
![](images/CICD流程/20250514194552.png)
![](images/CICD流程/20250514194615.png)
### 静态检查
- golangci-lint 静态检查 详细配置见 .golangci.yml
    - 使用golangci-lint的版本为v2.0.2
    - 本地安装
        - 参考https://golangci-lint.run/welcome/install/,建议使用`go install`的方式,命令`go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2`
    - 本地检查有两种方式
        - 可选方式: 结合编辑器集成,如自动保存时进行检查,集成到编辑器参考https://golangci-lint.run/welcome/integrations/
        - 必选方式: git hook检查,在commit时进行检查
    - 常用命令:
        - 运行所有检查: `golangci-lint run` 或 `make lint`
        - 运行单一检查: `golangci-lint run --default=none -E errcheck`
        - 修复: `golangci-lint run --fix` 或 `make lint-fix`
    
    - 绕过检查:
        - 对于一些特殊情况,需要绕过检查,严禁滥用
            - all: 忽略该行所有检查
            - 特定检查器,如unused: 忽略特定检查
            - 注意//和nolint之间不能有空格
            - 例如:
                - 忽略该行的检查,放在行尾:
                ```go
                    var bad_name int //nolint:unused
                ```
                - 忽略代码块的检查,在行的开始添加:
                ```go
                    /*nolint:all*/
                    var bad_name int
                ```
                - 忽略文件的检查,放在文件的开始:
                ```go
                    //nolint:all
                    package main
                ```
### 人工review
- 无论是master分支,还是dev分支,所有的提交都通过PR的方式,先经过本地静态检查没有问题后,开发人员提交PR,经过ci检查没有问题后,需要有至少一名其他成员进行review,review通过后,合并到目标分支,关于PR的使用,参考[git相关.md](git规范.md)

## 部署
- 配置采用配置文件+环境变量的方式进行部署,部分敏感配置通过环境变量来指定,所有具有WLINK_前缀的环境变量,都会对配置文件进行覆盖
- 当合并到目标分支后,会自动触发代码部署和sql部署
- 以下配置通过环境变量指定
    - WLINK_DATABASE_DEFAULT_LINK -> database.default.link
    - WLINK_REDIS_DEFAULT_ADDRESS -> redis.default.address
    - WLINK_REDIS_DEFAULT_PASS ->  redis.default.pass
    - WLINK_REDIS_CACHE_ADDRESS -> redis.cache.address
    - WLINK_REDIS_CACHE_PASS -> redis.cache.pass
- 参见[config.local.env](../manifest/config/config.local.env)


