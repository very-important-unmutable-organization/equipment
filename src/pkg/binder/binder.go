package binder

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

var ErrInvalidData = errors.New("invalid data")

type InputBinder struct {
	validate *validator.Validate
}

func NewInputBinder() *InputBinder {
	return &InputBinder{
		validate: validator.New(),
	}
}

func (b *InputBinder) Bind(r *http.Request, input interface{}) error {
	if err := render.Decode(r, input); err != nil {
		return ErrInvalidData
	}

	if err := TrimStrFields(input); err != nil {
		panic(err)
	}

	if err := b.validate.Struct(input); err != nil {
		return formatError(err)
	}
	return nil
}

func formatError(err error) error {
	validationErrors := err.(validator.ValidationErrors)
	formattedErrors := make([]string, len(validationErrors))
	for i, fieldError := range validationErrors {
		formattedErrors[i] = fieldError.Error()
	}
	return errors.New(strings.Join(formattedErrors, ", "))
}
