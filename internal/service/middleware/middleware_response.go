// Package middleware 中间件-响应处理
package middleware

import (
	"context"
	"fmt"
	"mime"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"

	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

var nilData = map[string]interface{}{}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
	maxParamLength          = 50
	paramTruncateSuffix     = "..."
	paramTruncateLength     = 47 // maxParamLength - len(paramTruncateSuffix)
	millisecondsDivisor     = 1000.0
)

// var needHiddenErrors = map[gcode.Code]struct{}{
// 	gcode.CodeInternalError:             {},
// 	gcode.CodeDbOperationError:          {},
// 	gcode.CodeInternalPanic:             {},
// 	gcode.CodeNotImplemented:            {},
// 	gcode.CodeNotSupported:              {},
// 	gcode.CodeOperationFailed:           {},
// 	gcode.CodeServerBusy:                {},
// 	gcode.CodeUnknown:                   {},
// 	gcode.CodeNecessaryPackageNotImport: {},
// }

// streamContentTypes 流式响应的content type
var streamContentTypes = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}

// DefaultResponse 返回给前端的响应结构
type DefaultResponse struct {
	Code    int         `json:"code"            dc:"错误码"`
	Message string      `json:"message"         dc:"错误信息"`
	Data    interface{} `json:"data"            dc:"返回数据"`
	Error   interface{} `json:"error,omitempty" dc:"当为debug模式时,error为堆栈信息"`
}

// HandlerResponse 是一个中间件处理函数,用于统一处理HTTP请求的响应格式
func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	_, span := st.GetTracer().NewSpan(r.Context(), "middleware_response")
	defer span.End()

	// 检查是否为流式响应
	if isStreamResponse(r) {
		return
	}
	// 如果响应已经写入,则清空响应
	if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
		r.Response.ClearBuffer()
	}

	isDebug := st.GetConfig().GetDebug(r.Context())
	if r.Header.Get("X-Debug") != "" {
		isDebug = true
	}
	err := r.GetError()
	// 如果没有错误,返回成功响应
	if err == nil {
		r.Response.WriteJsonExit(DefaultResponse{
			Code:    gcode.CodeOK.Code(),
			Message: g.I18n().T(r.Context(), "business.success"),
			Data:    r.GetHandlerResponse(),
		})
		return
	}
	// 构建错误响应
	resp := buildErrResponse(r.Context(), err, isDebug)

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

// buildErrResponse 构建响应结构
func buildErrResponse(ctx context.Context, err error, isDebug bool) DefaultResponse {
	var (
		data = nilData // 避免json序列化后,data为空时,显示为null
	)

	code := gerror.Code(err)
	resp := DefaultResponse{
		Code:    code.Code(),
		Message: "",
		Data:    data,
	}
	var ctxVariable interface{}

	if code.Code() == errctx.CodeWithCtx { // 自定义错误
		ctxVariable = code.Detail()
		msg := code.Message()
		if gstr.TrimAll(msg) != "" {
			resp.Message = msg
		}
	}

	var validError gvalid.Error
	if gerror.As(err, &validError) { // gf校验错误
		err = validError.FirstError()
		resp.Message = validError.Error()
		internalCode := gerror.Code(err)
		if internalCode.Code() == errctx.CodeWithCtx { // gf校验错误包裹了自定义错误
			resp.Message = internalCode.Message()
		}
	}

	if resp.Message == "" {
		resp.Message = g.I18n().T(ctx, "business.fail")
	}

	if isDebug {
		var validError gvalid.Error
		if gerror.As(err, &validError) {
			resp.Error = parseStack(fmt.Sprintf("%+v", validError.FirstError()), ctxVariable)
		} else {
			resp.Error = parseStack(gerror.Stack(err), ctxVariable)
		}
	}

	// 显示业务错误提示信息
	// if _, ok := needHiddenErrors[code]; !ok {
	// 	// 使用两次current的原因是当使用自定义验证类并且自定义验证类出现了内部错误,使用两次current可以获取到自定义验证类定义的错误信息,使用一次会获取到gvalid.Error
	// 	resp.Message = gerror.Current(gerror.Current(err)).Error()
	// }
	return resp
}

func parseStack(s string, ctxVariable interface{}) []string {
	if ctxVariable != nil {
		s = fmt.Sprintf("上下文变量: %+v\n%s", ctxVariable, s)
	}
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
