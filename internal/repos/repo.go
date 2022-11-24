package repos

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDB, NewRedis, NewMinio, NewTransaction, NewDBData, NewAccountRepo)
