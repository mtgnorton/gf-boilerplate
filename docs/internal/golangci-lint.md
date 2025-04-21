## staickcheck参考配置
```yaml
linters:
  settings:
    staticcheck:
      # 允许的点导入（. import）白名单
      # 默认值: ["github.com/mmcloughlin/avo/build", "github.com/mmcloughlin/avo/operand", "github.com/mmcloughlin/avo/reg"]
      dot-import-whitelist:
        - fmt

      # 允许使用的缩写词（全大写标识符）
      # 默认值: ["ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "QPS", "RAM", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "GID", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS", "SIP", "RTP", "AMQP", "DB", "TS"]
      initialisms: ["ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "QPS", "RAM", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "GID", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS", "SIP", "RTP", "AMQP", "DB", "TS"]

      # 允许的HTTP状态码白名单
      # 默认值: ["200", "400", "404", "500"]
      http-status-code-whitelist: ["200", "400", "404", "500"]

      # 启用的检查规则（SAxxxx/STxxxx/Sxxxx等）
      # 示例（禁用某些检查）: [ "all", "-SA1000", "-SA1001"]
      # 默认值: ["all", "-ST1000", "-ST1003", "-ST1016", "-ST1020", "-ST1021", "-ST1022"]
      checks:
        # SA1000: 无效的正则表达式
        - SA1000
        # SA1001: 无效的模板
        - SA1001
        # SA1002: 'time.Parse'中使用无效的格式
        - SA1002
        # SA1003: 'encoding/binary'函数的不支持参数
        - SA1003
        # SA1004: 'time.Sleep'中使用可疑的小型无类型常量
        - SA1004
        # SA1005: 'exec.Command'的第一个参数无效
        - SA1005
        # SA1006: 'Printf'第一个参数是动态的且没有其他参数
        - SA1006
        # SA1007: 'net/url.Parse'中的无效URL
        - SA1007
        # SA1008: 'http.Header'映射中的非规范键
        - SA1008
        # SA1010: '(*regexp.Regexp).FindAll'调用时'n == 0'，这将始终返回零结果
        - SA1010
        # SA1011: "strings"包中的各种方法需要有效的UTF-8，但提供了无效输入
        - SA1011
        # SA1012: 传递了nil的'context.Context'，应考虑使用'context.TODO'
        - SA1012
        # SA1013: 'io.Seeker.Seek'的第一个参数是whence常量，但它应该是第二个参数
        - SA1013
        # SA1014: 向'Unmarshal'或'Decode'传递了非指针值
        - SA1014
        # SA1015: 以会导致泄漏的方式使用'time.Tick'。应考虑使用'time.NewTicker'，仅在测试、命令和无尽函数中使用'time.Tick'
        - SA1015
        # SA1016: 捕获无法捕获的信号
        - SA1016
        # SA1017: 与'os/signal.Notify'一起使用的通道应该是缓冲的
        - SA1017
        # SA1018: 以'n == 0'调用'strings.Replace'，这不会做任何事情
        - SA1018
        # SA1019: 使用已弃用的函数、变量、常量或字段
        - SA1019
        # SA1020: 与'net.Listen'相关函数一起使用的无效host:port对
        - SA1020
        # SA1021: 使用'bytes.Equal'比较两个'net.IP'
        - SA1021
        # SA1023: 在'io.Writer'实现中修改缓冲区
        - SA1023
        # SA1024: 字符串cutset包含重复字符
        - SA1024
        # SA1025: 无法正确使用'(*time.Timer).Reset'的返回值
        - SA1025
        # SA1026: 无法序列化通道或函数
        - SA1026
        # SA1027: 对64位变量的原子访问必须是64位对齐的
        - SA1027
        # SA1028: 'sort.Slice'只能用于切片
        - SA1028
        # SA1029: 在'context.WithValue'调用中使用不适当的键
        - SA1029
        # SA1030: 'strconv'函数调用中的无效参数
        - SA1030
        # SA1031: 传递给编码器的重叠字节切片
        - SA1031
        # SA1032: 'errors.Is'的参数顺序错误
        - SA1032
        # SA2000: 在goroutine内部调用'sync.WaitGroup.Add'，导致竞态条件
        - SA2000
        # SA2001: 空临界区，您是否想延迟解锁？
        - SA2001
        # SA2002: 在goroutine中调用'testing.T.FailNow'或'SkipNow'，这是不允许的
        - SA2002
        # SA2003: 锁定后立即延迟'Lock'，可能应该延迟'Unlock'
        - SA2003
        # SA3000: 'TestMain'没有调用'os.Exit'，隐藏了测试失败
        - SA3000
        # SA3001: 在基准测试中分配给'b.N'会扭曲结果
        - SA3001
        # SA4000: 二元运算符的两边有相同的表达式
        - SA4000
        # SA4001: '&*x'被简化为'x'，它不会复制'x'
        - SA4001
        # SA4003: 将无符号值与负值进行比较是无意义的
        - SA4003
        # SA4004: 循环无条件地在一次迭代后退出
        - SA4004
        # SA4005: 永远不会被观察到的字段赋值。您是否想使用指针接收器？
        - SA4005
        # SA4006: 分配给变量的值在被覆盖之前从未被读取。忘记的错误检查或死代码？
        - SA4006
        # SA4008: 循环条件中的变量从未改变，您是否递增了错误的变量？
        - SA4008
        # SA4009: 函数参数在其第一次使用之前被覆盖
        - SA4009
        # SA4010: 'append'的结果永远不会被观察到
        - SA4010
        # SA4011: 没有效果的break语句。您是否想跳出外部循环？
        - SA4011
        # SA4012: 将值与NaN比较，尽管没有值等于NaN
        - SA4012
        # SA4013: 双重否定布尔值（'!!b'）等同于'b'。这要么是冗余的，要么是拼写错误
        - SA4013
        # SA4014: if/else if链有重复的条件且没有副作用；如果条件第一次不匹配，第二次也不会匹配
        - SA4014
        # SA4015: 对从整数转换的浮点数调用像'math.Ceil'这样的函数没有用处
        - SA4015
        # SA4016: 某些位操作，如'x ^ 0'，没有用处
        - SA4016
        # SA4017: 丢弃没有副作用的函数的返回值，使调用无意义
        - SA4017
        # SA4018: 变量的自赋值
        - SA4018
        # SA4019: 同一文件中有多个相同的构建约束
        - SA4019
        # SA4020: 类型切换中无法到达的case子句
        - SA4020
        # SA4021: "x = append(y)"等同于"x = y"
        - SA4021
        # SA4022: 将变量的地址与nil比较
        - SA4022
        # SA4023: 接口值与无类型nil的不可能比较
        - SA4023
        # SA4024: 检查内置函数的不可能返回值
        - SA4024
        # SA4025: 字面量的整数除法结果为零
        - SA4025
        # SA4026: Go常量无法表示负零
        - SA4026
        # SA4027: '(*net/url.URL).Query'返回一个副本，修改它不会改变URL
        - SA4027
        # SA4028: 'x % 1'总是为零
        - SA4028
        # SA4029: 尝试排序切片的无效方式
        - SA4029
        # SA4030: 生成随机数的无效尝试
        - SA4030
        # SA4031: 检查永远不会为nil的值是否为nil
        - SA4031
        # SA4032: 将'runtime.GOOS'或'runtime.GOARCH'与不可能的值比较
        - SA4032
        # SA5000: 赋值给nil映射
        - SA5000
        # SA5001: 在检查可能的错误之前延迟'Close'
        - SA5001
        # SA5002: 空for循环（"for {}"）会自旋并可能阻塞调度器
        - SA5002
        # SA5003: 无限循环中的defer永远不会执行
        - SA5003
        # SA5004: "for { select { ..."带有空的default分支会自旋
        - SA5004
        # SA5005: 终结器引用被终结的对象，阻止垃圾回收
        - SA5005
        # SA5007: 无限递归调用
        - SA5007
        # SA5008: 无效的结构体标签
        - SA5008
        # SA5009: 无效的Printf调用
        - SA5009
        # SA5010: 不可能的类型断言
        - SA5010
        # SA5011: 可能的nil指针解引用
        - SA5011
        # SA5012: 向期望偶数大小的函数传递奇数大小的切片
        - SA5012
        # SA6000: 在循环中使用'regexp.Match'或相关函数，应该使用'regexp.Compile'
        - SA6000
        # SA6001: 当通过字节切片索引映射时错过的优化机会
        - SA6001
        # SA6002: 在'sync.Pool'中存储非指针值会分配内存
        - SA6002
        # SA6003: 在遍历之前将字符串转换为rune切片
        - SA6003
        # SA6005: 使用'strings.ToLower'或'strings.ToUpper'进行低效的字符串比较
        - SA6005
        # SA6006: 使用io.WriteString写入'[]byte'
        - SA6006
        # SA9001: 在range循环中的defer可能不会在您期望的时候运行
        - SA9001
        # SA9002: 使用看起来应该是八进制的非八进制'os.FileMode'
        - SA9002
        # SA9003: if或else分支中的空体
        - SA9003
        # SA9004: 只有第一个常量有显式类型
        - SA9004
        # SA9005: 尝试序列化没有公共字段也没有自定义序列化的结构体
        - SA9005
        # SA9006: 对固定大小的整数值进行可疑的位移
        - SA9006
        # SA9007: 删除不应该被删除的目录
        - SA9007
        # SA9008: 类型断言的'else'分支可能没有读取正确的值
        - SA9008
        # SA9009: 无效的Go编译器指令
        - SA9009
        # ST1000: 不正确或缺少包注释
        - ST1000
        # ST1001: 不鼓励使用点导入
        - ST1001
        # ST1003: 选择不当的标识符
        - ST1003
        # ST1005: 错误字符串格式不正确
        - ST1005
        # ST1006: 选择不当的接收器名称
        - ST1006
        # ST1008: 函数的错误值应该是其最后一个返回值
        - ST1008
        # ST1011: 为'time.Duration'类型的变量选择不当的名称
        - ST1011
        # ST1012: 为错误变量选择不当的名称
        - ST1012
        # ST1013: 应该使用常量表示HTTP错误代码，而不是魔术数字
        - ST1013
        # ST1015: switch的default case应该是第一个或最后一个case
        - ST1015
        # ST1016: 使用方法接收器名称保持一致
        - ST1016
        # ST1017: 不要使用尤达条件（Yoda conditions）
        - ST1017
        # ST1018: 避免在字符串字面量中使用零宽度和控制字符
        - ST1018
        # ST1019: 多次导入同一个包
        - ST1019
        # ST1020: 导出函数的文档应该以函数名开头
        - ST1020
        # ST1021: 导出类型的文档应该以类型名开头
        - ST1021
        # ST1022: 导出变量或常量的文档应该以变量名开头
        - ST1022
        # ST1023: 变量声明中的冗余类型
        - ST1023
        # S1000: 使用普通的通道发送或接收而不是单case的select
        - S1000
        # S1001: 用copy调用替换for循环
        - S1001
        # S1002: 省略与布尔常量的比较
        - S1002
        # S1003: 用'strings.Contains'替换'strings.Index'调用
        - S1003
        # S1004: 用'bytes.Equal'替换'bytes.Compare'调用
        - S1004
        # S1005: 删除不必要的空白标识符使用
        - S1005
        # S1006: 使用"for { ... }"表示无限循环
        - S1006
        # S1007: 通过使用原始字符串字面量简化正则表达式
        - S1007
        # S1008: 简化返回布尔表达式
        - S1008
        # S1009: 省略对切片、映射和通道的冗余nil检查
        - S1009
        # S1010: 省略默认的切片索引
        - S1010
        # S1011: 使用单个'append'连接两个切片
        - S1011
        # S1012: 用'time.Since(x)'替换'time.Now().Sub(x)'
        - S1012
        # S1016: 使用类型转换而不是手动复制结构体字段
        - S1016
        # S1017: 用'strings.TrimPrefix'替换手动修剪
        - S1017
        # S1018: 使用"copy"滑动元素
        - S1018
        # S1019: 通过省略冗余参数简化"make"调用
        - S1019
        # S1020: 在类型断言中省略冗余的nil检查
        - S1020
        # S1021: 合并变量声明和赋值
        - S1021
        # S1023: 省略冗余的控制流
        - S1023
        # S1024: 用'time.Until(x)'替换'x.Sub(time.Now())'
        - S1024
        # S1025: 不必要地使用'fmt.Sprintf("%s", x)'
        - S1025
        # S1028: 用'fmt.Errorf'简化错误构造
        - S1028
        # S1029: 直接遍历字符串
        - S1029
        # S1030: 使用'bytes.Buffer.String'或'bytes.Buffer.Bytes'
        - S1030
        # S1031: 省略循环周围的冗余nil检查
        - S1031
        # S1032: 使用'sort.Ints(x)'、'sort.Float64s(x)'和'sort.Strings(x)'
        - S1032
        # S1033: 围绕"delete"调用的不必要的保护
        - S1033
        # S1034: 使用类型断言的结果简化case
        - S1034
        # S1035: 在'net/http.Header'的方法调用中冗余调用'net/http.CanonicalHeaderKey'
        - S1035
        # S1036: 围绕映射访问的不必要的保护
        - S1036
        # S1037: 复杂的休眠方式
        - S1037
        # S1038: 打印格式化字符串的不必要复杂方式
        - S1038
        # S1039: 不必要地使用'fmt.Sprint'
        - S1039
        # S1040: 断言为当前类型的类型断言
        - S1040
        # QF1001: 应用德摩根定律
        - QF1001
        # QF1002: 

```


