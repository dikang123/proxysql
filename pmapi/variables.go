package pmapi

import (
	"Pdbs/src/admin/users"
	"database/sql"
	"net/http"
	"proxysql-master/admin/variables"

	"github.com/gin-gonic/gin"
)

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c *gin.Context) {
	var tmpvars variables.Variables
	var aryvars []variables.Variables
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")
	/*
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		if limit == 0 {
			limit = 10
		}

		if page == 0 {
			page = 1
		}

		skip := (page - 1) * limit
	*/

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

		aryvars, err = tmpvars.GetProxySqlVariables(pmapi.Apidb)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, aryvars)
	}

}

func (pmapi *PMApi) UpdateOneVariables(c *gin.Context) {

	var tmpvars variables.Variables
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")

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

		if err := c.Bind(&tmpvars); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}

		_, err := tmpvars.UpdateOneVariable(pmapi.Apidb)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": "UpdateOneVariable Failed"})
		}
		c.JSON(http.StatusOK, gin.H{"result": "OK"})
	}

}
