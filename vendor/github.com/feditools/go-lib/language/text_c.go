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

// TextClientID returns a translated phrase
func (l *Localizer) TextClientID(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClientID")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "ClientID",
			Description: "the common phrase for client IDs",
			One:         "Client ID",
			Other:       "Client IDs",
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

// TextClientSecret returns a translated phrase
func (l *Localizer) TextClientSecret(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClientSecret")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "ClientSecret",
			Description: "the common phrase for client secrets",
			One:         "Client Secret",
			Other:       "Client Secrets",
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

// TextCreate returns a translated phrase
func (l *Localizer) TextCreate() *LocalizedString {
	lg := logger.WithField("func", "TextCreate")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Create",
			Description: "the common phrase for create",
			Other:       "Create",
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
