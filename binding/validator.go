package binding

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

// StructValidator is the minimal interface which needs to be implemented in order for it
// to be used as the validator engine for ensuring the correctness of the request.
type StructValidator interface {
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is a slice|array, the validation should be performed travel on every element.
	// If the received type is not a struct or slice|array, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
	ValidateStruct(any) error

	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
	Engine() any
}

var Validator StructValidator = &defaultValidator{}

func validate(obj any) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *defaultValidator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Pointer:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		validateErr := make(SliceValidationError, 0)
		for i := 0; i < value.Len(); i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateErr = append(validateErr, err)
			}
		}
		if len(validateErr) == 0 {
			return nil
		}
		return validateErr
	default:
		return nil
	}
}

func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyInit()
	return v.validate.Struct(obj)
}

func (v *defaultValidator) Engine() any {
	v.lazyInit()
	return v.validate
}

func (v *defaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}

type SliceValidationError []error

// Error concatenates all error elements in SliceValidationError into a single string separated by \n.
func (err SliceValidationError) Error() string {
	n := len(err)
	if n == 0 {
		return ""
	}

	var builder strings.Builder
	if err[0] != nil {
		fmt.Fprintf(&builder, "[%d]: %s", 0, err[0].Error())
	}
	if n > 1 {
		for i := 1; i < n; i++ {
			if err[i] != nil {
				builder.WriteString("\n")
				fmt.Fprintf(&builder, "[%d]: %s", i, err[i].Error())
			}
		}
	}
	return builder.String()
}
