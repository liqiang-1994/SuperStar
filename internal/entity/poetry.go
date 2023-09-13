package entity

type Poetry struct {
	Id                  int64  `gorm:"column:id;primary_key;serial"`
	Author              string `gorm:"column:author;size:100"`
	AuthorId            int64  `gorm:"column:author_id"`
	Style               string `gorm:"column:style;size:200"`
	Contents            string `gorm:"column:contents"`
	Rhythmic            string `gorm:"column:rhythmic"`
	Section             string `gorm:"column:section"`
	Notes               string `gorm:"column:notes"`
	Strains             string `gorm:"column:strains"`
	Favorite            int64  `gorm:"column:favorite"`
	Kind                string `gorm:"column:kind;index;size:100"`
	TraditionalContents string `gorm:"column:traditional_contents"`
	Dynasty             string `gorm:"column:dynasty;size:100"`
}
