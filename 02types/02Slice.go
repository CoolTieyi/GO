package main

import "fmt"

func main() {
	//Slice的元素是间接引用的 一个Slice甚至可以包含自身
	s1 := make([]int, 5)
	s2 := make([]int, 5, 10) //make([]T,len,cap) len目前长度 cap最大容量
	s3 := []int{}
	s4 := []int{1, 2, 3} //切片不必声明长度 因为他是动态的
	fmt.Println(s1, s2, s3, s4)
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[:2])       // [0,2)
	fmt.Println(len(s) == 0) //判断切片是否为空  不应该用s==nil判断
}
