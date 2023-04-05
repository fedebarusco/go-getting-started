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
3. Run the code: 
    ```console
    go run .
    ```