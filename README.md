# Go Lang

## Variable Declaration

````go
varName := value

var varName2 string

// constant variable declaration
const var varName3 = value

// or

const var varName4 float32 = 34.56
````

## Common Data types
````go
int int8 int16 int32 int64
bool string complex64 complex128
uint uint8 uint16 uint32 uint64
float32 float64
rune - for unicode characters
````


## Taking User Input

```go
var age int
fmt.Printf("Enter your age: ")
fmt.Scan(&age)
```
& is used to pass the reference or pointer address of a variable



## Strigns

```go
var userName string = "John Doe"

// or

userName := "John Doe"
```

###### String Operation functions

> len(userName): Gives the length of the string

> userName[0]: Gives the character at index 0

> userName[0:4]: Gives the characters from index 0 to 4

> userName[0:]: Gives the characters from index 0 to end

> userName[:4]: Gives the characters from start to index 4

> userName[:]: Gives the characters from start to end

> userName + "Hello": Concatenates the string

> fmt.Sprintf("Hello %s", userName): Concatenates the string

> strings.ToUpper(userName): Converts the string to uppercase

> strings.ToLower(userName): Converts the string to lowercase

> strings.Contains(userName, "John"): Checks if the string contains the substring

> strings.Replace(userName, "John", "Jane", 1): Replaces the substring with another substring, where 1 is the number of occurences to be replaced

> strings.Split(userName, " "): Splits the string into array of strings based on the delimiter

> strings.Join([]string{"Hello", "World"}, " "): Joins the array of strings into a single string based on the delimiter



## Array and Slices

Array definition and initialization
```go
var userNames = [10]string{}

// or

userNames := [10]string{}
```
Here 10 is the size and string is the datatype and in the {} braces we pass
our comma seperated values

Without initialization
```go
var userNames [50]string
```

##### Array Operation functions

> len(userNames): Gives the length of the array

### Slices - It is an array with dynamic size

###### Declaration
```go
var userNames []string
```
###### Operation functions
* Append - append() - Adds value at the end of slice
* Length - len() - Get the length of the slice
```go
append(userName, "John Doe")
len(userName)
```


## Maps

```go
var userNames = map[string]string{}

// or

userNames := map[string]string{}
```

###### Operation functions

* Inserting values
```go
userNames["John"] = "John Doe"
```

* Deleting values
```go
delete(userNames, "John")
```

* Length of map
```go
len(userNames)
```

* Check if key exists
```go

if val, ok := userNames["John"]; ok {
    fmt.Printf("The value is %s\n", val)
}
```

* Iterate over map
```go

for key, value := range userNames {
    fmt.Printf("The key is %s and value is %s\n", key, value)
}
```

## Conditionals

* If else

```go
if condition {
    // code
} else {
    // code
}
```

* If else if

```go

if condition {
    // code
} else if condition {
    // code
} else {
    // code
}
```

* Switch case

```go

switch condition {
case condition:
    // code
case condition:

default:
    // code
}
```



## Loops
> In Go We have only for loop for all the different use cases

* Infinite Loop Syntax

```go
	count := 0
	for {
		fmt.Printf("The count is %d\n", count)
		count++
	}
```

* Looping through Slice/Array

```go
for index, value := range userNames {
    fmt.Printf("The index is %d and value is %s\n", index, value)
}
```

* Looping through Map

```go
for key, value := range userNames {
    fmt.Printf("The key is %d and value is %s\n", key, value)
}
```

* Looping through String

```go
for index, value := range userNames {
    fmt.Printf("The index is %d and value is %s\n", index, value)
}
```

* General For Loop

```go
for i := 0; i < 10; i++ {
    fmt.Printf("The index is %d\n", i)
}
```

* For Loop As a While Loop

```go

for condition {
    // code
}
```


## Pointers

```go
var age int = 23
var agePointer *int = &age
```

* Dereferencing a pointer

```go

fmt.Printf("The age is %d\n", *agePointer)
```

* Passing pointer to a function

```go

func changeAge(agePointer *int) {
    *agePointer = 24
}

changeAge(&age)
fmt.Printf("The age is %d\n", age)
```

* Creating a pointer to a struct

```go

type User struct {
    name string
    age int
    address string
}

var user User
var userPointer *User = &user
```

* Dereferencing a pointer to a struct

```go

fmt.Printf("The age is %d\n", (*userPointer).age)
```

* Dereferencing a pointer to a struct with shorthand

```go

fmt.Printf("The age is %d\n", userPointer.age)
```

* Creating a pointer to a struct with new keyword

```go

var userPointer *User = new(User)
```


## Functions

```go
func functionName() {
    // function body
}
```

* Function with parameters

```go

func functionName(param1 string, param2 int) {
    // function body
}
```

* Function with return type

```go
func functionName(param1 string, param2 int) int {
    // function body
    return 0
}
```

