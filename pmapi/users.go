package pmapi

import (
	"database/sql"
	"log"
	"net/http"
	"proxysql-master/admin/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pmapi *PMApi) DeleteOneUser(c *gin.Context) {
	/*新建一个用户实例*/
	var tmpusr users.Users
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

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

		if err := c.Bind(&tmpusr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->DeleteOneUser->DeleteOneUser tmpusr", tmpusr)

		_, err := tmpusr.DeleteOneUser(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->DeleteOneUser->DeleteOneUser Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) CreateOneUser(c *gin.Context) {
	/*新建一个用户实例*/
	var tmpusr users.Users
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

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

		if err := c.Bind(&tmpusr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->CreateOneUser->AddOneUser tmpusr", tmpusr)

		_, err := tmpusr.AddOneUser(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneUser->AddOneUser Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) ListAllUsers(c *gin.Context) {

	var tmpusr users.Users
	var aryusr []users.Users

	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	if len(hostname) == 0 || hostname == "undefined" {
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

		aryusr, err = tmpusr.FindAllUserInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, aryusr)
	}

}

/*更新用户信息的patch方法*/
func (pmapi *PMApi) UpdateOneUser(c *gin.Context) {

	/*新建一个用户实例*/
	var tmpusr users.Users
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

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

		if err := c.Bind(&tmpusr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->UpdateOneUser->UpdateOneUser tmpusr", tmpusr)

		_, err := tmpusr.UpdateOneUserInfo(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->UpdateOneUser->UpdateOneUser Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}
