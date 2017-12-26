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
		Proxy_port            uint64 `db:"proxy_port" json:"proxy_port"`
		Digest                string `db:"digest" json:"digest"`
		Match_digest          string `db:"match_digest" json:"match_digest"`
		Match_pattern         string `db:"match_pattern" json:"match_pattern"`
		Negate_match_pattern  uint64 `db:"negate_match_pattern" json:"negate_match_pattern"`
		FlagOUT               uint64 `db:"flagOUT" json:"flagOUT"`
		Replace_pattern       string `db:"replace_pattern" json:"replace_pattern"`
		Destination_hostgroup uint64 `db:"destination_hostgroup" json:"destination_hostgroup"`
		Cache_ttl             uint64 `db:"cache_ttl" json:"cache_ttl"`
		Reconnect             uint64 `db:"reconnect" json:"reconnect"`
		Timeout               uint64 `db:"timeout" json:"timeout"`
		Retries               uint64 `db:"retries" json:"retries"`
		Delay                 uint64 `db:"delay" json:"delay"`
		Mirror_flagOUT        uint64 `db:"mirror_flagOUT" json:"mirror_flagOUT"`
		Mirror_hostgroup      uint64 `db:"mirror_hostgroup" json:"mirror_hostgroup"`
		Error_msg             string `db:"error_msg" json:"error_msg"`
		Log                   uint64 `db:"log" json:"log"`
		Apply                 uint64 `db:"apply" json:"apply"`
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
func NewQr(username string, destination_hostgroup uint64) (*QueryRules, error) {
	newqr := new(QueryRules)

	newqr.Username = username
	newqr.Destination_hostgroup = destination_hostgroup

	newqr.Schemaname = "NULL"
	newqr.FlagIN = 0
	newqr.Client_addr = "NULL"
	newqr.Proxy_addr = "NULL"
	newqr.Proxy_port = 0
	newqr.Digest = "NULL"
	newqr.Match_digest = "NULL"
	newqr.Match_pattern = "NULL"
	newqr.Negate_match_pattern = 0
	newqr.FlagOUT = 0
	newqr.Replace_pattern = "NULL"
	newqr.Cache_ttl = 0
	newqr.Reconnect = 0
	newqr.Timeout = 0
	newqr.Retries = 0
	newqr.Delay = 0
	newqr.Mirror_flagOUT = 0
	newqr.Mirror_hostgroup = 0
	newqr.Error_msg = "NULL"
	newqr.Log = 0
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
		qr.Schemaname = schema_name
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
		qr.Client_addr = client_addr
	}
}

// set qr proxy_addr
func (qr *QueryRules) SetQrProxyAddr(proxy_addr string) {
	if proxy_addr == "" || len(proxy_addr) == 0 {
		qr.Proxy_addr = "NULL"
	} else {
		qr.Proxy_addr = proxy_addr
	}
}

// set qr proxy_port
func (qr *QueryRules) SetProxyPort(proxy_port uint64) {
	qr.Proxy_port = proxy_port
}

// set qr digest
func (qr *QueryRules) SetQrDigest(digest string) {
	if digest == "" || len(digest) == 0 {
		qr.Digest = "NULL"
	} else {
		qr.Digest = digest
	}
}

// set qr match_digest
func (qr *QueryRules) SetQrMatchDigest(match_digest string) {
	if match_digest == "" || len(match_digest) == 0 {
		qr.Match_digest = "NULL"
	} else {
		qr.Match_digest = match_digest
	}
}

// set qr match_pattern
func (qr *QueryRules) SetQrMatchPattern(match_pattern string) {
	if match_pattern == "" || len(match_pattern) == 0 {
		qr.Match_pattern = "NULL"
	} else {
		qr.Match_pattern = match_pattern
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
func (qr *QueryRules) SetQrFlagOut(flag_out uint64) {
	qr.FlagOUT = flag_out
}

// set qr replace_pattern
func (qr *QueryRules) SetQrReplacePattern(replace_pattern string) {
	if replace_pattern == "" || len(replace_pattern) == 0 {
		qr.Replace_pattern = "NULL"
	} else {
		qr.Replace_pattern = replace_pattern
	}
}

// set qr destination_hostgroup
func (qr *QueryRules) SetQrDestHostGroup(destination_hostgroup uint64) {
	qr.Destination_hostgroup = destination_hostgroup
}

// set qr cache_ttl
func (qr *QueryRules) SetQrCacheTTL(cache_ttl uint64) {
	qr.Cache_ttl = cache_ttl
}

// set qr reconnect
func (qr *QueryRules) SetQrReconnect(reconnect uint64) {
	qr.Reconnect = reconnect
}

// set qr timeout
func (qr *QueryRules) SetQrTimeOut(timeout uint64) {
	qr.Timeout = timeout
}

// set qr retries
func (qr *QueryRules) SetQrRetries(retries uint64) {
	qr.Retries = retries
}

// set qr delay
func (qr *QueryRules) SetQrDelay(delay uint64) {
	qr.Delay = delay
}

// set qr mirror_flagout
func (qr *QueryRules) SetQrMirrorFlagOUT(mirror_flagout uint64) {
	qr.Mirror_flagOUT = mirror_flagout
}

// set qr mirror_hostgroup
func (qr *QueryRules) SetQrMirrorHostgroup(mirror_hostgroup uint64) {
	qr.Mirror_hostgroup = mirror_hostgroup
}

// set qr error_msg
func (qr *QueryRules) SetQrErrorMsg(error_msg string) {
	if error_msg == "" || len(error_msg) == 0 {
		qr.Error_msg = "NULL"
	} else {
		qr.Error_msg = error_msg
	}
}

// set qr log
func (qr *QueryRules) SetQrLog(log uint64) {
	qr.Log = log
}

// add a new query rules.
func (qr *QueryRules) AddOneQr(db *sql.DB) error {

	Query := fmt.Sprintf(StmtAddOneQr, qr.Username)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err) //add user failed
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
func (qr *QueryRules) UpdateOneQrInfo(db *sql.DB) error {

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
