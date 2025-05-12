package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-boilerplate/apibackend/auth/access"
)

func (c *ControllerAccess) Logout(ctx context.Context, req *access.LogoutReq) (res *access.LogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
