package auth

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"

	"gf-boilerplate/apibackend/auth/access"
	"gf-boilerplate/internal/consts/enum"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
)

func (c *ControllerAccess) Login(ctx context.Context, req *access.LoginReq) (res *access.LoginRes, err error) {
	ctx, span := st.GetTracer().NewSpan(ctx, "Login")
	defer span.End()

	var (
		columns     = dao.Member.Columns()
		existMember *entity.Member
	)
	err = dao.Member.Ctx(ctx).Where(columns.Username, req.Username).Scan(&existMember)
	if err != nil {
		return nil, err
	}
	if existMember == nil {
		return nil, errctx.New(ctx, "cg.access.login.user_not_found")
	}
	currentPassword := gmd5.MustEncryptString(req.Password + existMember.Salt)
	if currentPassword != existMember.PasswordHash {
		return nil, errctx.New(ctx, "cg.access.login.password_error")
	}
	if existMember.Status != string(enum.StatusNormal) {
		return nil, errctx.New(ctx, "cg.access.login.user_disabled")
	}
	return
}
