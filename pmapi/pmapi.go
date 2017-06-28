package pmapi

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"proxysql-master/admin/cmd"
	"proxysql-master/admin/queryrules"
	"proxysql-master/admin/servers"
	"proxysql-master/admin/status"
	"proxysql-master/admin/users"
	"proxysql-master/admin/variables"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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
	pmapi.Echo.PATCH("/api/v1/users", pmapi.UpdateOneUserInfo)
	pmapi.Echo.DELETE("/api/v1/users/:username", pmapi.DeleteOneUser)

	/*Server Services*/
	pmapi.Echo.GET("/api/v1/servers", pmapi.ListAllServers)
	pmapi.Echo.GET("/api/v1/servers/:hostgroup", pmapi.ListServerByHostgroup)
	pmapi.Echo.PUT("/api/v1/servers", pmapi.ListOneServer)
	pmapi.Echo.POST("/api/v1/servers", pmapi.CreateServer)
	pmapi.Echo.PUT("/api/v1/servers/status", pmapi.UpdateOneServerStatus)
	pmapi.Echo.PUT("/api/v1/servers/weight", pmapi.UpdateOneServerWeight)
	pmapi.Echo.PUT("/api/v1/servers/maxconnection", pmapi.UpdateOneServerMC)
	pmapi.Echo.PATCH("/api/v1/servers", pmapi.UpdateOneServerInfo)
	pmapi.Echo.DELETE("/api/v1/servers", pmapi.DeleteOneServers)

	/*Query Rules*/
	pmapi.Echo.GET("/api/v1/queryrules", pmapi.ListAllQueryRules)
	pmapi.Echo.GET("/api/v1/queryrules/:ruleid", pmapi.ListOneQueryRule)
	pmapi.Echo.POST("/api/v1/queryrules", pmapi.CreateQueryRules)
	pmapi.Echo.PUT("/api/v1/queryrules/status", pmapi.UpdateOneQueryRulesStatus)
	pmapi.Echo.PUT("/api/v1/queryrules/username", pmapi.UpdateOneQueryRulesUser)
	pmapi.Echo.PUT("/api/v1/queryrules/schemaname", pmapi.UpdateOneQueryRulesSchema)
	pmapi.Echo.PUT("/api/v1/queryrules/clientaddr", pmapi.UpdateOneQueryRulesClient)
	pmapi.Echo.PUT("/api/v1/queryrules/matchdigest", pmapi.UpdateOneQueryRulesMatchDigest)
	pmapi.Echo.PUT("/api/v1/queryrules/digest", pmapi.UpdateOneQueryRulesDigest)
	pmapi.Echo.PUT("/api/v1/queryrules/matchpattern", pmapi.UpdateOneQueryRulesMatchPattern)
	pmapi.Echo.PUT("/api/v1/queryrules/replacepattern", pmapi.UpdateOneQueryRulesReplacePattern)
	pmapi.Echo.PUT("/api/v1/queryrules/desthostgroup", pmapi.UpdateOneQueryRulesDestHostgroup)
	pmapi.Echo.PUT("/api/v1/queryrules/errmsg", pmapi.UpdateOneQueryRulesErrmsg)
	pmapi.Echo.PATCH("/api/v1/queryrules", pmapi.UpdateOneQueryRulesInfo)
	pmapi.Echo.DELETE("/api/v1/queryrules/:ruleid", pmapi.DeleteOneQueryRules)

	/*Scheduler*/
	/*
		pmapi.Echo.GET("/api/v1/scheduler", pmapi.ListAllScheduler)
		pmapi.Echo.GET("/api/v1/scheduler/:id", pmapi.ListSchedulerById)
		pmapi.Echo.POST("/api/v1/scheduler", pmapi.CreateScheduler)
		pmapi.Echo.PUT("/api/v1/scheduler/status", pmapi.UpdateOneSchedulerStatus)
		pmapi.Echo.PUT("/api/v1/scheduler/interval", pmapi.UpdateOneSchedulerInterval)
		pmapi.Echo.DELETE("/api/v1/scheduler/:id", pmapi.DeleteOneScheduler)
	*/

	/*ProxySQL admin API*/
	pmapi.Echo.GET("/api/v1/cmd/readonly", pmapi.SetProxySQLReadonly)
	pmapi.Echo.GET("/api/v1/cmd/readwrite", pmapi.SetProxySQLReadwrite)
	pmapi.Echo.GET("/api/v1/cmd/start", pmapi.SetProxySQLStart)
	pmapi.Echo.GET("/api/v1/cmd/restart", pmapi.SetProxySQLRestart)
	pmapi.Echo.GET("/api/v1/cmd/stop", pmapi.SetProxySQLStop)
	pmapi.Echo.GET("/api/v1/cmd/pause", pmapi.SetProxySQLPause)
	pmapi.Echo.GET("/api/v1/cmd/resume", pmapi.SetProxySQLResume)
	pmapi.Echo.GET("/api/v1/cmd/shutdown", pmapi.SetProxySQLShutdown)
	pmapi.Echo.GET("/api/v1/cmd/flushlogs", pmapi.SetProxySQLFlogs)
	pmapi.Echo.GET("/api/v1/cmd/kill", pmapi.SetProxySQLKill)

}

