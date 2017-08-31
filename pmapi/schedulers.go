package pmapi

import (
	"log"
	"net/http"
	"proxysql-master/admin/schedulers"
	"strconv"

	"github.com/labstack/echo"
)

/*与调取器相关的api函数*/
func (pmapi *PMApi) ListAllScheduler(c echo.Context) error {

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	ret, err := schedulers.FindAllSchedulerInfo(pmapi.Apidb, limit, skip)
	if err != nil {
		log.Print("ListAllScheduler->qr.FindAllSchdulerInfo ", err)
		return c.JSON(http.StatusExpectationFailed, "ListAllSchduler ExpectationFailed")
	}
	log.Print("ret=", ret)
	return c.JSON(http.StatusOK, ret)
}

func (pmapi *PMApi) CreateScheduler(c echo.Context) error {
	args := struct {
		FileName   string `json:"filename"`
		IntervalMs int64  `json:"interval_ms"`
	}{}

	schld := new(schedulers.Schedulers)

	if err := c.Bind(&args); err != nil {
		return err
	}

	schld.FileName = args.FileName
	schld.IntervalMs = args.IntervalMs

	sret := schld.AddOneScheduler(pmapi.Apidb)
	if sret == 1 {
		log.Print("pmapi.go->CreateScheduler->AddOneScheduler Failed")
		return c.JSON(http.StatusExpectationFailed, "CreateScheduler Failed")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) DeleteOneScheduler(c echo.Context) error {
	schld := new(schedulers.Schedulers)
	schld.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

	sret := schld.DeleteOneScheduler(pmapi.Apidb)
	if sret != 0 {
		return c.JSON(http.StatusExpectationFailed, "DeleteOneScheduler Failed")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) UpdateOneScheduler(c echo.Context) error {
	args := struct {
		Id         int64  `json:"id" db:"id"`
		Active     int64  `json:"active" db:"active"`
		IntervalMs int64  `json:"interval_ms" db:"interval_ms"`
		FileName   string `json:"filename" db:"filename"`
		Arg1       string `json:"arg1" db:"arg1"`
		Arg2       string `json:"arg2" db:"arg2"`
		Arg3       string `json:"arg3" db:"arg3"`
		Arg4       string `json:"arg4" db:"arg4"`
		Arg5       string `json:"arg5" db:"arg5"`
		Comment    string `json:"comment" db:"comment"`
	}{}

	schld := new(schedulers.Schedulers)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneScheduler->c.Bind ", err)
		return err
	}

	schld.Id = args.Id
	schld.Active = args.Active
	schld.IntervalMs = args.IntervalMs
	schld.FileName = args.FileName
	schld.Arg1 = args.Arg1
	schld.Arg2 = args.Arg2
	schld.Arg3 = args.Arg3
	schld.Arg4 = args.Arg4
	schld.Arg5 = args.Arg5
	schld.Comment = args.Comment

	log.Print("pmapi->UpdateOneScheduler->schld: ", schld)

	sret := schld.UpdateOneSchedulerInfo(pmapi.Apidb)
	if sret != 0 {
		return c.JSON(http.StatusExpectationFailed, "UpdateOneScheduler Failed")
	}
	return c.JSON(http.StatusOK, "OK")

}
