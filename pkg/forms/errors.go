package forms

type errors map[string][]string

func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	errorString := e[field]

	if len(errorString) == 0 {
		return ""
	}

	return errorString[0]
}
