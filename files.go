package login

import "embed"

// Files contains static files required by the application
//go:embed web/static/css/default.min.css
//go:embed web/static/css/error.min.css
//go:embed web/static/css/login.min.css
//go:embed web/template/*
var Files embed.FS
