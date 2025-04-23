package middleware

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gvalid"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/service/errctx"
)

func Test_buildErrResponse(t *testing.T) {
	ctx := context.Background()

	// 预期错误->手动抛出 关闭debug
	gtest.C(t, func(t *gtest.T) {
		err := errctx.New("错误提示")
		resp := buildErrResponse(ctx, err, false)
		t.Assert(resp, DefaultResponse{
			Code:    errctx.CodeWithCtx,
			Message: "错误提示",
			Data:    nilData,
			Error:   nil,
		})
	})

	// 以下开启debug

	// 预期错误->手动抛出
	gtest.C(t, func(t *gtest.T) {
		err := errctx.New("错误提示")
		resp := buildErrResponse(ctx, err, true)
		t.Assert(resp.Code, errctx.CodeWithCtx)
		t.Assert(resp.Message, "错误提示")
		t.Assert(resp.Data, nilData)
		t.AssertNE(resp.Error, nil)
	})

	// 预期错误->内置的gf校验错误
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name string `v:"required"`
		}
		user := User{}
		//  gf校验错误
		err := g.Validator().Data(user).Run(ctx)
		resp := buildErrResponse(ctx, err, true)
		t.Assert(resp.Code, gcode.CodeValidationFailed.Code())
		t.Assert(resp.Message, "The Name field is required")
		t.Assert(resp.Data, nilData)
		t.AssertNE(resp.Error, nil)
	})
	// 预期错误->自定义的校验错误
	gtest.C(t, func(t *gtest.T) {
		customRuleFunc := func(ctx context.Context, in gvalid.RuleFuncInput) (e error) {
			return errctx.New("自定义错误")
		}
		gvalid.RegisterRule("custom-rule", customRuleFunc)
		type User struct {
			Name string `v:"custom-rule"`
		}
		user := User{}
		err := g.Validator().Data(user).Run(ctx)
		resp := buildErrResponse(ctx, err, true)
		t.Assert(resp.Code, gcode.CodeValidationFailed.Code())
		t.Assert(resp.Message, "自定义错误")
		t.Assert(resp.Data, nilData)
		t.AssertNE(resp.Error, nil)
	})

	// 非预期错误
	gtest.C(t, func(t *gtest.T) {
		//nolint:errcheck
		g.Config().GetAdapter().(*gcfg.AdapterFile).SetFileName(gtest.DataPath("config.test.yaml"))
		subCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(10))
		defer cancel()
		_, err := g.Model("not-exist-table").Ctx(subCtx).Count()

		resp := buildErrResponse(subCtx, err, true)
		t.Assert(resp.Code, gcode.CodeDbOperationError.Code())
		t.Assert(resp.Message, "Operation failed")
		t.Assert(resp.Data, nilData)
		t.AssertNE(resp.Error, nil)
	})
}
