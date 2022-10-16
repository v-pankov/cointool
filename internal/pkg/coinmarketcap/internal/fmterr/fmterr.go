package fmterr

import (
	"fmt"
	"strings"

	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/payload"
)

func FormatError(statusCode int, statusPayload *payload.Status) error {
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
		"error response: %s",
		strings.Join(msgParts, ", "),
	)
}
