package servers

import (
	"database/sql"
	"fmt"
	"log"
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
	StmtAddOneServers          = `INSERT INTO mysql_servers(hostgroup_id,hostname,port) VALUES(%d,%q,%d)`
	StmtDeleteOneServers       = `DELETE FROM mysql_servers WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtActiveOneServer        = `UPDATE mysql_servers SET status='ONLINE' WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtSoftDisactiveOneServer = `UPDATE mysql_servers SET status='OFFLINE_SOFT' WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtHardDisactiveOneServer = `UPDATE mysql_servers SET status='OFFLINE_HARD' WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtUpdateOneServerWeight  = `UPDATE mysql_servers SET weight=%d WHERE hostgroup_id = %d AND hostname=%q AND port=%d`
	StmtUpdateOneServerMc      = `UPDATE mysql_servers SET max_connections=%d WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtServerExists           = `SELECT count(*) FROM mysql_servers WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtFindOneServer          = `SELECT ifnull(hostgroup_id,0) as hostgroup_id,ifnull(hostname,"") as hostname,ifnull(port,0) as port,ifnull(status,"") as status,ifnull(weight,0) as weight,ifnull(compression,0) as compression,ifnull(max_connections,0) as max_connections,ifnull(max_replication_lag,0) as max_replication_lag,ifnull(use_ssl,0) as use_ssl,ifnull(max_latency_ms,0) as max_latency_ms,ifnull(comment,"") as comment FROM mysql_servers WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtFindAllServer          = `SELECT ifnull(hostgroup_id,0) as hostgroup_id,ifnull(hostname,"") as hostname,ifnull(port,0) as port,ifnull(status,"") as status,ifnull(weight,0) as weight,ifnull(compression,0) as compression,ifnull(max_connections,0) as max_connections,ifnull(max_replication_lag,0) as max_replication_lag,ifnull(use_ssl,0) as use_ssl,ifnull(max_latency_ms,0) as max_latency_ms,ifnull(comment,"") as comment FROM mysql_servers limit %d offset %d`
	StmtFindServersByHostgroup = `SELECT ifnull(hostgroup_id,0) as hostgroup_id,ifnull(hostname,"") as hostname,ifnull(port,0) as port,ifnull(status,"") as status,ifnull(weight,0) as weight,ifnull(compression,0) as compression,ifnull(max_connections,0) as max_connections,ifnull(max_replication_lag,0) as max_replication_lag,ifnull(use_ssl,0) as use_ssl,ifnull(max_latency_ms,0) as max_latency_ms,ifnull(comment,"") as comment FROM mysql_servers WHERE hostgroup_id=%d`
)

func (srvs *Servers) ServerExists(db *sql.DB) bool {
	st := fmt.Sprintf(StmtServerExists, srvs.HostGroupId, srvs.HostName, srvs.Port)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("ServerExists:", err)
	}
	var ServerCount uint64
	for rows.Next() {
		err = rows.Scan(&ServerCount)
		if err != nil {
			log.Print("ServerExists:", err)
		}
	}
	if ServerCount == 0 {
		return false
	} else {
		return true
	}
}

func (srvs *Servers) AddOneServers(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == false {
		st := fmt.Sprintf(StmtAddOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) DeleteOneServers(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtDeleteOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) ActiveOneServer(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtActiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) SoftDisactiveOneServer(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtSoftDisactiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) HardDisactiveOneServer(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtHardDisactiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) UpdateOneServerWeight(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerWeight, srvs.Weight, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) UpdateOneServerMc(db *sql.DB) int {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerMc, srvs.MaxConnections, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) FindOneServersInfo(db *sql.DB) Servers {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtFindOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		rows, err := db.Query(st)
		if err != nil {
			log.Print("FindOneServerInfo:", err)
		}
		for rows.Next() {
			err = rows.Scan(
				&srvs.HostGroupId,
				&srvs.HostName,
				&srvs.Port,
				&srvs.Status,
				&srvs.Weight,
				&srvs.Compression,
				&srvs.MaxConnections,
				&srvs.MaxReplicationLag,
				&srvs.UseSsl,
				&srvs.MaxLatencyMs,
				&srvs.Comment,
			)
		}
	} else {
		log.Print("FindOneServerInfo: Server not exists")
	}
	return *srvs
}

func FindAllServerInfo(db *sql.DB, limit int64, skip int64) []Servers {
	var allserver []Servers
	var tmpserver Servers
	st := fmt.Sprintf(StmtFindAllServer, limit, skip)
	rows, err := db.Query(st)
	log.Printf(StmtFindAllServer, limit, skip)
	if err != nil {
		log.Print("FindAllServerInfo:", err)
	}

	defer rows.Close()

	for rows.Next() {
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
		allserver = append(allserver, tmpserver)
	}

	return allserver
}

func (srvs *Servers) FindServersInfoByHostgroup(db *sql.DB) []Servers {
	var allserver []Servers
	var tmpserver Servers

	st := fmt.Sprintf(StmtFindServersByHostgroup, srvs.HostGroupId)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("FindServersInfoByHostgroup:", err)
	}

	defer rows.Close()

	for rows.Next() {
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
		allserver = append(allserver, tmpserver)
	}

	return allserver
}
