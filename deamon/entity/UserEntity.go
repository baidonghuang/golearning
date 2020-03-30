package entity

type UserEntity struct {
	Id   string
	name string //小写属性值为私有变量，不能通过对象导航方式直接设置
	Age  int    //大写的方式为公开变量，可以将属性直接暴露给外部调用
}

func (o UserEntity) GetName() string {
	return o.name
}

/**
定义公开函数设置私有成员变量的值
*/
func (o *UserEntity) SetName(name string) {
	o.name = name
}
