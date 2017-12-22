
# Golang ProxySQL Library

-----

### 1.introduce

A ProxySQL Go library.

### 2. Requirements

1. Go 1.7 or higher.
1. ProxySQL 1.3.x

### 3. Installation

Simple install the package to your $GOPATH with the go tool from shell:

    # go get -u github.com/imSQL/proxysql

Make sure git command is installed on your OS.

### 4. Usage

example:

list all mysql_users .

	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}
