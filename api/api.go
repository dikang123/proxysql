package api

import (
	"fmt"
)

func deleteOneUser(c echo.Context) error {
	user := new(users.Users)
	user.Username = c.Param("username")
	fmt.Println(user.Username)
	dret := user.DeleteOneUser(db)
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

func createUser(c echo.Context) error {
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

	cret := user.AddOneUser(db)
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

func listOneUser(c echo.Context) error {
	user := new(users.Users)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.Username = c.Param("username")
	return c.JSON(http.StatusOK, user.FindOneUserInfo(db))
}

func listAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users.FindAllUserInfo(db))
}
