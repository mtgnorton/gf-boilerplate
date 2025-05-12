package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerMember) GetPageList(
	ctx context.Context,
	req *member.GetPageListReq,
) (res *member.GetPageListRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "GetPageList")
	defer span.End()

	columns := dao.Member.Columns()
	res = &member.GetPageListRes{}
	res.Page = req.Page
	res.Size = req.Size

	err = dao.Member.Ctx(ctx).
		Fields(columns).
		Page(req.Offset(), req.Size).
		OrderDesc(columns.Id).
		ScanAndCount(&res.List, &res.Total, false)
	return res, err
}
