package fit2cloud_go_sdk

import (
	"crypto"
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
	. "renh.yang/fit2cloud-go-sdk/model"
	//"fmt"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Fit2CloudClient struct {
	c          *oauth.Consumer
	RootApiUrl string
}

func NewClient(consumerKey, secret, restApiUrl string) *Fit2CloudClient {
	provider := oauth.ServiceProvider{
		RequestTokenUrl:   "",
		AuthorizeTokenUrl: "",
		AccessTokenUrl:    "",
		BodyHash:          true,
	}
	consumer := oauth.NewCustomConsumer(consumerKey, secret, crypto.SHA1, provider, nil)
	return &Fit2CloudClient{
		c:          consumer,
		RootApiUrl: restApiUrl,
	}

}

/**
 * 获取当前用户所有集群信息
 *
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) GetClusters() ([]Cluster, error) {
	geturl := client.RootApiUrl + "/clusters"

	cli, err := client.c.MakeHttpClient(&oauth.AccessToken{})
	if err != nil {
		return nil, err
	}

	resp, err := cli.Get(geturl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	var clusters []Cluster
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &clusters)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}
	return clusters, nil
}

/* 获取指定集群下所有虚机组信息
*
* @param clusterId	集群ID
* @return
* @throws Fit2CloudException
 */
func (client *Fit2CloudClient) GetClusterRoles(clusterId int64) ([]ClusterRole, error) {
	strcid := strconv.FormatInt(clusterId, 10)
	geturl := client.RootApiUrl + "/cluster/" + strcid + "/roles"

	cli, err := client.c.MakeHttpClient(&oauth.AccessToken{})
	if err != nil {
		return nil, err
	}

	resp, err := cli.Get(geturl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	var clusterRoles []ClusterRole
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &clusterRoles)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}
	return clusterRoles, nil
}

/**
 * 创建虚机
 *
 * @param clusterId	虚机所在的集群
 * @param clusterRoleId	虚机所在的虚机组
 * @param launchConfigurationId	创建虚机所使用的模板
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) LaunchServer(clusterId, clusterRoleId, launchConfigurationId int64) (*Server, error) {
	strcid := strconv.FormatInt(clusterId, 10)
	strcrid := strconv.FormatInt(clusterRoleId, 10)
	strlcid := strconv.FormatInt(launchConfigurationId, 10)
	launchUrl := client.RootApiUrl + "/launchserver/cluster/" + strcid + "/clusterrole/" + strcrid
	//cli, err := client.c.MakeHttpClient(&oauth.AccessToken{})
	transport, err := client.c.MakeRoundTripper(&oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	cli := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(10 * time.Minute),
	}

	form := url.Values{"launchConfigurationId": {strlcid}}
	resp, err := cli.PostForm(launchUrl, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// process the response
	if resp.StatusCode == 200 {
		var server Server
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &server)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}

		return &server, nil
	} else {
		return nil, errors.New(resp.Status)
	}
}

/**
 * 创建虚机. 此接口将立刻返回虚机ID,而后台将异步创建虚机. 之后通过返回的虚机信息中的ID来查询完整的虚机信息
 *
 * @param clusterId	虚机所在的集群
 * @param clusterRoleId	虚机所在的虚机组
 * @param launchConfigurationId	创建虚机所使用的模板
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) LaunchServerAsync(clusterId, clusterRoleId, launchConfigurationId int64) (*Server, error) {
	strcid := strconv.FormatInt(clusterId, 10)
	strcrid := strconv.FormatInt(clusterRoleId, 10)
	strlcid := strconv.FormatInt(launchConfigurationId, 10)
	launchUrl := client.RootApiUrl + "/launchserver/async/cluster/" + strcid + "/clusterrole/" + strcrid
	cli, err := client.c.MakeHttpClient(&oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	form := url.Values{"launchConfigurationId": {strlcid}}
	resp, err := cli.PostForm(launchUrl, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// process the response
	if resp.StatusCode == 200 {
		var server Server
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &server)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}

		return &server, nil
	} else {
		return nil, errors.New(resp.Status)
	}
}

/**
 * 获取虚机创建模板列表
 *
 * @param cloudCredentialId	可选,云帐号ID
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) GetLaunchconfiguration(cloudCredentialId int64) ([]LaunchConfiguration, error) {
	geturl := client.RootApiUrl + "/launchconfigurations"

	/*cli,err := client.c.MakeHttpClient(&oauth.AccessToken{})
	if err!= nil{
		return nil,err
	} */
	form := make(map[string]string)
	if cloudCredentialId > 0 {
		form["cloudCredentialId"] = strconv.FormatInt(cloudCredentialId, 10)
	}
	resp, err := client.c.Get(geturl, form, &oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var config []LaunchConfiguration
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &config)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		return config, nil

	} else {
		return nil, errors.New(resp.Status)
	}
}

