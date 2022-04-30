package fedi

import (
	"github.com/feditools/login/internal/log"
)

type empty struct{}

var logger = log.WithPackageField(empty{})
