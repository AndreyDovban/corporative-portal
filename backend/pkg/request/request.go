package request

import (
	"backend/pkg/logger"
	"backend/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (T, error) {
	var payload T

	payload, err := Decode[T](r.Body)
	if err != nil {
		response.Json(*w, err.Error(), 402)
		logger.ERROR(err.Error(), 402)

		return payload, err
	}

	err = IsValid(payload)
	if err != nil {
		response.Json(*w, err.Error(), 402)
		logger.ERROR(err.Error(), 402)
		return payload, err
	}

	return payload, nil
}
