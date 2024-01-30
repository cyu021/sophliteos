package system

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/repo"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
)

type PasswordApi struct{}

func (b *PasswordApi) PasswordMod(c *gin.Context) {
	req := response.PasswordModify{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &req)

	err := service.Valid(c.Request, req)
	if err != nil {
		// panic(err)
		errStr := fmt.Sprintf("%v", err)
		logger.Error("error: %s", errStr)
		c.JSON(http.StatusUnprocessableEntity, handle.FailWithMsg(1, errStr))
		return
	}
	code := check(req.NewPassword)
	if code != buserr.Ok {
		c.JSON(http.StatusOK, handle.Fail(code, "密码无效"))
		return
	}
	user := service.GetUser(service.Token(c.Request))
	if user.Password == md5Value(md5Value(req.Password)) {
		user.Password = md5Value(md5Value(req.NewPassword))
		repo.SaveUser(user)
		c.JSON(http.StatusOK, handle.Ok())
		return
	} else {
		logger.Error("旧密码错误")
		c.JSON(http.StatusOK, handle.Fail(buserr.PwdNotEqErr, "旧密码错误"))
		return
	}

}

func md5Value(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 长度8-16位，大小写字母，数字，特征字符至少3种
func check(pwd string) buserr.Code {
	if len(pwd) < 8 {
		return buserr.PwdValidErr
	}
	if len(pwd) > 16 {
		return buserr.PwdValidErr
	}
	var level = 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	if level < 3 {
		return buserr.PwdValidErr
	}
	return buserr.Ok
}
