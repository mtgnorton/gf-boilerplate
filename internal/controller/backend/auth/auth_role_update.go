package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerRole) Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Update")
	defer span.End()

	var (
		oldRole *entity.Role
	)

	err = dao.Role.Ctx(ctx).WherePri(req.Id).Scan(&oldRole)
	if err != nil {
		return nil, err
	}
	if oldRole == nil {
		return nil, errctx.New(ctx, "cg.role.not_found")
	}
	_, err = dao.Role.Ctx(ctx).WherePri(req.Id).Data(req.RoleArg).Update()
	if err != nil {
		return nil, err
	}
	return &role.UpdateRes{Id: req.Id}, nil
}
