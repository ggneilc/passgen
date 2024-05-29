*Date: 2024-01-05 21:38*
*type:* #oneshot
*Category*: [[Projects]]

*Dictionary, Random, and Command Line execution*

## Problem Description

Create a golang command line function (or honestly we might just do it in bash) that when going to the command line you can type 'passgen "identifier"' and it will create a new secured password for you and assign it to the identifier (e.x. 'goog' for google accounts) and you can then call passgen -r "identifier" to retrieve the password at the identifier for copy and paste. 


## Tools

1. Written in Go
2. Maps (Pythong Dictionary)
3. String generation
4. File io
5. Command line execution

### Maps


> [!tip] Map Creation
> You create a map with `map[string]int`, however if you just do a basic variable assignment, then it points to a null address, so use the `make()` function to automatically allocate space for the map!

Maps are very intuitive to use, here are the basic features learned from this little project:

1. simple creation with `m := make(map[string]int)`
2. simple removal with `delete(m, key)`
3. simple addition with `m["key"] = value`

you can also first check if the map contains a key before creation with `value, ok := m[key]`, where `value` takes the value at the key *if it exists*, and 0 otherwise. `ok` is true *if it exists*, and false otherwise.

You can also use this shorthand expression for error handling in go:

```go
if password, ok := passwords[key]; ok {
	fmt.Println("Key exists!")
}
```

### Functions

You create a function with a basic `func my_function() {}`, but in order to return a value, you need to add a type after the function: `func generateString() string {}`.


### Random

You first need to import golangs installed `math/rand` package, and from there it needs to be seeded to avoid repeatable behavior, which is easy to seed with the time of the function call: 
```go
s1 := rand.NewSource(time.Now().UnixNano())
r1 := rand.New(s1)
```

from there you can either create new integers with `r1.Int()`, or a integer between a certain value with `r1.Intn(10)`.

### Command Line Executable

This is how all programs become globally executable: **simply move them to the bin (*binaries*) folder on your machine**.

In windows this is usually under `C:\Windows\System32`
In Linux this is under `usr/local/bin`

once the executable (binaries) are added to the bin (binaries) folder, it is recognized as a global command and can be executed from anywhere in the terminal!