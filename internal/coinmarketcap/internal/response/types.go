package response

type StatusPayload struct {
	ErrorCode    *int    `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}
