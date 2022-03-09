package validators

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

func ValidatePassword (fl validator.FieldLevel) bool {
	//TODO: Modificare la regex, ?= non supportato in go
	pattern := `^(?=.*[a-z])(?=.*[A-Z])^(?=.*[0-9])(?=.*[.,:;!?#@&_\\-])[a-zA-Z0-9.,:;!?#@&_\\-]{8,}$`
	ok, err := regexp.MatchString(pattern, fl.Field().String())
	if err != nil {
		log.Fatalln(err)
	}

	return ok
}
