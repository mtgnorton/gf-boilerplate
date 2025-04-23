package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
)

func (c *ControllerRole) GetPageList(
	ctx context.Context,
	req *role.GetPageListReq,
) (res *role.GetPageListRes, err error) {
	columns := dao.Role.Columns()
	res = &role.GetPageListRes{}
	res.Page = req.Page
	res.Size = req.Size

	err = dao.Role.Ctx(ctx).
		Fields(columns).
		Page(req.Offset(), req.Size).
		OrderDesc(columns.Id).
		ScanAndCount(&res.List, &res.Total, false)
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}
	return res, nil
}
