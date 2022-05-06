package models

// Link represents a link in an api response.
type Link struct {
	Rel  string `json:"rel"`
	HRef string `json:"href"`
}
