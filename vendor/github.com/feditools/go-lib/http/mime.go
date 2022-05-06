package http

const (
	// MimeAll matches any mime type.
	MimeAll = `*/*`
	// MimeAppJRDJSON represents a JSON Resource Descriptor type.
	MimeAppJRDJSON = `application/jrd+json`
	// MimeAppJSON represents a JavaScript Object Notation type.
	MimeAppJSON = `application/json`
	// MimeAppActivityJSON represents a JSON activity pub action type.
	MimeAppActivityJSON = `application/activity+json`
	// MimeAppActivityLDJSON represents JSON-based Linked Data for activity streams type.
	MimeAppActivityLDJSON = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
	// MimeTextHTML represents a html type.
	MimeTextHTML = `text/html`
)
