package cmd

import (
	"database/sql"
	"log"
)

const (
	CmdProxyReadOnly               = `PROXYSQL READONLY`
	CmdProxyReadWrite              = `PROXYSQL READWRITE`
	CmdProxyStart                  = `PROXYSQL START`
	CmdProxyRestart                = `PROXYSQL RESTART`
	CmdProxyStop                   = `PROXYSQL STOP`
	CmdProxyPause                  = `PROXYSQL PAUSE`
	CmdProxyResume                 = `PROXYSQL RESUME`
	CmdProxyShutdown               = `PROXYSQL SHUTDOWN`
	CmdProxyFlushLogs              = `PROXYSQL FLUSH LOGS`
	CmdProxyKill                   = `PROXYSQL KILL`
	CmdLoadUserToRuntime           = `LOAD MYSQL USERS TO RUNTIME`
	CmdSaveUserToDisk              = `SAVE MYSQL USERS TO DISK`
	CmdLoadServerToRuntime         = `LOAD MYSQL SERVERS TO RUNTIME`
	CmdSaveServerToDisk            = `SAVE MYSQL SERVERS TO DISK`
	CmdLoadQueryRulesToRuntime     = `LOAD MYSQL QUERY RULES TO RUNTIME`
	CmdSaveQueryRulesToDisk        = `SAVE MYSQL QUERY RULES TO DISK`
	CmdLoadSchedulerToRuntime      = `LOAD SCHEDULER TO RUNTIME`
	CmdSaveSchedulerToDisk         = `SAVE SCHEDULER TO DISK`
	CmdLoadMySQLVariablesToRuntime = `LOAD MYSQL VARIABLES TO RUNTIME`
	CmdSaveMySQLVariablesToDisk    = `SAVE MYSQL VARIABLES TO DISK`
	CmdLoadAdminVariablesToRuntime = `LOAD ADMIN VARIABLES TO RUNTIME`
	CmdSaveAdminVariablesToDisk    = `SAVE ADMIN VARIABLES TO DISK`
)

func ProxyReadOnly(db *sql.DB) (int, error) {
	log.Print("ProxyReadOnly: ", CmdProxyReadOnly)
	_, err := db.Query(CmdProxyReadOnly)
	if err != nil {
		log.Print("ProxyReadOnly:", err)
		return 1, err
	}
	log.Print("ProxyReadOnly: Success")
	return 0, nil
}

func ProxyReadWrite(db *sql.DB) (int, error) {
	log.Print("ProxyReadWrite: ", CmdProxyReadWrite)
	_, err := db.Query(CmdProxyReadWrite)
	if err != nil {
		log.Print("ProxyReadWrite:", err)
		return 1, err
	}
	log.Print("ProxyReadWrite: Success")
	return 0, nil
}

func ProxyStart(db *sql.DB) (int, error) {
	log.Print("ProxyStart: ", CmdProxyStart)
	_, err := db.Query(CmdProxyStart)
	if err != nil {
		log.Print("ProxyStart:", err)
		return 1, err
	}
	log.Print("ProxyStart: Success")
	return 0, nil
}

func ProxyRestart(db *sql.DB) (int, error) {
	log.Print("ProxyRestart: ", CmdProxyRestart)
	_, err := db.Query(CmdProxyRestart)
	if err != nil {
		log.Print("ProxyRestart:", err)
		return 1, err
	}
	log.Print("ProxyRestart: Success")
	return 0, nil
}

func ProxyStop(db *sql.DB) (int, error) {
	log.Print("ProxyStop: ", CmdProxyStop)
	_, err := db.Query(CmdProxyStop)
	if err != nil {
		log.Print("ProxyStop:", err)
		return 1, err
	}
	log.Print("ProxyStop: Success")
	return 0, nil
}

func ProxyPause(db *sql.DB) (int, error) {
	log.Print("ProxyPause: ", CmdProxyStop)
	_, err := db.Query(CmdProxyPause)
	if err != nil {
		log.Print("ProxyPause:", err)
		return 1, err
	}
	log.Print("ProxyPause: Success")
	return 0, nil
}

func ProxyResume(db *sql.DB) (int, error) {
	log.Print("ProxyResume: ", CmdProxyResume)
	_, err := db.Query(CmdProxyResume)
	if err != nil {
		log.Print("ProxyResume:", err)
		return 1, err
	}
	log.Print("ProxyResume: Success")
	return 0, nil
}

func ProxyShutdown(db *sql.DB) (int, error) {
	log.Print("ProxyShutdown: ", CmdProxyShutdown)
	_, err := db.Query(CmdProxyShutdown)
	if err != nil {
		log.Print("ProxyShutdown:", err)
		return 1, err
	}
	log.Print("ProxyShutdown: Success")
	return 0, nil
}

func ProxyFlushLogs(db *sql.DB) (int, error) {
	log.Print("ProxyFlushLogs: ", CmdProxyFlushLogs)
	_, err := db.Query(CmdProxyFlushLogs)
	if err != nil {
		log.Print("ProxyFlushLogs:", err)
		return 1, err
	}
	log.Print("ProxyFlushLogs: Success")
	return 0, nil
}

