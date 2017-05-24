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

