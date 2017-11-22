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
### Embedding Structs

* Go provides the ability to *embed on struct inside of another*
* Example: 
  * A `struct` of *type* `person` can have the following properties
    * `firstName` which is of *type* `string`
    * `lastName` which is of *type* `string`
    * `contact` which is of *type* `contactInfo`
  * And `contact` is actually a `struct` of *type* `contactInfo` with the following properties
    * `email` which is of *type* `string`
    * `zip` which is of *type* `int`
  ```
  	alex := person {
		firstName: "Alex",
		lastName: "Anderson",
		contact: contactInfo{
			email: "aanderson@gmail.com",
			zip: 12345,
		  },
	  }
  ```
  * When using the `{}` syntax for *creating an instance*, we DO need to use commas
  * When declaring a *multi-line structure*, every single line *MUST HAVE a comma* - even the last property declaration

  ### Structs with Receiver Functions

  * We can set up functions that take a receiver with structs
    * This allows any *instances* created from a struct to have similar functionality and behavior
  
  ### Pass By Value

  * How **RAM**, or memory, works on a machine:
    * Memory on a local machine can be thought of little boxes where each box in the computer's memory can store some data and each of those boxes has a discrete address
      * When your program wants to retreive data from the computer's memory, it finds the data according to that address
  * Go is referred to as a **pass by value language**
    * **Pass by value** means that whenever we pass some value into a function, Go will take that value - or *struct* -, copy all of the data in that struct and then place in into a *new box in the computer's memory*
      * So, in our `updateName()` function, we DID NOT update the original struct of `alex`, we ONLY updated the **copy that was made for that particular function call**

### Structs With Pointers

* To fix the issue with have with *pass by value*, we can use **pointers**

### Pointer Operations

* `&variableName` - The `&` is an *operator* that **gets the memory address** of the value the `variableName` is pointing at
  * This is NOT a direct reference to the struct; but a *reference to the memory address of the struct*
* `*pointer` - The `*` is an *operator* that gets the value the memory address, or **pointer**, is referencing
  * So, in the updated `updateName()` function, `pointerToPerson` refers to the *memory address* that `alex` is stored at
* **NOTE:**
  * `*person` - This is a **type description**; it means we're working with a *pointer to a value of type person*
    * When used with a *type*, the `*` should NOT be thought of as an operator but as a *type description*
    * Our `updateName()` function can ONLY be called with a receiver of a *type* of **pointer** to a `person`
  * `*pointerToPerson` - This is an **operator**; it means we want to *manipulate the value the pointer is referencing*
    * We take the memory address that is being referenced and retreives the actual value that is stored at that address
* **SUMMARY**
  * `*variableName` - Turns *address* into a *value*
  * `&variableName` - Turns *value* into an *address*
  * `*typeName` - *Type declaration* saying we are looking for a *pointer of typeName*

### Pointer Shortcut

* Go allows us to take a shortcut when using *pointers* by allowing us to write *receiver functions*: 
  * If we use a variable of a specific type and our receiver takes a **pointer** to a value of that same type, Go will automatically turn that variable into a pointer

### Gotchas With Pointers
```
package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	
	updateSlice(mySlice)
	
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

=> [Bye There How Are You]
```
* When working with **slices**, the `updateSlice()` changed the value of the first element in our slice WITHOUT using a *pointer*

### Reference vs Value Types

* Arrays in Go
  * Primitive Data Structure
  * Can't be resized
  * Rarely used directly
* Slices in Go
  * Can grow and shrink
  * Used 99% of the time for lists of elements
* *Internally*, when we create a slice, Go creates 2 separate data structures, a **slice** and an **array**
  * The **slice** is a data structure that has 3 elements inside it:
    1. A pointer to the **head**
      * Points to the underlying *array* that represents the *actual list of elements*
    2. A **capacity** number
      * How many elements it *can contain*
    3. A **length** number
      * How many elements it *currently contains*
  * In your machine, the *slice* is stored in one address in memory and the *array* is stored in another address in memory
    * NOTE: The `mySlice` variable is *NOT pointing at the actual array that contains the elements*
      * `mySlice` returns the *slice data structure*
* So, when we call the `updateSlice()` function, Go still behaves as a *pass by value* language
  * It makes a copy of the *slice data structure*
    * But the copy of the slice data structure still points to the original array in memory

* Value Types - *Use pointers to change these things in a function*
  * `int`
  * `float`
  * `string`
  * `bool`
  * `struct`
