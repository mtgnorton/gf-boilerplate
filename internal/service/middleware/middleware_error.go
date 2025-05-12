package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"

	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

func HandleError(r *ghttp.Request) {
	r.Middleware.Next()
	_, span := st.GetTracer().NewSpan(r.Context(), "middleware_error")
	defer span.End()

	isDebug := st.GetConfig().GetDebug(r.Context())
	if r.Header.Get("X-Debug") != "" {
		isDebug = true
	}
	// 检查是否发生了critical错误
	// if checkAndClearCritical(r) {
	// 	return
	// }
	handleErrorLog(r, isDebug)
}

// handleErrorLog 处理错误日志
func handleErrorLog(r *ghttp.Request, isDebug bool) {
	// 检查是否发生了critical错误
	// isCritical := checkAndClearCritical(r)
	err := r.GetError()
	if err == nil {
		return
	}
	code := gerror.Code(err)

	if code == gcode.CodeInternalPanic {
		c := errorWithRequestInfo(err, r, nil)
		g.Log().Stack(false).Stdout(true).Error(r.Context(), c)
		go func() {
			notifyErr := st.GetNotifier().Send(r.Context(), c)
			if notifyErr != nil {
				g.Log().Warningf(r.Context(), "发送告警通知失败: %+v", notifyErr)
			}
		}()
		return
	}
	var ctxVariable interface{}
	if code.Code() == errctx.CodeWithCtx { // 携带上下文变量
		ctxVariable = code.Detail()
		code = lowErrCode(err)
	}

	var validError gvalid.Error
	if code == gcode.CodeValidationFailed {
		if gerror.As(err, &validError) {
			err = validError.FirstError()
		}
	}

	isExpectedError := code == gcode.CodeValidationFailed || code == gcode.CodeNil // No error code specified.

	switch {
	case isExpectedError && isDebug:
		g.Log().
			Infof(r.Context(), "Expected error,code: %v,stack: %+v", code, errorWithRequestInfo(err, r, ctxVariable))
	case isExpectedError && !isDebug:
		g.Log().Infof(r.Context(), "Expected error,code: %v,message: %v", code, err)
	case !isExpectedError:
		g.Log().
			Stack(false).
			Stdout(true).
			Errorf(r.Context(), "Unexpected error,code: %v,message: %v", code, errorWithRequestInfo(err, r, ctxVariable))
	}
}

// errorWithRequestInfo 接管gf框架的默认错误处理,添加参数信息
func errorWithRequestInfo(err error, r *ghttp.Request, ctxVariable interface{}) string {
	var (
		code          = gerror.Code(err)
		scheme        = r.GetSchema()
		codeDetail    = code.Detail()
		codeDetailStr string
		params        string
	)

	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}

	// 获取和处理请求参数
	if r.Method == "GET" {
		params = ""
	} else {
		params = processRequestParams(r)
	}

	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v",
"params: %s"
"ctxVariable: %+v"
`,

		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.Sub(r.EnterTime))/millisecondsDivisor,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr, params, ctxVariable,
	)

	// 根据服务器配置决定是否添加错误堆栈信息
	if stack := gerror.Stack(err); stack != "" {
		content += "\nStack:\n" + stack
	} else {
		content += ", " + err.Error()
	}
	return content
}

// processRequestParams 处理请求参数
func processRequestParams(r *ghttp.Request) string {
	var bodyMap map[string]interface{}

	// 处理文件上传请求
	if r.GetMultipartForm() != nil && len(r.GetMultipartForm().File) > 0 {
		bodyMap = r.GetMap()
		// 处理文件字段
		for field := range r.GetMultipartForm().File {
			if _, ok := bodyMap[field]; ok {
				bodyMap[field] = "[file]"
			}
		}
	} else {
		// 处理普通请求
		unmarshalErr := json.Unmarshal([]byte(r.GetBodyString()), &bodyMap)
		if unmarshalErr != nil {
			return r.GetBodyString()
		}
	}

	// 处理所有参数
	processedMap, ok := processParams(bodyMap).(map[string]interface{})
	if !ok {
		return r.GetBodyString()
	}

	// 序列化处理后的参数
	paramsBytes, marshalErr := json.Marshal(processedMap)
	if marshalErr != nil {
		return r.GetBodyString()
	}

	return string(paramsBytes)
}

// processParams 处理请求参数
func processParams(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		if len(v) > maxParamLength {
			return v[:paramTruncateLength] + paramTruncateSuffix
		}
		return v
	case map[string]interface{}:
		result := make(map[string]interface{}, len(v))
		for key, val := range v {
			result[key] = processParams(val)
		}
		return result
	default:
		return v
	}
}

func lowErrCode(err error) gcode.Code {
	for {
		err = gerror.Unwrap(err)
		if err == nil {
			return gcode.CodeNil
		}
		if gerror.Code(err) != gcode.CodeNil {
			return gerror.Code(err)
		}
	}
}
