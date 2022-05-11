package fedi

import (
	"context"
	"encoding/json"
	"fmt"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/login/internal/fedi/models"
	"github.com/feditools/login/internal/http"
	"net/url"
)

// GetWellknownNodeInfo retrieves wellknown nodeinfo from a federated instance
func (f *Fedi) GetWellknownNodeInfo(ctx context.Context, domain string) (*models.NodeInfo, error) {
	l := logger.WithField("func", "GetWellknownNodeInfo")
	nodinfoURI := fmt.Sprintf("https://%s/.well-known/nodeinfo", domain)
	v, err, _ := f.requestGroup.Do(nodinfoURI, func() (interface{}, error) {
		// do request
		resp, err := http.Get(ctx, nodinfoURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())
			return nil, err
		}

		nodeinfo := new(models.NodeInfo)
		err = json.NewDecoder(resp.Body).Decode(nodeinfo)
		if err != nil {
			l.Errorf("decode json: %s", err.Error())
			return nil, err
		}

		return nodeinfo, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())
		return nil, err
	}

	nodeinfo := v.(*models.NodeInfo)
	return nodeinfo, nil
}

// GetWellknownWebFinger retrieves wellknown web finger resource from a federated instance
func (f *Fedi) GetWellknownWebFinger(ctx context.Context, username, domain string) (*models.WebFinger, error) {
	l := logger.WithField("func", "GetWellknownWebFinger")
	webfingerURI := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s@%s", domain, username, domain)
	v, err, _ := f.requestGroup.Do(webfingerURI, func() (interface{}, error) {
		// do request
		resp, err := http.Get(ctx, webfingerURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())
			return nil, err
		}

		webfinger := new(models.WebFinger)
		err = json.NewDecoder(resp.Body).Decode(webfinger)
		if err != nil {
			l.Errorf("decode json: %s", err.Error())
			return nil, err
		}

		return webfinger, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())
		return nil, err
	}

	webfinger := v.(*models.WebFinger)
	return webfinger, nil
}

// findActorURI parses a webfinger document for an actor uri
func findActorURI(webfinger *models.WebFinger) (*url.URL, error) {
	var actorURIstr string
	for _, link := range webfinger.Links {
		if link.Rel == "self" || link.Type == libhttp.MimeAppActivityJSON {
			actorURIstr = link.HRef
			break
		}
	}
	if actorURIstr == "" {
		return nil, nil
	}

	actorURI, err := url.Parse(actorURIstr)
	if err != nil {
		return nil, fmt.Errorf("invalid actor uri: %s", err.Error())
	}

	return actorURI, err
}
