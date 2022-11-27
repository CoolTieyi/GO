package main

import "fmt"

//func exchange(a, b *int) {
//	var c int
//	c = *a
//	*a = *b
//	*b = c
//}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fs1 := f()
	fmt.Println(fs1)
	fs2 := f2()
	fmt.Println(fs2)
	fs3 := f3()
	fmt.Println(fs3)

	//
	//var name2 string
	//var num float64
	//fmt.Scanf("%s\n%f", &name2, &num) //输入时严格保证“%s\n%f”格式
	//fmt.Printf("姓名%s\n学号%.0f", name2, num)
	//
	//var scan1 string
	//var scan2 string
	//fmt.Println("&scan1:", &scan1)
	//fmt.Println("请输入s1:")
	//fmt.Scanln(scan1)
	//fmt.Println("请输入s2:")
	//fmt.Println("&scan2:", &scan2)
	//fmt.Scanf(scan2)
	//fmt.Println(scan1, scan2)
	//a := 1
	//b := 2
	//fmt.Println("a b ", a, b)
	//ex(&a, &b)
	//fmt.Println("a b ", a, b)

	//var t int
	//println("&t", &t)
	//t = a
	//println("&t t", &t, t)
	//
	//println("a b: ", a, b)
	//println("&a &b: ", &a, &b)
	//pa := &a
	//pa = &b
	////&a = &b	//不可以
	//println(pa)
	//println(*pa)
	//println(a)
	//println(&a)

	//a = 1
	//println("&a &b: ", &a, &b)
	//b = 2
	//println("&a &b: ", &a, &b)
	//var ptra *int = (*int)0x454546321	//手动指定地址应该不成
	//println(ptra)
	//println("a b: ", a, b)
	//println("&a &b: ", &a, &b)
	//ptra := &a
	//println("ptra: ", ptra)
	//println("&ptra: ", &ptra)
	//ptrb := &b
	//println("ptrb: ", ptrb)
	//println("&ptrb: ", &ptrb)
	//var t int
	//println("&t", &t)
	//t = a
	//println("&t", &t)
	//a = b
	//b = t
	//println("a b t: ", a, b, t)
	//println("&a &b &t: ", &a, &b, &t)
}
