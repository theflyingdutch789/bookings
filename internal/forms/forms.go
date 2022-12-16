package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New creates a custom form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if the form field is empty or not
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "Has nothing in the passed field")
		return false
	}
	return true

}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This filed is important")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field has to be minimum %d length", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}

}
