package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		//我们调用 t 的 Errorf 方法打印一条消息并使测试失败。
		//f 表示格式化，它允许我们构建一个字符串，并将值插入占位符值 %q 中。当你的测试失败时，它能够让你清楚是什么错误导致的。
		//这里的 f 对应 float64，.2 表示输出 2 位小数。
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//原版只有測試方形(沒有圓形)
/*
func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
*/

//加入長方形的測試
//不同的包可以有函数名相同的函数。
//所以我们可以在一个新的包里创建函数 Area(Circle)。
//但是感觉有点大才小用了
//
/* func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.16

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
*/

//改使用方法,为新类型定义方法
/* func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.16

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

} */

//引入接口,重構,把got,want,和顯示error的訊息取出來,做成一個func重複使用
/* func TestArea(t *testing.T) {

	//傳入一個Shape 类型
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.16)

	})

}
*/

//引入「表格驱动测试」
func TestArea(t *testing.T) {

	//创建了一个匿名的结构体。我们用含有两个域 shape 和 want 的 []struct 声明了一个结构体切片。
	//然后我们用测试用例去填充这个数组了。
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}

	//傳入一個Shape 类型
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.16)

	})

}
