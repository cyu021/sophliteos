package service

import (
	"sophliteos/pkg/repo"

	"net/http"

	"strings"
	"time"
)

func SaveOptLog(request *http.Request, operationName, content string) {

	user := GetUser(Token(request))

	if operationName == "登录" {
		repo.SaveLog(repo.OptLog{
			UserName:         "admin",
			CreatedTime:      time.Now(),
			OperationType:    strings.Split(request.RequestURI, "?")[0],
			OperationContent: content,
			OperationIP:      request.RemoteAddr[0:strings.LastIndex(request.RemoteAddr, ":")],
			OperationFunc:    operationName,
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
