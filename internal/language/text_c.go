package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextClient returns a translated phrase
func (l *Localizer) TextClient(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClient")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Client",
			Description: "the common phrase for client",
			One:         "Client",
			Other:       "Clients",
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
