package fedi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feditools/login/internal/fedi/models"
	"github.com/feditools/login/internal/http"
)

// GetWellknownNodeInfo retrieves wellknown nodeinfo from a federated instance
func (f *Fedi) GetWellknownNodeInfo(ctx context.Context, domain string) (*models.NodeInfo, error) {
	l := logger.WithField("func", "GetWellknownNodeInfo")
	v, err, _ := f.nodeinfoRequestGroup.Do(domain, func() (interface{}, error) {
		// get nodeinfo
		resp, err := http.Get(ctx, fmt.Sprintf("https://%s/.well-known/nodeinfo", domain))
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
