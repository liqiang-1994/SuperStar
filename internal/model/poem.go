package model

type PoemListReq struct {
	Type     int    `json:"type"` // 0:诗人 1:诗词 2:
	KeyWord  string `json:"keyWord"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Dynasty  string `json:"dynasty"`
}
