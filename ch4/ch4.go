// universe block:
// aca estan definidas todas las funciones/palabras que usamos,
// lo que significa que son shadowables
package main

//package block:
import (
	//file block:
	"fmt"
	"math/rand"
)

func main() {
	//every pair of {} is a block
	fmt.Println("ch4")
	blocks_ex()
	if_ex()
	for_forwaays()
	labeling_ex()
	switch_ex()
	excercieseses()
}

func excercieseses() {
	fmt.Println("excercieseses()")
	fmt.Println("ex_1()")
	var numeritos []int
	for i := range 100 {
		fmt.Println(i)
		numeritos = append(numeritos, int(rand.Int31n(100)))
	}
	fmt.Println(numeritos)
	fmt.Println("ex_2()")
	for _, s := range numeritos {
		switch {
		case s%3 == 0 && s%2 == 0:
			fmt.Println("Six!")
		case s%2 == 0:
			fmt.Println("Two!")
		case s%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("never mind... ")
		}
	}

	/**
	es un ejercicio de shadowing. 
	la respuesta es nada, 
	y es porque metio una reasignacion dentro del for loop.
	*/

}

func switch_ex() {
	fmt.Println("switch_ex()")
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	//usando break con for y switch...
	//Hay que ponerle un label al for para poder salir desde el case del switch
loop:
	for i := range 10 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	//blank switch
	//te permite mover la validacion a cada case
	//si siempre haces la misma validacion contra distintos valores, pasar al switch trad
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	switch_fizzbuzz()

}

func switch_fizzbuzz() {
	for i := 1; i <= 30; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func labeling_ex() {
	fmt.Println("labeling_ex()")
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

	/*
			outer:
		    for _, outerVal := range outerValues {
		        for _, innerVal := range outerVal {
		            // process innerVal
		            if invalidSituation(innerVal) {
		                continue outer
		            }
		        }
		        // here we have code that runs only when all of the
		        // innerVal values were sucessfully processed
		    }*/

	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v)
	}
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}
}

func for_forwaays() {
	fmt.Println("for_forwaays()")
	//only looping keyword!
	//4 formas de usarlo:

	//Classic C
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Condition-only
	//podemos dejar fuera de la definicion del for el primer
	//y ultimo bloke, dejando solo la def
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
	fmt.Println()
	//infinite
	stop := 0
	for {
		fmt.Println("Hello")
		if stop == 10 {
			break
		}
		stop++
	}
	fmt.Println()
	not_idiomatic_fizzbuzz()
	fmt.Println()
	idiomatic_fizzbuzz()
	fmt.Println()

	//for-range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	fmt.Println()
	for _, v := range evenVals {
		fmt.Println("Ignored i:", v)
	}
	fmt.Println()
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
	fmt.Println()

	m := map[string]int{
		"a": 1,
		"b": 3,
		"c": 4,
		"d": 2,
		"e": 6,
	}

	for i := 0; i < 5; i++ {
		fmt.Println("loop ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	fmt.Println()
	//old school C loop written in idiomatic Go
	for i := range 5 {
		fmt.Println("loop ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	fmt.Println()

	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	fmt.Println("for range copia el valor de lo que le pasas a range.")
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

func idiomatic_fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

func not_idiomatic_fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func if_ex() {
	fmt.Println("if_ex()")
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	if m := rand.Intn(10); m == 0 {
		fmt.Println("That's too low")
	} else if m > 5 {
		fmt.Println("That's too big:", m)
	} else {
		fmt.Println("That's a good number:", m)
	}
	//fmt.Println("m: ",m) //Compilation error: undefined: m

	//Esta forma de definir variables tambien puede shadowear variables
	//n dentro del if esta sombreando a la n anterior
	if n := rand.Intn(10); n == 0 {
		fmt.Println("2That's too low")
	} else if n > 5 {
		fmt.Println("2That's too big:", n)
	} else {
		fmt.Println("2That's a good number:", n)
	}
}

func blocks_ex() {
	fmt.Println("blocks_ex()")
	//"Each place where a declaration occurs is called a block."
	fmt.Println("Shadowing a variable")
	x := 10
	if x > 5 {
		fmt.Println(x)
		//aca adentro, se crea una nueva variable que "pisa" al antiguo x
		x, y := 5, 20
		fmt.Println(x, y)
	}
	//Aca afuera, x tiene el valor de este blocke
	fmt.Println(x)

}
