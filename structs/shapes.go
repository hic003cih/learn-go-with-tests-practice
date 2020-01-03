package main

import "math"

//使用保留字 struct 来定义自己的类型
//一个通过 struct 定义出来的类型是一些已命名的域的集合，这些域用来保存数据。
//封装长方形的信息
/* type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
} */

//給類型加一些方法
type Rectangle struct {
	Width  float64
	Height float64
}

//方法接收者的语法 func(receiverName ReceiverType) MethodName(args)
//当方法被这种类型的变量调用时，数据的引用通过变量 receiverName 获得。在其他许多编程语言中这些被隐藏起来并且通过 this 来获得接收者。
//把类型的第一个字母作为接收者变量是 Go 语言的一个惯例。
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//定義Shape類型(interface)
type Shape interface {
	Area() float64
}

/* //算周長
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// 算面積
func Area(width float64, height float64) float64 {
	return width * height
}
*/

//使用struct的版本
//算周長
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// 算面積
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
