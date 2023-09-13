package entity

type Idiom struct {
	Id           int64  `gorm:"column:id;primary_key;serial" json:"id"`
	Word         string `gorm:"column:word;index;size:100" json:"word"`
	Pinyin       string `gorm:"column:pinyin;size:100" json:"pinyin"`
	Example      string `gorm:"column:example;" json:"example"`
	Explanation  string `gorm:"column:explanation;" json:"explanation"`
	Derivation   string `gorm:"column:derivation;" json:"derivation"`
	Abbreviation string `gorm:"column:abbreviation;index" json:"abbreviation"`
}
