package main

import "fmt"

type Employee struct {
	name string
	age int
}

/*
	这两个[方法]都有接收者, 一个接收者是一个类型, 一个接收者是一个指针类型

	可以看到函数和方法的区别, 就是一个有接收者一个没有接收者, 指定了接收者
	之后, 调用方法的时候必须要 接收者.func() 的形式来进行函数的调用, 函数
	则可以 直接调用就行, 如果没有参数连括号都不用加上去 直接 func 就行, 
	有函数就func(args) 如此进行调用

	 
	而指针接收者是不是复制一份值过去, 而是将this, 将自己本体作为接收者, 方
	法可以用本体中的一些东西, 改变也是改变本体中的东西

	因为方法中制定了要使用Employee类型的指针, 所以调用的时候是(&e).changeAge(args)
	这样来用, 但是golang允许我们省略&符号, 其内部自动将e.changeAge 解析成(&e).changeAge

	那么问题来了:
	[在变量中使用]
	var i int = 1
	*i 无法这么使用, 除非i本来就是指针,那么*i表示的就是指针对应的值, 现在情况看起来i是变量, 所以不能这么用
	&i

	[在类型中使用]
	*Employee 指的是Employee类型的指针, 方法中该参数作为接收者即表示操作的是接收者本体, 方法中作为参数表示
		方法中使用该参数操作的都是本体
	&Employee
		获取Employee类型的地址
			Employee{} 和 &Employee{} 的区别
			应该从左往右看.. 两边都是创建了一个实例.. 左边的是真实的实例对象, 右边是创建实例后的实例地址



*/

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func changeName(newName string) {

}

func changeAge(newAge int) {

}

func main () {
	e := Employee{
		name : "zhans",
		age : 50,
	}



	fmt.Println("name before change: %s", e.name)
	e.changeName("zhanghansen")
	fmt.Println("name after change: %s", e.name)


	fmt.Println("age before change: %s", e.age)
	e.changeAge(20)
	fmt.Println("age after change: %s", e.age)

	// -----------------------------------
	bytes := byte(3)

	fmt.Println("bytes: %v", bytes)



}