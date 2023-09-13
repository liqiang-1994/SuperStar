package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewPoetryRepo(db *Data) services.PoetryRepo {
	return &poetryRepo{db: db}
}

type poetryRepo struct {
	db *Data
}

func (p poetryRepo) QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Poetry, error) {
	var poetry []entity.Poetry
	sql := p.db.DB(ctx).Offset((params.PageNum - 1) * params.PageSize).Limit(params.PageSize)
	if params.KeyWord != "" {
		sql = sql.Where("author = ?", params.KeyWord).
			Or("contents LIKE ?", "%"+params.KeyWord+"%").
			Or("rhythmic LIKE ?", "%"+params.KeyWord+"%")
	}

	err := sql.Order("id asc").Find(&poetry).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return poetry, err
}
