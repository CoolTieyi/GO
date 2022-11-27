package main

import (
	"fmt"
)

func def1() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func def2() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}
func def3() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 1
}

func main() {

	//defer
	/*
		defer 用于资源释放 在函数返回之前调用
		1. 多个defer的话 类似于栈 先进后出
		2. 实际使用的话 要注意:
			return其实并不是一条原子语句 是由赋值和返回组成的 分为两步:
			1)返回值= xxx 2) return
			但是有defer了之后变成
			1)返回值= xxx 2) defer 3)return
		汇编代码为
			(call runtime.deferreturn)  //有defer时
			and xx SP
			return
	*/
	fmt.Println(def1())
	fmt.Println(def2())
	fmt.Println(def3())
	/*
		def1()：result = 0
				func(){ result++ }()
				return	//返回1
		def2()：func f() (r int){
				t:=5
				r=t
				func(){t=t+5} }
				return	//defer被插入到赋值和返回之间 这里r没有被修改过 所以返回1
		def3()：result = 0
				func(){ result++ }()
				return	//1
	*/
	fmt.Println("===============")

	//variadic function可变函数 参数数量可变的函数称为为可变参数函数
	fmt.Println(sum())        //0
	fmt.Println(sum(3))       //3
	fmt.Println(sum(4, 5, 6)) //15
	val := []int{4, 5, 6}
	fmt.Println(sum(val...)) //15	跟上面一样 一定得有这个...

}

func sum(vals ...int) int { //vals被看作[]int的切片
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
