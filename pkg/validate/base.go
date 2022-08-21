package validate

import (
	"bytes"
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
		return ut.Add("min", "{0} is less than min", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("max", trans, func(ut uniTranslator.Translator) error {
		return ut.Add("max", "{0} is greater than max", true)
	}, func(ut uniTranslator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())

		return t
	})

	err := v.validate.Struct(u)
	//nolint:errorlint,forcetypeassert,goerr113
	if err != nil {
		var errStr bytes.Buffer
		for index, e := range err.(validator.ValidationErrors) {
			if index+1 == len(err.(validator.ValidationErrors)) {
				errStr.WriteString(formatString(e.Translate(trans)))
			} else {
				errStr.WriteString(formatString(e.Translate(trans)) + ",")
			}
		}

		return errors.New(errStr.String())
	}

	return nil
}

func formatString(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, "!", ""))
}
