package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextReturnURI returns a translated phrase
func (l *Localizer) TextReturnURI() *LocalizedString {
	lg := logger.WithField("func", "TextReturnURI")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "ReturnURI",
			Description: "the common phrase for return uri",
			Other:       "Return URI",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