func (pmapi *PMApi) RunApiService() {
	pmapi.Echo.Logger.Fatal(pmapi.Echo.Start(pmapi.ApiHost))
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

/*更新用户信息的patch方法*/
func (pmapi *PMApi) UpdateOneUserInfo(c echo.Context) error {

	args := struct {
		UserName              string `json:"username"`
		Password              string `json:"password"`
		Active                uint64 `json:"active"`
		UseSsl                uint64 `json:"use_ssl"`
		DefaultHostgroup      uint64 `json:"default_hostgroup"`
		DefaultSchema         string `json:"default_schema"`
		SchemaLocked          uint64 `json:"schema_locked"`
		TransactionPersistent uint64 `json:"transaction_persistent"`
		FastForward           uint64 `json:"fast_forward"`
		Backend               uint64 `json:"backend"`
		Frontend              uint64 `json:"frontend"`
		MaxConnections        uint64 `json:"max_connections"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Password = args.Password
	user.Active = args.Active
	user.UseSsl = args.UseSsl
	user.DefaultHostgroup = args.DefaultHostgroup
	user.DefaultSchema = args.DefaultSchema
	user.SchemaLocked = args.SchemaLocked
	user.TransactionPersistent = args.TransactionPersistent
	user.FastForward = args.FastForward
	user.Backend = args.Backend
	user.Frontend = args.Frontend
	user.MaxConnections = args.MaxConnections

	user.UpdateOneUserInfo(pmapi.Apidb)
	return c.JSON(http.StatusOK, "OK")
}

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

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	ret, err := qr.FindAllQr(pmapi.Apidb, limit, skip)
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

	args := struct {
		UserName string `json:"username"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		return err
	}

	qr.Username = args.UserName
	log.Print("CreateQueryRules: ", qr)

	cret := qr.AddOneQr(pmapi.Apidb)
	if cret == 1 {
		return c.JSON(http.StatusExpectationFailed, "CreateQueryRules->AddOneQr->db.Query error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) UpdateOneQueryRulesStatus(c echo.Context) error {
	args := struct {
		RuleId int64 `json:"rule_id"`
		Status int64 `json:"active"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesStatus->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	log.Print("UpdateOneQueryRulesStatus->args.Status:", args.Status)

	if args.Status == 0 {
		qret := qr.DisactiveOneQr(pmapi.Apidb)
		if qret == 1 {
			return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesStatus->DisactiveOneQr error")
		}
		return c.JSON(http.StatusOK, "Disactive OK")
	}
	qret := qr.ActiveOneQr(pmapi.Apidb)
	if qret == 1 {
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesStatus->ActiveOneQr error")
	}
	return c.JSON(http.StatusOK, "Active OK")
}

//更新一个查询规则中的用户名
func (pmapi *PMApi) UpdateOneQueryRulesUser(c echo.Context) error {
	args := struct {
		RuleId   int64  `json:"rule_id"`
		Username string `json:"username"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesUser->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Username = args.Username

	qret := qr.UpdateOneQrUn(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesUser->qr.UpdateOneQrUn err")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesUser Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) UpdateOneQueryRulesSchema(c echo.Context) error {
	args := struct {
		RuleId     int64  `json:"rule_id"`
		Schemaname string `json:"schemaname"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesSchea->c.Bind", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Schemaname = args.Schemaname

	qret := qr.UpdateOneQrSn(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRules->UpdateOneQrSn Err")
		return c.JSON(http.StatusExpectationFailed, "UPdateOneQueryRulesSchema Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新一个规则的客户端地址
func (pmapi *PMApi) UpdateOneQueryRulesClient(c echo.Context) error {
	args := struct {
		RuleId     int64  `json:"rule_id"`
		ClientAddr string `json:"client_addr"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesClient->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Client_addr = args.ClientAddr

	qret := qr.UpdateOneQrCa(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesClient->UpdateOneQrCa Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesClient->qr.UpdateOneQrCa  Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新查询规则的digest列
func (pmapi *PMApi) UpdateOneQueryRulesDigest(c echo.Context) error {
	args := struct {
		RuleId int64  `json:"rule_id"`
		Digest string `json:"digest"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesDigest->c.Bind :", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Digest = args.Digest

	qret := qr.UpdateOneQrDg(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesDigest->qr.UpdateOneQrDg Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesDigest->qr.UpdateOneQrDg Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新查询规则的match_digest列
func (pmapi *PMApi) UpdateOneQueryRulesMatchDigest(c echo.Context) error {
	args := struct {
		RuleId      int64  `json:"rule_id"`
		MatchDigest string `json:"match_digest"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesMatchDigest->c.Bind :", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Match_digest = args.MatchDigest

	qret := qr.UpdateOneQrMd(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesMatchDigest->qr.UpdateOneQrMd Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesMatchDigest->qr.UpdateOneQrMd Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新规则匹配内容
func (pmapi *PMApi) UpdateOneQueryRulesMatchPattern(c echo.Context) error {
	args := struct {
		RuleId       int64  `json:"rule_id"`
		MatchPattern string `json:"match_pattern"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesMatchPattern->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Match_pattern = args.MatchPattern

	qret := qr.UpdateOneQrMp(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesMatchPattern->qr.UpdateOneQrMp Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesMatchPattern->qr.UpdateOneQrMp Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新替换内容
func (pmapi *PMApi) UpdateOneQueryRulesReplacePattern(c echo.Context) error {
	args := struct {
		RuleId         int64  `json:"rule_id"`
		ReplacePattern string `json:"replace_pattern"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesReplacePattern->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Replace_pattern = args.ReplacePattern

	qret := qr.UpdateOneQrRp(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesReplacePattern->qr.UpdateOneQrRp Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesReplacePattern->qr.UpdateOneQrRp Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新规则的默认主机组
func (pmapi *PMApi) UpdateOneQueryRulesDestHostgroup(c echo.Context) error {
	args := struct {
		RuleId               int64 `json:"rule_id"`
		DestinationHostgroup int64 `json:"destination_hostgroup"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesDestHostgroup->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Destination_hostgroup = args.DestinationHostgroup

	qret := qr.UpdateOneQrDh(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesDestHostgroup->qr.UpdateOneQrDh Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesDestHostgroup->qr.UpdateOneQrDh  Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新一个规则的错误消息内容
func (pmapi *PMApi) UpdateOneQueryRulesErrmsg(c echo.Context) error {
	args := struct {
		RuleId int64  `json:"rule_id"`
		ErrMsg string `json:"error_msg"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesErrmsg->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Error_msg = args.ErrMsg

	qret := qr.UpdateOneQrEm(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesErrmsg->qr.UpdateOneQrEm Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesErrmsg->qr.UpdateOneQrEm Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) DeleteOneQueryRules(c echo.Context) error {
	qr := new(queryrules.QueryRules)
	qr.Rule_id, _ = strconv.ParseInt(c.Param("ruleid"), 10, 64)
	qret := qr.DeleteOneQr(pmapi.Apidb)
	if qret == 1 {
		log.Print("DeleteOneQueryRules->qr.DeleteOneQr Error")
		return c.JSON(http.StatusExpectationFailed, "DeleteOneQueryRules->qr.DeleteOneQr Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) SetProxySQLReadonly(c echo.Context) error {
	cret := cmd.ProxyReadOnly(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLReadwrite(c echo.Context) error {
	cret := cmd.ProxyReadWrite(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStart(c echo.Context) error {
	cret := cmd.ProxyStart(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLRestart(c echo.Context) error {
	cret := cmd.ProxyRestart(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStop(c echo.Context) error {
	cret := cmd.ProxyStop(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLPause(c echo.Context) error {
	cret := cmd.ProxyPause(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLResume(c echo.Context) error {
	cret := cmd.ProxyResume(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLShutdown(c echo.Context) error {
	cret := cmd.ProxyShutdown(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLFlogs(c echo.Context) error {
	cret := cmd.ProxyFlushLogs(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLKill(c echo.Context) error {
	cret := cmd.ProxyKill(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}
