package main

import (
	"fmt"
	"strings"
)

func main() {

	//
	println("====5. strings.Join() : ====")
	s1 := []string{"foo", "bar", "baz"} //+ 连接代价高昂 用strings.Join()
	res := strings.Join(s1, ",")        //Join()是将数组arg1里面的东西，用arg2“”连接，最后返回一个字符串
	fmt.Printf("%s\tresult is %T\n", res, res)

	a := "a"
	b := "b"
	res2 := strings.Join([]string{a, b}, "~~")
	fmt.Println(res2)

	//
	println("==== range []string: ====")
	sli := []string{"gold", "silver", "bronze"}
	for x, _ := range sli {
		fmt.Print(sli[x], "\t")
	}

	//
	println("\n==== range string: ====")
	d := "hello!"
	for x, s := range d {
		fmt.Printf("%d,%c\n", x, s) // x:key s:value
	}
	c := 1
	print(c << 5)
}
