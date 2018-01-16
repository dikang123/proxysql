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

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		fmt.Println(err)
	}

	i := new(uint64)
	for *i = 0;*i<1000 ;*i++{
		newsch,err := proxysql.NewSch(strconv.Itoa(int(*i)),0)
		if err != nil {
			fmt.Println(err)
		}

		err = newsch.AddOneScheduler(db)
		if err != nil {
			fmt.Println(err)
		}

		
	}

	
}
