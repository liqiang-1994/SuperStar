package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewIdiomRepo(db *Data) services.IdiomRepo {
	return &idiomRepo{db: db}
}

type idiomRepo struct {
	db *Data
}

func (p *idiomRepo) QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Idiom, error) {
	var idiom []entity.Idiom
	sql := p.db.DB(ctx).Offset((params.PageNum - 1) * params.PageSize).Limit(params.PageSize)
	if params.KeyWord != "" {
		sql = sql.Where("word LIKE ?", "%"+params.KeyWord+"%").
			Or("explanation LIKE ?", "%"+params.KeyWord+"%").
			Or("derivation LIKE ?", "%"+params.KeyWord+"%").
			Or("abbreviation LIKE ?", "%"+params.KeyWord+"%")
	}

	err := sql.Order("id asc").Find(&idiom).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return idiom, err
}
