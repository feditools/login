package login

import "embed"

// Files contains static files required by the application
//go:embed locales/active.*.toml
var Files embed.FS
