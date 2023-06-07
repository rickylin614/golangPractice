package validation

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func ParsePayValidateResponse(resp string, validateLog *LogPayValidate) (jsonStr string, err error) {
	base64DecodeStr, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		validateLog.Remark = "First part base64 decodeString failed."
		return "", err
	}

	urlDecodeStr, err := url.QueryUnescape(string(base64DecodeStr))
	if err != nil {
		validateLog.Remark = "First part queryUnescape failed."
		return "", err
	}
	respData := struct {
		Sc    string `json:"Sc"`
		MerNo string `json:"MerNo"`
		Data  string `json:"data"`
	}{}
	err = json.Unmarshal([]byte(urlDecodeStr), &respData)
	if err != nil {
		return "", err
	}
	base64DecodeStr, err = base64.StdEncoding.DecodeString(respData.Data)
	if err != nil {
		validateLog.Remark = "Second part base64 decodeString failed."
		return "", err
	}
	urlDecodeStr, err = url.QueryUnescape(string(base64DecodeStr))
	if err != nil {
		validateLog.Remark = "Second part queryUnescape failed."
		return "", err
	}
	return urlDecodeStr, nil
}
