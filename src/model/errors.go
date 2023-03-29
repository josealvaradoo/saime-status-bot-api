package model

import "errors"

var (
	ErrStatusNotExists       = errors.New("no exists any status")
	ErrStatusCanNotBeUpdated = errors.New("cannot update status")
)
