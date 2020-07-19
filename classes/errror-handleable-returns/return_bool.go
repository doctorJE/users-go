package errrorhandleablereturns

import "github.com/doctorJE/users-go/classes/error"

// ErrorHandleableReturnBool struct
type ErrorHandleableReturnBool struct {
	isTrue          bool
	errorHandleable ErrorHandleable
}

// NewReturnBool: 實例化 ReturnBool
func NewReturnBool(isTrue bool, internalError *error.InternalError) ErrorHandleableReturnBool {
	var errorHandleableReturnBool ErrorHandleableReturnBool
	errorHandleableReturnBool.SetIsTrue(isTrue)

	var err error.InternalError
	if internalError == nil {
		err = error.NewInternalError(error.NoError)
	} else {
		err = *internalError
	}

	errorHandleableReturnBool.SetError(err)
	return errorHandleableReturnBool
}

// SetIsTrue: 設定是否為 true
func (returnBool *ErrorHandleableReturnBool) SetIsTrue(isTrue bool) {
	returnBool.isTrue = isTrue
}

// IsTrue: 是否為 true
func (returnBool *ErrorHandleableReturnBool) IsTrue() bool {
	return returnBool.isTrue
}

// SetError: 設定錯誤
func (returnBool *ErrorHandleableReturnBool) SetError(internalError error.InternalError) {
	returnBool.errorHandleable.setError(internalError)
}

// GetError: 取得錯誤
func (returnBool *ErrorHandleableReturnBool) GetError() error.InternalError {
	return returnBool.errorHandleable.getError()
}

// HasError: 是否有錯誤
func (returnBool *ErrorHandleableReturnBool) HasError() bool {
	return returnBool.errorHandleable.hasError()
}
