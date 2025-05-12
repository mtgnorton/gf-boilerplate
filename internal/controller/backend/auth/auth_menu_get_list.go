package auth

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"gf-boilerplate/apibackend/auth/menu"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/entity"
)

func (c *ControllerMenu) GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error) {
	columns := dao.Menu.Columns()
	var menus []*entity.Menu
	if subErr := dao.Menu.Ctx(ctx).OrderAsc(columns.Sort).OrderAsc(columns.Id).Scan(&menus); subErr != nil {
		return nil, subErr
	}
	// 构建pid->[]*entity.Menu的map
	menuMap := make(map[uint64][]*entity.Menu)
	for _, m := range menus {
		menuMap[m.Pid] = append(menuMap[m.Pid], m)
	}
	res = &menu.GetListRes{}
	res.List, err = buildTree(0, menuMap)
	return
}

// buildTree 根据父ID递归构建菜单树形结构
//
// @param pid uint64 父菜单ID
// @param menuMap map[uint64][]*entity.Menu 以父ID为key的菜单映射表
//
// @return list []*menu.TreeItem 菜单树形结构列表
// @return err error 错误信息
//
// @note
// 1. 通过递归方式构建树形结构
// 2. menuMap中key为父ID,value为该父ID下的所有子菜单
//
// @example
//
//	menuMap := map[uint64][]*entity.Menu{
//	  0: []*entity.Menu{{Id:1, Name:"系统管理"}, {Id:2, Name:"用户管理"}},
//	  1: []*entity.Menu{{Id:3, Name:"菜单管理"}},
//	}
//
// tree, _ := buildTree(0, menuMap)
//
// 输出树形结构:
// - 系统管理
//   - 菜单管理
//
// - 用户管理
func buildTree(pid uint64, menuMap map[uint64][]*entity.Menu) (list []*menu.TreeItem, err error) {
	for _, m := range menuMap[pid] {
		item := &menu.TreeItem{}
		if subErr := gconv.Struct(m, item); subErr != nil {
			return nil, subErr
		}
		item.Children, err = buildTree(m.Id, menuMap)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
