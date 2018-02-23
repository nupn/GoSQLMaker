package main

import (
	"bufio"
	//"fmt"
	"io/ioutil"
	"os"
)


func main() {

	//Way1 : Write With IoUtil => "io/ioutil"
	bytes := []byte("selet * from")
	bytes = append(bytes, []byte("\nselet * from")...)
	err := ioutil.WriteFile("./TEST.sql", bytes, 0)
	if err != nil{
		panic(err)
	}



	f, err := os.Create("./Test2.sql")
	if err != nil {
	}

	defer f.Close()


	//Way2 : Write With OS
	_, err = f.WriteString("Select * from ")
	if err != nil {

	}

	//way 3 : use Bufio
	w := bufio.NewWriter(f)
	_, err = w.WriteString("buffered\n")
	w.Flush()
}