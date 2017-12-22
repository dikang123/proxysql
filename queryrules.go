package proxysql

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
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
	/*new query rules*/
	StmtAddOneQr = `
	INSERT 
	INTO 
		mysql_query_rules(username) 
	VALUES(%q)`

	/*delete a query rules*/
	StmtDeleteOneQr = `
	DELETE 
	FROM 
		mysql_query_rules 
	WHERE rule_id = %d`

	/*query all query rules.*/
	StmtFindAllQr = `
	SELECT 
		ifnull(rule_id,0) as rule_id,
		ifnull(active,0) as active,
		ifnull(username,"") as username,
		ifnull(schemaname,"") as schemaname,
		ifnull(flagIN,0) as flagIN,
		ifnull(client_addr,"") as client_addr,
		ifnull(proxy_addr,"") as proxy_addr,
		ifnull(proxy_port,0) as proxy_port,
		ifnull(digest,"") as digest,
		ifnull(match_digest,"") as match_digest,
		ifnull(match_pattern,"") as match_pattern,
		ifnull(negate_match_pattern,0) as negate_match_pattern,
		ifnull(flagOUT,0) as flagOUT,
		ifnull(replace_pattern,"") as replace_pattern,
		ifnull(destination_hostgroup,0) as destination_hostgroup,
		ifnull(cache_ttl,0) as cache_ttl,
		ifnull(reconnect,0) as reconnect,
		ifnull(timeout,0) as timeout,
		ifnull(retries,0) as retries,
		ifnull(delay,0) as delay,
		ifnull(mirror_flagOUT,0) as mirror_flagOUT,
		ifnull(mirror_hostgroup,0) as mirror_hostgroup,
		ifnull(error_msg,"") as error_msg,
		ifnull(log,0) as log,
		ifnull(apply,0) as apply,
		ifnull(comment,"") as comment 
	FROM mysql_query_rules 
	LIMIT %d 
	OFFSET %d`

	/*update a query rules.*/
	StmtUpdateOneQr = `
	UPDATE 
		mysql_query_rules 
	SET 
		active=%d,
		username=%s,
		schemaname=%s,
		client_addr=%s,
		digest=%s,
		match_digest=%s,
		match_pattern=%s,
		replace_pattern=%s,
		destination_hostgroup=%d,
		cache_ttl=%d,
		error_msg=%s,
		apply=%d 
	WHERE 
		rule_id=%d`

	StmtUpdateOneQrNoCache = `
	UPDATE 
		mysql_query_rules 
	SET 
		active=%d,
		username=%s,
		schemaname=%s,
		client_addr=%s,
		digest=%s,
		match_digest=%s,
		match_pattern=%s,
		replace_pattern=%s,
		destination_hostgroup=%d,
		error_msg=%s,
		apply=%d 
	WHERE 
		rule_id=%d`
)

// select * from mysql_query_rules limit n offset n
func (qr *QueryRules) FindAllQr(db *sql.DB, limit int64, skip int64) ([]QueryRules, error) {

	var AllQr []QueryRules
	Query := fmt.Sprintf(StmtFindAllQr, limit, skip)

	// exec query statement
	rows, err := db.Query(Query)
	if err != nil {
		return []QueryRules{}, errors.Trace(err)
	}
	defer rows.Close()

	// scan results.
	for rows.Next() {

		var tmpqr QueryRules

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
			continue
		}

		AllQr = append(AllQr, tmpqr)
	}
	return AllQr, nil
}

// add a new query rules.
func (qr *QueryRules) AddOneQr(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtAddOneQr, qr.Username)

	_, err := db.Exec(Query)
	if err != nil {
		return 1, errors.Trace(err)
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return 0, nil
}

//delete a query rules.
func (qr *QueryRules) DeleteOneQr(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtDeleteOneQr, qr.Rule_id)

	_, err := db.Exec(Query)
	if err != nil {
		return 1, errors.Trace(err)
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return 0, nil
}

func convertString(cs string) string {
	var cstmp string
	if cs == "" {
		cstmp = "NULL"
	} else {
		cstmp = fmt.Sprintf("'%s'", cs)
	}
	return cstmp
}

//update a query rules.
func (qr *QueryRules) UpdateOneQrInfo(db *sql.DB) (int, error) {

	var Query string
	qr.Username = convertString(qr.Username)
	qr.Schemaname = convertString(qr.Schemaname)
	qr.Client_addr = convertString(qr.Client_addr)
	qr.Digest = convertString(qr.Digest)
	qr.Match_digest = convertString(qr.Match_digest)
	qr.Match_pattern = convertString(qr.Match_pattern)
	qr.Replace_pattern = convertString(qr.Replace_pattern)
	qr.Error_msg = convertString(qr.Error_msg)

	if qr.Cache_ttl == 0 {
		Query = fmt.Sprintf(StmtUpdateOneQrNoCache,
			qr.Active,
			qr.Username,
			qr.Schemaname,
			qr.Client_addr,
			qr.Digest,
			qr.Match_digest,
			qr.Match_pattern,
			qr.Replace_pattern,
			qr.Destination_hostgroup,
			qr.Error_msg,
			qr.Active,
			qr.Rule_id)
	} else {
		Query = fmt.Sprintf(StmtUpdateOneQr,
			qr.Active,
			qr.Username,
			qr.Schemaname,
			qr.Client_addr,
			qr.Digest,
			qr.Match_digest,
			qr.Match_pattern,
			qr.Replace_pattern,
			qr.Destination_hostgroup,
			qr.Cache_ttl,
			qr.Error_msg,
			qr.Active,
			qr.Rule_id)

	}

	_, err := db.Exec(Query)
	if err != nil {
		return 1, errors.Trace(err)
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return 0, nil
}
