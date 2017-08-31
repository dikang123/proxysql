package pmapi

import (
	"net/http"
	"proxysql-master/admin/variables"

	"github.com/labstack/echo"
)

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c echo.Context) error {

	return c.JSON(http.StatusOK, variables.GetProxySqlVariables(pmapi.Apidb))
}
