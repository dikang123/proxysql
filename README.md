
# Golang ProxySQL Library

-----

### 1.introduce

-----

A ProxySQL Go library.

### 2. Requirements

-----

1. Go 1.7 or higher.
1. ProxySQL 1.3.x

### 3. Installation

Simple install the package to your $GOPATH with the go tool from shell:

    # go get -u github.com/imSQL/proxysql

Make sure git command is installed on your OS.

### 4. Usage

-----

example:

list all mysql_users .

	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}
	
	conn.SetCharacter("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}


### Donate

-----

If you like the project and want to buy me a cola, you can through:

| PayPal                                                                                                               | 微信                                                                 |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [![](https://www.paypalobjects.com/webstatic/paypalme/images/pp_logo_small.png)](https://www.paypal.me/taylor840326) | ![](https://github.com/taylor840326/blog/raw/master/imgs/weixin.png) |


