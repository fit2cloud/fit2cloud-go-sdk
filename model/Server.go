package model

import (
//"fmt"
//"encoding/json"
)

type Server struct {
	Id                    int64  `json:"id"`
	Description           string `json:"description"`
	CluseterId            int64  `json:"clusterId"`
	ClusterName           string `json:"clusterName"`
	ClusterRoleId         int64  `json:"clusterRoleId"`
	ClusterRoleName       string `json:"clusterRoleName"`
	ImageId               string `json:"imageId"`
	VmId                  string `json:"vmId"`
	VmStatus              string `json:"vmStatus"`
	VmType                string `json:"vmType"`
	VmTypeDescription     string `json:"vmTypeDescription"`
	Status                string `json:"status"`
	Region                string `json:"region"`
	RemoteIP              string `json:"remoteIP"`
	LocalIP               string `json:"localIP"`
	Created               int64  `json:"created"`
	HeartbeatStatus       string `json:"heartbeatStatus"`
	Name                  string `json:"name"`
	AlertType             string `json:"alertType"`
	Zone                  string `json:"zone"`
	Hostname              string `json:"hostname"`
	KeyPasswordId         int64  `json:"keyPasswordId"`
	RabbitmqQueue         string `json:"rabbitmqQueue"`
	SshPort               int64  `json:"sshPort"`
	CustomData            string `json:"customData"`
	CredentialId          int64  `json:"credentialId"`
	LaunchConfigurationId int64  `json:"launchConfigurationId"`
	MachineId             string `json:"machineId"`
	Os                    string `json:"os"`
	AgentVersion          string `json:"agentVersion"`
	Tags                  []Tag  `json:"tags"`
}
