package utils

import "reflect"

type EmptyObject map[string]string

type ResponseSuccess struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ResponseError struct {
	Status   string `json:"status"`
	Code     string `json:"code"`
	DataFail any    `json:"data_fail"`
	Message  any    `json:"message"`
}

type ResponseFail struct {
	Status   string `json:"status" default:"fail"`
	Data     any    `json:"data"`
	DataFail any    `json:"data_fail"`
	Message  string `json:"message"`
}

// ApiResponseSuccess is a function to return response success
func ApiResponseSuccess(status string, data any) *ResponseSuccess {
	// return empty string if data is nil
	if isNilFixed(data) {
		data = EmptyObject{}
	}
	// return array if data is empty
	if isArrayEmpty(data) {
		data = []EmptyObject{}
	}
	return &ResponseSuccess{
		Status: status,
		Data:   data,
	}
}

// ApiResponseError is a function to return response error
func ApiResponseError[T any](message T, code string) *ResponseError {
	return &ResponseError{
		Status:   "error",
		Code:     code,
		DataFail: message,
		Message:  message,
	}
}

// ApiResponseFail is a function to return response fail
func ApiResponseFail(data any, message string) *ResponseFail {
	if isNilFixed(data) {
		data = EmptyObject{}
	}

	if isArrayEmpty(data) {
		data = []EmptyObject{}
	}

	return &ResponseFail{
		Status:   "fail",
		Data:     data,
		DataFail: data,
		Message:  message,
	}
}

// IsNilFixed is a fixed version of IsNil from reflection package
func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Interface:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}

// IsArrayEmpty is a function to check if an array is empty
func isArrayEmpty(i interface{}) bool {
	if i == nil {
		return true
	}

	if reflect.TypeOf(i).Kind() == reflect.Slice {
		return reflect.ValueOf(i).Len() == 0
	}

	return false
}
