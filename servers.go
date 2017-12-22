package proxysql

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
)

type Servers struct {
	HostGroupId       uint64 `db:"hostgroup_id,omitempty" json:"hostgroup_id"`
	HostName          string `db:"hostname" json:"hostname"`
	Port              uint64 `db:"port" json:"port"`
	Status            string `db:"status" json:"status"`
	Weight            uint64 `db:"weight" json:"weight"`
	Compression       uint64 `db:"compression" json:"compression"`
	MaxConnections    uint64 `db:"max_connections" json:"max_connections"`
	MaxReplicationLag uint64 `db:"max_replication_lag" json:"max_replication_lag"`
	UseSsl            uint64 `db:"use_ssl" json:"use_ssl"`
	MaxLatencyMs      uint64 `db:"max_latency_ms" json:"max_latency_ms"`
	Comment           string `db:"comment" json:"comment"`
}

const (
	/*add a new backends.*/
	StmtAddOneServers = `
	INSERT 
	INTO 
		mysql_servers(
			hostgroup_id,
			hostname,
			port
		) 
	VALUES(%d,%q,%d)`

	/*delete a backend*/
	StmtDeleteOneServers = `
	DELETE 
	FROM 
		mysql_servers 
	WHERE 
		hostgroup_id=%d 
	AND hostname=%q 
	AND port=%d`

	/*update a backends*/
	StmtUpdateOneServer = `
	UPDATE 
		mysql_servers 
	SET 
		status=%q,
		weight=%d,
		compression=%d,
		max_connections=%d,
		max_replication_lag=%d,
		use_ssl=%d,
		max_latency_ms=%d,
		comment=%q 
	WHERE 
		hostgroup_id=%d 
	AND hostname=%q 
	AND port=%d`

	/*list all mysql_servers*/
	StmtFindAllServer = `
	SELECT 
		ifnull(hostgroup_id,0) as hostgroup_id,
		ifnull(hostname,"") as hostname,
		ifnull(port,0) as port,
		ifnull(status,"") as status,
		ifnull(weight,0) as weight,
		ifnull(compression,0) as compression,
		ifnull(max_connections,0) as max_connections,
		ifnull(max_replication_lag,0) as max_replication_lag,
		ifnull(use_ssl,0) as use_ssl,
		ifnull(max_latency_ms,0) as max_latency_ms,
		ifnull(comment,"") as comment 
	FROM 
		mysql_servers 
	LIMIT %d 
	OFFSET %d`
)

/*list all mysql_servers*/
func (srvs *Servers) FindAllServerInfo(db *sql.DB, limit int64, skip int64) ([]Servers, error) {

	var allserver []Servers

	Query := fmt.Sprintf(StmtFindAllServer, limit, skip)

	rows, err := db.Query(Query)
	if err != nil {
		return []Servers{}, errors.Trace(err)
	}
	defer rows.Close()

	for rows.Next() {

		var tmpserver Servers

		err = rows.Scan(
			&tmpserver.HostGroupId,
			&tmpserver.HostName,
			&tmpserver.Port,
			&tmpserver.Status,
			&tmpserver.Weight,
			&tmpserver.Compression,
			&tmpserver.MaxConnections,
			&tmpserver.MaxReplicationLag,
			&tmpserver.UseSsl,
			&tmpserver.MaxLatencyMs,
			&tmpserver.Comment,
		)

		if err != nil {
			continue
		}

		allserver = append(allserver, tmpserver)
	}

	return allserver, nil
}

/*add a new backend*/
func (srvs *Servers) AddOneServers(db *sql.DB) error {

	Query := fmt.Sprintf(StmtAddOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	LoadServerToRuntime(db)
	SaveServerToDisk(db)

	return nil
}

/*delete a backend*/
func (srvs *Servers) DeleteOneServers(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	LoadServerToRuntime(db)
	SaveServerToDisk(db)

	return nil
}

//更新后端服务全部信息
func (srvs *Servers) UpdateOneServerInfo(db *sql.DB) error {

	Query := fmt.Sprintf(StmtUpdateOneServer,
		srvs.Status,
		srvs.Weight,
		srvs.Compression,
		srvs.MaxConnections,
		srvs.MaxReplicationLag,
		srvs.UseSsl,
		srvs.MaxLatencyMs,
		srvs.Comment,
		srvs.HostGroupId,
		srvs.HostName,
		srvs.Port)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	LoadServerToRuntime(db)
	SaveServerToDisk(db)

	return nil
}
