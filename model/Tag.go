package model

type Tag struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Value           string `json:"value"`
	ServerId        int64  `json:"serverId"`
	ClusterRoleId   int64  `json:"clusterRoleId`
	ClusterId       int64  `json:"clusterId"`
	ServerName      string `json:"serverName"`
	ClusterName     string `json:"clusterName"`
	ClusterRoleName string `json:"clusterRoleName"`
}
