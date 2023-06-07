package validation

type LogPayValidate struct {
	ID       string
	IP       string
	SdkReq   string
	ReqURL   string
	ReqParam string
	Err      error
	Remark   string
}
