package main

//导出名
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
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
	//包、导出名、函数
	fmt.Println("=======包、导出名、函数======")
	fmt.Println("Hello, World")
	fmt.Println("Hello, ", rand.Intn(10))
	fmt.Println("Hello, ", math.Pi)
	fmt.Println("Hello, ", add(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y := split(17)
	fmt.Println(x, y)

	//变量
	fmt.Println("=======变量======")
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
	fmt.Println("=======数据类型======")
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//类型转换
	fmt.Println("=======类型转换======")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	// i := 42
	// f := float64(i)
	// u := uint(f)
	fmt.Println(u)

	//类型推导
	fmt.Println("=======类型推导======")
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	//常量
	fmt.Println("=======常量======")
	const P = 3.14
	const Q = "abc"
	const W = true
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//for循环
	fmt.Println("=======for======")
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 20 {
		sum += sum
	}
	fmt.Println(sum)

	// for {

	// }

	//if
	fmt.Println("=======if======")
	if sum > 30 {
		fmt.Println("YES")
	}

	if xx := math.Pow(1, 2); xx < 20 {
		fmt.Println("xx is ", xx)
	} else {
		fmt.Println("HAHA")
	}
	// fmt.Println(xx)

	//练习 循环与函数
	fmt.Println("=======练习 循环与函数======")
	fmt.Println(mySqrt(2, false))
	fmt.Println(mySqrt(2, true))

	//switch
	fmt.Println("=======switch======")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows")
		fallthrough
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	switch x {
	case 0:
		fmt.Println("x is 0")
	case int(mySqrt(4, true)):
		fmt.Println("x is ", x)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//defer
	fmt.Println("=======defer======")
	defer fmt.Println("Final")
	fmt.Println("Not Final")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Not Final 2")

	//指针
	fmt.Println("=======指针======")
	// var po int
	// po = 10
	// var p *int
	// p = &po
	// *p = *p - 1
	// fmt.Println(po)

	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

func mySqrt(x float64, isTimes bool) float64 {
	z := 1.0
	if isTimes {
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
			fmt.Println("z is ", z)
		}
		return z
	}
	temp := x
	loops := 0
	for math.Abs(z-temp) > 0.00000001 {
		temp = z
		z -= (z*z - x) / (2 * z)
		loops++
		fmt.Println("z is ", z)
	}
	fmt.Println("loop times is ", loops)
	return z
}
