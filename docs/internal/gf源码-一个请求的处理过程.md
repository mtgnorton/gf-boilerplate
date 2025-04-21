1. 执行s.Run, 源码位置net/ghttp/ghttp_server.go:445,然后执行s.Start,在Start方法中有如下代码
    ```go
        // Default HTTP handler.
        if s.config.Handler == nil {
            s.config.Handler = s.ServeHTTP
        }
    ```
    s.ServeHttp代码位置:net/ghttp/ghttp_server_handler.go:31
2. ServeHttp处理请求,方法精简如下:
    ```go
    request   = newRequest(s, r, w)    // Create a new request object.

    defer s.handleAfterRequestDone(request)

    s.callHookHandler(HookBeforeServe, request)

    // 根据请求路径,获取对应的处理器
    request.handlers,
            request.serveHandler,
            request.hasHookHandler,
            request.hasServeHandler = s.getHandlersWithCache(request)


    // HOOK - BeforeServe
    s.callHookHandler(HookBeforeServe, request)

    if len(request.handlers) > 0 {
            // Dynamic service.
            request.Middleware.Next()
    }
    // HOOK - AfterServe
    if !request.IsExited() {
        s.callHookHandler(HookAfterServe, request)
    }

    // HOOK - BeforeOutput
    if !request.IsExited() {
        s.callHookHandler(HookBeforeOutput, request)
    }

    // Response handling.
    s.handleResponse(request, sessionId)

    // HOOK - AfterOutput
    if !request.IsExited() {
        s.callHookHandler(HookAfterOutput, request)
    }
    ```


3. 关键结构体Request字段如下:
    ```go
    type Request struct {
        *http.Request                  // 内嵌标准库的HTTP请求对象
        Server     *Server             // 关联的服务器实例
        Cookie     *Cookie             // Cookie管理对象
        Session    *gsession.Session   // 会话管理对象
        Response   *Response           // 当前请求对应的响应对象
        Router     *Router             // 匹配到的路由对象（注意：在HOOK处理器中不可用）
        EnterTime  *gtime.Time         // 请求开始时间（毫秒级）
        LeaveTime  *gtime.Time         // 请求结束时间（毫秒级）
        Middleware *middleware         // 中间件管理器
        StaticFile *staticFile         // 静态文件服务对象

        // ===================================================================
        // 以下为内部使用的私有属性
        // ===================================================================

        handlers        []*HandlerItemParsed   // 当前请求匹配的所有处理器（包含业务处理器、钩子和中间件）
        serveHandler    *HandlerItemParsed     // 实际执行业务逻辑的处理器（非钩子或中间件）
        handlerResponse interface{}            // 处理器返回的响应对象
        hasHookHandler  bool                   // 标记是否存在钩子处理器（性能优化用）
        hasServeHandler bool                   // 标记是否存在业务处理器（性能优化用）
        parsedQuery     bool                   // 标记是否已解析GET参数
        parsedBody      bool                   // 标记是否已解析请求体
        parsedForm      bool                   // 标记是否已解析表单数据（PUT/POST/PATCH方法）
        paramsMap       map[string]interface{} // 自定义参数映射表
        routerMap       map[string]string      // 路由参数映射表（无路由参数时为nil）
        queryMap        map[string]interface{} // URL查询参数映射表（无查询字符串时为nil）
        formMap         map[string]interface{} // 表单参数映射表（无表单数据时为nil）
        bodyMap         map[string]interface{} // 请求体参数映射表（无内容时为nil）
        error           error                  // 当前请求执行过程中的错误
        exitAll         bool                   // 标记当前请求是否已完全退出
        parsedHost      string                 // 解析后的主机名（供GetHost方法使用）
        clientIp        string                 // 解析后的客户端IP（供GetClientIp方法使用）
        bodyContent     []byte                 // 原始请求体内容
        isFileRequest   bool                   // 标记当前是否为文件请求
        viewObject      *gview.View            // 响应专用的模板引擎对象
        viewParams      gview.Params           // 响应专用的模板变量
        originUrlPath   string                 // 客户端传递的原始URL路径
    }
    ```

