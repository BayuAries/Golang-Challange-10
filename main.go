package main

import (
	"Challange-7/app"
	"Challange-7/config"
)

func init() {
	// load env
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	err := config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
