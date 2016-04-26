package model

type LaunchConfiguration struct {
	Id                      int64  `json:"id"`
	GroupId                 int64  `json:"groupId"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	CredentialId            int64  `json:"credentialId"`
	CredentialName          string `json:"credentialName"`
	CloudProvider           string `json:"cloudProvider"`
	SshPort                 int    `json:"sshPort"`
	KeyPasswordId           int64  `json:"keyPasswordId"`
	LaunchConfiguration     string `json:"launchConfiguration"`
	Created                 int64  `json:"created"`
	AgentInstallMethod      string `json:"agentInstallMethod"`
	InitScriptId            int64  `json:"initScriptId"`
	KeyPasswordName         string `json:"keyPasswordName"`
	InitScriptName          string `json:"initScriptName"`
	Status                  string `json:"status"`
	ApplicationRevisionId   int64  `json:"applicationRevisionId"`
	ApplicationRevisionName string `json:"applicationRevisionName"`
	PluginDesc              string `json:"pluginDesc"`
	PluginIcon              string `json:"pluginIcon"`
}
