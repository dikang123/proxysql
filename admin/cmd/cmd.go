package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	CmdProxyReadOnly           = `PROXYSQL READONLY`
	CmdProxyReadWrite          = `PROXYSQL READWRITE`
	CmdProxyStart              = `PROXYSQL START`
	CmdProxyRestart            = `PROXYSQL RESTART`
	CmdProxyStop               = `PROXYSQL STOP`
	CmdProxyPause              = `PROXYSQL PAUSE`
	CmdProxyResume             = `PROXYSQL RESUME`
	CmdProxyShutdown           = `PROXYSQL SHUTDOWN`
	CmdProxyFlushLogs          = `PROXYSQL FLUSH LOGS`
	CmdProxyKill               = `PROXYSQL KILL`
	CmdLoadUserToRuntime       = `LOAD MYSQL USERS TO RUNTIME`
	CmdSaveUserToDisk          = `SAVE MYSQL USERS TO DISK`
	CmdLoadServerToRuntime     = `LOAD MYSQL SERVERS TO RUNTIME`
	CmdSaveServerToDisk        = `SAVE MYSQL SERVERS TO DISK`
	CmdLoadQueryRulesToRuntime = `LOAD MYSQL QUERY RULES TO RUNTIME`
	CmdSaveQueryRulesToDisk    = `SAVE MYSQL QUERY RULES TO DISK`
	CmdLoadSchedulerToRuntime  = `LOAD SCHEDULER TO RUNTIME`
	CmdSaveSchedulerToDisk     = `SAVE SCHEDULER TO DISK`
)

func ProxyReadOnly(db *sql.DB) {
	_, err := db.Query(CmdProxyReadOnly)
	if err != nil {
		log.Fatal("ProxyReadOnly:", err)
	}
}

func ProxyReadWrite(db *sql.DB) {
	_, err := db.Query(CmdProxyReadWrite)
	if err != nil {
		log.Fatal("ProxyReadWrite:", err)
	}
}

func ProxyStart(db *sql.DB) {
	_, err := db.Query(CmdProxyStart)
	if err != nil {
		log.Fatal("ProxyStart:", err)
	}
}

func ProxyRestart(db *sql.DB) {
	_, err := db.Query(CmdProxyRestart)
	if err != nil {
		log.Fatal("ProxyRestart:", err)
	}
}

func ProxyStop(db *sql.DB) {
	_, err := db.Query(CmdProxyStop)
	if err != nil {
		log.Fatal("ProxyStop:", err)
	}
}

func ProxyPause(db *sql.DB) {
	_, err := db.Query(CmdProxyPause)
	if err != nil {
		log.Fatal("ProxyPause:", err)
	}
}

func ProxyResume(db *sql.DB) {
	_, err := db.Query(CmdProxyResume)
	if err != nil {
		log.Fatal("ProxyResume:", err)
	}
}

func ProxyShutdown(db *sql.DB) {
	_, err := db.Query(CmdProxyShutdown)
	if err != nil {
		log.Fatal("ProxyShutdown:", err)
	}
}

func ProxyFlushLogs(db *sql.DB) {
	_, err := db.Query(CmdProxyFlushLogs)
	if err != nil {
		log.Fatal("ProxyFlushLogs:", err)
	}
}

func ProxyKill(db *sql.DB) {
	_, err := db.Query(CmdProxyKill)
	if err != nil {
		log.Fatal("ProxyKill:", err)
	}
}

func LoadUserToRuntime(db *sql.DB) {
	_, err := db.Query(CmdLoadUserToRuntime)
	if err != nil {
		log.Fatal("LoadUserToRuntime:", err)
	}
}

func SaveUserToDisk(db *sql.DB) {
	_, err := db.Query(CmdSaveUsertoDisk)
	if err != nil {
		log.Fatal("SaveUserToDisk", err)
	}
}

func LoadServerToRuntime(db *sql.DB) {
	_, err := db.Query(CmdLoadServerToRuntime)
	if err != nil {
		log.Fatal("LoadServerToRuntime:", err)
	}
}

func SaveServerToDisk(db *sql.DB) {
	_, err := db.Query(CmdSaveServerToDisk)
	if err != nil {
		log.Fatal("SaveServerToDisk:", err)
	}
}

func LoadQueryRulesToRuntime(db *sql.DB) {
	_, err := db.Query(CmdLoadQueryRulesToRuntime)
	if err != nil {
		log.Fatal("LoadQueryRulesToRuntime:", err)
	}
}

func SaveQueryRulesToDisk(db *sql.DB) {
	_, err := db.Query(CmdSaveQueryRulesToDisk)
	if err != nil {
		log.Fatal("SaveQueryRulesToDisk:", err)
	}
}

func LoadSchedulerToRuntime(db *sql.DB) {
	_, err := db.Query(CmdLoadSchedulerToRuntime)
	if err != nil {
		log.Fatal("LoadSchedulerToRuntime:", err)
	}
}

func SaveSchedulerToDisk(db *sql.DB) {
	_, err := db.Query(CmdSaveSchedulerToDisk)
	if err != nil {
		log.Fatal("SaveSchedulerToDisk:", err)
	}
}
