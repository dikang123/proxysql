package pmapi

import (
	"Pdbn/src/dbusers"
	"database/sql"
	"log"
	"net/http"
	"proxysql-master/admin/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pmapi *PMApi) DeleteOneUser(c *gin.Context) {
	user := new(users.Users)
	user.Username = c.Param("username")
	dret := user.DeleteOneUser((pmapi.Apidb))
	switch dret {
	case 0:
		c.JSON(http.StatusOK, user)
	case 1:
		c.JSON(http.StatusExpectationFailed, "Failed")
	case 2:
		c.JSON(http.StatusFound, "Exists")
	default:
		c.JSON(http.StatusOK, "Nothing")

	}

}

func (pmapi *PMApi) CreateOneUser(c *gin.Context) {
	/*新建一个用户实例*/
	var tmpusr users.Users
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

		if err := c.Bind(&tmpusr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->CreateOneUser->AddOneUser tmpusr", tmpusr)

		_, err := tmpusr.AddOneUser(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneUser->AddOneUser Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
	}
}

func (pmapi *PMApi) ListAllUsers(c *gin.Context) {

	var tmpusr users.Users
	var aryusr []users.Users

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

		aryusr, err = tmpusr.FindAllUserInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, aryusr)
	}

}

/*更新用户信息的patch方法*/
func (pmapi *PMApi) UpdateOneUser(c *gin.Context) {

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
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
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

	log.Print("pmapi->UpdateOneUserInfo->user :", user)

	user.UpdateOneUserInfo(pmapi.Apidb)
	c.JSON(http.StatusOK, "OK")
}
