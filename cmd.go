package proxysql

// proxysql admin commands

import (
	"database/sql"

	"github.com/juju/errors"
)

// define commands
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

//set proxysql to readonly mode.
func ProxyReadOnly(db *sql.DB) error {
	_, err := db.Exec(CmdProxyReadOnly)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// set proxysql to readwrite mode.
func ProxyReadWrite(db *sql.DB) error {
	_, err := db.Exec(CmdProxyReadWrite)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// start proxysql child process.
func ProxyStart(db *sql.DB) error {
	_, err := db.Exec(CmdProxyStart)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// restart proxysql process.
func ProxyRestart(db *sql.DB) error {
	_, err := db.Exec(CmdProxyRestart)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// stop proxysql child process.
func ProxyStop(db *sql.DB) error {
	_, err := db.Exec(CmdProxyStop)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// pause proxysql
func ProxyPause(db *sql.DB) error {
	_, err := db.Exec(CmdProxyPause)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// resume proxysql
func ProxyResume(db *sql.DB) error {
	_, err := db.Exec(CmdProxyResume)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// shutdown proxysql
func ProxyShutdown(db *sql.DB) error {
	_, err := db.Exec(CmdProxyShutdown)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// flush proxysql logs to file
func ProxyFlushLogs(db *sql.DB) error {
	_, err := db.Exec(CmdProxyFlushLogs)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// kill child process.
func ProxyKill(db *sql.DB) error {
	_, err := db.Exec(CmdProxyKill)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

//execute load mysql users to runtime.
func LoadUserToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadUserToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

//execute save mysql users to disk.
func SaveUserToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveUserToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute load mysql servers to runtime.
func LoadServerToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadServerToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute save mysql servers to disk.
func SaveServerToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveServerToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute load mysql query rules to runtime.
func LoadQueryRulesToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadQueryRulesToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute save mysql query rules to disk.
func SaveQueryRulesToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveQueryRulesToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute load schedulers to runtime.
func LoadSchedulerToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadSchedulerToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute save schedulers to disk.
func SaveSchedulerToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveSchedulerToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute  load mysql variables to runtime.
func LoadMySQlVariablesToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadMySQLVariablesToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute load admin variables to runtime.
func LoadAdminVariablesToRuntime(db *sql.DB) error {
	_, err := db.Exec(CmdLoadAdminVariablesToRuntime)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute save mysql variables to runtime.
func SaveMySQLVariablesToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveMySQLVariablesToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// execute save admin variables to disk.
func SaveAdminVariablesToDisk(db *sql.DB) error {
	_, err := db.Exec(CmdSaveAdminVariablesToDisk)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
