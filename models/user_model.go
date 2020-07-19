package models

import (
	"github.com/astaxie/beego/orm"
	databaseResources "github.com/doctorJE/users-go/classes/database-resources"
	"github.com/doctorJE/users-go/classes/error"
	errrorhandleablereturns "github.com/doctorJE/users-go/classes/errror-handleable-returns"
)

// InsertUser: 新增使用者
func InsertUser(userObject *databaseResources.UserObject, hashedPassword string) errrorhandleablereturns.ErrorHandleableReturnBool {
	queryUser := databaseResources.User{
		Account:  userObject.GetAccount(),
		Password: hashedPassword,
	}

	userId, queryError := orm.NewOrm().Insert(&queryUser)
	if queryError != nil {
		internalError := error.NewInternalError(error.DatabaseError)
		internalError.SetMessage(queryError.Error())
		return errrorhandleablereturns.NewReturnBool(false, &internalError)
	}

	userObject.SetID(userId)
	return errrorhandleablereturns.NewReturnBool(true, nil)
}
