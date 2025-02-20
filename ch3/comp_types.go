package main

import (
	"fmt"
	"slices"
)

func main() {

	var x [3]int
	//{0,0,0}

	var x2 = [3]int{10, 20, 30}

	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(x, x2, x3)

	//matrices:
	var m [2][3]int
	//access:
	m[0][2] = 1

	fmt.Println(m)

	fmt.Println(m[0][2])

	//len:
	fmt.Println(len(x3))

	slices_ex()
	func_ex()
	capacity_ex()
	make_ex()
	slicing_ex()
	copy_ex()
	strings_runes_bytes()
	maps_ex()
	set_ex()
	structs_ex()
	ch3_ex()

}

func ch3_ex() {
	fmt.Println("ch3_ex()")

	fmt.Println("ch3_ex_1()")
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	s1 := greetings[:2]
	s2 := greetings[1:4]
	s3 := greetings[3:]

	fmt.Println("greetings: ", greetings, " -- s1: ", s1, " -- s2: ", s2, " -- s3: ", s3)
	fmt.Println("ch3_ex_2()")
	message := "+Hi ߑ頡nd ߑ被"
	fmt.Println("message: ", message)
	var fourthRune string = message[3:4]
	fmt.Println("4thrune: ", fourthRune, " **")
	fmt.Println("ch3_ex_3()")
	type employee struct {
		firstName string
		lastName  string
		id        int
	}
	sanMartin := employee{
		firstName: "Jose",
		lastName:  "de San Martin",
		id:        0,
	}
	rosas := employee{"Juan Manuel", "de Rosas", 1}
	var quiroga employee
	quiroga.firstName = "Facundo"
	quiroga.lastName = "Quiroga"
	quiroga.id = 2
	fmt.Println("sanMartin: ", sanMartin)
	fmt.Println("rosas: ", rosas)
	fmt.Println("quiroga: ", quiroga)

}

func structs_ex() {
	fmt.Println("structs_ex()")
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}

	fmt.Println("fred: ", fred, " bob: ", bob)

	julia := person{
		"julia", 40, "Kat",
	}
	beth := person{
		age:  30,
		name: "Beth",
	}
	bob.name = "Bob"
	fmt.Println("julia: ", julia, " beth: ", beth)
	fmt.Println("fred: ", fred, " bob: ", bob)

	var aperson struct {
		name string
		age  int
		pet  string
	}

	aperson.name = "bob"
	aperson.age = 50
	aperson.pet = "hund"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println("pet: ", pet)
	fmt.Println("aperson: ", aperson)

	/*
		type nombre struct{... campos ...}
		es mas parecido a una clase que hace falta instanciar
		var nombre struct{... campos ...}
		es como una def e instanciacion de una de esa estructura
	*/

	/*
		structs son comparables si y solo si todos sus campos son comparables
	*/

}

func set_ex() {
	fmt.Println("set_ex()")
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	/*
		We wrote 11 values into intSet, but the length of intSet is 8,
		because you cannot have duplicate keys in a map
	*/
}

func maps_ex() {
	fmt.Println("maps_ex")
	fmt.Println("map[KeyType]ValueType")
	var nilMap map[string]int
	fmt.Println(len(nilMap))

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println("teams: ", teams)

	ages := make(map[int][]string, 10)
	fmt.Println("ages: ", ages)
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	fmt.Println("totalwins: ", totalWins)
	/*
		map1 == map2 o map1 != map2 => no existe
			The key for a map can be any comparable type.
			This means you cannot use a slice or a map as the key for a map
	*/

	//the comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	/**
	the comma ok idiom significa que cuando busco un elemnto en un mapa y
	este esta presente, la funcion de 'get' devuelve el valor y true en ok
	O
	devuelve el valor 0 del type y false en ok.
	*/

	delete(m, "hello")
	fmt.Println("m: ", m)

	m["foo"] = 1
	m["bar"] = 2
	fmt.Println("m: ", m)
	clear(m)
	fmt.Println("m: ", m)

}

func strings_runes_bytes() {
	fmt.Println("strings_runes_bytes")
	var s string = "Hello there"
	var b byte = s[6]
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println("s:", s, "b:", b, "s2:", s2, "s3:", s3, "s4:", s4)
	//b vale 116, el valor de t en UTF-8
}

func copy_ex() {
	fmt.Println("copy_ex")
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println("x:", x, "y:", y, "num:", num)
	//Num es el numero de elem copiados. Y es el slice destino, x el og
	z := make([]int, 2)
	copy(z, x[2:])
	fmt.Println("x:", x, "z:", z)
	x = []int{1, 2, 3, 4}
	num = copy(x[:3], x[1:])
	fmt.Println(x, num)
	// Cambia los valores sobre X

}

func slicing_ex() {
	fmt.Println("slicing_ex")
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	/*
		When you take a slice from a slice,
		you are not making a copy of the data.
		Instead, you now have two variables that are sharing memory.

		x and e point to the same memory loc.
	*/

	e = append(e, "EEEEE")
	fmt.Println("e:", e)
	fmt.Println("x:", x)

	x2 := []string{"a", "b", "c", "d"}
	y2 := x2[:2]
	z2 := x2[1:]
	x2[1] = "y2"
	y2[0] = "x2"
	z2[1] = "z2"
	fmt.Println("x2:", x2)
	fmt.Println("y2:", y2)
	fmt.Println("z2:", z2)
}

func make_ex() {
	fmt.Println("Make")
	x := make([]int, 5)
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 5, 10)

	fmt.Println(y, len(y), cap(y))

	z := make([]int, 0, 10)
	z = append(z, 5, 6, 7, 8)
	fmt.Println(z)
}

func capacity_ex() {
	fmt.Println("Capacity")
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	fmt.Println("Emptying A Slice")
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))
}

func func_ex() {
	ss := []string{"a", "b", "c"}
	ss1 := []string{"d", "e", "f"}

	fmt.Println(len(ss))

	ss = append(ss, "appended")
	fmt.Println(len(ss))

	ss3 := append(ss, ss1...)
	fmt.Println(ss3)

	/**
	Go is a call by value language.
	Every time you pass a parameter to a function,
	Go makes a copy of the value that’s passed in.
	Passing a slice to the append function actually
	passes a copy of the slice to the function.
	The function adds the values to the copy of the slice
	and returns the copy. You then assign the returned
	slice back to the variable in the calling function.
	*/
}

func slices_ex() {
	fmt.Println("Slices parecen ser el reemplazo del ArrayList clasico de JAVA!")
	var s = []int{10, 20, 30}
	var s1 = []int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(s, s1)

	//todo igual...

	var s2 []int
	fmt.Println(s2, s2 == nil)

	s3 := []int{1, 2, 3, 4, 5}
	s4 := []int{1, 2, 3, 4, 5}
	s5 := []int{1, 2, 3, 4, 5, 6}
	ss := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(s3, s4))
	fmt.Println(slices.Equal(s3, s5))
	//fmt.Println(slices.Equal(s3, ss)) not compile
	fmt.Println(ss)
}
