## api 接口和控制器生成
  - 接口文件路径:
      - `/api/模块/子模块/v1/user.go`
      - 如`/api/admin/user/v1/user.go`

  - 接口文件结构:
    - 接口命名:
      - api 定义的结构体名称需要满足 `操作+Req` 及 `操作+Res` 的命名方式规范,这里的操作是当前 API 模块的操作名称，通常对应 CRUD 是： Create、 Read、 Update、 Delete。
      
      | 操作名称 | 建议命名 | 方法 | 备注 |
      | ------- | ------- | ---- | ---- |
      | 查询列表 | GetListReq/Res | GET | 通常是从数据库中分页查询数据记录 |
      | 查询详情 | GetOneReq/Res | GET | 通常接口需要传递主键条件，从数据库中查询记录详情 |
      | 创建资源 | CreateReq/Res | POST | 通常是往数据表中插入一条或多条数据记录 |
      | 修改资源 | UpdateReq/Res | PUT | 通常是按照一定条件修改数据表中的一条或多条数据记录 |
      | 删除资源 | DeleteReq/Res | DELETE | 通常是按照一定条件删除数据表中的一条或多条数据记录 |

    - 接口文件示例:
      ```go
      // 创建用户
      type CreateReq struct {
          g.Meta   `path:"/user" tags:"用户" method:"post" summary:"创建用户"`
          Username string `v:"required#请输入用户名" dc:"用户名"`
          Password string `v:"required#请输入密码" dc:"密码"`
      }
      type CreateRes struct {
          Id int64 `json:"id" dc:"用户ID"`
      } 
      ```
 - 生成命令
    `gf gen ctrl` 或 `make ctrl`
## dao 生成
  - 首先配置 `hack/config.yaml`下的`gfcli`选项
  - 数据表存放路径:
    - `manifest/migration/`
  - 在本地执行`make migrate` 生成数据表结构
  - 生成命令
    `gf gen dao` 或 `make dao`
  - 生成文件
    - `/internal/dao/internal/user.go` 封装对数据表user crud的访问,**不可修改**
    - `/internal/model/do` 执行查询时业务模型到数据模型的转换,大多字段都为interface{},**不可修改**
    - `/internal/model/entity` 数据模型,数据模型即与数据表一一对应的数据结构,**不可修改**
    - `/internal/dao/user.go` 对`/internal/dao/internal/user.go`的简单封装,用于其他模块调用,**可修改**

  - 自定义业务模型
    - 存放在`/internal/model/` 目录下