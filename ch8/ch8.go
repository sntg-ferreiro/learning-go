package main

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("|******|")
	fmt.Println("| Ch8! |")
	fmt.Println("|******|")
	stack_ex()
	fmt.Println("|******|")
	kustom_map_ex()
	fmt.Println("|******|")
	generics_interfaces_ex()
	fmt.Println("|******|")
	div_rem_ex()
	fmt.Println("|******|")
	binary_tree_ex()
	fmt.Println("|******|")
	pothole_ex()
	fmt.Println("|******|")
	ex_ex()
}

func ex_ex() {
	fmt.Println("ex_ex()")
	fmt.Println("ex_ex(1)")
	var dd int = 40
	fmt.Println(DobleDoble(dd))
	var df float64 = 5.0 / 3.0
	fmt.Println(df)
	fmt.Println(DobleDoble(df))
	fmt.Println("ex_ex(2)")
	var pp MyPrintable = 10
	PrettyPrint(pp)
	fmt.Println("ex_ex(3)")
	var ls *List[int] = NewList[int]()
	ls.Add(10)
	ls.Add(11)
	ls.Add(12)
	ls.Add(13)
	ls.Add(14)

	fmt.Println(ls.String())
	fmt.Println("index(14):", ls.Index(14))

	ls.Insert(15, 0)
	ls.Insert(100, 20)
	ls.Insert(99, 19)
	ls.Insert(101, 20)
	ls.Insert(98, 10)
	fmt.Println(ls.String())
	fmt.Println("index(15):", ls.Index(15))
	fmt.Println("ls.size:", ls.size)

}

type List[T comparable] struct {
	head *ElemList[T]
	size int
}

func (l *List[T]) Insert(val T, idx int) {
	l.head.Insert(val, idx)
	l.size++
}

func (l List[T]) String() string {
	return fmt.Sprintf("[%v", l.head.String())
}

func NewList[T comparable]() *List[T] {
	return &List[T]{head: nil}
}

func (l *List[T]) Add(val T) {
	if l.head == nil {
		l.head = &ElemList[T]{val: val, next: nil, index: 0}
	} else {
		l.head.Add(val)
	}
	l.size++
}

func (l *List[T]) Index(val T) int {
	return l.head.Index(val)
}

type ElemList[T comparable] struct {
	val   T
	index int
	next  *ElemList[T]
}

func (e *ElemList[T]) Insert(val T, idx int) {
	if e.index == idx {
		//nuevo valor en este idx, Add(viejo valor)?
		e.Add(e.val)
		e.val = val
	} else if e.index < idx {
		if e.next == nil {
			e.next = &ElemList[T]{val: val, index: idx, next: nil}
		} else {
			e.next.Insert(val, idx)
		}
	} else {
		e.next = &ElemList[T]{val: e.val, index: e.index, next: e.next}
		e.val = val
		e.index = idx
	}
}

func (e ElemList[T]) String() string {
	out := ""
	if e.next == nil {
		out = fmt.Sprintf("%v:%d]", e.val, e.index)
	} else {
		out = fmt.Sprintf("%v:%d,%v", e.val, e.index, e.next.String())
	}
	return out
}

func (e *ElemList[T]) Index(val T) int {
	if e.val == val {
		return e.index
	}
	if e.next == nil {
		return -1
	}
	return e.next.Index(val)
}

func (e *ElemList[T]) Add(val T) {
	if e.next == nil {
		e.next = &ElemList[T]{val: val, next: nil, index: e.index + 1}
		return
	}
	e.next.Add(val)
}

type Printable interface {
	fmt.Stringer
	~int
}

type MyPrintable int

func (mp MyPrintable) String() string {
	return fmt.Sprintf("%d", mp)
}

func PrettyPrint[T Printable](val T) {
	fmt.Println(val.String())
}

type Doublerer interface {
	~int | ~float32 | ~float64
}

func DobleDoble[T Doublerer](val T) T {
	return (val + val) * (val + val)
}