* Reference Types - *DO NOT use pointers to change these things in a function*
  * `slices`
  * `maps`
  * `channels`
  * `pointers`
  * `functions`

## Maps
### What is a Map?

* A **map** is a collection of **key-value** pairs
  * You can think of it similarily like a *hash in Ruby* or an *object* in JavaScript
* Both `keys` and `values` are **statically-typed**
  * This means that all the **keys** must be the **same type**
  * Also, all the **values** must be the **same type**
    * However, the **keys and values can be different types**

### Manipulating Maps

* The **zero-value** of a map is just an *empty map*
* To add values to a map, we can use **square-brackets syntax**
```
mapName[key] = value
```
* Maps DOES NOT allow **dot syntax**
```
mapName.key = value
```
* This is because all the keys are **values of a specific type** NOT variables

### Iterating Over Maps
```
func printMap(c map[string]string) {
	for color, hex := range c {
		
	}
}
```
* `c` - the argument name
* `map[string]string` - *type* of the argument
* `color` - represents the **key** for each iteration
* `hex` - represents the **value** for each iteration

### Differences Between Maps and Structs

* **Map**
  * All keys must be the same type
  * All values must be the same type
  * Keys are indexed - we can iterate over them
  * Used to represent a collection of related properties
  * Don't need to know all the keys at compile time
    * Able to add/remove properties at anytime
  * **Reference Type**

* **Struct**
  * Values can be different types
  * Keys don't support indexing
    * You cannot iterate through all the different properties of a struct
  * You need to know all the different fields at compile time
    * Cannot add/remove properties freely
  * Used to represent a "thing" with a lot of different properties
  * **Value Type**

## Interfaces
### Purpose of Interfaces

* We know that:
  * Every value has a *type*
  * Every function has to specify what type the *type of its arguments and/or its receiver* 
* Does that mean:
  * **Every function we ever write has to be rewritten to accomodate different types even if the logic is identical?**
* This is a small part of the problems that *interfaces* solves for us 
  * They make it easier to reuse code between different parts of our codebase

### Problems Without Interfaces

**LOOK AT CODE**

### Interfaces in Practice
```
type bot interface { // 
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
```
  * Layout for defining an **interface**
  * First, we create a new custom *type* called `bot` that is of *type* `interface`
    * `bot` is used for the example here to layout a more generic version of `englishBot` and `spanishBot`
  * Inside of the definition for the custom *type* `bot`, we define a function called `getGreeting()` that returns a `string`
    * By puttiing this function within the `bot` `interface`, we are saying that ***any *type* defined within the program that also has a `getGreeting()` function AND also returns a value of *type* `string`, they NOW can also be defined as *type* `bot`**
  * Lastly, we define the function `printGreeting()` that takes an argument of *type* `bot`
    * Since the `bot` `interface` allows other *types* defined in the program to ALSO be considered as *type* `bot`, all the other *objects* also have access to the `printGreeting()` function

* **SUMMARY:** we use **interfaces** to *define a method set or function set* 

### Rules of Interfaces
```
type bot interface { // 
	getGreeting(string, int) (string, error)
  getBotVersion() float
  respondToUser(user) string
}
```
  * This is just another example of an `interface` with added *argument type* list and *return type* list
    * Note, we do not need to add variable names for the arguments or returns, only the *types*
  * Additionally, we can add additional functionality 
  * **NOTE:** Whatever we declare within the interface, a *type* must ALSO have ALL the same functionality and behavior to be qualified as the *type* declared in the `interface`

* When creating an application, we can define **two different types**:
  * **Concrete Types**
    * Able to *create a value** directly; we are able to access it, change it, and create new copies of it, etc. 
    * Examples:
      * `map`
      * `struct`
      * `int`
      * `string`
      * `englishBot`
  * **Interface Types**
    * CANNOT create a value directly
    * Example:
      * `bot`

### Extra Interface Notes

* Interfaces are NOT *generic types*
  * Go DOES NOT support *generic types* unlike other languages
* Interfaces are **implicit**
  * We DON'T need to manually say that our custom types satsify an interface
  * Go *implements* our custom types with an interfaces when everything is satisfied within the interface
* Interfaces are a *contract* to help us *manage types*
  * GARBAGE IN -> GARBAGE OUT
    * If a custom type's implementation of a function is broken, then the interface will NOT help us
  * Interfaces should NOT be used as a test for your code validity
* Interfaces are tough. Step #1 is understanding how to read them!
  * Understand how to read interfaces in the standard library. 
  * Writing your own interfaces is tough and requires practice/experience

