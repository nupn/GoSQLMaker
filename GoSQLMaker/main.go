package main


import (
	"fmt"
)

func main() {

	p := GetConfig()

	fmt.Println(p.General.SQLServer)
	fmt.Println(p.Tables[0].SQL)
	
	fmt.Println(p.toString())
	fmt.Println(toJson(p))
}
