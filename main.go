package main

import (
	"golearning/deamon/controller"
	"net/http"
)

func main() {
	// url 与 handle 绑定
	http.HandleFunc("/user/list", controller.UserList)
	http.HandleFunc("/user/create", controller.CreateUser)
	http.HandleFunc("/user/update", controller.UpdateUser)
	http.HandleFunc("/user/detail", controller.UserDetail)
	http.HandleFunc("/user/redis", controller.TestRedis)

	// 启动WEB服务器侦听
	http.ListenAndServe(":8088", nil)
}
