# medogo
List of notes and code base on what I learn from udemy "Go: The Complete Developer's Guide (Golang)"

Since I come from PHP world so some notes here will be using analogy from PHP ◔_◔
## Hello World
1. install Golang: https://golang.org/doc/install
2. setup local: https://golang.org/doc/install#install
3. install VSCode: Need to install Go extension then to trigger install go analysis tool just create new file and set language to "Go" it will trigger VSCode to install the analysis tool
4. create simple hello world
5. `go run main.go` it will compile and run the code
6. `go build main.go` it will just compile the code, it will generate executeable file from main.go
7. `package main` and `func main()` should have both in same file e.g. `main.go`
8. there is executeable package e.g. `package main` and re-useable package e.g. package that created by another dev
9. use `import fmt` is standard Go lib for "format", usually use for print text
10. go to https://pkg.go.dev/std for more info about Go standard packages

## Cards
1. to declare variable we can use `var <variable name> <data type>` e.g. `var card string = "Five of Diamonds"`
2. use `=` to assign value to variable
3. use `:=` to declare and assign value to variable e.g `card := "Ace of Spades"`, Go can identify the variable data type base on assigned value
4. It will throw error if use `:=` to the same variable that we already declare before
5. we can initialize variable outside function but we cannot assign value to it e.g
```go
// this is valid
package main

import "fmt"

var number int

func main() {
    number = 5
    fmt.Println(number)
}

---

// this is invalid
package main

import "fmt"

var number int = 5

func main() {
    fmt.Println(number)
}

```
6. function in Go can return data type e.g.
```go
func newCard() string {
    return "Five of Diamonds"
}
```
7. we can call a function in other go file as long as it using same package name e.g.
```go
// main.go
package main

import "fmt"

func main() {
    number = 5
    fmt.Println(printState())
}

---

// state.go
// in this file we don't need to declare `func main()` anymore since it already exist in main.go
package main

import "fmt"

func printState() {
    number = 5
    fmt.Println(printState())
}
```
8. there is 2 type of array in Go:
    - Array: array with fixed length of list
    - Slice: array that can dynamically added or removed at will, notes: slice must contain element with same data type
9. `for i, card := range cards` the reason why we use `:=`(declare and assign) it because `range cards` will always return 2 value. the first value is the index and the second value is copy of the element at that index
10. we can define a "user-defined" data type by declare `type <type name> <data type>` e.g.
```go
// main.go
package main

import "fmt"

func main() {
    // we declare and assign "cards" variable with type "deck" from deck.go
    cards := deck{"Six of Spades", "Five of Diamonds"}
    for i, card := range cards {
        fmt.Println(i, card)
    }
}

---

// deck.go
package main

// we can assume `type` is like `class` in PHP (kinda ¯\_(ツ)_/¯)
type deck []string
```
11. to run multiple go file, we can execute `go run {file1} {file2}` e.g. `go run main.go deck.go`
12. we can create a "receiver" function when using`type` e.g.
```go
// deck.go
package main

import "fmt"

type deck []string

// `d` here can be assumed like `$this` in PHP, so this function here can be assumed simillar to instance method e.g. `$this->print()`
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```
13. when a variable assign with the data `type` that variable can access the receiver method from that types e.g.
```go
// main.go
package main

func main() {
    cards := deck{"Ace of Diamonds", newCard()}
    // add new element to array
    cards = append(cards, "Six of Spades")

    // call receiver from type "deck"
    cards.print()
}
```
14. case like `for i, card := range cards` the `range cards` always return 2 values so we need to assign it to variables, but since in our code we're not using `i` Go will complain this un-used variable. To solve this issue we can replace `i` with `_` (underscore) this will tell Go to ignore the `i` e.g. `for _, card := range cards`
15. to use array slice in Go we can do `cards[<position of index> : <range of array>]` in PHP this like `array_slice`
```go
//                    0        1         2         3
fruits := []string{"apple", "banana", "grape", "orange"}

fruits[0:2]     // output: "apple", "banana"
fruits[:2]      // output: "apple", "banana" this equal to above syntax, this mean slice array from begining of array to range 2
fruits[2:]      // output: "grape", "orange" this mean slice array from index 2 to end of array

```
16. a function in Go can return multiple value example like below
```go
// main.go
package main

import "fmt"

func main() {
    fruits := []string{"apple", "banana", "grape", "orange"}

    sliceOne, sliceTwo := sliceFruit(fruits, 2)
    fmt.Println(sliceOne)
    fmt.Println(sliceTwo)
}

func sliceFruit(f []string, offset int) ([]string, []string){
    return f[:2], f[2:]
}
```
17. `[]byte` turn value into slice of byte e.g.
```go
// main.go
package main

import "fmt"

func main() {
    text := "Hello World"
    fmt.Println([]byte(text)) // output: [72 101 108 108 111 32 87 111 114 108 100]
}
```
18. Go have `strings` package that can convert slice of string to string, in PHP it same like `explode` function
19. to use multiple packages in Go we just need to add parentheses after import e.g.
```go
// main.go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fruits := []string{"apple", "banana", "grape", "orange"}
    text := strings.Join(fruits, ",")
    fmt.Println(text)   // output: "apple,banana,grape,orange"
}
```
20. to make a test in Go, just create a file with format `<filename>_test.go` e.g. `deck_test.go`
21. to run the test just execute `go test`
22. when running `go test` then get error `go: cannot find main module, but found` there is 3 solution here
- execute `go env -w GO111MODULE=auto`
- execute `cd <path to project> && go mod init <project name>`
- execute `go test main.go <file for test>.go <file for test>_test.go` e.g. `go test main.go deck.go deck_test.go`
## Struct
1. in Go we can create a custom data structure using `struct` e.g.
```go
type person struct {
    firstName string
    lastName string
}
```
2. there is 3 ways to declare struct in Go
```go
// main.go
package main

import "fmt"

type person struct {
    firstName   string
    lastName    string
}

func main() {
    // option #1
    p := person{"John", "Doe"}

    // option #2
    pp := person{firstName: "John", lastName: "Doe"}

    // option #3
    var ppp person
    ppp.firstName: "John"
    ppp.lastName: "Doe"
}

```
3. we can call link struct inside another struct(nested) e.g.
```go
// main.go
package main

import "fmt"

type contactInfo struct {
    email   string
}

type person struct {
    firstName   string
    lastName    string
    contactInfo
}

func main() {
    p := person{
        firstName: "John",
        lastName: "Doe",
        contactInfo: contactInfo{
            email: "john@mail.com",
        },
    }
}
```
4. operator `&` and `*` is use in Go pointer
5. operator `&` is used tell Go that we want the variable memory address e.g.
```go
// main.go
package main

import "fmt"

func main() {
    x := 10
    y := &x

    y = 0

    fmt.Println(x) //   output: 10
    fmt.Println(y) //   output:0xc000a410
}
```
6. operator `*` is used tell Go that we want the value from the pointer e.g.
```go
// main.go
package main

import "fmt"

func main() {
    x := 10
    y := &x

    fmt.Println(x) //   output: 10

    // by doing this it means that we want pointer y which is poiting to variable x change the value to 0
    *y = 0

    fmt.Println(x) //   output: 0
}
```
7. but if we see a receiver function like this `func (p *person) update()` an operator `*` in front of data type (we call it type description) this mean that this receiver function can be called by variable that assigned with pointer for data type person or the variable it self
```go
// main.go
package main

import "fmt"

type person struct {
    firstName   string
    lastName    string
}

func main() {
    p := person{
        firstName: "John",
        lastName: "Doe",
    }

    // option #1 : still valid
    pPointer := &p
    pPointer.update("Mark")
    p.print()

    // option #2 : recommended
    p.update("Mark")
    p.print()

}

// as long as the receiver function using pointer person (*person),
// Go will know that in this receiver function we want to access the pointer for data type person
func (p *person) update(n string) {
    (*p).firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```
