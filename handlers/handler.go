package handlers

import "github.com/google/wire"

var Provider = wire.NewSet(NewAccountHandler, NewLoginHandler,
	NewPoemHandler, NewCircleHandler, NewTagHandler, NewStorageHandler)
