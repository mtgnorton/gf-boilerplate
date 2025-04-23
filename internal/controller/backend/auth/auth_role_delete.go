package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
)

func (c *ControllerRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	_, err = dao.Role.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}

	return &role.DeleteRes{Id: req.Id}, nil
}
