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
