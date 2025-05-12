package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerMember) GetList(ctx context.Context, req *member.GetListReq) (res *member.GetListRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "GetList")
	defer span.End()

	res = &member.GetListRes{}
	columns := dao.Member.Columns()
	// 获取管理员列表
	err = dao.Member.Ctx(ctx).
		Fields(columns).
		OrderDesc(columns.Id).
		Scan(&res.List)

	return
}
