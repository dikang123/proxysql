package cmd

import (
        "database/sql"
        "fmt"
        "log"
        "os"
)

const (
    CmdProxyReadOnly = `PROXYSQL READONLY`
    CmdProxyReadWrite = `PROXYSQL READWRITE`
    CmdProxyStart = `PROXYSQL START`
    CmdProxyRestart = `PROXYSQL RESTART`
    CmdProxyStop = `PROXYSQL STOP`
    CmdProxyPause = `PROXYSQL PAUSE`
    CmdProxyResume = `PROXYSQL RESUME`
    CmdProxyShutdown = `PROXYSQL SHUTDOWN`
    CmdProxyFlushLogs= `PROXYSQL FLUSH LOGS`
    CmdProxyKill = `PROXYSQL KILL`
    CmdLoadUserToRuntime = `LOAD MYSQL USERS TO RUNTIME`
    CmdSaveUserToDisk = `SAVE MYSQL USERS TO DISK`
    CmdLoadServerToRuntime = `LOAD MYSQL SERVERS TO RUNTIME`
    CmdSaveServerToDisk = `SAVE MYSQL SERVERS TO DISK`
    CmdLoadQueryRulesToRuntime = `LOAD MYSQL QUERY RULES TO RUNTIME`
    CmdSaveQueryRulesToDisk = `SAVE MYSQL QUERY RULES TO DISK`
    CmdLoadSchedulerToRuntime = `LOAD SCHEDULER TO RUNTIME`
    CmdSaveSchedulerToDisk = `SAVE SCHEDULER TO DISK`
)
