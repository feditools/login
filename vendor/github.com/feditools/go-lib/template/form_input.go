package template

import "github.com/feditools/go-lib/language"

// FormInput is a templated form input.
type FormInput struct {
	ID           string
	Type         string
	Name         string
	Placeholder  string
	Label        *language.LocalizedString
	LabelClass   string
	Value        string
	WrappedClass string
	Disabled     bool
	Required     bool
	Validation   *FormValidation
}

// FormValidation is a validation response to a form input.
type FormValidation struct {
	Valid    bool
	Response string
}
