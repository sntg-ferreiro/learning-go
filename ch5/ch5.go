package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ch5")
	fmt.Println(div(150, 15))
	simulating_named_and_optional_ex()
	variadic_inputs_ex()
	multiple_returns_ex()
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

func div(numerator int, denominator int) int {
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
