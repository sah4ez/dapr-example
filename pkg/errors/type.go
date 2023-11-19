package errors

import (
	"fmt"
	"net/http"

	"github.com/seniorGolang/json"
)

var errorsMap = make(map[string]Error)

type Error struct {
	publicErr  *Error
	Tr         string      `json:"tr"`
	Msg        string      `json:"msg"`
	Cause      string      `json:"cause,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode *int        `json:"statusCode,omitempty"`
}

func new(msg, trKey string) Error {
	errorsMap[trKey] = Error{Msg: msg, Tr: trKey}
	return errorsMap[trKey]
}

func (e Error) MarshalJSON() ([]byte, error) {

	type jError Error
	if e.publicErr != nil {
		return json.Marshal((jError)(*e.publicErr))
	}
	return json.Marshal((jError)(e))
}

func (e Error) WithPublic(err Error) Error {
	e.publicErr = &err
	return e
}

func (e Error) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e Error) WithCode(code int) Error {
	e.StatusCode = &code
	return e
}

func (e Error) Code() int {

	if e.publicErr != nil {
		return e.publicErr.Code()
	}
	if e.StatusCode != nil {
		return *e.StatusCode
	}
	return http.StatusOK
}

func (e Error) Error() (errStr string) {

	if e.publicErr != nil {
		return e.publicErr.Error()
	}
	if e.Cause != "" {
		errStr = fmt.Sprintf("%s cause: %s", errStr, e.Cause)
	}
	return e.Msg + errStr
}

func (e Error) SetCause(format string, a ...interface{}) Error {
	e.Cause = fmt.Sprintf(format, a...)
	return e
}

func Map() (errors map[string]Error) {

	errors = make(map[string]Error)
	for k, v := range errorsMap {
		errors[k] = v
	}
	return
}
