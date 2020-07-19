package error

// BaseError interface
type BaseError interface {
	GetCode() int32
	SetMessage(string)
	GetMessage() string
}

// Error struct
type Error struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
}

// 基礎錯誤
const (
	NoError int32 = 0
)

// getBaseErrorCodeMappingList: 取得基礎錯誤代碼映照訊息
func getBaseErrorCodeMappingList() map[int32]string {
	messages := make(map[int32]string)
	messages[NoError] = ""

	return messages
}
