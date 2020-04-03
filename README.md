# go-lang-learning
Learning GoLang

```
The `package main` is used to build `executable` package, anything other means we are creating a `re-usable` code.

> The purpose of `package` in Go is to group together code with similar purpose

> We use `main` in a package to tell Go that ewe want it to be turned into a file that can be executed
```

```
The `import "fmt"` statement means, is used to give out package access to code that is in another package, so giving the `main` package access to the `fmt` library, which is a std-lib in go to print to terminal

go-stlib:
* fmt
* debug
* encoding
* math
* io
* crypto

We can aslo import other peoples code just like python libararies.. `golang.org/pkg` it has all the list of std-libs in Go.

> We use `import` to give our package access to code written in another package. Here we are giving our package access to code in `fmt` package
```

```
The `func main(){ fmt.Println("Hi!")}` is a function in Go eqvuivalent to `def` in python

Definitions:

`func` -> Tell's we are going to define a function.
`main` -> Setting the name of the function.
`()`   -> List of args to pass to the function.
`{}`   -> Function body calling the function runs this space.

> The `main` function will be automatically run when the program is executed.
```

```
How the main.go is organized?
1. Package declaration
2. Import required packages
3. Declare functions, var declarations and core logic of the code
```

## Variables
```
Go is a Static Type language

Dynamic Types are:
JavaScript, Ruby, Python:

e.g:
var number = 123
number = "abcd"
The interpretor won't care if we keep changing the assignments(where we don't care of the assignment of a variable, like one place we put that as string and in other we put it as int)

StaticTypes are:
C++, Go, Java: Doing the above operation will throw a big error, so we have to define the type of the varibale when we are assigning, once we set a var as string, we can't change it in the future
```

```
Basic Go Types:
--------------

Type | Example
bool | true, false
string| "Hello" "How are you"
int | 0, -1000, 99999
float64 | 10.00001, 0.00009, -100.003

there are other types too...
```

```
> You can initialize variables by the following manner:
 * short hand:
   `:=` Go will automatically detect the variable type by its assignment
   e.g: pi := 3.14
 * long hand:
   `=` you have to explicitly tell the var type
   e.g: var pi float = 3.14

* you can initialize or set `variables` outside a `function`, but you can only assign a value to it inside a `function`
```

## functions declaration
```(go)
func newCard() string {

}
```

```
func:    -> key word in Go to define a function
newCard: -> Define a function called 'newCard'
string:  -> When executed, this function will return a value of type 'string'
```