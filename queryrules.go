package proxysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/juju/errors"
)

type (
	QueryRules struct {
		Rule_id               uint64 `db:"rule_id" json:"rule_id"`
		Active                uint64 `db:"active" json:"active"`
		Username              string `db:"username" json:"username"`
		Schemaname            string `db:"schemaname" json:"schemaname"`
		FlagIN                uint64 `db:"flagIN" json:"flagIN"`
		Client_addr           string `db:"client_addr" json:"client_addr"`
		Proxy_addr            string `db:"proxy_addr" json:"proxy_addr"`
		Proxy_port            string `db:"proxy_port" json:"proxy_port"`
		Digest                string `db:"digest" json:"digest"`
		Match_digest          string `db:"match_digest" json:"match_digest"`
		Match_pattern         string `db:"match_pattern" json:"match_pattern"`
		Negate_match_pattern  uint64 `db:"negate_match_pattern" json:"negate_match_pattern"`
		FlagOUT               string `db:"flagOUT" json:"flagOUT"`
		Replace_pattern       string `db:"replace_pattern" json:"replace_pattern"`
		Destination_hostgroup string `db:"destination_hostgroup" json:"destination_hostgroup"`
		Cache_ttl             string `db:"cache_ttl" json:"cache_ttl"`
		Reconnect             string `db:"reconnect" json:"reconnect"`
		Timeout               string `db:"timeout" json:"timeout"`
		Retries               string `db:"retries" json:"retries"`
		Delay                 string `db:"delay" json:"delay"`
		Mirror_flagOUT        string `db:"mirror_flagOUT" json:"mirror_flagOUT"`
		Mirror_hostgroup      string `db:"mirror_hostgroup" json:"mirror_hostgroup"`
		Error_msg             string `db:"error_msg" json:"error_msg"`
		Log                   string `db:"log" json:"log"`
		Apply                 uint64 `db:"apply" json:"apply"`
		Comment               string `db:"comment" json:"comment"`
	}
)

const (
	/*new query rules*/
	StmtAddOneQr = `
	INSERT 
	INTO 
		mysql_query_rules(rule_id,username) 
	VALUES(%d,%s)`

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

	/*find last insert rule_id*/
	StmtFindLastRuleId = `
	SELECT 
		max(rule_id)
	FROM mysql_query_rules
	WHERE 
		username = %s`

	/*update a query rules.*/
	StmtUpdateOneQr = `
	UPDATE 
		mysql_query_rules 
	SET 
		active=%d,
		username=%s,
		schemaname=%s,
		flagIN=%d,
		client_addr=%s,
		proxy_addr=%s,
		proxy_port=%s,
		digest=%s,
		match_digest=%s,
		match_pattern=%s,
		negate_match_pattern=%d,
		flagOUT=%s,
		replace_pattern=%s,
		destination_hostgroup=%s,
		cache_ttl=%s,
		reconnect=%s,
		timeout=%s,
		retries=%s,
		delay=%s,
		mirror_flagOUT=%s,
		mirror_hostgroup=%s,
		error_msg=%s,
		log=%s,
		apply=%d,
		comment=%s
	WHERE 
		rule_id=%d`
)

