package enum

type MenuType string

const (
	MenuTypeDirectory MenuType = "directory"
	MenuTypeMenu      MenuType = "menu"
	MenuTypeButton    MenuType = "button"
)
