package model

import "SuperStar/internal/entity"

type CreateCircleReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Signature   string `json:"signature"`
}

type CircleDetail struct {
	Circle entity.Circle   `json:"circle"`
	Follow []entity.TagRel `json:"follow"`
}
