package service

import (
	"sophliteos/pkg/repo"

	"net/http"

	"strings"
	"time"
)

func SaveOptLog(request *http.Request, operationName, content string) {

	user := GetUser(Token(request))

	if strings.HasPrefix(operationName, "Login-") { //"登录" {
		eles := strings.Split(operationName, "-")
		repo.SaveLog(repo.OptLog{
			UserName:         eles[1], //"admin",
			CreatedTime:      time.Now(),
			OperationType:    strings.Split(request.RequestURI, "?")[0],
			OperationContent: content,
			OperationIP:      request.RemoteAddr[0:strings.LastIndex(request.RemoteAddr, ":")],
			OperationFunc:    eles[0],
		})
		return
	}

	if user == nil {
		return
	}
	repo.SaveLog(repo.OptLog{
		UserName:         user.UserName,
		CreatedTime:      time.Now(),
		OperationType:    strings.Split(request.RequestURI, "?")[0],
		OperationContent: content,
		OperationIP:      request.RemoteAddr[0:strings.LastIndex(request.RemoteAddr, ":")],
		OperationFunc:    operationName,
	})

}
