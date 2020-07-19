package converter

import error "github.com/doctorJE/users-go/classes/error"

// result struct
type result struct {
	IsOK *bool `json:"IsOK"`
}

// Output struct
type Output struct {
	*error.Error
	Result  *result `json:"Result"`
}

// ConvertOutput: 轉換為輸出結果
func ConvertOutput(isOK *bool, apiError *error.APIError) Output {
	var resultResponse *result
	if isOK != nil {
		resultResponse = &result{
			IsOK: isOK,
		}
	}

	if apiError == nil {
		apiNoError := error.NewAPIError(error.NoError)
		return Output{
			Error: apiNoError.ToResponseStruct(),
			Result: resultResponse,
		}
	}

	return Output{
		Error: apiError.ToResponseStruct(),
		Result: resultResponse,
	}
}
