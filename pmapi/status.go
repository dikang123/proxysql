package pmapi

import (
	"net/http"
	"proxysql-master/admin/status"

	"github.com/labstack/echo"
)

//查询出ProxySQL状态信息
func (pmapi *PMApi) ListPStatus(c echo.Context) error {
	ps := new(status.PsStatus)

	return c.JSON(http.StatusOK, ps.GetProxySqlStatus(pmapi.Apidb))
}
