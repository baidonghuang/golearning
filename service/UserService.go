package service

import (
	"errors"
	"fmt"
	"golearning/dao"
	"golearning/entity"
	"golearning/exception"
	"strconv"
)

// 成员变量
var userCount = 1

// 全局变量/静态变量
var ServiceName = "Old User Service"

func UserList() []entity.UserEntity {
	return dao.List()
}

// 公开函数，新建用户
func SaveUser(user entity.UserEntity) (result string, err string) {
	// 1、表单校验
	err = checkUser(user)
	if err != "" {
		fmt.Println(err)
		return
	}
	dao.Save(user)
	result = "user id is :" + user.Id + "My name is :" + user.GetName() + ", My age is " + strconv.Itoa(user.Age) + ", My number is " + strconv.Itoa(userCount)
	fmt.Println(result)
	userCount++
	return
}

//私有函数
func checkUser(user entity.UserEntity) (msg string) {

	if user.Age < 18 {
		// 返回系统异常
		return errors.New("未满十八岁，不允许调用保存").Error()
	}

	if user.GetName() == "" {
		//返回自定义异常
		error := exception.ServiceException{ErrorMsg: "ServiceException 姓名不能为空", ErrorCode: 10001}
		return error.Error()
	}

	return ""
}
