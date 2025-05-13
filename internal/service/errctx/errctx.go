package errctx

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const CodeWithCtx = 1001

// New 创建一个新的错误对象
//
// 函数介绍:
//
//	创建一个带有错误码、错误信息和详情的新错误对象
//
// 参数说明:
//
//	ctx: 上下文对象,用于国际化错误信息
//	message: 错误信息,支持国际化
//	details: 可选的错误详情
//
// 返回值说明:
//
//	error: 包含错误码、错误信息和详情的错误对象
//
// 注意事项:
//  1. message为空时不会进行国际化转换
//  2. 只支持传入一个detail参数
//
// 举例:
//
//	return errctx.New(ctx, "cg.role.not_found")
func New(ctx context.Context, message string, details ...interface{}) error {
	var detail interface{}
	if len(details) == 0 {
		detail = nil
	} else {
		detail = details[0]
	}
	if message != "" {
		message = g.I18n().T(ctx, message)
	}
	return gerror.NewCodeSkip(gcode.New(CodeWithCtx, message, detail), 1, message)
}

// Wrap 包装一个已有的错误对象
//
// 函数介绍:
//
//	将已有错误包装成带有错误码、错误信息和详情的新错误对象
//
// 参数说明:
//
//	ctx: 上下文对象,用于国际化错误信息
//	err: 原始错误对象
//	messageAndDetails: 可选的错误信息和详情,可以传入0-2个参数:
//	  - 不传参数: 保持原始错误信息
//	  - 传1个参数: 第一个参数作为错误信息
//	  - 传2个参数: 第一个参数作为错误信息,第二个参数作为详情
//
// 返回值说明:
//
//	error: 包含原始错误和新增错误信息的错误对象
//
// 注意事项:
//  1. message为空时不会进行国际化转换
//  2. 最多支持传入两个参数
//
// 举例:
//
//	return errctx.Wrap(ctx, err, "cg.fail")
func Wrap(ctx context.Context, err error, messageAndDetails ...interface{}) error {
	var (
		message string
		detail  interface{}
	)

	switch len(messageAndDetails) {
	case 0:
		// 不设置message和detail
	case 1:
		// 仅设置message
		if msg, ok := messageAndDetails[0].(string); ok {
			message = msg
		}
	//nolint:mnd
	case 2:
		// 设置message和detail
		if msg, ok := messageAndDetails[0].(string); ok {
			message = msg
		}
		detail = messageAndDetails[1]
	}
	if message != "" {
		message = g.I18n().T(ctx, message)
	}
	return gerror.WrapCodeSkip(gcode.New(CodeWithCtx, message, detail), 1, err, message)
}
