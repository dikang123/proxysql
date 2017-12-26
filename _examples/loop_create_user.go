package main

import (
	"strconv"
	"fmt"
	"github.com/imSQL/proxysql"
)

func main() {
	conn, err := proxysql.NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		fmt.Println(err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		newuser, err := proxysql.NewUser("devtest"+strconv.Itoa(i), "devtest", 0, "dev")
		if err != nil {
			fmt.Println(err)
		}

		err = newuser.AddOneUser(db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
