package model

type CloudCredential struct {
	Id            int64  `json:"id"`
	GroupId       int64  `json:"groupId"`
	Name          string `json:"name"`
	Credential    string `json:"credential"`
	Status        string `json:"status"`
	Created       int64  `json:"created"`
	ServerCount   int64  `json:"serverCount"`
	LbCount       int    `json:"lbCount"`
	RdsCount      int    `json:"rdsCount"`
	SyncStatus    string `json:"SyncStatus"`
	LastUpdate    int64  `json:"LastUpdate"`
	CloudProvider string `json:"cloudProvider"`
	CustomData    string `json:"customData"`
	Description   string `json:"description"`
	PluginDes     string `json:"pluginDesc"`
	PluginIcon    string `json:"pluginIcon"`
	ImageCount    int64  `json:"imageCount"`
}
