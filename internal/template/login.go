package template

// Login contains the variables for the "login" template.
type Login struct {
	Common

	Image string

	FormError   string
	FormAccount string
}
