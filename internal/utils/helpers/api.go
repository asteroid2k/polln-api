package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/asteroid2k/polln-api/internal/utils/constants"
	"github.com/rs/zerolog/log"
)

type AppResponse struct {
	Status  int
	Data    any
	Headers http.Header
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SendJSON(w http.ResponseWriter, payload AppResponse) error {
	if payload.Status < 100 {
		payload.Status = 202
	}
	if payload.Data == nil {
		payload.Data = map[string]any{"message": "OK"}
	}

	json, err := json.Marshal(payload.Data)
	if err != nil {
		return err
	}
	json = append(json, '\n')
	for key, value := range payload.Headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	w.Write(json)
	return nil
}

func ParseJSON(w http.ResponseWriter, data io.ReadCloser, dst any) bool {
	err := json.NewDecoder(data).Decode(dst)
	if err == nil {
		return true
	}
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError

	switch {
	case errors.As(err, &syntaxError), errors.Is(err, io.ErrUnexpectedEOF):
		errmsg := "body contains malformed JSON %v"
		if syntaxError.Offset > 0 {
			errmsg = fmt.Sprintf(errmsg, fmt.Sprintf(" (at character %d)", syntaxError.Offset))
		}
		SendJSON(w, NewBadJSONResponse(errmsg))

	case errors.As(err, &unmarshalTypeError):
		if unmarshalTypeError.Field != "" {
			valErrors := []ValidationError{{Field: unmarshalTypeError.Field, Message: "Invalid type"}}
			SendJSON(w, NewValidationErrorResponse(valErrors, nil))
			break
		}
		SendJSON(w, NewBadJSONResponse(fmt.Sprintf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)))

	case errors.Is(err, io.EOF):
		SendJSON(w, NewBadJSONResponse("body must not be empty"))

	case errors.As(err, &invalidUnmarshalError):
		log.Fatal().Msg("invalid unmarshal destination")
		fallthrough
	default:
		SendJSON(w, AppResponse{Status: 500, Data: map[string]string{"message": constants.GenericError}})
	}
	return false
}

func NewValidationErrorResponse(errors []ValidationError, meta any) AppResponse {
	data := map[string]any{
		"message": constants.ValidationError,
		"code":    "VALIDATION_ERROR",
		"errors":  errors,
		"meta":    meta,
	}
	return AppResponse{Status: 422, Data: data}
}

func NewBadJSONResponse(meta any) AppResponse {
	data := map[string]any{
		"message": constants.BadJSONError,
		"code":    "INVALID_JSON",
		"meta":    meta,
	}
	return AppResponse{Status: 422, Data: data}
}
