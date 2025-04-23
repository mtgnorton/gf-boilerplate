package valid

import (
	"github.com/gogf/gf/v2/util/gvalid"
)

func RegisterAll() {
	gvalid.RegisterRule("unique-field", UniqueField)
	gvalid.RegisterRule("unique-field-add", UniqueFieldAdd)
	gvalid.RegisterRule("unique-field-update", UniqueFieldUpdate)
	gvalid.RegisterRule("exist-record", ExistRecord)
}
