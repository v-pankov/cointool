package requests

import (
	"coinconv/internal/coinmarketcap/internal/responses"
	"fmt"
	"strings"
)

func fmtBadStatusError(statusCode int, statusPayload *responses.StatusPayload) error {
	msgParts := make([]string, 0, 3)

	msgParts = append(
		msgParts,
		fmt.Sprintf(
			"http_status=%d",
			statusCode,
		),
	)

	if statusPayload.ErrorCode != nil {
		msgParts = append(
			msgParts,
			fmt.Sprintf(
				"error_code=%d",
				*statusPayload.ErrorCode,
			),
		)
	}

	if statusPayload.ErrorMessage != nil {
		msgParts = append(
			msgParts,
			"error_message="+*statusPayload.ErrorMessage,
		)
	}

	return fmt.Errorf(
		"unexpected status: %s",
		strings.Join(msgParts, ","),
	)
}
