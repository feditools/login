package template

import "github.com/feditools/go-lib/language"

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

type FormValidation struct {
	Valid    bool
	Response string
}
