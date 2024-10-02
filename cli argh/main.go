package main

import (
	"flag"
	"fmt"
)

func main() {
	// os.args takes argumnts passed to comand line
	// for index, data := range os.Args {
	// 	fmt.Println(index, data)
	// }

	// flag get command line flags line -pass
	//flag.String("name", "default value", "desc")
	username := flag.String("username", "", "used as a username")
	pass := flag.String("pass", "", "get pass")
	flag.Parse()

	fmt.Printf("\nusername :%v \npassword :%v", *username, *pass)
}
