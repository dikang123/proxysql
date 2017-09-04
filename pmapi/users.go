package pmapi

import (
	"fmt"
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
	args := struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}{}

	user := new(users.Users)
	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	user.Username = args.UserName
	user.Password = args.PassWord

	fmt.Println(args)

	cret := user.AddOneUser((pmapi.Apidb))
	switch cret {
	case 0:
		c.JSON(http.StatusCreated, user)
	case 1:
		c.JSON(http.StatusExpectationFailed, "Failed")
	case 2:
		c.JSON(http.StatusFound, "Exists")
	default:
		c.JSON(http.StatusOK, "OK")
	}
}

func (pmapi *PMApi) ListAllUsers(c *gin.Context) {

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	c.JSON(http.StatusOK, users.FindAllUserInfo(pmapi.Apidb, limit, skip))
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
