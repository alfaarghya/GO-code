<div align="center">

<img src="./public/go.webp" height=100/>

</div>

## How to run codes in this repo

```bash
    git clone https://github.com/alfaarghya/GO-code.git

    cd <dir_name>

    go run .
```

## Day 1 - Say Hello

### Why GO?

Go (often called Golang) is a statically typed, compiled programming language developed by Google in 2009. It was created by Robert Griesemer, Rob Pike, and Ken Thompson to Go address common challenges in software engineering, such as efficiency, simplicity, and scalability, especially in large-scale systems.

#### Key Features of Go:

1. **Simplicity**: Go’s syntax is simple and minimalistic, designed to reduce complexity without compromising functionality. It has a clean structure and avoids features like inheritance and operator overloading, making it easy to read, write, and maintain code.
2. **Performance**: Go is a compiled language, so it offers near-C-like performance. It compiles quickly into native machine code, resulting in efficient execution without the need for a virtual machine. This makes it well-suited for high-performance applications and backend services.
3. **Concurrency**: Go provides built-in concurrency support through goroutines and channels. Goroutines are lightweight threads that allow you to run functions concurrently with minimal resource overhead, which is ideal for building scalable, concurrent applications and distributed systems.
4. **Garbage Collection**: Go includes an efficient garbage collector to manage memory, which helps reduce the complexity of memory management and lets developers focus on writing functional code without manual memory allocation or deallocation.
5. **Cross-Platform**: Go is designed to be highly portable. A single codebase can be compiled for multiple operating systems, like Windows, macOS, and Linux, without any changes, and it produces a single executable binary, simplifying deployment.
6. **Rich Standard Library**: Go has a powerful standard library with packages for web development, file handling, testing, and more, making it easy to accomplish common programming tasks without relying heavily on external libraries.

#### Common Use Cases for Go:

- **Backend Web Development**: Many high-performance web servers and microservices are built in Go, as it’s efficient and handles concurrent requests well. Popular web frameworks in Go include Gin and Echo.
- **Cloud-Native Applications**: Go is widely used in cloud-native environments. Docker, Kubernetes, and other cloud infrastructure tools are written in Go because it’s ideal for distributed systems and microservices.
- **Command-Line Tools**: Go’s ability to compile into a single binary makes it popular for building CLI tools. Its simplicity and efficiency make it suitable for automation tools and DevOps utilities.
- **Networking and Distributed Systems**: Go is used to build systems that require high concurrency and low latency, such as networking applications, API servers, and data-processing pipelines.

### **Get started**

#### Download & Install

    https://go.dev/doc/install

#### Let's start code

Get started with Hello, World.

1. Go to the desire directory & create a new directory as `hello` & open in VS code

   ```bash
   mkdir hello
   cd hello
   code .
   ```

2. Enable dependency tracking for your code

   When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

   In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be `github.com/mymodule`. If you plan to publish your module for others to use, the module path *must* be a location from which Go tools can download your module.

   For the purposes of practice, just use `example/hello`.

   ```bash
   go mod init practice/hello
   ```

3. Now create a new GO file, `hello.go`
4. Write a simple `hello world!` code

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

5. Explanation
   - `package main`: Every Go program is part of a package. The `main` package is where the Go runtime looks for the entry point of the program.
   - `import "fmt"`: This line imports the `fmt` package, which provides formatting functions like `Println`.
   - `func main()`: The `main` function is the entry point of the program, where execution starts.
   - `fmt.Println`: This function prints the message to the console.
6. Run Your Code

   ```bash
   go run hello.go
   ```

### Taking Input

```go
    package main

    import "fmt"

    func main() {
    	var name string //variable

    	fmt.Println("What is your name?")
    	fmt.Print("> ")
    	fmt.Scan(&name) //taking input from user

    	fmt.Println(">> Greetings, "+name)

    }
```

## Day 2 - Variables & Data Type

### Declare Variables

1. **With `var` keyword:**

   **syntax**

   ```go
   var *variable1 type = value //with type
   var variable2 = value //with out type*
   ```

   **📑 Note**

   1. can declare without assigning value.
   2. after assigning a value(suppose string), can’t re-assign with a different datatype value(int or float32 or bool).
   3. after assigning a value(suppose string), can re-assign with the same datatype value(string).
   4. `var` can be used inside a function or outside of a function, in other words as a local variable or global variable

2. **with `:=` sign:**

   **syntax**

   ```go
   variablename := value
   ```

   **📑 Note**

   1. can’t declare without assigning value.
   2. after assigning a value(suppose string), can’t re-assign with a different datatype value(int or float32 or bool).
   3. after assigning a value(suppose string), can re-assign with the same datatype value(string).
   4. `:=` can be used inside a function only, in other words as a local variable.

