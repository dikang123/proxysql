package variables

import (
	"database/sql"
	"fmt"
	//"fmt"
	"log"
	"proxysql-master/admin/cmd"
)

type (
	PsVariables struct {
		Admin_admin_credentials                int64 `db:"admin_admin_credentials" json:"admin_admin_credentials"`
		Admin_hash_passwords                   int64 `db:"admin_hash_passwords" json:"admin_hash_passwords"`
		Admin_mysql_ifaces                     int64 `db:"admin_mysql_ifaces" json:"admin_mysql_ifaces"`
		Admin_read_only                        int64 `db:"admin_read_only" json:"admin_read_only"`
		Admin_refresh_interval                 int64 `db:"admin_refresh_interval" json:"admin_refresh_interval"`
		Admin_stats_credentials                int64 `db:"admin_stats_credentials" json:"admin_stats_credentials"`
		Admin_telnet_admin_ifaces              int64 `db:"admin_telnet_admin_ifaces" json:"admin_telnet_admin_ifaces"`
		Admin_telnet_stats_ifaces              int64 `db:"admin_telnet_stats_ifaces" json:"admin_telnet_stats_ifaces"`
		Admin_version                          int64 `db:"admin_version" json:"admin_version"`
		Mysql_client_found_rows                int64 `db:"mysql_client_found_rows" json:"mysql_client_found_rows"`
		Mysql_commands_stats                   int64 `db:"mysql_commands_stats" json:"mysql_commands_stats"`
		Mysql_connect_retries_delay            int64 `db:"mysql_connect_retries_delay" json:"mysql_connect_retries_delay"`
		Mysql_connect_retries_on_failure       int64 `db:"mysql_connect_retries_on_failure" json:"mysql_connect_retries_on_failure"`
		Mysql_connect_timeout_server           int64 `db:"mysql_connect_timeout_server" json:"mysql_connect_timeout_server"`
		Mysql_connect_timeout_server_max       int64 `db:"mysql_connect_timeout_server_max" json:"mysql_connect_timeout_server_max"`
		Mysql_connection_max_age_ms            int64 `db:"mysql_connection_max_age_ms" json:"mysql_connection_max_age_ms"`
		Mysql_default_charset                  int64 `db:"mysql_default_charset" json:"mysql_default_charset"`
		Mysql_default_max_latency_ms           int64 `db:"mysql_default_max_latency_ms" json:"mysql_default_max_latency_ms"`
		Mysql_default_query_delay              int64 `db:"mysql_default_query_delay" json:"mysql_default_query_delay"`
		Mysql_default_query_timeout            int64 `db:"mysql_default_query_timeout" json:"mysql_default_query_timeout"`
		Mysql_default_reconnect                int64 `db:"mysql_default_reconnect" json:"mysql_default_reconnect"`
		Mysql_default_schema                   int64 `db:"mysql_default_schema" json:"mysql_default_schema"`
		Mysql_default_sql_mode                 int64 `db:"mysql_default_sql_mode" json:"mysql_default_sql_mode"`
		Mysql_default_time_zone                int64 `db:"mysql_default_time_zone" json:"mysql_default_time_zone"`
		Mysql_enforce_autocommit_on_reads      int64 `db:"mysql_enforce_autocommit_on_reads" json:"mysql_enforce_autocommit_on_reads"`
		Mysql_eventslog_filename               int64 `db:"mysql_eventslog_filename" json:"mysql_eventslog_filename"`
		Mysql_eventslog_filesize               int64 `db:"mysql_eventslog_filesize" json:"mysql_eventslog_filesize"`
		Mysql_free_connections_pct             int64 `db:"mysql_free_connections_pct" json:"mysql_free_connections_pct"`
		Mysql_have_compress                    int64 `db:"mysql_have_compress" json:"mysql_have_compress"`
		Mysql_init_connect                     int64 `db:"mysql_init_connect" json:"mysql_init_connect"`
		Mysql_interfaces                       int64 `db:"mysql_interfaces" json:"mysql_interfaces"`
		Mysql_long_query_time                  int64 `db:"mysql_long_query_time" json:"mysql_long_query_time"`
		Mysql_max_allowed_packet               int64 `db:"mysql_max_allowed_packet" json:"mysql_max_allowed_packet"`
		Mysql_max_connections                  int64 `db:"mysql_max_connections" json:"mysql_max_connections"`
		Mysql_max_stmts_cache                  int64 `db:"mysql_max_stmts_cache" json:"mysql_max_stmts_cache"`
		Mysql_max_stmts_per_connection         int64 `db:"mysql_max_stmts_per_connection" json:"mysql_max_stmts_per_connection"`
		Mysql_max_transaction_time             int64 `db:"mysql_max_transaction_time" json:"mysql_max_transaction_time"`
		Mysql_monitor_connect_interval         int64 `db:"mysql_monitor_connect_interval" json:"mysql_monitor_connect_interval"`
		Mysql_monitor_connect_timeout          int64 `db:"mysql_monitor_connect_timeout" json:"mysql_monitor_connect_timeout"`
		Mysql_monitor_enabled                  int64 `db:"mysql_monitor_enabled" json:"mysql_monitor_enabled"`
		Mysql_monitor_history                  int64 `db:"mysql_monitor_history" json:"mysql_monitor_history"`
		Mysql_monitor_password                 int64 `db:"mysql_monitor_password" json:"mysql_monitor_password"`
		Mysql_monitor_ping_interval            int64 `db:"mysql_monitor_ping_interval" json:"mysql_monitor_ping_interval"`
		Mysql_monitor_ping_max_failures        int64 `db:"mysql_monitor_ping_max_failures" json:"mysql_monitor_ping_max_failures"`
		Mysql_monitor_ping_timeout             int64 `db:"mysql_monitor_ping_timeout" json:"mysql_monitor_ping_timeout"`
		Mysql_monitor_query_interval           int64 `db:"mysql_monitor_query_interval" json:"mysql_monitor_query_interval"`
		Mysql_monitor_query_timeout            int64 `db:"mysql_monitor_query_timeout" json:"mysql_monitor_query_timeout"`
		Mysql_monitor_read_only_interval       int64 `db:"mysql_monitor_read_only_interval" json:"mysql_monitor_read_only_interval"`
		Mysql_monitor_read_only_timeout        int64 `db:"mysql_monitor_read_only_timeout" json:"mysql_monitor_read_only_timeout"`
		Mysql_monitor_replication_lag_interval int64 `db:"mysql_monitor_replication_lag_interval" json:"mysql_monitor_replication_lag_interval"`
		Mysql_monitor_replication_lag_timeout  int64 `db:"mysql_monitor_replication_lag_timeout" json:"mysql_monitor_replication_lag_timeout"`
		Mysql_monitor_slave_lag_when_null      int64 `db:"mysql_monitor_slave_lag_when_null" json:"mysql_monitor_slave_lag_when_null"`
		Mysql_monitor_username                 int64 `db:"mysql_monitor_username" json:"mysql_monitor_username"`
		Mysql_monitor_writer_is_also_reader    int64 `db:"mysql_monitor_writer_is_also_reader" json:"mysql_monitor_writer_is_also_reader"`
		Mysql_multiplexing                     int64 `db:"mysql_multiplexing" json:"mysql_multiplexing"`
		Mysql_ping_interval_server_msec        int64 `db:"mysql_ping_interval_server_msec" json:"mysql_ping_interval_server_msec"`
		Mysql_ping_timeout_server              int64 `db:"mysql_ping_timeout_server" json:"mysql_ping_timeout_server"`
		Mysql_poll_timeout                     int64 `db:"mysql_poll_timeout" json:"mysql_poll_timeout"`
		Mysql_poll_timeout_on_failure          int64 `db:"mysql_poll_timeout_on_failure" json:"mysql_poll_timeout_on_failure"`
		Mysql_query_cache_size_MB              int64 `db:"mysql_query_cache_size_MB" json:"mysql_query_cache_size_MB"`
		Mysql_query_digests                    int64 `db:"mysql_query_digests" json:"mysql_query_digests"`
		Mysql_query_digests_lowercase          int64 `db:"mysql_query_digests_lowercase" json:"mysql_query_digests_lowercase"`
		Mysql_query_digests_max_digest_length  int64 `db:"mysql_query_digests_max_digest_length" json:"mysql_query_digests_max_digest_length"`
		Mysql_query_digests_max_query_length   int64 `db:"mysql_query_digests_max_query_length" json:"mysql_query_digests_max_query_length"`
		Mysql_query_processor_iterations       int64 `db:"mysql_query_processor_iterations" json:"mysql_query_processor_iterations"`
		Mysql_query_retries_on_failure         int64 `db:"mysql_query_retries_on_failure" json:"mysql_query_retries_on_failure"`
		Mysql_server_capabilities              int64 `db:"mysql_server_capabilities" json:"mysql_server_capabilities"`
		Mysql_server_version                   int64 `db:"mysql_server_version" json:"mysql_server_version"`
		Mysql_servers_stats                    int64 `db:"mysql_servers_stats" json:"mysql_servers_stats"`
		Mysql_session_idle_ms                  int64 `db:"mysql_session_idle_ms" json:"mysql_session_idle_ms"`
		Mysql_session_idle_show_processlist    int64 `db:"mysql_session_idle_show_processlist" json:"mysql_session_idle_show_processlist"`
		Mysql_sessions_sort                    int64 `db:"mysql_sessions_sort" json:"mysql_sessions_sort"`
		Mysql_shun_on_failures                 int64 `db:"mysql_shun_on_failures" json:"mysql_shun_on_failures"`
		Mysql_shun_recovery_time_sec           int64 `db:"mysql_shun_recovery_time_sec" json:"mysql_shun_recovery_time_sec"`
		Mysql_ssl_p2s_ca                       int64 `db:"mysql_ssl_p2s_ca" json:"mysql_ssl_p2s_ca"`
		Mysql_ssl_p2s_cert                     int64 `db:"mysql_ssl_p2s_cert" json:"mysql_ssl_p2s_cert"`
		Mysql_ssl_p2s_cipher                   int64 `db:"mysql_ssl_p2s_cipher" json:"mysql_ssl_p2s_cipher"`
		Mysql_ssl_p2s_key                      int64 `db:"mysql_ssl_p2s_key" json:"mysql_ssl_p2s_key"`
		Mysql_stacksize                        int64 `db:"mysql_stacksize" json:"mysql_stacksize"`
		Mysql_threads                          int64 `db:"mysql_threads" json:"mysql_threads"`
		Mysql_threshold_query_length           int64 `db:"mysql_threshold_query_length" json:"mysql_threshold_query_length"`
		Mysql_threshold_resultset_size         int64 `db:"mysql_threshold_resultset_size" json:"mysql_threshold_resultset_size"`
		Mysql_wait_timeout                     int64 `db:"mysql_wait_timeout" json:"mysql_wait_timeout"`
	}
	Variables struct {
		VariablesName string `db:"Variable_name" json:"Variable_name"`
		Value         int64  `db:"Value" json:"Value"`
	}
)

