package entity

import "time"

type User struct {
	Id         int64     `gorm:"column:id;primary_key;" json:"id"`
	UserName   string    `gorm:"column:user_name;index:user_name,unique;size:100" json:"userName"`
	Age        int8      `gorm:"column:age;" json:"age"`
	Address    string    `gorm:"column:address;" json:"address"`
	Phone      string    `gorm:"column:phone;index:phone,unique" json:"phone"`
	Gender     int8      `gorm:"column:gender;" json:"gender"`
	Password   string    `gorm:"column:pass;" json:"pass"`
	Signature  string    `gorm:"column:signature;" json:"signature"`
	Avatar     string    `gorm:"column:avatar;" json:"avatar"`
	Status     int8      `gorm:"column:status;" json:"status"`
	Follow     int64     `gorm:"column:follow;" json:"follow"`
	Watch      int64     `gorm:"column:watch;" json:"watch"`
	Up         int64     `gorm:"column:up;" json:"up"`
	LoginType  int8      `gorm:"column:login_type;" json:"loginType"`
	Nation     string    `gorm:"column:nation;" json:"nation"`
	CreateTime time.Time `gorm:"column:create_time;" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;" json:"updateTime"`
}

type UserRel struct {
	Id           int64     `gorm:"column:id;primary_key" json:"id"`
	UserId       int64     `gorm:"column:user_id;index" json:"userId"`
	UserName     string    `gorm:"column:user_name;size:100" json:"userName"`
	UserAvatar   string    `gorm:"column:user_avatar;size:100" json:"userAvatar"`
	FollowId     int64     `gorm:"column:follow_id;index" json:"followId"`
	FollowName   string    `gorm:"column:follow_name;size:100" json:"followName"`
	FollowAvatar string    `gorm:"column:follow_avatar;size:100" json:"followAvatar"`
	Status       int8      `gorm:"column:status;" json:"status"`
	CreateTime   time.Time `gorm:"column:create_time;" json:"createTime"`
}
