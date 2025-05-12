// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GooseDbVersion is the golang structure for table goose_db_version.
type GooseDbVersion struct {
	Id        uint64      `json:"id"        orm:"id"         ` //
	VersionId int64       `json:"versionId" orm:"version_id" ` //
	IsApplied int         `json:"isApplied" orm:"is_applied" ` //
	Tstamp    *gtime.Time `json:"tstamp"    orm:"tstamp"     ` //
}
