package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	x, y float64
}

func Distance(p Point, q Point) float64 { //function
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//Distance p是方法的接收器 receiver
func (p Point) Distance(q Point) float64 { //method
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//ScaleBy 基于指针对象的方法
func (p *Point) ScaleBy(factor float64) { //方法名为(*point).ScaleBy
	p.x *= factor
	p.y *= factor
}

func main() {

	//函数名字前加一个变量 => 方法
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 函数调用
	fmt.Println(p.Distance(q))  // 方法调用		//"p.Distance"叫做选择器

	//Pointer 基于指针对象的方法
	//接收器是指针 那一定要用指针调用方法
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	s := Point{1, 2}
	(&s).ScaleBy(3)
	fmt.Println(s)

	s.ScaleBy(2) // 只能变量
	fmt.Println(s)

	/*
		ColoredPoint并不是一个Point 但他"has a"Point
		是不是只要是Point类型的就可以调用Distance()和scaleBy()方法啊
		所以这个接收器才很有意义 对 就算在另外的函数体里面也行
	*/

	//Struct Embedding
	red := color.RGBA{255, 0, 0, 255}
	var u = ColoredPoint{Point{1, 2}, Point{3, 4}, red}
	u.a.ScaleBy(5) //u是Color类型 但是却可以用ScaleBy() 是因为CPstruct里面包含Point 而ScaleBy()和Distance()的接收器是Point
	fmt.Println(u)

	//MethodValue
	//甚至可以再声明接收器的时候不传值
	sp := Point{1, 2}

	scalep := sp.ScaleBy
	scalep(2)
	scalep(3)
	scalep(4)
	fmt.Println(sp) //{1*2*3*4,2*2*3*4}

	//Method Expression
	medistance := Point.Distance  //method expression
	fmt.Println(medistance(p, q)) //会将第一个参数用作接收器 所以可以用不写选择器的方式来对齐进行调用
	fmt.Printf("%T\n", medistance)

	mescale := (*Point).ScaleBy //method expression
	mescale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", mescale)
	/*
		这个medistance实际上制定了Point对象为接收器的一个方法 func (p Point) Distance(),
		但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数
		即其需要用第一个额外参数指定接收器 后面排列Distance方法的参数
		=== 看起来本书中函数和方法的区别就是指有没有接收器 而不像其他语言一样值有没有返回值 ===
	*/

}

type ColoredPoint struct {
	a     Point
	b     Point
	color color.RGBA
}

/*
ColoredPoint
一个ColoredPoint并不是一个Point，但他"has a"Point，并且它有从Point类里引入的Distance和ScaleBy方法

func (p ColoredPoint) Distance(q Point) float64 {	//接收器值为p.Point
	return p.Point.Distance(q)
}
func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}
*/

type Path []Point

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y + q.y}
}

//TranslateBy 甚至还可以这么用 这go对方法的调用也太随便了
/*
	当根据变量来决定调用一个类型的哪个函数时 方法表达式就很有用的
	你可以根据选择来调用接收器不同的方法 下例op可以有add 或 sub
	Path.TranslateBy 会为其Path中的每一个Point来调用对应的方法
*/
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
