package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pedropmartiniano/mvc-golang/src/configuration/restErr"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()

		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *restErr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return restErr.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errCauses := []restErr.Causes{}

		for _, e := range jsonValidationError {
			cause := restErr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errCauses = append(errCauses, cause)
		}

		return restErr.NewBadRequestValidationError("Some fields are invalid", errCauses)
	}

	return restErr.NewBadRequestError("Error trying to convert fields")
}
