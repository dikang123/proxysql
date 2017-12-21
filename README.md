
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

    import "github.com/imSQL/proxysql"

    proxysql.FindAllUserInfo(db,0,0)
