package valid

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
)

// argNumber 规则参数数量形式如 rule:arg 的形式,如 unique-field:role,id 则argNumber为2
const argNumber = 2

// UniqueField 检查唯一字段是否存在,当id存在时类似update,不存在时类似add
func UniqueField(ctx context.Context, in gvalid.RuleFuncInput) (e error) {
	// 解析规则参数
	ruleArray := strings.Split(in.Rule, ":")
	if len(ruleArray) != argNumber {
		return gerror.New("unique-field格式错误")
	}

	// 解析表和列
	params := strings.Split(ruleArray[1], ",")
	if len(params) == 0 {
		return gerror.New("unique-field 至少需要一个表名参数")
	}

	field := in.Field
	table := params[0]
	if len(params) > 1 {
		field = params[1]
	}

	// 获取要检查的值
	value := in.Value.String()

	// 获取主键值
	id := in.Data.MapStrAny()["id"]

	model := g.Model(table)
	if id != nil {
		model = model.WhereNot("id", id)
	}
	model = model.Where(field, value)

	// 查询数据库判断是否存在
	count, err := model.Count()
	if err != nil {
		return gerror.Wrap(err, "内部错误")
	}

	if count > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, g.I18n().T(ctx, in.Message))
	}

	return nil
}

// UniqueFieldAdd 添加数据时，检查唯一字段是否存在
// in参数类似如下:
//
//	{
//	    Rule:      "unique-field:role",
//	    Message:   "角色名称已存在",
//	    Field:     "Name",
//	    ValueType: "string",
//	    Value:     "1sss",
//	    Data:      "{\"Code\":\"2@2\",\"Name\":\"1sss\",\"Status\":\"disabled\",\"id\":\"1\"}",
//	}
func UniqueFieldAdd(ctx context.Context, in gvalid.RuleFuncInput) (e error) {
	// 解析规则参数
	ruleArray := strings.Split(in.Rule, ":")
	if len(ruleArray) != argNumber {
		return gerror.New("unique-field格式错误")
	}

	// 解析表和列
	params := strings.Split(ruleArray[1], ",")
	if len(params) == 0 {
		return gerror.New("unique-field 至少需要一个表名参数")
	}

	field := in.Field
	table := params[0]
	if len(params) > 1 {
		field = params[1]
	}

	model := g.Model(table)
	id, ok := in.Data.MapStrAny()["id"]
	if ok {
		model = model.WhereNot("id", id)
	}

	// 获取要检查的值
	value := in.Value.String()
	// 查询数据库判断是否存在
	count, err := model.Where(field, value).Count()
	if err != nil {
		return gerror.Wrap(err, "内部错误")
	}
	if count > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, g.I18n().T(ctx, in.Message))
	}

	return nil
}

// UniqueFieldUpdate 更新数据时，检查唯一字段是否存在,默认检查主键id
func UniqueFieldUpdate(ctx context.Context, in gvalid.RuleFuncInput) (e error) {
	// 解析规则参数
	ruleArray := strings.Split(in.Rule, ":")
	if len(ruleArray) != argNumber {
		return gerror.New("unique-field格式错误")
	}

	// 解析表和列
	params := strings.Split(ruleArray[1], ",")
	if len(params) == 0 {
		return gerror.New("unique-field 至少需要一个表名参数")
	}

	field := in.Field
	table := params[0]
	if len(params) > 1 {
		field = params[1]
	}

	// 获取要检查的值
	value := in.Value.String()

	// 获取主键值
	id := in.Data.MapStrAny()["id"]
	if id == nil {
		return gerror.New("unique-field-update 缺少主键id")
	}
	// 查询数据库判断是否存在
	count, err := g.Model(table).Where(field, value).WhereNot("id", id).Count()
	if err != nil {
		return gerror.Wrap(err, "内部错误")
	}

	if count > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, g.I18n().T(ctx, in.Message))
	}

	return nil
}