4. 中间件调用和具体的业务handler通过`request.Middleware.Next()`调用
    ```go
    	var item *HandlerItemParsed
    	var loop = true

    	for loop {
		// Check whether the request is excited.
		if m.request.IsExited() || m.handlerIndex >= len(m.request.handlers) {
			break
		}
		item = m.request.handlers[m.handlerIndex]
		// Filter the HOOK handlers, which are designed to be called in another standalone procedure.
		if item.Handler.Type == HandlerTypeHook {
			m.handlerIndex++
			continue
		}
		var ctx = m.request.Context()
		gutil.TryCatch(ctx, func(ctx context.Context) {
			// Execute bound middleware array of the item if it's not empty.
			if m.handlerMDIndex < len(item.Handler.Middleware) {
				md := item.Handler.Middleware[m.handlerMDIndex]
				m.handlerMDIndex++
				niceCallFunc(func() {
					md(m.request)
				})
				loop = false
				return
			}
			m.handlerIndex++

			switch item.Handler.Type {
			// Service object.
			case HandlerTypeObject:
				m.served = true
				if m.request.IsExited() {
					break
				}
				if !m.request.IsExited() {
					m.callHandlerFunc(item.Handler.Info)
				}
			// Service handler.
			case HandlerTypeHandler:
				m.served = true
				if m.request.IsExited() {
					break
				}
				niceCallFunc(func() {
					m.callHandlerFunc(item.Handler.Info)
				})
			}
		}, func(ctx context.Context, exception error) {
			if gerror.HasStack(exception) {
				// It's already an error that has stack info.
				m.request.error = exception
			} else {
				m.request.error = gerror.WrapCodeSkip(gcode.CodeInternalError, 1, exception, "")
			}
			m.request.Response.WriteStatus(http.StatusInternalServerError, exception)
			loop = false
		})
	}

    ```
- handlerIndex和handlerMDIndex的作用如下
    1. middleware.Next() 被调用。
    2. handlerIndex 指向当前的 HandlerItemParsed。
    3. 如果当前的 HandlerItemParsed 有绑定的 middleware，则 handlerMDIndex 从 0 开始，逐个执行这些 middleware。
    4. 每执行完一个 handler 绑定的 middleware，handlerMDIndex 递增。
    5. 当 handlerMDIndex 达到 handler 绑定的 middleware 数组的末尾时，才会执行这个 handler 本身。
    6. handler 执行完毕后，handlerIndex 递增，指向下一个 HandlerItemParsed，然后重复步骤 3。
- m.request.handlers包含的handler如下:
   - 假设定义如下hook和middleware
        ```go 
        s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
                fmt.Println("HookBeforeServe")
            })
            s.BindHookHandler("/*", ghttp.HookAfterServe, func(r *ghttp.Request) {
                fmt.Println("HookAfterServe")
            })
            s.Group("/", func(group *ghttp.RouterGroup) {
                group.Middleware(middleware.HandlerResponse)
                group.Bind(
                    organization.NewV1(),
                )
            })
            // handlerResponse为后置中间件代码如下:
            func HandlerResponse(r *ghttp.Request) {
            r.Middleware.Next()
            if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {
                return
            }
            //... 其他代码
            }
            // 请求处理函数
            func (c *ControllerV1) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {

                insertId, err := dao.User.Ctx(ctx).Data(do.User{
                    Name:   req.Name,
                    Status: v1.StatusOK,
                    Age:    req.Age,
                }).InsertAndGetId()
                if err != nil {
                    return nil, err
                }
                res = &v1.UserCreateRes{
                    Id: uint(insertId),
                }
                return                
            }
        ```
    则m.request.handlers包含的handler如下4个handle:
    - 1. 内置的链路追踪: github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing
    - 2. 请求处理函数: gf-boilerplate/internal/controller/backend/organization.(*ControllerV1).UserCreate,该handle内部的middleware为[middleware.HandlerResponse]
    - 3. HookBeforeServe 
    - 4. HookAfterServe
    但是Hook在这里不会处理,会直接跳过
