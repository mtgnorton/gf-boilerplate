## 此文件包含所有可用的配置选项
## 及其默认值。

# 参见 https://github.com/golangci/golangci-lint#config-file
# 参见 https://golangci-lint.run/usage/configuration/

# 分析运行的选项。
run:
  # 分析超时时间，例如 30s, 5m。
  # 默认值: 1m
  timeout: 5m
  # 当至少发现一个问题时的退出代码。
  # 默认值: 1
  issues-exit-code: 6
  # 是否包含测试文件。
  # 默认值: true
  tests: false
  # 所有 linter 使用的构建标签列表。
  # 默认值: []
  build-tags: []
  # modules-download-mode 是 GolangCI 的一个配置参数，用于控制 Go 模块的下载行为。
  # 它主要用于在 CI/CD 环境中优化模块下载的性能和稳定性。
  # readonly：只读模式，不修改模块缓存。在 CI/CD 环境中，使用 readonly 模式可以避免重复下载模块，提高构建速度。

  modules-download-mode: readonly
  # 允许多个 golangci-lint 实例并行运行。
  # 如果为 false，golangci-lint 在启动时获取文件锁。
  # 默认值: false
  allow-parallel-runners: true
  # 允许多个 golangci-lint 实例运行，但围绕锁进行序列化。
  # 如果为 false，当 golangci-lint 在启动时无法获取文件锁时，它会以错误退出。
  # 默认值: false
  allow-serial-runners: true
  # 定义 Go 版本限制。
  # 主要与 go1.18 以来的泛型支持相关。
  # 默认值: 使用 go.mod 文件中的 Go 版本，回退到环境变量 `GOVERSION`，再回退到 1.17
  go: '1.23'
  # 可以同时执行 golangci-lint 的操作系统线程数（`GOMAXPROCS`）。
  # 如果明确设置为 0（即非默认值），则 golangci-lint 将自动设置该值以匹配 Linux 容器 CPU 配额。
  # 默认值: 机器中的逻辑 CPU 数量
  concurrency: 4


# 主要 linter 配置。
# 参见 https://golangci-lint.run/usage/linters
linters:
  # 禁用所有默认启用的 linter。
  disable-all: true
  # 自定义启用我们想要使用的 linter。
  enable:
    - errcheck      # Errcheck 是一个用于检查 go 程序中未检查错误的程序。
    - errchkjson    # 检查传递给 JSON 编码函数的类型。报告不支持的类型，并可选择报告可以省略返回错误检查的情况。
    - funlen        # 用于检测长函数的工具
    - gofmt         # Gofmt 检查代码是否经过 gofmt 格式化。默认情况下，此工具使用 -s 选项运行，以检查代码简化
    - goimports     # 检查导入语句是否按照 'goimport' 命令格式化。在自动修复模式下重新格式化导入。
    - gci           # Gci 控制 Go 包导入顺序，使其始终确定性。
    - goconst       # 查找可以替换为常量的重复字符串
    - gocritic      # 提供检查错误、性能和风格问题的诊断。
    - gosimple      # Go 源代码的 linter，专门用于简化代码
    - govet         # Vet 检查 Go 源代码并报告可疑结构，例如参数与格式字符串不匹配的 Printf 调用
    - misspell      # 在注释中查找常见的英语拼写错误
    - nolintlint    # 报告格式不正确或不充分的 nolint 指令
    - revive        # 快速、可配置、可扩展、灵活且美观的 Go linter。golint 的替代品。
    - staticcheck   # 它是 staticcheck 的一组规则。它与 staticcheck 二进制文件不同。
    - typecheck     # 像 Go 编译器的前端一样，解析和类型检查 Go 代码
    - usestdlibvars # 一个检测可能使用 Go 标准库中变量/常量的 linter。
    - whitespace    # 用于检测前导和尾随空白的工具


