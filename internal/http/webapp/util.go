package webapp

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	nethttp "net/http"
	"net/http/httputil"

	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/web"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

// auth helpers

func (m *Module) authRequireLoggedIn(w nethttp.ResponseWriter, r *nethttp.Request) (*models.FediAccount, bool) {
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)

	if r.Context().Value(http.ContextKeyAccount) == nil {
		// Save current page
		if r.URL.Query().Encode() == "" {
			us.Values[http.SessionKeyLoginRedirect] = r.URL.Path
		} else {
			us.Values[http.SessionKeyLoginRedirect] = r.URL.Path + "?" + r.URL.Query().Encode()
		}
		err := us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return nil, true
		}

		// redirect to login
		nethttp.Redirect(w, r, path.Login, nethttp.StatusFound)

		return nil, true
	}
	account := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)

	return account, false
}

// language helpers

func getPageLang(query, header, defaultLang string) string {
	l := logger.WithField("func", "getPageLang")

	if query != "" {
		t, _, err := language.ParseAcceptLanguage(query)
		if err == nil {
			l.Debugf("query languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())

				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	if header != "" {
		t, _, err := language.ParseAcceptLanguage(header)
		if err == nil {
			l.Debugf("header languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())

				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	l.Debugf("returning default language: %s", defaultLang)

	return defaultLang
}

// signature caching

func getSignature(filePath string) (string, error) {
	l := logger.WithField("func", "getSignature")

	file, err := web.Files.Open(filePath)
	if err != nil {
		l.Errorf("opening file: %s", err.Error())

		return "", err
	}

	// read it
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// hash it
	h := sha512.New384()
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}

func (m *Module) getSignatureCached(filePath string) (string, error) {
	if sig, ok := m.readCachedSignature(filePath); ok {
		return sig, nil
	}
	sig, err := getSignature(filePath)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(filePath, sig)

	return sig, nil
}

func (m *Module) readCachedSignature(filePath string) (string, bool) {
	m.sigCacheLock.RLock()
	defer m.sigCacheLock.RUnlock()

	val, ok := m.sigCache[filePath]

	return val, ok
}

func (m *Module) writeCachedSignature(filePath string, sig string) {
	m.sigCacheLock.Lock()
	defer m.sigCacheLock.Unlock()

	m.sigCache[filePath] = sig
}

// debug.
func dumpRequest(l *logrus.Entry, header string, r *nethttp.Request) {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		l.Warnf("dump request: %s", err.Error())

		return
	}

	l.Debugf("%s: %s", header, string(data))
}
