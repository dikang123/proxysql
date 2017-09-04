package pmapi

import (
	"log"
	"net/http"
	"proxysql-master/admin/queryrules"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pmapi *PMApi) CreateOneQueryRules(c *gin.Context) {

	args := struct {
		UserName string `json:"username"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	qr.Username = args.UserName
	log.Print("CreateQueryRules: ", qr)

	cret := qr.AddOneQr(pmapi.Apidb)
	if cret == 1 {
		c.JSON(http.StatusExpectationFailed, "CreateQueryRules->AddOneQr->db.Query error")
	}
	c.JSON(http.StatusOK, "OK")
}

/*Patch方法的查询规则更新函数*/
func (pmapi *PMApi) UpdateOneQueryRules(c *gin.Context) {
	args := struct {
		RuleId                int64  `json:"rule_id"`
		Active                int64  `db:"active" json:"active"`
		Username              string `db:"username" json:"username"`
		Schemaname            string `db:"schemaname" json:"schemaname"`
		FlagIN                int64  `db:"flagIN" json:"flagIN"`
		Client_addr           string `db:"client_addr" json:"client_addr"`
		Proxy_addr            string `db:"proxy_addr" json:"proxy_addr"`
		Proxy_port            int64  `db:"proxy_port" json:"proxy_port"`
		Digest                string `db:"digest" json:"digest"`
		Match_digest          string `db:"match_digest" json:"match_digest"`
		Match_pattern         string `db:"match_pattern" json:"match_pattern"`
		Negate_match_pattern  int64  `db:"negate_match_pattern" json:"negate_match_pattern"`
		FlagOUT               int64  `db:"flagOUT" json:"flagOUT"`
		Replace_pattern       string `db:"replace_pattern" json:"replace_pattern"`
		Destination_hostgroup int64  `db:"destination_hostgroup" json:"destination_hostgroup"`
		Cache_ttl             int64  `db:"cache_ttl" json:"cache_ttl"`
		Reconnect             int64  `db:"reconnect" json:"reconnect"`
		Timeout               int64  `db:"timeout" json:"timeout"`
		Retries               int64  `db:"retries" json:"retries"`
		Delay                 int64  `db:"delay" json:"delay"`
		Mirror_flagOUT        int64  `db:"mirror_flagOUT" json:"mirror_flagOUT"`
		Mirror_hostgroup      int64  `db:"mirror_hostgroup" json:"mirror_hostgroup"`
		Error_msg             string `db:"error_msg" json:"error_msg"`
		Log                   int64  `db:"log" json:"log"`
		Apply                 int64  `db:"apply" json:"apply"`
		Comment               string `db:"comment" json:"comment"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesErrmsg->c.Bind ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	qr.Rule_id = args.RuleId
	qr.Active = args.Active
	qr.Username = args.Username
	qr.Schemaname = args.Schemaname
	qr.FlagIN = args.FlagIN
	qr.Client_addr = args.Client_addr
	qr.Proxy_addr = args.Proxy_addr
	qr.Proxy_port = args.Proxy_port
	qr.Digest = args.Digest
	qr.Match_digest = args.Match_digest
	qr.Match_pattern = args.Match_pattern
	qr.Negate_match_pattern = args.Negate_match_pattern
	qr.FlagOUT = args.FlagOUT
	qr.Replace_pattern = args.Replace_pattern
	qr.Destination_hostgroup = args.Destination_hostgroup
	qr.Cache_ttl = args.Cache_ttl
	qr.Reconnect = args.Reconnect
	qr.Timeout = args.Timeout
	qr.Retries = args.Retries
	qr.Delay = args.Delay
	qr.Mirror_flagOUT = args.Mirror_flagOUT
	qr.Mirror_hostgroup = args.Mirror_hostgroup
	qr.Error_msg = args.Error_msg
	qr.Log = args.Log
	qr.Apply = args.Apply
	qr.Comment = args.Comment

	qr.UpdateOneQrInfo(pmapi.Apidb)
	c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) DeleteOneQueryRules(c *gin.Context) {
	qr := new(queryrules.QueryRules)
	qr.Rule_id, _ = strconv.ParseInt(c.Param("ruleid"), 10, 64)
	qret := qr.DeleteOneQr(pmapi.Apidb)
	if qret == 1 {
		log.Print("DeleteOneQueryRules->qr.DeleteOneQr Error")
		c.JSON(http.StatusExpectationFailed, "DeleteOneQueryRules->qr.DeleteOneQr Error")
	}
	c.JSON(http.StatusOK, "OK")
}

//查询出所有查询规则
func (pmapi *PMApi) ListAllQueryRules(c *gin.Context) {
	qr := new(queryrules.QueryRules)

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	ret, err := qr.FindAllQr(pmapi.Apidb, limit, skip)
	if err != nil {
		log.Print("ListAllQueryRules->qr.FindAllQr ", err)
		c.JSON(http.StatusExpectationFailed, "ListAllQueryRules ExpectationFailed")
	}
	c.JSON(http.StatusOK, ret)
}
