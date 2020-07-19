package error

//資源驗證錯誤
const (
	APIInvalidInput                int32 = 1
	APIIncorrectUsernameOrPassword int32 = 2
)

//資源已存在
const (
	APIAccountHasExisted int32 = 11
)

//資源不存在
const (
	APIUserNotFound int32 = 21
)

//伺服器端錯誤
const (
	APIInternalServerError int32 = 31
)

//加解密失敗
const (
	APIEncryptionError int32 = 41
)

// getAPIErrorCodeMappingMessage: 取得 API 錯誤代碼映照訊息
func getAPIErrorCodeMappingMessage(code int32) string {
	baseErrorCodeList := getBaseErrorCodeMappingList()
	messages := make(map[int32]string)
	for key, value := range baseErrorCodeList {
		messages[key] = value
	}

	//資源驗證錯誤
	messages[APIInvalidInput] = "無效的參數。"
	messages[APIIncorrectUsernameOrPassword] = "Login Failed"

	//資源已存在
	messages[APIAccountHasExisted] = "使用者名稱已存在。"

	//資源不存在
	messages[APIUserNotFound] = "使用者不存在。"

	//伺服器端錯誤
	messages[APIInternalServerError] = "內部伺服器錯誤。"

	//加解密失敗
	messages[APIEncryptionError] = "加密發生錯誤。"

	if code != NoError && messages[code] == "" {
		return "未知的錯誤。"
	}
	return messages[code]
}

// APIError struct
type APIError struct {
	error Error
}

// NewAPIError: 實體化 APIError
func NewAPIError(code int32) APIError {
	var apiError APIError
	apiError.setCode(code)

	message := getAPIErrorCodeMappingMessage(code)
	apiError.SetMessage(message)

	return apiError
}

// setCode: 設定錯誤代碼
func (apiError *APIError) setCode(errorCode int32) {
	apiError.error.Code = errorCode
}

// GetCode: 取得錯誤代碼
func (apiError *APIError) GetCode() int32 {
	return apiError.error.Code
}

// SetMessage: 設定錯誤訊息
func (apiError *APIError) SetMessage(message string) {
	apiError.error.Message = message
}

// GetMessage: 取得錯誤訊息
func (apiError *APIError) GetMessage() string {
	return apiError.error.Message
}

// ToResponseStruct: 產出回應結構
func (apiError *APIError) ToResponseStruct() (errorResponse *Error) {
	errorResponse = &Error{
		Code:    apiError.GetCode(),
		Message: apiError.GetMessage(),
	}
	return
}
