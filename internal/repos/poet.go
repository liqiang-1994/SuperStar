package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewPoetRepo(db *Data) services.PoetRepo {
	return &poetRepo{db: db}
}

type poetRepo struct {
	db *Data
}

func (p *poetRepo) QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Poet, error) {
	var poets []entity.Poet
	sql := p.db.DB(ctx).Offset((params.PageNum - 1) * params.PageSize).Limit(params.PageSize)
	if params.KeyWord != "" {
		sql = sql.Where("poet_name LIKE ?", "%"+params.KeyWord+"%")
	}
	if params.Dynasty != "" {
		sql = sql.Where("dynasty = ?", params.Dynasty)
	}
	err := sql.Order("id asc").Find(&poets).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return poets, err
}
