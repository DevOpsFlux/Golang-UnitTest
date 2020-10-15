package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mssql", "server=127.0.0.1;user id=dev;password=@NTdev;database=devdb")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리 : QueryRow()
	var DName, DLocation string
	err = db.QueryRow("SELECT Top 1 Id, Name, Location FROM TestSchema.Employees;").Scan(&DName, &DLocation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(DName, DLocation)
	/*
		// 복수 Row를 갖는 SQL 쿼리 : Query()
		rows, err := db.Query("SELECT au_lname, au_fname FROM authors WHERE au_lname=?", "Ringer")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)

		for rows.Next() {
			err := rows.Scan(&lname, &fname)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(fname, lname)
		}
	*/
}
