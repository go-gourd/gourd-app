// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int32                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserName   string                `gorm:"column:user_name;not null" json:"user_name"`
	Mobile     string                `gorm:"column:mobile;not null" json:"mobile"`
	Password   string                `gorm:"column:password;not null" json:"password"`
	CreateTime int32                 `gorm:"column:create_time;type:int unsigned;default:0;autoCreateTime" json:"create_time"`
	UpdateTime int32                 `gorm:"column:update_time;type:int unsigned;default:0;autoUpdateTime" json:"update_time"`
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time;type:int unsigned;default:0;autoCreateTime" json:"delete_time"`
	Extend     string                `gorm:"column:extend" json:"extend"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}