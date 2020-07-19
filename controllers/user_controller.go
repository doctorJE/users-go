package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/doctorJE/users-go/classes/converter"
	databaseResources "github.com/doctorJE/users-go/classes/database-resources"
	"github.com/doctorJE/users-go/classes/error"
	"github.com/doctorJE/users-go/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// response: 回應
func (this *UserController) response(isOk *bool, apiError *error.APIError, statusCode int) {
	this.Ctx.ResponseWriter.WriteHeader(statusCode)
	this.Data["json"] = converter.ConvertOutput(isOk, apiError)
	this.ServeJSON()
}

// Create: 建立使用者
func (this *UserController) Create() {
	var user databaseResources.User
	data := this.Ctx.Input.RequestBody
	unmarshalError := json.Unmarshal(data, &user)
	if unmarshalError != nil {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	if len(user.Account) == 0 ||
		len(user.Account) > 50 ||
		len(user.Password) == 0 ||
		len(user.Password) > 50 {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	userObject := databaseResources.NewUserObject()
	userObject.Load(&user)

	isUserExistedReturns := models.IsAccountExisted(user.Account)
	if isUserExistedReturns.HasError() {
		isOk := false
		apiError := error.NewAPIError(error.APIInternalServerError)
		this.response(&isOk, &apiError, http.StatusInternalServerError)
		return
	}
	if isUserExistedReturns.IsTrue() {
		isOk := false
		apiError := error.NewAPIError(error.APIAccountHasExisted)
		this.response(&isOk, &apiError, http.StatusConflict)
		return
	}

	hashedPasswordByte, hashedError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if hashedError != nil {
		isOk := false
		apiError := error.NewAPIError(error.APIInternalServerError)
		this.response(&isOk, &apiError, http.StatusInternalServerError)
		return
	}

	insertUserReturns := models.InsertUser(&userObject, string(hashedPasswordByte))
	if insertUserReturns.HasError() {
		isOk := false

		internalError := insertUserReturns.GetError()
		if internalError.GetCode() == error.ResourceNotFound {
			apiError := error.NewAPIError(error.APIUserNotFound)
			this.response(&isOk, &apiError, http.StatusNotFound)
			return
		}
		apiError := error.NewAPIError(error.APIInternalServerError)
		this.response(&isOk, &apiError, http.StatusInternalServerError)
		return
	}

	isOk := true
	this.response(&isOk, nil, http.StatusOK)
	return
}

// Delete: 刪除
func (this *UserController) Delete() {
	var user databaseResources.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	if len(user.Account) == 0 ||
		len(user.Account) > 50 {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	deleteUserReturns := models.DeleteByAccount(user.Account)
	if deleteUserReturns.HasError() {
		isOk := false
		internalError := deleteUserReturns.GetError()
		if internalError.GetCode() == error.ResourceNotFound {
			apiError := error.NewAPIError(error.APIUserNotFound)
			this.response(&isOk, &apiError, http.StatusNotFound)
			return
		}
		apiError := error.NewAPIError(error.APIInternalServerError)
		this.response(&isOk, &apiError, http.StatusInternalServerError)
		return
	}

	isOk := true
	this.response(&isOk, nil, http.StatusOK)
	return
}

// ChangePassword: 變更密碼
func (this *UserController) ChangePassword() {
	var user databaseResources.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	if len(user.Account) == 0 ||
		len(user.Account) > 50 ||
		len(user.Password) == 0 ||
		len(user.Password) > 50 {
		isOk := false
		apiError := error.NewAPIError(error.APIInvalidInput)
		this.response(&isOk, &apiError, http.StatusBadRequest)
		return
	}

	userObject := databaseResources.NewUserObject()
	userObject.Load(&user)

	getUserReturns := models.GetByAccount(userObject)
	if getUserReturns.HasError() {
		isOk := false
		internalError := getUserReturns.GetError()
		if internalError.GetCode() == error.ResourceNotFound {
			apiError := error.NewAPIError(error.APIUserNotFound)
			this.response(&isOk, &apiError, http.StatusNotFound)
			return
		}

		apiError := error.NewAPIError(error.APIInternalServerError)
		this.response(&isOk, &apiError, http.StatusInternalServerError)
		return
	}
	existedUserObject := getUserReturns.GetUser()

	if !existedUserObject.IsPasswordCorrect(user.Password) {
		hashedPasswordByte, encryptedError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if encryptedError != nil {
			isOk := false
			apiError := error.NewAPIError(error.APIEncryptionError)
			this.response(&isOk, &apiError, http.StatusInternalServerError)
			return
		}

		updatedPasswordReturns := models.UpdatePassword(existedUserObject, string(hashedPasswordByte))
		if updatedPasswordReturns.HasError() {
			isOk := false
			internalError := updatedPasswordReturns.GetError()
			if internalError.GetCode() == error.ResourceNotFound {
				apiError := error.NewAPIError(error.APIUserNotFound)
				this.response(&isOk, &apiError, http.StatusNotFound)
				return
			}
			apiError := error.NewAPIError(error.APIInternalServerError)
			apiError.SetMessage(internalError.GetMessage())
			this.response(&isOk, &apiError, http.StatusInternalServerError)
			return
		}
	}

	isOk := true
	this.response(&isOk, nil, http.StatusOK)
	return
}
