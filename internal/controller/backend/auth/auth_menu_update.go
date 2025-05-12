package auth

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-boilerplate/apibackend/auth/menu"
	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/errctx"
)

func (c *ControllerMenu) Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error) {
	var (
		data            do.Menu
		current, parent *entity.Menu
	)
	res = &menu.UpdateRes{}
	if err = gconv.Scan(req, &data); err != nil {
		return res, err
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Menu.Ctx(ctx).WherePri(req.Id).Scan(&current)
		if err != nil {
			return err
		}
		if req.Pid == current.Id {
			return errctx.New(ctx, "cg.menu.pid_equal_id")
		}
		if current.Pid == req.Pid {
			_, err = dao.Menu.Ctx(ctx).WherePri(req.Id).Update(data)
			return err
		}

		var (
			newLevel = consts.MenuFirstLevel
			newTree  = consts.MenuFirstLevelTree
		)
		if req.Pid != 0 {
			err = dao.Menu.Ctx(ctx).WherePri(req.Pid).Scan(&parent)
			if err != nil {
				return err
			}
			if parent == nil {
				return errctx.New(ctx, "cg.menu.parent_not_found")
			}
			newLevel = parent.Level + 1
			newTree = gstr.Trim(gconv.String(req.Pid)+","+parent.Tree, ",")
		}

		if err = refreshChildrenTree(ctx, current.Id, newLevel, newTree); err != nil {
			return err
		}
		data.Tree = newTree
		data.Level = newLevel
		_, err = dao.Menu.Ctx(ctx).WherePri(req.Id).Update(data)
		return err
	})
	return &menu.UpdateRes{
		Id: req.Id,
	}, err
}

// refreshChildrenTree 递归刷新菜单子节点的树形结构和层级
//
// @param ctx context.Context 上下文
// @param pid uint64 父菜单ID
// @param pLevel uint 父菜单层级
// @param pTree string 父菜单树形结构字符串
//
// @return error 错误信息
//
// @note
// 1. 当修改菜单的父级ID时,需要同步更新其所有子菜单的tree和level
// 2. tree字段格式为: "3,2,1" 表示从父节点到根节点的ID路径
// 3. level字段表示当前节点的层级,从1开始
func refreshChildrenTree(ctx context.Context, pid uint64, pLevel uint, pTree string) error {
	columns := dao.Menu.Columns()
	var children []entity.Menu
	err := dao.Menu.Ctx(ctx).Where(columns.Pid, pid).Scan(&children)
	if err != nil {
		return err
	}
	if len(children) == 0 {
		return nil
	}
	updateIDs := make([]uint64, len(children))
	for i, child := range children {
		updateIDs[i] = child.Id
		if err = refreshChildrenTree(ctx, child.Id, pLevel+1, gstr.Trim(gconv.String(pid)+","+pTree, ",")); err != nil {
			return err
		}
	}
	if len(updateIDs) > 0 {
		updateData := g.Map{
			columns.Tree:  gstr.Trim(gconv.String(pid)+","+pTree, ","),
			columns.Level: pLevel + 1,
		}
		_, err = dao.Menu.Ctx(ctx).Where(columns.Id, updateIDs).Update(updateData)
		if err != nil {
			return err
		}
	}
	return nil
}
