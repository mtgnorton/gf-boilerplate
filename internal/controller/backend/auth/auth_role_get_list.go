package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
)

func (c *ControllerRole) GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error) {
	res = &role.GetListRes{}
	columns := dao.Role.Columns()

	err = dao.Role.Ctx(ctx).
		Fields(columns).
		OrderDesc(columns.Id).
		Scan(&res.List)
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}
	return res, nil
}
