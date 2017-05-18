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

type MySQLConnectionsPool struct {
	HostGroup     uint64 `db:"hostgroup" json:"hostgroup"`
	SrvHost       string `db:"srv_host" json:"srv_host"`
	SrvPort       uint64 `db:"srv_port" json:"srv_port"`
	Status        string `db:"status" json:"status"`
	ConnUsed      uint64 `db:"ConnUsed" json:"ConnUsed"`
	ConnFree      uint64 `db:"ConnFree" json:"ConnFree"`
	ConnOK        uint64 `db:"ConnOK" json:"ConnOK"`
	ConnERR       uint64 `db:"ConnERR" json:"ConnERR"`
	Queries       uint64 `db:"Queries" json:"Queries"`
	BytesDataSent uint64 `db:"Bytes_data_sent" json:"Bytes_data_sent"`
	BytesDataRecv uint64 `db:"Bytes_data_recv" json:"Bytes_data_recv"`
	LatencyUs     uint64 `db:"Latency_us" json:"Latency_us"`
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
	if ServerExists == 0 {
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
			log.Fatal("AddOneServers:", err)
		}
	}
	return 0
}

func (srvs *Servers) DeleteOneServers(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtDeleteOneServers, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("DeleteOneServers:", err)
		}
	}
}

func (srvs *Servers) ActiveOneServer(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtActiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("ActiveOneServer:", err)
		}
	} else {
		log.Fatal("ActiveOneServer,Server not exists")
	}
}

func (srvs *Servers) SoftDisactiveOneServer(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtSoftDisactiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("SoftDisactiveOneServer:", err)
		}
	} else {
		log.Fatal("SoftDisactiveOneServer,Server not exists")
	}
}

func (srvs *Servers) HardDisactiveOneServer(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtHardDisactiveOneServer, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("HardDisactiveOneServer:", err)
		}
	} else {
		log.Fatal("HardDisactiveOneServer,Server not exists")
	}
}

func (srvs *Servers) UpdateOneServerWeight(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerWeight, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("UpdateOneServerWeight:", err)
		}
	}
}

func (srvs *Servers) UpdateOneServerMc(db *sql.DB) {
	if isexist := srvs.ServerExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneServerMc, srvs.HostGroupId, srvs.HostName, srvs.Port)
		_, err := db.Query(st)
		if err != nil {
			log.Fatal("UpdateOneServerMc:", err)
		}
	}
}
