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
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "ddddd@hoga.com"
	// u.Password = "ddddd"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@hoe.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

		// user, _ := models.GetUser(2)
		// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("DDD first Todo")

	// get all todos
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// get todos by user
	userD, _ := models.GetUser(3)
	todos, _ := userD.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}


}