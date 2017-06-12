package main

import (
	"log"
	//"database/sql"
	//"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
	//"log"
	//	"net/http"
	"proxysql-master/pmapi"
	//	"proxysql-master/admin/servers"
	//"proxysql-master/admin/cmd"
	//	"proxysql-master/admin/users"
	//"os"
	//"proxysql-master/admin/status"
	"os"
	"syscall"
)

func main() {
	//var err error
	pmapiv1 := new(pmapi.PMApi)

	pmapiv1.ApiLogcwd = "/tmp/pm.log"
	pmapiv1.ApiLogfd, pmapiv1.ApiErr = os.OpenFile(pmapiv1.ApiLogcwd, syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0755)
	if pmapiv1.ApiErr != nil {
		log.Fatal("Open Log File Failed", pmapiv1.ApiErr)
	}
	log.SetOutput(pmapiv1.ApiLogfd)
	log.Printf("%s", "test")

	pmapiv1.Echo = echo.New()
	//e := pmapiv1.Echo
	pmapiv1.PMuser = "admin"
	pmapiv1.PMpass = "admin"
	pmapiv1.PMhost = "172.18.7.204:6032"
	pmapiv1.PMdb = "main"
	pmapiv1.MakePMdbi()

	pmapiv1.RegisterMiddleware()

	pmapiv1.RegisterDBInterface()

	pmapiv1.RegisterServices()

	pmapiv1.RunApiService()

}
