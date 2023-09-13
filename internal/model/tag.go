package model

type CreateTagReq struct {
	Name string `json:"name"`
}

type CreateTagRelReq struct {
	TagId        int64 `json:"tagId"`
	ResourceId   int64 `json:"resourceId"`
	ResourceType int64 `json:"resourceType"`
}
