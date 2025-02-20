package main

import "fmt"

const konst_x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	var x int = 10
	x *= 2

	fmt.Println(x)

	var y int = 20

	fmt.Println(x == y)

	var x2 int = 10
	var y2 float64 = 30.2
	var sum1 float64 = float64(x2) + y2
	var sum2 int = x2 + int(y2)

	fmt.Println(sum1, sum2)

	//Assign with no specific type
	var noTypeInt = 10

	//Assign zero value
	var zeroValue int

	fmt.Println(noTypeInt, zeroValue)

	//multiple assignments of same type
	var a, b, c int = 1, 2, 3

	fmt.Println(a + b + c)

	//zero values
	var d, e int
	fmt.Println(d, e)

	//diff types
	var f, g = 100, "STRINGS!"
	fmt.Println(f, g)

	//declaration list:
	var (
		h    int
		i        = 20
		j    int = 30
		k, l     = 40, "otra string"
		m, n string
	)
	fmt.Println(h, i, j, k, l, m, n)

	//short declaration
	x3 := 10
	x3, y3 := 30, "ooootra striiing"
	fmt.Println(x3, y3)
	//':=' no es legal fuera de func

	const konst_y = "Hello"

	fmt.Println(konst_x)
	fmt.Println(konst_y)

	//konst_x = konst_x + 1 wont compile
	//konst_y = "bbye" wont compile

	fmt.Println(konst_x)
	fmt.Println(konst_y)
	ex1()
	ex2()
	ex3()
}

func ex1() {
	i := 20
	var f float64 = float64(i)
	fmt.Println(i, f)
}

func ex2() {
	const value = 20
	var i int = value
	var f float64 = value

	fmt.Println(i, f)
}

func ex3() {
	var b byte
	var smallI int32
	var bigI uint64

	smallI = 2147483647
	bigI = 18446744073709551615
	b = 255

	fmt.Println(b, smallI, bigI)
	b++
	smallI++
	bigI++
	fmt.Println(b, smallI, bigI)
}
