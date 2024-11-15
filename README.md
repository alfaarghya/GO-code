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
