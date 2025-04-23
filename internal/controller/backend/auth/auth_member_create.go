package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/member"
)

func (c *ControllerMember) Create(ctx context.Context, req *member.CreateReq) (res *member.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
