package queryrules

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type (
	QueryRules struct {
		Rule_id               int64  `db:"rule_id" json:"Rule_id"`
		Active                int64  `db:"active" json:"Active"`
		Username              string `db:"username" json:"Username"`
		Schemaname            string `db:"schemaname" json:"Schemaname"`
		FlagIN                int64  `db:"flagIN" json:"FlagIN"`
		Client_addr           string `db:"client_addr" json:"Client_addr"`
		Proxy_addr            string `db:"proxy_addr" json:"Proxy_addr"`
		Proxy_port            int64  `db:"proxy_port" json:"Proxy_port"`
		Digest                string `db:"digest" json:"Digest"`
		Match_digest          string `db:"match_digest" json:"Match_digest"`
		Match_pattern         string `db:"match_pattern" json:"Match_pattern"`
		Negate_match_pattern  int64  `db:"negate_match_pattern" json:"Negate_match_pattern"`
		FlagOUT               int64  `db:"flagOUT" json:"FlagOUT"`
		Replace_pattern       string `db:"replace_pattern" json:"Replace_pattern"`
		Destination_hostgroup int64  `db:"destination_hostgroup" json:"Destination_hostgroup"`
		Cache_ttl             int64  `db:"cache_ttl" json:"Cache_ttl"`
		Reconnect             int64  `db:"reconnect" json:"Reconnect"`
		Timeout               int64  `db:"timeout" json:"Timeout"`
		Retries               int64  `db:"retries" json:"Retries"`
		Delay                 int64  `db:"delay" json:"Delay"`
		Mirror_flagOUT        int64  `db:"mirror_flagOUT" json:"Mirror_flagOUT"`
		Mirror_hostgroup      int64  `db:"mirror_hostgroup" json:"Mirror_hostgroup"`
		Error_msg             string `db:"error_msg" json:"Error_msg"`
		Log                   int64  `db:"log" json:"Log"`
		Apply                 int64  `db:"apply" json:"Apply"`
		Comment               string `db:"comment" json:"Comment"`
	}
)

const (
	StmtQrExists       = `SELECT COUNT(*) FROM mysql_query_rules WHERE rule_id = %d`
	StmtAddOneQr       = `INSERT INTO mysql_query_rules(username) VALUES(%q)`
	StmtDeleteOneQr    = `DELETE FROM mysql_query_rules WHERE rule_id = %d`
	StmtActiveOneQr    = `UPDATE mysql_query_rules SET active =1 AND apply=1 WHERE rule_id=%d`
	StmtDisactiveOneQr = `UPDATE mysql_query_rules SET active =0 AND  apply=0 WHERE rule_id=%d`
	StmtFindOneQr      = `select ifnull(rule_id,0),ifnull(active,0),ifnull(username,""),ifnull(schemaname,""),ifnull(flagIN,0),ifnull(client_addr,""),ifnull(proxy_addr,""),ifnull(proxy_port,0),ifnull(digest,""),ifnull(match_digest,""),ifnull(match_pattern,""),ifnull(negate_match_pattern,0),ifnull(flagOUT,0),ifnull(replace_pattern,""),ifnull(destination_hostgroup,0),ifnull(cache_ttl,0),ifnull(reconnect,0),ifnull(timeout,0),ifnull(retries,0),ifnull(delay,0),ifnull(mirror_flagOUT,0),ifnull(mirror_hostgroup,0),ifnull(error_msg,""),ifnull(log,0),ifnull(apply,0),ifnull(comment,"") from mysql_query_rules WHERE rule_id = %d`
	StmtFindAllQr      = `select ifnull(rule_id,0),ifnull(active,0),ifnull(username,""),ifnull(schemaname,""),ifnull(flagIN,0),ifnull(client_addr,""),ifnull(proxy_addr,""),ifnull(proxy_port,0),ifnull(digest,""),ifnull(match_digest,""),ifnull(match_pattern,""),ifnull(negate_match_pattern,0),ifnull(flagOUT,0),ifnull(replace_pattern,""),ifnull(destination_hostgroup,0),ifnull(cache_ttl,0),ifnull(reconnect,0),ifnull(timeout,0),ifnull(retries,0),ifnull(delay,0),ifnull(mirror_flagOUT,0),ifnull(mirror_hostgroup,0),ifnull(error_msg,""),ifnull(log,0),ifnull(apply,0),ifnull(comment,"") from mysql_query_rules limit %d offset %d`
	StmtUpdateOneQrUn  = `UPDATE mysql_query_rules SET username =%q WHERE rule_id = %d`
	StmtUpdateOneQrSn  = `UPDATE mysql_query_rules SET schemaname = %q WHERE rule_id = %d`
	StmtUpdateOneQrCa  = `UPDATE mysql_query_rules SET client_addr = %q WHERE rule_id = %d`
	StmtUpdateOneQrMd  = `UPDATE mysql_query_rules SET match_digest = %q WHERE rule_id = %d`
	StmtUpdateOneQrMp  = `UPDATE mysql_query_rules SET match_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrRp  = `UPDATE mysql_query_rules SET replace_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrDh  = `UPDATE mysql_query_rules SET destination_hostgroup = %q WHERE rule_id = %d`
	StmtUpdateOneQrEm  = `UPDATE mysql_query_rules SET error_msg = %q WHERE rule_id = %d`
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
	log.Print("FindAllQr:", StmtFindAllQr)
	rows, err := db.Query(StmtFindAllQr, limit, skip)
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
	return 0
}
