package pmapi

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"proxysql-master/admin/queryrules"
	"proxysql-master/admin/servers"
	"proxysql-master/admin/status"
	"proxysql-master/admin/users"
	"proxysql-master/admin/variables"
	"strconv"
)

type PMApi struct {
	PMuser    string
	PMpass    string
	PMhost    string
	PMdb      string
	PMdbi     string
	Apidb     *sql.DB
	ApiHost   string
	ApiLogfd  *os.File
	ApiLogcwd string
	ApiErr    error
	*echo.Echo
}

func (pmapi *PMApi) MakePMdbi() {
	pmapi.PMdbi = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", pmapi.PMuser, pmapi.PMpass, pmapi.PMhost, pmapi.PMdb)
}

func (pmapi *PMApi) RegisterDBInterface() {
	var err error
	pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
	if err != nil {
		log.Print("sql.Open()", err)
	}
}

func (pmapi *PMApi) DestoryDBInterface() {
	defer pmapi.Apidb.Close()
}

func (pmapi *PMApi) RegisterMiddleware() {
	pmapi.Echo.Use(mw.Logger())
	pmapi.Echo.Use(mw.Recover())
}

func (pmapi *PMApi) RegisterServices() {
	/*Dashboard*/
	pmapi.Echo.GET("/api/v1/status", pmapi.ListPStatus)
	pmapi.Echo.GET("/api/v1/variables", pmapi.ListPsVariables)

	/*User Services*/
	pmapi.Echo.GET("/api/v1/users", pmapi.ListAllUsers)
	pmapi.Echo.GET("/api/v1/users/:username", pmapi.ListOneUser)
	pmapi.Echo.POST("/api/v1/users", pmapi.CreateUser)
	pmapi.Echo.PUT("/api/v1/users/passwd", pmapi.UpdateOneUserPass)
	pmapi.Echo.PUT("/api/v1/users/status", pmapi.UpdateOneUserStatus)
	pmapi.Echo.PUT("/api/v1/users/hostgroup", pmapi.UpdateOneUserDH)
	pmapi.Echo.PUT("/api/v1/users/schema", pmapi.UpdateOneUserDS)
	pmapi.Echo.PUT("/api/v1/users/maxconnection", pmapi.UpdateOneUserMC)
	pmapi.Echo.DELETE("/api/v1/users/:username", pmapi.DeleteOneUser)

	/*Server Services*/
	pmapi.Echo.GET("/api/v1/servers", pmapi.ListAllServers)
	pmapi.Echo.GET("/api/v1/servers/:hostgroup", pmapi.ListServerByHostgroup)
	pmapi.Echo.PUT("/api/v1/servers", pmapi.ListOneServer)
	pmapi.Echo.POST("/api/v1/servers", pmapi.CreateServer)
	pmapi.Echo.PUT("/api/v1/servers/status", pmapi.UpdateOneServerStatus)
	pmapi.Echo.PUT("/api/v1/servers/weight", pmapi.UpdateOneServerWeight)
	pmapi.Echo.PUT("/api/v1/servers/maxconnection", pmapi.UpdateOneServerMC)
	pmapi.Echo.DELETE("/api/v1/servers", pmapi.DeleteOneServers)

	/*Query Rules*/
	pmapi.Echo.GET("/api/v1/queryrules", pmapi.ListAllQueryRules)
	pmapi.Echo.GET("/api/v1/queryrules/:ruleid", pmapi.ListOneQueryRule)
	pmapi.Echo.POST("/api/v1/queryrules", pmapi.CreateQueryRules)
	pmapi.Echo.PUT("/api/v1/queryrules/status/:ruleid", pmapi.UpdateOneQueryRulesStatus)
	pmapi.Echo.PUT("/api/v1/queryrules/username/:ruleid", pmapi.UpdateOneQueryRulesUser)
	pmapi.Echo.PUT("/api/v1/queryrules/schemaname/:ruleid", pmapi.UpdateOneQueryRulesSchema)
	pmapi.Echo.PUT("/api/v1/queryrules/clientaddr/:ruleid", pmapi.UpdateOneQueryRulesClient)
	pmapi.Echo.PUT("/api/v1/queryrules/matchdigest/:ruleid", pmapi.UpdateOneQueryRulesMatchDigest)
	pmapi.Echo.PUT("/api/v1/queryrules/matchpattern/:ruleid", pmapi.UpdateOneQueryRulesMatchPattern)
	pmapi.Echo.PUT("/api/v1/queryrules/replacepattern/:ruleid", pmapi.UpdateOneQueryRulesReplacePattern)
	pmapi.Echo.PUT("/api/v1/queryrules/desthostgroup/:ruleid", pmapi.UpdateOneQueryRulesDestHostgroup)
	pmapi.Echo.PUT("/api/v1/queryrules/errmsg/:ruleid", pmapi.UpdateOneQueryRulesErrmsg)
	pmapi.Echo.DELETE("/api/v1/queryrules/:id", pmapi.DeleteOneQueryRules)

	/*Scheduler*/
	/*
		pmapi.Echo.GET("/api/v1/scheduler", pmapi.ListAllScheduler)
		pmapi.Echo.GET("/api/v1/scheduler/:id", pmapi.ListSchedulerById)
		pmapi.Echo.POST("/api/v1/scheduler", pmapi.CreateScheduler)
		pmapi.Echo.PUT("/api/v1/scheduler/status", pmapi.UpdateOneSchedulerStatus)
		pmapi.Echo.PUT("/api/v1/scheduler/interval", pmapi.UpdateOneSchedulerInterval)
		pmapi.Echo.DELETE("/api/v1/scheduler/:id", pmapi.DeleteOneScheduler)
	*/

}

