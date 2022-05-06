package models

// NodeInfo is a federated nodeinfo object
type NodeInfo struct {
	Links []Link `json:"links"`
}

// NodeInfo20 is a federated nodeinfo 2.0 object
type NodeInfo20 struct {
	Software struct {
		Name string `json:"name"`
	} `json:"software"`
}
