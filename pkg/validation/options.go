package validation

import (
	"reflect"
	"strings"

	enLocales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Option validation option
type Option interface {
	apply(*option)
}

// option implement
type option struct {
	validator *validator.Validate
	uni       *ut.UniversalTranslator
	trans     *ut.Translator
}

type optionFn func(*option)

func (optFn optionFn) apply(opt *option) {
	optFn(opt)
}

// WithValidator set validator
func WithValidator(v *validator.Validate) Option {
	return optionFn(func(opt *option) {
		opt.validator = v
	})
}

// WithUniversalTranslator set UniversalTranslator
func WithUniversalTranslator(uni *ut.UniversalTranslator) Option {
	return optionFn(func(opt *option) {
		opt.uni = uni
	})
}

// WithTranslator set Translator
func WithTranslator(trans *ut.Translator) Option {
	return optionFn(func(opt *option) {
		opt.trans = trans
	})
}

func getDefaultOption() *option {
	v := validator.New()

	translator := enLocales.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil
	}

	if err := enTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonTag := fld.Tag.Get("json")
		if jsonTag == "" {
			return fld.Name
		}

		name := strings.SplitN(jsonTag, ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &option{
		validator: v,
		uni:       uni,
		trans:     &trans,
	}
}

func getOption(opts ...Option) *option {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}
