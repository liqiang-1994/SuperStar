package services

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"github.com/bwmarrin/snowflake"
	"time"
)

type CircleRepo interface {
	CreateCircle(ctx context.Context, circle *entity.Circle) error
	CreateCircleRel(ctx context.Context, circle *entity.CircleUserRel) error
	QueryCircleById(ctx context.Context, circleId string) (*entity.Circle, error)
}

type CircleService struct {
	circleRepo CircleRepo
	tm         Transaction
	GID        *snowflake.Node
}

func NewCircleService(circle CircleRepo, tm Transaction, GID *snowflake.Node) *CircleService {
	return &CircleService{circleRepo: circle, tm: tm, GID: GID}
}

func (s *CircleService) QueryCircleDetail(ctx context.Context, circleId string) (*model.CircleDetail, error) {
	circle, err := s.circleRepo.QueryCircleById(ctx, circleId)
	if err != nil {
		return nil, err
	}
	if circle == nil {
		return nil, nil
	}
	detail := &model.CircleDetail{
		Circle: *circle,
	}
	return detail, nil
}

func (s *CircleService) CreateCircle(ctx context.Context, m *model.CreateCircleReq, user *model.UserResponse) (bool, error) {
	circleId := s.GID.Generate().Int64()
	circle := &entity.Circle{
		Id:          circleId,
		CircleName:  m.Name,
		Description: m.Description,
		Signature:   m.Signature,
		Number:      1,
		CreateId:    user.Id,
		CreateTime:  time.Now(),
	}
	circleRel := &entity.CircleUserRel{
		Id:         s.GID.Generate().Int64(),
		CircleId:   circleId,
		UserId:     user.Id,
		UserName:   user.UserName,
		UserAvatar: user.Avatar,
		Status:     1,
		CreateTime: time.Now(),
	}
	err := s.tm.ExecTx(ctx, func(ctx context.Context) error {
		if err := s.circleRepo.CreateCircle(ctx, circle); err != nil {
			return err
		}
		if err := s.circleRepo.CreateCircleRel(ctx, circleRel); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
