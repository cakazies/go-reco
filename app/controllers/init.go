package controllers

type (
	// Response struct for response json
	Response struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Status  string      `json:"status"`
	}
)

func response(data interface{}, message string, status string) Response {
	result := Response{}
	result.Message = message
	result.Data = data
	result.Status = status
	return result
}
