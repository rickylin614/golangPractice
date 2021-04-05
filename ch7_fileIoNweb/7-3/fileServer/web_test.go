package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, req *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
	// return string(e)
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, req *http.Request) error {
	return testingUserError("user error!!")
}

func fileNotExistError(writer http.ResponseWriter, req *http.Request) error {
	return os.ErrNotExist
}

func permissionError(writer http.ResponseWriter, req *http.Request) error {
	return os.ErrPermission
}

func unknowError(writer http.ResponseWriter, req *http.Request) error {
	return errors.New("Unknow error")
}

func noError(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

var tests = []struct {
	handler appHandler
	code    int
	message string
}{
	// TODO: Add test cases.
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error!!"},
	{fileNotExistError, 404, "Not Found"},
	{permissionError, 403, "Forbidden"},
	{unknowError, 500, "Internal Server Error"},
	{noError, 200, ""},
}

func Test_errWrapper(t *testing.T) {
	// type args struct {
	// 	handler appHandler
	// }

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			// if got := errWrapper(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
			// 	// t.Errorf("errWrapper() = %v, want %v", got, tt.want)
			// }
			f := errWrapper(tt.handler)
			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
			f(resp, req)
			verifyResponse(resp.Result(), tt.code, tt.message, t)
		})
	}
}

func Test_errWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.handler) // f =  func(http.ResponseWriter, *http.Request)
		// HandlerFunc a interface perform ServeHTTP
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL) //執行get動作
		verifyResponse(resp, tt.code, tt.message, t)
	}

}

func verifyResponse(resp *http.Response, wantCode int, wantMsg string, t *testing.T) {
	body, _ := ioutil.ReadAll(resp.Body)
	b := strings.Trim(string(body), "\n")
	if resp.StatusCode != wantCode || b != wantMsg {
		t.Errorf("Error!! code: %d message: %s wantCode:%d ,wantMsg:%s", resp.StatusCode, b, wantCode, wantMsg)
	}
}
