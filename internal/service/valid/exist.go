package valid

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"

	"gf-boilerplate/internal/service/errctx"
)

// ExistRecord 检查ID是否存在
func ExistRecord(ctx context.Context, in gvalid.RuleFuncInput) (e error) {
	// 解析规则参数
	ruleArray := strings.Split(in.Rule, ":")
	if len(ruleArray) != argNumber {
		return gerror.New("exist-id格式错误")
	}
	// 获取表名
	table := ruleArray[1]

	// 获取要检查的ID值
	id := in.Value.Int64()

	// 查询数据库判断是否存在
	count, err := g.Model(table).Where("id", id).Count()
	if err != nil {
		return errctx.Wrap(err, "")
	}

	if count == 0 {
		return errctx.New(g.I18n().T(ctx, in.Message))
	}

	return nil
}
