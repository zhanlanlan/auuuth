package models

import "auuuth/utils"

// User ...
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// ...USERPREFIX
var (
	userPrefix = []byte("user")
)

// Marshal ...
func (u *User) Marshal() ([]byte, error) {
	return utils.MarshallOrErr(u)
}

// IntoRecord ...
func (u *User) IntoRecord() Record {
	pk := CreateCompositeKey(primaryPrefix, userPrefix, []byte(u.ID))

	record := NewRecord(pk, u)

	// 添加 user_name 字段的索引
	idx1 := CreateCompositeKey(indexPrefix, userPrefix, []byte(u.UserName))
	record.AddIndex(idx1)

	return record
}
