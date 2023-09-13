package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewCircleRepo(db *Data) services.CircleRepo {
	return &circleRepo{db: db}
}

type circleRepo struct {
	db *Data
}

func (c *circleRepo) QueryCircleById(ctx context.Context, circleId string) (*entity.Circle, error) {
	circle := &entity.Circle{}
	err := c.db.DB(ctx).Where("id = ?", circleId).First(circle).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return circle, nil
}

func (c *circleRepo) CreateCircleRel(ctx context.Context, circle *entity.CircleUserRel) error {
	return c.db.DB(ctx).Create(circle).Error
}

func (c *circleRepo) CreateCircle(ctx context.Context, circle *entity.Circle) error {
	return c.db.DB(ctx).Create(circle).Error
}
