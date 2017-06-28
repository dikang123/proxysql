package queryrules

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"proxysql-master/admin/cmd"
)

type (
	QueryRules struct {
		Rule_id               int64  `db:"rule_id" json:"rule_id"`
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
	}
)

const (
	StmtQrExists       = `SELECT COUNT(*) FROM mysql_query_rules WHERE rule_id = %d`
	StmtAddOneQr       = `INSERT INTO mysql_query_rules(username) VALUES(%q)`
	StmtDeleteOneQr    = `DELETE FROM mysql_query_rules WHERE rule_id = %d`
	StmtActiveOneQr    = `UPDATE mysql_query_rules SET active =1 ,apply=1 WHERE rule_id=%d`
	StmtDisactiveOneQr = `UPDATE mysql_query_rules SET active =0 ,apply=0 WHERE rule_id=%d`
	StmtFindOneQr      = `select ifnull(rule_id,0) as rule_id,ifnull(active,0) as active,ifnull(username,"") as username,ifnull(schemaname,"") as schemaname,ifnull(flagIN,0) as flagIN,ifnull(client_addr,"") as client_addr,ifnull(proxy_addr,"") as proxy_addr,ifnull(proxy_port,0) as proxy_port,ifnull(digest,"") as digest,ifnull(match_digest,"") as match_digest,ifnull(match_pattern,"") as match_pattern,ifnull(negate_match_pattern,0) as negate_match_pattern,ifnull(flagOUT,0) as flagOUT,ifnull(replace_pattern,"") as replace_pattern,ifnull(destination_hostgroup,0) as destination_hostgroup,ifnull(cache_ttl,0) as cache_ttl,ifnull(reconnect,0) as reconnect,ifnull(timeout,0) as timeout,ifnull(retries,0) as retries,ifnull(delay,0) as delay,ifnull(mirror_flagOUT,0) as mirror_flagOUT,ifnull(mirror_hostgroup,0) as mirror_hostgroup,ifnull(error_msg,"") as error_msg,ifnull(log,0) as log,ifnull(apply,0) as apply,ifnull(comment,"") as comment from mysql_query_rules WHERE rule_id = %d`
	StmtFindAllQr      = `select ifnull(rule_id,0) as rule_id,ifnull(active,0) as active,ifnull(username,"") as username,ifnull(schemaname,"") as schemaname,ifnull(flagIN,0) as flagIN,ifnull(client_addr,"") as client_addr,ifnull(proxy_addr,"") as proxy_addr,ifnull(proxy_port,0) as proxy_port,ifnull(digest,"") as digest,ifnull(match_digest,"") as match_digest,ifnull(match_pattern,"") as match_pattern,ifnull(negate_match_pattern,0) as negate_match_pattern,ifnull(flagOUT,0) as flagOUT,ifnull(replace_pattern,"") as replace_pattern,ifnull(destination_hostgroup,0) as destination_hostgroup,ifnull(cache_ttl,0) as cache_ttl,ifnull(reconnect,0) as reconnect,ifnull(timeout,0) as timeout,ifnull(retries,0) as retries,ifnull(delay,0) as delay,ifnull(mirror_flagOUT,0) as mirror_flagOUT,ifnull(mirror_hostgroup,0) as mirror_hostgroup,ifnull(error_msg,"") as error_msg,ifnull(log,0) as log,ifnull(apply,0) as apply,ifnull(comment,"") as comment from mysql_query_rules limit %d offset %d`
	StmtUpdateOneQrUn  = `UPDATE mysql_query_rules SET username =%q WHERE rule_id = %d`
	StmtUpdateOneQrSn  = `UPDATE mysql_query_rules SET schemaname = %q WHERE rule_id = %d`
	StmtUpdateOneQrCa  = `UPDATE mysql_query_rules SET client_addr = %q WHERE rule_id = %d`
	StmtUpdateOneQrMd  = `UPDATE mysql_query_rules SET match_digest = %q WHERE rule_id = %d`
	StmtUpdateOneQrDg  = `UPDATE mysql_query_rules SET digest = %q WHERE rule_id = %d`
	StmtUpdateOneQrMp  = `UPDATE mysql_query_rules SET match_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrRp  = `UPDATE mysql_query_rules SET replace_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrDh  = `UPDATE mysql_query_rules SET destination_hostgroup = %d WHERE rule_id = %d`
	StmtUpdateOneQrEm  = `UPDATE mysql_query_rules SET error_msg = %q WHERE rule_id = %d`
	StmtUpdateOneQr    = `UPDATE mysql_query_rules SET active=%d,username=%q,schemaname=%q,client_addr=%q,digest=%q,match_digest=%q,match_pattern=%q,replace_pattern=%q,destination_hostgroup=%d,cache_ttl=%d,error_msg=%q,apply=%d WHERE rule_id=%d`
)

