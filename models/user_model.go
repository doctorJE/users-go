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

// DeleteByAccount: 透過 account 刪除使用者
func DeleteByAccount(account string) errrorhandleablereturns.ErrorHandleableReturnBool {
	effectedNum, queryError := orm.NewOrm().QueryTable("user").Filter("account", account).Delete()
	if queryError != nil {
		if queryError == orm.ErrNoRows {
			internalError := error.NewInternalError(error.ResourceNotFound)
			return errrorhandleablereturns.NewReturnBool(false, &internalError)
		}
		internalError := error.NewInternalError(error.DatabaseError)
		internalError.SetMessage(queryError.Error())
		return errrorhandleablereturns.NewReturnBool(false, &internalError)
	}

	isDeleted := effectedNum > 0
	return errrorhandleablereturns.NewReturnBool(isDeleted, nil)
}

// IsAccountExisted: 帳號是否存在
func IsAccountExisted(account string) errrorhandleablereturns.ErrorHandleableReturnBool {
	countUser, queryError := orm.NewOrm().QueryTable("user").Filter("account", account).Count()
	if queryError != nil {
		if queryError == orm.ErrNoRows {
			return errrorhandleablereturns.NewReturnBool(false, nil)
		}
		internalError := error.NewInternalError(error.DatabaseError)
		internalError.SetMessage(queryError.Error())
		return errrorhandleablereturns.NewReturnBool(false, &internalError)
	}

	isExisted := countUser > 0
	return errrorhandleablereturns.NewReturnBool(isExisted, nil)
}

// GetByAccount: 透過 account 取得使用者
func GetByAccount(userObject databaseResources.UserObject) errrorhandleablereturns.ErrorHandleableReturnUser {
	queryUser := databaseResources.User{
		Account: userObject.GetAccount(),
	}

	queryError := orm.NewOrm().Read(&queryUser, "Account")
	if queryError != nil {
		if queryError == orm.ErrNoRows {
			internalError := error.NewInternalError(error.ResourceNotFound)
			return errrorhandleablereturns.NewReturnUser(userObject, &internalError)
		}
		internalError := error.NewInternalError(error.DatabaseError)
		internalError.SetMessage(queryError.Error())
		return errrorhandleablereturns.NewReturnUser(userObject, &internalError)
	}

	userObject.Load(&queryUser)
	return errrorhandleablereturns.NewReturnUser(userObject, nil)
}
