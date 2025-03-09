package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

func main() {
	fmt.Println("ch7!")
	fmt.Println("*********************************")
	methods_ex()
	fmt.Println("*********************************")
	counter_ex()
	fmt.Println("*********************************")
	intTree_ex()
	fmt.Println("*********************************")
	adder_ex()
	fmt.Println("*********************************")
	iota_ex()
	fmt.Println("*********************************")
	embed_ex()
	fmt.Println("*********************************")
	understanding_implicit_interfaces_ex()
	fmt.Println("*********************************")
	types_assertions_switches_ex()
	fmt.Println("*********************************")
	greeting_service_ex()
	fmt.Println("*********************************")
	ex_ex()

}

func ex_ex() {
	fmt.Println("ex_ex()")
	t1 := NewTeam("SanLo", []string{"fred", "fred1", "fred2"})
	t3 := NewTeam("Bover", []string{"Bob", "Bob1", "Bob2"})
	t4 := NewTeam("Rica", []string{"Pat", "Pat1", "Pat2"})

	l1 := NewLeague()
	l1.addTeam(t1)
	l1.addTeam(t4)
	l1.addTeam(t3)

	l1.MatchResult(t1.Name, t3.Name, 2, 0)
	l1.MatchResult(t1.Name, t3.Name, 2, 0)
	l1.MatchResult(t1.Name, t4.Name, 2, 0)
	l1.MatchResult(t3.Name, t4.Name, 2, 0)
	l1.MatchResult(t3.Name, t4.Name, 2, 0)
	l1.MatchResult(t3.Name, t4.Name, 2, 0)
	l1.MatchResult(t3.Name, t4.Name, 2, 0)
	l1.MatchResult(t3.Name, t4.Name, 0, 0)
	l1.MatchResult(t3.Name, t4.Name, 0, 1)

	ranks := l1.Ranking()
	fmt.Println(l1)
	fmt.Println(ranks)

	writer := os.Stdout
	RankPrinter(l1, writer)

}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	for _, v := range ranking {
		io.WriteString(w, v)
		//io.WriteString(w, "\n")
		w.Write([]byte("\n"))
	}
}

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(t1, t2 string, scoreTeam1, scoreTeam2 int) {
	if _, ok := l.Teams[t1]; !ok {
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		return
	}
	if scoreTeam1 == scoreTeam2 {
		return
	}
	if scoreTeam1 > scoreTeam2 {
		l.addWinToTeam(t1)
	} else if scoreTeam2 > scoreTeam1 {
		l.addWinToTeam(t2)
	}
}

func (l League) Ranking() []string {
	keys := make([]string, 0, len(l.Wins))
	for k := range l.Wins {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return l.Wins[keys[i]] > l.Wins[keys[j]]
	})
	return keys
}

func (l *League) addTeam(t Team) {
	l.Teams[t.Name] = t
}

func (l *League) addWinToTeam(t string) {
	l.Wins[t]++
}

func NewTeam(n string, ps []string) Team {
	return Team{
		Name:    n,
		Players: ps,
	}
}

func NewLeague() League {
	return League{
		Teams: map[string]Team{},
		Wins:  map[string]int{},
	}
}

/**
EjerciciosEjerciciosEjerciciosEjerciciosEjerciciosEjerciciosEjerciciosEjerciciosEjercicios
*/

func greeting_service_ex() {
	fmt.Println("greeting_service_ex()")
	l := LoggerAdapter(LogOutPut2)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	//http.ListenAndServe(":8080", nil)
	//descomentar lo de arriba para probar el servicio http
}

func LogOutput(msg string) {
	fmt.Println(msg)
}

func LogOutPut2(msg string) {
	fmt.Println("this is the second logger >", msg)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
			"4": "Bob",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type GreetingLogic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic GreetingLogic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	msg, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(msg))
}

func NewController(l Logger, logic GreetingLogic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

/**
CREANDO UN APIREST ARRIBA,
DOC DEL CAPI ABAJO
*/

func types_assertions_switches_ex() {
	fmt.Println("types_assertions_switches_ex()")
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // << -- !! type assertion !!
	fmt.Println(i2 + 1)
	//i3 := i.(string) // PanicS!
	//fmt.Println(i3) // PanicS!
	//i4 := i.(int) // PanicS!
	//fmt.Println(i4 + 2) // PanicS!
	i5, ok := i.(int)
	if !ok {
		fmt.Println(fmt.Errorf("unexpected type: %v", i5))
	} else {
		fmt.Println(i5 + 2)
	}
	doThings(i) // << -- !! Type Switch !!
}

func doThings(i any) {
	switch j := i.(type) {
	case nil:
	// i is nil, type of j is any
	case int:
		// j is of type int
		fmt.Println("// j is of type int")
	case MyInt:
		// j is of type MyInt
		fmt.Println("// j is of type MyInt")
	case io.Reader:
	// j is of type io.Reader
	case string:
	// j is a string
	case bool, rune:
	// i is either a bool or rune, so j is of type any
	default:
		// no idea what i is, so j is of type any
		fmt.Println(j)
	}
}

type MyInt int

func understanding_implicit_interfaces_ex() {
	fmt.Println("understanding_implicit_interfaces_ex()")
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
	c2 := Client{
		L: LogicProvider2{Datita: "trogolidita"},
	}
	c2.Program()
	fmt.Println("C:", c)
	fmt.Println("C2:", c2)
}

type LogicProvider2 struct {
	Datita string
}

func (lp LogicProvider2) Process(data string) string {
	return data + " *** business logic *** " + lp.Datita
}

/*
*
nothing is declared on LogicProvider to indicate that it meets the interface
*/
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return data + " *** business logic"
}

