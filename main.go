package main

import (
	"fmt"
	"log"
	"github.com/labstack/echo"
	"proxysql-master/pmapi"
	"flag"
	"os"
	"strings"
	"syscall"
)

var (
	apiSource = flag.String("s", "admin/admin@localhost:6032/main", "ProxySQL Connection URI address.")
	apiPort   = flag.Int64("p", 6031, "Api port.")
	apiLog    = flag.String("l", "/tmp/pm.log", "api log file.")
)

func main() {

	flag.Parse()
	if len(os.Args) <= 2 {
		log.Fatal(`
	    Usage of proxysql-master:
	      -l string
	        api log file.  (default "/tmp/pm.log")
	      -p int
	        Api port.  (default 6031)
	      -s string
         	ProxySQL Connection URI address. (default "admin/admin@localhost:6032/main")
	    `)
	}

	//新建Api实例
	pmapiv1 := new(pmapi.PMApi)
	//设定api的运行主机和端口
	pmapiv1.ApiHost = fmt.Sprintf(":%d", *apiPort)

	//设定api的日志路径
	pmapiv1.ApiLogcwd = *apiLog
	pmapiv1.ApiLogfd, pmapiv1.ApiErr = os.OpenFile(pmapiv1.ApiLogcwd, syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0644)
	if pmapiv1.ApiErr != nil {
		log.Fatal("Open Log File Failed", pmapiv1.ApiErr)
	}
	log.SetOutput(pmapiv1.ApiLogfd)

	pmapiv1.Echo = echo.New()
	pmapiv1.PMuser = strings.Split(strings.Split(*apiSource, "@")[0], "/")[0]
	pmapiv1.PMpass = strings.Split(strings.Split(*apiSource, "@")[0], "/")[1]
	pmapiv1.PMhost = strings.Split(strings.Split(*apiSource, "@")[1], "/")[0]
	pmapiv1.PMdb = strings.Split(strings.Split(*apiSource, "@")[1], "/")[1]

	pmapiv1.MakePMdbi()

	pmapiv1.RegisterMiddleware()

	pmapiv1.RegisterDBInterface()

	pmapiv1.RegisterServices()

	pmapiv1.RunApiService()

}
