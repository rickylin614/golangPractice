package myHandle

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//a new Error defined
type UserError string

//implements Error() method by UserError
func (e UserError) Error() string {
	return e.Message()
}

//implements Message() method by UserError
func (e UserError) Message() string {
	return string(e)
}

// define a http client and return file content (listener)
//  	e.g. HandleFileList() ddd
func HandleFileList(writer http.ResponseWriter, req *http.Request) error {

	if strings.Count(req.URL.Path, prefix) == 0 {
		// return errors.New("Path must start with " + prefix)
		return UserError("Path must start with " + prefix)
	}
	path := req.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// log.Println(err)
		// http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		// log.Println(err)
		return err
	}
	_, err = writer.Write(all)
	if err != nil {
		// log.Println(err)
		return err
	}
	return nil
}
