package validatePayload

import (
	"github.com/go-playground/validator/v10"
)

func Validate(payload any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(payload)
	//var validationErrors validator.ValidationErrors
	//errors.As(err, &validationErrors)
	if err != nil {
		return err
	}
	return nil
}
