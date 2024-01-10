package methods

import "encoding/json"

// JsonMethod 实现redis序列化、反序列化接口方法
type JsonMethod struct {
}

// MarshalBinary 支持json序列化
func (m *JsonMethod) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// UnmarshalBinary 支持json反序列化
func (m *JsonMethod) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
