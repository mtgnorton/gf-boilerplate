// Package middleware 中间件-响应处理
package middleware

import (
	"encoding/json"
	"fmt"
	"mime"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"

	"gf-boilerplate/internal/service/global"
)

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
	maxParamLength          = 50
	paramTruncateSuffix     = "..."
	paramTruncateLength     = 47 // maxParamLength - len(paramTruncateSuffix)
	millisecondsDivisor     = 1000.0
)

var needHiddenErrors = map[gcode.Code]struct{}{
	gcode.CodeInternalError:             {},
	gcode.CodeDbOperationError:          {},
	gcode.CodeInternalPanic:             {},
	gcode.CodeNotImplemented:            {},
	gcode.CodeNotSupported:              {},
	gcode.CodeOperationFailed:           {},
	gcode.CodeServerBusy:                {},
	gcode.CodeUnknown:                   {},
	gcode.CodeNecessaryPackageNotImport: {},
}

// streamContentTypes 流式响应的content type
var streamContentTypes = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}

// DefaultHandlerResponse 返回给前端的响应结构
type DefaultHandlerResponse struct {
	Code    int         `json:"code"            dc:"错误码"`
	Message string      `json:"message"         dc:"错误信息"`
	Data    interface{} `json:"data"            dc:"返回数据"`
	Error   interface{} `json:"error,omitempty" dc:"当为debug模式时,error为堆栈信息"`
}

// HandlerResponse 是一个中间件处理函数,用于统一处理HTTP请求的响应格式
func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 检查是否为流式响应
	if isStreamResponse(r) {
		return
	}
	// 处理错误日志
	handleErrorLog(r)

	// 构建响应
	resp := buildResponse(r)

	// 返回响应
	r.Response.WriteJsonExit(resp)
}

// isStreamResponse 检查是否为流式响应
func isStreamResponse(r *ghttp.Request) bool {
	mediaType, _, err := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	if err != nil {
		return false
	}
	for _, ct := range streamContentTypes {
		if mediaType == ct {
			return true
		}
	}
	return false
}

// checkAndClearCritical 检查是否发生了critical错误
func checkAndClearCritical(r *ghttp.Request) bool {
	if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
		r.Response.ClearBuffer()
		return true
	}
	return false
}

// buildResponse 构建响应结构
func buildResponse(r *ghttp.Request) DefaultHandlerResponse {
	var (
		err     = r.GetError()
		data    = map[string]interface{}{} // 避免json序列化后,data为空时,显示为null
		isDebug = global.GetConfig().GetDebug(r.Context())
	)

	// 如果没有错误,返回成功响应
	if err == nil {
		return DefaultHandlerResponse{
			Code:    gcode.CodeOK.Code(),
			Message: g.I18n().T(r.Context(), "business.success"),
			Data:    r.GetHandlerResponse(),
		}
	}

	code := gerror.Code(err)
	resp := DefaultHandlerResponse{
		Code:    code.Code(),
		Message: g.I18n().T(r.Context(), "business.fail"),
		Data:    data,
	}

	if isDebug {
		resp.Error = parseStack(gerror.Stack(err))
	}

	// 显示业务错误提示信息
	if _, ok := needHiddenErrors[code]; !ok {
		resp.Message = gerror.Current(err).Error()
	}

	return resp
}

// handleErrorLog 处理错误日志
func handleErrorLog(r *ghttp.Request) {

	// 检查是否发生了critical错误
	isCritical := checkAndClearCritical(r)

	err := r.GetError()
	if err == nil {
		return
	}

	code := gerror.Code(err)
	isBusinessError := code == gcode.CodeNil

	if isBusinessError {
		g.Log().Info(r.Context(), "Business error: %v", err)
		return
	}

	// 处理非业务错误
	c := errorLogContent(err, r)
	g.Log().Stack(false).Stdout(true).Error(r.Context(), c)

	// 如果是critical错误,发送告警通知
	if isCritical {
		go func() {
			notifyErr := global.GetNotifier(r.Context()).Send(r.Context(), c)
			if notifyErr != nil {
				g.Log().Warningf(r.Context(), "发送告警通知失败: %+v", notifyErr)
			}
		}()
	}
}

func parseStack(s string) []string {
	stacks := gstr.Split(s, "\n")
	var result []string
	for i := 0; i < len(stacks); i++ {
		if stacks[i] == "" {
			continue
		}
		result = append(result, gstr.Replace(stacks[i], "\t", "--> "))
	}
	return result
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

// errorLogContent 接管gf框架的默认错误处理,添加参数信息
func errorLogContent(err error, r *ghttp.Request) string {
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
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v", "params: 
%s
"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.Sub(r.EnterTime))/millisecondsDivisor,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr, params,
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