issues:
  exclude-rules:
    # 排除测试文件中的 context.Context 检查
    - path: _test\.go
      text: "context.Context should be the first parameter of a function"
      linters:
        - revive
    # 排除测试文件中的导出函数返回未导出类型的问题
    - path: _test\.go
      text: "exported func.*returns unexported type.*which can be annoying to use"
      linters:
        - revive
    # https://github.com/go-critic/go-critic/issues/926
    # 忽略 gocritic 检查工具报告的关于 unnecessaryDefer 的问题。
    # gocritic 是一个 Go 代码检查工具，它会检查代码中是否存在不必要的 defer 语句。
    # 在某些情况下，开发者可能有意使用 defer，即使它看起来不必要（例如为了代码一致性或可读性）。
    - linters:
        - gocritic
      text: "unnecessaryDefer:"


# https://golangci-lint.run/usage/linters
linters-settings:
  # https://golangci-lint.run/usage/linters/#misspell
  misspell:
    locale: US
    ignore-words:
      - cancelled

  goimports:
    # 以逗号分隔的前缀列表，如果设置，检查具有给定前缀的导入路径
    # 是否在第三方包之后分组。
    # 默认值: ""
    local-prefixes: gf-boilerplate
  gci:
    # 要比较的部分配置。如果 `custom-order` 为 `true`，则按照 `sections` 选项的顺序。
    # 部分名称不区分大小写，可能包含括号中的参数。
    # 部分的默认顺序是 `standard > default > custom > blank > dot > alias > localmodule`，
    # 解释如下
    # standard 标准库部分,例如：import "fmt"、import "os"。
    # blank：空白导入部分,例如：import _ "github.com/example/package"。
    # default：默认部分，包含所有无法匹配到其他部分的导入语句.例如：一些第三方库的导入语句。
    # dot：点导入部分例如：import . "github.com/example/package"。
    # prefix(github.com/gogf/gf)：自定义部分，例如：import "github.com/gogf/gf/v2"。
    # localmodule: 本地库, 例如 "gf-boilerplate/internal/cmd"

    #   假设有以下导入语句：
    #
    # import (
    #     "fmt"
    #     "os"
    #     _ "github.com/example/package"
    #     . "github.com/example/dotpackage"
    #     "github.com/gogf/gf/v2"
    #     "github.com/thirdparty/package"
    # )
    # 经过 gci 处理后，导入语句会按照以下顺序排列：
    #
    # import (
    #     "fmt"
    #     "os"
    #     "github.com/thirdparty/package"
    #     "github.com/gogf/gf/v2"
    #     _ "github.com/example/package"
    #     . "github.com/example/dotpackage"
    # )

    sections:
      - standard 
      - blank 
      - default 
      - dot 
      # - alias 
      - prefix(github.com/gogf/gf) # 自定义部分：将所有具有指定前缀的导入分组。
      - localmodule # 本地模块部分：包含所有本地包。除非明确启用，否则此部分不存在。

    # 跳过生成的文件。
    # 默认值: true
    skip-generated: true
    # 启用部分的自定义顺序。
    # 如果为 `true`，使部分顺序与 `sections` 的顺序相同。
    # 默认值: false
    custom-order: true
    # 为自定义部分取消词法排序。
    # 默认值: false
    no-lex-order: false
  # https://golangci-lint.run/usage/linters/#revive
  # https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md
  revive:
    ignore-generated-header: true
    severity: error
    rules:
      - name: atomic
      - name: line-length-limit
        severity: error
        arguments: [ 380 ]
      - name: unhandled-error
        severity: warning
        disabled: true
        arguments: []
      - name: var-naming
        severity: warning
        disabled: true
        arguments:
          # AllowList
          - [ "ID","URL","IP","HTTP","JSON","API","UID","Id","Api","Uid","Http","Json","Ip","Url" ]
          # DenyList
          - [ "VM" ]
      - name: string-format
        severity: warning
        disabled: false
        arguments:
          - - 'core.WriteError[1].Message'
            - '/^([^A-Z]|$)/'
            - must not start with a capital letter
          - - 'fmt.Errorf[0]'
            - '/(^|[^\.!?])$/'
            - must not end in punctuation
          - - panic
            - '/^[^\n]*$/'
            - must not contain line breaks
      - name: function-result-limit
        severity: warning
        disabled: false
        arguments: [ 4 ]

  # https://golangci-lint.run/usage/linters/#funlen
  funlen:
    # 检查函数中的行数。
    # 如果小于 0，禁用检查。
    # 默认值: 60
    lines: 340
    # 检查函数中的语句数（即代码语句的数量，如赋值、条件判断、循环等）。
    # 如果小于 0，禁用检查。
    # 默认值: 40
    statements: -1

  # https://golangci-lint.run/usage/linters/#goconst
  goconst:
    # 字符串常量的最小长度。
    # 默认值: 3
    min-len: 4
    # 触发问题的常量字符串出现的最小次数。
    # 默认值: 3
    # 为了后续优化，该值被减少。
    min-occurrences: 30
    # 忽略测试文件。
    # 默认值: false
    ignore-tests: true
    # 查找与值匹配的现有常量。
    # 默认值: true
    match-constant: false
    # 也搜索重复的数字。
    # 默认值: false
    numbers: true
    # 最小值，仅适用于 goconst.numbers
    # 默认值: 3
    min: 5
    # 最大值，仅适用于 goconst.numbers
    # 默认值: 3
    max: 20
    # 当常量不用作函数参数时忽略。
    # 默认值: true
    ignore-calls: false

  # https://golangci-lint.run/usage/linters/#gocritic
  gocritic:
    disabled-checks:
      - ifElseChain
      - assignOp
      - appendAssign
      - singleCaseSwitch
      - regexpMust
      - typeSwitchVar
      - elseif

  # https://golangci-lint.run/usage/linters/#gosimple
  gosimple:
    # Sxxxx 检查 https://staticcheck.io/docs/configuration/options/#checks
    # 默认值: ["*"]
    checks: [
      "all", "-S1000", "-S1001", "-S1002", "-S1008", "-S1009", "-S1016", "-S1023", "-S1025", "-S1029", "-S1034", "-S1040"
    ]

  # https://golangci-lint.run/usage/linters/#govet
  govet:
    # 报告关于遮蔽变量的信息。
    # 默认值: false
    # check-shadowing: true
    # 每个分析器的设置。
    settings:
      # 分析器名称，运行 `go tool vet help` 查看所有分析器。
      printf:
        # 要检查的打印函数名称的逗号分隔列表（除了默认值，请参见 `go tool vet help printf`）。
        # 默认值: []
        funcs:
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Infof
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Warnf
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Errorf
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Fatalf
        # shadow:
        # 是否严格对待遮蔽；可能会产生噪音。
        # 默认值: false
        # strict: false
      unusedresult:
        # 必须使用其结果的函数的逗号分隔列表
        # （除了默认值 context.WithCancel,context.WithDeadline,context.WithTimeout,context.WithValue,
        # errors.New,fmt.Errorf,fmt.Sprint,fmt.Sprintf,sort.Reverse）
        # 默认值 []
        funcs:
          - pkg.MyFunc
          - context.WithCancel
        # 类型为 func() string 的方法的名称的逗号分隔列表，其结果必须使用
        # （除了默认值 Error,String）
        # 默认值 []
        stringmethods:
          - MyMethod
    # 启用所有分析器。
    # 默认值: false
    enable-all: true
    # 按名称禁用分析器。
    # 运行 `go tool vet help` 查看所有分析器。
    # 默认值: []
    disable:
      - asmdecl
      - assign
      - atomic
      - atomicalign
      - bools
      - buildtag
      - cgocall
      - composites
      - copylocks
      - deepequalerrors
      - errorsas
      - fieldalignment
      - findcall
      - framepointer
      - httpresponse
      - ifaceassert
      - loopclosure
      - lostcancel
      - nilfunc
      - nilness
      - reflectvaluecompare
      - shift
      - shadow
      - sigchanyzer
      - sortslice
      - stdmethods
      - stringintconv
      - structtag
      - testinggoroutine
      - tests
      - unmarshal
      - unreachable
      - unsafeptr
      - unusedwrite

  # https://golangci-lint.run/usage/linters/#staticcheck
  staticcheck:
    # SAxxxx 检查 https://staticcheck.io/docs/configuration/options/#checks
    # 默认值: ["*"]
    checks: [ "all","-SA1019","-SA4015","-SA1029","-SA1016","-SA9003","-SA4006","-SA6003" ]