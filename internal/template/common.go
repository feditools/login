package template

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/login/internal/models"
)

// Common contains the variables used in nearly every template
type Common struct {
	Language  string
	Localizer *language.Localizer

	Account       *models.FediAccount
	FooterScripts []Script
	HeadLinks     []HeadLink
	LogoSrcDark   string
	LogoSrcLight  string
	NavBar        Navbar
	PageTitle     string
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

// SetLogoSrc sets the src for the logo image
func (t *Common) SetLogoSrc(dark, light string) {
	t.LogoSrcDark = dark
	t.LogoSrcLight = light
	return
}

// SetNavbar sets the top level navbar used by the template
func (t *Common) SetNavbar(nodes Navbar) {
	t.NavBar = nodes
	return
}

// SetAccount sets the currently logged-in account
func (t *Common) SetAccount(account *models.FediAccount) {
	t.Account = account
	return
}