// select * from mysql_query_rules limit n offset n
func FindAllQr(db *sql.DB, limit uint64, skip uint64) ([]QueryRules, error) {

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

// new mysql query rules
func NewQr(username string, destination_hostgroup string) (*QueryRules, error) {
	newqr := new(QueryRules)

	if username == "" {
		return nil, errors.BadRequestf(username)
	}
	if destination_hostgroup == "" {
		return nil, errors.BadRequestf(destination_hostgroup)
	}
	newqr.Username = fmt.Sprintf("\"%s\"", username)
	newqr.Destination_hostgroup = fmt.Sprintf("\"%s\"", destination_hostgroup)

	newqr.Schemaname = "NULL"
	newqr.FlagIN = 0
	newqr.Client_addr = "NULL"
	newqr.Proxy_addr = "NULL"
	newqr.Proxy_port = "NULL"
	newqr.Digest = "NULL"
	newqr.Match_digest = "NULL"
	newqr.Match_pattern = "NULL"
	newqr.Negate_match_pattern = 0
	newqr.FlagOUT = "NULL"
	newqr.Replace_pattern = "NULL"
	newqr.Cache_ttl = "NULL"
	newqr.Reconnect = "NULL"
	newqr.Timeout = "NULL"
	newqr.Retries = "NULL"
	newqr.Delay = "NULL"
	newqr.Mirror_flagOUT = "NULL"
	newqr.Mirror_hostgroup = "NULL"
	newqr.Error_msg = "NULL"
	newqr.Log = "NULL"
	newqr.Apply = 0
	newqr.Active = 0
	newqr.Comment = "NULL"

	return newqr, nil
}

// set qr rule_id
func (qr *QueryRules) SetQrRuleid(rule_id uint64) {
	qr.Rule_id = rule_id
}

// set qr active
func (qr *QueryRules) SetQrActive(active uint64) {
	switch active {
	case 0:
		qr.Active = 0
	case 1:
		qr.Active = 1
	default:
		qr.Active = 1
	}
}

// set qr apply
func (qr *QueryRules) SetQrApply(apply uint64) {
	switch apply {
	case 0:
		qr.Apply = 0
	case 1:
		qr.Apply = 1
	default:
		qr.Apply = 1
	}
}

// set qr schemaname
func (qr *QueryRules) SetQrSchemaname(schema_name string) {
	if schema_name == "" || len(schema_name) == 0 {
		qr.Schemaname = "NULL"
	} else {
		qr.Schemaname = fmt.Sprintf("\"%s\"", schema_name)
	}
}

// set qr flagIN
func (qr *QueryRules) SetQrFlagIN(flag_in uint64) {
	qr.FlagIN = flag_in
}

// set qr client_addr
func (qr *QueryRules) SetQrClientAddr(client_addr string) {
	if client_addr == "" || len(client_addr) == 0 {
		qr.Client_addr = "NULL"
	} else {
		qr.Client_addr = fmt.Sprintf("\"%s\"", client_addr)
	}
}

// set qr proxy_addr
func (qr *QueryRules) SetQrProxyAddr(proxy_addr string) {
	if proxy_addr == "" || len(proxy_addr) == 0 {
		qr.Proxy_addr = "NULL"
	} else {
		qr.Proxy_addr = fmt.Sprintf("\"%s\"", proxy_addr)
	}
}

// set qr proxy_port
func (qr *QueryRules) SetProxyPort(proxy_port string) {
	if proxy_port == "" || len(proxy_port) == 0 {
		qr.Proxy_port = "NULL"
	} else {
		qr.Proxy_port = fmt.Sprintf("\"%s\"", proxy_port)
	}
}

// set qr digest
func (qr *QueryRules) SetQrDigest(digest string) {
	if digest == "" || len(digest) == 0 {
		qr.Digest = "NULL"
	} else {
		qr.Digest = fmt.Sprintf("\"%s\"", digest)
	}
}

// set qr match_digest
func (qr *QueryRules) SetQrMatchDigest(match_digest string) {
	if match_digest == "" || len(match_digest) == 0 {
		qr.Match_digest = "NULL"
	} else {
		qr.Match_digest = fmt.Sprintf("\"%s\"", match_digest)
	}
}

// set qr match_pattern
func (qr *QueryRules) SetQrMatchPattern(match_pattern string) {
	if match_pattern == "" || len(match_pattern) == 0 {
		qr.Match_pattern = "NULL"
	} else {
		qr.Match_pattern = fmt.Sprintf("\"%s\"", match_pattern)
	}
}

// set qr mnegate_match_pattern
func (qr *QueryRules) SetQrNegateMatchPattern(negate_match_pattern uint64) {
	switch negate_match_pattern {
	case 0:
		qr.Negate_match_pattern = 0
	case 1:
		qr.Negate_match_pattern = 1
	default:
		qr.Negate_match_pattern = 0
	}
}

// set qr flagout
func (qr *QueryRules) SetQrFlagOut(flag_out string) {
	if flag_out == "" || len(flag_out) == 0 {
		qr.FlagOUT = "NULL"
	} else {
		qr.FlagOUT = fmt.Sprintf("\"%s\"", flag_out)
	}
}

// set qr replace_pattern
func (qr *QueryRules) SetQrReplacePattern(replace_pattern string) {
	if replace_pattern == "" || len(replace_pattern) == 0 {
		qr.Replace_pattern = "NULL"
	} else {
		qr.Replace_pattern = fmt.Sprintf("\"%s\"", replace_pattern)
	}
}

// set qr destination_hostgroup
func (qr *QueryRules) SetQrDestHostGroup(destination_hostgroup string) {
	if destination_hostgroup == "" || len(destination_hostgroup) == 0 {
		qr.Destination_hostgroup = "NULL"
	} else {
		qr.Destination_hostgroup = fmt.Sprintf("\"%s\"", destination_hostgroup)
	}
}

// set qr cache_ttl
func (qr *QueryRules) SetQrCacheTTL(cache_ttl string) {
	if cache_ttl == "" || len(cache_ttl) == 0 {
		qr.Cache_ttl = "NULL"
	} else {
		qr.Cache_ttl = fmt.Sprintf("\"%s\"", cache_ttl)
	}
}

// set qr reconnect
func (qr *QueryRules) SetQrReconnect(reconnect string) {
	if reconnect == "" || len(reconnect) == 0 {
		qr.Reconnect = "NULL"
	} else {
		qr.Reconnect = fmt.Sprintf("\"%s\"", reconnect)
	}
}

// set qr timeout
func (qr *QueryRules) SetQrTimeOut(timeout string) {
	if timeout == "" || len(timeout) == 0 {
		qr.Timeout = "NULL"
	} else {
		qr.Timeout = fmt.Sprintf("\"%s\"", timeout)
	}
}

// set qr retries
func (qr *QueryRules) SetQrRetries(retries string) {
	if retries == "" || len(retries) == 0 {
		qr.Retries = "NULL"
	} else {
		qr.Retries = fmt.Sprintf("\"%s\"", retries)
	}
}

// set qr delay
func (qr *QueryRules) SetQrDelay(delay string) {
	if delay == "" || len(delay) == 0 {
		qr.Delay = "NULL"
	} else {
		qr.Delay = fmt.Sprintf("\"%s\"", delay)
	}
}

// set qr mirror_flagout
func (qr *QueryRules) SetQrMirrorFlagOUT(mirror_flagout string) {
	if mirror_flagout == "" || len(mirror_flagout) == 0 {
		qr.Mirror_flagOUT = "NULL"
	} else {
		qr.Mirror_flagOUT = fmt.Sprintf("\"%s\"", mirror_flagout)
	}
}

// set qr mirror_hostgroup
func (qr *QueryRules) SetQrMirrorHostgroup(mirror_hostgroup string) {
	if mirror_hostgroup == "" || len(mirror_hostgroup) == 0 {
		qr.Mirror_hostgroup = "NULL"
	} else {
		qr.Mirror_hostgroup = fmt.Sprintf("\"%s\"", mirror_hostgroup)
	}
}

// set qr error_msg
func (qr *QueryRules) SetQrErrorMsg(error_msg string) {
	if error_msg == "" || len(error_msg) == 0 {
		qr.Error_msg = "NULL"
	} else {
		qr.Error_msg = fmt.Sprintf("\"%s\"", error_msg)
	}
}

// set qr log
func (qr *QueryRules) SetQrLog(log string) {
	if log == "" || len(log) == 0 {
		qr.Log = "NULL"
	} else {
		qr.Log = fmt.Sprintf("\"%s\"", log)
	}
}

// add a new query rules.
func (qr *QueryRules) AddOneQr(db *sql.DB) error {

	Query := fmt.Sprintf(StmtAddOneQr, qr.Rule_id, qr.Username)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err) //add user failed
	}

	Query = fmt.Sprintf(StmtFindLastRuleId, qr.Username)
	rows := db.QueryRow(Query)

	/*
		FIX:
		It will always return 0 when you use sql.Result.LastInsertId() function to get last inserted row id.
		the proxysql not support transaction.
		So,I Query a max(id) after insert a row.
	*/
	err = rows.Scan(&qr.Rule_id)
	if err != nil {
		return errors.Trace(err)
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return nil
}

