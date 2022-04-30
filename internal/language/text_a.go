package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextAccount returns a translated phrase.
func (l *Localizer) TextAccount(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAccount")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Account",
			Description: "the common phrase for account",
			One:         "Account",
			Other:       "Accounts",
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
