package auth

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
	"gf-boilerplate/internal/service/errctx"
)

func (c *ControllerMember) Create(ctx context.Context, req *member.CreateReq) (res *member.CreateRes, err error) {
	var (
		data do.Member
	)
	res = &member.CreateRes{}

	// 生成密码盐和密码哈希
	salt := grand.S(consts.PasswordSaltLength)
	data.PasswordHash = gmd5.MustEncryptString(req.Password + salt)
	data.Salt = salt

	// 复制其他字段
	if err = gconv.Scan(req, &data); err != nil {
		return
	}

	// 创建管理员
	res.Id, err = dao.Member.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, errctx.New(ctx, "cg.member.create_failed")
	}
	return
}
