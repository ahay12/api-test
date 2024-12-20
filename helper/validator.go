package helper

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ConstructErrors(err error) []ErrorField {
	errorList, success := err.(validator.ValidationErrors)
	errorFields := []ErrorField{}
	if success {
		for _, value := range errorList {
			fieldIDString := SetFirstLetterToLowerCase(value.StructField())
			fieldValue := value.Value().(string)
			fieldErrorCaused := value.Translate(nil)
			fieldErrorMessage := "Invalid Value"
			switch value.Tag() {
			case "required":
				fieldErrorMessage = "Required"
			case "alpha":
				fieldErrorMessage = "Only letters allowed"
				// case "product_name":
				// 	fieldErrorMessage = "Product name must be unique"
			}
			errorFields = append(errorFields, ErrorField{
				fieldIDString, fieldValue, fieldErrorCaused, fieldErrorMessage})
		}
	}
	return errorFields
}

func SetFirstLetterToLowerCase(toBeConvertedString string) string {
	fieldNameBytes := []rune(toBeConvertedString)
	fieldNameBytes[0] = unicode.ToLower(fieldNameBytes[0])
	return string(fieldNameBytes)
}
