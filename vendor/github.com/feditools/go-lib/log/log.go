package log

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

// WithPackageField creates a new logrus entry with the package name added as a field.
func WithPackageField(m interface{}) *logrus.Entry {
	return logrus.WithField("module", "go-lib").WithField("package", strings.ReplaceAll(strings.TrimPrefix(reflect.TypeOf(m).PkgPath(), "github.com/feditools/go-lib/"), "/", "."))
}
