package repos

import (
	"SuperStar/internal/utils"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDB, NewRedis, NewMinio, NewGID, utils.NewSmsClient,
	NewTransaction, NewDBData, NewAccountRepo, NewCircleRepo, NewTagRepo,
	NewIdiomRepo, NewPoetRepo, NewPoetryRepo, NewSayingRepo)