type Logic interface {
	Process(data string) string
}
type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	data := "get data from somewhere"
	fmt.Println(c.L.Process(data))
}

func embed_ex() {
	fmt.Println("embed_ex()")
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)     // prints 12345
	fmt.Println(m.Desc()) // prints Bob Bobson (12345)
	e := Employee{
		Name: "Pat Pattinson",
		ID:   "54321",
	}
	fmt.Println(e.Desc())
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Desc() string {
	return fmt.Sprintf("%s $(%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) findNewEmpl() []Employee {
	return []Employee{}
}

func iota_ex() {
	/**
	iota-based enumerations make sense only when you care
	about being able to differentiate between a set
	of values and don’t particularly care what the value is behind the scenes.
	If the actual value matters, specify it explicitly.
	*/
	fmt.Println("iota_ex()")
	type MailCategory int
	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Advertisements
	)
	fmt.Println("MailCategories:")
	fmt.Println("Uncategorized:", Uncategorized)
	fmt.Println("Personal:", Personal)
	fmt.Println("Spam:", Spam)
	fmt.Println("Social:", Social)
	fmt.Println("Advertisements:", Advertisements)
}

/**
Functions Versus Methods
Anytime your logic depends on values that are configured at startup or changed while
your program is running, those values should be stored in a struct,
and that logic should be implemented as a method.
If your logic depends only on the input parameters, it should be a function.
*/

func adder_ex() {
	fmt.Println("adder_ex()")
	myAdder := Adder{start: 100}
	fmt.Println(myAdder.AddTo(5)) // 105

	//Create a function from the instance
	f1 := myAdder.AddTo
	fmt.Println(f1(10)) //110

	//Create a func from the Type
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 5000)) //5100
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func intTree_ex() {
	fmt.Println("intTree_ex()")
	var tree *IntTree
	tree = tree.Insert(1)
	tree = tree.Insert(2)
	tree = tree.Insert(3)
	tree = tree.Insert(4)
	tree = tree.Insert(5)
	tree = tree.Insert(5)
	tree = tree.Insert(50)
	tree = tree.Insert(60)
	s := tree.String()
	fmt.Println(s)

	fmt.Println(fmt.Sprintf("Contains %d: %v", 6, tree.Contains(6)))
	fmt.Println(fmt.Sprintf("Contains %d: %v", 5, tree.Contains(5)))
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}

	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}

	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func (it *IntTree) String() string {
	if it == nil {
		return "--"
	}
	return fmt.Sprintf("%d \n L: %v \n R: %v", it.val, it.left.String(), it.right.String())
}

func counter_ex() {
	fmt.Println("counter_ex()")
	var counter Counter
	fmt.Println("Initial Counter: ", counter.String())
	counter.Increment()
	fmt.Println("Counter: ", counter.String())
	counter.Increment()
	fmt.Println("Counter: ", counter.String())
	counter.Increment()
	fmt.Println("Counter: ", counter.String())

	cPointer := &Counter{}
	fmt.Println("cPointer.String(): ", cPointer.String())
	cPointer.Increment()
	fmt.Println("cPointer.String(): ", cPointer.String())

}

func methods_ex() {
	fmt.Println("methods_ex()")
	p1 := Person{
		FirstName: "Bob",
		LastName:  "Bobson",
		Age:       31,
	}

	fmt.Println(p1.String())
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, last Update: %v", c.total, c.lastUpdated)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

/**Metodos:
- Solo se pueden definir en el package block
- Tienen un elemento mas que es el "receiver"
- No se suele usar 'this' o 'self' para ese param

func (r Receiver) nombre(p1 params) (o output){
	Body of Func
}
*/

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

/**
la palabra reservada 'type' le dice al programa que lo siguiente
es un alias para un tipo propio del lenguaje:

type Score int
type Converter func(string)Score
type TeamScores map[string]Score

Converter devuelte un Score, que en este caso,
Score es solo un Int.
*/

/**
Un metodo es una funcion 'atada' a un typo.
*/
