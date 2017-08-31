package pmapi

import (
	"fmt"
	"net/http"
	"proxysql-master/admin/servers"
	"strconv"

	"github.com/labstack/echo"
)

/*返回所有后端数据库服务器的信息*/
func (pmapi *PMApi) ListAllServers(c echo.Context) error {
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit
	return c.JSON(http.StatusOK, servers.FindAllServerInfo(pmapi.Apidb, limit, skip))
}

/*查询指定主机组中主机的信息*/
func (pmapi *PMApi) ListServerByHostgroup(c echo.Context) error {
	server := new(servers.Servers)

	server.HostGroupId, _ = strconv.ParseUint(c.Param("hostgroup"), 10, 64)
	fmt.Println(server.HostGroupId)
	return c.JSON(http.StatusOK, server.FindServersInfoByHostgroup(pmapi.Apidb))
}

//通过参数主机组、主机名和端口查出一个主机的信息
func (pmapi *PMApi) ListOneServer(c echo.Context) error {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
	}{}

	server := new(servers.Servers)
	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port

	cret := server.FindOneServersInfo(pmapi.Apidb)
	return c.JSON(http.StatusOK, cret)
}

/*创建一个新的后端数据库服务节点*/
func (pmapi *PMApi) CreateServer(c echo.Context) error {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port

	cret := server.AddOneServers(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "CreateServer Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "Server exists")
	default:
		return c.JSON(http.StatusOK, "CreateServer ???")

	}
}

/*更新一个后端服务的状态*/
func (pmapi *PMApi) UpdateOneServerStatus(c echo.Context) error {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
		Status      string `json:"status"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port
	server.Status = args.Status

	switch server.Status {
	case "SOFT_OFFLINE":
		cret := server.SoftDisactiveOneServer(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "SoftDisactiveOneServer Failed")
		case 2:
			return c.JSON(http.StatusExpectationFailed, "Server not exists")
		default:
			return c.JSON(http.StatusExpectationFailed, "SoftDisactiveOneServer other return value")
		}

	case "HARD_OFFLINE":
		cret := server.HardDisactiveOneServer(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "HardDisactiveOneServer Failed")
		case 2:
			return c.JSON(http.StatusExpectationFailed, "Server not exists")
		default:
			return c.JSON(http.StatusExpectationFailed, "HardDisactiveOneServer other return value")
		}

	case "ONLINE":
		cret := server.ActiveOneServer(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "ActiveOneServer Failed")
		case 2:
			return c.JSON(http.StatusExpectationFailed, "Server not exists")
		default:
			return c.JSON(http.StatusExpectationFailed, "ActiveOneServer other return value")
		}
	default:
		return c.JSON(http.StatusOK, "UpdateOneServerStatus other status")
	}
}

/*更改指定后端服务器的权重*/
func (pmapi *PMApi) UpdateOneServerWeight(c echo.Context) error {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
		Weight      uint64 `json:"weight"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port
	server.Weight = args.Weight

	cret := server.UpdateOneServerWeight(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneServerWeight Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "Server not exists")
	default:
		return c.JSON(http.StatusOK, "UpdateOneServerWeight ???")
	}
}

/*更改指定服务器的最大连接数*/
func (pmapi *PMApi) UpdateOneServerMC(c echo.Context) error {
	args := struct {
		HostGroupId    uint64 `json:"hostgroup_id"`
		HostName       string `json:"hostname"`
		Port           uint64 `json:"port"`
		MaxConnections uint64 `json:"max_connections"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port
	server.MaxConnections = args.MaxConnections

	cret := server.UpdateOneServerMc(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneServerMc Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "Server not exists")
	default:
		return c.JSON(http.StatusOK, "UpdateOneServerMC ???")

	}
}

/*删除指定服务器*/
func (pmapi *PMApi) DeleteOneServers(c echo.Context) error {
	args := struct {
		HostGroupId uint64 `json:"hostgroup_id"`
		HostName    string `json:"hostname"`
		Port        uint64 `json:"port"`
	}{}

	server := new(servers.Servers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	server.HostGroupId = args.HostGroupId
	server.HostName = args.HostName
	server.Port = args.Port

	cret := server.DeleteOneServers(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "DeleteOneServer Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "Server not exists")
	default:
		return c.JSON(http.StatusOK, "DeleteOneServers ???")

	}
}

/*更新服务信息的patch函数*/
func (pmapi *PMApi) UpdateOneServerInfo(c echo.Context) error {
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
		return err
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

	return c.JSON(http.StatusOK, "OK")
}
