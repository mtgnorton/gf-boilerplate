package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
)

func (c *ControllerRole) GetOne(ctx context.Context, req *role.GetOneReq) (res *role.GetOneRes, err error) {
	err = dao.Role.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}

	return res, nil
}
