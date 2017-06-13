package queryrules

import (
	"database/sql"
	"fmt"
	"log"
)

type (
	QueryRules struct {
		Rule_id               string `db:"Rule_id" json:"Rule_id"`
		Active                string `db:"Active" json:"Active"`
		Username              string `db:"Username" json:"Username"`
		Schemaname            string `db:"Schemaname" json:"Schemaname"`
		FlagIN                string `db:"FlagIN" json:"FlagIN"`
		Client_addr           string `db:"Client_addr" json:"Client_addr"`
		Proxy_addr            string `db:"Proxy_addr" json:"Proxy_addr"`
		Proxy_port            string `db:"Proxy_port" json:"Proxy_port"`
		Digest                string `db:"Digest" json:"Digest"`
		Match_digest          string `db:"Match_digest" json:"Match_digest"`
		Match_pattern         string `db:"Match_pattern" json:"Match_pattern"`
		Negate_match_pattern  string `db:"Negate_match_pattern" json:"Negate_match_pattern"`
		FlagOUT               string `db:"FlagOUT" json:"FlagOUT"`
		Replace_pattern       string `db:"Replace_pattern" json:"Replace_pattern"`
		Destination_hostgroup string `db:"Destination_hostgroup" json:"Destination_hostgroup"`
		Cache_ttl             string `db:"Cache_ttl" json:"Cache_ttl"`
		Reconnect             string `db:"Reconnect" json:"Reconnect"`
		Timeout               string `db:"Timeout" json:"Timeout"`
		Retries               string `db:"Retries" json:"Retries"`
		Delay                 string `db:"Delay" json:"Delay"`
		Mirror_flagOUT        string `db:"Mirror_flagOUT" json:"Mirror_flagOUT"`
		Mirror_hostgroup      string `db:"Mirror_hostgroup" json:"Mirror_hostgroup"`
		Error_msg             string `db:"Error_msg" json:"Error_msg"`
		Log                   string `db:"Log" json:"Log"`
		Apply                 string `db:"Apply" json:"Apply"`
		Comment               string `db:"Comment" json:"Comment"`
	}
)

const (
	StmtAddOneQr       = `INSERT INTO mysql_query_rules(username) VALUES(%q)`
	StmtDeleteOneQr    = `DELETE FROM mysql_query_rules WHERE rule_id = %d`
	StmtActiveOneQr    = `UPDATE mysql_query_rules SET active =1 AND apply=1 WHERE rule_id=%d`
	StmtDisactiveOneQr = `UPDATE mysql_query_rules SET active 0 AND  apply=0 WHERE rule_id=%d`
	StmtFindOneQr      = `SELECT * FROM mysql_query_rules WHERE rule_id = %d`
	StmtFindAllQr      = `SELECT * FROM mysql_query_rules`
	StmtUpdateOneQrUn  = `UPDATE mysql_query_rules SET username =%q WHERE rule_id = %d`
	StmtUpdateOneQrSn  = `UPDATE mysql_query_rules SET schemaname = %q WHERE rule_id = %d`
	StmtUpdateOneQrCa  = `UPDATE mysql_query_rules SET client_addr = %q WHERE rule_id = %d`
	StmtUpdateOneQrMd  = `UPDATE mysql_query_rules SET match_digest = %q WHERE rule_id = %d`
	StmtUpdateOneQrMp  = `UPDATE mysql_query_rules SET match_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrRp  = `UPDATE mysql_query_rules SET replace_pattern = %q WHERE rule_id = %d`
	StmtUpdateOneQrDh  = `UPDATE mysql_query_rules SET destination_hostgroup = %q WHERE rule_id = %d`
	StmtUpdateOneQrEm  = `UPDATE mysql_query_rules SET error_msg = %q WHERE rule_id = %d`
)

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
func (qr *QueryRules) FindOneQr(db *sql.DB) int {
	st := fmt.Sprintf(StmtFindOneQr, qr.Rule_id)
	log.Print("FindOneQr: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("FindOneQr: ", err)
		return 1
	}
	log.Print("FindOneQr: Success")
	return 0
}

//获取所有查询规则的内容
func (qr *QueryRules) FindAllQr(db *sql.DB) int {
	log.Print("FindAllQr:", StmtFindAllQr)
	_, err := db.Query(StmtFindAllQr)
	if err != nil {
		log.Print("FindAllQr: ", err)
		return 1
	}
	log.Print("FindAllQr: Success")
	return 0
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
	st := fmt.Sprint(StmtUpdateOneQrSn, qr.Schemaname, qr.Rule_id)
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
