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

### Creating a New Deck

* Remember, in Go, whenever you want a function to return a value, you must declare the *type* of the value you are expecting to be returned
* The plan for creating a new deck:
  1. Create an empty deck
  2. Create a list of all possible *suits*
  3. Create a list of all possible *card values*
  4. Create 2 nested `for` loops for both the list of suits and list of card values
  5. Add a new card for every *value of suit* to the deck
* This will help us avoid having to manually type all 52 cards of the deck
* *Go is vanilla compared to other languages and a lot of concepts and patterns are done in similar ways in Go*
```
	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
```
* **NOTE:**
  * If you save the code above as is, an error will pop up informing that we created the variables `i` and `j` but did not actually use it in our code
    * In Go, whenever we have some variable that we do not actually need to use, we can replace it with an `_`
    * This lets Go know that we understand there should be a variable in that place, but we don't want to use it

### Slice Range Syntax

* For the `deal()` method, each *hand* that is being dealt out will still be of *type* `deck`
  * This means we are going to just be splitting up our `deck` slice into smaller slices
* Breakdown on **slices:**
  * Slices are **zerod-indexed**
  * To access specific elements, the syntax is similar to other languages
    ```
    examlpleSlice[1]
    ```
* For the `deal()` method we are creating, we want to take the existing slice and select a **range** of elements to create a new slice
  * Go has a *helper* for us to do this
  ```
  exampleSlice[startIndexIncluding:upToNotIncluding]
  ```
  * Example:
  ```
  fruits := ["apple", "banana", "orange", "guava"]
  fruits[0:2]
  => ["apple", "banana"]
  ```
  * **OPTIONALLY**
    * You can leave off numbers on either side of the colon to have Go infer that we automatically want to start from the beginning OR go up to the end of the slice
  ```
  fruits := ["apple", "banana", "orange", "guava"]
  fruits[:2]
  => ["apple", "banana"]

  //OR

  fruits[2:]
  => ["orange", "guava"]

  ```

### Multiple Return Values

* When writing our function, `deal()`, we need to write out the *types* for each argument
* This means that our `deal()` function can only take arguments of *type* `deck` and *type* `int`
* `d` and `handSize` are the variable names we use to reference those values within the function
  * Again, it is by convention that we use the abbreviated `d` to reference the *deck object* of *type* `deck`
* Because we know want to return 2 separate slices, we have to annotate both *return types* of both elements are want to return
```
func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}
```
* The next thing to work out is capturing both return values and storing them in their own variables
  * We can do this with some syntax
```
hand, remainingCards := deal(cards, 5)
```
* **NOTE:** 
  * After calling our `deal()` function, it is important to note that the original `cards` variable DOES NOT change or get modified. We created two new references that point at subsections of the 'cards' slice. We never directly modified the slice that 'cards' is pointing at.

### Byte Slices

* `saveToFile()` is a function where we are going to save a copy of the current deck to our local hardDrive
* `newDeckFromFile()` is a function that will load the file saved on the hard drive and create a new deck
* Whenever we want to work with the underlying hardware or operating system that our program is running on, it's best to make use of Go's standard library of [packages]('https://golang.org/pkg/')
* [ioutil]('https://golang.org/pkg/io/ioutil/') is a package that implements common operations for working with the hard drive on our local machine
  * The function we are specifically looking at is `WriteFile()`
  ```
  func WriteFile(filename string, data []byte, perm os.FileMode) error
  ```
  * Looking at the argument list, we can see that we have an argument `data` and it is of *type* `[]byte`
    * When we see the `[]`, we should be thinking of a `slice`
    * This is referred to as a `slice` of *type* `byte`, or a **byte slice**
  * `perm os.FileMode` - `perm` is short for *permissions* and refers to the permissions used to create the file

* **What is a byte slice?**
  ```
  "Hi there!" // string 
  => [72 105 32 116 104 101 114 101 33] // byte slice
  ```
  * Whenever you hear the term *byte slice*, you can automatically think of *string*
  * Every element in the *byte slice* refers to an **ascii character code** 
  * [ASCII]('http://www.asciitable.com/')

### Deck to String

