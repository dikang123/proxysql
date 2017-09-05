package pmapi

import (
	"Pdbs/src/admin/users"
	"database/sql"
	"net/http"
	"proxysql-master/admin/servers"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*返回所有后端数据库服务器的信息*/
func (pmapi *PMApi) ListAllServers(c *gin.Context) {

	var tmpserver servers.Servers
	var aryservers []servers.Servers

	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []users.Users{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		defer pmapi.Apidb.Close()

		aryservers, err = tmpserver.FindAllServerInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, aryusr)
	}

}

/*创建一个新的后端数据库服务节点*/
func (pmapi *PMApi) CreateOneServer(c *gin.Context) {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port

	cret := server.AddOneServers(pmapi.Apidb)
	switch cret {
	case 0:
		c.JSON(http.StatusOK, "OK")
	case 1:
		c.JSON(http.StatusExpectationFailed, "CreateServer Failed")
	case 2:
		c.JSON(http.StatusExpectationFailed, "Server exists")
	default:
		c.JSON(http.StatusOK, "CreateServer ???")

	}
}

/*删除指定服务器*/
func (pmapi *PMApi) DeleteOneServers(c *gin.Context) {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port

	cret := server.DeleteOneServers(pmapi.Apidb)
	switch cret {
	case 0:
		c.JSON(http.StatusOK, "OK")
	case 1:
		c.JSON(http.StatusExpectationFailed, "DeleteOneServer Failed")
	case 2:
		c.JSON(http.StatusExpectationFailed, "Server not exists")
	default:
		c.JSON(http.StatusOK, "DeleteOneServers ???")

	}
}

/*更新服务信息的patch函数*/
func (pmapi *PMApi) UpdateOneServer(c *gin.Context) {
	args := struct {
		HostGroupId       uint64 `json:"hostgroup_id"`
		HostName          string `json:"hostname"`
		Port              uint64 `json:"port"`
		Status            string `json:"status"`
		Weight            uint64 `json:"weight"`
		Compression       uint64 `json:"compression"`
		MaxConnections    uint64 `json:"max_connections"`
		MaxReplicationLag uint64 `json:"max_replication_lag"`
		UseSsl            uint64 `json:"use_ssl"`
		MaxLatencyMs      uint64 `json:"max_latency_ms"`
		Comment           string `json:"comment"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port
	server.Status = args.Status
	server.Weight = args.Weight
	server.Compression = args.Compression
	server.MaxConnections = args.MaxConnections
	server.MaxReplicationLag = args.MaxReplicationLag
	server.UseSsl = args.UseSsl
	server.MaxLatencyMs = args.MaxLatencyMs
	server.Comment = args.Comment

	server.UpdateOneServerInfo(pmapi.Apidb)

	c.JSON(http.StatusOK, "OK")
}
