package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

//https://github.com/denisenkom/go-mssqldb/blob/master/examples/simple/simple.go 참고

func main() {
	db, err := sql.Open("mssql", "server=175.207.4.91;port=1433;user id=nupn56;password=Dudah256u;database=UserContent")
	if err != nil{
		log.Fatal(err)
		fmt.Println("Fail To Connect")
	}
	defer db.Close();

	var userid string
	var point int
	rows, err := db.Query("SELECT * from [UserContent].[dbo].[FxDailyMissionPoint]")
	if err != nil{
		log.Fatal(err)
		fmt.Println("Fail To Query")
	}

	for rows.Next() {
		err := rows.Scan(&userid, &point);
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userid, point);
	}

	defer rows.Close();
} 