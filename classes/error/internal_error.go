package error

//資源驗證錯誤
const (
	InvalidInput int32 = 1000
)

//資源不存在
const (
	ResourceNotFound int32 = 2000
)

//伺服器端錯誤
const (
	DatabaseError           int32 = 3000
)

// getInternalErrorCodeMappingMessage: 取得內部錯誤代碼映照訊息
func getInternalErrorCodeMappingMessage(code int32) string {
	baseErrorCodeList := getBaseErrorCodeMappingList()
	messages := make(map[int32]string)
	for key, value := range baseErrorCodeList {
		messages[key] = value
	}

	//資源驗證錯誤
	messages[InvalidInput] = "無效的參數。"

	//資源不存在
	messages[ResourceNotFound] = "不存在的資料。"

	//伺服器端錯誤
	messages[DatabaseError] = "資料庫錯誤。"

	if code != NoError && messages[code] == "" {
		return "未知的錯誤。"
	}
	return messages[code]
}

// InternalError struct
type InternalError struct {
	error Error
}

// NewInternalError: 實體化 InternalError
func NewInternalError(code int32) InternalError {
	var internalError InternalError
	internalError.setCode(code)

	message := getInternalErrorCodeMappingMessage(code)
	internalError.SetMessage(message)

	return internalError
}

// setCode: 設定錯誤代碼
func (internalError *InternalError) setCode(code int32) {
	internalError.error.Code = code
}

// GetCode: 取得錯誤代碼
func (internalError *InternalError) GetCode() int32 {
	return internalError.error.Code
}

// SetMessage: 設定錯誤訊息
func (internalError *InternalError) SetMessage(message string) {
	internalError.error.Message = message
}

// GetMessage: 取得錯誤訊息
func (internalError *InternalError) GetMessage() string {
	return internalError.error.Message
}
