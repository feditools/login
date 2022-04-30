package template

import (
	"github.com/feditools/login/internal/language"
)

// Common contains the variables used in nearly every template
type Common struct {
	Language  string
	Localizer *language.Localizer

	FooterScripts []Script
	HeadLinks     []HeadLink
	NavBar        Navbar
	PageTitle     string
	AccountID     int64
}

// AddHeadLink adds a headder link to the template
func (t *Common) AddHeadLink(l HeadLink) {
	if t.HeadLinks == nil {
		t.HeadLinks = []HeadLink{}
	}
	t.HeadLinks = append(t.HeadLinks, l)
	return
}

// AddFooterScript adds a footer script to the template
func (t *Common) AddFooterScript(s Script) {
	if t.FooterScripts == nil {
		t.FooterScripts = []Script{}
	}
	t.FooterScripts = append(t.FooterScripts, s)
	return
}

// SetLanguage sets the template's default language
func (t *Common) SetLanguage(l string) {
	t.Language = l
	return
}

// SetLocalizer sets the localizer the template will use to generate text
func (t *Common) SetLocalizer(l *language.Localizer) {
	t.Localizer = l
	return
}

// SetNavbar sets the top level navbar used by the template
func (t *Common) SetNavbar(nodes Navbar) {
	t.NavBar = nodes
	return
}

// SetAccountID sets the currently logged-in account
func (t *Common) SetAccountID(accountID int64) {
	t.AccountID = accountID
	return
}