## formatters-gci
```yaml
formatters:
  settings:
    gci:
      # 用于分组的导入区域配置（区域名称不区分大小写，可包含括号参数）。
      # 默认区域顺序为：`standard > default > custom > blank > dot > alias > localmodule`。
      # 若 `custom-order` 为 `true`，则完全按照 `sections` 列表的顺序排列。
      # 默认值: ["standard", "default"]
      sections:
        - standard    # 标准区域：包含所有标准库（如 "fmt"、"net/http"）。
        - default     # 默认区域：未匹配到其他区域的导入项。
        - prefix(github.com/org/project)  # 自定义区域：匹配指定前缀的导入（如项目自有包）。
        - blank       # 空导入区域：如 `_ "package"`。需显式启用才会生效。
        - dot         # 点导入区域：如 `. "package"`。需显式启用才会生效。
        - alias       # 别名导入区域：如 `alias "package"`。需显式启用才会生效。
        - localmodule # 本地模块区域：本地包的导入。需显式启用才会生效。

      # 禁止行内注释（即导入语句同一行的注释）。
      # 默认值: false
      no-inline-comments: true

      # 禁止前缀注释（即导入语句上方的独立注释行）。
      # 默认值: false
      no-prefix-comments: true

      # 启用自定义区域顺序。
      # 若为 `true`，则严格按 `sections` 列表的顺序排列区域。
      # 默认值: false
      custom-order: true

      # 禁用自定义区域的字典序排序（按代码书写顺序保留）。
      # 默认值: false
      no-lex-order: true

```