### The HTTP Package

* To practice more with interfaces, we're going to write a small program that *makes an HTTP request* to `https://google.com` and then we'll *print that response* in the terminal
* [net](https://golang.org/pkg/net/) package - Network Interface package 
  * Includes [HTTP](https://golang.org/pkg/net/http/) package
    * Under *type* [Response](https://golang.org/pkg/net/http/#Response), we can find the function [Get()](https://golang.org/pkg/net/http/#Get)
    ```
    func Get(url string) (resp *Response, err error)
    ```

### Reading The Docs
```
$ go run main.go
&{200 OK 200 HTTP/1.1 1 1 map[X-Xss-Protection:[1; mode=block] X-Frame-Options:[SAMEORIGIN] Date:[Mon, 20 Nov 2017 23:16:28 GMT] Expires:[-1] Content-Type:[text/html; charset=ISO-8859-1] Server:[gws] Set-Cookie:[1P_JAR=2017-11-20-23; expires=Wed, 20-Dec-2017 23:16:28 GMT; path=/; domain=.google.com NID=117=vvz_hb_BAlYaMJQvdKrv345Wb_BAQ9L4-zrEeZkEmeXc3prQBNssPqwfVkxYoWM4fEaa8MJhQjiE44bcW5n93EPAv5N_yZkYMkNUHxsWPXIO3fNmuLru0eKzltHz_yyi; expires=Tue, 22-May-2018 23:16:28 GMT; path=/; domain=.google.com; HttpOnly] Cache-Control:[private, max-age=0] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."]] 0xc4200e26c0 -1 [] false true map[] 0xc4200fa100 <nil>}
```

  * Looking at this output for our program the way it is, we don't really see the **body** of the **response** that we were expecting
    * The **body** is the **HTML** that renders the Google homepage

* Looking at the docs we can click through what exactly the `*Response` is pointing to
  ```
  func Get(url string) (resp *Response, err error)
  ```
  * `*Response` is a **pointer** to a *response object*
  * In the [docs](https://golang.org/pkg/net/http/#Response), we can see that the `*Response` pointer is pointing at a `Response` `struct`
  ```
  type Response struct {
  ```
  * Looking closer at the docs, we find that `Body` is of *type* `ReadCloser`
  ```
  // The Body is automatically dechunked if the server replied
  // with a "chunked" Transfer-Encoding.
  Body io.ReadCloser
  ```
  * When we click on [ReadCloser](https://golang.org/pkg/io/#ReadCloser), we find that it is an `interface`
  ```
  type ReadCloser interface {
        Reader
        Closer
  }
  ```
  * NOTE: This `ReadCloser` interface looks very different than the format that we first saw when dealing with *interfaces*
  
  * By clicking on [Reader](https://golang.org/pkg/io/#Reader) and [Closer](https://golang.org/pkg/io/#Closer), we find out that they are also *interfaces* with syntax that looks familiar to what we've seen already
```
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
```
type Closer interface {
        Close() error
}
```
### More Interface Syntax
```
type Response struct {
  Status string
  StatusCode int
  Body io.ReadCloser
}
```
  * If we specify an `interface` as a *value* inside of a `struct`, we're saying that the  `Body` field can have ANY value assigned to it AS LONG as it fulfills the `interface` `ReadCloser`
    * By clicking through, we then find out that to fulfill the `interface`, we just need to define a `struct` that has the functions `Read()` and `Close()` that takes the same *types* of arguments and returns the same *types* of return-values

```
type ReadCloser interface {
      Reader
      Closer
}
```
  * In Go, we can take **multiple interfaces** and assemble them together to create another `interface`
    * Both `Reader` and `Closer` are *interfaces* 

### Interface Review

* In general, we use *interfaces* because in some conditions, they allow us to reuse some code 

### The Reader Interface

* The `Reader` `interface` is used to solve the issue of handling different *types* of data input and their different return types - i.e. `[]string`, `int`, `float64`, `jpeg`, etc. - that all have similar functionality
* So, no matter what the *source of input* is, the `Reader` interface allows us to handle that input without having to write custom functions to handle the different *types* and return the output as a *byte slice*

### More on the Reader Interface
```
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
* In the background of the `Reader` interface, the `Read()` function will be given a *byte slice* as an argument and then the `Read()` will inject some data into the *byte slice*
  * This works because as we learned earlier, with *pointers*, if we pass a `slice` into a *function*, that function can freely modify the slice and the original slice gets updated
* As for the two *return values* from the `Read()` function, the `int` value which gets returned is the **number of *bytes* that was read, or injected, into the slice** 
  * The second return value is just another `error`

### Working with the Read Function
```
bs := make([]byte, 99999)

resp.Body.Read(bs)
fmt.Println(string(bs))
```
* `make()`
  * Built-in function in the Go standard library
  * Takes an argument of a *type*
  * If we are creating a `slice`, we can add a second argument of *type* `int`, to be the *number of elements, or empty spaces, we want in our `slice`
  * The reason why initialize the *byte slice* with so many empty spaces is because the `Read` function is NOT set up to automatically resize the *byte slice* if it gets full; it will simply fill the *byte slice* until it reaches capacity

* Next, we'll take the the `resp` struct, access the `Body` property - this is the value that *implements* the `Reader` `interface` - and access the `Read()` function
* Then, we'll pass in our *byte slice* into the `Read()` function and the `Read()` should fill the `bs` with the HTML data
* Remember, when we deal with *byte slices*, we can essentially think of them as a string, so we can use *type conversion* when printing out the output of `bs`
```
$ go run main.go
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en"><head><meta content="Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for." name="description"><meta content="noodp" name="robots"><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color..........
```
* **NOTE:** In reality, we don't need to really create the *byte slice* in this way, Go does have some helper functions to implement the `Read()` function a little bit easier

### The Writer Interface
**INSTEAD OF WRITING OUR CODE LIKE THIS**
```
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```
**THIS ONE-LINER DOES THE SAME THING**
```
io.Copy(os.Stdout, resp.Body)
```
* The `Writer` `interface` is used to take some information from within our program and send that data out to some form/channel/method of output 
* For example:
  * We take a *byte slice* 
  * Pass it to a *value that implements the* `Writer` *interface*
  * The `Writer` interface is essentially describing something that can take info and send it outside of our program 
  * Sources of Outputs:
    * Outgoing HTTP request
    * Text file on a hard drive
    * Image file on a hard drive
    * Terminal 
* In order to use the `Writer` interface, we need to find something in the standard library that *implements* the `Writer` interface, and use that to log out all the data that we're receiving fromt the `Reader`

### The io.Copy Function

* In the [io](https://golang.org/pkg/io/) package, we can find the docs on the *type* [Writer](https://golang.org/pkg/io/#Writer)
```
type Writer interface {
        Write(p []byte) (n int, err error)
}
```
* Although similar to the `Reader` interface, the *byte slice* that is passed here to the `Write()` function is truly being used as a *source of input*
  * The `Write()` receives the *byte slice*, takes the data within the *byte slice*, and sends all that information to some output channel

* Looking more closely at the [docs](https://golang.org/pkg/io/#Copy) for the `io.Copy()`
```
(func Copy(dst Writer, src Reader) (written int64, err error)) // dst == destination
```
```
io.Copy(os.Stdout, resp.Body)
```
* The `Copy()` expects the first argument to be some value that implements the `Writer` interface
  * Remember, the `Writer` interface can be thought of something that is taking data and sending it outside of our application
* The `Copy()` also expects the second argument to be some value that implements the `Reader` interface
* Essentially, the `Copy()` as being used to take some information from outside our application and write/copy it to some outside channel
  * The `os.Stdnout` is simply a value that implements the `Writer` interface
    * `os.Stdout` is a *value* that is exported by the `os` package and it is of *type* `*File`
    * `File` has a function called `Write()`
      * Therefore it impements the `Writer` interface

### The Implemenatation of io.Copy

* In VSCode specifically, we can click on functions using **CMD + click** and get to the source code for that function

### A Custom Writer

```
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
```
## Channels and Go Routines
### Website Status Checker

* **Channels** and **Go Routines** are both structures inside of Go that are used to handle *concurrent programming*
* We're going to write a small program that will take a list of websites and make an HTTP `GET` request to each of them
  * The program will make sure the websites are up and responding to HTTP traffic

### Printing Site Status
### Serial Link Checking

* Right now, our program is working in a sequential/serial flow; we take a link, make a request, and wait for the response before moving onto the next link in the slice
* Instead of this, we can try to create a more *parallel* solution
  * We make a request to each link within our slice and print out the responses as soon as we receive them

### Go Routines 

* When we execute a Go program, we automatically create 1 *Go Routine*
* A **Go Routine** can be thought of as something that exists inside of a running program, or **process**
  * The Go Routine takes every line of code inside of the program and executes them one-by-one 
* Inside of our program the `http.Get(link)` is referred to as a **blocking call**
  * This is because this piece of code does take some time to execute
  * While it is being executed, the *main Go Routine* cannot do anything else
```
go checkLink(link)
```
  * Here, in the updated version of the code, we added a new keyword `go`
  * `go` - This keyword means to *run the function inside of a new **Go Routine**
    * The *Go runtime* will automatically create a new Go Routine to run the `checkLink()`

### Theory of Go Routines

* A **Go Routine** can be thought of as something that executes lines of code, line-by-line, inside of a *function call*
* Behind the scenes, in your local machine/operating system, there is something called the **Go Scheduler**
  * The **Go Scheduler** works with *one CPU* on your local machine, even if you're running a dual-core machine
  * So, even though we are *launching multiple Go Routines*, only one is being executed or running at any given time
  * The purpose of the **Go Scheduler** is to *monitor the code that is running inside each of the Go Routines
    * When the Scheduler detects that all the code inside of a Go Routine has been executed or it has reached a *blocking call*, it will *pause* the current Go Routine and move on to the next one in the chain
      * The Scheduler runs one routine until it finishes or makes a blocking call
  * By default, Go will attempt to use only one *CPU core* 
    * If we override this setting - by letting Go use multiple CPU cores - each CPU core can *run one Go Routine at a time*
    * So, the Go Scheduler will assign each CPU core a Go Routine and allow each CPU to run one Go routine each
      * The Go Scheduler runs one thread on each "logical" core

* **Concurrency is NOT parallelism** - a phrase that is often discussed in the Go community
  * A program IS *concurrent* if it has the ability to *load up* multiple Go Routines at a time
    * **Concurrency** - can have multiple threads executing code. If one thread blocks, another thread is picked up and worked on
    * This refers to a Go program only working with one CPU core
  * **Parallelism** can only be achieved by working with *multiple CPU cores* on a machine
    * Multiple threads *executed* at the exact same time. Requires multiple CPUs

* When running a program:
  * The **main Go Routine** is created when the *program is launched*
  * **Child Go Routines** are created when we use the `go` keyword
    * These child Go Routines behave slightly different that the main Go Routine

### Channels

* NOTE: The `go` keyword is ONLY used in front of function calls
* If we run our program the way it is right now, nothing will be logged out to the terminal
* The **main Go Routine** is the single routine inside the program the controls when the program should exit or quit
  * This means that the main Go Routine *exits the program* before the **child Go Routines** have a chance to complete
* We can fix this issues using **channels**
  * **Channels** are used to *communicate between different, running Go Routines**
  * We'll use one channel to make sure the main Go Routine is aware of when each of the child Go Routines have been completed
  * Channels are the ONLY way to communicate between different Go Routines
  * We treat channels just like any other values in Go
    * We create them essentially the same way we create a `struct`, a `slice`, an `int`, a `string`, etc. 
    * Channels are also **typed**, just like any other variable
      * This is referring to the *data* that is passed into the channel must all be the same *type*

### Channel Implementation
```
	c := make(chan string)
```
  * Here, we making a new channel, `chan`, that will take data of type `string`
  * Even though we are using the channel for *communication*, we still treat it as a *value*
  * This means we need to update our code for us to pass in `c` as an argument to our `checkLink()`
```
  func checkLink(link string, c chan string)
```
  * When we defined `checkLink()` with the added argument:
    * `c` represents the argument variable name
    * `chan` says that the argument is a of type channel
    * `string` refers to the type of data the channel will process

* Sending data with Channels:
```
  channel <- 5
```
  * Send the value `5` into this channel

```
  myNumber <- channel
```
  * Wait for a value to be sent into the channel. When we get one, assign it to the variable `myNumber`

```
  fmt.Printlin(<- channel)
```
  * Wait for a value to be sent into the channel. When we get one, log it out immediately

* NOTE: If the output we get in the terminal still runs only one `checkLink()` Go Routine before the program exited

### Blocking Channels

* When we executed our program, we only got one log statement before the program exited
  * This is because *whenever we wait for a message to come through the channel, it is a blocking call*
  * So, receiving messages from a channel is a **blocking line of code**

### Receiving Messages
```
for i := 0; i < len(links); i++ { 
		fmt.Println(<-c)
	}
```
  * We use this `for` loop to listen for the messages from our channel up to the length of our `slice`

### Repeating Routines

* What is we want to conitually ping each link for than once until we get an `error`?


