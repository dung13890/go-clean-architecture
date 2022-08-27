package validate

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	uniTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var errTranslateNotFound = errors.New("translator not found")

// CustomValidate is struct used to validate
type CustomValidate struct {
	validate *validator.Validate
}

// NewValidate is function that return new validate
func NewValidate() *CustomValidate {
	return &CustomValidate{validate: validator.New()}
}

// Validate is used to validate interface
func (v *CustomValidate) Validate(u interface{}) error {
	translator := en.New()
	uni := uniTranslator.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		return errTranslateNotFound
	}

	_ = v.validate.RegisterTranslation("required", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("email", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("email", "{0} invalid email!", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("number", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("number", "{0} is not number!", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("number", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("min", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("min", "{0} is less than min!", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("max", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("max", "{0} is greater than max!", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())

		return t
	})

	if err := v.validate.Struct(u); err != nil {
		errArr := []string{}
		validationErrs := validator.ValidationErrors{}
		if errors.As(err, &validationErrs) {
			for _, e := range validationErrs {
				errArr = append(errArr, e.Translate(trans))
			}
		}

		errStr := strings.Join(errArr, ";")
		//nolint:goerr113
		return errors.New(errStr)
	}

	return nil
}
