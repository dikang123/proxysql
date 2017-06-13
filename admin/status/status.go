package status

import (
	"database/sql"
	//"fmt"
	"log"
)

type (
	PsStatus struct {
		Active_Transactions          int64 `db:"Active_Transactions" json:"Active_Transactions"`
		Backend_query_time_nsec      int64 `db:"Backend_query_time_nsec" json:"Backend_query_time_nsec"`
		Client_Connections_aborted   int64 `db:"Client_Connections_aborted" json:"Client_Connections_aborted"`
		Client_Connections_connected int64 `db:"Client_Connections_connected" json:"Client_Connections_connected"`
		Client_Connections_created   int64 `db:"Client_Connections_created" json:"Client_Connections_created"`
		Client_Connections_non_idle  int64 `db:"Client_Connections_non_idle" json:"Client_Connections_non_idle"`
		Com_autocommit               int64 `db:"Com_autocommit" json:"Com_autocommit"`
		Com_autocommit_filtered      int64 `db:"Com_autocommit_filtered" json:"Com_autocommit_filtered"`
		Com_commit                   int64 `db:"Com_commit" json:"Com_commit"`
		Com_commit_filtered          int64 `db:"Com_commit_filtered" json:"Com_commit_filtered"`
		Com_rollback                 int64 `db:"Com_rollback" json:"Com_rollback"`
		Com_rollback_filtered        int64 `db:"Com_rollback_filtered" json:"Com_rollback_filtered"`
		Com_stmt_close               int64 `db:"Com_stmt_close" json:"Com_stmt_close"`
		Com_stmt_execute             int64 `db:"Com_stmt_execute" json:"Com_stmt_execute"`
		Com_stmt_prepare             int64 `db:"Com_stmt_prepare" json:"Com_stmt_prepare"`
		ConnPool_get_conn_failure    int64 `db:"ConnPool_get_conn_failure" json:"ConnPool_get_conn_failure"`
		ConnPool_get_conn_immediate  int64 `db:"ConnPool_get_conn_immediate" json:"ConnPool_get_conn_immediate"`
		ConnPool_get_conn_success    int64 `db:"ConnPool_get_conn_success" json:"ConnPool_get_conn_success"`
		ConnPool_memory_bytes        int64 `db:"ConnPool_memory_bytes" json:"ConnPool_memory_bytes"`
		MySQL_Monitor_Workers        int64 `db:"MySQL_Monitor_Workers" json:"MySQL_Monitor_Workers"`
		MySQL_Thread_Workers         int64 `db:"MySQL_Thread_Workers" json:"MySQL_Thread_Workers"`
		ProxySQL_Uptime              int64 `db:"ProxySQL_Uptime" json:"ProxySQL_Uptime"`
		Queries_backends_bytes_recv  int64 `db:"Queries_backends_bytes_recv" json:"Queries_backends_bytes_recv"`
		Queries_backends_bytes_sent  int64 `db:"Queries_backends_bytes_sent" json:"Queries_backends_bytes_sent"`
		Query_Cache_Entries          int64 `db:"Query_Cache_Entries" json:"Query_Cache_Entries"`
		Query_Cache_Memory_bytes     int64 `db:"Query_Cache_Memory_bytes" json:"Query_Cache_Memory_bytes"`
		Query_Cache_Purged           int64 `db:"Query_Cache_Purged" json:"Query_Cache_Purged"`
		Query_Cache_bytes_IN         int64 `db:"Query_Cache_bytes_IN" json:"Query_Cache_bytes_IN"`
		Query_Cache_bytes_OUT        int64 `db:"Query_Cache_bytes_OUT" json:"Query_Cache_bytes_OUT"`
		Query_Cache_count_GET        int64 `db:"Query_Cache_count_GET" json:"Query_Cache_count_GET"`
		Query_Cache_count_GET_OK     int64 `db:"Query_Cache_count_GET_OK" json:"Query_Cache_count_GET_OK"`
		Query_Cache_count_SET        int64 `db:"Query_Cache_count_SET" json:"Query_Cache_count_SET"`
		Query_Processor_time_nsec    int64 `db:"Query_Processor_time_nsec" json:"Query_Processor_time_nsec"`
		Questions                    int64 `db:"Questions" json:"Questions"`
		SQLite3_memory_bytes         int64 `db:"SQLite3_memory_bytes" json:"SQLite3_memory_bytes"`
		Server_Connections_aborted   int64 `db:"Server_Connections_aborted" json:"Server_Connections_aborted"`
		Server_Connections_connected int64 `db:"Server_Connections_connected" json:"Server_Connections_connected"`
		Server_Connections_created   int64 `db:"Server_Connections_created" json:"Server_Connections_created"`
		Servers_table_version        int64 `db:"Servers_table_version" json:"Servers_table_version"`
		Slow_queries                 int64 `db:"Slow_queries" json:"Slow_queries"`
		Stmt_Active_Total            int64 `db:"Stmt_Active_Total" json:"Stmt_Active_Total"`
		Stmt_Active_Unique           int64 `db:"Stmt_Active_Unique" json:"Stmt_Active_Unique"`
		Stmt_Max_Stmt_id             int64 `db:"Stmt_Max_Stmt_id" json:"Stmt_Max_Stmt_id"`
		Mysql_backend_buffers_bytes  int64 `db:"mysql_backend_buffers_bytes" json:"mysql_backend_buffers_bytes"`
		Mysql_frontend_buffers_bytes int64 `db:"mysql_frontend_buffers_bytes" json:"mysql_frontend_buffers_bytes"`
		Mysql_session_internal_bytes int64 `db:"mysql_session_internal_bytes" json:"mysql_session_internal_bytes"`
	}
	Variables struct {
		VariablesName string `db:"Variable_name" json:"Variable_name"`
		Value         int64  `db:"Value" json:"Value"`
	}
)

