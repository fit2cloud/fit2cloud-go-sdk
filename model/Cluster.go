package model

type Cluster struct {
	Id                        int64  `json:"id"`
	EnvType                   string `json:"envType"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	RabbitmqExchange          string `json:"rabbitmqExchange"`
	Created                   int64  `json:"created"`
	ServerNumber              int    `json:"serverNumber"`
	RoleNumber                int    `json:"roleNumber"`
	OnlineRunningServerNumber int    `json:"onlineRunningServerNumber"`
}
