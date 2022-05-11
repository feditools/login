package template

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// Common contains the variables used in nearly every template
type Common struct {
	Language  string
	Localizer *liblanguage.Localizer

	Account *models.FediAccount

	Alerts        *[]libtemplate.Alert
	FooterScripts []libtemplate.Script
	HeadLinks     []libtemplate.HeadLink
	LogoSrcDark   string
	LogoSrcLight  string
	NavBar        Navbar
	PageTitle     string
}

// AddHeadLink adds a headder link to the template
func (t *Common) AddHeadLink(l libtemplate.HeadLink) {
	if t.HeadLinks == nil {
		t.HeadLinks = []libtemplate.HeadLink{}
	}
	t.HeadLinks = append(t.HeadLinks, l)
	return
}

// AddFooterScript adds a footer script to the template
func (t *Common) AddFooterScript(s libtemplate.Script) {
	if t.FooterScripts == nil {
		t.FooterScripts = []libtemplate.Script{}
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
func (t *Common) SetLocalizer(l *liblanguage.Localizer) {
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
