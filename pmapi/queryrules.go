package pmapi

import (
	"log"
	"net/http"
	"proxysql-master/admin/queryrules"
	"proxysql-master/admin/variables"
	"strconv"

	"github.com/labstack/echo"
)

func (pmapi *PMApi) CreateQueryRules(c echo.Context) error {

	args := struct {
		UserName string `json:"username"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		return err
	}

	qr.Username = args.UserName
	log.Print("CreateQueryRules: ", qr)

	cret := qr.AddOneQr(pmapi.Apidb)
	if cret == 1 {
		return c.JSON(http.StatusExpectationFailed, "CreateQueryRules->AddOneQr->db.Query error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) UpdateOneQueryRulesStatus(c echo.Context) error {
	args := struct {
		RuleId int64 `json:"rule_id"`
		Status int64 `json:"active"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesStatus->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	log.Print("UpdateOneQueryRulesStatus->args.Status:", args.Status)

	if args.Status == 0 {
		qret := qr.DisactiveOneQr(pmapi.Apidb)
		if qret == 1 {
			return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesStatus->DisactiveOneQr error")
		}
		return c.JSON(http.StatusOK, "Disactive OK")
	}
	qret := qr.ActiveOneQr(pmapi.Apidb)
	if qret == 1 {
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesStatus->ActiveOneQr error")
	}
	return c.JSON(http.StatusOK, "Active OK")
}

//更新一个查询规则中的用户名
func (pmapi *PMApi) UpdateOneQueryRulesUser(c echo.Context) error {
	args := struct {
		RuleId   int64  `json:"rule_id"`
		Username string `json:"username"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesUser->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Username = args.Username

	qret := qr.UpdateOneQrUn(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesUser->qr.UpdateOneQrUn err")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesUser Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) UpdateOneQueryRulesSchema(c echo.Context) error {
	args := struct {
		RuleId     int64  `json:"rule_id"`
		Schemaname string `json:"schemaname"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesSchea->c.Bind", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Schemaname = args.Schemaname

	qret := qr.UpdateOneQrSn(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRules->UpdateOneQrSn Err")
		return c.JSON(http.StatusExpectationFailed, "UPdateOneQueryRulesSchema Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新一个规则的客户端地址
func (pmapi *PMApi) UpdateOneQueryRulesClient(c echo.Context) error {
	args := struct {
		RuleId     int64  `json:"rule_id"`
		ClientAddr string `json:"client_addr"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesClient->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Client_addr = args.ClientAddr

	qret := qr.UpdateOneQrCa(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesClient->UpdateOneQrCa Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesClient->qr.UpdateOneQrCa  Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新查询规则的digest列
func (pmapi *PMApi) UpdateOneQueryRulesDigest(c echo.Context) error {
	args := struct {
		RuleId int64  `json:"rule_id"`
		Digest string `json:"digest"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesDigest->c.Bind :", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Digest = args.Digest

	qret := qr.UpdateOneQrDg(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesDigest->qr.UpdateOneQrDg Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesDigest->qr.UpdateOneQrDg Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新查询规则的match_digest列
func (pmapi *PMApi) UpdateOneQueryRulesMatchDigest(c echo.Context) error {
	args := struct {
		RuleId      int64  `json:"rule_id"`
		MatchDigest string `json:"match_digest"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesMatchDigest->c.Bind :", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Match_digest = args.MatchDigest

	qret := qr.UpdateOneQrMd(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesMatchDigest->qr.UpdateOneQrMd Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesMatchDigest->qr.UpdateOneQrMd Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新规则匹配内容
func (pmapi *PMApi) UpdateOneQueryRulesMatchPattern(c echo.Context) error {
	args := struct {
		RuleId       int64  `json:"rule_id"`
		MatchPattern string `json:"match_pattern"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesMatchPattern->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Match_pattern = args.MatchPattern

	qret := qr.UpdateOneQrMp(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesMatchPattern->qr.UpdateOneQrMp Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesMatchPattern->qr.UpdateOneQrMp Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新替换内容
func (pmapi *PMApi) UpdateOneQueryRulesReplacePattern(c echo.Context) error {
	args := struct {
		RuleId         int64  `json:"rule_id"`
		ReplacePattern string `json:"replace_pattern"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesReplacePattern->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Replace_pattern = args.ReplacePattern

	qret := qr.UpdateOneQrRp(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesReplacePattern->qr.UpdateOneQrRp Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesReplacePattern->qr.UpdateOneQrRp Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新规则的默认主机组
func (pmapi *PMApi) UpdateOneQueryRulesDestHostgroup(c echo.Context) error {
	args := struct {
		RuleId               int64 `json:"rule_id"`
		DestinationHostgroup int64 `json:"destination_hostgroup"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesDestHostgroup->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Destination_hostgroup = args.DestinationHostgroup

	qret := qr.UpdateOneQrDh(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesDestHostgroup->qr.UpdateOneQrDh Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesDestHostgroup->qr.UpdateOneQrDh  Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

//更新一个规则的错误消息内容
func (pmapi *PMApi) UpdateOneQueryRulesErrmsg(c echo.Context) error {
	args := struct {
		RuleId int64  `json:"rule_id"`
		ErrMsg string `json:"error_msg"`
	}{}

	qr := new(queryrules.QueryRules)
	if err := c.Bind(&args); err != nil {
		log.Print("UpdateOneQueryRulesErrmsg->c.Bind ", err)
		return err
	}

	qr.Rule_id = args.RuleId
	qr.Error_msg = args.ErrMsg

	qret := qr.UpdateOneQrEm(pmapi.Apidb)
	if qret == 1 {
		log.Print("UpdateOneQueryRulesErrmsg->qr.UpdateOneQrEm Error")
		return c.JSON(http.StatusExpectationFailed, "UpdateOneQueryRulesErrmsg->qr.UpdateOneQrEm Error")
	}
	return c.JSON(http.StatusOK, "OK")
}

/*Patch方法的查询规则更新函数*/
func (pmapi *PMApi) UpdateOneQueryRulesInfo(c echo.Context) error {
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
		return err
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
	return c.JSON(http.StatusOK, "OK")
}

func (pmapi *PMApi) DeleteOneQueryRules(c echo.Context) error {
	qr := new(queryrules.QueryRules)
	qr.Rule_id, _ = strconv.ParseInt(c.Param("ruleid"), 10, 64)
	qret := qr.DeleteOneQr(pmapi.Apidb)
	if qret == 1 {
		log.Print("DeleteOneQueryRules->qr.DeleteOneQr Error")
		return c.JSON(http.StatusExpectationFailed, "DeleteOneQueryRules->qr.DeleteOneQr Error")
	}
	return c.JSON(http.StatusOK, "OK")
}
func (pmapi *PMApi) UpdateOneVariables(c echo.Context) error {
	args := struct {
		VariableName  string `json:"variable_name"`
		VariableValue string `json:"variable_value"`
	}{}

	if err := c.Bind(&args); err != nil {
		return err
	}

	psv := new(variables.Variables)
	log.Print("UpdateOneVariables", args)

	psv.VariablesName = args.VariableName
	psv.Value = args.VariableValue

	pret, _ := psv.UpdateOneVariable(pmapi.Apidb)
	if pret == 1 {
		return c.JSON(http.StatusExpectationFailed, "UpdateOneVariable Failed")
	}
	return c.JSON(http.StatusOK, "OK")
}

//查询出所有查询规则
func (pmapi *PMApi) ListAllQueryRules(c echo.Context) error {
	qr := new(queryrules.QueryRules)

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)

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
		return c.JSON(http.StatusExpectationFailed, "ListAllQueryRules ExpectationFailed")
	}
	return c.JSON(http.StatusOK, ret)
}

//查询出一个规则的内容
func (pmapi *PMApi) ListOneQueryRule(c echo.Context) error {
	qr := new(queryrules.QueryRules)
	if err := c.Bind(qr); err != nil {
		return err
	}
	qr.Rule_id, _ = strconv.ParseInt(c.Param("ruleid"), 10, 64)
	log.Print("ListOneQueryRule->qr.Rule_id = ", qr.Rule_id)

	ret, err := qr.FindOneQr(pmapi.Apidb)
	if err != nil {
		log.Print("ListOneQueryRules: ", err)
		return c.JSON(http.StatusExpectationFailed, "QueryRuler Exec Error")
	}

	return c.JSON(http.StatusOK, ret)
}
