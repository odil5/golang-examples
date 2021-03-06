package main


import (
		"database/sql"
		"io/ioutil"
		"fmt"
		"strings"

		"github.com/BurntSushi/toml"
		_ "github.com/go-sql-driver/mysql"
	)



type App struct {
	MYSQL	MYSQL
	POSTGRE POSTGRE
}


type MYSQL struct {
	DSN	string
}


type POSTGRE struct {
	Login1	string
	Passw1	string
}


func main() {

	var app App

	f, err := ioutil.ReadFile("conf.toml")
	if err != nil{
		panic(err)
	}

	c := strings.NewReader(string(f))
	_ , err = toml.DecodeReader(c , &app)
	if err != nil{
		panic(err)
	}
	fmt.Println(app)

	
	db, err := sql.Open("mysql", app.MYSQL.DSN)
	if err != nil{
		panic(err)
	}

	rows, err := db.Query("select `id` from max_units")
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	//var ids = make(map[int]int)
	for rows.Next() {
		var id int
		fmt.Println(id)

		rows.Scan(&id)
	//	ids[id] = id
	}
	
	



//fmt.Println(app.MYSQL.Login)
	
/*
	a:=1;
	b:=8

	if a > b {
		fmt.Println(" 1 ")
	}else{
		fmt.Println(" 0 ")
	}


	fmt.Println(shouldEscape('?'))
*/

}