* *Type conversion* is where you take a *type* of value and turn it to another *type*
```
[]byte("Hi there!")
```
* `[]byte` refers to our *desired type*
* `"Hi there!"` is the value we have and whose *type* we want to change to a *byte*
* To create a *byte slice*:
  * We'll take our `deck`, which is really just a `slice` of *type* `string`, combine all the strings into one string, and turn that string into the *byte slice*
    * First, we need to create a helper function, `toString()`, to turn a `deck` into a single `string`

### Joining a Slice of Strings

* To combine our deck into a single string, we can use the [`Join()`]('https://golang.org/pkg/strings/#Join') function built into the standard library
  * To use the `Join()` function, you first need to `import` the strings package
  * **Syntax** for *importing* multiple packages
  ```
    import (
    "fmt"
    "strings"
  )
  ```

### Saving Data to the Hard Drive

```
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```
  * We use `error` as a *return* value type because that is what is going to happen with the `WriteFile()` function 
  * `0666` is a standard *permissions* code

### Reading From the Hard Drive

* For our `newDeckFromFile()` function, we are going to use the [`ReadFile()`]('https://golang.org/pkg/io/ioutil/#example_ReadFile') function from the `ioutil` standard libary package
  * We want to be able to specify the name of the file to open
  * We are going to take the *byte slice* and split the elements from a single string back into a `deck`
* The `error` object is a value of *type* `error`
  * If something went wrong during the function call, the `error` object would be be populated with an actual value 
  * If everything was *successful*, the value of the `error` object would be `nil`
    * In Go, if there is an `error` object being written, usually you would add an `if` statement to check if it has an actual value or if it is `nil`
    * *Error handling* can be a little tricky because it depends on the specific error that occured
      * You can ask yourself this common sense question when trying to figure out how to handle an error:
        * "If something goes wrong here, what do I really want to happen?"
