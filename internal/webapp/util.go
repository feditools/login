package webapp

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/feditools/login"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/gorilla/sessions"
	"golang.org/x/text/language"
	"io/ioutil"
	nethttp "net/http"
)

// auth helpers

func (m *Module) authRequireLoggedIn(w nethttp.ResponseWriter, r *nethttp.Request) (*models.FediAccount, bool) {
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)

	if r.Context().Value(http.ContextKeyAccount) == nil {
		// Save current page
		if r.URL.Query().Encode() == "" {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path
		} else {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path + "?" + r.URL.Query().Encode()
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

func getSignature(path string) (string, error) {
	l := logger.WithField("func", "getSignature")

	file, err := login.Files.Open(path)
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
	h.Write(data)
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}

func (m *Module) getSignatureCached(path string) (string, error) {
	if sig, ok := m.readCachedSignature(path); ok {
		return sig, nil
	}
	sig, err := getSignature(path)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(path, sig)
	return sig, nil
}

func (m *Module) readCachedSignature(path string) (string, bool) {
	m.sigCacheLock.RLock()
	defer m.sigCacheLock.RUnlock()

	val, ok := m.sigCache[path]
	return val, ok
}

func (m *Module) writeCachedSignature(path string, sig string) {
	m.sigCacheLock.Lock()
	defer m.sigCacheLock.Unlock()

	m.sigCache[path] = sig
	return
}
