package request

import (
	"backend/pkg/logger"
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		logger.ERROR(err)
		return payload, err
	}
	return payload, nil
}
