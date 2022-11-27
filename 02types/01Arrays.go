package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a1 := [3]int{}
	a2 := [3]int{1, 2, 3}
	a3 := [...]int{1} //数组一定要声明长度或...
	sa := a2[1:2]
	fmt.Println(sa, "切片可以是简介引用的数组 但是数目也固定了")
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "?", GBP: "?", RMB: " ￥"} //只有常量可以这么玩
	fmt.Print(a1[0], a2, a3, RMB, symbol[RMB])                     // 1 [1 2 3] [1] 3 ￥, 3是iota记得数

	r := [...]int{3, 99: 1, 2} //第一个为2 第99个为1 第100为1 中间的初始值0
	fmt.Println("\n", r[0], r[98], r[99], r[100], r[96])

	//Sum256对任意字节的slice类型 生成消息摘要(256bit) 所以对应[32]byte
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("c1:%x\nc2:%x\nc1==c2:%t\nType(c1)=%T", c1, c2, c1 == c2, c1) // %x十六进制打印数组或slice全部元素 %t打印布尔型  %T显示数据类型
}