func pothole_ex() {
	fmt.Println("pothole_ex()")
	var a int = 10
	var b int = 10
	Comparer(a, b)

	var a2 ThingerInt = 10
	var b2 ThingerInt = 10
	Comparer(a2, b2)

	/*
		var a3 ThingerSlice = []int{1, 2, 3}
		var b3 ThingerSlice = []int{1, 2, 3}
			Comparer(a3,b3)// compile fails: "ThingerSlice does not satisfy comparable"
	*/

	var a4 Thinger = a2
	var b4 Thinger = b2
	Comparer(a4, b4)
	/*
		var a5 Thinger = a3
		var b5 Thinger = b3
		Comparer(a5, b5) //Compiles, panic
	*/

}

type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("Equal!")
	}
}

func binary_tree_ex() {
	fmt.Println("binary_tree_ex()")
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	t1.Add(100)
	fmt.Println(t1.Contains(10))
	fmt.Println(t1.Contains(40))

	fmt.Println("structs 1:")
	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30}))
	fmt.Println(t2.Contains(Person{"Fred", 25}))
	//fmt.Println(t2.ToString())

	fmt.Println("structs 2:")
	t3 := NewTree(Person.Order)
	t3.Add(Person{"Bob", 30})
	t3.Add(Person{"Maria", 35})
	t3.Add(Person{"Bob", 50})
	//fmt.Println(t3.ToString())
	fmt.Println(t3.Contains(Person{"Bob", 30}))
	fmt.Println(t3.Contains(Person{"Fred", 25}))
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("{Name: %s, Age: %d}", p.Name, p.Age)
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

/*
	func (t Tree[T]) ToString() string {
		return t.root.ToString()
	}
*/
func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

/*
	func (n *Node[T]) ToString() string {
		if n == nil {
			return ""
		}
		return n.left.ToString() + "<-" + n.val.String() + "->" + n.right.ToString()
	}
*/
func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

func div_rem_ex() {
	fmt.Println("div_rem_ex()")
	var a, b int
	a = 10
	b = 4
	var ua uint = 18_446_744_073_709_551_615
	var ub uint = 9_223_372_036_854_775_808
	fmt.Println(divAndRem(ua, ub))
	fmt.Println(divAndRem(a, b))
}

type Integerer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func divAndRem[T Integerer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("0000")
	}
	return num / denom, num % denom, nil
}

func generics_interfaces_ex() {
	fmt.Println("generics_interfaces_ex()")
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindKloser(pair2Da, pair2Db)
	fmt.Println(closer)
	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindKloser(pair3Da, pair3Db)
	fmt.Println(closer2)

}

type Pair[T fmt.Stringer] struct {
	A T
	B T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

type Point2D struct {
	X, Y int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}
func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func FindKloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.A.Diff(pair1.B)
	d2 := pair2.A.Diff(pair2.B)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

func kustom_map_ex() {
	fmt.Println("kustom_map_ex()")
	var sints []int
	sints = append(sints, 14)
	sints = append(sints, 12)
	sints = append(sints, 13)
	sints = append(sints, 11)
	sints = append(sints, 10)

	mapres := KMap(sints, func(v int) string {
		return strconv.Itoa(v) + "algo"
	})
	fmt.Println(mapres)
	redres := KReduce(sints, "", func(ival int, sval string) string {
		return strconv.Itoa(ival) + sval
	})
	fmt.Println(redres)
	filres := KFilter(sints, func(val int) bool {
		return val <= 12
	})
	fmt.Println(filres)
}

func KFilter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// VAL, ACC
func KReduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}

func KMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func stack_ex() {
	fmt.Println("stack_ex()")

	var intStack Stack[int]

	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack)
	fmt.Println("Contains(5)", intStack.Contains(5))
	fmt.Println("Contains(3)", intStack.Contains(3))
	for i := range intStack.vals {
		v, ok := intStack.Pop()
		if ok {
			fmt.Println("POP!:", v)
		} else {
			fmt.Println("err on Pop!", i)
		}
	}
	fmt.Println("Contains(3)", intStack.Contains(3))

}

type Stack[T comparable] struct {
	vals []T
}

// Con Stack[T any] --> No compila porque no sabemos si T es comparable
func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val { //
			return true
		}
	}
	return false
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}
