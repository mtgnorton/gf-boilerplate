// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/access"
	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/apibackend/auth/menu"
	"gf-boilerplate/apibackend/auth/role"
)

type IAuthAccess interface {
	Login(ctx context.Context, req *access.LoginReq) (res *access.LoginRes, err error)
	Logout(ctx context.Context, req *access.LogoutReq) (res *access.LogoutRes, err error)
}

type IAuthMember interface {
	Create(ctx context.Context, req *member.CreateReq) (res *member.CreateRes, err error)
	Update(ctx context.Context, req *member.UpdateReq) (res *member.UpdateRes, err error)
	Delete(ctx context.Context, req *member.DeleteReq) (res *member.DeleteRes, err error)
	GetOne(ctx context.Context, req *member.GetOneReq) (res *member.GetOneRes, err error)
	GetList(ctx context.Context, req *member.GetListReq) (res *member.GetListRes, err error)
	GetPageList(ctx context.Context, req *member.GetPageListReq) (res *member.GetPageListRes, err error)
}

type IAuthMenu interface {
	Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error)
	Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error)
	Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error)
	GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error)
}

type IAuthRole interface {
	Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error)
	Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error)
	Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error)
	GetOne(ctx context.Context, req *role.GetOneReq) (res *role.GetOneRes, err error)
	GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error)
	GetPageList(ctx context.Context, req *role.GetPageListReq) (res *role.GetPageListRes, err error)
}
