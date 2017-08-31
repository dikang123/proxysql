package pmapi

import (
	"net/http"
	"proxysql-master/admin/cmd"

	"github.com/labstack/echo"
)

func (pmapi *PMApi) SetProxySQLReadonly(c echo.Context) error {
	cret := cmd.ProxyReadOnly(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLReadwrite(c echo.Context) error {
	cret := cmd.ProxyReadWrite(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStart(c echo.Context) error {
	cret := cmd.ProxyStart(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLRestart(c echo.Context) error {
	cret := cmd.ProxyRestart(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLStop(c echo.Context) error {
	cret := cmd.ProxyStop(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLPause(c echo.Context) error {
	cret := cmd.ProxyPause(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLResume(c echo.Context) error {
	cret := cmd.ProxyResume(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLShutdown(c echo.Context) error {
	cret := cmd.ProxyShutdown(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLFlogs(c echo.Context) error {
	cret := cmd.ProxyFlushLogs(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}

func (pmapi *PMApi) SetProxySQLKill(c echo.Context) error {
	cret := cmd.ProxyKill(pmapi.Apidb)
	if cret == 0 {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusExpectationFailed, "Failed")
}
