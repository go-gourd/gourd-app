package model

type User struct {
	ID       uint
	Nickname string
	Phone    string
}

func (User) TableName() string {
	return "user"
}
