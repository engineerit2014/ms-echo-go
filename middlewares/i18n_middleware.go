package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/laironacosta/ms-echo-go/enums"
	"github.com/laironacosta/ms-echo-go/translators"
)

type (
	I18nMiddlewareInterface interface {
		HandlerError(next echo.HandlerFunc) echo.HandlerFunc
	}

	I18nMiddleware struct {
		i18n translators.I18nInterface
	}
)

func NewI18nMiddleware(i18n translators.I18nInterface) I18nMiddlewareInterface {
	return &I18nMiddleware{
		i18n,
	}
}

func (i *I18nMiddleware) HandlerError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Request().Header.Get(enums.LangHeader)
		dictLang := i.i18n.GetDictLang(lang)

		c.SetRequest(c.Request().WithContext(setDictLangInContext(c.Request().Context(), dictLang)))
		return next(c)
	}
}

func setDictLangInContext(ctx context.Context, dictLang translators.DictLang) context.Context {
	return context.WithValue(ctx, enums.I18nKey, dictLang)
}
