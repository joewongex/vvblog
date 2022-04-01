package service

import (
	"vvblog/errors/vcode"
	"vvblog/errors/verror"

	"github.com/go-playground/validator/v10"
)

type validation struct {
	validate *validator.Validate
}

var V = validation{
	validate: validator.New(),
}

func (v *validation) CheckPositiveInt(num interface{}, name string) {
	if err := v.validate.Var(num, "required,min=1"); err != nil {
		panic(verror.NewCode(vcode.CodeInvalidParameter, name+"必须为正整数"))
	}
}
