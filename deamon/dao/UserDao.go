package dao

import (
	entity2 "golearning/deamon/entity"
)

var userMap = map[string]entity2.UserEntity{}

func Save(userEntity entity2.UserEntity) {
	userMap[userEntity.Id] = userEntity
}

func GetById(id string) entity2.UserEntity {
	return userMap[id]
}

func Update(userEntity entity2.UserEntity) {
	userMap[userEntity.Id] = userEntity
}

func DeleteById(id string) {
	delete(userMap, id)
}

func List() []entity2.UserEntity {
	println(len(userMap))
	users := []entity2.UserEntity{}
	for _, v := range userMap {
		users = append(users, v)
	}
	return users
}