//delete a query rules.
func (qr *QueryRules) DeleteOneQr(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneQr, qr.Rule_id)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(strconv.Itoa(int(qr.Rule_id)))
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return nil
}

//update a query rules.
func (qr *QueryRules) UpdateOneQrInfo(db *sql.DB) error {

	var Query string

	Query = fmt.Sprintf(StmtUpdateOneQr,
		qr.Active,
		qr.Username,
		qr.Schemaname,
		qr.FlagIN,
		qr.Client_addr,
		qr.Proxy_addr,
		qr.Proxy_port,
		qr.Digest,
		qr.Match_digest,
		qr.Match_pattern,
		qr.Negate_match_pattern,
		qr.FlagOUT,
		qr.Replace_pattern,
		qr.Destination_hostgroup,
		qr.Cache_ttl,
		qr.Reconnect,
		qr.Timeout,
		qr.Retries,
		qr.Delay,
		qr.Mirror_flagOUT,
		qr.Mirror_hostgroup,
		qr.Error_msg,
		qr.Log,
		qr.Apply,
		qr.Comment,
		qr.Rule_id)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(strconv.Itoa(int(qr.Rule_id)))
	}

	LoadQueryRulesToRuntime(db)
	SaveQueryRulesToDisk(db)

	return nil
}
