package main

//导出名
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

//函数定义
func add(x int, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//数值常量
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello, World")
	fmt.Println("Hello, ", rand.Intn(10))
	fmt.Println("Hello, ", math.Pi)
	fmt.Println("Hello, ", add(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y := split(17)
	fmt.Println(x, y)

	//变量
	var java, python bool
	var str, boolean, number = "abc", true, 12
	var d, e int = 1, 2
	fmt.Println(java)
	fmt.Println(python)
	fmt.Println(str)
	fmt.Println(boolean)
	fmt.Println(number)
	fmt.Println(d, e)

	//数据类型
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//类型转换
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	// i := 42
	// f := float64(i)
	// u := uint(f)
	fmt.Println(u)

	//类型推导
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	//常量
	const P = 3.14
	const Q = "abc"
	const W = true
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
