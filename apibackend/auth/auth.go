// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"gf-boilerplate/apibackend/auth/member"
	"gf-boilerplate/apibackend/auth/role"
)

type IAuthMember interface {
	Create(ctx context.Context, req *member.CreateReq) (res *member.CreateRes, err error)
	Update(ctx context.Context, req *member.UpdateReq) (res *member.UpdateRes, err error)
}

type IAuthRole interface {
	Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error)
	Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error)
	Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error)
	GetOne(ctx context.Context, req *role.GetOneReq) (res *role.GetOneRes, err error)
	GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error)
	GetPageList(ctx context.Context, req *role.GetPageListReq) (res *role.GetPageListRes, err error)
}
