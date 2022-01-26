package entity

import "errors"

var (
	ErrSomethingWentWrong  = errors.New("something went wrong")
	ErrNoDataMatched       = errors.New("no record found")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
