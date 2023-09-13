package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewSayingRepo(db *Data) services.SayingRepo {
	return &sayingRepo{db: db}
}

type sayingRepo struct {
	db *Data
}

func (p *sayingRepo) QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Saying, error) {
	var sayings []entity.Saying
	sql := p.db.DB(ctx).Offset((params.PageNum - 1) * params.PageSize).Limit(params.PageSize)
	if params.KeyWord != "" {
		sql = sql.Where("riddle LIKE ?", "%"+params.KeyWord+"%").
			Or("answer LIKE ?", "%"+params.KeyWord+"%")
	}
	err := sql.Order("id asc").Find(&sayings).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return sayings, err
}
