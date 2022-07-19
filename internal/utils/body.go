package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
)

func ValidateBody(body io.Reader, validate *validator.Validate, params interface{}) ValidationError {
	if err := json.NewDecoder(body).Decode(params); err != nil {
		return ValidationError{
			"body": []string{"invalid request body"},
		}
	}
	var errs []string
	if err := validate.Struct(params); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("%s should be %s", err.Field(), err.Tag()))
		}
	}
	if len(errs) > 0 {
		return ValidationError{"body": errs}
	}
	return nil
}

func ValidateStruct(validate *validator.Validate, params interface{}) ValidationError {
	var errs []string
	if err := validate.Struct(params); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("%s should be %s", err.Field(), err.Tag()))
		}
	}
	if len(errs) > 0 {
		return ValidationError{"body": errs}
	}
	return nil
}

func SendResponse(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Unable to encode to json", err)
		return
	}
}

type ValidationError map[string][]string

func SendError(w http.ResponseWriter, code int, error string) {
	SendResponse(w, code, map[string]interface{}{
		"errors": []string{error},
	})
}

func SendErrors(w http.ResponseWriter, code int, errors map[string][]string) {
	SendResponse(w, code, map[string]interface{}{
		"errors": errors,
	})
}
