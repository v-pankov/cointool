package request

import (
	"fmt"
	"strings"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/response"
)

func fmtBadStatusError(statusCode int, statusPayload *response.StatusPayload) error {
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
			fmt.Sprintf(`error_message="%s"`, *statusPayload.ErrorMessage),
		)
	}

	return fmt.Errorf(
		"unexpected status: %s",
		strings.Join(msgParts, ", "),
	)
}