func (pmapi *PMApi) RunApiService() {
	pmapi.Echo.Logger.Fatal(pmapi.Echo.Start(":3333"))
}

func (pmapi *PMApi) DeleteOneUser(c echo.Context) error {
	user := new(users.Users)
	user.Username = c.Param("username")
	dret := user.DeleteOneUser((pmapi.Apidb))
	switch dret {
	case 0:
		return c.JSON(http.StatusOK, user)
	case 1:
		return c.JSON(http.StatusExpectationFailed, "Failed")
	case 2:
		return c.JSON(http.StatusFound, "Exists")
	default:
		return c.JSON(http.StatusOK, "Nothing")

	}

}

func (pmapi *PMApi) CreateUser(c echo.Context) error {
	args := struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}{}

	user := new(users.Users)
	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Password = args.PassWord

	fmt.Println(args)

	cret := user.AddOneUser((pmapi.Apidb))
	switch cret {
	case 0:
		return c.JSON(http.StatusCreated, user)
	case 1:
		return c.JSON(http.StatusExpectationFailed, "Failed")
	case 2:
		return c.JSON(http.StatusFound, "Exists")
	default:
		return c.JSON(http.StatusOK, "OK")
	}
}

func (pmapi *PMApi) ListOneUser(c echo.Context) error {
	user := new(users.Users)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.Username = c.Param("username")
	return c.JSON(http.StatusOK, user.FindOneUserInfo((pmapi.Apidb)))
}

func (pmapi *PMApi) ListAllUsers(c echo.Context) error {

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	return c.JSON(http.StatusOK, users.FindAllUserInfo(pmapi.Apidb, limit, skip))
}

