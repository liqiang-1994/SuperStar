package entity

import "time"

type Comment struct {
	Id           int64     `gorm:"column:id;primary_key" json:"id"`
	Content      string    `gorm:"column:content" json:"content"`
	Up           int64     `gorm:"column:up" json:"up"`
	IpAddress    string    `gorm:"column:ip_address;size:50" json:"ipAddress"`
	CreateId     int64     `gorm:"column:create_id" json:"createId"`
	CreateName   string    `gorm:"column:create_name;size:100" json:"createName"`
	CreateAvatar string    `gorm:"column:create_avatar;size:200" json:"createAvatar"`
	ResourceId   int64     `gorm:"column:resource_id;index" json:"resourceId"`
	ParentId     int64     `gorm:"column:parent_id;index" json:"parentId"`
	ResourceType int64     `gorm:"column:resource_type" json:"resourceType"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
}
