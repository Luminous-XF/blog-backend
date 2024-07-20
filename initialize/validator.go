package initialize

import "blog-backend/pkg/validator"

func initValidator() {
	if validate := validator.InitValidator(); validate == nil {
		panic("Validator initialization failed")
	}
}
