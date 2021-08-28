# medogo
List of notes and code base on what I learn from udemy "Go: The Complete Developer's Guide (Golang)"
## Hello World
1. install Golang: https://golang.org/doc/install
2. setup local: https://golang.org/doc/install#install
3. install VSCode: Need to install Go extension then to trigger install go analysis tool just create new file and set language to "Go" it will trigger VSCode to install the analysis tool
4. create simple hello world
Notes:
- `go run main.go` it will compile and run the code
- `go build main.go` it will just compile the code, it will generate executeable file from main.go
- `package main` and `func main()` should have both in same file e.g. `main.go`
- there is executeable package e.g. `package main` and re-useable package e.g. package that created by another dev
- use `import fmt` is standard Go lib for "format", usually use for print text
- go to https://pkg.go.dev/std for more info about Go standard packages

## Cards
1. to declare variable we can use `var {variable name} {data type}`
2. use `=` to assign value to variable
3. use `:=` to declare and assign value to variable e.g `card := "Ace of Spades"`, Go can identify the variable data type base on assigned value
4. It will throw error if use `:=` to the same variable that we already declare before
5. we can initialize variable outside function but we cannot assign value to it e.g
```
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
6. other then `func main()` any other new function should return data type e.g.
```
func newCard() string {
    return "Five of Diamonds"
}
```
7. we can call all function in other go file as long as it using same package name e.g.
```
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