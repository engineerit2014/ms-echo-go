package translators

import (
	"github.com/laironacosta/ms-echo-go/enums"
	"sync"
)

type DictLangInterface interface {
	GetMsg(code string) string
}

type DictLang map[string]string

type I18nInterface interface {
	GetDictLang(lang string) DictLang
}

type I18n struct {
	i18n        map[string]DictLang
	defaultLang string
}

var (
	i18n I18n
	once sync.Once
)

func NewI18n(defaultLang string) I18n {
	once.Do(func() {
		i18n = initI18n(defaultLang)
	})

	return i18n
}

func initI18n(defaultLang string) I18n {
	translators := make(map[string]DictLang, 2)
	translators[enums.SpanishLang] = DictSpanish
	translators[enums.EnglishLang] = DictEnglish

	return I18n{translators, defaultLang}
}

func (t I18n) GetDictLang(lang string) DictLang {
	if val, ok := t.i18n[lang]; ok {
		return val
	}

	return t.i18n[t.defaultLang]
}

func (d DictLang) GetMsg(code string) string {
	return d[code]
}
