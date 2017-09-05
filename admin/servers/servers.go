package servers

import (
	"database/sql"
	"fmt"
	"log"
	"proxysql-master/admin/cmd"
	//"os"
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
	/*新建一个后端服务*/
	StmtAddOneServers = `
	INSERT 
	INTO 
		mysql_servers(
			hostgroup_id,
			hostname,
			port
		) 
	VALUES(%d,%q,%d)`

	/*删除一个后端服务*/
	StmtDeleteOneServers = `
	DELETE 
	FROM 
		mysql_servers 
	WHERE 
		hostgroup_id=%d 
	AND hostname=%q 
	AND port=%d`

	/*更新一个后端服务*/
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

	/*查询出所有后端服务信息*/
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

/*查询所有后端服务的信息*/
func (srvs *Servers) FindAllServerInfo(db *sql.DB, limit int64, skip int64) ([]Servers, error) {

	/*定义一个新的变量，保存所有后端服务信息*/
	var allserver []Servers

	Query := fmt.Sprintf(StmtFindAllServer, limit, skip)
	log.Print("admin->servers.go->FindAllServerInfo->Query: ", Query)

	rows, err := db.Query(Query)
	if err != nil {
		log.Print("admin->servers.go->FindAllServerInfo Failed:", err)
		return []Servers{}, err
	}
	defer rows.Close()

	/*得出查询结果*/
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
			log.Print("admin->servers.go->FindAllServerInfo-rows.Scan Failed:", err)
			continue
		}

		log.Print("admin->servers.go->FindAllServerInfo-tmpserver: ", tmpserver)
		allserver = append(allserver, tmpserver)
	}

	return allserver, nil
}

func (srvs *Servers) AddOneServers(db *sql.DB) int {
	st := fmt.Sprintf(StmtAddOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)
	_, err := db.Query(st)
	if err != nil {
		return 1
	}
	cmd.LoadServerToRuntime(db)
	cmd.SaveServerToDisk(db)
	return 0
}

func (srvs *Servers) DeleteOneServers(db *sql.DB) int {
	st := fmt.Sprintf(StmtDeleteOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)
	_, err := db.Query(st)
	if err != nil {
		return 1
	}
	cmd.LoadServerToRuntime(db)
	cmd.SaveServerToDisk(db)
	return 0
}

//更新后端服务全部信息
func (srvs *Servers) UpdateOneServerInfo(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneServer, srvs.Status, srvs.Weight, srvs.Compression, srvs.MaxConnections, srvs.MaxReplicationLag, srvs.UseSsl, srvs.MaxLatencyMs, srvs.Comment, srvs.HostGroupId, srvs.HostName, srvs.Port)
	log.Print("servers->UpdateOneServerInfo->st: ", st)
	_, err := db.Query(st)
	if err != nil {
		return 1
	}
	cmd.LoadServerToRuntime(db)
	cmd.SaveServerToDisk(db)
	return 0
}
