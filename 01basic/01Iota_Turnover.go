package main

import "fmt"

func main() {

	//常量计数器
	const ia = iota
	const ib = iota
	const (
		ia2 = iota //iota在一组常量里面是会计数的
		ib2 = "achar"
		ic2        //不赋值的话跟上面一个一样
		id2 = iota //恢复计数，这里的计数并不受上面的赋值影响
	)
	print("\nconst:", "\nia:", ia, "\tib:", ib, "\tic:", ic2,
		"\nia2:", ia2, "\tib2:", ib2, "\tic2:", ic2, "\tid2:", id2)

	//变量交换
	exA := 10
	exB := 20
	exA, exB = exB, exA //exchange core

	//类型转换
	tc := "string"
	fmt.Println([]rune(tc)) //字符串转化要用数组接收	//golang好像没有类似char的类型吧？
	var doub float64 = 4.2
	fmt.Println(int(doub))

}
