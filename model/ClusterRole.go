package model

type ClusterRole struct {
	Id                        int64  `json:"id"`
	ClusterId                 int64  `json:"clusterId"`
	ClusterName               string `json:"clusterName"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	OnlineRunningServerNumber int    `json:"onlineRunningServerNumber"`
	VmNumber                  int    `json:"vmNumber"`
}
