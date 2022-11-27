package main

import "fmt"

func main() {
	type student struct {
		name string
		Num  float64 //大写可以导出
	}
	var stu student
	stu.name = "tt"
	var ptrnum *float64 = &stu.Num
	*ptrnum = 210381
	//也可以点操作符和指针一起工作 指针可以直接用点操作符
	var stu2 *student = &stu //stu2是stu的指针
	stu2.name += "yy"        // 等于 (*stu2).name += "yy"
	fmt.Println("stu: ", stu)
	//Notice: 	一个名为S的结构体 不能再包含S类型的成员: 因为一个聚合的值不能包含它自身
	//			但! S类型的结构体可以包含*S指针类型的成员 例如下面的插入排序(我还没看懂呢 慢慢看)

	//type tree struct {
	//	value       int
	//	left, right *tree
	//}
	//
	//// Sort sorts values in place.
	//func Sort(values []int) {
	//	var root *tree
	//	for _, v := range values {
	//		root = add(root, v)
	//	}
	//	appendValues(values[:0], root)
	//}
	//
	//// appendValues appends the elements of t to values in order
	//// and returns the resulting slice.
	//func appendValues(values []int, t *tree) []int {
	//	if t != nil {
	//	values = appendValues(values, t.left)
	//	values = append(values, t.value)
	//	values = appendValues(values, t.right)
	//}
	//	return values
	//}
	//
	//func add(t *tree, value int) *tree {
	//	if t == nil {
	//	// Equivalent to return &tree{value: value}.
	//	t = new(tree)
	//	t.value = value
	//	return t
	//}
	//	if value < t.value {
	//	t.left = add(t.left, value)
	//} else {
	//	t.right = add(t.right, value)
	//}
	//	return t
	//}
}
