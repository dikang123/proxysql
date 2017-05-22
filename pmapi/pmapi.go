package pmapi

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"proxysql-master/admin/users"
)

type PMApi struct {
	PMuser  string
	PMpass  string
	PMhost  string
	PMdb    string
	PMdbi   string
	Apidb   *sql.DB
	ApiHost string
	*echo.Echo
}

func (pmapi *PMApi) MakePMdbi() {
	pmapi.PMdbi = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", pmapi.PMuser, pmapi.PMpass, pmapi.PMhost, pmapi.PMdb)
}

func (pmapi *PMApi) RegisterDBInterface() {
	var err error
	pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
	if err != nil {
		log.Fatal("sql.Open()", err)
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
	/*User Services*/
	pmapi.Echo.GET("/users", pmapi.ListAllUsers)
	pmapi.Echo.GET("/users/:username", pmapi.ListOneUser)
	pmapi.Echo.POST("/users", pmapi.CreateUser)
	//pmapi.Echo.PUT("/users/:username", pmapi.updateUsers)
	pmapi.Echo.DELETE("/users/:username", pmapi.DeleteOneUser)
	/*Server Services*/

	/*Query Rules*/

	/*Scheduler*/
}

func (pmapi *PMApi) RunApiService() {
	pmapi.Echo.Logger.Fatal(pmapi.Echo.Start(":3333"))
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
		return c.String(http.StatusOK, "Nothing")

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

	cret := user.AddOneUser((pmapi.Apidb))
	switch cret {
	case 0:
		return c.JSON(http.StatusCreated, user)
	case 1:
		return c.String(http.StatusExpectationFailed, "Failed")
	case 2:
		return c.String(http.StatusFound, "Exists")
	default:
		return c.String(http.StatusOK, "OK")
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
	return c.JSON(http.StatusOK, users.FindAllUserInfo((pmapi.Apidb)))
}
