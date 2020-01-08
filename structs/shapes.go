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
//声明结构体以创建我们自己的类型，让我们把数据集合在一起并达到简化代码的目地
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

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//定義Shape類型(interface)
//声明接口，这样我们可以定义适合不同参数类型的函数（参数多态）
//接口是把负责从系统的其他部分隐藏起来的伟大工具。在我们的测试中，
//辅助函数的代码不需要知道具体的几何形状，只需要知道获取它的面积即可。
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
