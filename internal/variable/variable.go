package variable

import "github.com/gogf/gf/v2/os/gtime"

type PaginationReq struct {
	Page int `json:"page" example:"10" d:"1"  v:"min:1#页码最小值不能低于1"                      dc:"当前页码"`
	Size int `json:"size" example:"1"  d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}

func (p *PaginationReq) Offset() int {
	return (p.Page - 1) * p.Size
}

type PaginationRes struct {
	Total int `json:"total" example:"100" dc:"总条数"`
	Page  int `json:"page"  example:"1"   dc:"当前页码"`
	Size  int `json:"size"  example:"10"  dc:"每页数量"`
}

type Times struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
