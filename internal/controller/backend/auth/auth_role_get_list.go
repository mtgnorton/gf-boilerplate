package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerRole) GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "GetList")
	defer span.End()

	res = &role.GetListRes{}
	columns := dao.Role.Columns()

	err = dao.Role.Ctx(ctx).
		Fields(columns).
		OrderDesc(columns.Id).
		Scan(&res.List)

	return res, err
}
