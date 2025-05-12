package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerMember) GetOne(ctx context.Context, req *member.GetOneReq) (res *member.GetOneRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "GetOne")
	defer span.End()

	res = &member.GetOneRes{}

	// 获取管理员信息
	err = dao.Member.Ctx(ctx).WherePri(req.Id).Scan(&res)

	return res, err
}
