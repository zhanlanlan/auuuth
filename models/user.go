package models

import (
	"auuuth/utils"
)

// User ...
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`

	xxxIndexs [][]byte
}

// ...USERPREFIX
const (
	userPrefix = "user"
)

// Store 保存在leveldb中
func (u *User) Store() error {
	data, err := Get([]byte(u.ID))
	if err != nil {
		return err
	} else if data != nil {
		// data already exist
		// it's fine just cover
		// 覆盖掉就好了，数据已经存在也没关系
	}

	data, err = utils.MarshallOrErr(u)
	if err != nil {
		return err
	}

	batch := NewBatch()

	// 主键指向具体的数据
	pk := CreateCompositeKey(primaryPrefix, []string{userPrefix, u.ID}...)
	batch.Put([]byte(pk), data)

	// 索引指向主键
	for _, idxKey := range u.xxxIndexs {
		batch.Put(idxKey, []byte(pk))
	}

	return batch.Write()
}

// AddIndex 添加一个索引
// 索引的添加应该在 Store 函数调用前。
func (u *User) AddIndex(idx string) {
	// 真是个弱智， byte数组和string转来转去的，全给搞成byte数组得了。
	key := CreateCompositeKey(indexPrefix, []string{userPrefix, idx}...)

	u.xxxIndexs = append(u.xxxIndexs, []byte(key))
}
