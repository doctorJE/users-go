package databaseresource

import (
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       int64  `orm:"column(id);pk"`
	Account  string `orm:"column(account);unique;"`
	Password string `orm:"column(password);"`
}

// UserObject struct
type UserObject struct {
	user User
}

// NewUserObject 實體化 UserObject
func NewUserObject() UserObject {
	var userObject UserObject
	return userObject
}

// Load: 載入使用者
func (userObject *UserObject) Load(userInstance *User) {
	userObject.user = *userInstance
}

// SetID: 設定編號
func (userObject *UserObject) SetID(id int64) {
	userObject.user.ID = id
}

// GetID: 取得編號
func (userObject *UserObject) GetID() int64 {
	return userObject.user.ID
}

// SetAccount: 設定使用者名稱
func (userObject *UserObject) SetAccount(account string) {
	userObject.user.Account = account
}

// GetAccount: 取得使用者名稱
func (userObject *UserObject) GetAccount() string {
	return userObject.user.Account
}

// SetHashedPassword: 設定 bcrypt 雜湊的密碼
func (userObject *UserObject) SetHashedPassword(hashedPassword string) {
	userObject.user.Password = hashedPassword
}

// IsPasswordCorrect: 驗證密碼是否正確
func (userObject *UserObject) IsPasswordCorrect(password string) bool {
	comparedError := bcrypt.CompareHashAndPassword(
		[]byte(userObject.user.Password),
		[]byte(password),
	)
	if comparedError != nil {
		return false
	}
	return true
}
