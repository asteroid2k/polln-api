package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog/log"
)

type AppValidator struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func NewValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	en := en.New()

	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	return validate, trans
}

func (val *AppValidator) ValidateStruct(data any) ([]ValidationError, bool) {
	err := val.Validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Fatal().Msg("invalid validation error")
			return []ValidationError{}, false
		}
		var vErrors []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			vErrors = append(vErrors, ValidationError{Field: err.Field(), Message: err.Translate(val.Translator)})
		}

		return vErrors, false
	}
	return nil, true
}
