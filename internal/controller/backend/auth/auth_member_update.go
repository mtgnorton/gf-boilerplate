package auth

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/internal/dao"
	"gf-boilerplate/internal/model/do"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/errctx"
)

func (c *ControllerMember) Update(ctx context.Context, req *member.UpdateReq) (res *member.UpdateRes, err error) {
	var (
		data      do.Member
		oldMember *entity.Member
	)
	res = &member.UpdateRes{}

	if err = dao.Member.Ctx(ctx).Where(dao.Member.Columns().Id, req.Id).Scan(&oldMember); err != nil {
		return nil, errctx.New(ctx, "cg.member.not_found")
	}

	// 复制更新字段
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	// 如果更新了密码,需要重新生成密码哈希
	if req.Password != "" {
		data.PasswordHash = gmd5.MustEncryptString(req.Password + oldMember.Salt)
	}

	// 更新数据
	_, err = dao.Member.Ctx(ctx).Where(dao.Member.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, errctx.New(ctx, "cg.member.update_failed")
	}

	res.Id = req.Id
	return
}
