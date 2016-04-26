package fit2cloud_go_sdk

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	//"time"
)

var serverId = int64(43)

const (
	defaultkey      = "cmVuaC55YW5nQGhhaWhhbmd5dW4uY29t"
	defaultSecret   = "f8396069-c34c-4718-89fd-629b8bc31407"
	defaultEndpoint = "http://fit2cloud.haihangyun.net:8080/rest"
)

func TestNewClient(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)

	assert.NotNil(t, client)
}

func TestGetClusters(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	clusters, err := client.GetClusters()
	t.Logf("%v\n", clusters)
	assert.Nil(t, err)
}

func TestGetClusterRoles(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	clusterRoles, err := client.GetClusterRoles(5)
	t.Logf("%v\n", clusterRoles)
	assert.Nil(t, err)
}
func TestGetCloudCredential(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	creds, err := client.GetCloudCredential()
	t.Logf("%v\n", creds)
	assert.Nil(t, err)
}
func TestGetLaunchconfiguration(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	cnf, err := client.GetLaunchconfiguration(2)
	t.Logf("%v\n", cnf)
	assert.Nil(t, err)
}
/*func TestLaunchServer(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	starttime := time.Now()
	server, err := client.LaunchServer(5, 5, 7)
	endtime := time.Now()
	t.Error(starttime.Format("2016-01-21-13-10-10"))
	t.Error(endtime.Format("2016-01-21-13-10-10"))
	if err != nil {
	} else {
		serverId = server.Id
	}
	t.Logf("%v\n", server)
	assert.Nil(t, err)
}*/
func TestStopServer(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	res := client.StopServer(serverId)
	assert.True(t, res)
}
func TestStartServer(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	server, err := client.StartServer(serverId)
	t.Logf("%v\n", server)
	assert.Nil(t, err)
}
func TestTerminateServer(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	res := client.TerminateServer(serverId)
	assert.True(t, res)
}
func TestGetServers(t *testing.T) {
	client := NewClient(defaultkey, defaultSecret, defaultEndpoint)
	servers, err := client.GetServers(5, 5, "", "", 0, 0, false)
	t.Logf("%v\n", servers)
	assert.Nil(t, err)
}
