package main

import 	(
			"fmt"
			"strings"
			"io/ioutil"
			"github.com/BurntSushi/toml"
		)


type App struct {
	FIRST_SOME_PARM		FIRST_SOME_PARM
	SECOND_SOME_PARM	SECOND_SOME_PARM
}


type FIRST_SOME_PARM struct {
	USER	string
	PASW	string
}


type SECOND_SOME_PARM struct {
	LOGIN	string
	PASS	string
}

func main() {

	f1 := "Demo how to read from conf.toml"
	fmt.Println(f1)
	
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

	fmt.Println("User ="+app.FIRST_SOME_PARM.USER+" Pasw="+app.FIRST_SOME_PARM.PASW)
	fmt.Println("User ="+app.SECOND_SOME_PARM.LOGIN+" Pasw="+app.SECOND_SOME_PARM.PASS)

}


