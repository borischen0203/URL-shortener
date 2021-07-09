package validators

// "url-shortener/dto"
// 	"github.com/gookit/validate"
// "url-shortener/dto"
import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func NotEmptyValidator(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 0
}

func IsValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", value)
	return match
}
