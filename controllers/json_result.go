package controllers

import (
	"fmt"
	"github.com/albrow/go-data-parser"
	"runtime"
)

type JSONResult struct {
	Short    string                 `json:",omitempty"` // a short explanation of the response (usually one or two words). for internal use only
	Messages []string               `json:",omitempty"` // any non-error messages which should be displayed to the user
	Errors   []string               `json:",omitempty"` // any errors that may have occured with the request and should be displayed to the user
	Keys     []string               `json:",omitempty"` // keys corresponding to the errors in the case of validation errors
	Data     map[string]interface{} `json:",omitempty"` // formatted json data corresponding to the type of request. (e.g. Users.Index will reutrn an array of users)
	From     string                 `json:",omitempty"` // for errors, the file and line number at which the error originated
}

func newJSONError(short string, err error) JSONResult {
	_, file, line, _ := runtime.Caller(1)
	return JSONResult{
		Short:  short,
		Errors: []string{err.Error()},
		From:   fmt.Sprintf("%s:%d", file, line),
	}
}

func newJSONValidationError(val *data.Validator) JSONResult {
	return JSONResult{
		Short:  "validationError",
		Keys:   val.Fields(),
		Errors: val.Messages(),
	}
}

func newJSONOk() JSONResult {
	return JSONResult{Short: "ok"}
}

func newJSONMessage(msg string) JSONResult {
	return JSONResult{
		Short:    "ok",
		Messages: []string{msg},
	}
}

func newJSONData(data map[string]interface{}) JSONResult {
	return JSONResult{
		Short: "ok",
		Data:  data,
	}
}

func newJSONDataAndMessage(data map[string]interface{}, msg string) JSONResult {
	return JSONResult{
		Short:    "ok",
		Data:     data,
		Messages: []string{msg},
	}
}
