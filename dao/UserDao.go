package dao

import "golearning/entity"

var userMap = map[string]entity.UserEntity{}

func Save(userEntity entity.UserEntity) {
	userMap[userEntity.Id] = userEntity
}

func GetById(id string) entity.UserEntity {
	return userMap[id]
}

func Update(userEntity entity.UserEntity) {
	userMap[userEntity.Id] = userEntity
}

func DeleteById(id string) {
	delete(userMap, id)
}

func List() []entity.UserEntity {
	println(len(userMap))
	users := []entity.UserEntity{}
	for _, v := range userMap {
		users = append(users, v)
	}
	return users
}
