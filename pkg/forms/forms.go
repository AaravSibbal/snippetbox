package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

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

func (f *Form) MinLength(field string, minLen int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) < minLen {
		f.Errors.Add(field, fmt.Sprintf("The field is too short minimum length is %d", minLen))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
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

	return len(f.Errors) == 0
}
