package repos

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDB, NewTransaction, NewDBData, NewAccountRepo)
