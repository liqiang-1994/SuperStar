package entity

import "time"

type Account struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type TAccount struct {
	Id         int64     `gorm:"column:id;" json:"id"`
	UserName   string    `gorm:"column:user_name;" json:"user_name"`
	Age        int       `gorm:"column:age;" json:"age"`
	Address    string    `gorm:"column:address;" json:"address"`
	Phone      string    `gorm:"column:phone;" json:"phone"`
	Gender     int8      `gorm:"column:gender;" json:"gender"`
	Pass       string    `gorm:"column:pass;" json:"pass"`
	Signature  string    `gorm:"column:signature;" json:"signature"`
	Avatar     string    `gorm:"column:avatar;" json:"avatar"`
	Status     int       `gorm:"column:status;" json:"status"`
	Nation     string    `gorm:"column:nation;" json:"nation"`
	CreateTime time.Time `gorm:"column:create_time;" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;" json:"update_time"`
}
