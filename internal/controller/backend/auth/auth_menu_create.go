package auth

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-boilerplate/apibackend/auth/menu"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/errctx"
)

func (c *ControllerMenu) Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error) {
	var (
		columns = dao.Menu.Columns()
		data    do.Menu
	)
	res = &menu.CreateRes{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data.Level = 1
		if req.Pid != 0 {
			var parent *entity.Menu
			err = dao.Menu.Ctx(ctx).Where(columns.Id, req.Pid).Scan(&parent)
			if err != nil {
				return err
			}
			if parent == nil {
				return errctx.New(ctx, "cg.menu.parent_not_found")
			}
			data.Level = parent.Level + 1
			data.Tree = gstr.Trim(gconv.String(req.Pid)+","+parent.Tree, ",")
		}
		res.Id, err = dao.Menu.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		return nil
	})
	return
}
