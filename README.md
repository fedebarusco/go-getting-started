# Go: Getting Started
Introduction to Go programming language ([Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)).

Once the repository has been cloned, proceed with the following steps: 

1. From the project root directory, execute the following command to enable dependency tracking for your code by creating a go.mod file giving it the name of the module your code will be in. The name is the module's module path, for instance:
    ```console
    go mod init go-getting-started
    ```
2. Launch the command to locate and download the module(s) that contains the imported package(s): 
    ```console
    go mod tidy
    ```
    `go mod tidy` ensures that the `go.mod` file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to `go.sum` and removes unnecessary entries. For more information about the go mod tidy command, please refer to the [official documentation](https://go.dev/ref/mod#go-mod-tidy).  
3. Run the code: 
    ```console
    go run .
    ```
4. To build the program into binaries, use the go build command:
    ```console
    go build hello.go
    ```
    And execute the built binary directly:
    ```console
    ./hello
    ```

## Understanding the go.mod File

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a `go.mod` file that tracks the modules that provide those packages. That `go.mod` file stays with your code, including in your source code repository.

To enable dependency tracking for your code by creating a `go.mod` file, run the `go mod init [module-path]` command, giving it the name of the module your code will be in. The name is the module's module path.

In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path, see Managing dependencies. 

For more information about the go.mod file: [go-mod-init](https://go.dev/ref/mod#go-mod-init).