package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/menu"
	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerMenu) Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Delete")
	defer span.End()

	var (
		columns = dao.Menu.Columns()
	)

	code, err := dao.Menu.Ctx(ctx).WherePri(req.Id).Fields(columns.Code).Value()
	if err != nil {
		return nil, err
	}
	if code.String() == consts.RoleSuperAdmin {
		return nil, errctx.New(ctx, "cg.menu.cannot_delete_super_admin")
	}

	hashChildren, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Pid, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if hashChildren > 0 {
		return nil, errctx.New(ctx, "cg.menu.delete_has_children")
	}

	_, err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, req.Id).Delete()
	return &menu.DeleteRes{
		Id: req.Id,
	}, err
}
