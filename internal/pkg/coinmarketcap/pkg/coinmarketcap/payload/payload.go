package payload

type Status struct {
	ErrorCode    *int    `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}

type StatusPayload struct {
	Status Status `json:"status"`
}
