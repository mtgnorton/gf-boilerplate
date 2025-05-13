package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerRole) GetOne(ctx context.Context, req *role.GetOneReq) (res *role.GetOneRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "GetOne")
	defer span.End()

	err = dao.Role.Ctx(ctx).WherePri(req.Id).Scan(&res)

	return res, err
}
