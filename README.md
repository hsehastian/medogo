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
2. 

