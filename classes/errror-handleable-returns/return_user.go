package errrorhandleablereturns

import (
	databaseresource "github.com/doctorJE/users-go/classes/database-resources"
	"github.com/doctorJE/users-go/classes/error"
)

// ErrorHandleableReturnUser struct
type ErrorHandleableReturnUser struct {
	user            databaseresource.UserObject
	errorHandleable ErrorHandleable
}

// NewReturnUser: 實例化 ReturnUser
func NewReturnUser(user databaseresource.UserObject, internalError *error.InternalError) ErrorHandleableReturnUser {
	var returnUser ErrorHandleableReturnUser
	returnUser.SetUser(user)

	var err error.InternalError
	if internalError == nil {
		err = error.NewInternalError(error.NoError)
	} else {
		err = *internalError
	}
	returnUser.SetError(err)
	return returnUser
}

// SetUser: 設定使用者
func (returnUser *ErrorHandleableReturnUser) SetUser(user databaseresource.UserObject) {
	returnUser.user = user
}

// GetUser: 取得使用者
func (returnUser *ErrorHandleableReturnUser) GetUser() databaseresource.UserObject {
	return returnUser.user
}

// SetError: 設定錯誤
func (returnUser *ErrorHandleableReturnUser) SetError(internalError error.InternalError) {
	returnUser.errorHandleable.setError(internalError)
}

// GetError: 取得錯誤
func (returnUser *ErrorHandleableReturnUser) GetError() error.InternalError {
	return returnUser.errorHandleable.getError()
}

// HasError: 是否有錯誤
func (returnUser *ErrorHandleableReturnUser) HasError() bool {
	return returnUser.errorHandleable.hasError()
}