const (
	StmtMySQLStatus = `SHOW MYSQL STATUS`
)

func (ps *PsStatus) GetProxySqlStatus(db *sql.DB) PsStatus {

	var tmp Variables
	rows, err := db.Query(StmtMySQLStatus)
	if err != nil {
		log.Print("db.Query", StmtMySQLStatus)
	}
	for rows.Next() {
		tmp = Variables{}
		err = rows.Scan(&tmp.VariablesName, &tmp.Value)
		switch tmp.VariablesName {
		case "Active_Transactions":
			ps.Active_Transactions = tmp.Value
		case "Backend_query_time_nsec":
			ps.Backend_query_time_nsec = tmp.Value
		case "Client_Connections_aborted":
			ps.Client_Connections_aborted = tmp.Value
		case "Client_Connections_connected":
			ps.Client_Connections_connected = tmp.Value
		case "Client_Connections_created":
			ps.Client_Connections_created = tmp.Value
		case "Client_Connections_non_idle":
			ps.Client_Connections_non_idle = tmp.Value
		case "Com_autocommit":
			ps.Com_autocommit = tmp.Value
		case "Com_autocommit_filtered":
			ps.Com_autocommit_filtered = tmp.Value
		case "Com_commit":
			ps.Com_commit = tmp.Value
		case "Com_commit_filtered":
			ps.Com_commit_filtered = tmp.Value
		case "Com_rollback":
			ps.Com_rollback = tmp.Value
		case "Com_rollback_filtered":
			ps.Com_rollback_filtered = tmp.Value
		case "Com_stmt_close":
			ps.Com_stmt_close = tmp.Value
		case "Com_stmt_execute":
			ps.Com_stmt_execute = tmp.Value
		case "Com_stmt_prepare":
			ps.Com_stmt_prepare = tmp.Value
		case "ConnPool_get_conn_failure":
			ps.ConnPool_get_conn_failure = tmp.Value
		case "ConnPool_get_conn_immediate":
			ps.ConnPool_get_conn_immediate = tmp.Value
		case "ConnPool_get_conn_success":
			ps.ConnPool_get_conn_success = tmp.Value
		case "ConnPool_memory_bytes":
			ps.ConnPool_memory_bytes = tmp.Value
		case "MySQL_Monitor_Workers":
			ps.MySQL_Monitor_Workers = tmp.Value
		case "MySQL_Thread_Workers":
			ps.MySQL_Thread_Workers = tmp.Value
		case "ProxySQL_Uptime":
			ps.ProxySQL_Uptime = tmp.Value
		case "Queries_backends_bytes_recv":
			ps.Queries_backends_bytes_recv = tmp.Value
		case "Queries_backends_bytes_sent":
			ps.Queries_backends_bytes_sent = tmp.Value
		case "Query_Cache_Entries":
			ps.Query_Cache_Entries = tmp.Value
		case "Query_Cache_Memory_bytes":
			ps.Query_Cache_Memory_bytes = tmp.Value
		case "Query_Cache_Purged":
			ps.Query_Cache_Purged = tmp.Value
		case "Query_Cache_bytes_IN":
			ps.Query_Cache_bytes_IN = tmp.Value
		case "Query_Cache_bytes_OUT":
			ps.Query_Cache_bytes_OUT = tmp.Value
		case "Query_Cache_count_GET":
			ps.Query_Cache_count_GET = tmp.Value
		case "Query_Cache_count_GET_OK":
			ps.Query_Cache_count_GET_OK = tmp.Value
		case "Query_Cache_count_SET":
			ps.Query_Cache_count_SET = tmp.Value
		case "Query_Processor_time_nsec":
			ps.Query_Processor_time_nsec = tmp.Value
		case "Questions":
			ps.Questions = tmp.Value
		case "SQLite3_memory_bytes":
			ps.SQLite3_memory_bytes = tmp.Value
		case "Server_Connections_aborted":
			ps.Server_Connections_aborted = tmp.Value
		case "Server_Connections_connected":
			ps.Server_Connections_connected = tmp.Value
		case "Server_Connections_created":
			ps.Server_Connections_created = tmp.Value
		case "Servers_table_version":
			ps.Servers_table_version = tmp.Value
		case "Slow_queries":
			ps.Slow_queries = tmp.Value
		case "Stmt_Active_Total":
			ps.Stmt_Active_Total = tmp.Value
		case "Stmt_Active_Unique":
			ps.Stmt_Active_Unique = tmp.Value
		case "Stmt_Max_Stmt_id":
			ps.Stmt_Max_Stmt_id = tmp.Value
		case "mysql_backend_buffers_bytes":
			ps.Mysql_backend_buffers_bytes = tmp.Value
		case "mysql_frontend_buffers_bytes":
			ps.Mysql_frontend_buffers_bytes = tmp.Value
		case "mysql_session_internal_bytes":
			ps.Mysql_session_internal_bytes = tmp.Value
		default:
			log.Print("GetProxySqlStatus()", tmp.VariablesName)
		}
	}
	return *ps
}
