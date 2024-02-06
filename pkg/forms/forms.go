package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

func (f *Form) MaxLength(field string, maxLen int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > maxLen {
		f.Errors.Add(field, fmt.Sprintf("this field is too long (maximum is %d)", maxLen))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}

	}

	f.Errors.Add(field, "this field is invalid")
}

func (f *Form) Valid() bool {
	if len(f.Errors) == 0 {
		return true
	}
	return false
}
