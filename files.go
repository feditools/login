package login

import "embed"

// Files contains static files required by the application
//go:embed locales/active.*.toml
//go:embed web/static/*
//go:embed web/template/*
var Files embed.FS
