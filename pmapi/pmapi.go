package pmapi

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	Router    *gin.Engine
}

func (pmapi *PMApi) MakePMdbi() {
	pmapi.PMdbi = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", pmapi.PMuser, pmapi.PMpass, pmapi.PMhost, pmapi.PMdb)
}

func (pmapi *PMApi) RegisterServices() {

	/*初始化gin实例*/
	pmapi.Router = gin.Default()

	/*Dashboard*/
	pmapi.Router.GET("/api/v1/status", pmapi.ListPStatus)

	/*Variables*/
	pmapi.Router.GET("/api/v1/variables", pmapi.ListPsVariables)
	pmapi.Router.PUT("/api/v1/variables", pmapi.UpdateOneVariables)

	/*User Services*/
	pmapi.Router.GET("/api/v1/users", pmapi.ListAllUsers)
	pmapi.Router.POST("/api/v1/users", pmapi.CreateOneUser)
	pmapi.Router.PUT("/api/v1/users", pmapi.UpdateOneUser)
	pmapi.Router.DELETE("/api/v1/users", pmapi.DeleteOneUser)

	/*Server Services*/
	pmapi.Router.GET("/api/v1/servers", pmapi.ListAllServers)
	pmapi.Router.POST("/api/v1/servers", pmapi.CreateOneServer)
	pmapi.Router.PUT("/api/v1/servers", pmapi.UpdateOneServer)
	pmapi.Router.DELETE("/api/v1/servers", pmapi.DeleteOneServers)

	/*Query Rules*/
	pmapi.Router.GET("/api/v1/queryrules", pmapi.ListAllQueryRules)
	pmapi.Router.POST("/api/v1/queryrules", pmapi.CreateOneQueryRules)
	pmapi.Router.PUT("/api/v1/queryrules", pmapi.UpdateOneQueryRules)
	pmapi.Router.DELETE("/api/v1/queryrules", pmapi.DeleteOneQueryRules)

	/*Scheduler*/
	pmapi.Router.GET("/api/v1/schedulers", pmapi.ListAllScheduler)
	pmapi.Router.POST("/api/v1/schedulers", pmapi.CreateOneScheduler)
	pmapi.Router.PUT("/api/v1/schedulers", pmapi.UpdateOneScheduler)
	pmapi.Router.DELETE("/api/v1/schedulers", pmapi.DeleteOneScheduler)

	/*ProxySQL admin API*/
	pmapi.Router.GET("/api/v1/cmd/readonly", pmapi.SetProxySQLReadonly)
	pmapi.Router.GET("/api/v1/cmd/readwrite", pmapi.SetProxySQLReadwrite)
	pmapi.Router.GET("/api/v1/cmd/start", pmapi.SetProxySQLStart)
	pmapi.Router.GET("/api/v1/cmd/restart", pmapi.SetProxySQLRestart)
	pmapi.Router.GET("/api/v1/cmd/stop", pmapi.SetProxySQLStop)
	pmapi.Router.GET("/api/v1/cmd/pause", pmapi.SetProxySQLPause)
	pmapi.Router.GET("/api/v1/cmd/resume", pmapi.SetProxySQLResume)
	pmapi.Router.GET("/api/v1/cmd/shutdown", pmapi.SetProxySQLShutdown)
	pmapi.Router.GET("/api/v1/cmd/flushlogs", pmapi.SetProxySQLFlogs)
	pmapi.Router.GET("/api/v1/cmd/kill", pmapi.SetProxySQLKill)

}

func (pmapi *PMApi) RunApiService() {
	pmapi.Router.Run(":3333")
}
