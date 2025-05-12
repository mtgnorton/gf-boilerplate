package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/role"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Delete")
	defer span.End()
	_, err = dao.Role.Ctx(ctx).WherePri(req.Id).Delete()

	return &role.DeleteRes{Id: req.Id}, err
}