func (pmapi *PMApi) UpdateOneUserStatus(c echo.Context) error {

	args := struct {
		UserName string `json:"username"`
		Active   uint64 `json:"active"`
	}{}

	user := new(users.Users)
	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Active = args.Active

	switch args.Active {
	case 0:
		cret := user.DisactiveOneUser(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser Failed")
		case 2:

			//return c.JSON(http.StatusExpectationFailed, "User not exists")
			return c.JSON(http.StatusExpectationFailed, args.UserName)
		default:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser ??")
		}
	case 1:
		cret := user.ActiveOneUser(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser Failed")
		case 2:
			return c.JSON(http.StatusExpectationFailed, "User not exists")
		default:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser ??")
		}

	default:
		return c.JSON(http.StatusExpectationFailed, "active?")
	}

}

func (pmapi *PMApi) UpdateOneUserDH(c echo.Context) error {

	args := struct {
		UserName         string `json:"username"`
		DefaultHostgroup uint64 `json:"default_hostgroup"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.DefaultHostgroup = args.DefaultHostgroup

	cret := user.UpdateOneUserDh(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUser Hostgroup Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDH ???")

	}

}

func (pmapi *PMApi) UpdateOneUserDS(c echo.Context) error {
	args := struct {
		UserName      string `json:"username"`
		DefaultSchema string `json:"default_schema"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.DefaultSchema = args.DefaultSchema

	cret := user.UpdateOneUserDs(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDS Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDS ???")

	}
}

func (pmapi *PMApi) UpdateOneUserMC(c echo.Context) error {
	args := struct {
		UserName       string `json:"username"`
		MaxConnections uint64 `json:"max_connections"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.MaxConnections = args.MaxConnections

	cret := user.UpdateOneUserMc(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc ???")

	}
}

func (pmapi *PMApi) UpdateOneUserPass(c echo.Context) error {
	args := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Password = args.Password

	cret := user.UpdateOneUserPass(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserPass Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc ???")

	}
}

/*返回所有后端数据库服务器的信息*/
func (pmapi *PMApi) ListAllServers(c echo.Context) error {
	return c.JSON(http.StatusOK, servers.FindAllServerInfo(pmapi.Apidb))
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

//查询出ProxySQL状态信息
func (pmapi *PMApi) ListPStatus(c echo.Context) error {
	ps := new(status.PsStatus)

	return c.JSON(http.StatusOK, ps.GetProxySqlStatus(pmapi.Apidb))
}

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c echo.Context) error {
	ps := new(variables.PsVariables)

	return c.JSON(http.StatusOK, ps.GetProxySqlVariables(pmapi.Apidb))
}

//查询出所有查询规则
func (pmapi *PMApi) ListAllQueryRules(c echo.Context) error {
	qr := new(queryrules.QueryRules)
	ret, err := qr.FindAllQr(pmapi.Apidb)
	if err != nil {
		log.Print("ListAllQueryRules->qr.FindAllQr ", err)
		return c.JSON(http.StatusExpectationFailed, "ListAllQueryRules ExpectationFailed")
	}
	return c.JSON(http.StatusOK, ret)
}

//查询出一个规则的内容
func (pmapi *PMApi) ListOneQueryRule(c echo.Context) error {
	qr := new(queryrules.QueryRules)
	if err := c.Bind(qr); err != nil {
		return err
	}
	qr.Rule_id, _ = strconv.ParseInt(c.Param("ruleid"), 10, 64)
	log.Print("ListOneQueryRule->qr.Rule_id = ", qr.Rule_id)

	ret, err := qr.FindOneQr(pmapi.Apidb)
	if err != nil {
		log.Print("ListOneQueryRules: ", err)
		return c.JSON(http.StatusExpectationFailed, "QueryRuler Exec Error")
	}

	return c.JSON(http.StatusOK, ret)
}

func (pmapi *PMApi) CreateQueryRules(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesUser(c echo.Context) error { return c.JSON(http.StatusOK, "OK") }
func (pmapi *PMApi) UpdateOneQueryRulesSchema(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesClient(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesMatchDigest(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesMatchPattern(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesReplacePattern(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesDestHostgroup(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneQueryRulesErrmsg(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) DeleteOneQueryRules(c echo.Context) error { return c.JSON(http.StatusOK, "OK") }
