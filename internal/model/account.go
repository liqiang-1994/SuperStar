package model

import (
	"SuperStar/internal/entity"
	"time"
)

type LoginReq struct {
	UserName  string `json:"userName"`
	CheckCode string `json:"checkCode"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Id        int64     `json:"id,omitempty"`
	UserName  string    `json:"name,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *entity.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}
}
