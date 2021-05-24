package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//带有接收者的函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()
}