- 执行绑定到指定路由的middleware,UserCreate的对应的Middleware为[middleware.HandlerResponse]:
    ```go
    if m.handlerMDIndex < len(item.Handler.Middleware) {
            md := item.Handler.Middleware[m.handlerMDIndex]
            m.handlerMDIndex++
            niceCallFunc(func() {
                md(m.request)
            })
            loop = false
            return
        }

    ```
    - 此时递归执行Next方法,Next方法通过handlerIndex和handlerMDIndex记住了当前执行到哪个handler,执行到哪个middleware,此时handlerIndex=1,handlerMDIndex=1,因为不满足`m.handlerMDIndex < len(item.Handler.Middleware)`,会执行到
        ```go
            if !m.request.IsExited() {
                m.callHandlerFunc(item.Handler.Info)
            }
        
            func (m *middleware) callHandlerFunc(funcInfo handlerFuncInfo) {
                niceCallFunc(func() {
                    funcInfo.Func(m.request)
                })
            }
            func niceCallFunc(f func()) {
                defer func() {
                    if exception := recover(); exception != nil {
                        switch exception {
                        case exceptionExit, exceptionExitAll:
                            return

                        default:
                            if v, ok := exception.(error); ok && gerror.HasStack(v) {
                                // It's already an error that has stack info.
                                panic(v)
                            }
                            // Create a new error with stack info.
                            // Note that there's a skip pointing the start stacktrace
                            // of the real error point.
                            if v, ok := exception.(error); ok {
                                if gerror.Code(v) != gcode.CodeNil {
                                    panic(v)
                                } else {
                                    panic(gerror.WrapCodeSkip(
                                        gcode.CodeInternalPanic, 1, v, "exception recovered",
                                    ))
                                }
                            } else {
                                panic(gerror.NewCodeSkipf(
                                    gcode.CodeInternalPanic, 1, "exception recovered: %+v", exception,
                                ))
                            }
                        }
                    }
                }()
                f()
            }

            // item.Handler.Info类型如下
            // handlerFuncInfo 包含处理器函数地址及其反射类型信息
            handlerFuncInfo struct { 
                // func 的实际位置net/ghttp/ghttp_server_service_handler.go:242
                // 这里通过反射获取到UserCreate相关参数,并执行UserCreate
                Func            HandlerFunc      // 处理器函数地址
                Type            reflect.Type     // 处理器的反射类型信息（用于功能扩展）
                Value           reflect.Value    // 处理器的反射值信息（用于功能扩展）
                IsStrictRoute   bool             // 是否启用严格路由匹配
                ReqStructFields []gstructs.Field // 请求体结构体字段信息
            }



        ```
    -  业务处理函数UserCreate在上述`item.Handler.Info.Func`中被调用,代码如下:
        ```go

            func createRouterFunc(funcInfo handlerFuncInfo) func(r *Request) {
                return func(r *Request) {
                    // 通过反射生成参数
                    if funcInfo.Type.NumIn() == 2 {
                        // ...相关处理
                        inputValues = append(inputValues, inputObject)
                    }
                    // 通过反射执行UserCreate
                    results := funcInfo.Value.Call(inputValues)
                    switch len(results) {
                    case 2: // 规范的返回值
                        r.handlerResponse = results[0].Interface()
                        if !results[1].IsNil() {
                            if err, ok = results[1].Interface().(error); ok {
                                r.error = err
                            }
                        }
                    }
                }
            }
        ```
    - 如果UserCreate内部发生panic,panic的第一次捕获位置为niceCallFunc,具体位置:net/ghttp/ghttp_func.go:54,会继续向上抛出异常,下一次的捕获位置为Next方法调用时的TryCatch方法,具体位置:net/ghttp/ghttp_middleware.go:102,将异常写入到response中,并返回500错误,loop=false,退出此次循环,退出循环后,会返回上一次的Next方法,继续执行中间件HandlerResponse方法中`r.Middleware.Next()`下方代码,执行到`if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {`return返回
        ```go
        func niceCallFunc(f func()) {
            defer func(){
                    panic(gerror.NewCodeSkipf(
                            gcode.CodeInternalPanic, 1, "exception recovered: %+v", exception,
                        ))
            }
        }
        // 在Next方法中调用TryCatch方法,具体位置:net/ghttp/ghttp_middleware.go:102,
        func(ctx context.Context, exception error) {
                if gerror.HasStack(exception) {
                    m.request.error = exception
                } 
                m.request.Response.WriteStatus(http.StatusInternalServerError, exception)
                loop = false
        }
    ```
- Middlware.Next中的逐层退出,只到Middleware.Next全部退出,回到ServerHttp方法中,继续执行handleAfterRequestDone方法
    ```go
        func (s *Server) handleAfterRequestDone(request *Request) {
        request.LeaveTime = gtime.Now()
        if request.error != nil {
            // 走这里,处理错误日志,打印到控制台
            s.handleErrorLog(request.error, request)
        } else {
            if exception := recover(); exception != nil {
                // 继续捕获异常
            }
        }
        // 记录访问日志
        s.handleAccessLog(request)
        // 关闭session,如果session存在,则更新TTL
        if err := request.Session.Close(); err != nil {
            intlog.Errorf(request.Context(), `%+v`, err)
        }

        // 关闭请求和响应体
        // 释放文件描述符
        err := request.Request.Body.Close()
        if err != nil {
            intlog.Errorf(request.Context(), `%+v`, err)
        }
        if request.Request.Response != nil {
            // 关闭响应,后续无法再写入,这里只是关闭文件描述符,并没有关闭连接
            err = request.Request.Response.Body.Close()
            if err != nil {
                intlog.Errorf(request.Context(), `%+v`, err)
            }
        }
    }

    ```
- 执行最终回到go 源码net/http/server.go:2092处
    ```go
    	serverHandler{c.server}.ServeHTTP(w, w.req)
		inFlightResponse = nil
		w.cancelCtx()
		if c.hijacked() {
			return
		}
        // 这里执行真正向客户端写入数据
		w.finishRequest()
		c.rwc.SetWriteDeadline(time.Time{})
		if !w.shouldReuseConnection() {
			if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
				c.closeWriteAndWait()
			}
            // 这里退出
			return
		}    
    ```

