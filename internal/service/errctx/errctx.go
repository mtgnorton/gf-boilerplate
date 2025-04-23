package errctx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const CodeWithCtx = 1001

// New 一般用于预期错误, 且需要记录上下文相关信息
func New(message string, details ...interface{}) error {
	var detail interface{}
	if len(details) == 0 {
		detail = nil
	} else {
		detail = details[0]
	}
	return gerror.NewCode(gcode.New(CodeWithCtx, message, detail), message)
}

// Wrap 一般用于非预期错误, 且需要记录上下文相关信息
func Wrap(err error, message string, details ...interface{}) error {
	var detail interface{}
	if len(details) == 0 {
		detail = nil
	} else {
		detail = details[0]
	}
	return gerror.WrapCode(gcode.New(CodeWithCtx, message, detail), err, message)
}
