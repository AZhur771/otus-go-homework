package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var ErrFieldNotValid = errors.New("validation error")

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errorMessage strings.Builder

	for _, validationError := range v {
		errorMessage.WriteString(fmt.Sprintf("%s: %v\n", validationError.Field, validationError.Err))
	}

	return errorMessage.String()
}

// StringValidator (len/regexp/in).
type StringValidator struct{}

func (v StringValidator) ValidateLen(value string, length int) (bool, error) {
	if len(value) != length {
		return false, fmt.Errorf("should be equal to %v chars long: %w", length, ErrFieldNotValid)
	}

	return true, nil
}

func (v StringValidator) ValidateRegexp(value string, re *regexp.Regexp) (bool, error) {
	if !re.MatchString(value) {
		return false, fmt.Errorf("should match %v regexp: %w", re, ErrFieldNotValid)
	}

	return true, nil
}

func (v StringValidator) ValidateIn(value string, in []string) (bool, error) {
	for _, v := range in {
		if v == value {
			return true, nil
		}
	}

	return false, fmt.Errorf("value should be in %v: %w", in, ErrFieldNotValid)
}

// NumberValidator (min/max/in).
type NumberValidator struct{}

func (v NumberValidator) ValidateMin(value int, min int) (bool, error) {
	if value < min {
		return false, fmt.Errorf("value should not be less than %v: %w", min, ErrFieldNotValid)
	}

	return true, nil
}

func (v NumberValidator) ValidateMax(value int, max int) (bool, error) {
	if value > max {
		return false, fmt.Errorf("value should be less or equal to %v: %w", max, ErrFieldNotValid)
	}

	return true, nil
}

func (v NumberValidator) ValidateIn(value int, in []int) (bool, error) {
	for _, v := range in {
		if v == value {
			return true, nil
		}
	}

	return false, fmt.Errorf("value should be in %v: %w", in, ErrFieldNotValid)
}

func validateString(tag string, fieldName string, fieldValue reflect.Value) (bool, error) {
	validationErrors := make(ValidationErrors, 0)

	value := fieldValue.String()
	validator := StringValidator{}

	for _, t := range strings.Split(tag, "|") {
		ts := strings.Split(t, ":")

		isValid := true
		var validationError error

		switch ts[0] {
		case "len":
			num, err := strconv.Atoi(ts[1])
			if err != nil {
				return false, fmt.Errorf("failed to convert string to int %w", err)
			}

			isValid, validationError = validator.ValidateLen(value, num)
		case "regexp":
			re, err := regexp.Compile(ts[1])
			if err != nil {
				return false, fmt.Errorf("failed to compile regexp %w", err)
			}

			isValid, validationError = validator.ValidateRegexp(value, re)
		case "in":
			in := strings.Split(ts[1], ",")

			isValid, validationError = validator.ValidateIn(value, in)
		}

		if !isValid {
			validationErrors = append(validationErrors, ValidationError{
				Field: fieldName,
				Err:   validationError,
			})
		}
	}

	if len(validationErrors) > 0 {
		return false, validationErrors
	}

	return true, nil
}

func validateInt(tag string, fieldName string, fieldValue reflect.Value) (bool, error) {
	validationErrors := make(ValidationErrors, 0)

	value := int(fieldValue.Int())
	validator := NumberValidator{}

	for _, t := range strings.Split(tag, "|") {
		ts := strings.Split(t, ":")

		isValid := true
		var validationError error

		switch ts[0] {
		case "min":
			num, err := strconv.Atoi(ts[1])
			if err != nil {
				return false, fmt.Errorf("failed to convert string to int %w", err)
			}

			isValid, validationError = validator.ValidateMin(value, num)
		case "max":
			num, err := strconv.Atoi(ts[1])
			if err != nil {
				return false, fmt.Errorf("failed to convert string to int %w", err)
			}

			isValid, validationError = validator.ValidateMax(value, num)
		case "in":
			inOfStrings := strings.Split(ts[1], ",")

			in := make([]int, len(inOfStrings))

			for i, v := range inOfStrings {
				num, err := strconv.Atoi(v)
				if err != nil {
					return false, fmt.Errorf("failed to convert string to int %w", err)
				}

				in[i] = num
			}

			isValid, validationError = validator.ValidateIn(value, in)
		}

		if !isValid {
			validationErrors = append(validationErrors, ValidationError{
				Field: fieldName,
				Err:   validationError,
			})
		}
	}

	if len(validationErrors) > 0 {
		return false, validationErrors
	}

	return true, nil
}

func validateSlice(tag string, fieldName string, fieldValue reflect.Value) (bool, error) {
	validationErrors := make(ValidationErrors, 0)

	for i := 0; i < fieldValue.Len(); i++ {
		value := fieldValue.Index(i)
		isValid, err := validateTag(tag, fieldName, value)
		if !isValid {
			var ve ValidationErrors
			if errors.As(err, &ve) {
				validationErrors = append(validationErrors, ve...)
			} else {
				return false, fmt.Errorf("slice loop error: %w", err)
			}
		}
	}

	if len(validationErrors) > 0 {
		return false, validationErrors
	}

	return true, nil
}

func validateTag(tag string, fieldName string, fieldValue reflect.Value) (bool, error) {
	var isValid bool
	var err error

	//nolint:exhaustive
	switch fieldValue.Kind() {
	case reflect.String:
		isValid, err = validateString(tag, fieldName, fieldValue)
	case reflect.Int:
		isValid, err = validateInt(tag, fieldName, fieldValue)
	case reflect.Slice:
		isValid, err = validateSlice(tag, fieldName, fieldValue)
	}

	return isValid, err
}

func Validate(v interface{}) error {
	// validation errors
	validationErrors := make(ValidationErrors, 0)

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("not a struct, received %T", v)
	}

	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := rv.Type().Field(i)
		tag := fieldType.Tag.Get("validate")

		// skip empty/ignored tags
		if tag == "" || tag == "-" {
			continue
		}

		rvField := rv.Field(i)

		// field is exported and value can be accessed
		if rvField.CanInterface() {
			isValidTag, err := validateTag(tag, fieldType.Name, fieldValue)

			if !isValidTag {
				var ve ValidationErrors
				if errors.As(err, &ve) {
					// validation errors
					validationErrors = append(validationErrors, ve...)
				} else {
					// program errors
					return fmt.Errorf("program error: %w", err)
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
