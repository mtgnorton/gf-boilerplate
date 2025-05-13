package model

type MenuTreeItem struct {
	Id    uint64 `json:"id"` // 菜单ID
	Pid   uint64 `json:"pid"`
	Level uint   `json:"level"`
	Tree  string `json:"tree"`
}