* Function with multiple return types

```go
func functionName(param1 string, param2 int) (int, string) {
    // function body
    return 0, "Hello"
}
```

* Function with named return types

```go

func functionName(param1 string, param2 int) (a int, b string) {
    // function body
    a = 0
    b = "Hello"
    return
}
```

* Function with variadic parameters

```go
func functionName(param1 string, param2 ...int) {
    // function body
}
```


* Function with anonymous function

```go

func functionName(param1 string, param2 int) {
    // function body
    func() {
        // anonymous function body
    }()
}
```

* Function with anonymous function and return type

```go

func functionName(param1 string, param2 int) int {
    // function body
    return func() int {
        // anonymous function body
        return 0
    }()
}
```


## types and structs

```go
type User struct {
    name string
    age int
    address string
}
```

* Creating a struct

```go
var user User
```

* Creating a struct with values

```go

user := User{
    name: "John Doe",
    age: 23,
    address: "New York",
}
```

* Creating a struct with values without keys

```go

user := User{
    "John Doe",
    23,
    "New York",
}
```

* Defining Function on Struct

```go

func (u User) getAge() int {
    return u.age
}
```

* Defining Function on Struct with pointer

```go

func (u *User) getAge() int {
    return u.age
}
```


## Interfaces

```go

type User interface {
    getAge() int
}
```

* Implementing Interface

```go

type Person struct {
    name string
    age int
    address string
}

func (p Person) getAge() int {
    return p.age
}
```

* Implementing Interface with pointer

```go

type Person struct {
    name string
    age int
    address string
}

func (p *Person) getAge() int {
    return p.age
}
```

* Implementing Interface with pointer

```go

type Person struct {
    name string
    age int
    address string
}

func (p *Person) getAge() int {
    return p.age
}
```

## Defer

* Defer is used to delay the execution of a function until the surrounding function returns

```go

func main() {
    defer fmt.Println("Hello")
    fmt.Println("World")
}
```

## Panic and Recover

* Panic is used to throw an error and stop the execution of the program

```go

func main() {
    panic("Error")
}
```

* Recover is used to recover from the panic and continue the execution of the program

```go

func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("Error")
}
```


## Concurrency

* Goroutines are lightweight threads managed by the Go runtime

```go

func main() {
    go func() {
        fmt.Println("Hello")
    }()
    fmt.Println("World")
}
```

* Channels are used to communicate between goroutines

```go

func main() {
    channel := make(chan string)
    go func() {
        channel <- "Hello"
    }()
    fmt.Println(<-channel)
}
```

* Buffered Channels are used to communicate between goroutines

```go

func main() {
    channel := make(chan string, 2)
    go func() {
        channel <- "Hello"
        channel <- "World"
    }()
    fmt.Println(<-channel)
    fmt.Println(<-channel)
}
```

* Select is used to select between different channels

```go

func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)
    go func() {
        channel1 <- "Hello"
    }()
    go func() {
        channel2 <- "World"
    }()
    select {
    case msg1 := <-channel1:
        fmt.Println(msg1)
    case msg2 := <-channel2:
        fmt.Println(msg2)
    }
}
```

* Select with default is used to select between different channels

```go

func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)
    go func() {
        channel1 <- "Hello"
    }()
    go func() {
        channel2 <- "World"
    }()
    select {
    case msg1 := <-channel1:
        fmt.Println(msg1)
    case msg2 := <-channel2:
        fmt.Println(msg2)
    default:
        fmt.Println("Default")
    }
}
```


## File Handling

* Creating a file

```go

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    file.WriteString("Hello World")
    file.Close()
}
```

* Reading a file

```go

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    // Open the file
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)

    // Read the file line by line
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

```

* Deleting a file

```go

func main() {
    err := os.Remove("test.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

* Creating a directory

```go

func main() {
    err := os.Mkdir("test", 0755)
    if err != nil {
        log.Fatal(err)
    }
}
```

* Deleting a directory

```go

func main() {
    err := os.Remove("test")
    if err != nil {
        log.Fatal(err)
    }
}
```

* Renaming a file

```go

func main() {
    err := os.Rename("test.txt", "test2.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

* Checking if file or directory exists

```go

func main() {
    if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
        log.Fatal(err)
    }
}
```


## HTTP Server

* Creating a HTTP Server

```go

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })
    http.ListenAndServe(":8080", nil)
}
```

* Creating a HTTP Server with templates

```go

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, "Hello World")
    })
    http.ListenAndServe(":8080", nil)
}
```

* Creating a HTTP Server with templates and static files

```go

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, "Hello World")
    })
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8080", nil)
}
```

* Creating a HTTP Server with templates and static files and custom handler

```go

type CustomHandler struct {
    http.Handler
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, "Hello World")
    })
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.Handle("/custom/", &CustomHandler{})
    http.ListenAndServe(":8080", nil)
}
```





