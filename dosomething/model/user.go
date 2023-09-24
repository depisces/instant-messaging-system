package model

type User struct {
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Age  string `gorm:"default:0"`
}

func (User) TableName() string {
	return "user"
}