//查询指定规则id是否存在
func (qr *QueryRules) QrExists(db *sql.DB) int {
	st := fmt.Sprintf(StmtQrExists, qr.Rule_id)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("QrExists: ", err)
		return 2
	}
	var QrCount uint64
	for rows.Next() {
		err = rows.Scan(&QrCount)
		if err != nil {
			log.Print("QrExists rows.Next: ", err)
			return 3
		}
	}
	if QrCount == 0 {
		return 0
	} else {
		return 1
	}
}

//添加一个新的查询规则
func (qr *QueryRules) AddOneQr(db *sql.DB) int {
	st := fmt.Sprintf(StmtAddOneQr, qr.Username)
	log.Print("AddOneQr: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("AddOneQr: ", err)
		return 1
	}
	log.Print("AddOneQr: Add Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//删除一个已有的查询规则
func (qr *QueryRules) DeleteOneQr(db *sql.DB) int {
	st := fmt.Sprintf(StmtDeleteOneQr, qr.Rule_id)
	log.Print("DeleteOneQr: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("DeleteOneQr: ", err)
		return 1
	}
	log.Print("DeleteOneQr: Delete Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//激活一个查询规则
func (qr *QueryRules) ActiveOneQr(db *sql.DB) int {
	st := fmt.Sprintf(StmtActiveOneQr, qr.Rule_id)
	log.Print("ActiveOneQr: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("ActiveOneQr: ", err)
		return 1
	}
	log.Print("ActiveOneQr: ActiveOneQr Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//关闭一个查询规则
func (qr *QueryRules) DisactiveOneQr(db *sql.DB) int {
	st := fmt.Sprintf(StmtDisactiveOneQr, qr.Rule_id)
	log.Print("DisactiveOneQr: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("DisactiveOneQr: ", err)
		return 1
	}
	log.Print("DisactiveOneQr: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//获取一个查询规则内容
func (qr *QueryRules) FindOneQr(db *sql.DB) (QueryRules, error) {
	if qr.QrExists(db) == 2 {
		log.Print("FindOneQr->QrExists == 2")
		return QueryRules{}, errors.New("QueryRuler Exec Error")
	}
	if qr.QrExists(db) == 0 {
		log.Print("FindOneQr->QrExists == 0")
		return QueryRules{}, errors.New("Query Rules Not Exists")
	}
	var tmpqr QueryRules
	st := fmt.Sprintf(StmtFindOneQr, qr.Rule_id)
	log.Print("FindOneQr: ", st)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("FindOneQr: ", err)
		return QueryRules{}, err
	}
	log.Print("FindOneQr: Success")
	for rows.Next() {
		tmpqr = QueryRules{}
		err = rows.Scan(
			&tmpqr.Rule_id,
			&tmpqr.Active,
			&tmpqr.Username,
			&tmpqr.Schemaname,
			&tmpqr.FlagIN,
			&tmpqr.Client_addr,
			&tmpqr.Proxy_addr,
			&tmpqr.Proxy_port,
			&tmpqr.Digest,
			&tmpqr.Match_digest,
			&tmpqr.Match_pattern,
			&tmpqr.Negate_match_pattern,
			&tmpqr.FlagOUT,
			&tmpqr.Replace_pattern,
			&tmpqr.Destination_hostgroup,
			&tmpqr.Cache_ttl,
			&tmpqr.Reconnect,
			&tmpqr.Timeout,
			&tmpqr.Retries,
			&tmpqr.Delay,
			&tmpqr.Mirror_flagOUT,
			&tmpqr.Mirror_hostgroup,
			&tmpqr.Error_msg,
			&tmpqr.Log,
			&tmpqr.Apply,
			&tmpqr.Comment,
		)

		if err != nil {
			log.Print("FindOneQr err1 ", err.Error())
			continue
		}

	}

	err = rows.Err()
	if err != nil {
		log.Print("FindOneQr err 2:", err.Error())
	}
	log.Print("FindOneQr: Success")
	return tmpqr, nil
}

//获取所有查询规则的内容
func (qr *QueryRules) FindAllQr(db *sql.DB, limit int64, skip int64) ([]QueryRules, error) {
	var AllQr []QueryRules
	log.Printf(StmtFindAllQr, limit, skip)
	st := fmt.Sprintf(StmtFindAllQr, limit, skip)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("FindAllQr: ", err)
		return []QueryRules{}, errors.New("FindAllQr db.Query Exec Error")
	}
	log.Print("FindAllQr: Success")
	for rows.Next() {
		var tmpqr QueryRules = QueryRules{}
		err = rows.Scan(
			&tmpqr.Rule_id,
			&tmpqr.Active,
			&tmpqr.Username,
			&tmpqr.Schemaname,
			&tmpqr.FlagIN,
			&tmpqr.Client_addr,
			&tmpqr.Proxy_addr,
			&tmpqr.Proxy_port,
			&tmpqr.Digest,
			&tmpqr.Match_digest,
			&tmpqr.Match_pattern,
			&tmpqr.Negate_match_pattern,
			&tmpqr.FlagOUT,
			&tmpqr.Replace_pattern,
			&tmpqr.Destination_hostgroup,
			&tmpqr.Cache_ttl,
			&tmpqr.Reconnect,
			&tmpqr.Timeout,
			&tmpqr.Retries,
			&tmpqr.Delay,
			&tmpqr.Mirror_flagOUT,
			&tmpqr.Mirror_hostgroup,
			&tmpqr.Error_msg,
			&tmpqr.Log,
			&tmpqr.Apply,
			&tmpqr.Comment,
		)
		log.Printf("tmpqr = %#v", tmpqr)
		AllQr = append(AllQr, tmpqr)
	}
	return AllQr, nil
}

//更新一个查询规则的用户名称
func (qr *QueryRules) UpdateOneQrUn(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrUn, qr.Username, qr.Rule_id)
	log.Print("UpdateOneQrUn: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrUn: ", err)
		return 1
	}
	log.Print("UpdateOneQrUn: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的默认模式名
func (qr *QueryRules) UpdateOneQrSn(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrSn, qr.Schemaname, qr.Rule_id)
	log.Print("UpdateOneQrSn: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrSn: ", err)
		return 1
	}
	log.Print("UpdateOneQrSn: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的客户端地址
func (qr *QueryRules) UpdateOneQrCa(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrCa, qr.Client_addr, qr.Rule_id)
	log.Print("UpdateOneQrCa: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrCa: ", err)
		return 1
	}
	log.Print("UpdateOneQrCa: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的digest号
func (qr *QueryRules) UpdateOneQrDg(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrDg, qr.Digest, qr.Rule_id)
	log.Print("UpdateOneQrDg: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrDg: ", err)
		return 1
	}
	log.Print("UpdateOneQrDg: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的匹配语句
func (qr *QueryRules) UpdateOneQrMd(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrMd, qr.Match_digest, qr.Rule_id)
	log.Print("UpdateOneQrMd: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrMd: ", err)
		return 1
	}
	log.Print("UpdateOneQrMd: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的匹配正则
func (qr *QueryRules) UpdateOneQrMp(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrMp, qr.Match_pattern, qr.Rule_id)
	log.Print("UpdateOneQrMp: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrMp: ", err)
		return 1
	}
	log.Print("UpdateOneQrMp: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的替换正则
func (qr *QueryRules) UpdateOneQrRp(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrRp, qr.Replace_pattern, qr.Rule_id)
	log.Print("UpdateOneQrRp: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrRp: ", err)
		return 1
	}
	log.Print("UpdateOneQrRp: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的默认主机组
func (qr *QueryRules) UpdateOneQrDh(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrDh, qr.Destination_hostgroup, qr.Rule_id)
	log.Print("UpdateOneQrDh: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrDh: ", err)
		return 1
	}
	log.Print("UpdateOneQrDh: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//更新一个查询规则的错误信息
func (qr *QueryRules) UpdateOneQrEm(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQrEm, qr.Error_msg, qr.Rule_id)
	log.Print("UpdateOneQrEm: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneQrEm: ", st)
		return 1
	}
	log.Print("UpdateOneQrEm: Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}

//添加一个Patch方法更新查询规则信息
func (qr *QueryRules) UpdateOneQrInfo(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneQr, qr.Active, qr.Username, qr.Schemaname, qr.Client_addr, qr.Digest, qr.Match_digest, qr.Match_pattern, qr.Replace_pattern, qr.Destination_hostgroup, qr.Cache_ttl, qr.Error_msg, qr.Apply, qr.Rule_id)
	log.Print("queryrules->UpdateOneQrInfo->st: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("queryrules->UpdateOneQrInfo->err :", err)
		return 1
	}
	log.Print("queryrules->UpdateOneQrInfo-> Success")
	cmd.LoadQueryRulesToRuntime(db)
	cmd.SaveQueryRulesToDisk(db)
	return 0
}
