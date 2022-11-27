package main

import (
	"fmt"
	"math"
)

func main() {

	//bool可能会有短路情况 如果运算符左边值已经可以确定整个布尔表达式的值 那么运算符右边的值将不再被求值
	fmt.Println("===== Bool:====")
	s := "x"
	fmt.Println(s != "" && s[0] == 'x') //永远安全
	fmt.Println(s == "" && s[0] == 'x') //这里左边运算符已经可以决定结果了 所以右边不会被判断
	fmt.Println(s != "" || s[0] != 'x') //同上

	//complex 复数	complex64 complex128
	fmt.Println("===== complex:====")
	var x complex128 = complex(1, 2) // x:=1+2i
	var y complex128 = complex(3, 4) // y:=3+4i
	fmt.Println(x * y)               //-5+10i
	fmt.Println(real(x * y))         //-5
	fmt.Println(imag(x * y))         //10

	//无穷与非数
	fmt.Println("==== z,Inf,Nan:====")
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //+Inf 正无穷 -Inf负无穷 NaN非数
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan, nan != nan) //false false false true

	//如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败 如：
	/*
		func somefunc()(value float64 , ok bool){
			//.......
			if failed{
				return 0, false
			}
			return result, true
		}
	*/

	//
	println("==== 0 0x :=====")
	a1 := 0666 //0开头是八进制 0x或0X开头是16进制
	a2 := 9
	fmt.Printf("%d %[1]o %#[1]o \n", a1) //[1]是第一个操作数，# 是printf在对%o,%x,%X生成0、0x、0X前缀
	fmt.Printf("%d %[1]o %#[1]o \n", a2) //%d 十进制整数
	a3 := int64(0xdeadbeef)              //%o 八进制整数
	fmt.Printf("%d %[1]x %#[1]x \n", a3) //%x 十六进制整数	//补充:%g是浮点数 %e带指数 %f控制精读

}
