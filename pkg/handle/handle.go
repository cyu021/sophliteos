package handle

import (
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/utils/i18n"
)

func Result(code int, result interface{}, msg string) response.Result {
	return response.Result{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

func Ok() response.Result {
	return Result(buserr.Ok, nil, "ok")
}
func OkWithMsg(msg string) response.Result {
	return Result(buserr.Ok, nil, msg)
}

func Success(result interface{}) response.Result {
	return Result(buserr.Ok, result, "ok")
}

func Error(error string) response.Result {
	return response.Result{
		Code: buserr.Err,
		Msg:  error,
	}
}

func Fail(code int, msg string) response.Result {
	return Result(code, nil, msg)
}

func FailWithMsg(code int, msg string) response.Result {
	return Result(code, nil, msg)
}

func HandleError(err error, codes ...interface{}) {
	if err != nil {
		if len(codes) > 0 {
			logger.Error("%v\n%s", codes, err.Error())
		} else {
			// panic(err.Error())
		}
	}
}

func Handle(ssmResult dto.SsmResult, code int) response.Result {
	var result response.Result
	if ssmResult.ErrorCode != buserr.Ok {
		logger.Error("%s %s %s", i18n.GetString(i18n.Zh, code), ssmResult.ErrorCode, ssmResult.ErrorMessage)
		panic(code)
	} else {
		result.Code = ssmResult.Code
		result.Msg = ssmResult.Msg
	}
	return result
}

func HandleResult(ssmResult dto.SsmResult, err error, code int) response.Result {
	HandleError(err, code)
	return Handle(ssmResult, code)
}
