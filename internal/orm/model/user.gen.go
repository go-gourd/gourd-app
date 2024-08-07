// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"

	"gorm.io/plugin/soft_delete"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int32                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserName   string                `gorm:"column:user_name;not null;comment:用户名" json:"user_name"`                     // 用户名
	Mobile     string                `gorm:"column:mobile;not null;comment:手机号" json:"mobile"`                           // 手机号
	Password   string                `gorm:"column:password;not null;comment:密码" json:"password"`                        // 密码
	CreateTime uint                  `gorm:"column:create_time;not null;autoCreateTime;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime uint                  `gorm:"column:update_time;not null;autoUpdateTime;comment:更新时间" json:"update_time"` // 更新时间
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time;not null;comment:删除时间" json:"delete_time"`                // 删除时间
}

// MarshalBinary 支持json序列化
func (m *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// UnmarshalBinary 支持json反序列化
func (m *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
