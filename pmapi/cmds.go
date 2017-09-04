package pmapi

import (
	"net/http"
	"proxysql-master/admin/cmd"

	"github.com/gin-gonic/gin"
)

func (pmapi *PMApi) SetProxySQLReadonly(c *gin.Context) {
	cret := cmd.ProxyReadOnly(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, gin.H{"result": "OK"})
	}
	c.JSON(http.StatusExpectationFailed, gin.H{"result": "Failed"})
}

func (pmapi *PMApi) SetProxySQLReadwrite(c *gin.Context) {
	cret := cmd.ProxyReadWrite(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStart(c *gin.Context) {
	cret := cmd.ProxyStart(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLRestart(c *gin.Context) {
	cret := cmd.ProxyRestart(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStop(c *gin.Context) {
	cret := cmd.ProxyStop(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLPause(c *gin.Context) {
	cret := cmd.ProxyPause(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLResume(c *gin.Context) {
	cret := cmd.ProxyResume(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLShutdown(c *gin.Context) {
	cret := cmd.ProxyShutdown(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLFlogs(c *gin.Context) {
	cret := cmd.ProxyFlushLogs(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLKill(c *gin.Context) {
	cret := cmd.ProxyKill(pmapi.Apidb)
	if cret == 0 {
		c.JSON(http.StatusOK, "OK")
	}
	c.JSON(http.StatusExpectationFailed, "Failed")
}
