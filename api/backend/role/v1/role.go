package role

type RoleCreateReq struct {
	Name string `json:"name" v:"required#角色名称不能为空"`
	Code string `json:"code" v:"required#角色编码不能为空"`
}

type RoleCreateRes struct {
	Id string `json:"id"`
}