func ProxyKill(db *sql.DB) (int, error) {
	log.Print("ProxyKill: ", CmdProxyKill)
	_, err := db.Query(CmdProxyKill)
	if err != nil {
		log.Print("ProxyKill:", err)
		return 1, err
	}
	log.Print("ProxyKill: Success")
	return 0, nil
}

func LoadUserToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadUserToRuntime: ", CmdLoadUserToRuntime)
	_, err := db.Query(CmdLoadUserToRuntime)
	if err != nil {
		log.Print("LoadUserToRuntime:", err)
		return 1, err
	}
	log.Print("LoadUserToRuntime: Success")
	return 0, nil
}

func SaveUserToDisk(db *sql.DB) (int, error) {
	log.Print("SaveUserToDisk: ", CmdSaveUserToDisk)
	_, err := db.Query(CmdSaveUserToDisk)
	if err != nil {
		log.Print("SaveUserToDisk", err)
		return 1, err
	}
	log.Print("SaveUserToDisk: Success")
	return 0, nil
}

func LoadServerToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadServerToRuntime: ", CmdLoadServerToRuntime)
	_, err := db.Query(CmdLoadServerToRuntime)
	if err != nil {
		log.Print("LoadServerToRuntime:", err)
		return 1, err
	}
	log.Print("LoadServerToRuntime: Success")
	return 0, nil
}

func SaveServerToDisk(db *sql.DB) (int, error) {
	log.Print("SaveServerToDisk: ", CmdSaveServerToDisk)
	_, err := db.Query(CmdSaveServerToDisk)
	if err != nil {
		log.Print("SaveServerToDisk:", err)
		return 1, err
	}
	log.Print("SaveServerToDisk: Success")
	return 0, nil
}

func LoadQueryRulesToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadQueryRulesToRuntime: ", CmdLoadQueryRulesToRuntime)
	_, err := db.Query(CmdLoadQueryRulesToRuntime)
	if err != nil {
		log.Print("LoadQueryRulesToRuntime:", err)
		return 1, err
	}
	log.Print("LoadQueryRulesToRuntime: Success")
	return 0, nil
}

func SaveQueryRulesToDisk(db *sql.DB) (int, error) {
	log.Print("SaveQueryRulesToDisk: ", CmdSaveQueryRulesToDisk)
	_, err := db.Query(CmdSaveQueryRulesToDisk)
	if err != nil {
		log.Print("SaveQueryRulesToDisk:", err)
		return 1, err
	}
	log.Print("SaveQueryRulesToDisk: Success")
	return 0, nil
}

func LoadSchedulerToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadSchedulerToRuntime: ", CmdLoadSchedulerToRuntime)
	_, err := db.Query(CmdLoadSchedulerToRuntime)
	if err != nil {
		log.Print("LoadSchedulerToRuntime:", err)
		return 1, err
	}
	log.Print("LoadSchedulerToRuntime: Success")
	return 0, nil
}

func SaveSchedulerToDisk(db *sql.DB) (int, error) {
	log.Print("SaveSchedulerToDisk: ", CmdSaveSchedulerToDisk)
	_, err := db.Query(CmdSaveSchedulerToDisk)
	if err != nil {
		log.Print("SaveSchedulerToDisk:", err)
		return 1, err
	}
	log.Print("SaveSchedulerToDisk: Success")
	return 0, nil
}

func LoadMySQlVariablesToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadMySQLVariablesToRunTime: ", CmdLoadMySQLVariablesToRuntime)
	_, err := db.Query(CmdLoadMySQLVariablesToRuntime)
	if err != nil {
		log.Print("LoadMySQlVariablesToRunTime: ", err)
		return 1, err
	}
	log.Print("LoadMySQlVariablesToRunTime: Success")
	return 0, nil
}

func LoadAdminVariablesToRuntime(db *sql.DB) (int, error) {
	log.Print("LoadAdminVariablesToRunTime: ", CmdLoadAdminVariablesToRuntime)
	_, err := db.Query(CmdLoadAdminVariablesToRuntime)
	if err != nil {
		log.Print("LoadAdminVariablesToRunTime: ", err)
		return 1, err
	}
	log.Print("LoadAdminVariablesToRunTime: Success")
	return 0, nil
}

func SaveMySQLVariablesToDisk(db *sql.DB) (int, error) {
	log.Print("SaveMySQLVariablesToDisk: ", CmdSaveMySQLVariablesToDisk)
	_, err := db.Query(CmdSaveMySQLVariablesToDisk)
	if err != nil {
		log.Print("SaveMySQLVariablesToDisk:", err)
		return 1, err
	}
	log.Print("SaveMySQLVariablesToDisk: Success")
	return 0, nil
}

func SaveAdminVariablesToDisk(db *sql.DB) (int, error) {
	log.Print("SaveAdminVariablesToDisk: ", CmdSaveAdminVariablesToDisk)
	_, err := db.Query(CmdSaveAdminVariablesToDisk)
	if err != nil {
		log.Print("SaveAdminVariablesToDisk:", err)
		return 1, err
	}
	log.Print("SaveAdminVariablesToDisk: Success")
	return 0, nil
}
