package main

import (
	"regexp"

	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

func regexValidation(pattern string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(pattern)
		return re.MatchString(fl.Field().String())
	}
}

func dateValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(fl.Field().String())
}

func RegisterValidations(v *validator.Validate) {
	v.RegisterValidation("retailerRegex", regexValidation(`^[\w\s\-&]+$`))
	v.RegisterValidation("shortDescRegex", regexValidation(`^[\w\s\-]+$`))
	v.RegisterValidation("priceRegex", regexValidation(`^\d+\.\d{2}$`))
	v.RegisterValidation("time24", regexValidation(`^(2[0-3]|[01]?[0-9]):([0-5]?[0-9])$`))
	v.RegisterValidation("date", dateValidation)

}

func generateUniqueID() string {
	var uniqueID string
	for {
		uniqueID = uuid.New().String()

		_, exists := pointsMap[uniqueID]

		if !exists {
			break
		}
	}
	return uniqueID
}
