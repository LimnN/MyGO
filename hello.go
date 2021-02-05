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
	var po int
	po = 10
	var p *int
	p = &po
	*p = *p - 1
	fmt.Println(*p)

	// i, j := 42, 2701

	// p := &i         // 指向 i
	// fmt.Println(*p) // 通过指针读取 i 的值
	// *p = 21         // 通过指针设置 i 的值
	// fmt.Println(i)  // 查看 i 的值

	// p = &j         // 指向 j
	// *p = *p / 37   // 通过指针对 j 进行除法运算
	// fmt.Println(j) // 查看 j 的值

	//结构体
	fmt.Println("=======结构体======")
	type Vertex struct {
		X int
		Y int
	}

	var st Vertex
	st = Vertex{3, 4}
	st.X = 9
	np := &st
	fmt.Println(Vertex{1, 2})
	fmt.Println(st.X)
	fmt.Println(np.Y)

	//数组
	fmt.Println("=======数组======")
	var list1 [2]string
	list1[0] = "hello"
	list1[1] = "world"
	list2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(list1)
	fmt.Println(list2)

	//切片
	fmt.Println("=======切片======")
	var slice1 []int = list2[2:3]
	slice2 := list2[3:4]
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1[0] = 9
	slice2[0] = 9
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(list2)

	var slicenil []int
	fmt.Println(slicenil, len(slicenil), cap(slicenil))
	if nil == slicenil {
		fmt.Println("nil slice")
	}

	slice3 := make([]int, 5)
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice4 := make([]int, 0, 5)
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice0 := make([]int, 0, 0)
	slice5 := append(slice0, 1, 2, 3)
	fmt.Println(slice0, len(slice0), cap(slice0))
	fmt.Println(slice5, len(slice5), cap(slice5))

	fmt.Println("++++++++++")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i1 := range pow {
		fmt.Println(i1)
	}

	for _, v1 := range pow {
		fmt.Println(v1)
	}
	fmt.Println("++++++++++")

	// //映射
	// fmt.Println("=======切片======")
	// var m1 map[string]int
	// m1["abc"] = 12
	// fmt.Println(m1["abc"])
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
