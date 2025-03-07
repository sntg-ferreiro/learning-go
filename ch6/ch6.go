package main

import (
	"fmt"
)

func main() {
	fmt.Println("ch6")
	fmt.Println("************************************")
	primer_ex()
	fmt.Println("************************************")
	struct_ex()
	fmt.Println("************************************")
	update_ex()
	fmt.Println("************************************")
	makeFoo_ex()
	fmt.Println("************************************")
	main_excercises_ex()
	fmt.Println("************************************")
}

func main_excercises_ex() {
	fmt.Println("main_excercises_ex()")
	p1 := makePerson("bob", "sager", 40)
	pp := makePersonPointer(p1.fName, p1.lName, p1.Age)
	fmt.Println("p1:", p1, "pp:", pp)
	ss := []string{"uno", "dos", "tres", "cuatro"}

	s := "cinco"
	fmt.Println("Before -> ss:", ss, "s:", s)
	updateSlice(s, ss)
	fmt.Println("Update -> ss:", ss, "s:", s)
	growSlice(s, ss)
	fmt.Println("Grow -> ss:", ss, "s:", s)
}

func updateSlice(s string, ss []string) {
	ss[len(ss)-1] = s
	fmt.Println("Update: ", ss)
}

func growSlice(s string, ss []string) {
	ss = append(ss, s)
	fmt.Println("Grow: ", ss)
}

func makePersonPointer(fname, lname string, age int) *Person {
	p := Person{
		fName: fname,
		lName: lname,
		Age:   age,
	}
	return &p
}

func makePerson(fname, lname string, age int) Person {
	p := Person{
		fName: fname,
		lName: lname,
		Age:   age,
	}
	return p
}

type Person struct {
	fName string
	lName string
	Age   int
}

func makeFoo_ex() {
	fmt.Println("makeFoo_ex()")
	//var f *Foo
	//DontMakeFoo(f)
	f2, err := MakeFoo()
	//fmt.Println("Dont make foo:", *f) //panic
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("Make foo:", f2)
	}
}

func DontMakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func update_ex() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func struct_ex() {
	fmt.Println("struct_ex()")
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p)
}

func makePointer[T any](t T) *T {
	return &t
}

func primer_ex() {
	fmt.Println("primer_ex()")
	var x int32 = 10
	var y bool = true
	//'&':
	pointerX := &x //adress operator: "la direccion de..."
	pointerY := &y
	//'*':
	//indirection operator: "esta variable es un pointer a... "
	// O
	// returns the "pointed to" value
	var pointerZ *string

	fmt.Println(x, y, pointerX, pointerY, pointerZ)
	fmt.Println("x: ", x, "pointerX: ", pointerX)
	fmt.Println("x: ", x, "*pointerX: ", *pointerX)

	var nilPointer *int
	fmt.Println(nilPointer == nil) //true
	//fmt.Println(*nilPointer)       //panic
}
