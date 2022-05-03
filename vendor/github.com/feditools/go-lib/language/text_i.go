package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextInvalidURI returns a translated phrase
func (l *Localizer) TextInvalidURI(count int) *LocalizedString {
	lg := logger.WithField("func", "TextInvalidURI")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "InvalidURI",
			Description: "the common phrase for invalid uri",
			One:         "Invalid URI",
			Other:       "Invalid URIs",
		},
		PluralCount: count,
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