const (
	StmtGlobalVariables   = `SHOW GLOBAL VARIABLES`
	StmtUpdateOneVariable = `UPDATE global_variables SET variable_value=%q WHERE variable_name = %q`
)

func (vars *Variables) UpdateOneVariable(db *sql.DB) (int, error) {
	st := fmt.Sprintf(StmtUpdateOneVariable, vars.Value, vars.VariablesName)
	log.Print("variables.go->UpdateOneVariable->st:", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("UpdateOneVariable->db.Query: ", err)
		return 1, err
	}
	cmd.LoadMySQlVariablesToRuntime(db)
	cmd.LoadAdminVariablesToRuntime(db)
	cmd.SaveMySQLVariablesToDisk(db)
	cmd.SaveAdminVariablesToDisk(db)
	return 0, nil
}

func (ps *PsVariables) GetProxySqlVariables(db *sql.DB) PsVariables {
	var tmp Variables
	log.Print("Execution: ", StmtGlobalVariables)
	rows, err := db.Query(StmtGlobalVariables)
	if err != nil {
		log.Print("StmtGlobalVariables Msg:", err)
		return PsVariables{}
	}

	for rows.Next() {
		tmp = Variables{}
		err = rows.Scan(&tmp.VariablesName, &tmp.Value)
		switch tmp.VariablesName {
		case "admin-admin_credentials":
			ps.Admin_admin_credentials = tmp.Value
		case "admin-hash_passwords":
			ps.Admin_hash_passwords = tmp.Value
		case "admin-mysql_ifaces":
			ps.Admin_mysql_ifaces = tmp.Value
		case "admin-read_only":
			ps.Admin_read_only = tmp.Value
		case "admin-refresh_interval":
			ps.Admin_refresh_interval = tmp.Value
		case "admin-stats_credentials":
			ps.Admin_stats_credentials = tmp.Value
		case "admin-telnet_admin_ifaces":
			ps.Admin_telnet_admin_ifaces = tmp.Value
		case "admin-telnet_stats_ifaces":
			ps.Admin_telnet_stats_ifaces = tmp.Value
		case "admin-version":
			ps.Admin_version = tmp.Value
		case "mysql-client_found_rows":
			ps.Mysql_client_found_rows = tmp.Value
		case "mysql-commands_stats":
			ps.Mysql_commands_stats = tmp.Value
		case "mysql-connect_retries_delay":
			ps.Mysql_connect_retries_delay = tmp.Value
		case "mysql-connect_retries_on_failure":
			ps.Mysql_connect_retries_on_failure = tmp.Value
		case "mysql-connect_timeout_server":
			ps.Mysql_connect_timeout_server = tmp.Value
		case "mysql-connect_timeout_server_max":
			ps.Mysql_connect_timeout_server_max = tmp.Value
		case "mysql-connection_max_age_ms":
			ps.Mysql_connection_max_age_ms = tmp.Value
		case "mysql-default_charset":
			ps.Mysql_default_charset = tmp.Value
		case "mysql-default_max_latency_ms":
			ps.Mysql_default_max_latency_ms = tmp.Value
		case "mysql-default_query_delay":
			ps.Mysql_default_query_delay = tmp.Value
		case "mysql-default_query_timeout":
			ps.Mysql_default_query_timeout = tmp.Value
		case "mysql-default_reconnect":
			ps.Mysql_default_reconnect = tmp.Value
		case "mysql-default_schema":
			ps.Mysql_default_schema = tmp.Value
		case "mysql-default_sql_mode":
			ps.Mysql_default_sql_mode = tmp.Value
		case "mysql-default_time_zone":
			ps.Mysql_default_time_zone = tmp.Value
		case "mysql-enforce_autocommit_on_reads":
			ps.Mysql_enforce_autocommit_on_reads = tmp.Value
		case "mysql-eventslog_filename":
			ps.Mysql_eventslog_filename = tmp.Value
		case "mysql-eventslog_filesize":
			ps.Mysql_eventslog_filesize = tmp.Value
		case "mysql-free_connections_pct":
			ps.Mysql_free_connections_pct = tmp.Value
		case "mysql-have_compress":
			ps.Mysql_have_compress = tmp.Value
		case "mysql-init_connect":
			ps.Mysql_init_connect = tmp.Value
		case "mysql-interfaces":
			ps.Mysql_interfaces = tmp.Value
		case "mysql-long_query_time":
			ps.Mysql_long_query_time = tmp.Value
		case "mysql-max_allowed_packet":
			ps.Mysql_max_allowed_packet = tmp.Value
		case "mysql-max_connections":
			ps.Mysql_max_connections = tmp.Value
		case "mysql-max_stmts_cache":
			ps.Mysql_max_stmts_cache = tmp.Value
		case "mysql-max_stmts_per_connection":
			ps.Mysql_max_stmts_per_connection = tmp.Value
		case "mysql-max_transaction_time":
			ps.Mysql_max_transaction_time = tmp.Value
		case "mysql-monitor_connect_interval":
			ps.Mysql_monitor_connect_interval = tmp.Value
		case "mysql-monitor_connect_timeout":
			ps.Mysql_monitor_connect_timeout = tmp.Value
		case "mysql-monitor_enabled":
			ps.Mysql_monitor_enabled = tmp.Value
		case "mysql-monitor_history":
			ps.Mysql_monitor_history = tmp.Value
		case "mysql-monitor_password":
			ps.Mysql_monitor_password = tmp.Value
		case "mysql-monitor_ping_interval":
			ps.Mysql_monitor_ping_interval = tmp.Value
		case "mysql-monitor_ping_max_failures":
			ps.Mysql_monitor_ping_max_failures = tmp.Value
		case "mysql-monitor_ping_timeout":
			ps.Mysql_monitor_ping_timeout = tmp.Value
		case "mysql-monitor_query_interval":
			ps.Mysql_monitor_query_interval = tmp.Value
		case "mysql-monitor_query_timeout":
			ps.Mysql_monitor_query_timeout = tmp.Value
		case "mysql-monitor_read_only_interval":
			ps.Mysql_monitor_read_only_interval = tmp.Value
		case "mysql-monitor_read_only_timeout":
			ps.Mysql_monitor_read_only_timeout = tmp.Value
		case "mysql-monitor_replication_lag_interval":
			ps.Mysql_monitor_replication_lag_interval = tmp.Value
		case "mysql-monitor_replication_lag_timeout":
			ps.Mysql_monitor_replication_lag_timeout = tmp.Value
		case "mysql-monitor_slave_lag_when_null":
			ps.Mysql_monitor_slave_lag_when_null = tmp.Value
		case "mysql-monitor_username":
			ps.Mysql_monitor_username = tmp.Value
		case "mysql-monitor_writer_is_also_reader":
			ps.Mysql_monitor_writer_is_also_reader = tmp.Value
		case "mysql-multiplexing":
			ps.Mysql_multiplexing = tmp.Value
		case "mysql-ping_interval_server_msec":
			ps.Mysql_ping_interval_server_msec = tmp.Value
		case "mysql-ping_timeout_server":
			ps.Mysql_ping_timeout_server = tmp.Value
		case "mysql-poll_timeout":
			ps.Mysql_poll_timeout = tmp.Value
		case "mysql-poll_timeout_on_failure":
			ps.Mysql_poll_timeout_on_failure = tmp.Value
		case "mysql-query_cache_size_MB":
			ps.Mysql_query_cache_size_MB = tmp.Value
		case "mysql-query_digests":
			ps.Mysql_query_digests = tmp.Value
		case "mysql-query_digests_lowercase":
			ps.Mysql_query_digests_lowercase = tmp.Value
		case "mysql-query_digests_max_digest_length":
			ps.Mysql_query_digests_max_digest_length = tmp.Value
		case "mysql-query_digests_max_query_length":
			ps.Mysql_query_digests_max_query_length = tmp.Value
		case "mysql-query_processor_iterations":
			ps.Mysql_query_processor_iterations = tmp.Value
		case "mysql-query_retries_on_failure":
			ps.Mysql_query_retries_on_failure = tmp.Value
		case "mysql-server_capabilities":
			ps.Mysql_server_capabilities = tmp.Value
		case "mysql-server_version":
			ps.Mysql_server_version = tmp.Value
		case "mysql-servers_stats":
			ps.Mysql_servers_stats = tmp.Value
		case "mysql-session_idle_ms":
			ps.Mysql_session_idle_ms = tmp.Value
		case "mysql-session_idle_show_processlist":
			ps.Mysql_session_idle_show_processlist = tmp.Value
		case "mysql-sessions_sort":
			ps.Mysql_sessions_sort = tmp.Value
		case "mysql-shun_on_failures":
			ps.Mysql_shun_on_failures = tmp.Value
		case "mysql-shun_recovery_time_sec":
			ps.Mysql_shun_recovery_time_sec = tmp.Value
		case "mysql-ssl_p2s_ca":
			ps.Mysql_ssl_p2s_ca = tmp.Value
		case "mysql-ssl_p2s_cert":
			ps.Mysql_ssl_p2s_cert = tmp.Value
		case "mysql-ssl_p2s_cipher":
			ps.Mysql_ssl_p2s_cipher = tmp.Value
		case "mysql-ssl_p2s_key":
			ps.Mysql_ssl_p2s_key = tmp.Value
		case "mysql-stacksize":
			ps.Mysql_stacksize = tmp.Value
		case "mysql-threads":
			ps.Mysql_threads = tmp.Value
		case "mysql-threshold_query_length":
			ps.Mysql_threshold_query_length = tmp.Value
		case "mysql-threshold_resultset_size":
			ps.Mysql_threshold_resultset_size = tmp.Value
		case "mysql-wait_timeout":
			ps.Mysql_wait_timeout = tmp.Value
		default:
			log.Print("GetProxySqlVariables() ", tmp.VariablesName)
		}
	}
	log.Printf("GetProxySqlVariables tmp variables =%#v", *ps)
	return *ps

}
