package pmapi

import (
	"fmt"
	"log"
	"net/http"
	"proxysql-master/admin/users"
	"strconv"

	"github.com/labstack/echo"
)

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
		return c.JSON(http.StatusOK, "Nothing")

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

	fmt.Println(args)

	cret := user.AddOneUser((pmapi.Apidb))
	switch cret {
	case 0:
		return c.JSON(http.StatusCreated, user)
	case 1:
		return c.JSON(http.StatusExpectationFailed, "Failed")
	case 2:
		return c.JSON(http.StatusFound, "Exists")
	default:
		return c.JSON(http.StatusOK, "OK")
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

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	return c.JSON(http.StatusOK, users.FindAllUserInfo(pmapi.Apidb, limit, skip))
}

func (pmapi *PMApi) UpdateOneUserStatus(c echo.Context) error {

	args := struct {
		UserName string `json:"username"`
		Active   uint64 `json:"active"`
	}{}

	user := new(users.Users)
	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Active = args.Active

	switch args.Active {
	case 0:
		cret := user.DisactiveOneUser(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser Failed")
		case 2:

			//return c.JSON(http.StatusExpectationFailed, "User not exists")
			return c.JSON(http.StatusExpectationFailed, args.UserName)
		default:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser ??")
		}
	case 1:
		cret := user.ActiveOneUser(pmapi.Apidb)
		switch cret {
		case 0:
			return c.JSON(http.StatusOK, "OK")
		case 1:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser Failed")
		case 2:
			return c.JSON(http.StatusExpectationFailed, "User not exists")
		default:
			return c.JSON(http.StatusExpectationFailed, "DisactiveOneUser ??")
		}

	default:
		return c.JSON(http.StatusExpectationFailed, "active?")
	}

}

func (pmapi *PMApi) UpdateOneUserDH(c echo.Context) error {

	args := struct {
		UserName         string `json:"username"`
		DefaultHostgroup uint64 `json:"default_hostgroup"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.DefaultHostgroup = args.DefaultHostgroup

	cret := user.UpdateOneUserDh(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUser Hostgroup Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDH ???")

	}

}

func (pmapi *PMApi) UpdateOneUserDS(c echo.Context) error {
	args := struct {
		UserName      string `json:"username"`
		DefaultSchema string `json:"default_schema"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.DefaultSchema = args.DefaultSchema

	cret := user.UpdateOneUserDs(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDS Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserDS ???")

	}
}

func (pmapi *PMApi) UpdateOneUserMC(c echo.Context) error {
	args := struct {
		UserName       string `json:"username"`
		MaxConnections uint64 `json:"max_connections"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.MaxConnections = args.MaxConnections

	cret := user.UpdateOneUserMc(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc ???")

	}
}

func (pmapi *PMApi) UpdateOneUserPass(c echo.Context) error {
	args := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	user := new(users.Users)

	if err := c.Bind(&args); err != nil {
		return err
	}

	user.Username = args.UserName
	user.Password = args.Password

	cret := user.UpdateOneUserPass(pmapi.Apidb)
	switch cret {
	case 0:
		return c.JSON(http.StatusOK, "OK")
	case 1:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserPass Failed")
	case 2:
		return c.JSON(http.StatusExpectationFailed, "User not exists")
	default:
		return c.JSON(http.StatusExpectationFailed, "UpdateOneUserMc ???")

	}
}

/*更新用户信息的patch方法*/
func (pmapi *PMApi) UpdateOneUserInfo(c echo.Context) error {

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
		return err
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
	return c.JSON(http.StatusOK, "OK")
}
