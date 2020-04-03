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