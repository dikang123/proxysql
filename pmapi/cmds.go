package pmapi

import (
	"net/http"
	"proxysql-master/admin/cmd"

	"github.com/gin-gonic/gin"
)

func (pmapi *PMApi) SetProxySQLReadonly(c *gin.Context) {
	_, err := cmd.ProxyReadOnly(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLReadwrite(c *gin.Context) {
	_, err := cmd.ProxyReadWrite(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStart(c *gin.Context) {
	_, err := cmd.ProxyStart(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLRestart(c *gin.Context) {
	_, err := cmd.ProxyRestart(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStop(c *gin.Context) {
	_, err := cmd.ProxyStop(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLPause(c *gin.Context) {
	_, err := cmd.ProxyPause(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLResume(c *gin.Context) {
	_, err := cmd.ProxyResume(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLShutdown(c *gin.Context) {
	_, err := cmd.ProxyShutdown(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLFlogs(c *gin.Context) {
	_, err := cmd.ProxyFlushLogs(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLKill(c *gin.Context) {
	_, err := cmd.ProxyKill(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
