package entity

import "time"

type Circle struct {
	Id          int64     `gorm:"column:id;primary_key" json:"id"`
	CircleName  string    `gorm:"column:circle_name;size:100" json:"circleName"`
	Number      int64     `gorm:"column:number;" json:"number"`
	PostNumber  int64     `gorm:"column:post_number;" json:"postNumber"`
	Hot         int32     `gorm:"column:hot;" json:"hot"`
	Description string    `gorm:"column:description;" json:"description"`
	Signature   string    `gorm:"column:signature;size:100" json:"signature"`
	TypeName    string    `gorm:"column:type_name;size:100" json:"typeName"`
	Avatar      string    `gorm:"column:avatar;size:100" json:"avatar"`
	CreateId    int64     `gorm:"column:create_id;" json:"createId"`
	CreateTime  time.Time `gorm:"column:create_time;" json:"createTime"`
	UpdateId    int64     `gorm:"column:update_id;" json:"updateId"`
	UpdateTime  time.Time `gorm:"column:update_time;" json:"updateTime"`
}

type CircleUserRel struct {
	Id         int64     `gorm:"column:id;primary_key" json:"id"`
	CircleId   int64     `gorm:"column:circle_id;index" json:"circleId"`
	UserId     int64     `gorm:"column:user_id;index" json:"userId"`
	UserName   string    `gorm:"column:user_name;size:100" json:"userName"`
	UserAvatar string    `gorm:"column:user_avatar;size:100" json:"userAvatar"`
	Status     int8      `gorm:"column:status;" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;" json:"createTime"`
}
