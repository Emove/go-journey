package model

type User struct {
	*Model
	Name   string `gorm:"column:name"`
	Gender int    `gorm:"column:gender" parser:"enum:1-male,2-female"`
	Addr   string `gorm:"column:addr"`
	Age    *int   `gorm:"column:age"`
}

func (User) TableName() string {
	return "user"
}
