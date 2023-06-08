package validate

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	et "github.com/go-playground/validator/v10/translations/en"
)

var (
	errTranslateNotFound = errors.New("translator not found")
	translationOverride  = map[string]string{}
)

// CustomValidate is struct used to validate
type CustomValidate struct {
	validate *validator.Validate
}

// NewValidate is function that return new validate
func NewValidate() *CustomValidate {
	return &CustomValidate{validate: validator.New()}
}

// RegisterAlias is used to register list of alias
func (v *CustomValidate) RegisterAlias(list map[string]string) {
	for alias, tags := range list {
		v.validate.RegisterAlias(alias, tags)
	}
}

// RegisterTranslationOverride is used to register list of translation override
func (*CustomValidate) RegisterTranslationOverride(list map[string]string) {
	for tag, translation := range list {
		translationOverride[tag] = translation
	}
}

// Validate is used to validate interface
func (v *CustomValidate) Validate(u interface{}) error {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		return errTranslateNotFound
	}
	_ = et.RegisterDefaultTranslations(v.validate, trans)

	_ = v.validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} invalid email!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("number", trans, func(ut ut.Translator) error {
		return ut.Add("number", "{0} is not number!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("number", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} is less than min!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())

		return t
	})

	_ = v.validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} is greater than max!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())

		return t
	})

	if len(translationOverride) > 0 {
		for tag, translation := range translationOverride {
			_ = v.validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
				return ut.Add(tag, translation, true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(tag, fe.Field(), fe.Param())

				return t
			})
		}
	}

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
