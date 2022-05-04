package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextAccount returns a translated phrase
func (l *Localizer) TextAccount(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAccount")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Account",
			One:   "Account",
			Other: "Accounts",
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

// TextAddOauth20Client returns a translated phrase
func (l *Localizer) TextAddOauth20Client(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAddOauth20Client")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddOauth20Client",
			One:   "Add OAuth 2.0 Client",
			Other: "Add OAuth 2.0 Clients",
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
