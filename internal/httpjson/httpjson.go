package httpjson

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const maxBodySize = 1 << 20

func Decode(r *http.Request, dst any) error {
	if r.Body == nil {
		return errors.New("request body is required")
	}

	decoder := json.NewDecoder(io.LimitReader(r.Body, maxBodySize))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func Write(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
