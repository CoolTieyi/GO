package main

func main() {

	//Scanner
	//法1:
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//input := scanner.Text()
	//fmt.Printf("%q\n", input) //说明这里input是字符串类型的
	//fmt.Printf("%q\n", "123") //%q 字面值 可以是双引号-字符串 也可以是单引号-整数
	//fmt.Printf("%q\n", 123)   //整数的话就是ASCII码转化了

	//法2:
	//因为这scanf()函数实现的时候也是形参操作的, 所以需要我们传入地址 不然我们的赋值只是给形参赋值
	//var scan1 string
	//var scan2 string
	//fmt.Println("请输入s1:")
	//fmt.Scanln(&scan1)
	//fmt.Println("请输入s2:")
	//fmt.Scanln(&scan2)
	//fmt.Println("scan1:", scan1, "scan2", scan2)

	//交换值 实参形参 地址 变量 指针 5年了 今儿才算学到点门道
	//a := 1
	//b := 2
	//fmt.Println("a b ", a, b)
	//ex(&a, &b)
	//fmt.Println("a b ", a, b)

}

//func ex(c, d *int) {
//	var e int //你要改变你得改变地址里的内容啊
//	e = *c    //要不然你改形参对应的地址 不还是没用吗
//	*c = *d   //出了函数啥都没了 对实参没影响
//	*d = e
//}
