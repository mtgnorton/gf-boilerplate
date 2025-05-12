package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerMember) Delete(ctx context.Context, req *member.DeleteReq) (res *member.DeleteRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Delete")
	defer span.End()

	// 删除管理员
	_, err = dao.Member.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, errctx.New(ctx, "cg.member.delete_failed")
	}

	return &member.DeleteRes{Id: req.Id}, err
}
