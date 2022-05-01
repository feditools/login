package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextHome returns a translated phrase.
func (l *Localizer) TextHome(count int) *LocalizedString {
	lg := logger.WithField("func", "TextHomeShort")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HomeShort",
			Description: "a single word representation of home, as in home page.",
			One:         "Home",
			Other:       "Homes",
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
