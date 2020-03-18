package controller

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"golearning/cache"
	"golearning/dao"
	"golearning/entity"
	"golearning/service"
	"net/http"
	"strconv"
)

/**
修改用户handle函数
*/
func UserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---->  controller's method 'UserList' is called")
	json, _ := json.Marshal(service.UserList())
	fmt.Println("用户列表：" + string(json))
	w.Write(json)
}

/**
用户详情handle函数
*/
func UserDetail(w http.ResponseWriter, r *http.Request) {
	//获取参数之前必须调用ParseForm
	r.ParseForm()
	if len(r.Form["id"]) > 0 {
		fmt.Println(r.FormValue("id"))
		fmt.Println(r.Form["id"][0])
	}

	user := dao.GetById(r.FormValue("id"))
	userJson, _ := json.Marshal(user)
	fmt.Println("用户详情：" + string(userJson))
	//header.set 必须在Write和WriteHeader之前。
	w.Header().Set("name", "hbd")
	w.Write(userJson)
	w.WriteHeader(200)
}

/**
新建用户handle函数
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("---->  controller's method 'CreateUser' is called")

	//获取参数之前必须调用ParseForm
	r.ParseForm()
	if len(r.Form["name"]) > 0 {
		fmt.Println(r.FormValue("name"))
		fmt.Println(r.Form["name"][0])
		fmt.Println(r.PostFormValue("name"))
	}

	var user entity.UserEntity

	// 设置年龄
	age, err := strconv.Atoi(r.PostFormValue("age"))
	if err != nil {
		w.Write([]byte("年龄不允许为空"))
		w.WriteHeader(500)
	} else {
		uid := uuid.Must(uuid.NewV4())
		user = entity.UserEntity{Id: uid.String(), Age: age}
	}

	// 设置姓名
	name := r.PostFormValue("name")
	if name != "" {
		user.SetName(name)
	}
	resultMsg, errMsg := service.SaveUser(user)
	if errMsg != "" {
		w.Write([]byte(errMsg))
		w.WriteHeader(500)
	}
	w.Write([]byte(resultMsg))
}

/**
修改用户handle函数
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---->  controller's method 'UpdateUser' is called")
	user := entity.UserEntity{Age: 1}
	service.SaveUser(user)
}

/**
修改用户handle函数
*/
func TestRedis(w http.ResponseWriter, r *http.Request) {
	cache.SetOne("go-client-key", "hello redis client")
	val := cache.GetOne("go-client-key")
	fmt.Println("get ", val)
}
