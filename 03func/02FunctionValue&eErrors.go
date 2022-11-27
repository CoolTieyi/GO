package main

import "fmt"

func funca(x int) int {
	return x * x
}

func main() {
	//Errors
	//err是接口类型 返回nil成功  返回non-nil失败
	//fmt.Printf("%v",err)

	//function value
	var fv = funca     //等于f:=square	//函数声明可以这么声明的啊!
	fmt.Println(fv(3)) //调用的时候直接传字面量
	fmt.Println(fv(3) == fv(3))

	i2 := func(x, y int) int { return x + y }(1, 2) //func(入参列表) 返回值类型 {函数体} (实参)  //就相当于定义完直接调用了
	fmt.Println(i2)
}
