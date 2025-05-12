package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerRole) Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Create")
	defer span.End()
	id, err := dao.Role.Ctx(ctx).Data(do.Role{
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
	}).InsertAndGetId()

	return &role.CreateRes{Id: id}, err
}
