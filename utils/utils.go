package utils

import "encoding/json"

// MarshallOrErr 封装后的序列化函数，
// 后续可以更方便地替换成别的序列化方式。
func MarshallOrErr(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
