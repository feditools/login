package fedi

import (
	"context"
	"errors"
	"github.com/feditools/login/internal/models"
)

// GenerateFediInstanceFromDomain created a FediInstance object by querying the apis of the federated instance
func (f *Fedi) GenerateFediInstanceFromDomain(ctx context.Context, domain string) (*models.FediInstance, error) {
	l := logger.WithField("func", "GenerateFediInstanceFromDomain")

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
		return nil, errors.New("missing nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, domain, nodeinfoURI)
	if err != nil {
		l.Errorf("get nodeinfo 2.0: %s", err.Error())
		return nil, err
	}

	// get actor uri
	webfinger, err := f.GetWellknownWebFinger(ctx, domain, domain)
	if err != nil {
		return nil, err
	}
	actorURI, err := FindActorURI(webfinger)
	if err != nil {
		return nil, err
	}
	if actorURI == nil {
		return nil, errors.New("missing actor uri")
	}

	return &models.FediInstance{
		Domain:         domain,
		ActorURI:       actorURI.String(),
		ServerHostname: nodeinfoURI.Host,
		Software:       nodeinfo.Software.Name,
	}, nil
}
