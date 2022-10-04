package main

import (
	"privy/app"
)

// func init() {
// 	godotenv.Load()
// 	if err := infra.InitMysql(); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	app.StartApplication()
}
