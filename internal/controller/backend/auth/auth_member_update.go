package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/member"
)

func (c *ControllerMember) Update(ctx context.Context, req *member.UpdateReq) (res *member.UpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
