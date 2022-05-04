package template

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/token"
	"html/template"
	"io/ioutil"
	"strings"
)

const templateDir = "web/template"

// InitTemplate are the functions a template implementing Common will have
type InitTemplate interface {
	AddHeadLink(l libtemplate.HeadLink)
	AddFooterScript(s libtemplate.Script)
	SetAccount(account *models.FediAccount)
	SetLanguage(l string)
	SetLocalizer(l *language.Localizer)
	SetLogoSrc(dark, light string)
	SetNavbar(nodes Navbar)
}

// New creates a new template
func New(t *token.Tokenizer) (*template.Template, error) {
	tpl := template.New("")
	tpl.Funcs(template.FuncMap{
		"dec": func(i int) int {
			i--
			return i
		},
		"htmlSafe": func(html string) template.HTML {
			/* #nosec G203 */
			return template.HTML(html)
		},
		"inc": func(i int) int {
			i++
			return i
		},
		"token": t.GetToken,
	})

	dir, err := login.Files.ReadDir(templateDir)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		filePath := templateDir + "/" + d.Name()
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".gohtml") {
			continue
		}

		// open it
		file, err := login.Files.Open(filePath)
		if err != nil {
			return nil, err
		}

		// read it
		tmplData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(tmplData))
		if err != nil {
			return nil, err
		}
	}

	return tpl, nil
}