/**
 * 获取云帐号列表
 *
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) GetCloudCredential() ([]CloudCredential, error) {
	geturl := client.RootApiUrl + "/cloudcredentials"

	/*cli,err := client.c.MakeHttpClient(&oauth.AccessToken{})
	if err!= nil{
		return nil,err
	} */
	resp, err := client.c.Get(geturl, nil, &oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var credentials []CloudCredential
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &credentials)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		return credentials, nil

	} else {
		return nil, errors.New(resp.Status)
	}

}

/**
 * 删除虚机
 *
 * @param serverId	虚机ID
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) TerminateServer(serverId int64) bool {
	geturl := client.RootApiUrl + "/terminateserver/server/" + strconv.FormatInt(serverId, 10)

	resp, err := client.c.Post(geturl, nil, &oauth.AccessToken{})
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false
		}
		if "true" == string(body) {
			return true
		}
		return false
	} else {
		return false
	}

}

/**
 * 启动处于停止中状态的虚机
 *
 * @param serverId	虚机ID
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) StartServer(serverId int64) (*Server, error) {
	launchUrl := client.RootApiUrl + "/startserver/server/" + strconv.FormatInt(serverId, 10)
	resp, err := client.c.PostForm(launchUrl, nil, &oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// process the response
	if resp.StatusCode == 200 {
		var server Server
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &server)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}

		return &server, nil
	} else {
		return nil, errors.New(resp.Status)
	}
}

/**
 * 停止虚机
 *
 * @param serverId	虚机ID
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) StopServer(serverId int64) bool {
	launchUrl := client.RootApiUrl + "/stopserver/server/" + strconv.FormatInt(serverId, 10)
	resp, err := client.c.PostForm(launchUrl, nil, &oauth.AccessToken{})
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	// process the response
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false
		}
		if "true" == string(body) {
			return true
		}
		return false
	} else {
		return false
	}
}

/**
 * 所有参数均非必须, 若所有参数均为null, 则返回当前用户拥有的所有虚机
 *
 * @param clusterId	集群ID,(可选)
 * @param clusterRoleId	虚机组ID,(可选)
 * @param sort	排序字段,(可选)
 * @param order	排序方式,(可选)
 * @param pageSize	分页大小,(可选,默认9999)
 * @param pageNum	分页编号,(可选,默认1)
 * @param showTerminated	是否显示已关闭虚机
 * @return
 * @throws Fit2CloudException
 */
func (client *Fit2CloudClient) GetServers(clusterId int64, clusterRole int64, sort string, order string, pageSize int, pageNum int, showTerminated bool) ([]Server, error) {
	geturl := client.RootApiUrl + "/servers"

	form := make(map[string]string)
	if clusterId > 0 {
		form["clusterId"] = strconv.FormatInt(clusterId, 10)
	}
	if clusterRole > 0 {
		form["clusterRoleId"] = strconv.FormatInt(clusterRole, 10)
	}

	if sort != "" {
		form["sort"] = sort
	}

	if order != "" {
		form["order"] = order
	}

	if pageSize > 0 {
		form["pageSize"] = strconv.Itoa(pageSize)
	}

	if pageSize > 0 {
		form["pageNum"] = strconv.Itoa(pageNum)
	}

	form["showTerminated"] = strconv.FormatBool(showTerminated)
	resp, err := client.c.Get(geturl, form, &oauth.AccessToken{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var servers []Server
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &servers)
		if err != nil {
			//panic(err.Error())
			return nil, err
		}
		return servers, nil

	} else {
		return nil, errors.New(resp.Status)
	}
}
