package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
)

func (c *ControllerRole) Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error) {
	_, err = dao.Role.Ctx(ctx).WherePri(req.Id).Data(req.RoleArg).Update()
	return &role.UpdateRes{Id: req.Id}, err
}
