package pmapi

import (
	"net/http"

	"github.com/imSQL/proxysql-master/admin/status"

	"github.com/gin-gonic/gin"
)

//查询出ProxySQL状态信息
func (pmapi *PMApi) ListPStatus(c *gin.Context) {
	ps := new(status.PsStatus)

	c.JSON(http.StatusOK, ps.GetProxySqlStatus(pmapi.Apidb))
}
