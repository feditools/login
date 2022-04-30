package fedi

import (
	"context"
	"errors"
	"fmt"
	"github.com/feditools/login/internal/models"
	"net/url"
	"strings"
)

// GetLoginURL retrieves an oauth url for a federated instance
func (f *Fedi) GetLoginURL(ctx context.Context, act string) (*url.URL, error) {
	l := logger.WithField("func", "GetLoginURL")
	_, domain, err := SplitAccount(act)
	if err != nil {
		l.Errorf("split account: %s", err.Error())
		return nil, err
	}

	// try to get instance from the database
	instance, err := f.db.ReadFediInstanceByDomain(ctx, domain)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}
	if instance != nil {
		u, err := f.loginURLForInstance(ctx, instance)
		if err != nil {
			l.Errorf("get login url: %s", err.Error())
			return nil, err
		}
		return u, nil
	}

	// get nodeinfo endpoints from well-known location
	wkni, err := f.GetWellknownNodeInfo(ctx, domain)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())
		return nil, err
	}

	// check for nodeinfo 2.0 schema
	nodeinfoURI, err := findNodeInfo20URI(wkni)
	if err != nil {
		return nil, err
	}
	if nodeinfoURI == nil {
		return nil, errors.New("no nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, domain, nodeinfoURI)
	if err != nil {
		l.Errorf("get nodeinfo 2.0: %s", err.Error())
		return nil, err
	}

	// create instance for db
	newInstance := &models.FediInstance{
		Domain:         domain,
		ServerHostname: nodeinfoURI.Host,
		Software:       nodeinfo.Software.Name,
	}
	err = f.db.CreateFediInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		return nil, err
	}

	u, err := f.loginURLForInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("get login url: %s", err.Error())
		return nil, err
	}
	return u, nil
}

// SplitAccount splits a federated account into a username and domain
func SplitAccount(act string) (string, string, error) {
	actFragments := strings.Split(strings.ToLower(act), "@")

	switch len(actFragments) {
	case 2:
		return actFragments[0], actFragments[1], nil
	case 3:
		return actFragments[1], actFragments[2], nil
	default:
		return "", "", errors.New("invalid account format")
	}
}

func (f *Fedi) loginURLForInstance(ctx context.Context, instance *models.FediInstance) (*url.URL, error) {
	l := logger.WithField("func", "loginURLForInstance")

	if _, ok := f.helpers[Software(instance.Software)]; !ok {
		return nil, fmt.Errorf("no helper for '%s'", instance.Software)
	}

	if instance.ClientID == "" || len(instance.ClientSecret) == 0 {
		clientID, clientSecret, err := f.helpers[SoftwareMastodon].RegisterApp(ctx, instance)
		if err != nil {
			l.Errorf("registering app: %s", err.Error())
			return nil, err
		}
		l.Debugf("got app: %s, %s", clientID, clientSecret)
		instance.ClientID = clientID
		err = instance.SetClientSecret(clientSecret)
		if err != nil {
			l.Errorf("setting secret: %s", err.Error())
			return nil, err
		}

		err = f.db.UpdateFediInstance(ctx, instance)
		if err != nil {
			l.Errorf("db update: %s", err.Error())
			return nil, err
		}
	}

	return f.helpers[SoftwareMastodon].MakeLoginURL(ctx, instance)
}
