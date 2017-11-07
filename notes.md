# Complete Dev Guide to Go (Golang) - Udemy

## Environment Setup

1. Install [Go]('https://goland.org/dl/') - *Runtime to build and execute Go programs*
2. Install [VSCode]('https://code.visualstudio.com/') - *Code editor with one of the best Go integrations*
3. Confirgure VSCode - *Install Golang extension*
4. Write Code!

## A Simple Start
### Five Important Questions

* These are 5 basic questions we can ask to get a better sense of what the `hello, world` program is doing
1. How do we run the code in our project?
  * `$ go run *file_name*`

    | Go CLI |
    | ------ |
    | `$ go build` - Compiles a bunch of go source code files (This command creates an *executable*  |
    | `$ go run` - Compiles and executes go files |
    | `$ go fmt` - Formats all the code in each file in the current directory   |
    | `$ go install` - Compiles and "installs" a package (Used to handle *dependecies) |
    | `$ go get` - Downloads the raw source code of someone else's package (Used to handle *dependecies)   |          
    | `$ go test` - Runs any tests associated with the current project    |  
            
2. What does `package main` mean?
  * In Go, package == project == workspace
  * A `package` is a collection of common source code files
    * A `package` can have many files contained in it, each of them ending with the extension `.go`
    * The only requirement of every file in a package is the first line of each file must declare which package it belongs to
    * Two different types of packages
        * **Executable**
          * Generates a file that we can run
          * Used primarily for doing something 
        * **Reusable**
          * Code used as "helpers"
          * Good place to store reusable logic
          * Similar to code dependencies or libraries
          * How to know if you need to make an *executable* or *reusable* file?
          * The name of the package you use determines whether or not your file is *executable* or *reusable* 
      * `main` is specifically used to create an *executable* type package
        * Anything other name used would create a *reusable* type package
        * Anytime you create an *executable* type package, it must always contain a `func main()` inside of it
3. What does `import "fmt"` mean?
  * `import` keyword is used to access code from another package in the current package
  * `fmt` is a standard library package that is included in Go by default
    * *fmt* is short for **format**
    * This library is used to print all sorts of information to the terminal
  * By default, the `main` package the we create has NO access to the standard, default packages included with Go
    * To gain access to the other packages from within our package, we use the `import` statement to form a link between the packages
    * We can also `import` packages that are not a part of the default, standard library; packages authored by a 3rd party
      * Link to documentation regarding Go's standard library [packages]('https://golang.org/pkg')
4. What's that `func` thing?
  * `func` is short for function
    * They work very similarly to functions in other programming languages
5. How is the `main.go` file organized?
  * We are going to always use the exact same pattern when creating a file
    * `package main` - Package declaration
    * `import "fmt"` - Import other packages that we need
    * `func main() {}` - Declare functions, tell Go to do things

**NOTE:**
  * You can access course diagrams [here]('https://github.com/StephenGrider/GoCasts/tree/master/diagrams')
    * Open the folder containing the set of diagrams you want to edit
    * Click on the ‘.xml’ file
    * Click the ‘raw’ button
    * Copy the URL
    * Go to https://www.draw.io/
    * On the ‘Save Diagrams To…’ window click ‘Decide later’ at the bottom
    * Click ‘File’ -> ‘Import From’ -> ‘URL’
    * Paste the link to the XML file
    * Tada!

## Deeper Into Go

* We are going to build a functional, deck of playing cards
* The functions our package will include are:
  * newDeck()
    * Creates a list of playing cards. Essentially an array of strings
  * print()
    * Logs out the contents of a deck of playing cards
  * shuffle()
    * Shuffles all the cards in a deck
  * deal()
    * Creates a hand of cards
  * saveToFile()
    * Save a list of cards to a file on the local machine
  * newDeckFromFile()
    * Load a list of cards from the local machine

### Variable Declarations

```
  var card string = "Ace of Spades"
```
  * `var` - keyword used to create a new **variable**
  * `card` - name of the variable 
  * `string` - only a **string** value will be assigned to this variable
    * This is called **type declaration**
  * `"Ace of Spades"` - value of the variable

* **Go** is referred to as a *statically-typed language*
  * Examples of **static** languages:
    * C++
    * Java
    * Go
  * Examples of **dynamic** languages:
    * JavaScript
    * Ruby
    * Python

* **Dynamically-typed languages** are languages where developers essentially do NOT care about what values are assigned to which variable
  * You can redfine a specific variable different *types* of values 

* *Note:* In VSCode, if you try to use `fmt` functionality without first declaring `import fmt`, VSCode will automatically input the import statement for you

* Basic data types in Go:
  * bool
  * string
  * int
  * float64 

* **NOTE:**
  * You can define variables another way
`	card := "Ace of Spades"`
  * This is because the Go compiler can *infer* what we are trying to do with the variable `card` 
  * ONLY use the `:=` syntax when **defining a new variable**
    * To **reassign** a variable's value, you simply use the `=` operator
  * ALSO a variable that is initialized AND assigned a value OUTSIDE of a function is not scoped to be accessed from INSIDE the function
    * HOWEVER a variable that is initialized OUTSIDE of a function and later ASSIGNED a value from INSIDE the function is accessible from with the function
  
### Functions and Return Types

```
func newCard() string {
	return "Five of Diamonds"
}
```
* It is necesarry to let the Go compiler know that the function `newCard()` will have a return value of type string
  * If we remove the `string` keyword, we get an error:
  ```
  too many arguments to return
    have (string)
    want ()
  ```
  * **NOTE:** 
    * Every function that returns a value must indicate what *type of value* it is returning
    * Files in the SAME package can freely call functions defined in other files.
  
### Slices and For Loops

* Go has 2 basic data structures for handling *lists*
  * **Array**
    * Fixed length list of things
  * **Slice** 
    * An array that can grow or shrink
  * Both *arrays* and *slices* must be **defined with a data tyoe**
    * Every element within the list has to be of the same data type

* Syntax for creating a new slice:
```
  cards := []string{}
```
  * `cards` is the variable that we will assign our slice to as the its value
  * `[]string{}` declares that we are creating a slice, with type `string`
    * All elements that we are inserting go within the `{}`
  * To add elements into our slice
  ```
  exampleSlice = append(exampleSlice, desiredElement)
  ```
  * The `append()` does NOT modify the slice, it returns a new slice with the updated elements and reassigns the variable `cards` to the value of the new slice

* How do we iterate over a slice?
```
for i, card := range cards {
		fmt.Println(i, card)
	}
```
  * `for` is the keyword to define that we are entering a *loop*
  * `i` represents the index; you can name this anything you want, it does not need to be "i"
  * `card` represents the current element in the iteration; again, this can be named anything you want
  * `range cards` represents the slice and loops over each element
  * `fmt.Println(i, card)` is run one time for each element in the slice
  * The reason why we use the `:=` to instantiate a new card for every iteration of the `for` loop is because through every iteration, we are "throwing away" the previous index and instance that was declared

* **NOTE:**
  * If using the `for` loop in the way we just learned, you need to also print the `index` along with the actual element 
    * If you try to just print the element, you'll throw an error

### OO Approach vs Go Approach

* Go is NOT an object-oriented language; so there is no idea of *classes*
* Base Data Types in Go
  * string
  * integer
  * float
  * array
  * map
* We can "extend" a **base type** and add some extra functionality to it
  * `type deck []string`
    * Here, we are going to create a new *type*, called `deck`, that will essentially be a `slice` of type `string`
    * This tells Go we want to create a `slice` of type `string` and add a bunch of functions specifically made to work with it
* Functions with `deck` as a *receiver*
  * A function with a receiver is like a *method* - a function that belongs to an *instance*
* This is a common pattern that is used in a lot of Go programs

* Notes on File Structure:
  * `main.go`
    * Code to create and manipulate a deck
  * `deck.go`
    * Code that describes what a deck is and how it works
  * `deck_test.go`
    * Code to automatically test a deck

* **NOTE:** Every single file inside of a package must declare the package name on the first line at the top

### Custom Type Declarations

* To reiterate, we are going to create a new *type*, a `deck` type, which will be a `slice` of type `string`, so we can add functionality specific to the `deck` type
  * This is called **type declaration**
* By default, our `deck` type will have all the functionality of a `slice` of type `string`; similar to inheritance from Object-Oriented programmming
* Because of this, we can replace `[]string` in our code to simply say `deck` and it should function normally
* Because we are using code written in one file, `deck.go`, and running it another file, `main.go`, when we are executing the program, we need to run both files 
  ```
  $ go run main.go deck.go
  ```
* This *compiles* both files and executes the code in both of them

### Receiver Functions
```
  func (d deck) print() {
    for i, card := range d {
      fmt.Println(i, card)
    }
  }
```
  * `(d deck)` - This syntax is called a **receiver**
  * The entire thing is referred to as a *receiver on a function*
  * For this example, the *receiver* is of *type* `deck` and the function's name is `print()`
  * This makes it so ANY variable of *type* `deck` now gets access to the `print()` method
  * *The receiver sets up methods on variables that are created*
```
(d deck)
```
  * `d` refers to the actual copy, or **instance**, of the *deck* we're working on
  * `deck` allows ANY variable of *type* `deck`to have access to the `print()` method
  * Compared to an object-oriented language, `d` is similar to keywords such as `this` or `self`
    * In Go, we don't use keywords such as `this` or `self`; you always refer to the receiver as the actual thing that it is
  * By convention, the **receiver** is always referred to with a *1 or 2 letter abbreviation of that matches the type of the receiver*
  * This is a common pattern that is widely used in the Go language