8. gotchas with pointer, see example below
```go
// main.go
package main

import "fmt"

func main() {
    s := []string{"Hi", "There", "How", "Are", "You"}

    updateSlice(s)

    fmt.Println(s)  // output: {"Holla", "There", "How", "Are", "You"}

    // as you can see above the value of the slice changed, but as we learn from notes #7
    // to update the value we need to passing the pointer
    // the reason why this work for the slice it because in Go create Array and Slice data in the memory
    // #1: it store the slice data e.g. pointer to the head, capacity of the current slice, length of the slice
    // #2: it store the Array value
    // when we pass the slice to the `updateSlice` function Go will copy slice data only (#1) but not the Array value,
    // the copy of the slice data still pointing to the Array value thats why `updateSlice` can update the element value (pass by reference)
}

func updateSlice(s []string) {
    s[0] = "Holla"
}
```
9. to differentiate which data type that we need to handle with pointer and which one is not, check the list below.
- Value Type: a data type when you need to update the value in the function you need to pass it as a pointer, it consists of: int, float, string, bool, structs
- Reference Type: a data type when you need to update the value in the function you just update it directly without need the pointer, it consists of: slices, maps, channels, pointers, functions
## Map
1. to do `map` in Go we can do `map[<data type>]<dada type>` e.g. `map[string]string`
2. there is 3 way to declare map in Go e.g.
```go
// main.go
package main

import "fmt"

func main() {
    // option #1 use var
    var colors map[string]string
    colors["red"] = "#ff0000"

    // option #2 use make
    colors := make(map[string]string)
    colors["red"] = "#ff0000"

    // option #3 declare and assign
    colors := map[string]string{
        "red": "#ff0000",
    }

    fmt.Println(colors)
}
```
3. to delete key in map we can do `delete(<map>, <key>)` e.g. `delete(colors, "red")`
## Interface
1. in Go we can create interface data type using `interface` e.g.
```go
type bot interface {
    getGreeting() string
}
```
2. `interface` in Go have different behaviour compared to PHP or any other oop language, usually in oop the class need to `implements` the interface so the class can have a same function that defined by interface but interface in Go is more like we add the `type` as a member of `interface` and as long as the `type` have receiver function use same function name that defined in `interface` also as long as it in the same package it considered already implement the `interface` (°ロ°)!
```go
// main.go
package main

import "fmt"

type car interface {
    hasGas() bool
}

type honda struct{}
type toyota struct{}
type ducati struct{}

func main() {

    h := honda{}
    t := toyota{}
    d := ducati{}

    startEngine(h)  // this is valid because type `honda` have receiver function that implement interface `car`
    startEngine(t)  // this is valid because type `toyota` have receiver function that implement interface `car`
    startEngine(d)  // this is error because type `ducati` dont have receiver function that implement interface `car`
}

func startEngine(c car) string {
    if c.hasGas() {
        fmt.Println("Engine started!")
    }else {
        fmt.Println("Engine die!")
    }
}
// honda receiver function
func (honda) hasGas() bool {
    return true
}
// toyota receiver function
func (toyota) hasGas() bool {
    return false
}
```
### Reader Interface
1. I dont know what the heck is this why it so complicated (╯°□°)╯︵ ┻━┻ (after I understand more I will add more notes here)
