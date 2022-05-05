package models

// NodeInfo is a federated nodeinfo object
type NodeInfo struct {
	Links []struct {
		Rel  string `json:"rel"`
		HRef string `json:"href"`
	} `json:"links"`
}

// NodeInfo20 is a federated nodeinfo 2.0 object
type NodeInfo20 struct {
	Software struct {
		Name string `json:"name"`
	} `json:"software"`
}
