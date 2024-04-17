package dao

import "errors"

var (
	ErrUpdateNothing = errors.New("update nothing")
)

const (
	PublishVideoIgnore = -1
	PublishVideoYes    = 1
	PublishVideoNo     = 0
)
