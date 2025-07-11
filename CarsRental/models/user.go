package models

type User struct {
	ID       int
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}
