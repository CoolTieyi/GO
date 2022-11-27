package main

import (
	"fmt"
	"sort"
)

func main() {
	//Map
	m1 := make(map[string]int) //map[key_type]value_type
	m1["k1"] = 1
	m1["key2"] = 2
	delete(m1, "key2") //delete
	fmt.Println(m1)

	_, ok := m1["keyer"] //查找不存在的key会将value返回为false 可用来判断key是否在map里
	fmt.Println(ok)

	m2 := map[string]int{
		"n0": 20, "n1": 20, "n2": 20, "n3": 20, "n4": 20,
		"n5": 20, "n6": 20, "n7": 20, "n8": 20, "n9 ": 20,
	}

	fmt.Print("直接遍历m2(未排序):") //Map的迭代顺序是不确定的
	for name, age := range m2 {
		fmt.Printf("%s:%d  ", name, age)
	}

	//顺序遍历需要对key排序
	names := make([]string, 0, len(m2)) //这么定义切片更合适 因为容量已知了
	for name, _ := range m2 {
		names = append(names, name)
	}
	sort.Strings(names) //直接sort. 排序
	fmt.Print("\n 排序后：")
	for _, name := range names {
		fmt.Printf("%s:%d ", name, m2[name])
	}
	fmt.Print("\n")
}