* We can quit our program entirely by using the [os](https://golang.org/pkg/os/) package
  * In the `os` package, we will use the [Exit()](https://golang.org/pkg/os/#Exit) function

### Error Handling

* To convert our *byte slice* back into a `deck`, we can look in the [strings](https://golang.org/pkg/strings/) package for a function named [Split()](https://golang.org/pkg/strings/#Split)
  * Works similarily like `Join()`

### Shuffling a Deck

* To manually *shuffle* through our deck, we are going to **loop through our deck one time**
* At every interval, we will **generate a random number between 0 and `len(cards) - 1`**
* Then, we'll **swap the current card and the card at `cards[randomNumber]`**
* To generate a *random number*, we can look inside the [Math](https://golang.org/pkg/math/)  package, find the [rand](https://golang.org/pkg/math/rand/) sub-package, and use the `Intn()` function
  * We pass the `Intn()` function a number and it will generate a random number between 0 and the number that we passed in

### Random Number Generator

* The way we initially wrote the `newPosition` logic randomizes the shuffle exactly the same every time
* This is because `rand.Intn()` is referred to as a **pseudo-random generator**
  * Go, by default, uses a random number generator that depends on a **seed value**
  * The **seed-value** is the *source of randomness*
  * By utilizing the same seed every time, the **generator** produces the same sequence 
  * To fix this, we need to *generate a new seed value* to truly randomize the sequence

* Looking through the [rand](https://golang.org/pkg/math/rand/#Int) docs, we see documentation regarding the *type* `Rand`
  * The *type* `Rand` is defined as a *source of random numbers*
  * By using creating a value of type `Rand`, we can *specifcy the seed* 
  * To utilize the `New()` function for *type* `Rand`
    * Create a new `int64` > `Source` > `Rand`
      * For the value of type `int64`
        * [Time](https://golang.org/pkg/time/) package > [`Now()`](https://golang.org/pkg/time/#Now) > [type `Time`](https://golang.org/pkg/time/#Time) > `UnixNano()` function
          * This function returns a value of type `int64`
          * The function will take the current time and return the value as a type `int64`

### Testing With Go

* Go testing is **NOT** like RSpec, Mocha, Jasmine, Selenium, etc...
  * Unlike other testing frameworks, Go provides a small interface; only a small set of functions
  * In practice, it's really using regular Go code to test 
* To make a test, creating a new file which ends in `_test.go`
```
deck_test.go
```
* To run all tests in a package:
```
$ go test
```
* NOTE: Even test files require that you specify the package name it belongs to at the top of the file

### Writing Useful Tests

* Deciding What To Test:
  * For `newDeck()`:
    * Verify length of `deck`
    * Verify the value of first card // It should always be the `"Ace of Spades"`
    * Verify the value of last card // It should always be the `"King of Clubs"`
* When creating a test for a function, we create a *new function* within the *test file* with the name **Test** and the function name
  * Example:
  ```
  // in deck.go
  func newDeck() {
    // logic here
  }
  // inside deck_test.go
  func TestNewDeck() {
    // logic to verify deck length
    // logic to verify value of first card
    // logic to verify value of last card
  }
  ```
* **Test Handler** - `t *testing.T`
```
func TestNewDeck(t *testing.T)
```

* **Formatted string**
```
t.Errorf("Expected deck length of 52, but got %v", len(d))
```
* We are *injecting* the value of `len(d)` where the placeholder `%v` is at in the string

### Testing File I/O

* **EDGE CASE**
  * Create a deck > Save to a file > File created! > Attempt to load file > Crash!
* **NOTE**
  * When writing tests with Go, it is important to *manually take care of clean-up*; **remove the temporary testing file that gets created**
* Testing `saveToDeck()` and `newDeckFromFile()`
  1. Delete any files in current working directory with the name "_decktesting"
  2. Create a deck
  3. Save a file; `_decktesting`
  4. Load from file
  5. Assert deck length
  6. Delete any files in current working directory with the name "_decktesting"
* To delete a file, check [os](https://golang.org/pkg/os/) package and [Remove()](https://golang.org/pkg/os/#Remove) function
* When testing, it is important to make tests fail at least 1 time to check for false positives

### Project Review

* For every file created, we list the `package` name that we are working with
* We `import` packages we need to use for additonal functionality
  * In VSCode, packages automatically get added to the import list for us when using their code
* We can create custom *types* 
  * They are essentially an abstraction from the basic data types found in Go - for example, the `deck` type we created is essentially a *slice of type string* - but we were then able to add additional functionality that is specific to our `deck` type
* We create *receiver functions* - or a *function that accepts a receiver* - which essentially allows us to *call that specific function on any value created of the same type*
* **Slices** can thought of or accessed similar to *arrays* used in other languages
  * But there are some additional behaviors particular to *slices*:
    * Able to select a range within a slice simply by using a `:`

## Organizing Data with Structs
### Structs in Go

* In the `deck` project, we used *strings* to represent each card
  * This is problematic if we considered *using the cards in a more functional way*; i.e. playing a game of poker, etc.
    * We would have to use *string manipulation* to figure out the value or suit of each card
* To fix this issue, we could have used a *data structure* called a `Struct`
  * A `Struct` is a data structure that represents a *collection of different properties that are related together or have some type of common purpose*
  * For the card example, we could have created a `Struct` of *type* `card` that has two different *properties* - a `suit` of *type* `string` and a `value` of *type* `string`
    * You can think of a `Struct` as a *plain JavaScript object*
    * You can also think of a `Struct` as a *Ruby hash*

### Defining Structs

* When first creating a `struct`, we need to *define all the properties that `struct` might have
* When defining properties for a `struct`, you DO NOT need to use any semicolons or commas, etc, to separate properties
```
type person struct { // Creates a new type, "person", that is really of type "struct"
	firstName string // First name property for type "person"
	lastName string // Last name property for type "person"
}
```

### Declaring Structs

* When you create a variable in Go and you DO NOT assign it a value, Go *assigns the variable a* **zero value** *to each of the properties of the struct
```
var alex person 
```
* Zero Values in Go:
  * `string` == `""`
  * `int` == `0`
  * `float` == `0`
  * `bool` == `false`
* **NOTE**
  ```
  fmt.Printf("%+v", alex)
  ```
  * `Printf()` is another type of log statement used for *string interpolation*
  * Using `%+v` will log out the property names along with its value
  * SUMMARY:
    * There are 3 different ways for *declaring a struct*
  ```
  alex := person{"Alex", "Anderson"} // Relies on order of definition of fields for assignment // NOT RECOMMENDED WAY

  alex := person{firstName: "Alex", lastName: "Anderson"} // Another way for defining a struct // MORE RELIABLE

  var alex person            // Third way to declare a new struct // Properties will be of "zero value"
  alex.firstName = "Alex"    // Updates firstName property
  alex.lastName = "Anderson" // Updates lastName property
  ```