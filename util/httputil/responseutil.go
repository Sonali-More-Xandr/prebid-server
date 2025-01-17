package httputil

import (
	"fmt"
	"net/http"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/errortypes"
)

func CheckResponseStatusCodeForErrors(response *adapters.ResponseData) error {
	if response.StatusCode == http.StatusBadRequest {
		return &errortypes.BadInput{
			Message: fmt.Sprintf("Unexpected status code: %d. Run with request.debug = 1 for more info", response.StatusCode),
		}
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d. Run with request.debug = 1 for more info", response.StatusCode)
	}

	return nil
}

func IsResponseStatusCodeNoContent(response *adapters.ResponseData) bool {
	return response.StatusCode == http.StatusNoContent
}
