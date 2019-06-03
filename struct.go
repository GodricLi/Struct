package main

import (
	"fmt"
)

//struct是一种用户自定义数据类型，可以包含复杂的数据类型
//使用type定义struct类型，先定义结构
type User struct {
	Name string
	Age  int
	Sex  string
}

func defind_struct() {

	//初始化方式一
	var user User
	user.Age = 18
	user.Name = "godric"
	user.Sex = "fm"

	// 访问属性
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user)

	//初始化方式二：
	// user2:=User{}
	var user2 User = User{
		Name: "alex",
		Age:  19,
		Sex:  "ml",
	}
	fmt.Println(user2.Sex)
}

func point_struct() {
	//结构体类型的指针
	var user *User = &User{
		Name: "rain",
		Age:  12,
		Sex:  "男",
	}
	// 访问属性，可以直接使用user.Age
	fmt.Println((*user).Age)
	fmt.Println(user.Age)

	//new方法
	var user2 *User = new(User)
	user2.Age = 14
	user2.Name = "lee"
	user2.Sex = "sss"
	fmt.Println(user2)
}

func structure_struct(name string, age int, sex string) *User {
	//结构体没有构造函数，必须自己实现,返回指针类型
	// 方式一：使用new来分配内存地址
	user := new(User)
	user.Name = name
	user.Age = age
	user.Sex = sex

	// 方式二：
	/*	user2 := &User{
		Name : name,
		Age : age,
		Sex : sex,
	}*/
	return user
}

//定义一个匿名结构体,没有名字的字段，访问复制时使用字段的类型
type As struct {
	name string
	age int
	int
}

func anonymous_struct()  {
	//匿名结构体
	a := As{
		name:"alex",
		age : 18,
		int:66,//访问复制时使用字段的类型
	}
	fmt.Println(a,a.int)

}

type address struct{
	Province string
	City string
}

func nest_struct()  {
	//结构体的嵌套
	type nest struct{
		name string
		age int
		addr address
	}
	//嵌套结构域体的初始化
	//方式一：
	user := nest{
		name : "godric",
		age : 18,
		addr : address{
			Province : "beij",
			City : "beij",

		},
	}
	fmt.Println(user)

	// 当嵌套体为匿名字段时，直接使用嵌套内部的字段，user.字段不在nest里面时，会自动往嵌套的address里面去找
	//当两个结构存在字段相同时默认访问外部字段，访问嵌套结构体字段需要指定结构体user.address.Province
	type nest2 struct{
		name string
		age int
		address
	}
	var user2 nest2
	user2.name="alex"
	user2.age=19
	user2.Province="beij2"
	user2.City="beij2"

	fmt.Printf("%#v\n",user2)
}

func main() {
	// defind_struct()
	// point_struct()
	user := structure_struct("godric", 18, "男")
	fmt.Println(user)
	anonymous_struct()
	nest_struct()
}
