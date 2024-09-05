package entity

import (
	"github.com/joomcode/errorx"
)

var errors = errorx.NewNamespace("apperr")

var (
	ReqErr  = errors.NewType("BADREQUEST_ERR")
	AuthErr = errors.NewType("AUTHENTICATION_ERR")
	ValErr  = errors.NewType("VALIDATION_ERR")
)

var (
	Code = errorx.RegisterProperty("status_code")
)
