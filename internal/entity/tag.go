package entity

import "time"

type Tag struct {
	Id         int64     `gorm:"column:id;primary_key" json:"id"`
	Name       string    `gorm:"column:name;size:100" json:"name"`
	CreateId   int64     `gorm:"column:create_id" json:"createId"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateId   int64     `gorm:"column:update_id" json:"updateId"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

type TagRel struct {
	Id           int64     `gorm:"column:id;primary_key" json:"id"`
	TagId        int64     `gorm:"column:tag_id;index" json:"tagId"`
	ResourceId   int64     `gorm:"column:resource_id;index" json:"resourceId"`
	ResourceType int64     `gorm:"column:resource_type;" json:"resourceType"`
	CreateId     int64     `gorm:"column:create_id" json:"createId"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
}
