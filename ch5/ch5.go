package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	varadd = func(i int, j int) int { return i + j }

	varsub = func(i int, j int) int { return i - j }

	varmul = func(i int, j int) int { return i * j }

	vardiv = func(i int, j int) int { return i / j }
)

func main() {
	fmt.Println("ch5")
	fmt.Println(divv(150, 15))
	fmt.Println("*******************************")
	simulating_named_and_optional_ex()
	fmt.Println("*******************************")
	variadic_inputs_ex()
	fmt.Println("*******************************")
	multiple_returns_ex()
	fmt.Println("*******************************")
	named_returns_ex()
	fmt.Println("*******************************")
	calculator_ex()
	fmt.Println("*******************************")
	anonymous_ex()
	fmt.Println("*******************************")
	var_func_ex()
	fmt.Println("*******************************")
	closures_ex()
	fmt.Println("*******************************")
	func_as_params_ex()
	fmt.Println("*******************************")
	ret_func_ex()
	fmt.Println("*******************************")
	deferExample()
	fmt.Println("*******************************")
	go_is_call_by_value_ex()
	fmt.Println("*******************************")
	ex_ec()
	fmt.Println("*******************************")
	//defer_ex()
	fmt.Println("*******************************")
}

func ex_ec() {
	fmt.Println("ex_ec()")
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}
	len, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("FileLen: ", len)
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(s1 string) func(string) string {
	return func(s2 string) string {
		return s1 + " " + s2
	}
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func fileLen(fname string) (int, error) {
	len := 0
	file, closer, err := getFile(fname)
	if err != nil {
		return 0, err
	}
	defer closer()
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		len += count
		//os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return len, err

}

func go_is_call_by_value_ex() {
	fmt.Println("go_is_call_by_value_ex()")
	p := person{}
	i := 2
	s := "Hello!"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s2 := []int{1, 2, 3}
	modSlice(s2)
	fmt.Println(s2)
}

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "goodbay"
	p.name = "bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func defer_ex() {
	fmt.Println("defer_ex()")
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	/*
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fmt.Println("defer being called")
		defer f.Close()
		/*
			a function call runs immediately,
			but defer delays the invocation until the surrounding
			function exits.
	*/
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

func ret_func_ex() {
	fmt.Println("ret_func_ex()")
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := range 3 {
		fmt.Println(threeBase(i), twoBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func func_as_params_ex() {
	fmt.Println("func_as_params_ex()")
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

}

func closures_ex() {
	fmt.Println("closures_ex()")
	/*
		Functions declared inside of functions are special; they are closures
		functions declared inside of functions are able to access and modify variables declared in the outer function
	*/
	a := 20
	f := func() {
		fmt.Println(a)
		a++
	}
	for i := range 10 {
		fmt.Println(i)
		f()
		fmt.Println(a)
	}
	b := 15
	t := func() {
		fmt.Println(b)
		b := 3000
		fmt.Println(b)
	}
	t()
	fmt.Println(b)

}

func var_func_ex() {
	fmt.Println("var_func_ex()")
	x := varadd(2, 3)
	fmt.Println(x)
	change_var_add()
	y := varadd(2, 3)
	fmt.Println(y)
}

func change_var_add() {
	varadd = func(i, j int) int { return i + j + j }
}

func anonymous_ex() {
	fmt.Println("anonymous_ex()")
	f := func(j int) {
		fmt.Println("printing", j, "from inside the anon func")
	}
	for i := range 5 {
		f(i)
		func(j int) {
			fmt.Println("printing", j, "from inside the 2nd anon func")
		}(i)
	}

}

func calculator_ex() {
	fmt.Println("calculator_ex()")
	/**
	we can use type to define the signature of funcs
	var opMap = map[string]func(int, int) int
	==
	var opMap = map[string]opFuncCalcType
	*/
	type opFuncCalcType func(int, int) (int, error)

	var opMap = map[string]opFuncCalcType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(result)
	}
}

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("Otra vez 0 pibe!")
	}
	return i / j, nil
}

func named_returns_ex() {
	fmt.Println("named_returns_ex()")
	x, y, z := divNremNamed(16, 3)
	fmt.Println(x, y, z)
	x, y, z = divNremNamed(16, 0)
	fmt.Println(x, y, z)

}

func divNremNamed(num, den int) (res, rem int, err error) {
	if den == 0 {
		err = errors.New("por 0 de vuelta, pibe!")
		return res, rem, err
	}
	res, rem = num/den, num%den
	return res, rem, err
}

func multiple_returns_ex() {
	fmt.Println("multiple_returns_ex()")
	fmt.Println(divNrem(15, 3))
	fmt.Println(divNrem(16, 3))
	fmt.Println(divNrem(15, 0))
	res, rem, err := divNrem(16, 3)
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	fmt.Println(res, rem)
	//Use _ whenever you don’t need to read a value that’s returned by a function.
	res2, _, err2 := divNrem(15, 3)
	if err2 != nil {
		fmt.Printf("%s", err2.Error())
		os.Exit(1)
	}
	fmt.Println(res2)
}

func divNrem(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("no se puede dividir por 0, pibe!")
	}
	return num / denom, num % denom, nil
}

func simulating_named_and_optional_ex() {
	fmt.Println("simulating_named_and_optional_ex()")
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	return nil
}

func divv(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func variadic_inputs_ex() {
	fmt.Println("variadic_inputs_ex()")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
	//cuando le paso un slice tengo que "desarmarlo" usando '...'

}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
