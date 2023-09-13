package entity

type Saying struct {
	Id     int64  `gorm:"column:id;primary_key;serial" json:"id"`
	Riddle string `gorm:"column:riddle;index" json:"riddle"`
	Answer string `gorm:"column:answer;" json:"answer"`
}
