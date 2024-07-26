package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"sophliteos/pkg/repo"
	"sophliteos/pkg/utils/i18n"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/patrickmn/go-cache"
)

const (
	acceptLanguage = "Accept-Language"
	authorization  = "Authorization"
	contentType    = "Content-Type"
	multipart      = "multipart/form-data"
	Pattern        = "2006-01-02 15:04:05"
)

var tokenCache *cache.Cache

func init() {
	tokenCache = cache.New(2*time.Hour, 5*time.Minute)
}

func Token(request *http.Request) string {
	return request.Header.Get(authorization)
}

func GetUser(token string) *repo.User {
	user, found := tokenCache.Get(token)
	if found {
		return user.(*repo.User)
	} else {
		user, _ := repo.QueryUserWithToken(token)
		return user
	}
}

func SetUser(token string, user *repo.User) {
	tokenCache.Set(token, user, 2*time.Hour)
}

func RemoveUser(token string) {
	tokenCache.Delete(token)
}

func IsMultiPartRequest(request *http.Request) bool {
	return strings.Contains(request.Header.Get(contentType), multipart)
}

func GetLang(request *http.Request) string {
	var lang = request.Header.Get(acceptLanguage)
	if len(lang) > 0 {
		return lang
	}
	return i18n.Zh
}

func NotBlank(request *http.Request, values ...string) error {
	lang := GetLang(request)
	for i := 0; i < len(values); i = i + 2 {
		value := values[i+1]
		if len(value) <= 0 {
			return errors.New(values[i] + i18n.GetString(lang, value))
		}
	}
	return nil
}

func Valid(request *http.Request, req interface{}) error {
	en := en.New()
	zh := zh.New()
	uni := ut.New(en, zh)
	trans, _ := uni.GetTranslator(GetLang(request))
	validate := validator.New()
	zhtrans.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errText := fmt.Sprintf("Invalid param!%v", removeStructName(errs.Translate(trans)))
		// logger.Error(errText)
		return errors.New(errText)
	}
	return err
}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}
