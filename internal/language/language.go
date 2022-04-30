package language

import (
	"github.com/BurntSushi/toml"
	"github.com/feditools/login"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"strings"
)

// DefaultLanguage is the default language of the application
var DefaultLanguage = language.English

// Module represent the language module for translating text
type Module struct {
	lang       language.Tag
	langBundle *i18n.Bundle
}

// New creates a new language module
func New() (*Module, error) {
	l := logger.WithField("func", "New")

	module := Module{
		lang:       DefaultLanguage,
		langBundle: i18n.NewBundle(DefaultLanguage),
	}

	module.langBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	dir, err := login.Files.ReadDir("locales")
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".toml") {
			continue
		}
		l.Debugf("loading language file: %s", d.Name())

		// open it
		file, err := login.Files.Open("locales/" + d.Name())
		if err != nil {
			return nil, err
		}

		// read it
		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		module.langBundle.MustParseMessageFileBytes(buffer, d.Name())
	}

	return &module, nil
}

// Language returns the default language
func (m Module) Language() language.Tag { return m.lang }
