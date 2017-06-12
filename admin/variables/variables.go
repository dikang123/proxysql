package variables

import (
	"database/sql"
	"fmt"
	//"log"
)

type (
	PsVariables struct {
		admin_admin_credentials                int64 `db:"admin-admin_credentials" json:"admin-admin_credentials"`
		admin_hash_passwords                   int64 `db:"admin-hash_passwords" json:"admin-hash_passwords"`
		admin_mysql_ifaces                     int64 `db:"admin-mysql_ifaces" json:"admin-mysql_ifaces"`
		admin_read_only                        int64 `db:"admin-read_only" json:"admin-read_only"`
		admin_refresh_interval                 int64 `db:"admin-refresh_interval" json:"admin-refresh_interval"`
		admin_stats_credentials                int64 `db:"admin-stats_credentials" json:"admin-stats_credentials"`
		admin_telnet_admin_ifaces              int64 `db:"admin-telnet_admin_ifaces" json:"admin-telnet_admin_ifaces"`
		admin_telnet_stats_ifaces              int64 `db:"admin-telnet_stats_ifaces" json:"admin-telnet_stats_ifaces"`
		admin_version                          int64 `db:"admin-version" json:"admin-version"`
		mysql_client_found_rows                int64 `db:"mysql-client_found_rows" json:"mysql-client_found_rows"`
		mysql_commands_stats                   int64 `db:"mysql-commands_stats" json:"mysql-commands_stats"`
		mysql_connect_retries_delay            int64 `db:"mysql-connect_retries_delay" json:"mysql-connect_retries_delay"`
		mysql_connect_retries_on_failure       int64 `db:"mysql-connect_retries_on_failure" json:"mysql-connect_retries_on_failure"`
		mysql_connect_timeout_server           int64 `db:"mysql-connect_timeout_server" json:"mysql-connect_timeout_server"`
		mysql_connect_timeout_server_max       int64 `db:"mysql-connect_timeout_server_max" json:"mysql-connect_timeout_server_max"`
		mysql_connection_max_age_ms            int64 `db:"mysql-connection_max_age_ms" json:"mysql-connection_max_age_ms"`
		mysql_default_charset                  int64 `db:"mysql-default_charset" json:"mysql-default_charset"`
		mysql_default_max_latency_ms           int64 `db:"mysql-default_max_latency_ms" json:"mysql-default_max_latency_ms"`
		mysql_default_query_delay              int64 `db:"mysql-default_query_delay" json:"mysql-default_query_delay"`
		mysql_default_query_timeout            int64 `db:"mysql-default_query_timeout" json:"mysql-default_query_timeout"`
		mysql_default_reconnect                int64 `db:"mysql-default_reconnect" json:"mysql-default_reconnect"`
		mysql_default_schema                   int64 `db:"mysql-default_schema" json:"mysql-default_schema"`
		mysql_default_sql_mode                 int64 `db:"mysql-default_sql_mode" json:"mysql-default_sql_mode"`
		mysql_default_time_zone                int64 `db:"mysql-default_time_zone" json:"mysql-default_time_zone"`
		mysql_enforce_autocommit_on_reads      int64 `db:"mysql-enforce_autocommit_on_reads" json:"mysql-enforce_autocommit_on_reads"`
		mysql_eventslog_filename               int64 `db:"mysql-eventslog_filename" json:"mysql-eventslog_filename"`
		mysql_eventslog_filesize               int64 `db:"mysql-eventslog_filesize" json:"mysql-eventslog_filesize"`
		mysql_free_connections_pct             int64 `db:"mysql-free_connections_pct" json:"mysql-free_connections_pct"`
		mysql_have_compress                    int64 `db:"mysql-have_compress" json:"mysql-have_compress"`
		mysql_init_connect                     int64 `db:"mysql-init_connect" json:"mysql-init_connect"`
		mysql_interfaces                       int64 `db:"mysql-interfaces" json:"mysql-interfaces"`
		mysql_long_query_time                  int64 `db:"mysql-long_query_time" json:"mysql-long_query_time"`
		mysql_max_allowed_packet               int64 `db:"mysql-max_allowed_packet" json:"mysql-max_allowed_packet"`
		mysql_max_connections                  int64 `db:"mysql-max_connections" json:"mysql-max_connections"`
		mysql_max_stmts_cache                  int64 `db:"mysql-max_stmts_cache" json:"mysql-max_stmts_cache"`
		mysql_max_stmts_per_connection         int64 `db:"mysql-max_stmts_per_connection" json:"mysql-max_stmts_per_connection"`
		mysql_max_transaction_time             int64 `db:"mysql-max_transaction_time" json:"mysql-max_transaction_time"`
		mysql_monitor_connect_interval         int64 `db:"mysql-monitor_connect_interval" json:"mysql-monitor_connect_interval"`
		mysql_monitor_connect_timeout          int64 `db:"mysql-monitor_connect_timeout" json:"mysql-monitor_connect_timeout"`
		mysql_monitor_enabled                  int64 `db:"mysql-monitor_enabled" json:"mysql-monitor_enabled"`
		mysql_monitor_history                  int64 `db:"mysql-monitor_history" json:"mysql-monitor_history"`
		mysql_monitor_password                 int64 `db:"mysql-monitor_password" json:"mysql-monitor_password"`
		mysql_monitor_ping_interval            int64 `db:"mysql-monitor_ping_interval" json:"mysql-monitor_ping_interval"`
		mysql_monitor_ping_max_failures        int64 `db:"mysql-monitor_ping_max_failures" json:"mysql-monitor_ping_max_failures"`
		mysql_monitor_ping_timeout             int64 `db:"mysql-monitor_ping_timeout" json:"mysql-monitor_ping_timeout"`
		mysql_monitor_query_interval           int64 `db:"mysql-monitor_query_interval" json:"mysql-monitor_query_interval"`
		mysql_monitor_query_timeout            int64 `db:"mysql-monitor_query_timeout" json:"mysql-monitor_query_timeout"`
		mysql_monitor_read_only_interval       int64 `db:"mysql-monitor_read_only_interval" json:"mysql-monitor_read_only_interval"`
		mysql_monitor_read_only_timeout        int64 `db:"mysql-monitor_read_only_timeout" json:"mysql-monitor_read_only_timeout"`
		mysql_monitor_replication_lag_interval int64 `db:"mysql-monitor_replication_lag_interval" json:"mysql-monitor_replication_lag_interval"`
		mysql_monitor_replication_lag_timeout  int64 `db:"mysql-monitor_replication_lag_timeout" json:"mysql-monitor_replication_lag_timeout"`
		mysql_monitor_slave_lag_when_null      int64 `db:"mysql-monitor_slave_lag_when_null" json:"mysql-monitor_slave_lag_when_null"`
		mysql_monitor_username                 int64 `db:"mysql-monitor_username" json:"mysql-monitor_username"`
		mysql_monitor_writer_is_also_reader    int64 `db:"mysql-monitor_writer_is_also_reader" json:"mysql-monitor_writer_is_also_reader"`
		mysql_multiplexing                     int64 `db:"mysql-multiplexing" json:"mysql-multiplexing"`
		mysql_ping_interval_server_msec        int64 `db:"mysql-ping_interval_server_msec" json:"mysql-ping_interval_server_msec"`
		mysql_ping_timeout_server              int64 `db:"mysql-ping_timeout_server" json:"mysql-ping_timeout_server"`
		mysql_poll_timeout                     int64 `db:"mysql-poll_timeout" json:"mysql-poll_timeout"`
		mysql_poll_timeout_on_failure          int64 `db:"mysql-poll_timeout_on_failure" json:"mysql-poll_timeout_on_failure"`
		mysql_query_cache_size_MB              int64 `db:"mysql-query_cache_size_MB" json:"mysql-query_cache_size_MB"`
		mysql_query_digests                    int64 `db:"mysql-query_digests" json:"mysql-query_digests"`
		mysql_query_digests_lowercase          int64 `db:"mysql-query_digests_lowercase" json:"mysql-query_digests_lowercase"`
		mysql_query_digests_max_digest_length  int64 `db:"mysql-query_digests_max_digest_length" json:"mysql-query_digests_max_digest_length"`
		mysql_query_digests_max_query_length   int64 `db:"mysql-query_digests_max_query_length" json:"mysql-query_digests_max_query_length"`
		mysql_query_processor_iterations       int64 `db:"mysql-query_processor_iterations" json:"mysql-query_processor_iterations"`
		mysql_query_retries_on_failure         int64 `db:"mysql-query_retries_on_failure" json:"mysql-query_retries_on_failure"`
		mysql_server_capabilities              int64 `db:"mysql-server_capabilities" json:"mysql-server_capabilities"`
		mysql_server_version                   int64 `db:"mysql-server_version" json:"mysql-server_version"`
		mysql_servers_stats                    int64 `db:"mysql-servers_stats" json:"mysql-servers_stats"`
		mysql_session_idle_ms                  int64 `db:"mysql-session_idle_ms" json:"mysql-session_idle_ms"`
		mysql_session_idle_show_processlist    int64 `db:"mysql-session_idle_show_processlist" json:"mysql-session_idle_show_processlist"`
		mysql_sessions_sort                    int64 `db:"mysql-sessions_sort" json:"mysql-sessions_sort"`
		mysql_shun_on_failures                 int64 `db:"mysql-shun_on_failures" json:"mysql-shun_on_failures"`
		mysql_shun_recovery_time_sec           int64 `db:"mysql-shun_recovery_time_sec" json:"mysql-shun_recovery_time_sec"`
		mysql_ssl_p2s_ca                       int64 `db:"mysql-ssl_p2s_ca" json:"mysql-ssl_p2s_ca"`
		mysql_ssl_p2s_cert                     int64 `db:"mysql-ssl_p2s_cert" json:"mysql-ssl_p2s_cert"`
		mysql_ssl_p2s_cipher                   int64 `db:"mysql-ssl_p2s_cipher" json:"mysql-ssl_p2s_cipher"`
		mysql_ssl_p2s_key                      int64 `db:"mysql-ssl_p2s_key" json:"mysql-ssl_p2s_key"`
		mysql_stacksize                        int64 `db:"mysql-stacksize" json:"mysql-stacksize"`
		mysql_threads                          int64 `db:"mysql-threads" json:"mysql-threads"`
		mysql_threshold_query_length           int64 `db:"mysql-threshold_query_length" json:"mysql-threshold_query_length"`
		mysql_threshold_resultset_size         int64 `db:"mysql-threshold_resultset_size" json:"mysql-threshold_resultset_size"`
		mysql_wait_timeout                     int64 `db:"mysql-wait_timeout" json:"mysql-wait_timeout"`
	}
	Variables struct {
		VariablesName string `db:"Variable_name" json:"Variable_name"`
		Value         int64  `db:"Value" json:"Value"`
	}
)

const (
	StmtGlobalVariables = `SHOW GLOBAL VARIABLES`
)

func (ps *PsVariables) GetProxySqlVariables(db *sql.DB) PsVariables {
	var tmp Variables

}
