package forms

type errors map[string][]string

// Adds an error message for the given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Gets an error message for the given form field
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}
	return string(es[0])
}
