package errrorhandleablereturns

import "github.com/doctorJE/users-go/classes/error"

// ErrorHandleable struct
type ErrorHandleable struct {
	internalError error.InternalError
}

// setError: 設定錯誤
func (errorHandleable *ErrorHandleable) setError(internalError error.InternalError) {
	errorHandleable.internalError = internalError
}

// getError: 取得錯誤
func (errorHandleable *ErrorHandleable) getError() error.InternalError {
	return errorHandleable.internalError
}

// hasError: 是否有錯誤
func (errorHandleable *ErrorHandleable) hasError() bool {
	if errorHandleable.internalError.GetCode() == error.NoError {
		return false
	}
	return true
}
