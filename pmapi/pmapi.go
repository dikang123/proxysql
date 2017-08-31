package pmapi

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	/*Variables*/
	pmapi.Echo.GET("/api/v1/variables", pmapi.ListPsVariables)
	pmapi.Echo.PUT("/api/v1/variables", pmapi.UpdateOneVariables)

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
	pmapi.Echo.GET("/api/v1/schedulers", pmapi.ListAllScheduler)
	pmapi.Echo.POST("/api/v1/schedulers", pmapi.CreateScheduler)
	pmapi.Echo.PUT("/api/v1/schedulers", pmapi.UpdateOneScheduler)
	pmapi.Echo.DELETE("/api/v1/schedulers/:id", pmapi.DeleteOneScheduler)

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
