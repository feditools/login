package fedi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	nethttp "net/http"
	"net/url"

	"github.com/feditools/login/internal/fedi/models"
	"github.com/feditools/login/internal/http"
)

// findNodeInfo20URI parses a nodeinfo document for a nodeinfo 2.0 uri.
func findNodeInfo20URI(nodeinfo *models.NodeInfo) (*url.URL, error) {
	var nodeinfoURIstr string
	for _, link := range nodeinfo.Links {
		if link.Rel == NodeInfo20Schema {
			nodeinfoURIstr = link.HRef

			break
		}
	}
	if nodeinfoURIstr == "" {
		return nil, nil
	}

	nodeinfoURI, err := url.Parse(nodeinfoURIstr)
	if err != nil {
		return nil, fmt.Errorf("invalid nodeinfo 2.0 uri: %s", err.Error())
	}

	return nodeinfoURI, err
}

// GetNodeInfo20 retrieves wellknown nodeinfo from a federated instance.
func (f *Fedi) GetNodeInfo20(ctx context.Context, domain string, infoURI *url.URL) (*models.NodeInfo2, error) {
	l := logger.WithField("func", "GetNodeInfo20")
	v, err, _ := f.requestGroup.Do(infoURI.String(), func() (interface{}, error) {
		// check cache
		cache, err := f.kv.GetFediNodeInfo(ctx, domain)
		if err != nil && err.Error() != "redis: nil" {
			l.Errorf("redis get: %s", err.Error())

			return nil, err
		}
		if err == nil {
			return unmarshalNodeInfo20(cache)
		}

		// get nodeinfo
		resp, err := http.Get(ctx, infoURI.String())
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}
		if resp.StatusCode != nethttp.StatusOK {
			return nil, fmt.Errorf("http status %s %d", infoURI, resp.StatusCode)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			l.Errorf("read body: %s", err.Error())

			return nil, err
		}
		bodyString := string(bodyBytes)

		// marshal
		nodeinfo, err := unmarshalNodeInfo20(bodyString)
		if err != nil {
			l.Errorf("marshal: %s", err.Error())

			return nil, err
		}

		// write cache
		err = f.kv.SetFediNodeInfo(ctx, domain, bodyString, f.nodeinfoCacheExp)
		if err != nil {
			l.Errorf("redis get: %s", err.Error())

			return nil, err
		}

		return nodeinfo, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	nodeinfo := v.(*models.NodeInfo2)

	return nodeinfo, nil
}

func unmarshalNodeInfo20(body string) (*models.NodeInfo2, error) {
	var nodeinfo *models.NodeInfo2
	if err := json.Unmarshal([]byte(body), &nodeinfo); err != nil {
		return nil, err
	}

	return nodeinfo, nil
}
