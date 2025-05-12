## 概述
- 后端框架: 
  - 系统整体基于GOFRAME2.x,GOFRAME文档地址: https://goframe.org/docs/core, 选用GOFRAME的主要原因如下:
    - 该框架具有从构建到代码组织等全流程规范,适合团队协作. 
    - 该框架具有很多开箱即用的底层组件,可以快速构建系统,让开发者可以专注于业务逻辑的实现.
- 中间件
  - mysql8.4.5
    - `docker run --name mysql845 -v /home/mysql845/data:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=16IvYBW1UFMtFWR -e MYSQL_DATABASE=goose-test  -d mysql:8.4.5`
  - redis 7.2.8
    - server 
    - `docker run --name redis728 \
    -p 6379:6379 \
    -v /home/redis728/data:/data \
    -d redis:7.2.8 \
    --save 60 1`
    - client 
     - `docker exec -it redis728 redis-cli`

  - kafka 4.x
  - 后续根据需求如果需要增加其他中间件,必须论证必要性,否则,一律不能增加
  - 无论是框架还是中间件,当有大的版本更新时,为了性能和安全性,会选择合适时间点进行升级
- 原则
  - 所有的规范和最佳实践都是为了提升开发效率,避免代码膨胀之后难以维护,所以,必须严格遵守规范
- 规范: 
  - [golang代码规范和最佳实践](docs/golang代码规范和最佳实践.md)
  - [gf-代码最佳实践](docs/gf代码最佳实践.md)
  - [错误处理规范](docs/错误处理规范.md)
  - [数据库规范](docs/数据库规范.md)
  - [git规范](docs/git规范.md)
  - [CI CD流程](docs/CICD流程.md)
  - [日志规范](docs/日志规范.md)
  - [时间处理](docs/时间处理.md)
  - [单元测试规范](docs/单元测试规范.md)
  - [链路追踪和指标](docs/链路追踪和指标.md)
- 需求
  - https://klteeqb691.feishu.cn/drive/folder/FYvmf3GZ0leaxvdKK8Yck2wHnQf
  
## 其他注意事项(严格遵守)
1. 每个开发必须具有独立的本地环境,开发完成后,必须自己在本地测试没有问题,才能进行PR
