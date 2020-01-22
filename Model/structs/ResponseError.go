package Model

type ResponseError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
