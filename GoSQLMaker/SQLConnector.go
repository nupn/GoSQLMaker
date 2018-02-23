package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type SQLConnector struct {
	Conn *sql.DB
}


func (p* SQLConnector) Connect(server string, port int, id string, pw string, db string) (error, string){

	connectStr := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s", server, port, id, pw, db)

	cn, err := sql.Open("mssql", connectStr)
	if err != nil{
		log.Fatal(err)
		return err, connectStr
	}

	p.Conn = cn
	return nil, ""
}

func (p* SQLConnector) Close(){
	defer p.Conn.Close();
}

//https://github.com/denisenkom/go-mssqldb/blob/master/examples/simple/simple.go 참고
/*
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
*/