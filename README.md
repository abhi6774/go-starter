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
````


## Taking User Input

```go
var age int
fmt.Printf("Enter your age: ")
fmt.Scan(&age)
```
& is used to pass the reference or pointer address of a variable

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

#### Array Operation functions

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

## Loops
> In Go We have only for loop for all the different use cases


