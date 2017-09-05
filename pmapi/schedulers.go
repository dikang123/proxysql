package pmapi

import (
	"Pdbn/src/dbusers"
	"Pdbs/src/admin/users"
	"database/sql"
	"log"
	"net/http"
	"proxysql-master/admin/schedulers"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*与调取器相关的api函数*/
func (pmapi *PMApi) ListAllScheduler(c *gin.Context) {

	var tmpsch schedulers.Schedulers
	var arysch []schedulers.Schedulers
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

		arysch, err = tmpsch.FindAllSchedulerInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, arysch)
	}
}

func (pmapi *PMApi) CreateOneScheduler(c *gin.Context) {

	var tmpsch schedulers.Schedulers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []dbusers.Users{})
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

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->AddOneScheduler->AddOneScheduler tmpsch", tmpsch)

		_, err := tmpsch.AddOneScheduler(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->AddOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) DeleteOneScheduler(c *gin.Context) {
	var tmpsch schedulers.Schedulers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []dbusers.Users{})
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

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->DeleteOneScheduler->DeleteOneScheduler tmpsch", tmpsch)

		_, err := tmpsch.DeleteOneScheduler(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->DeleteOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) UpdateOneScheduler(c *gin.Context) {
	var tmpsch schedulers.Schedulers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []dbusers.Users{})
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

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->UpdateOneScheduler->UpdateOneScheduler tmpsch", tmpsch)

		_, err := tmpsch.UpdateOneSchedulerInfo(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->UpdateOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}

}
