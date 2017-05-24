package servers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Servers struct {
	HostGroupId       uint64 `db:"hostgroup_id,omitempty" json:"hostname_id"`
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
	StmtSoftDisactiveOneServer = `UPDATE mysql_servers SET status='SOFT_OFFLINE' WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtHardDisactiveOneServer = `UPDATE mysql_servers SET status='HARD_OFFLINE' WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtUpdateOneServerWeight  = `UPDATE mysql_servers SET weight=%d WHERE hostgroup_id = %d AND hostname=%q AND port=%d`
	StmtUpdateOneServerMc      = `UPDATE mysql_servers SET max_connections=%d WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtServerExists           = `SELECT count(*) FROM mysql_servers WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtFindOneServer          = `SELECT * FROM mysql_servers WHERE hostgroup_id=%d AND hostname=%q AND port=%d`
	StmtFindAllServer          = `SELECT * FROM mysql_servers`
)

func (srvs *Servers) ServerExists(db *sql.DB) bool {
	st := fmt.Sprintf(StmtServerExists, srvs.HostGroupId, srvs.HostName, srvs.Port)
	rows, err := db.Query(st)
	if err != nil {
		log.Fatal("ServerExists:", err)
	}
	var ServerCount uint64
	for rows.Next() {
		err = rows.Scan(&ServerCount)
		if err != nil {
			log.Fatal("ServerExists:", err)
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
	}
	return 2
}

func (srvs *Servers) DeleteOneServers(db *sql.DB) {
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

func (srvs *Servers) ActiveOneServer(db *sql.DB) {
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

func (srvs *Servers) SoftDisactiveOneServer(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtSoftDisactiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) HardDisactiveOneServer(db *sql.DB) {
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

func (srvs *Servers) UpdateOneServerWeight(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerWeight, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			return 1
		}
		return 0
	} else {
		return 2
	}
}

func (srvs *Servers) UpdateOneServerMc(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerMc, srvs.HostGroupId, srvs.HostName, srvs.Port)
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
			return 1
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
		return 2
	}
	return srvs
}

func FindAllServerInfo(db *sql.DB) []Servers {
	var allserver []Servers
	var tmpserver Servers

	rows, err := db.Query(StmtFindAllServer)
	if err != nil {
		return 1
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