3. **with `const` keyword:**

   **syntax**

   ```go
   const variable1 type = value //with type
   const variable2 = value //with out type
   ```

   **📑 Note**

   1. can’t declare without assigning value.
   2. after assigning a value(suppose string), can’t re-assign with a different datatype value(int or float32 or bool).
   3. after assigning a value(suppose string), can’t re-assign with the same datatype value(string).
   4. `const` can be used inside a function or outside of a function, in other words as a local variable or global variable

### Data Types

1.  **numeric**

    a. **Integer** - this data types are used to store a whole number without decimals, like 10, -110, or 1503450. The integer data type has two categories:

    - **Signed integers** - can store both positive and negative values

      ```go
      var x int = 500
      var y int = -4500
      ```

      Go has five keywords/types of signed integers -

      | type    | size                      | range                                                                                       |
      | :------ | :------------------------ | :------------------------------------------------------------------------------------------ |
      | `int`   | can be 32 bits or 64 bits | 2147483648 to 2147483647 in 32 bit OR -9223372036854775808 to 9223372036854775807 in 64 bit |
      | `int8`  | 8 bits or 1 byte          | -128 to 127                                                                                 |
      | `int16` | 16 bits or 4 byte         | -32768 to 32767                                                                             |
      | `int32` | 32 bits or 4 byte         | -2147483648 to 2147483647                                                                   |
      | `int64` | 64 bits or 8 byte         | -9223372036854775808 to 9223372036854775807                                                 |

    * **Unsigned integers** - can only store non-negative values

      ```go
      var x uint = 500
      var y uint = 4500
      ```

      Go has five keywords/types of signed integers -

      | type     | size                      | range                                                            |
      | :------- | :------------------------ | :--------------------------------------------------------------- |
      | `uint`   | can be 32 bits or 64 bits | 0 to 4294967295 in 32 bit OR 0 to 18446744073709551615 in 64 bit |
      | `uint8`  | 8 bits or 1 byte          | 0 to 255                                                         |
      | `uint16` | 16 bits or 4 byte         | 0 to 65535                                                       |
      | `uint32` | 32 bits or 4 byte         | 0 to 4294967295                                                  |
      | `uint64` | 64 bits or 8 byte         | 0 to 18446744073709551615                                        |

    b. **float -** this data types are used to store positive and negative numbers with a decimal point, like 10.10, -2.53, or 4397.34587.

    Go float data type has two keywords -

    | type      | size              | range                   |
    | --------- | ----------------- | ----------------------- |
    | `float32` | 32 bits or 4 byte | -3.4e+38 to 3.4e+38.    |
    | `float64` | 64 bits or 8 byte | -1.7e+308 to +1.7e+308. |

2.  **string**

    The `string` data type is used to store a sequence of characters (text). String values must be surrounded by double quotes.

    ```go
    var txt1 string = "Hello World"
    ```

3.  **boolean**

    A boolean data type is declared with the `bool` keyword and can only take the values `true` or `false`. The default value of a boolean data type is `false`.

    ```go
    var bool1 bool = true
    var bool2 bool
    ```

## Day 3 - Operations

### Arithmetic operation

| Operator | Name           | Description                      | example |
| -------- | -------------- | -------------------------------- | ------- |
| +        | Addition       | adds values together             | x + y   |
| -        | Substraction   | Subtracts one value from another | x - y   |
| \*       | Multiplication | Multiplies two values            | x \* y  |
| /        | Division       | divides one value by another     | x / y   |
| %        | Modulus        | return the divison remainder     | x % y   |

### **Comparison Operators**

| Operator | Name               | example |
| -------- | ------------------ | ------- |
| ==       | Equal tp           | x == y  |
| ! =      | not equal          | x ! = y |
| >        | greater than       | x > y   |
| <        | less than          | x < y   |
| > =      | greater than euqal | x > = y |
| < =      | less than equal    | x < = y |

### Logical operator

| Operator | Name        | Description                                                | example             |
| -------- | ----------- | ---------------------------------------------------------- | ------------------- |
| &&       | Logical AND | returns true if both statement are true                    | x < 11 && x < 20    |
|          | Logical OR  | returns true if one of the statements is true              |
| !        | Logical NOT | reverse the result, returns true if false or false if true | !(x < 11 && x < 20) |

### Bitwise operators

| Operator | Name                 | Description                                                                                             | example |
| -------- | -------------------- | ------------------------------------------------------------------------------------------------------- | ------- |
| &        | AND                  | Sets each bit to 1 if both bits are 1                                                                   | x & y   |
|          | OR                   | Sets each bit to 1 if one of two bits is 1                                                              |
| ^        | XOR                  | Sets each bit to 1 if only one of two bits is 1                                                         | x ^ y   |
| <<       | Zero fill left shift | Shift left by pushing zeros in from the right                                                           | x << 3  |
| >>       | zero fll right shift | Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off | x >> 4  |

```
📌 NOTE : for any missing operator check the code, can't use OR(|) operator because markdown table consider it as a row end

```
