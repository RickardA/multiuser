package repository

import "errors"

var ErrCouldNotGetObjectID = errors.New("could not convert object id to domain type")
var ErrIDNotFound = errors.New("could not find with matching id")
