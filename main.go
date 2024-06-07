package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")
	fmt.Println(models.Db)

	// テストユーザーを登録
	u := &models.User{}
	u.Name = "test"
	u.Email = "hoegfeoh@hoga.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
}