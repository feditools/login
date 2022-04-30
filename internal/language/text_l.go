package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextLogin returns a translated phrase
func (l *Localizer) TextLogin() *LocalizedString {
	lg := logger.WithField("func", "TextLogin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Login",
			Description: "the common phrase for logging in",
			Other:       "Login",
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

// TextLoginShort returns a translated phrase
func (l *Localizer) TextLoginShort() *LocalizedString {
	lg := logger.WithField("func", "TextLoginShort")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "LoginShort",
			Description: "a single word representation of logging in",
			Other:       "Login",
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
