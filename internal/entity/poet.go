package entity

type Poet struct {
	Id          int64  `gorm:"column:id;primary_key;serial" json:"id"`
	PoetName    string `gorm:"column:poet_name;index;size:100" json:"poetName"`
	Dynasty     string `gorm:"column:dynasty;size:100" json:"dynasty"`
	Description string `gorm:"column:description;" json:"description"`
	Style       string `gorm:"column:style;size:200" json:"style"`
	Favorite    int64  `gorm:"column:favorite;" json:"favorite"`
}
