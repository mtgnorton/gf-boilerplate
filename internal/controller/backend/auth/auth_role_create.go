package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
)

func (c *ControllerRole) Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error) {
	id, err := dao.Role.Ctx(ctx).Data(do.Role{
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}

	return &role.CreateRes{Id: id}, nil
}
