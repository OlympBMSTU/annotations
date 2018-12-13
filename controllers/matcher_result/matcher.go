package matcher_result

import (
	"encoding/json"
	"net/http"

	auth_result "github.com/OlympBMSTU/annotations/auth/result"
	http_res "github.com/OlympBMSTU/annotations/controllers/http_result"
	fs_result "github.com/OlympBMSTU/annotations/fstorage/result"
	root_result "github.com/OlympBMSTU/annotations/result"
	"github.com/OlympBMSTU/annotations/views/output"

)

func fillResult(info ResultInfo, body interface{}) http_res.HttpResult {
	var jsonRes output.ResultView
	jsonRes.SetData(body)
	jsonRes.SetStatus(info.Status)
	jsonRes.SetMessage(info.Message)

	val, err := json.Marshal(jsonRes)
	code := info.HttpCode
	var outHttpRes http_res.HttpResult

	if err != nil {
		code = http.StatusInternalServerError
	} else {
		outHttpRes.SetBody(val)
	}

	outHttpRes.SetStatus(code)
	return outHttpRes
}

// func MatchDbResult(res result.Result) http_res.HttpResult {
// 	var jsonRes output.ResultView
// 	info := mapHttpDbStatuses[res.GetStatus().GetCode()]
// 	if res.IsError() {
// 		jsonRes.SetData(nil)
// 	} else {
// 		jsonRes.SetData(res.GetData())
// 	}
// 	jsonRes.SetStatus(info.Status)
// 	jsonRes.SetMessage(info.Message)

// 	val, err := json.Marshal(jsonRes)
// 	code := info.HttpCode
// 	var outHttpRes http_res.HttpResult

// 	if err != nil {
// 		code = http.StatusInternalServerError
// 	} else {
// 		outHttpRes.SetBody(val)
// 	}

// 	outHttpRes.SetStatus(code)
// 	return outHttpRes
// }

func MatchResult(res root_result.Result) http_res.HttpResult {
	var infoRes ResultInfo
	var bodyData interface{} = nil
	// bodyData = nil/

	switch res.(type) {
	case fs_result.FSResult:
		infoRes = getAssociatedFsInfo(res)
		// return MatchFSResult(res)
	case auth_result.AuthResult:
		infoRes = getAssociatedAuthInfo(res)
		// return MatchAuthResult(res)
	default:
		// coorect this
		return http_res.ResultInernalSreverError()
	}

	return fillResult(infoRes, bodyData)
}
