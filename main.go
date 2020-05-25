// main.go

package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
		
	//List env variables (bash): printenv
	a.Run(":8010")
}
