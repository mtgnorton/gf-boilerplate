// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id     uint   `json:"id"     orm:"id"     description:""`            //
	Name   string `json:"name"   orm:"name"   description:"username"`    // username
	Status int    `json:"status" orm:"status" description:"user status"` // user status
	Age    uint   `json:"age"    orm:"age"    description:"user age"`    // user age
}
