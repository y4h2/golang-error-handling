package myerr

import "github.com/pkg/errors"

type UserError interface {
	error
	UserError()
}

func ToUserError(err error) (UserError, bool) {
	if userErr, ok := errors.Cause(err).(UserError); ok {
		return userErr, true
	}

	return nil, false
}
